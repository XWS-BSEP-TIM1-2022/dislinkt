// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.20.0
// source: job/jobService.proto

package job

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

// JobServiceClient is the client API for JobService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type JobServiceClient interface {
	GetRequest(ctx context.Context, in *JobIdRequest, opts ...grpc.CallOption) (*GetResponse, error)
	GetAllRequest(ctx context.Context, in *EmptyRequest, opts ...grpc.CallOption) (*JobsResponse, error)
	PostRequest(ctx context.Context, in *UserRequest, opts ...grpc.CallOption) (*GetResponse, error)
	DeleteRequest(ctx context.Context, in *JobIdRequest, opts ...grpc.CallOption) (*EmptyRequest, error)
	SearchJobsRequest(ctx context.Context, in *SearchRequest, opts ...grpc.CallOption) (*JobsResponse, error)
}

type jobServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewJobServiceClient(cc grpc.ClientConnInterface) JobServiceClient {
	return &jobServiceClient{cc}
}

func (c *jobServiceClient) GetRequest(ctx context.Context, in *JobIdRequest, opts ...grpc.CallOption) (*GetResponse, error) {
	out := new(GetResponse)
	err := c.cc.Invoke(ctx, "/job.JobService/GetRequest", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *jobServiceClient) GetAllRequest(ctx context.Context, in *EmptyRequest, opts ...grpc.CallOption) (*JobsResponse, error) {
	out := new(JobsResponse)
	err := c.cc.Invoke(ctx, "/job.JobService/GetAllRequest", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *jobServiceClient) PostRequest(ctx context.Context, in *UserRequest, opts ...grpc.CallOption) (*GetResponse, error) {
	out := new(GetResponse)
	err := c.cc.Invoke(ctx, "/job.JobService/PostRequest", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *jobServiceClient) DeleteRequest(ctx context.Context, in *JobIdRequest, opts ...grpc.CallOption) (*EmptyRequest, error) {
	out := new(EmptyRequest)
	err := c.cc.Invoke(ctx, "/job.JobService/DeleteRequest", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *jobServiceClient) SearchJobsRequest(ctx context.Context, in *SearchRequest, opts ...grpc.CallOption) (*JobsResponse, error) {
	out := new(JobsResponse)
	err := c.cc.Invoke(ctx, "/job.JobService/SearchJobsRequest", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// JobServiceServer is the server API for JobService service.
// All implementations must embed UnimplementedJobServiceServer
// for forward compatibility
type JobServiceServer interface {
	GetRequest(context.Context, *JobIdRequest) (*GetResponse, error)
	GetAllRequest(context.Context, *EmptyRequest) (*JobsResponse, error)
	PostRequest(context.Context, *UserRequest) (*GetResponse, error)
	DeleteRequest(context.Context, *JobIdRequest) (*EmptyRequest, error)
	SearchJobsRequest(context.Context, *SearchRequest) (*JobsResponse, error)
	mustEmbedUnimplementedJobServiceServer()
}

// UnimplementedJobServiceServer must be embedded to have forward compatible implementations.
type UnimplementedJobServiceServer struct {
}

func (UnimplementedJobServiceServer) GetRequest(context.Context, *JobIdRequest) (*GetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRequest not implemented")
}
func (UnimplementedJobServiceServer) GetAllRequest(context.Context, *EmptyRequest) (*JobsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllRequest not implemented")
}
func (UnimplementedJobServiceServer) PostRequest(context.Context, *UserRequest) (*GetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PostRequest not implemented")
}
func (UnimplementedJobServiceServer) DeleteRequest(context.Context, *JobIdRequest) (*EmptyRequest, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteRequest not implemented")
}
func (UnimplementedJobServiceServer) SearchJobsRequest(context.Context, *SearchRequest) (*JobsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchJobsRequest not implemented")
}
func (UnimplementedJobServiceServer) mustEmbedUnimplementedJobServiceServer() {}

// UnsafeJobServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to JobServiceServer will
// result in compilation errors.
type UnsafeJobServiceServer interface {
	mustEmbedUnimplementedJobServiceServer()
}

func RegisterJobServiceServer(s grpc.ServiceRegistrar, srv JobServiceServer) {
	s.RegisterService(&JobService_ServiceDesc, srv)
}

func _JobService_GetRequest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(JobIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JobServiceServer).GetRequest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/job.JobService/GetRequest",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JobServiceServer).GetRequest(ctx, req.(*JobIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _JobService_GetAllRequest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmptyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JobServiceServer).GetAllRequest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/job.JobService/GetAllRequest",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JobServiceServer).GetAllRequest(ctx, req.(*EmptyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _JobService_PostRequest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JobServiceServer).PostRequest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/job.JobService/PostRequest",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JobServiceServer).PostRequest(ctx, req.(*UserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _JobService_DeleteRequest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(JobIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JobServiceServer).DeleteRequest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/job.JobService/DeleteRequest",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JobServiceServer).DeleteRequest(ctx, req.(*JobIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _JobService_SearchJobsRequest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JobServiceServer).SearchJobsRequest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/job.JobService/SearchJobsRequest",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JobServiceServer).SearchJobsRequest(ctx, req.(*SearchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// JobService_ServiceDesc is the grpc.ServiceDesc for JobService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var JobService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "job.JobService",
	HandlerType: (*JobServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetRequest",
			Handler:    _JobService_GetRequest_Handler,
		},
		{
			MethodName: "GetAllRequest",
			Handler:    _JobService_GetAllRequest_Handler,
		},
		{
			MethodName: "PostRequest",
			Handler:    _JobService_PostRequest_Handler,
		},
		{
			MethodName: "DeleteRequest",
			Handler:    _JobService_DeleteRequest_Handler,
		},
		{
			MethodName: "SearchJobsRequest",
			Handler:    _JobService_SearchJobsRequest_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "job/jobService.proto",
}
