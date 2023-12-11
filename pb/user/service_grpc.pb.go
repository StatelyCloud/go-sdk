// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.1
// source: user/service.proto

package user

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
	User_Whoami_FullMethodName            = "/stately.User/Whoami"
	User_Enroll_FullMethodName            = "/stately.User/Enroll"
	User_EnrollMachineUser_FullMethodName = "/stately.User/EnrollMachineUser"
	User_CreateProject_FullMethodName     = "/stately.User/CreateProject"
	User_DescribeProject_FullMethodName   = "/stately.User/DescribeProject"
	User_ListProjects_FullMethodName      = "/stately.User/ListProjects"
	User_DeleteProject_FullMethodName     = "/stately.User/DeleteProject"
)

// UserClient is the client API for User service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserClient interface {
	// Whoami returns information about the user that calls it (based on the auth
	// token). This includes information about what organizations the user
	// belongs to, what projects they have access to, what roles(?) they can use,
	// etc. This is meant to be called from the Web Console or CLI in order to
	// populate some basic information in the UI and allow calling other APIs like
	// ListStores.
	Whoami(ctx context.Context, in *WhoamiRequest, opts ...grpc.CallOption) (*WhoamiResponse, error)
	// Enroll bootstraps a new User given a service principal ID from an auth
	// provider. This includes creating a user record for them, and a default
	// organization, project, and store for them to use. User information is
	// automatically read from the auth token.
	Enroll(ctx context.Context, in *EnrollRequest, opts ...grpc.CallOption) (*EnrollResponse, error)
	// EnrollMachineUser bootstraps a new machine user principal ID from an auth provider
	// and enrolls them in the orgnanization ID which was passed in the request.
	// Subsequent calls to this API will enroll the user in the orginization if a different
	// organization ID is passed, otherwise they will be a no-op.
	// ** THIS IS AN ADMIN ONLY API **
	EnrollMachineUser(ctx context.Context, in *EnrollMachineUserRequest, opts ...grpc.CallOption) (*EnrollMachineUserResponse, error)
	// CreateProject makes a new project within your organization. It will fail if
	// the project already exists or you don't have permission to create projects
	// in that organization.
	CreateProject(ctx context.Context, in *CreateProjectRequest, opts ...grpc.CallOption) (*CreateProjectResponse, error)
	// DescribeProject gets information about a project. It will fail if the project
	// does not exist or you don't have permission to describe projects in this
	// organization.
	DescribeProject(ctx context.Context, in *DescribeProjectRequest, opts ...grpc.CallOption) (*DescribeProjectResponse, error)
	// ListProjects lists all the projects in the selected organization. It will fail
	// if you don't have permission to list projects in this organization. You will
	// need to call DescribeProject on each project to get full details about it.
	ListProjects(ctx context.Context, in *ListProjectsRequest, opts ...grpc.CallOption) (*ListProjectsResponse, error)
	// DeleteProject schedules a project to be deleted, including all data within it.
	// This operation takes some time so it returns a handle to an operation that
	// you can check to see if it is complete. This will fail if the project does
	// not exist, if the project is already being deleted, or if you do not have
	// permission to delete project.
	DeleteProject(ctx context.Context, in *DeleteProjectRequest, opts ...grpc.CallOption) (*DeleteProjectResponse, error)
}

type userClient struct {
	cc grpc.ClientConnInterface
}

func NewUserClient(cc grpc.ClientConnInterface) UserClient {
	return &userClient{cc}
}

