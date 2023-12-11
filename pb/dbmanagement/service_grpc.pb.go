// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.1
// source: dbmanagement/service.proto

package dbmanagement

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
	Management_CreateStore_FullMethodName   = "/stately.Management/CreateStore"
	Management_DescribeStore_FullMethodName = "/stately.Management/DescribeStore"
	Management_ListStores_FullMethodName    = "/stately.Management/ListStores"
	Management_DeleteStore_FullMethodName   = "/stately.Management/DeleteStore"
)

// ManagementClient is the client API for Management service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ManagementClient interface {
	// CreateStore makes a new store within your project. It will fail if the
	// store already exists or you don't have permission to create stores in that
	// project.
	CreateStore(ctx context.Context, in *CreateStoreRequest, opts ...grpc.CallOption) (*CreateStoreResponse, error)
	// DescribeStore gets information about a store, including its current schema
	// (?). It will fail if the store does not exist or you don't have permission
	// to describe stores in this project.
	DescribeStore(ctx context.Context, in *DescribeStoreRequest, opts ...grpc.CallOption) (*DescribeStoreResponse, error)
	// ListStores lists all the stores in the selected project. It will fail if
	// you don't have permission to list stores in this project. You will need to
	// call DescribeStore on each store to get full details about it.
	ListStores(ctx context.Context, in *ListStoresRequest, opts ...grpc.CallOption) (*ListStoresResponse, error)
	// DeleteStore schedules a store to be deleted, including all data within it.
	// This operation takes some time so it returns a handle to an operation that
	// you can check to see if it is complete. This will fail if the store does
	// not exist, if the store is already being deleted, or if you do not have
	// permission to delete stores.
	DeleteStore(ctx context.Context, in *DeleteStoreRequest, opts ...grpc.CallOption) (*DeleteStoreResponse, error)
}

type managementClient struct {
	cc grpc.ClientConnInterface
}

func NewManagementClient(cc grpc.ClientConnInterface) ManagementClient {
	return &managementClient{cc}
}

func (c *managementClient) CreateStore(ctx context.Context, in *CreateStoreRequest, opts ...grpc.CallOption) (*CreateStoreResponse, error) {
	out := new(CreateStoreResponse)
	err := c.cc.Invoke(ctx, Management_CreateStore_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managementClient) DescribeStore(ctx context.Context, in *DescribeStoreRequest, opts ...grpc.CallOption) (*DescribeStoreResponse, error) {
	out := new(DescribeStoreResponse)
	err := c.cc.Invoke(ctx, Management_DescribeStore_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managementClient) ListStores(ctx context.Context, in *ListStoresRequest, opts ...grpc.CallOption) (*ListStoresResponse, error) {
	out := new(ListStoresResponse)
	err := c.cc.Invoke(ctx, Management_ListStores_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managementClient) DeleteStore(ctx context.Context, in *DeleteStoreRequest, opts ...grpc.CallOption) (*DeleteStoreResponse, error) {
	out := new(DeleteStoreResponse)
	err := c.cc.Invoke(ctx, Management_DeleteStore_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ManagementServer is the server API for Management service.
// All implementations should embed UnimplementedManagementServer
// for forward compatibility
type ManagementServer interface {
	// CreateStore makes a new store within your project. It will fail if the
	// store already exists or you don't have permission to create stores in that
	// project.
	CreateStore(context.Context, *CreateStoreRequest) (*CreateStoreResponse, error)
	// DescribeStore gets information about a store, including its current schema
	// (?). It will fail if the store does not exist or you don't have permission
	// to describe stores in this project.
	DescribeStore(context.Context, *DescribeStoreRequest) (*DescribeStoreResponse, error)
	// ListStores lists all the stores in the selected project. It will fail if
	// you don't have permission to list stores in this project. You will need to
	// call DescribeStore on each store to get full details about it.
	ListStores(context.Context, *ListStoresRequest) (*ListStoresResponse, error)
	// DeleteStore schedules a store to be deleted, including all data within it.
	// This operation takes some time so it returns a handle to an operation that
	// you can check to see if it is complete. This will fail if the store does
	// not exist, if the store is already being deleted, or if you do not have
	// permission to delete stores.
	DeleteStore(context.Context, *DeleteStoreRequest) (*DeleteStoreResponse, error)
}

// UnimplementedManagementServer should be embedded to have forward compatible implementations.
type UnimplementedManagementServer struct {
}

func (UnimplementedManagementServer) CreateStore(context.Context, *CreateStoreRequest) (*CreateStoreResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateStore not implemented")
}
func (UnimplementedManagementServer) DescribeStore(context.Context, *DescribeStoreRequest) (*DescribeStoreResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DescribeStore not implemented")
}
func (UnimplementedManagementServer) ListStores(context.Context, *ListStoresRequest) (*ListStoresResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListStores not implemented")
}
func (UnimplementedManagementServer) DeleteStore(context.Context, *DeleteStoreRequest) (*DeleteStoreResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteStore not implemented")
}

// UnsafeManagementServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ManagementServer will
// result in compilation errors.
type UnsafeManagementServer interface {
	mustEmbedUnimplementedManagementServer()
}

func RegisterManagementServer(s grpc.ServiceRegistrar, srv ManagementServer) {
	s.RegisterService(&Management_ServiceDesc, srv)
}

func _Management_CreateStore_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateStoreRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagementServer).CreateStore(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Management_CreateStore_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagementServer).CreateStore(ctx, req.(*CreateStoreRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Management_DescribeStore_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DescribeStoreRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagementServer).DescribeStore(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Management_DescribeStore_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagementServer).DescribeStore(ctx, req.(*DescribeStoreRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Management_ListStores_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListStoresRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagementServer).ListStores(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Management_ListStores_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagementServer).ListStores(ctx, req.(*ListStoresRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Management_DeleteStore_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteStoreRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagementServer).DeleteStore(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Management_DeleteStore_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagementServer).DeleteStore(ctx, req.(*DeleteStoreRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Management_ServiceDesc is the grpc.ServiceDesc for Management service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Management_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "stately.Management",
	HandlerType: (*ManagementServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateStore",
			Handler:    _Management_CreateStore_Handler,
		},
		{
			MethodName: "DescribeStore",
			Handler:    _Management_DescribeStore_Handler,
		},
		{
			MethodName: "ListStores",
			Handler:    _Management_ListStores_Handler,
		},
		{
			MethodName: "DeleteStore",
			Handler:    _Management_DeleteStore_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "dbmanagement/service.proto",
}
