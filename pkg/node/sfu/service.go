package sfu

import (
	"encoding/json"
	"fmt"
	"io"
	"sync"

	log "github.com/pion/ion-log"
	ion_sfu_log "github.com/pion/ion-sfu/pkg/logger"
	"github.com/pion/ion-sfu/pkg/middlewares/datachannel"
	ion_sfu "github.com/pion/ion-sfu/pkg/sfu"
	error_code "github.com/pion/ion/pkg/error"
	"github.com/pion/ion/proto/rtc"
	"github.com/pion/webrtc/v3"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func init() {
	ion_sfu_log.SetGlobalOptions(ion_sfu_log.GlobalConfig{V: 1})
}

type SFUService struct {
	rtc.UnimplementedRTCServer
	sfu   *ion_sfu.SFU
	mutex sync.RWMutex
	sigs  map[string]rtc.RTC_SignalServer
}

func NewSFUService(conf ion_sfu.Config) *SFUService {
	s := &SFUService{
		sigs: make(map[string]rtc.RTC_SignalServer),
	}
	sfu := ion_sfu.NewSFU(conf)
	dc := sfu.NewDatachannel(ion_sfu.APIChannelLabel)
	dc.Use(datachannel.SubscriberAPI)
	s.sfu = sfu
	return s
}

func (s *SFUService) RegisterService(registrar grpc.ServiceRegistrar) {
	rtc.RegisterRTCServer(registrar, s)
}

func (s *SFUService) Close() {
	log.Infof("SFU service closed")
}

func (s *SFUService) BroadcastStreamEvent(event *rtc.StreamEvent) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	for _, sig := range s.sigs {
		sig.Send(&rtc.Signalling{
			Payload: &rtc.Signalling_StreamEvent{
				StreamEvent: event,
			},
		})
	}
}

