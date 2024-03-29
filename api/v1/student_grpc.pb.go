// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.17.3
// source: api/v1/student.proto

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

// StudentServiceClient is the client API for StudentService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type StudentServiceClient interface {
	Create(ctx context.Context, in *CreateStudentRequest, opts ...grpc.CallOption) (*CreateStudentResponse, error)
	Read(ctx context.Context, in *ReadStudentRequest, opts ...grpc.CallOption) (*ReadStudentResponse, error)
	Update(ctx context.Context, in *UpdateStudentRequest, opts ...grpc.CallOption) (*UpdateStudentResponse, error)
	Delete(ctx context.Context, in *DeleteStudentRequest, opts ...grpc.CallOption) (*DeleteStudentResponse, error)
}

type studentServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewStudentServiceClient(cc grpc.ClientConnInterface) StudentServiceClient {
	return &studentServiceClient{cc}
}

func (c *studentServiceClient) Create(ctx context.Context, in *CreateStudentRequest, opts ...grpc.CallOption) (*CreateStudentResponse, error) {
	out := new(CreateStudentResponse)
	err := c.cc.Invoke(ctx, "/student.StudentService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *studentServiceClient) Read(ctx context.Context, in *ReadStudentRequest, opts ...grpc.CallOption) (*ReadStudentResponse, error) {
	out := new(ReadStudentResponse)
	err := c.cc.Invoke(ctx, "/student.StudentService/Read", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *studentServiceClient) Update(ctx context.Context, in *UpdateStudentRequest, opts ...grpc.CallOption) (*UpdateStudentResponse, error) {
	out := new(UpdateStudentResponse)
	err := c.cc.Invoke(ctx, "/student.StudentService/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *studentServiceClient) Delete(ctx context.Context, in *DeleteStudentRequest, opts ...grpc.CallOption) (*DeleteStudentResponse, error) {
	out := new(DeleteStudentResponse)
	err := c.cc.Invoke(ctx, "/student.StudentService/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StudentServiceServer is the server API for StudentService service.
// All implementations must embed UnimplementedStudentServiceServer
// for forward compatibility
type StudentServiceServer interface {
	Create(context.Context, *CreateStudentRequest) (*CreateStudentResponse, error)
	Read(context.Context, *ReadStudentRequest) (*ReadStudentResponse, error)
	Update(context.Context, *UpdateStudentRequest) (*UpdateStudentResponse, error)
	Delete(context.Context, *DeleteStudentRequest) (*DeleteStudentResponse, error)
	mustEmbedUnimplementedStudentServiceServer()
}

// UnimplementedStudentServiceServer must be embedded to have forward compatible implementations.
type UnimplementedStudentServiceServer struct {
}

func (UnimplementedStudentServiceServer) Create(context.Context, *CreateStudentRequest) (*CreateStudentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedStudentServiceServer) Read(context.Context, *ReadStudentRequest) (*ReadStudentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Read not implemented")
}
func (UnimplementedStudentServiceServer) Update(context.Context, *UpdateStudentRequest) (*UpdateStudentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedStudentServiceServer) Delete(context.Context, *DeleteStudentRequest) (*DeleteStudentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedStudentServiceServer) mustEmbedUnimplementedStudentServiceServer() {}

// UnsafeStudentServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to StudentServiceServer will
// result in compilation errors.
type UnsafeStudentServiceServer interface {
	mustEmbedUnimplementedStudentServiceServer()
}

func RegisterStudentServiceServer(s grpc.ServiceRegistrar, srv StudentServiceServer) {
	s.RegisterService(&StudentService_ServiceDesc, srv)
}

func _StudentService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateStudentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StudentServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/student.StudentService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StudentServiceServer).Create(ctx, req.(*CreateStudentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StudentService_Read_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReadStudentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StudentServiceServer).Read(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/student.StudentService/Read",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StudentServiceServer).Read(ctx, req.(*ReadStudentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StudentService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateStudentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StudentServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/student.StudentService/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StudentServiceServer).Update(ctx, req.(*UpdateStudentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StudentService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteStudentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StudentServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/student.StudentService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StudentServiceServer).Delete(ctx, req.(*DeleteStudentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// StudentService_ServiceDesc is the grpc.ServiceDesc for StudentService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var StudentService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "student.StudentService",
	HandlerType: (*StudentServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _StudentService_Create_Handler,
		},
		{
			MethodName: "Read",
			Handler:    _StudentService_Read_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _StudentService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _StudentService_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/v1/student.proto",
}
