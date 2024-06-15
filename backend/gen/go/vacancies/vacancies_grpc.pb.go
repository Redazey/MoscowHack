// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v5.27.1
// source: vacancies.proto

package vacancies

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.62.0 or later.
const _ = grpc.SupportPackageIsVersion8

const (
	VacanciesService_GetVacancies_FullMethodName         = "/vacancies.VacanciesService/GetVacancies"
	VacanciesService_GetVacanciesById_FullMethodName     = "/vacancies.VacanciesService/GetVacanciesById"
	VacanciesService_GetVacanciesByFilter_FullMethodName = "/vacancies.VacanciesService/GetVacanciesByFilter"
	VacanciesService_AddVacancies_FullMethodName         = "/vacancies.VacanciesService/AddVacancies"
	VacanciesService_DelVacancies_FullMethodName         = "/vacancies.VacanciesService/DelVacancies"
)

// VacanciesServiceClient is the client API for VacanciesService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type VacanciesServiceClient interface {
	GetVacancies(ctx context.Context, in *GetVacanciesRequest, opts ...grpc.CallOption) (*GetVacanciesResponse, error)
	GetVacanciesById(ctx context.Context, in *GetVacanciesByIdRequest, opts ...grpc.CallOption) (*GetVacanciesByIdResponse, error)
	GetVacanciesByFilter(ctx context.Context, in *GetFilterVacanciesRequest, opts ...grpc.CallOption) (*GetVacanciesResponse, error)
	AddVacancies(ctx context.Context, in *AddVacanciesRequest, opts ...grpc.CallOption) (*AddVacanciesResponse, error)
	DelVacancies(ctx context.Context, in *DelVacanciesRequest, opts ...grpc.CallOption) (*DelVacanciesResponse, error)
}

type vacanciesServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewVacanciesServiceClient(cc grpc.ClientConnInterface) VacanciesServiceClient {
	return &vacanciesServiceClient{cc}
}