func (s *SFUService) Signal(sigStream rtc.RTC_SignalServer) error {
	peer := ion_sfu.NewPeer(s.sfu)
	var streams []*rtc.Stream

	defer func() {
		if peer.Session() != nil {
			log.Infof("[S=>C] close: sid => %v, uid => %v", peer.Session().ID(), peer.ID())

			s.mutex.Lock()
			delete(s.sigs, peer.ID())
			s.mutex.Unlock()

			if len(streams) > 0 {
				event := &rtc.StreamEvent{
					State:   rtc.StreamEvent_REMOVE,
					Streams: streams,
				}
				s.BroadcastStreamEvent(event)
				log.Infof("broadcast stream event %v, state = REMOVE", streams)
			}
		}
	}()

	for {
		in, err := sigStream.Recv()

		if err != nil {
			peer.Close()

			if err == io.EOF {
				return nil
			}

			errStatus, _ := status.FromError(err)
			if errStatus.Code() == codes.Canceled {
				return nil
			}

			log.Errorf("%v signal error %d", fmt.Errorf(errStatus.Message()), errStatus.Code())
			return err
		}

		switch payload := in.Payload.(type) {
		case *rtc.Signalling_Join:
			sid := payload.Join.Sid
			uid := payload.Join.Uid
			log.Infof("[C=>S] join: sid => %v, uid => %v", sid, uid)

			// Notify user of new ice candidate
			peer.OnIceCandidate = func(candidate *webrtc.ICECandidateInit, target int) {
				log.Debugf("[S=>C] peer.OnIceCandidate: target = %v, candidate = %v", target, candidate.Candidate)
				bytes, err := json.Marshal(candidate)
				if err != nil {
					log.Errorf("OnIceCandidate error: %v", err)
				}
				err = sigStream.Send(&rtc.Signalling{
					Payload: &rtc.Signalling_Trickle{
						Trickle: &rtc.Trickle{
							Init:   string(bytes),
							Target: rtc.Target(target),
						},
					},
				})
				if err != nil {
					log.Errorf("OnIceCandidate send error: %v", err)
				}
			}

			// Notify user of new offer
			peer.OnOffer = func(o *webrtc.SessionDescription) {
				log.Debugf("[S=>C] peer.OnOffer: %v", o.SDP)
				err = sigStream.Send(&rtc.Signalling{
					Payload: &rtc.Signalling_Description{
						Description: &rtc.SessionDescription{
							Target: rtc.Target(rtc.Target_SUBSCRIBER),
							Sdp:    o.SDP,
							Type:   o.Type.String(),
						},
					},
				})
				if err != nil {
					log.Errorf("negotiation error: %v", err)
				}
			}

			joinConf := ion_sfu.JoinConfig{
				NoSubscribe:     false,
				NoPublish:       false,
				NoAutoSubscribe: true,
			}

			err = peer.Join(sid, uid, joinConf)
			if err != nil {
				switch err {
				case ion_sfu.ErrTransportExists:
					fallthrough
				case ion_sfu.ErrOfferIgnored:
					err = sigStream.Send(&rtc.Signalling{
						Payload: &rtc.Signalling_Error{
							Error: &rtc.Error{
								Code:   int32(error_code.InternalError),
								Reason: fmt.Sprintf("join error: %v", err),
							},
						},
					})
					if err != nil {
						log.Errorf("grpc send error: %v", err)
						return status.Errorf(codes.Internal, err.Error())
					}
				default:
					return status.Errorf(codes.Unknown, err.Error())
				}
			}

			peer.Publisher().OnPublisherTrack(func(track ion_sfu.PublisherTrack) {
				log.Debugf("peer.OnPublisherTrack: \nKind %v, \nUid: %v,  \nMsid: %v,\nTrackID: %v", track.Track.Kind(), uid, track.Track.Msid(), track.Track.ID())
			})

			sigStream.Send(&rtc.Signalling{
				Payload: &rtc.Signalling_Reply{
					Reply: &rtc.JoinReply{
						Success: true,
						Error:   nil,
					},
				},
			})

			streamMap := make(map[string]*rtc.Stream)
			for _, p := range peer.Session().Peers() {
				if peer.ID() != p.ID() {
					for _, pubTrack := range p.Publisher().PublisherTracks() {
						streamID := pubTrack.Track.StreamID()
						stream, found := streamMap[streamID]
						if !found {
							stream = &rtc.Stream{
								Uid:  uid,
								Msid: streamID,
							}
							streamMap[streamID] = stream
						}
						stream.Tracks = append(stream.Tracks, &rtc.Track{
							Id:    pubTrack.Track.ID(),
							Kind:  pubTrack.Track.Kind().String(),
							Muted: false,
							Rid:   pubTrack.Track.RID(),
						})
					}
				}
			}

			var otherStreams []*rtc.Stream
			for _, stream := range streamMap {
				otherStreams = append(otherStreams, stream)
			}

			event := &rtc.StreamEvent{
				State:   rtc.StreamEvent_ADD,
				Streams: otherStreams,
			}

			sigStream.Send(&rtc.Signalling{
				Payload: &rtc.Signalling_StreamEvent{
					StreamEvent: event,
				},
			})

			//TODO: Return error when the room is full, or locked, or permission denied

			s.mutex.Lock()
			s.sigs[peer.ID()] = sigStream
			s.mutex.Unlock()

		case *rtc.Signalling_Description:
			desc := webrtc.SessionDescription{
				SDP:  payload.Description.Sdp,
				Type: webrtc.NewSDPType(payload.Description.Type),
			}
			var err error = nil
			switch desc.Type {
			case webrtc.SDPTypeOffer:
				log.Debugf("[C=>S] description: offer %v", desc.SDP)
				answer, err := peer.Answer(desc)
				if err != nil {
					return status.Errorf(codes.Internal, fmt.Sprintf("answer error: %v", err))
				}

				// send answer
				log.Debugf("[S=>C] description: answer %v", answer.SDP)

				err = sigStream.Send(&rtc.Signalling{
					Payload: &rtc.Signalling_Description{
						Description: &rtc.SessionDescription{
							Target: rtc.Target(rtc.Target_PUBLISHER),
							Sdp:    answer.SDP,
							Type:   answer.Type.String(),
						},
					},
				})

				if err != nil {
					log.Errorf("grpc send error: %v", err)
					return status.Errorf(codes.Internal, err.Error())
				}

				newStreams, err := ParseSDP(peer.ID(), desc.SDP)
				if err != nil {
					log.Errorf("util.ParseSDP error: %v", err)
				}

				if len(newStreams) > 0 {
					event := &rtc.StreamEvent{
						Streams: newStreams,
						State:   rtc.StreamEvent_ADD,
					}
					streams = newStreams
					log.Infof("broadcast stream event %v, state = ADD", streams)
					s.BroadcastStreamEvent(event)
				}

			case webrtc.SDPTypeAnswer:
				log.Debugf("[C=>S] description: answer %v", desc.SDP)
				err = peer.SetRemoteDescription(desc)
			}

			if err != nil {
				switch err {
				case ion_sfu.ErrNoTransportEstablished:
					err = sigStream.Send(&rtc.Signalling{
						Payload: &rtc.Signalling_Error{
							Error: &rtc.Error{
								Code:   int32(error_code.UnsupportedMediaType),
								Reason: fmt.Sprintf("set remote description error: %v", err),
							},
						},
					})
					if err != nil {
						log.Errorf("grpc send error: %v", err)
						return status.Errorf(codes.Internal, err.Error())
					}
				default:
					return status.Errorf(codes.Unknown, err.Error())
				}
			}

		case *rtc.Signalling_Trickle:
			var candidate webrtc.ICECandidateInit
			err := json.Unmarshal([]byte(payload.Trickle.Init), &candidate)
			if err != nil {
				log.Errorf("error parsing ice candidate, error -> %v", err)
				err = sigStream.Send(&rtc.Signalling{
					Payload: &rtc.Signalling_Error{
						Error: &rtc.Error{
							Code:   int32(error_code.InternalError),
							Reason: fmt.Sprintf("unmarshal ice candidate error:  %v", err),
						},
					},
				})
				if err != nil {
					log.Errorf("grpc send error: %v", err)
					return status.Errorf(codes.Internal, err.Error())
				}
				continue
			}
			log.Debugf("[C=>S] trickle: target %v, candidate %v", int(payload.Trickle.Target), candidate.Candidate)
			err = peer.Trickle(candidate, int(payload.Trickle.Target))
			if err != nil {
				switch err {
				case ion_sfu.ErrNoTransportEstablished:
					log.Errorf("peer hasn't joined, error -> %v", err)
					err = sigStream.Send(&rtc.Signalling{
						Payload: &rtc.Signalling_Error{
							Error: &rtc.Error{
								Code:   int32(error_code.InternalError),
								Reason: fmt.Sprintf("trickle error:  %v", err),
							},
						},
					})
					if err != nil {
						log.Errorf("grpc send error: %v", err)
						return status.Errorf(codes.Internal, err.Error())
					}
				default:
					return status.Errorf(codes.Unknown, fmt.Sprintf("negotiate error: %v", err))
				}
			}

		case *rtc.Signalling_UpdateSettings:
			switch payload.UpdateSettings.Command.(type) {
			case *rtc.UpdateSettings_Subcription:
				subscription := payload.UpdateSettings.GetSubcription()
				subscribe := subscription.GetSubscribe()
				needNegotiate := false
				for _, trackId := range subscription.TrackIds {
					if subscribe {
						// Add down tracks
						for _, p := range peer.Session().Peers() {
							if p.ID() != peer.ID() {
								for _, track := range p.Publisher().PublisherTracks() {
									if track.Receiver.TrackID() == trackId {
										log.Debugf("Add RemoteTrack: %v to peer %v", trackId, peer.ID())
										peer.Publisher().GetRouter().AddDownTrack(peer.Subscriber(), track.Receiver)
										needNegotiate = true
									}
								}
							}
						}
					} else {
						// Remove down tracks
						for streamID, downTracks := range peer.Subscriber().DownTracks() {
							for _, downTrack := range downTracks {
								if downTrack != nil && downTrack.ID() == trackId {
									peer.Subscriber().RemoveDownTrack(streamID, downTrack)
									downTrack.Stop()
									needNegotiate = true
								}
							}
						}
					}
				}
				if needNegotiate {
					peer.Subscriber().Negotiate()
				}
			}
		}
	}
}