func (c *userClient) Whoami(ctx context.Context, in *WhoamiRequest, opts ...grpc.CallOption) (*WhoamiResponse, error) {
	out := new(WhoamiResponse)
	err := c.cc.Invoke(ctx, User_Whoami_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) Enroll(ctx context.Context, in *EnrollRequest, opts ...grpc.CallOption) (*EnrollResponse, error) {
	out := new(EnrollResponse)
	err := c.cc.Invoke(ctx, User_Enroll_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) EnrollMachineUser(ctx context.Context, in *EnrollMachineUserRequest, opts ...grpc.CallOption) (*EnrollMachineUserResponse, error) {
	out := new(EnrollMachineUserResponse)
	err := c.cc.Invoke(ctx, User_EnrollMachineUser_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) CreateProject(ctx context.Context, in *CreateProjectRequest, opts ...grpc.CallOption) (*CreateProjectResponse, error) {
	out := new(CreateProjectResponse)
	err := c.cc.Invoke(ctx, User_CreateProject_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) DescribeProject(ctx context.Context, in *DescribeProjectRequest, opts ...grpc.CallOption) (*DescribeProjectResponse, error) {
	out := new(DescribeProjectResponse)
	err := c.cc.Invoke(ctx, User_DescribeProject_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) ListProjects(ctx context.Context, in *ListProjectsRequest, opts ...grpc.CallOption) (*ListProjectsResponse, error) {
	out := new(ListProjectsResponse)
	err := c.cc.Invoke(ctx, User_ListProjects_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) DeleteProject(ctx context.Context, in *DeleteProjectRequest, opts ...grpc.CallOption) (*DeleteProjectResponse, error) {
	out := new(DeleteProjectResponse)
	err := c.cc.Invoke(ctx, User_DeleteProject_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServer is the server API for User service.
// All implementations should embed UnimplementedUserServer
// for forward compatibility
type UserServer interface {
	// Whoami returns information about the user that calls it (based on the auth
	// token). This includes information about what organizations the user
	// belongs to, what projects they have access to, what roles(?) they can use,
	// etc. This is meant to be called from the Web Console or CLI in order to
	// populate some basic information in the UI and allow calling other APIs like
	// ListStores.
	Whoami(context.Context, *WhoamiRequest) (*WhoamiResponse, error)
	// Enroll bootstraps a new User given a service principal ID from an auth
	// provider. This includes creating a user record for them, and a default
	// organization, project, and store for them to use. User information is
	// automatically read from the auth token.
	Enroll(context.Context, *EnrollRequest) (*EnrollResponse, error)
	// EnrollMachineUser bootstraps a new machine user principal ID from an auth provider
	// and enrolls them in the orgnanization ID which was passed in the request.
	// Subsequent calls to this API will enroll the user in the orginization if a different
	// organization ID is passed, otherwise they will be a no-op.
	// ** THIS IS AN ADMIN ONLY API **
	EnrollMachineUser(context.Context, *EnrollMachineUserRequest) (*EnrollMachineUserResponse, error)
	// CreateProject makes a new project within your organization. It will fail if
	// the project already exists or you don't have permission to create projects
	// in that organization.
	CreateProject(context.Context, *CreateProjectRequest) (*CreateProjectResponse, error)
	// DescribeProject gets information about a project. It will fail if the project
	// does not exist or you don't have permission to describe projects in this
	// organization.
	DescribeProject(context.Context, *DescribeProjectRequest) (*DescribeProjectResponse, error)
	// ListProjects lists all the projects in the selected organization. It will fail
	// if you don't have permission to list projects in this organization. You will
	// need to call DescribeProject on each project to get full details about it.
	ListProjects(context.Context, *ListProjectsRequest) (*ListProjectsResponse, error)
	// DeleteProject schedules a project to be deleted, including all data within it.
	// This operation takes some time so it returns a handle to an operation that
	// you can check to see if it is complete. This will fail if the project does
	// not exist, if the project is already being deleted, or if you do not have
	// permission to delete project.
	DeleteProject(context.Context, *DeleteProjectRequest) (*DeleteProjectResponse, error)
}

// UnimplementedUserServer should be embedded to have forward compatible implementations.
type UnimplementedUserServer struct {
}

func (UnimplementedUserServer) Whoami(context.Context, *WhoamiRequest) (*WhoamiResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Whoami not implemented")
}
func (UnimplementedUserServer) Enroll(context.Context, *EnrollRequest) (*EnrollResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Enroll not implemented")
}
func (UnimplementedUserServer) EnrollMachineUser(context.Context, *EnrollMachineUserRequest) (*EnrollMachineUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EnrollMachineUser not implemented")
}
func (UnimplementedUserServer) CreateProject(context.Context, *CreateProjectRequest) (*CreateProjectResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateProject not implemented")
}
func (UnimplementedUserServer) DescribeProject(context.Context, *DescribeProjectRequest) (*DescribeProjectResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DescribeProject not implemented")
}
func (UnimplementedUserServer) ListProjects(context.Context, *ListProjectsRequest) (*ListProjectsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListProjects not implemented")
}
func (UnimplementedUserServer) DeleteProject(context.Context, *DeleteProjectRequest) (*DeleteProjectResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteProject not implemented")
}

// UnsafeUserServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserServer will
// result in compilation errors.
type UnsafeUserServer interface {
	mustEmbedUnimplementedUserServer()
}

func RegisterUserServer(s grpc.ServiceRegistrar, srv UserServer) {
	s.RegisterService(&User_ServiceDesc, srv)
}

func _User_Whoami_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WhoamiRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).Whoami(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: User_Whoami_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).Whoami(ctx, req.(*WhoamiRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_Enroll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EnrollRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).Enroll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: User_Enroll_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).Enroll(ctx, req.(*EnrollRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_EnrollMachineUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EnrollMachineUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).EnrollMachineUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: User_EnrollMachineUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).EnrollMachineUser(ctx, req.(*EnrollMachineUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_CreateProject_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateProjectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).CreateProject(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: User_CreateProject_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).CreateProject(ctx, req.(*CreateProjectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_DescribeProject_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DescribeProjectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).DescribeProject(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: User_DescribeProject_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).DescribeProject(ctx, req.(*DescribeProjectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_ListProjects_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListProjectsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).ListProjects(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: User_ListProjects_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).ListProjects(ctx, req.(*ListProjectsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_DeleteProject_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteProjectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).DeleteProject(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: User_DeleteProject_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).DeleteProject(ctx, req.(*DeleteProjectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// User_ServiceDesc is the grpc.ServiceDesc for User service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var User_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "stately.User",
	HandlerType: (*UserServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Whoami",
			Handler:    _User_Whoami_Handler,
		},
		{
			MethodName: "Enroll",
			Handler:    _User_Enroll_Handler,
		},
		{
			MethodName: "EnrollMachineUser",
			Handler:    _User_EnrollMachineUser_Handler,
		},
		{
			MethodName: "CreateProject",
			Handler:    _User_CreateProject_Handler,
		},
		{
			MethodName: "DescribeProject",
			Handler:    _User_DescribeProject_Handler,
		},
		{
			MethodName: "ListProjects",
			Handler:    _User_ListProjects_Handler,
		},
		{
			MethodName: "DeleteProject",
			Handler:    _User_DeleteProject_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user/service.proto",
}
