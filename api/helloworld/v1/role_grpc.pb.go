// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.23.3
// source: api/helloworld/v1/role.proto

package v1

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
	Role_ListPage_FullMethodName         = "/api.helloworld.v1.Role/ListPage"
	Role_Add_FullMethodName              = "/api.helloworld.v1.Role/Add"
	Role_GetByCode_FullMethodName        = "/api.helloworld.v1.Role/GetByCode"
	Role_AssignPermission_FullMethodName = "/api.helloworld.v1.Role/AssignPermission"
	Role_DeleteById_FullMethodName       = "/api.helloworld.v1.Role/DeleteById"
)

// RoleClient is the client API for Role service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RoleClient interface {
	ListPage(ctx context.Context, in *RoleListRequest, opts ...grpc.CallOption) (*RoleListReply, error)
	Add(ctx context.Context, in *RoleSaveRequest, opts ...grpc.CallOption) (*RoleReply, error)
	GetByCode(ctx context.Context, in *RoleCodeRequest, opts ...grpc.CallOption) (*RoleReply, error)
	AssignPermission(ctx context.Context, in *RoleAssignPermRequest, opts ...grpc.CallOption) (*RoleOk, error)
	DeleteById(ctx context.Context, in *RoleDeleteRequest, opts ...grpc.CallOption) (*RoleOk, error)
}

type roleClient struct {
	cc grpc.ClientConnInterface
}

func NewRoleClient(cc grpc.ClientConnInterface) RoleClient {
	return &roleClient{cc}
}

func (c *roleClient) ListPage(ctx context.Context, in *RoleListRequest, opts ...grpc.CallOption) (*RoleListReply, error) {
	out := new(RoleListReply)
	err := c.cc.Invoke(ctx, Role_ListPage_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *roleClient) Add(ctx context.Context, in *RoleSaveRequest, opts ...grpc.CallOption) (*RoleReply, error) {
	out := new(RoleReply)
	err := c.cc.Invoke(ctx, Role_Add_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *roleClient) GetByCode(ctx context.Context, in *RoleCodeRequest, opts ...grpc.CallOption) (*RoleReply, error) {
	out := new(RoleReply)
	err := c.cc.Invoke(ctx, Role_GetByCode_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *roleClient) AssignPermission(ctx context.Context, in *RoleAssignPermRequest, opts ...grpc.CallOption) (*RoleOk, error) {
	out := new(RoleOk)
	err := c.cc.Invoke(ctx, Role_AssignPermission_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *roleClient) DeleteById(ctx context.Context, in *RoleDeleteRequest, opts ...grpc.CallOption) (*RoleOk, error) {
	out := new(RoleOk)
	err := c.cc.Invoke(ctx, Role_DeleteById_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RoleServer is the server API for Role service.
// All implementations must embed UnimplementedRoleServer
// for forward compatibility
type RoleServer interface {
	ListPage(context.Context, *RoleListRequest) (*RoleListReply, error)
	Add(context.Context, *RoleSaveRequest) (*RoleReply, error)
	GetByCode(context.Context, *RoleCodeRequest) (*RoleReply, error)
	AssignPermission(context.Context, *RoleAssignPermRequest) (*RoleOk, error)
	DeleteById(context.Context, *RoleDeleteRequest) (*RoleOk, error)
	mustEmbedUnimplementedRoleServer()
}

// UnimplementedRoleServer must be embedded to have forward compatible implementations.
type UnimplementedRoleServer struct {
}

func (UnimplementedRoleServer) ListPage(context.Context, *RoleListRequest) (*RoleListReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListPage not implemented")
}
func (UnimplementedRoleServer) Add(context.Context, *RoleSaveRequest) (*RoleReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Add not implemented")
}
func (UnimplementedRoleServer) GetByCode(context.Context, *RoleCodeRequest) (*RoleReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetByCode not implemented")
}
func (UnimplementedRoleServer) AssignPermission(context.Context, *RoleAssignPermRequest) (*RoleOk, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AssignPermission not implemented")
}
func (UnimplementedRoleServer) DeleteById(context.Context, *RoleDeleteRequest) (*RoleOk, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteById not implemented")
}
func (UnimplementedRoleServer) mustEmbedUnimplementedRoleServer() {}

// UnsafeRoleServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RoleServer will
// result in compilation errors.
type UnsafeRoleServer interface {
	mustEmbedUnimplementedRoleServer()
}

func RegisterRoleServer(s grpc.ServiceRegistrar, srv RoleServer) {
	s.RegisterService(&Role_ServiceDesc, srv)
}

func _Role_ListPage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RoleListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoleServer).ListPage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Role_ListPage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoleServer).ListPage(ctx, req.(*RoleListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Role_Add_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RoleSaveRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoleServer).Add(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Role_Add_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoleServer).Add(ctx, req.(*RoleSaveRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Role_GetByCode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RoleCodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoleServer).GetByCode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Role_GetByCode_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoleServer).GetByCode(ctx, req.(*RoleCodeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Role_AssignPermission_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RoleAssignPermRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoleServer).AssignPermission(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Role_AssignPermission_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoleServer).AssignPermission(ctx, req.(*RoleAssignPermRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Role_DeleteById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RoleDeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoleServer).DeleteById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Role_DeleteById_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoleServer).DeleteById(ctx, req.(*RoleDeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Role_ServiceDesc is the grpc.ServiceDesc for Role service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Role_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.helloworld.v1.Role",
	HandlerType: (*RoleServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListPage",
			Handler:    _Role_ListPage_Handler,
		},
		{
			MethodName: "Add",
			Handler:    _Role_Add_Handler,
		},
		{
			MethodName: "GetByCode",
			Handler:    _Role_GetByCode_Handler,
		},
		{
			MethodName: "AssignPermission",
			Handler:    _Role_AssignPermission_Handler,
		},
		{
			MethodName: "DeleteById",
			Handler:    _Role_DeleteById_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/helloworld/v1/role.proto",
}