// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.20.0
// source: connection/connectionService.proto

package connection

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

// ConnectionServiceClient is the client API for ConnectionService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ConnectionServiceClient interface {
	NewUserConnection(ctx context.Context, in *UserConnectionRequest, opts ...grpc.CallOption) (*UserConnectionResponse, error)
	ApproveConnection(ctx context.Context, in *UserConnectionRequest, opts ...grpc.CallOption) (*UserConnectionResponse, error)
	GetConnection(ctx context.Context, in *Connection, opts ...grpc.CallOption) (*Connection, error)
	ApproveAllConnection(ctx context.Context, in *UserIdRequest, opts ...grpc.CallOption) (*EmptyRequest, error)
	RejectConnection(ctx context.Context, in *UserConnectionRequest, opts ...grpc.CallOption) (*UserConnectionResponse, error)
	DeleteConnection(ctx context.Context, in *Connection, opts ...grpc.CallOption) (*UserConnectionResponse, error)
	GetAllConnections(ctx context.Context, in *UserIdRequest, opts ...grpc.CallOption) (*AllConnectionResponse, error)
	GetFollowings(ctx context.Context, in *UserIdRequest, opts ...grpc.CallOption) (*AllConnectionResponse, error)
	GetFollowers(ctx context.Context, in *UserIdRequest, opts ...grpc.CallOption) (*AllConnectionResponse, error)
	GetAllRequestConnectionsByUserId(ctx context.Context, in *UserIdRequest, opts ...grpc.CallOption) (*AllConnectionResponse, error)
	GetAllPendingConnectionsByUserId(ctx context.Context, in *UserIdRequest, opts ...grpc.CallOption) (*AllConnectionResponse, error)
	BlockUser(ctx context.Context, in *BlockUserRequest, opts ...grpc.CallOption) (*EmptyRequest, error)
	UnblockUser(ctx context.Context, in *BlockUserRequest, opts ...grpc.CallOption) (*EmptyRequest, error)
	IsBlocked(ctx context.Context, in *Block, opts ...grpc.CallOption) (*IsBlockedResponse, error)
	IsBlockedAny(ctx context.Context, in *Block, opts ...grpc.CallOption) (*IsBlockedResponse, error)
	Blocked(ctx context.Context, in *UserIdRequest, opts ...grpc.CallOption) (*BlockedResponse, error)
	BlockedBy(ctx context.Context, in *UserIdRequest, opts ...grpc.CallOption) (*BlockedResponse, error)
	// Blocked + BlockedBy
	BlockedAny(ctx context.Context, in *UserIdRequest, opts ...grpc.CallOption) (*BlockedResponse, error)
	ChangeMessageNotification(ctx context.Context, in *UserConnectionRequest, opts ...grpc.CallOption) (*UserConnectionResponse, error)
	ChangePostNotification(ctx context.Context, in *UserConnectionRequest, opts ...grpc.CallOption) (*UserConnectionResponse, error)
	ChangeCommentNotification(ctx context.Context, in *UserConnectionRequest, opts ...grpc.CallOption) (*UserConnectionResponse, error)
}

type connectionServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewConnectionServiceClient(cc grpc.ClientConnInterface) ConnectionServiceClient {
	return &connectionServiceClient{cc}
}

