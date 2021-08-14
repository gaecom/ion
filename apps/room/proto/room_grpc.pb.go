// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// RoomServiceClient is the client API for RoomService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RoomServiceClient interface {
	// Manager API
	// Room API
	CreateRoom(ctx context.Context, in *CreateRoomRequest, opts ...grpc.CallOption) (*CreateRoomReply, error)
	UpdateRoom(ctx context.Context, in *UpdateRoomRequest, opts ...grpc.CallOption) (*UpdateRoomReply, error)
	EndRoom(ctx context.Context, in *EndRoomRequest, opts ...grpc.CallOption) (*EndRoomReply, error)
	GetRooms(ctx context.Context, in *GetRoomsRequest, opts ...grpc.CallOption) (*GetRoomsReply, error)
	// Peer API
	AddPeer(ctx context.Context, in *AddPeerRequest, opts ...grpc.CallOption) (*AddPeerReply, error)
	UpdatePeer(ctx context.Context, in *UpdatePeerRequest, opts ...grpc.CallOption) (*UpdatePeerReply, error)
	RemovePeer(ctx context.Context, in *RemovePeerRequest, opts ...grpc.CallOption) (*RemovePeerReply, error)
	GetPeers(ctx context.Context, in *GetPeersRequest, opts ...grpc.CallOption) (*GetPeersReply, error)
	SetImportance(ctx context.Context, in *SetImportanceRequest, opts ...grpc.CallOption) (*SetImportanceReply, error)
}

type roomServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewRoomServiceClient(cc grpc.ClientConnInterface) RoomServiceClient {
	return &roomServiceClient{cc}
}

