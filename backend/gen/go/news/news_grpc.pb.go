// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.25.3
// source: news.proto

package news

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

// NewsServiceClient is the client API for NewsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type NewsServiceClient interface {
	GetNews(ctx context.Context, in *GetNewsRequest, opts ...grpc.CallOption) (*GetNewsResponse, error)
	GetNewsById(ctx context.Context, in *GetNewsByIdRequest, opts ...grpc.CallOption) (*GetNewsByIdResponse, error)
	GetNewsByCategory(ctx context.Context, in *GetNewsByCategoryRequest, opts ...grpc.CallOption) (*GetNewsResponse, error)
	AddNews(ctx context.Context, in *AddNewsRequest, opts ...grpc.CallOption) (*AddNewsResponse, error)
	DelNews(ctx context.Context, in *DelNewsRequest, opts ...grpc.CallOption) (*DelNewsResponse, error)
}

type newsServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewNewsServiceClient(cc grpc.ClientConnInterface) NewsServiceClient {
	return &newsServiceClient{cc}
}

func (c *newsServiceClient) GetNews(ctx context.Context, in *GetNewsRequest, opts ...grpc.CallOption) (*GetNewsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetNewsResponse)
	err := c.cc.Invoke(ctx, NewsService_GetNews_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *newsServiceClient) GetNewsById(ctx context.Context, in *GetNewsByIdRequest, opts ...grpc.CallOption) (*GetNewsByIdResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetNewsByIdResponse)
	err := c.cc.Invoke(ctx, NewsService_GetNewsById_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *newsServiceClient) GetNewsByCategory(ctx context.Context, in *GetNewsByCategoryRequest, opts ...grpc.CallOption) (*GetNewsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetNewsResponse)
	err := c.cc.Invoke(ctx, NewsService_GetNewsByCategory_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *newsServiceClient) AddNews(ctx context.Context, in *AddNewsRequest, opts ...grpc.CallOption) (*AddNewsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AddNewsResponse)
	err := c.cc.Invoke(ctx, NewsService_AddNews_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *newsServiceClient) DelNews(ctx context.Context, in *DelNewsRequest, opts ...grpc.CallOption) (*DelNewsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DelNewsResponse)
	err := c.cc.Invoke(ctx, NewsService_DelNews_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NewsServiceServer is the server API for NewsService service.
// All implementations must embed UnimplementedNewsServiceServer
// for forward compatibility
type NewsServiceServer interface {
	GetNews(context.Context, *GetNewsRequest) (*GetNewsResponse, error)
	GetNewsById(context.Context, *GetNewsByIdRequest) (*GetNewsByIdResponse, error)
	GetNewsByCategory(context.Context, *GetNewsByCategoryRequest) (*GetNewsResponse, error)
	AddNews(context.Context, *AddNewsRequest) (*AddNewsResponse, error)
	DelNews(context.Context, *DelNewsRequest) (*DelNewsResponse, error)
	mustEmbedUnimplementedNewsServiceServer()
}

// UnimplementedNewsServiceServer must be embedded to have forward compatible implementations.
type UnimplementedNewsServiceServer struct {
}

func (UnimplementedNewsServiceServer) GetNews(context.Context, *GetNewsRequest) (*GetNewsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetNews not implemented")
}
func (UnimplementedNewsServiceServer) GetNewsById(context.Context, *GetNewsByIdRequest) (*GetNewsByIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetNewsById not implemented")
}
func (UnimplementedNewsServiceServer) GetNewsByCategory(context.Context, *GetNewsByCategoryRequest) (*GetNewsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetNewsByCategory not implemented")
}
func (UnimplementedNewsServiceServer) AddNews(context.Context, *AddNewsRequest) (*AddNewsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddNews not implemented")
}
func (UnimplementedNewsServiceServer) DelNews(context.Context, *DelNewsRequest) (*DelNewsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DelNews not implemented")
}
func (UnimplementedNewsServiceServer) mustEmbedUnimplementedNewsServiceServer() {}

// UnsafeNewsServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to NewsServiceServer will
// result in compilation errors.
type UnsafeNewsServiceServer interface {
	mustEmbedUnimplementedNewsServiceServer()
}

func RegisterNewsServiceServer(s grpc.ServiceRegistrar, srv NewsServiceServer) {
	s.RegisterService(&NewsService_ServiceDesc, srv)
}

func _NewsService_GetNews_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetNewsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NewsServiceServer).GetNews(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/news.NewsService/GetNews",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NewsServiceServer).GetNews(ctx, req.(*GetNewsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NewsService_GetNewsById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetNewsByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NewsServiceServer).GetNewsById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/news.NewsService/GetNewsById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NewsServiceServer).GetNewsById(ctx, req.(*GetNewsByIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NewsService_GetNewsByCategory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetNewsByCategoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NewsServiceServer).GetNewsByCategory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/news.NewsService/GetNewsByCategory",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NewsServiceServer).GetNewsByCategory(ctx, req.(*GetNewsByCategoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NewsService_AddNews_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddNewsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NewsServiceServer).AddNews(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/news.NewsService/AddNews",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NewsServiceServer).AddNews(ctx, req.(*AddNewsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NewsService_DelNews_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DelNewsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NewsServiceServer).DelNews(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/news.NewsService/DelNews",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NewsServiceServer).DelNews(ctx, req.(*DelNewsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// NewsService_ServiceDesc is the grpc.ServiceDesc for NewsService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var NewsService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "news.NewsService",
	HandlerType: (*NewsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetNews",
			Handler:    _NewsService_GetNews_Handler,
		},
		{
			MethodName: "GetNewsById",
			Handler:    _NewsService_GetNewsById_Handler,
		},
		{
			MethodName: "GetNewsByCategory",
			Handler:    _NewsService_GetNewsByCategory_Handler,
		},
		{
			MethodName: "AddNews",
			Handler:    _NewsService_AddNews_Handler,
		},
		{
			MethodName: "DelNews",
			Handler:    _NewsService_DelNews_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "news.proto",
}
