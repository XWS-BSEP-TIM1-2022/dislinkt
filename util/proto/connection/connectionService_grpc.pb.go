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
	ApproveAllConnection(ctx context.Context, in *UserIdRequest, opts ...grpc.CallOption) (*EmptyRequest, error)
	RejectConnection(ctx context.Context, in *UserConnectionRequest, opts ...grpc.CallOption) (*UserConnectionResponse, error)
	DeleteConnection(ctx context.Context, in *Connection, opts ...grpc.CallOption) (*UserConnectionResponse, error)
	GetAllConnections(ctx context.Context, in *UserIdRequest, opts ...grpc.CallOption) (*AllConnectionResponse, error)
	GetFollowings(ctx context.Context, in *UserIdRequest, opts ...grpc.CallOption) (*AllConnectionResponse, error)
	GetFollowers(ctx context.Context, in *UserIdRequest, opts ...grpc.CallOption) (*AllConnectionResponse, error)
	GetAllRequestConnectionsByUserId(ctx context.Context, in *UserIdRequest, opts ...grpc.CallOption) (*AllConnectionResponse, error)
	GetAllPendingConnectionsByUserId(ctx context.Context, in *UserIdRequest, opts ...grpc.CallOption) (*AllConnectionResponse, error)
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

// ConnectionServiceServer is the server API for ConnectionService service.
// All implementations must embed UnimplementedConnectionServiceServer
// for forward compatibility
type ConnectionServiceServer interface {
	NewUserConnection(context.Context, *UserConnectionRequest) (*UserConnectionResponse, error)
	ApproveConnection(context.Context, *UserConnectionRequest) (*UserConnectionResponse, error)
	ApproveAllConnection(context.Context, *UserIdRequest) (*EmptyRequest, error)
	RejectConnection(context.Context, *UserConnectionRequest) (*UserConnectionResponse, error)
	DeleteConnection(context.Context, *Connection) (*UserConnectionResponse, error)
	GetAllConnections(context.Context, *UserIdRequest) (*AllConnectionResponse, error)
	GetFollowings(context.Context, *UserIdRequest) (*AllConnectionResponse, error)
	GetFollowers(context.Context, *UserIdRequest) (*AllConnectionResponse, error)
	GetAllRequestConnectionsByUserId(context.Context, *UserIdRequest) (*AllConnectionResponse, error)
	GetAllPendingConnectionsByUserId(context.Context, *UserIdRequest) (*AllConnectionResponse, error)
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
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "connection/connectionService.proto",
}