func (c *roomServiceClient) CreateRoom(ctx context.Context, in *CreateRoomRequest, opts ...grpc.CallOption) (*CreateRoomReply, error) {
	out := new(CreateRoomReply)
	err := c.cc.Invoke(ctx, "/room.RoomService/CreateRoom", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *roomServiceClient) UpdateRoom(ctx context.Context, in *UpdateRoomRequest, opts ...grpc.CallOption) (*UpdateRoomReply, error) {
	out := new(UpdateRoomReply)
	err := c.cc.Invoke(ctx, "/room.RoomService/UpdateRoom", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *roomServiceClient) EndRoom(ctx context.Context, in *EndRoomRequest, opts ...grpc.CallOption) (*EndRoomReply, error) {
	out := new(EndRoomReply)
	err := c.cc.Invoke(ctx, "/room.RoomService/EndRoom", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *roomServiceClient) GetRooms(ctx context.Context, in *GetRoomsRequest, opts ...grpc.CallOption) (*GetRoomsReply, error) {
	out := new(GetRoomsReply)
	err := c.cc.Invoke(ctx, "/room.RoomService/GetRooms", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *roomServiceClient) AddPeer(ctx context.Context, in *AddPeerRequest, opts ...grpc.CallOption) (*AddPeerReply, error) {
	out := new(AddPeerReply)
	err := c.cc.Invoke(ctx, "/room.RoomService/AddPeer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *roomServiceClient) UpdatePeer(ctx context.Context, in *UpdatePeerRequest, opts ...grpc.CallOption) (*UpdatePeerReply, error) {
	out := new(UpdatePeerReply)
	err := c.cc.Invoke(ctx, "/room.RoomService/UpdatePeer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *roomServiceClient) RemovePeer(ctx context.Context, in *RemovePeerRequest, opts ...grpc.CallOption) (*RemovePeerReply, error) {
	out := new(RemovePeerReply)
	err := c.cc.Invoke(ctx, "/room.RoomService/RemovePeer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *roomServiceClient) GetPeers(ctx context.Context, in *GetPeersRequest, opts ...grpc.CallOption) (*GetPeersReply, error) {
	out := new(GetPeersReply)
	err := c.cc.Invoke(ctx, "/room.RoomService/GetPeers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *roomServiceClient) SetImportance(ctx context.Context, in *SetImportanceRequest, opts ...grpc.CallOption) (*SetImportanceReply, error) {
	out := new(SetImportanceReply)
	err := c.cc.Invoke(ctx, "/room.RoomService/SetImportance", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RoomServiceServer is the server API for RoomService service.
// All implementations must embed UnimplementedRoomServiceServer
// for forward compatibility
type RoomServiceServer interface {
	// Manager API
	// Room API
	CreateRoom(context.Context, *CreateRoomRequest) (*CreateRoomReply, error)
	UpdateRoom(context.Context, *UpdateRoomRequest) (*UpdateRoomReply, error)
	EndRoom(context.Context, *EndRoomRequest) (*EndRoomReply, error)
	GetRooms(context.Context, *GetRoomsRequest) (*GetRoomsReply, error)
	// Peer API
	AddPeer(context.Context, *AddPeerRequest) (*AddPeerReply, error)
	UpdatePeer(context.Context, *UpdatePeerRequest) (*UpdatePeerReply, error)
	RemovePeer(context.Context, *RemovePeerRequest) (*RemovePeerReply, error)
	GetPeers(context.Context, *GetPeersRequest) (*GetPeersReply, error)
	SetImportance(context.Context, *SetImportanceRequest) (*SetImportanceReply, error)
	mustEmbedUnimplementedRoomServiceServer()
}

// UnimplementedRoomServiceServer must be embedded to have forward compatible implementations.
type UnimplementedRoomServiceServer struct {
}

func (UnimplementedRoomServiceServer) CreateRoom(context.Context, *CreateRoomRequest) (*CreateRoomReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateRoom not implemented")
}
func (UnimplementedRoomServiceServer) UpdateRoom(context.Context, *UpdateRoomRequest) (*UpdateRoomReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateRoom not implemented")
}
func (UnimplementedRoomServiceServer) EndRoom(context.Context, *EndRoomRequest) (*EndRoomReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EndRoom not implemented")
}
func (UnimplementedRoomServiceServer) GetRooms(context.Context, *GetRoomsRequest) (*GetRoomsReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRooms not implemented")
}
func (UnimplementedRoomServiceServer) AddPeer(context.Context, *AddPeerRequest) (*AddPeerReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddPeer not implemented")
}
func (UnimplementedRoomServiceServer) UpdatePeer(context.Context, *UpdatePeerRequest) (*UpdatePeerReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdatePeer not implemented")
}
func (UnimplementedRoomServiceServer) RemovePeer(context.Context, *RemovePeerRequest) (*RemovePeerReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemovePeer not implemented")
}
func (UnimplementedRoomServiceServer) GetPeers(context.Context, *GetPeersRequest) (*GetPeersReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPeers not implemented")
}
func (UnimplementedRoomServiceServer) SetImportance(context.Context, *SetImportanceRequest) (*SetImportanceReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetImportance not implemented")
}
func (UnimplementedRoomServiceServer) mustEmbedUnimplementedRoomServiceServer() {}

// UnsafeRoomServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RoomServiceServer will
// result in compilation errors.
type UnsafeRoomServiceServer interface {
	mustEmbedUnimplementedRoomServiceServer()
}

func RegisterRoomServiceServer(s grpc.ServiceRegistrar, srv RoomServiceServer) {
	s.RegisterService(&RoomService_ServiceDesc, srv)
}

func _RoomService_CreateRoom_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRoomRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoomServiceServer).CreateRoom(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/room.RoomService/CreateRoom",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoomServiceServer).CreateRoom(ctx, req.(*CreateRoomRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RoomService_UpdateRoom_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateRoomRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoomServiceServer).UpdateRoom(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/room.RoomService/UpdateRoom",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoomServiceServer).UpdateRoom(ctx, req.(*UpdateRoomRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RoomService_EndRoom_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EndRoomRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoomServiceServer).EndRoom(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/room.RoomService/EndRoom",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoomServiceServer).EndRoom(ctx, req.(*EndRoomRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RoomService_GetRooms_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRoomsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoomServiceServer).GetRooms(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/room.RoomService/GetRooms",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoomServiceServer).GetRooms(ctx, req.(*GetRoomsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RoomService_AddPeer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddPeerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoomServiceServer).AddPeer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/room.RoomService/AddPeer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoomServiceServer).AddPeer(ctx, req.(*AddPeerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RoomService_UpdatePeer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdatePeerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoomServiceServer).UpdatePeer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/room.RoomService/UpdatePeer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoomServiceServer).UpdatePeer(ctx, req.(*UpdatePeerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RoomService_RemovePeer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemovePeerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoomServiceServer).RemovePeer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/room.RoomService/RemovePeer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoomServiceServer).RemovePeer(ctx, req.(*RemovePeerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RoomService_GetPeers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPeersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoomServiceServer).GetPeers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/room.RoomService/GetPeers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoomServiceServer).GetPeers(ctx, req.(*GetPeersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RoomService_SetImportance_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetImportanceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoomServiceServer).SetImportance(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/room.RoomService/SetImportance",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoomServiceServer).SetImportance(ctx, req.(*SetImportanceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// RoomService_ServiceDesc is the grpc.ServiceDesc for RoomService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RoomService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "room.RoomService",
	HandlerType: (*RoomServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateRoom",
			Handler:    _RoomService_CreateRoom_Handler,
		},
		{
			MethodName: "UpdateRoom",
			Handler:    _RoomService_UpdateRoom_Handler,
		},
		{
			MethodName: "EndRoom",
			Handler:    _RoomService_EndRoom_Handler,
		},
		{
			MethodName: "GetRooms",
			Handler:    _RoomService_GetRooms_Handler,
		},
		{
			MethodName: "AddPeer",
			Handler:    _RoomService_AddPeer_Handler,
		},
		{
			MethodName: "UpdatePeer",
			Handler:    _RoomService_UpdatePeer_Handler,
		},
		{
			MethodName: "RemovePeer",
			Handler:    _RoomService_RemovePeer_Handler,
		},
		{
			MethodName: "GetPeers",
			Handler:    _RoomService_GetPeers_Handler,
		},
		{
			MethodName: "SetImportance",
			Handler:    _RoomService_SetImportance_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "apps/room/proto/room.proto",
}

// RoomSignalClient is the client API for RoomSignal service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RoomSignalClient interface {
	// Signal
	Signal(ctx context.Context, opts ...grpc.CallOption) (RoomSignal_SignalClient, error)
}

type roomSignalClient struct {
	cc grpc.ClientConnInterface
}

func NewRoomSignalClient(cc grpc.ClientConnInterface) RoomSignalClient {
	return &roomSignalClient{cc}
}

func (c *roomSignalClient) Signal(ctx context.Context, opts ...grpc.CallOption) (RoomSignal_SignalClient, error) {
	stream, err := c.cc.NewStream(ctx, &RoomSignal_ServiceDesc.Streams[0], "/room.RoomSignal/Signal", opts...)
	if err != nil {
		return nil, err
	}
	x := &roomSignalSignalClient{stream}
	return x, nil
}

type RoomSignal_SignalClient interface {
	Send(*Request) error
	Recv() (*Reply, error)
	grpc.ClientStream
}

type roomSignalSignalClient struct {
	grpc.ClientStream
}

func (x *roomSignalSignalClient) Send(m *Request) error {
	return x.ClientStream.SendMsg(m)
}

func (x *roomSignalSignalClient) Recv() (*Reply, error) {
	m := new(Reply)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// RoomSignalServer is the server API for RoomSignal service.
// All implementations must embed UnimplementedRoomSignalServer
// for forward compatibility
type RoomSignalServer interface {
	// Signal
	Signal(RoomSignal_SignalServer) error
	mustEmbedUnimplementedRoomSignalServer()
}

// UnimplementedRoomSignalServer must be embedded to have forward compatible implementations.
type UnimplementedRoomSignalServer struct {
}

func (UnimplementedRoomSignalServer) Signal(RoomSignal_SignalServer) error {
	return status.Errorf(codes.Unimplemented, "method Signal not implemented")
}
func (UnimplementedRoomSignalServer) mustEmbedUnimplementedRoomSignalServer() {}

// UnsafeRoomSignalServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RoomSignalServer will
// result in compilation errors.
type UnsafeRoomSignalServer interface {
	mustEmbedUnimplementedRoomSignalServer()
}

func RegisterRoomSignalServer(s grpc.ServiceRegistrar, srv RoomSignalServer) {
	s.RegisterService(&RoomSignal_ServiceDesc, srv)
}

func _RoomSignal_Signal_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(RoomSignalServer).Signal(&roomSignalSignalServer{stream})
}

type RoomSignal_SignalServer interface {
	Send(*Reply) error
	Recv() (*Request, error)
	grpc.ServerStream
}

type roomSignalSignalServer struct {
	grpc.ServerStream
}

func (x *roomSignalSignalServer) Send(m *Reply) error {
	return x.ServerStream.SendMsg(m)
}

func (x *roomSignalSignalServer) Recv() (*Request, error) {
	m := new(Request)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// RoomSignal_ServiceDesc is the grpc.ServiceDesc for RoomSignal service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RoomSignal_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "room.RoomSignal",
	HandlerType: (*RoomSignalServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Signal",
			Handler:       _RoomSignal_Signal_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "apps/room/proto/room.proto",
}