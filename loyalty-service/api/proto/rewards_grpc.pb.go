// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v5.26.1
// source: rewards.proto

package loyalty_service

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

// PosRewardServiceClient is the client API for PosRewardService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PosRewardServiceClient interface {
	CreatePosReward(ctx context.Context, in *CreatePosRewardRequest, opts ...grpc.CallOption) (*CreatePosRewardResponse, error)
	ReadPosReward(ctx context.Context, in *ReadPosRewardRequest, opts ...grpc.CallOption) (*ReadPosRewardResponse, error)
	UpdatePosReward(ctx context.Context, in *UpdatePosRewardRequest, opts ...grpc.CallOption) (*UpdatePosRewardResponse, error)
	DeletePosReward(ctx context.Context, in *DeletePosRewardRequest, opts ...grpc.CallOption) (*DeletePosRewardResponse, error)
	ReadAllPosRewards(ctx context.Context, in *ReadAllPosRewardsRequest, opts ...grpc.CallOption) (*ReadAllPosRewardsResponse, error)
}

type posRewardServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPosRewardServiceClient(cc grpc.ClientConnInterface) PosRewardServiceClient {
	return &posRewardServiceClient{cc}
}

func (c *posRewardServiceClient) CreatePosReward(ctx context.Context, in *CreatePosRewardRequest, opts ...grpc.CallOption) (*CreatePosRewardResponse, error) {
	out := new(CreatePosRewardResponse)
	err := c.cc.Invoke(ctx, "/pos.PosRewardService/CreatePosReward", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *posRewardServiceClient) ReadPosReward(ctx context.Context, in *ReadPosRewardRequest, opts ...grpc.CallOption) (*ReadPosRewardResponse, error) {
	out := new(ReadPosRewardResponse)
	err := c.cc.Invoke(ctx, "/pos.PosRewardService/ReadPosReward", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *posRewardServiceClient) UpdatePosReward(ctx context.Context, in *UpdatePosRewardRequest, opts ...grpc.CallOption) (*UpdatePosRewardResponse, error) {
	out := new(UpdatePosRewardResponse)
	err := c.cc.Invoke(ctx, "/pos.PosRewardService/UpdatePosReward", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *posRewardServiceClient) DeletePosReward(ctx context.Context, in *DeletePosRewardRequest, opts ...grpc.CallOption) (*DeletePosRewardResponse, error) {
	out := new(DeletePosRewardResponse)
	err := c.cc.Invoke(ctx, "/pos.PosRewardService/DeletePosReward", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *posRewardServiceClient) ReadAllPosRewards(ctx context.Context, in *ReadAllPosRewardsRequest, opts ...grpc.CallOption) (*ReadAllPosRewardsResponse, error) {
	out := new(ReadAllPosRewardsResponse)
	err := c.cc.Invoke(ctx, "/pos.PosRewardService/ReadAllPosRewards", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PosRewardServiceServer is the server API for PosRewardService service.
// All implementations must embed UnimplementedPosRewardServiceServer
// for forward compatibility
type PosRewardServiceServer interface {
	CreatePosReward(context.Context, *CreatePosRewardRequest) (*CreatePosRewardResponse, error)
	ReadPosReward(context.Context, *ReadPosRewardRequest) (*ReadPosRewardResponse, error)
	UpdatePosReward(context.Context, *UpdatePosRewardRequest) (*UpdatePosRewardResponse, error)
	DeletePosReward(context.Context, *DeletePosRewardRequest) (*DeletePosRewardResponse, error)
	ReadAllPosRewards(context.Context, *ReadAllPosRewardsRequest) (*ReadAllPosRewardsResponse, error)
	mustEmbedUnimplementedPosRewardServiceServer()
}

// UnimplementedPosRewardServiceServer must be embedded to have forward compatible implementations.
type UnimplementedPosRewardServiceServer struct {
}

func (UnimplementedPosRewardServiceServer) CreatePosReward(context.Context, *CreatePosRewardRequest) (*CreatePosRewardResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePosReward not implemented")
}
func (UnimplementedPosRewardServiceServer) ReadPosReward(context.Context, *ReadPosRewardRequest) (*ReadPosRewardResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReadPosReward not implemented")
}
func (UnimplementedPosRewardServiceServer) UpdatePosReward(context.Context, *UpdatePosRewardRequest) (*UpdatePosRewardResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdatePosReward not implemented")
}
func (UnimplementedPosRewardServiceServer) DeletePosReward(context.Context, *DeletePosRewardRequest) (*DeletePosRewardResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeletePosReward not implemented")
}
func (UnimplementedPosRewardServiceServer) ReadAllPosRewards(context.Context, *ReadAllPosRewardsRequest) (*ReadAllPosRewardsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReadAllPosRewards not implemented")
}
func (UnimplementedPosRewardServiceServer) mustEmbedUnimplementedPosRewardServiceServer() {}

// UnsafePosRewardServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PosRewardServiceServer will
// result in compilation errors.
type UnsafePosRewardServiceServer interface {
	mustEmbedUnimplementedPosRewardServiceServer()
}

func RegisterPosRewardServiceServer(s grpc.ServiceRegistrar, srv PosRewardServiceServer) {
	s.RegisterService(&PosRewardService_ServiceDesc, srv)
}

func _PosRewardService_CreatePosReward_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreatePosRewardRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PosRewardServiceServer).CreatePosReward(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pos.PosRewardService/CreatePosReward",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PosRewardServiceServer).CreatePosReward(ctx, req.(*CreatePosRewardRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PosRewardService_ReadPosReward_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReadPosRewardRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PosRewardServiceServer).ReadPosReward(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pos.PosRewardService/ReadPosReward",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PosRewardServiceServer).ReadPosReward(ctx, req.(*ReadPosRewardRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PosRewardService_UpdatePosReward_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdatePosRewardRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PosRewardServiceServer).UpdatePosReward(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pos.PosRewardService/UpdatePosReward",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PosRewardServiceServer).UpdatePosReward(ctx, req.(*UpdatePosRewardRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PosRewardService_DeletePosReward_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeletePosRewardRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PosRewardServiceServer).DeletePosReward(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pos.PosRewardService/DeletePosReward",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PosRewardServiceServer).DeletePosReward(ctx, req.(*DeletePosRewardRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PosRewardService_ReadAllPosRewards_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReadAllPosRewardsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PosRewardServiceServer).ReadAllPosRewards(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pos.PosRewardService/ReadAllPosRewards",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PosRewardServiceServer).ReadAllPosRewards(ctx, req.(*ReadAllPosRewardsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// PosRewardService_ServiceDesc is the grpc.ServiceDesc for PosRewardService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PosRewardService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pos.PosRewardService",
	HandlerType: (*PosRewardServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreatePosReward",
			Handler:    _PosRewardService_CreatePosReward_Handler,
		},
		{
			MethodName: "ReadPosReward",
			Handler:    _PosRewardService_ReadPosReward_Handler,
		},
		{
			MethodName: "UpdatePosReward",
			Handler:    _PosRewardService_UpdatePosReward_Handler,
		},
		{
			MethodName: "DeletePosReward",
			Handler:    _PosRewardService_DeletePosReward_Handler,
		},
		{
			MethodName: "ReadAllPosRewards",
			Handler:    _PosRewardService_ReadAllPosRewards_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "rewards.proto",
}