func (c *connectionServiceClient) NewUserConnection(ctx context.Context, in *UserConnectionRequest, opts ...grpc.CallOption) (*UserConnectionResponse, error) {
	out := new(UserConnectionResponse)
	err := c.cc.Invoke(ctx, "/connection.ConnectionService/NewUserConnection", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *connectionServiceClient) ApproveConnection(ctx context.Context, in *UserConnectionRequest, opts ...grpc.CallOption) (*UserConnectionResponse, error) {
	out := new(UserConnectionResponse)
	err := c.cc.Invoke(ctx, "/connection.ConnectionService/ApproveConnection", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *connectionServiceClient) GetConnection(ctx context.Context, in *Connection, opts ...grpc.CallOption) (*Connection, error) {
	out := new(Connection)
	err := c.cc.Invoke(ctx, "/connection.ConnectionService/GetConnection", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *connectionServiceClient) ApproveAllConnection(ctx context.Context, in *UserIdRequest, opts ...grpc.CallOption) (*EmptyRequest, error) {
	out := new(EmptyRequest)
	err := c.cc.Invoke(ctx, "/connection.ConnectionService/ApproveAllConnection", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *connectionServiceClient) RejectConnection(ctx context.Context, in *UserConnectionRequest, opts ...grpc.CallOption) (*UserConnectionResponse, error) {
	out := new(UserConnectionResponse)
	err := c.cc.Invoke(ctx, "/connection.ConnectionService/RejectConnection", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *connectionServiceClient) DeleteConnection(ctx context.Context, in *Connection, opts ...grpc.CallOption) (*UserConnectionResponse, error) {
	out := new(UserConnectionResponse)
	err := c.cc.Invoke(ctx, "/connection.ConnectionService/DeleteConnection", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *connectionServiceClient) GetAllConnections(ctx context.Context, in *UserIdRequest, opts ...grpc.CallOption) (*AllConnectionResponse, error) {
	out := new(AllConnectionResponse)
	err := c.cc.Invoke(ctx, "/connection.ConnectionService/GetAllConnections", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *connectionServiceClient) GetFollowings(ctx context.Context, in *UserIdRequest, opts ...grpc.CallOption) (*AllConnectionResponse, error) {
	out := new(AllConnectionResponse)
	err := c.cc.Invoke(ctx, "/connection.ConnectionService/GetFollowings", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *connectionServiceClient) GetFollowers(ctx context.Context, in *UserIdRequest, opts ...grpc.CallOption) (*AllConnectionResponse, error) {
	out := new(AllConnectionResponse)
	err := c.cc.Invoke(ctx, "/connection.ConnectionService/GetFollowers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *connectionServiceClient) GetAllRequestConnectionsByUserId(ctx context.Context, in *UserIdRequest, opts ...grpc.CallOption) (*AllConnectionResponse, error) {
	out := new(AllConnectionResponse)
	err := c.cc.Invoke(ctx, "/connection.ConnectionService/GetAllRequestConnectionsByUserId", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *connectionServiceClient) GetAllPendingConnectionsByUserId(ctx context.Context, in *UserIdRequest, opts ...grpc.CallOption) (*AllConnectionResponse, error) {
	out := new(AllConnectionResponse)
	err := c.cc.Invoke(ctx, "/connection.ConnectionService/GetAllPendingConnectionsByUserId", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *connectionServiceClient) BlockUser(ctx context.Context, in *BlockUserRequest, opts ...grpc.CallOption) (*EmptyRequest, error) {
	out := new(EmptyRequest)
	err := c.cc.Invoke(ctx, "/connection.ConnectionService/BlockUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *connectionServiceClient) UnblockUser(ctx context.Context, in *BlockUserRequest, opts ...grpc.CallOption) (*EmptyRequest, error) {
	out := new(EmptyRequest)
	err := c.cc.Invoke(ctx, "/connection.ConnectionService/UnblockUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *connectionServiceClient) IsBlocked(ctx context.Context, in *Block, opts ...grpc.CallOption) (*IsBlockedResponse, error) {
	out := new(IsBlockedResponse)
	err := c.cc.Invoke(ctx, "/connection.ConnectionService/IsBlocked", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *connectionServiceClient) IsBlockedAny(ctx context.Context, in *Block, opts ...grpc.CallOption) (*IsBlockedResponse, error) {
	out := new(IsBlockedResponse)
	err := c.cc.Invoke(ctx, "/connection.ConnectionService/IsBlockedAny", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *connectionServiceClient) Blocked(ctx context.Context, in *UserIdRequest, opts ...grpc.CallOption) (*BlockedResponse, error) {
	out := new(BlockedResponse)
	err := c.cc.Invoke(ctx, "/connection.ConnectionService/Blocked", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *connectionServiceClient) BlockedBy(ctx context.Context, in *UserIdRequest, opts ...grpc.CallOption) (*BlockedResponse, error) {
	out := new(BlockedResponse)
	err := c.cc.Invoke(ctx, "/connection.ConnectionService/BlockedBy", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *connectionServiceClient) BlockedAny(ctx context.Context, in *UserIdRequest, opts ...grpc.CallOption) (*BlockedResponse, error) {
	out := new(BlockedResponse)
	err := c.cc.Invoke(ctx, "/connection.ConnectionService/BlockedAny", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *connectionServiceClient) ChangeMessageNotification(ctx context.Context, in *UserConnectionRequest, opts ...grpc.CallOption) (*UserConnectionResponse, error) {
	out := new(UserConnectionResponse)
	err := c.cc.Invoke(ctx, "/connection.ConnectionService/ChangeMessageNotification", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *connectionServiceClient) ChangePostNotification(ctx context.Context, in *UserConnectionRequest, opts ...grpc.CallOption) (*UserConnectionResponse, error) {
	out := new(UserConnectionResponse)
	err := c.cc.Invoke(ctx, "/connection.ConnectionService/ChangePostNotification", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *connectionServiceClient) ChangeCommentNotification(ctx context.Context, in *UserConnectionRequest, opts ...grpc.CallOption) (*UserConnectionResponse, error) {
	out := new(UserConnectionResponse)
	err := c.cc.Invoke(ctx, "/connection.ConnectionService/ChangeCommentNotification", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ConnectionServiceServer is the server API for ConnectionService service.
// All implementations must embed UnimplementedConnectionServiceServer
// for forward compatibility
type ConnectionServiceServer interface {
	NewUserConnection(context.Context, *UserConnectionRequest) (*UserConnectionResponse, error)
	ApproveConnection(context.Context, *UserConnectionRequest) (*UserConnectionResponse, error)
	GetConnection(context.Context, *Connection) (*Connection, error)
	ApproveAllConnection(context.Context, *UserIdRequest) (*EmptyRequest, error)
	RejectConnection(context.Context, *UserConnectionRequest) (*UserConnectionResponse, error)
	DeleteConnection(context.Context, *Connection) (*UserConnectionResponse, error)
	GetAllConnections(context.Context, *UserIdRequest) (*AllConnectionResponse, error)
	GetFollowings(context.Context, *UserIdRequest) (*AllConnectionResponse, error)
	GetFollowers(context.Context, *UserIdRequest) (*AllConnectionResponse, error)
	GetAllRequestConnectionsByUserId(context.Context, *UserIdRequest) (*AllConnectionResponse, error)
	GetAllPendingConnectionsByUserId(context.Context, *UserIdRequest) (*AllConnectionResponse, error)
	BlockUser(context.Context, *BlockUserRequest) (*EmptyRequest, error)
	UnblockUser(context.Context, *BlockUserRequest) (*EmptyRequest, error)
	IsBlocked(context.Context, *Block) (*IsBlockedResponse, error)
	IsBlockedAny(context.Context, *Block) (*IsBlockedResponse, error)
	Blocked(context.Context, *UserIdRequest) (*BlockedResponse, error)
	BlockedBy(context.Context, *UserIdRequest) (*BlockedResponse, error)
	// Blocked + BlockedBy
	BlockedAny(context.Context, *UserIdRequest) (*BlockedResponse, error)
	ChangeMessageNotification(context.Context, *UserConnectionRequest) (*UserConnectionResponse, error)
	ChangePostNotification(context.Context, *UserConnectionRequest) (*UserConnectionResponse, error)
	ChangeCommentNotification(context.Context, *UserConnectionRequest) (*UserConnectionResponse, error)
	mustEmbedUnimplementedConnectionServiceServer()
}

// UnimplementedConnectionServiceServer must be embedded to have forward compatible implementations.
type UnimplementedConnectionServiceServer struct {
}

func (UnimplementedConnectionServiceServer) NewUserConnection(context.Context, *UserConnectionRequest) (*UserConnectionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NewUserConnection not implemented")
}
func (UnimplementedConnectionServiceServer) ApproveConnection(context.Context, *UserConnectionRequest) (*UserConnectionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ApproveConnection not implemented")
}
func (UnimplementedConnectionServiceServer) GetConnection(context.Context, *Connection) (*Connection, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetConnection not implemented")
}
func (UnimplementedConnectionServiceServer) ApproveAllConnection(context.Context, *UserIdRequest) (*EmptyRequest, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ApproveAllConnection not implemented")
}
func (UnimplementedConnectionServiceServer) RejectConnection(context.Context, *UserConnectionRequest) (*UserConnectionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RejectConnection not implemented")
}
func (UnimplementedConnectionServiceServer) DeleteConnection(context.Context, *Connection) (*UserConnectionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteConnection not implemented")
}
func (UnimplementedConnectionServiceServer) GetAllConnections(context.Context, *UserIdRequest) (*AllConnectionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllConnections not implemented")
}
func (UnimplementedConnectionServiceServer) GetFollowings(context.Context, *UserIdRequest) (*AllConnectionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFollowings not implemented")
}
func (UnimplementedConnectionServiceServer) GetFollowers(context.Context, *UserIdRequest) (*AllConnectionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFollowers not implemented")
}
func (UnimplementedConnectionServiceServer) GetAllRequestConnectionsByUserId(context.Context, *UserIdRequest) (*AllConnectionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllRequestConnectionsByUserId not implemented")
}
func (UnimplementedConnectionServiceServer) GetAllPendingConnectionsByUserId(context.Context, *UserIdRequest) (*AllConnectionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllPendingConnectionsByUserId not implemented")
}
func (UnimplementedConnectionServiceServer) BlockUser(context.Context, *BlockUserRequest) (*EmptyRequest, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BlockUser not implemented")
}
func (UnimplementedConnectionServiceServer) UnblockUser(context.Context, *BlockUserRequest) (*EmptyRequest, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UnblockUser not implemented")
}
func (UnimplementedConnectionServiceServer) IsBlocked(context.Context, *Block) (*IsBlockedResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IsBlocked not implemented")
}
func (UnimplementedConnectionServiceServer) IsBlockedAny(context.Context, *Block) (*IsBlockedResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IsBlockedAny not implemented")
}
func (UnimplementedConnectionServiceServer) Blocked(context.Context, *UserIdRequest) (*BlockedResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Blocked not implemented")
}
func (UnimplementedConnectionServiceServer) BlockedBy(context.Context, *UserIdRequest) (*BlockedResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BlockedBy not implemented")
}
func (UnimplementedConnectionServiceServer) BlockedAny(context.Context, *UserIdRequest) (*BlockedResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BlockedAny not implemented")
}
func (UnimplementedConnectionServiceServer) ChangeMessageNotification(context.Context, *UserConnectionRequest) (*UserConnectionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ChangeMessageNotification not implemented")
}
func (UnimplementedConnectionServiceServer) ChangePostNotification(context.Context, *UserConnectionRequest) (*UserConnectionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ChangePostNotification not implemented")
}
func (UnimplementedConnectionServiceServer) ChangeCommentNotification(context.Context, *UserConnectionRequest) (*UserConnectionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ChangeCommentNotification not implemented")
}
func (UnimplementedConnectionServiceServer) mustEmbedUnimplementedConnectionServiceServer() {}

// UnsafeConnectionServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ConnectionServiceServer will
// result in compilation errors.
type UnsafeConnectionServiceServer interface {
	mustEmbedUnimplementedConnectionServiceServer()
}

func RegisterConnectionServiceServer(s grpc.ServiceRegistrar, srv ConnectionServiceServer) {
	s.RegisterService(&ConnectionService_ServiceDesc, srv)
}

func _ConnectionService_NewUserConnection_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserConnectionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConnectionServiceServer).NewUserConnection(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/connection.ConnectionService/NewUserConnection",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConnectionServiceServer).NewUserConnection(ctx, req.(*UserConnectionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ConnectionService_ApproveConnection_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserConnectionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConnectionServiceServer).ApproveConnection(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/connection.ConnectionService/ApproveConnection",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConnectionServiceServer).ApproveConnection(ctx, req.(*UserConnectionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ConnectionService_GetConnection_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Connection)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConnectionServiceServer).GetConnection(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/connection.ConnectionService/GetConnection",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConnectionServiceServer).GetConnection(ctx, req.(*Connection))
	}
	return interceptor(ctx, in, info, handler)
}

func _ConnectionService_ApproveAllConnection_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConnectionServiceServer).ApproveAllConnection(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/connection.ConnectionService/ApproveAllConnection",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConnectionServiceServer).ApproveAllConnection(ctx, req.(*UserIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ConnectionService_RejectConnection_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserConnectionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConnectionServiceServer).RejectConnection(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/connection.ConnectionService/RejectConnection",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConnectionServiceServer).RejectConnection(ctx, req.(*UserConnectionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ConnectionService_DeleteConnection_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Connection)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConnectionServiceServer).DeleteConnection(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/connection.ConnectionService/DeleteConnection",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConnectionServiceServer).DeleteConnection(ctx, req.(*Connection))
	}
	return interceptor(ctx, in, info, handler)
}

func _ConnectionService_GetAllConnections_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConnectionServiceServer).GetAllConnections(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/connection.ConnectionService/GetAllConnections",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConnectionServiceServer).GetAllConnections(ctx, req.(*UserIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ConnectionService_GetFollowings_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConnectionServiceServer).GetFollowings(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/connection.ConnectionService/GetFollowings",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConnectionServiceServer).GetFollowings(ctx, req.(*UserIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ConnectionService_GetFollowers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConnectionServiceServer).GetFollowers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/connection.ConnectionService/GetFollowers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConnectionServiceServer).GetFollowers(ctx, req.(*UserIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ConnectionService_GetAllRequestConnectionsByUserId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConnectionServiceServer).GetAllRequestConnectionsByUserId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/connection.ConnectionService/GetAllRequestConnectionsByUserId",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConnectionServiceServer).GetAllRequestConnectionsByUserId(ctx, req.(*UserIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ConnectionService_GetAllPendingConnectionsByUserId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConnectionServiceServer).GetAllPendingConnectionsByUserId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/connection.ConnectionService/GetAllPendingConnectionsByUserId",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConnectionServiceServer).GetAllPendingConnectionsByUserId(ctx, req.(*UserIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ConnectionService_BlockUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BlockUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConnectionServiceServer).BlockUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/connection.ConnectionService/BlockUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConnectionServiceServer).BlockUser(ctx, req.(*BlockUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ConnectionService_UnblockUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BlockUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConnectionServiceServer).UnblockUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/connection.ConnectionService/UnblockUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConnectionServiceServer).UnblockUser(ctx, req.(*BlockUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ConnectionService_IsBlocked_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Block)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConnectionServiceServer).IsBlocked(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/connection.ConnectionService/IsBlocked",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConnectionServiceServer).IsBlocked(ctx, req.(*Block))
	}
	return interceptor(ctx, in, info, handler)
}

func _ConnectionService_IsBlockedAny_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Block)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConnectionServiceServer).IsBlockedAny(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/connection.ConnectionService/IsBlockedAny",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConnectionServiceServer).IsBlockedAny(ctx, req.(*Block))
	}
	return interceptor(ctx, in, info, handler)
}

func _ConnectionService_Blocked_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConnectionServiceServer).Blocked(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/connection.ConnectionService/Blocked",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConnectionServiceServer).Blocked(ctx, req.(*UserIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ConnectionService_BlockedBy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConnectionServiceServer).BlockedBy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/connection.ConnectionService/BlockedBy",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConnectionServiceServer).BlockedBy(ctx, req.(*UserIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ConnectionService_BlockedAny_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConnectionServiceServer).BlockedAny(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/connection.ConnectionService/BlockedAny",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConnectionServiceServer).BlockedAny(ctx, req.(*UserIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ConnectionService_ChangeMessageNotification_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserConnectionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConnectionServiceServer).ChangeMessageNotification(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/connection.ConnectionService/ChangeMessageNotification",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConnectionServiceServer).ChangeMessageNotification(ctx, req.(*UserConnectionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ConnectionService_ChangePostNotification_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserConnectionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConnectionServiceServer).ChangePostNotification(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/connection.ConnectionService/ChangePostNotification",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConnectionServiceServer).ChangePostNotification(ctx, req.(*UserConnectionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ConnectionService_ChangeCommentNotification_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserConnectionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConnectionServiceServer).ChangeCommentNotification(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/connection.ConnectionService/ChangeCommentNotification",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConnectionServiceServer).ChangeCommentNotification(ctx, req.(*UserConnectionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ConnectionService_ServiceDesc is the grpc.ServiceDesc for ConnectionService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ConnectionService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "connection.ConnectionService",
	HandlerType: (*ConnectionServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "NewUserConnection",
			Handler:    _ConnectionService_NewUserConnection_Handler,
		},
		{
			MethodName: "ApproveConnection",
			Handler:    _ConnectionService_ApproveConnection_Handler,
		},
		{
			MethodName: "GetConnection",
			Handler:    _ConnectionService_GetConnection_Handler,
		},
		{
			MethodName: "ApproveAllConnection",
			Handler:    _ConnectionService_ApproveAllConnection_Handler,
		},
		{
			MethodName: "RejectConnection",
			Handler:    _ConnectionService_RejectConnection_Handler,
		},
		{
			MethodName: "DeleteConnection",
			Handler:    _ConnectionService_DeleteConnection_Handler,
		},
		{
			MethodName: "GetAllConnections",
			Handler:    _ConnectionService_GetAllConnections_Handler,
		},
		{
			MethodName: "GetFollowings",
			Handler:    _ConnectionService_GetFollowings_Handler,
		},
		{
			MethodName: "GetFollowers",
			Handler:    _ConnectionService_GetFollowers_Handler,
		},
		{
			MethodName: "GetAllRequestConnectionsByUserId",
			Handler:    _ConnectionService_GetAllRequestConnectionsByUserId_Handler,
		},
		{
			MethodName: "GetAllPendingConnectionsByUserId",
			Handler:    _ConnectionService_GetAllPendingConnectionsByUserId_Handler,
		},
		{
			MethodName: "BlockUser",
			Handler:    _ConnectionService_BlockUser_Handler,
		},
		{
			MethodName: "UnblockUser",
			Handler:    _ConnectionService_UnblockUser_Handler,
		},
		{
			MethodName: "IsBlocked",
			Handler:    _ConnectionService_IsBlocked_Handler,
		},
		{
			MethodName: "IsBlockedAny",
			Handler:    _ConnectionService_IsBlockedAny_Handler,
		},
		{
			MethodName: "Blocked",
			Handler:    _ConnectionService_Blocked_Handler,
		},
		{
			MethodName: "BlockedBy",
			Handler:    _ConnectionService_BlockedBy_Handler,
		},
		{
			MethodName: "BlockedAny",
			Handler:    _ConnectionService_BlockedAny_Handler,
		},
		{
			MethodName: "ChangeMessageNotification",
			Handler:    _ConnectionService_ChangeMessageNotification_Handler,
		},
		{
			MethodName: "ChangePostNotification",
			Handler:    _ConnectionService_ChangePostNotification_Handler,
		},
		{
			MethodName: "ChangeCommentNotification",
			Handler:    _ConnectionService_ChangeCommentNotification_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "connection/connectionService.proto",
}
