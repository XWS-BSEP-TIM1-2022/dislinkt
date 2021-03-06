// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.20.0
// source: post/postService.proto

package post

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

// PostServiceClient is the client API for PostService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PostServiceClient interface {
	GetRequest(ctx context.Context, in *PostIdRequest, opts ...grpc.CallOption) (*PostResponse, error)
	GetAllRequest(ctx context.Context, in *EmptyRequest, opts ...grpc.CallOption) (*PostsResponse, error)
	GetAllFromUserRequest(ctx context.Context, in *UserPostsRequest, opts ...grpc.CallOption) (*PostsResponse, error)
	CreateRequest(ctx context.Context, in *PostRequest, opts ...grpc.CallOption) (*PostResponse, error)
	DeleteRequest(ctx context.Context, in *PostIdRequest, opts ...grpc.CallOption) (*EmptyRequest, error)
	GetCommentRequest(ctx context.Context, in *CommentIdRequest, opts ...grpc.CallOption) (*CommentResponse, error)
	GetAllCommentsRequest(ctx context.Context, in *EmptyRequest, opts ...grpc.CallOption) (*CommentsResponse, error)
	GetAllCommentsFromPostRequest(ctx context.Context, in *PostCommentsRequest, opts ...grpc.CallOption) (*CommentsResponse, error)
	CreateCommentRequest(ctx context.Context, in *CommentRequest, opts ...grpc.CallOption) (*CommentResponse, error)
	DeleteCommentRequest(ctx context.Context, in *CommentIdRequest, opts ...grpc.CallOption) (*EmptyRequest, error)
	GetReactionRequest(ctx context.Context, in *ReactionIdRequest, opts ...grpc.CallOption) (*ReactionResponse, error)
	GetAllReactionsRequest(ctx context.Context, in *EmptyRequest, opts ...grpc.CallOption) (*ReactionsResponse, error)
	GetAllReactionsFromPostRequest(ctx context.Context, in *PostReactionRequest, opts ...grpc.CallOption) (*ReactionsResponse, error)
	CreateReactionRequest(ctx context.Context, in *ReactionRequest, opts ...grpc.CallOption) (*ReactionResponse, error)
	DeleteReactionRequest(ctx context.Context, in *ReactionIdRequest, opts ...grpc.CallOption) (*EmptyRequest, error)
}

type postServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPostServiceClient(cc grpc.ClientConnInterface) PostServiceClient {
	return &postServiceClient{cc}
}