func (c *vacanciesServiceClient) GetVacancies(ctx context.Context, in *GetVacanciesRequest, opts ...grpc.CallOption) (*GetVacanciesResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetVacanciesResponse)
	err := c.cc.Invoke(ctx, VacanciesService_GetVacancies_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vacanciesServiceClient) GetVacanciesById(ctx context.Context, in *GetVacanciesByIdRequest, opts ...grpc.CallOption) (*GetVacanciesByIdResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetVacanciesByIdResponse)
	err := c.cc.Invoke(ctx, VacanciesService_GetVacanciesById_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vacanciesServiceClient) GetVacanciesByFilter(ctx context.Context, in *GetFilterVacanciesRequest, opts ...grpc.CallOption) (*GetVacanciesResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetVacanciesResponse)
	err := c.cc.Invoke(ctx, VacanciesService_GetVacanciesByFilter_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vacanciesServiceClient) AddVacancies(ctx context.Context, in *AddVacanciesRequest, opts ...grpc.CallOption) (*AddVacanciesResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AddVacanciesResponse)
	err := c.cc.Invoke(ctx, VacanciesService_AddVacancies_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vacanciesServiceClient) DelVacancies(ctx context.Context, in *DelVacanciesRequest, opts ...grpc.CallOption) (*DelVacanciesResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DelVacanciesResponse)
	err := c.cc.Invoke(ctx, VacanciesService_DelVacancies_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// VacanciesServiceServer is the server API for VacanciesService service.
// All implementations must embed UnimplementedVacanciesServiceServer
// for forward compatibility
type VacanciesServiceServer interface {
	GetVacancies(context.Context, *GetVacanciesRequest) (*GetVacanciesResponse, error)
	GetVacanciesById(context.Context, *GetVacanciesByIdRequest) (*GetVacanciesByIdResponse, error)
	GetVacanciesByFilter(context.Context, *GetFilterVacanciesRequest) (*GetVacanciesResponse, error)
	AddVacancies(context.Context, *AddVacanciesRequest) (*AddVacanciesResponse, error)
	DelVacancies(context.Context, *DelVacanciesRequest) (*DelVacanciesResponse, error)
	mustEmbedUnimplementedVacanciesServiceServer()
}

// UnimplementedVacanciesServiceServer must be embedded to have forward compatible implementations.
type UnimplementedVacanciesServiceServer struct {
}

func (UnimplementedVacanciesServiceServer) GetVacancies(context.Context, *GetVacanciesRequest) (*GetVacanciesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetVacancies not implemented")
}
func (UnimplementedVacanciesServiceServer) GetVacanciesById(context.Context, *GetVacanciesByIdRequest) (*GetVacanciesByIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetVacanciesById not implemented")
}
func (UnimplementedVacanciesServiceServer) GetVacanciesByFilter(context.Context, *GetFilterVacanciesRequest) (*GetVacanciesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetVacanciesByFilter not implemented")
}
func (UnimplementedVacanciesServiceServer) AddVacancies(context.Context, *AddVacanciesRequest) (*AddVacanciesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddVacancies not implemented")
}
func (UnimplementedVacanciesServiceServer) DelVacancies(context.Context, *DelVacanciesRequest) (*DelVacanciesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DelVacancies not implemented")
}
func (UnimplementedVacanciesServiceServer) mustEmbedUnimplementedVacanciesServiceServer() {}

// UnsafeVacanciesServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to VacanciesServiceServer will
// result in compilation errors.
type UnsafeVacanciesServiceServer interface {
	mustEmbedUnimplementedVacanciesServiceServer()
}

func RegisterVacanciesServiceServer(s grpc.ServiceRegistrar, srv VacanciesServiceServer) {
	s.RegisterService(&VacanciesService_ServiceDesc, srv)
}

func _VacanciesService_GetVacancies_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetVacanciesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VacanciesServiceServer).GetVacancies(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: VacanciesService_GetVacancies_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VacanciesServiceServer).GetVacancies(ctx, req.(*GetVacanciesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VacanciesService_GetVacanciesById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetVacanciesByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VacanciesServiceServer).GetVacanciesById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: VacanciesService_GetVacanciesById_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VacanciesServiceServer).GetVacanciesById(ctx, req.(*GetVacanciesByIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VacanciesService_GetVacanciesByFilter_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetFilterVacanciesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VacanciesServiceServer).GetVacanciesByFilter(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: VacanciesService_GetVacanciesByFilter_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VacanciesServiceServer).GetVacanciesByFilter(ctx, req.(*GetFilterVacanciesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VacanciesService_AddVacancies_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddVacanciesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VacanciesServiceServer).AddVacancies(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: VacanciesService_AddVacancies_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VacanciesServiceServer).AddVacancies(ctx, req.(*AddVacanciesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VacanciesService_DelVacancies_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DelVacanciesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VacanciesServiceServer).DelVacancies(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: VacanciesService_DelVacancies_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VacanciesServiceServer).DelVacancies(ctx, req.(*DelVacanciesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// VacanciesService_ServiceDesc is the grpc.ServiceDesc for VacanciesService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var VacanciesService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "vacancies.VacanciesService",
	HandlerType: (*VacanciesServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetVacancies",
			Handler:    _VacanciesService_GetVacancies_Handler,
		},
		{
			MethodName: "GetVacanciesById",
			Handler:    _VacanciesService_GetVacanciesById_Handler,
		},
		{
			MethodName: "GetVacanciesByFilter",
			Handler:    _VacanciesService_GetVacanciesByFilter_Handler,
		},
		{
			MethodName: "AddVacancies",
			Handler:    _VacanciesService_AddVacancies_Handler,
		},
		{
			MethodName: "DelVacancies",
			Handler:    _VacanciesService_DelVacancies_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "vacancies.proto",
}
