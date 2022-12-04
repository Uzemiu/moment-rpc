// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: moment.proto

package pb

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

// MomentRpcClient is the client API for MomentRpc service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MomentRpcClient interface {
	GetMoment(ctx context.Context, in *GetMomentReq, opts ...grpc.CallOption) (*GetMomentResp, error)
	QueryMoment(ctx context.Context, in *QueryMomentReq, opts ...grpc.CallOption) (*QueryMomentResp, error)
	InsertMoment(ctx context.Context, in *InsertMomentReq, opts ...grpc.CallOption) (*InsertMomentResp, error)
	UpdateMoment(ctx context.Context, in *UpdateMomentReq, opts ...grpc.CallOption) (*UpdateMomentResp, error)
	DeleteMoment(ctx context.Context, in *DeleteMomentReq, opts ...grpc.CallOption) (*DeleteMomentResp, error)
}

type momentRpcClient struct {
	cc grpc.ClientConnInterface
}

func NewMomentRpcClient(cc grpc.ClientConnInterface) MomentRpcClient {
	return &momentRpcClient{cc}
}

func (c *momentRpcClient) GetMoment(ctx context.Context, in *GetMomentReq, opts ...grpc.CallOption) (*GetMomentResp, error) {
	out := new(GetMomentResp)
	err := c.cc.Invoke(ctx, "/moment.MomentRpc/getMoment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *momentRpcClient) QueryMoment(ctx context.Context, in *QueryMomentReq, opts ...grpc.CallOption) (*QueryMomentResp, error) {
	out := new(QueryMomentResp)
	err := c.cc.Invoke(ctx, "/moment.MomentRpc/queryMoment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *momentRpcClient) InsertMoment(ctx context.Context, in *InsertMomentReq, opts ...grpc.CallOption) (*InsertMomentResp, error) {
	out := new(InsertMomentResp)
	err := c.cc.Invoke(ctx, "/moment.MomentRpc/insertMoment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *momentRpcClient) UpdateMoment(ctx context.Context, in *UpdateMomentReq, opts ...grpc.CallOption) (*UpdateMomentResp, error) {
	out := new(UpdateMomentResp)
	err := c.cc.Invoke(ctx, "/moment.MomentRpc/updateMoment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *momentRpcClient) DeleteMoment(ctx context.Context, in *DeleteMomentReq, opts ...grpc.CallOption) (*DeleteMomentResp, error) {
	out := new(DeleteMomentResp)
	err := c.cc.Invoke(ctx, "/moment.MomentRpc/deleteMoment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MomentRpcServer is the server API for MomentRpc service.
// All implementations must embed UnimplementedMomentRpcServer
// for forward compatibility
type MomentRpcServer interface {
	GetMoment(context.Context, *GetMomentReq) (*GetMomentResp, error)
	QueryMoment(context.Context, *QueryMomentReq) (*QueryMomentResp, error)
	InsertMoment(context.Context, *InsertMomentReq) (*InsertMomentResp, error)
	UpdateMoment(context.Context, *UpdateMomentReq) (*UpdateMomentResp, error)
	DeleteMoment(context.Context, *DeleteMomentReq) (*DeleteMomentResp, error)
	mustEmbedUnimplementedMomentRpcServer()
}

// UnimplementedMomentRpcServer must be embedded to have forward compatible implementations.
type UnimplementedMomentRpcServer struct {
}

func (UnimplementedMomentRpcServer) GetMoment(context.Context, *GetMomentReq) (*GetMomentResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMoment not implemented")
}
func (UnimplementedMomentRpcServer) QueryMoment(context.Context, *QueryMomentReq) (*QueryMomentResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryMoment not implemented")
}
func (UnimplementedMomentRpcServer) InsertMoment(context.Context, *InsertMomentReq) (*InsertMomentResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InsertMoment not implemented")
}
func (UnimplementedMomentRpcServer) UpdateMoment(context.Context, *UpdateMomentReq) (*UpdateMomentResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateMoment not implemented")
}
func (UnimplementedMomentRpcServer) DeleteMoment(context.Context, *DeleteMomentReq) (*DeleteMomentResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteMoment not implemented")
}
func (UnimplementedMomentRpcServer) mustEmbedUnimplementedMomentRpcServer() {}

// UnsafeMomentRpcServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MomentRpcServer will
// result in compilation errors.
type UnsafeMomentRpcServer interface {
	mustEmbedUnimplementedMomentRpcServer()
}

func RegisterMomentRpcServer(s grpc.ServiceRegistrar, srv MomentRpcServer) {
	s.RegisterService(&MomentRpc_ServiceDesc, srv)
}

func _MomentRpc_GetMoment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMomentReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MomentRpcServer).GetMoment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/moment.MomentRpc/getMoment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MomentRpcServer).GetMoment(ctx, req.(*GetMomentReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _MomentRpc_QueryMoment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryMomentReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MomentRpcServer).QueryMoment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/moment.MomentRpc/queryMoment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MomentRpcServer).QueryMoment(ctx, req.(*QueryMomentReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _MomentRpc_InsertMoment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InsertMomentReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MomentRpcServer).InsertMoment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/moment.MomentRpc/insertMoment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MomentRpcServer).InsertMoment(ctx, req.(*InsertMomentReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _MomentRpc_UpdateMoment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateMomentReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MomentRpcServer).UpdateMoment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/moment.MomentRpc/updateMoment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MomentRpcServer).UpdateMoment(ctx, req.(*UpdateMomentReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _MomentRpc_DeleteMoment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteMomentReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MomentRpcServer).DeleteMoment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/moment.MomentRpc/deleteMoment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MomentRpcServer).DeleteMoment(ctx, req.(*DeleteMomentReq))
	}
	return interceptor(ctx, in, info, handler)
}

// MomentRpc_ServiceDesc is the grpc.ServiceDesc for MomentRpc service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MomentRpc_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "moment.MomentRpc",
	HandlerType: (*MomentRpcServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "getMoment",
			Handler:    _MomentRpc_GetMoment_Handler,
		},
		{
			MethodName: "queryMoment",
			Handler:    _MomentRpc_QueryMoment_Handler,
		},
		{
			MethodName: "insertMoment",
			Handler:    _MomentRpc_InsertMoment_Handler,
		},
		{
			MethodName: "updateMoment",
			Handler:    _MomentRpc_UpdateMoment_Handler,
		},
		{
			MethodName: "deleteMoment",
			Handler:    _MomentRpc_DeleteMoment_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "moment.proto",
}