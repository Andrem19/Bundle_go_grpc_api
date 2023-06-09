// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.12.4
// source: service.proto

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

const (
	MyBundleProj_CreateUser_FullMethodName = "/pb.MyBundleProj/CreateUser"
	MyBundleProj_LoginUser_FullMethodName  = "/pb.MyBundleProj/LoginUser"
)

// MyBundleProjClient is the client API for MyBundleProj service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MyBundleProjClient interface {
	CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*CreateUserRsponse, error)
	LoginUser(ctx context.Context, in *LoginUserRequest, opts ...grpc.CallOption) (*LoginUserRsponse, error)
}

type myBundleProjClient struct {
	cc grpc.ClientConnInterface
}

func NewMyBundleProjClient(cc grpc.ClientConnInterface) MyBundleProjClient {
	return &myBundleProjClient{cc}
}

func (c *myBundleProjClient) CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*CreateUserRsponse, error) {
	out := new(CreateUserRsponse)
	err := c.cc.Invoke(ctx, MyBundleProj_CreateUser_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *myBundleProjClient) LoginUser(ctx context.Context, in *LoginUserRequest, opts ...grpc.CallOption) (*LoginUserRsponse, error) {
	out := new(LoginUserRsponse)
	err := c.cc.Invoke(ctx, MyBundleProj_LoginUser_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MyBundleProjServer is the server API for MyBundleProj service.
// All implementations must embed UnimplementedMyBundleProjServer
// for forward compatibility
type MyBundleProjServer interface {
	CreateUser(context.Context, *CreateUserRequest) (*CreateUserRsponse, error)
	LoginUser(context.Context, *LoginUserRequest) (*LoginUserRsponse, error)
	mustEmbedUnimplementedMyBundleProjServer()
}

// UnimplementedMyBundleProjServer must be embedded to have forward compatible implementations.
type UnimplementedMyBundleProjServer struct {
}

func (UnimplementedMyBundleProjServer) CreateUser(context.Context, *CreateUserRequest) (*CreateUserRsponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUser not implemented")
}
func (UnimplementedMyBundleProjServer) LoginUser(context.Context, *LoginUserRequest) (*LoginUserRsponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LoginUser not implemented")
}
func (UnimplementedMyBundleProjServer) mustEmbedUnimplementedMyBundleProjServer() {}

// UnsafeMyBundleProjServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MyBundleProjServer will
// result in compilation errors.
type UnsafeMyBundleProjServer interface {
	mustEmbedUnimplementedMyBundleProjServer()
}

func RegisterMyBundleProjServer(s grpc.ServiceRegistrar, srv MyBundleProjServer) {
	s.RegisterService(&MyBundleProj_ServiceDesc, srv)
}

func _MyBundleProj_CreateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MyBundleProjServer).CreateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MyBundleProj_CreateUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MyBundleProjServer).CreateUser(ctx, req.(*CreateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MyBundleProj_LoginUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MyBundleProjServer).LoginUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MyBundleProj_LoginUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MyBundleProjServer).LoginUser(ctx, req.(*LoginUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// MyBundleProj_ServiceDesc is the grpc.ServiceDesc for MyBundleProj service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MyBundleProj_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.MyBundleProj",
	HandlerType: (*MyBundleProjServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateUser",
			Handler:    _MyBundleProj_CreateUser_Handler,
		},
		{
			MethodName: "LoginUser",
			Handler:    _MyBundleProj_LoginUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service.proto",
}