func (c *postServiceClient) GetRequest(ctx context.Context, in *PostIdRequest, opts ...grpc.CallOption) (*PostResponse, error) {
	out := new(PostResponse)
	err := c.cc.Invoke(ctx, "/post.PostService/GetRequest", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *postServiceClient) GetAllRequest(ctx context.Context, in *EmptyRequest, opts ...grpc.CallOption) (*PostsResponse, error) {
	out := new(PostsResponse)
	err := c.cc.Invoke(ctx, "/post.PostService/GetAllRequest", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *postServiceClient) GetAllFromUserRequest(ctx context.Context, in *UserPostsRequest, opts ...grpc.CallOption) (*PostsResponse, error) {
	out := new(PostsResponse)
	err := c.cc.Invoke(ctx, "/post.PostService/GetAllFromUserRequest", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *postServiceClient) CreateRequest(ctx context.Context, in *PostRequest, opts ...grpc.CallOption) (*PostResponse, error) {
	out := new(PostResponse)
	err := c.cc.Invoke(ctx, "/post.PostService/CreateRequest", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *postServiceClient) DeleteRequest(ctx context.Context, in *PostIdRequest, opts ...grpc.CallOption) (*EmptyRequest, error) {
	out := new(EmptyRequest)
	err := c.cc.Invoke(ctx, "/post.PostService/DeleteRequest", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *postServiceClient) GetCommentRequest(ctx context.Context, in *CommentIdRequest, opts ...grpc.CallOption) (*CommentResponse, error) {
	out := new(CommentResponse)
	err := c.cc.Invoke(ctx, "/post.PostService/GetCommentRequest", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *postServiceClient) GetAllCommentsRequest(ctx context.Context, in *EmptyRequest, opts ...grpc.CallOption) (*CommentsResponse, error) {
	out := new(CommentsResponse)
	err := c.cc.Invoke(ctx, "/post.PostService/GetAllCommentsRequest", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *postServiceClient) GetAllCommentsFromPostRequest(ctx context.Context, in *PostCommentsRequest, opts ...grpc.CallOption) (*CommentsResponse, error) {
	out := new(CommentsResponse)
	err := c.cc.Invoke(ctx, "/post.PostService/GetAllCommentsFromPostRequest", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *postServiceClient) CreateCommentRequest(ctx context.Context, in *CommentRequest, opts ...grpc.CallOption) (*CommentResponse, error) {
	out := new(CommentResponse)
	err := c.cc.Invoke(ctx, "/post.PostService/CreateCommentRequest", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *postServiceClient) DeleteCommentRequest(ctx context.Context, in *CommentIdRequest, opts ...grpc.CallOption) (*EmptyRequest, error) {
	out := new(EmptyRequest)
	err := c.cc.Invoke(ctx, "/post.PostService/DeleteCommentRequest", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *postServiceClient) GetReactionRequest(ctx context.Context, in *ReactionIdRequest, opts ...grpc.CallOption) (*ReactionResponse, error) {
	out := new(ReactionResponse)
	err := c.cc.Invoke(ctx, "/post.PostService/GetReactionRequest", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *postServiceClient) GetAllReactionsRequest(ctx context.Context, in *EmptyRequest, opts ...grpc.CallOption) (*ReactionsResponse, error) {
	out := new(ReactionsResponse)
	err := c.cc.Invoke(ctx, "/post.PostService/GetAllReactionsRequest", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *postServiceClient) GetAllReactionsFromPostRequest(ctx context.Context, in *PostReactionRequest, opts ...grpc.CallOption) (*ReactionsResponse, error) {
	out := new(ReactionsResponse)
	err := c.cc.Invoke(ctx, "/post.PostService/GetAllReactionsFromPostRequest", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *postServiceClient) CreateReactionRequest(ctx context.Context, in *ReactionRequest, opts ...grpc.CallOption) (*ReactionResponse, error) {
	out := new(ReactionResponse)
	err := c.cc.Invoke(ctx, "/post.PostService/CreateReactionRequest", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *postServiceClient) DeleteReactionRequest(ctx context.Context, in *ReactionIdRequest, opts ...grpc.CallOption) (*EmptyRequest, error) {
	out := new(EmptyRequest)
	err := c.cc.Invoke(ctx, "/post.PostService/DeleteReactionRequest", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PostServiceServer is the server API for PostService service.
// All implementations must embed UnimplementedPostServiceServer
// for forward compatibility
type PostServiceServer interface {
	GetRequest(context.Context, *PostIdRequest) (*PostResponse, error)
	GetAllRequest(context.Context, *EmptyRequest) (*PostsResponse, error)
	GetAllFromUserRequest(context.Context, *UserPostsRequest) (*PostsResponse, error)
	CreateRequest(context.Context, *PostRequest) (*PostResponse, error)
	DeleteRequest(context.Context, *PostIdRequest) (*EmptyRequest, error)
	GetCommentRequest(context.Context, *CommentIdRequest) (*CommentResponse, error)
	GetAllCommentsRequest(context.Context, *EmptyRequest) (*CommentsResponse, error)
	GetAllCommentsFromPostRequest(context.Context, *PostCommentsRequest) (*CommentsResponse, error)
	CreateCommentRequest(context.Context, *CommentRequest) (*CommentResponse, error)
	DeleteCommentRequest(context.Context, *CommentIdRequest) (*EmptyRequest, error)
	GetReactionRequest(context.Context, *ReactionIdRequest) (*ReactionResponse, error)
	GetAllReactionsRequest(context.Context, *EmptyRequest) (*ReactionsResponse, error)
	GetAllReactionsFromPostRequest(context.Context, *PostReactionRequest) (*ReactionsResponse, error)
	CreateReactionRequest(context.Context, *ReactionRequest) (*ReactionResponse, error)
	DeleteReactionRequest(context.Context, *ReactionIdRequest) (*EmptyRequest, error)
	mustEmbedUnimplementedPostServiceServer()
}

// UnimplementedPostServiceServer must be embedded to have forward compatible implementations.
type UnimplementedPostServiceServer struct {
}

func (UnimplementedPostServiceServer) GetRequest(context.Context, *PostIdRequest) (*PostResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRequest not implemented")
}
func (UnimplementedPostServiceServer) GetAllRequest(context.Context, *EmptyRequest) (*PostsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllRequest not implemented")
}
func (UnimplementedPostServiceServer) GetAllFromUserRequest(context.Context, *UserPostsRequest) (*PostsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllFromUserRequest not implemented")
}
func (UnimplementedPostServiceServer) CreateRequest(context.Context, *PostRequest) (*PostResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateRequest not implemented")
}
func (UnimplementedPostServiceServer) DeleteRequest(context.Context, *PostIdRequest) (*EmptyRequest, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteRequest not implemented")
}
func (UnimplementedPostServiceServer) GetCommentRequest(context.Context, *CommentIdRequest) (*CommentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCommentRequest not implemented")
}
func (UnimplementedPostServiceServer) GetAllCommentsRequest(context.Context, *EmptyRequest) (*CommentsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllCommentsRequest not implemented")
}
func (UnimplementedPostServiceServer) GetAllCommentsFromPostRequest(context.Context, *PostCommentsRequest) (*CommentsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllCommentsFromPostRequest not implemented")
}
func (UnimplementedPostServiceServer) CreateCommentRequest(context.Context, *CommentRequest) (*CommentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCommentRequest not implemented")
}
func (UnimplementedPostServiceServer) DeleteCommentRequest(context.Context, *CommentIdRequest) (*EmptyRequest, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteCommentRequest not implemented")
}
func (UnimplementedPostServiceServer) GetReactionRequest(context.Context, *ReactionIdRequest) (*ReactionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetReactionRequest not implemented")
}
func (UnimplementedPostServiceServer) GetAllReactionsRequest(context.Context, *EmptyRequest) (*ReactionsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllReactionsRequest not implemented")
}
func (UnimplementedPostServiceServer) GetAllReactionsFromPostRequest(context.Context, *PostReactionRequest) (*ReactionsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllReactionsFromPostRequest not implemented")
}
func (UnimplementedPostServiceServer) CreateReactionRequest(context.Context, *ReactionRequest) (*ReactionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateReactionRequest not implemented")
}
func (UnimplementedPostServiceServer) DeleteReactionRequest(context.Context, *ReactionIdRequest) (*EmptyRequest, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteReactionRequest not implemented")
}
func (UnimplementedPostServiceServer) mustEmbedUnimplementedPostServiceServer() {}

// UnsafePostServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PostServiceServer will
// result in compilation errors.
type UnsafePostServiceServer interface {
	mustEmbedUnimplementedPostServiceServer()
}

func RegisterPostServiceServer(s grpc.ServiceRegistrar, srv PostServiceServer) {
	s.RegisterService(&PostService_ServiceDesc, srv)
}

func _PostService_GetRequest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PostIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostServiceServer).GetRequest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/post.PostService/GetRequest",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostServiceServer).GetRequest(ctx, req.(*PostIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PostService_GetAllRequest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmptyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostServiceServer).GetAllRequest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/post.PostService/GetAllRequest",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostServiceServer).GetAllRequest(ctx, req.(*EmptyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PostService_GetAllFromUserRequest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserPostsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostServiceServer).GetAllFromUserRequest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/post.PostService/GetAllFromUserRequest",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostServiceServer).GetAllFromUserRequest(ctx, req.(*UserPostsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PostService_CreateRequest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PostRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostServiceServer).CreateRequest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/post.PostService/CreateRequest",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostServiceServer).CreateRequest(ctx, req.(*PostRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PostService_DeleteRequest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PostIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostServiceServer).DeleteRequest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/post.PostService/DeleteRequest",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostServiceServer).DeleteRequest(ctx, req.(*PostIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PostService_GetCommentRequest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CommentIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostServiceServer).GetCommentRequest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/post.PostService/GetCommentRequest",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostServiceServer).GetCommentRequest(ctx, req.(*CommentIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PostService_GetAllCommentsRequest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmptyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostServiceServer).GetAllCommentsRequest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/post.PostService/GetAllCommentsRequest",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostServiceServer).GetAllCommentsRequest(ctx, req.(*EmptyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PostService_GetAllCommentsFromPostRequest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PostCommentsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostServiceServer).GetAllCommentsFromPostRequest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/post.PostService/GetAllCommentsFromPostRequest",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostServiceServer).GetAllCommentsFromPostRequest(ctx, req.(*PostCommentsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PostService_CreateCommentRequest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CommentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostServiceServer).CreateCommentRequest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/post.PostService/CreateCommentRequest",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostServiceServer).CreateCommentRequest(ctx, req.(*CommentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PostService_DeleteCommentRequest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CommentIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostServiceServer).DeleteCommentRequest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/post.PostService/DeleteCommentRequest",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostServiceServer).DeleteCommentRequest(ctx, req.(*CommentIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PostService_GetReactionRequest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReactionIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostServiceServer).GetReactionRequest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/post.PostService/GetReactionRequest",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostServiceServer).GetReactionRequest(ctx, req.(*ReactionIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PostService_GetAllReactionsRequest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmptyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostServiceServer).GetAllReactionsRequest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/post.PostService/GetAllReactionsRequest",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostServiceServer).GetAllReactionsRequest(ctx, req.(*EmptyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PostService_GetAllReactionsFromPostRequest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PostReactionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostServiceServer).GetAllReactionsFromPostRequest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/post.PostService/GetAllReactionsFromPostRequest",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostServiceServer).GetAllReactionsFromPostRequest(ctx, req.(*PostReactionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PostService_CreateReactionRequest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReactionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostServiceServer).CreateReactionRequest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/post.PostService/CreateReactionRequest",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostServiceServer).CreateReactionRequest(ctx, req.(*ReactionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PostService_DeleteReactionRequest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReactionIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostServiceServer).DeleteReactionRequest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/post.PostService/DeleteReactionRequest",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostServiceServer).DeleteReactionRequest(ctx, req.(*ReactionIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// PostService_ServiceDesc is the grpc.ServiceDesc for PostService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PostService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "post.PostService",
	HandlerType: (*PostServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetRequest",
			Handler:    _PostService_GetRequest_Handler,
		},
		{
			MethodName: "GetAllRequest",
			Handler:    _PostService_GetAllRequest_Handler,
		},
		{
			MethodName: "GetAllFromUserRequest",
			Handler:    _PostService_GetAllFromUserRequest_Handler,
		},
		{
			MethodName: "CreateRequest",
			Handler:    _PostService_CreateRequest_Handler,
		},
		{
			MethodName: "DeleteRequest",
			Handler:    _PostService_DeleteRequest_Handler,
		},
		{
			MethodName: "GetCommentRequest",
			Handler:    _PostService_GetCommentRequest_Handler,
		},
		{
			MethodName: "GetAllCommentsRequest",
			Handler:    _PostService_GetAllCommentsRequest_Handler,
		},
		{
			MethodName: "GetAllCommentsFromPostRequest",
			Handler:    _PostService_GetAllCommentsFromPostRequest_Handler,
		},
		{
			MethodName: "CreateCommentRequest",
			Handler:    _PostService_CreateCommentRequest_Handler,
		},
		{
			MethodName: "DeleteCommentRequest",
			Handler:    _PostService_DeleteCommentRequest_Handler,
		},
		{
			MethodName: "GetReactionRequest",
			Handler:    _PostService_GetReactionRequest_Handler,
		},
		{
			MethodName: "GetAllReactionsRequest",
			Handler:    _PostService_GetAllReactionsRequest_Handler,
		},
		{
			MethodName: "GetAllReactionsFromPostRequest",
			Handler:    _PostService_GetAllReactionsFromPostRequest_Handler,
		},
		{
			MethodName: "CreateReactionRequest",
			Handler:    _PostService_CreateReactionRequest_Handler,
		},
		{
			MethodName: "DeleteReactionRequest",
			Handler:    _PostService_DeleteReactionRequest_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "post/postService.proto",
}
