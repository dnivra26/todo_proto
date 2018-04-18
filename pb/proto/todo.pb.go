// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/todo.proto

/*
Package proto is a generated protocol buffer package.

It is generated from these files:
	proto/todo.proto

It has these top-level messages:
	Todo
	CreateTodoRequest
	CreateTodoResponse
	GetTodoRequest
	GetTodoResponse
*/
package proto

import proto1 "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto1.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto1.ProtoPackageIsVersion2 // please upgrade the proto package

type Todo struct {
	Id     string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Title  string `protobuf:"bytes,2,opt,name=title" json:"title,omitempty"`
	Status bool   `protobuf:"varint,3,opt,name=status" json:"status,omitempty"`
}

func (m *Todo) Reset()                    { *m = Todo{} }
func (m *Todo) String() string            { return proto1.CompactTextString(m) }
func (*Todo) ProtoMessage()               {}
func (*Todo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Todo) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Todo) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *Todo) GetStatus() bool {
	if m != nil {
		return m.Status
	}
	return false
}

type CreateTodoRequest struct {
	Todo *Todo `protobuf:"bytes,1,opt,name=todo" json:"todo,omitempty"`
}

func (m *CreateTodoRequest) Reset()                    { *m = CreateTodoRequest{} }
func (m *CreateTodoRequest) String() string            { return proto1.CompactTextString(m) }
func (*CreateTodoRequest) ProtoMessage()               {}
func (*CreateTodoRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *CreateTodoRequest) GetTodo() *Todo {
	if m != nil {
		return m.Todo
	}
	return nil
}

type CreateTodoResponse struct {
	Id string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
}

func (m *CreateTodoResponse) Reset()                    { *m = CreateTodoResponse{} }
func (m *CreateTodoResponse) String() string            { return proto1.CompactTextString(m) }
func (*CreateTodoResponse) ProtoMessage()               {}
func (*CreateTodoResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *CreateTodoResponse) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type GetTodoRequest struct {
	Id string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
}

func (m *GetTodoRequest) Reset()                    { *m = GetTodoRequest{} }
func (m *GetTodoRequest) String() string            { return proto1.CompactTextString(m) }
func (*GetTodoRequest) ProtoMessage()               {}
func (*GetTodoRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *GetTodoRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type GetTodoResponse struct {
	Todo *Todo `protobuf:"bytes,1,opt,name=todo" json:"todo,omitempty"`
}

func (m *GetTodoResponse) Reset()                    { *m = GetTodoResponse{} }
func (m *GetTodoResponse) String() string            { return proto1.CompactTextString(m) }
func (*GetTodoResponse) ProtoMessage()               {}
func (*GetTodoResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *GetTodoResponse) GetTodo() *Todo {
	if m != nil {
		return m.Todo
	}
	return nil
}

func init() {
	proto1.RegisterType((*Todo)(nil), "proto.Todo")
	proto1.RegisterType((*CreateTodoRequest)(nil), "proto.CreateTodoRequest")
	proto1.RegisterType((*CreateTodoResponse)(nil), "proto.CreateTodoResponse")
	proto1.RegisterType((*GetTodoRequest)(nil), "proto.GetTodoRequest")
	proto1.RegisterType((*GetTodoResponse)(nil), "proto.GetTodoResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for TodoService service

type TodoServiceClient interface {
	CreateTodo(ctx context.Context, in *CreateTodoRequest, opts ...grpc.CallOption) (*CreateTodoResponse, error)
	GetTodo(ctx context.Context, in *GetTodoRequest, opts ...grpc.CallOption) (*GetTodoResponse, error)
}

type todoServiceClient struct {
	cc *grpc.ClientConn
}

func NewTodoServiceClient(cc *grpc.ClientConn) TodoServiceClient {
	return &todoServiceClient{cc}
}

func (c *todoServiceClient) CreateTodo(ctx context.Context, in *CreateTodoRequest, opts ...grpc.CallOption) (*CreateTodoResponse, error) {
	out := new(CreateTodoResponse)
	err := grpc.Invoke(ctx, "/proto.TodoService/CreateTodo", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *todoServiceClient) GetTodo(ctx context.Context, in *GetTodoRequest, opts ...grpc.CallOption) (*GetTodoResponse, error) {
	out := new(GetTodoResponse)
	err := grpc.Invoke(ctx, "/proto.TodoService/GetTodo", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for TodoService service

type TodoServiceServer interface {
	CreateTodo(context.Context, *CreateTodoRequest) (*CreateTodoResponse, error)
	GetTodo(context.Context, *GetTodoRequest) (*GetTodoResponse, error)
}

func RegisterTodoServiceServer(s *grpc.Server, srv TodoServiceServer) {
	s.RegisterService(&_TodoService_serviceDesc, srv)
}

func _TodoService_CreateTodo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTodoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TodoServiceServer).CreateTodo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.TodoService/CreateTodo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TodoServiceServer).CreateTodo(ctx, req.(*CreateTodoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TodoService_GetTodo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTodoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TodoServiceServer).GetTodo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.TodoService/GetTodo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TodoServiceServer).GetTodo(ctx, req.(*GetTodoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _TodoService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.TodoService",
	HandlerType: (*TodoServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateTodo",
			Handler:    _TodoService_CreateTodo_Handler,
		},
		{
			MethodName: "GetTodo",
			Handler:    _TodoService_GetTodo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/todo.proto",
}

func init() { proto1.RegisterFile("proto/todo.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 226 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x4f, 0xbf, 0x4e, 0x86, 0x30,
	0x10, 0xb7, 0xf8, 0x7d, 0xa8, 0x47, 0x82, 0x7a, 0x51, 0x52, 0x59, 0x24, 0x8d, 0x03, 0x13, 0x26,
	0xe8, 0xe4, 0x8a, 0x89, 0x3b, 0xfa, 0x02, 0x68, 0x6f, 0x68, 0x62, 0x2c, 0xd2, 0xc3, 0xc7, 0xf0,
	0x99, 0x0d, 0xa5, 0x2a, 0xa2, 0x89, 0x13, 0xfc, 0xee, 0xf7, 0xb7, 0x70, 0xd4, 0x0f, 0x96, 0xed,
	0x25, 0x5b, 0x6d, 0x2b, 0xff, 0x8b, 0x5b, 0xff, 0x51, 0xb7, 0xb0, 0x79, 0xb0, 0xda, 0x62, 0x0a,
	0x91, 0xd1, 0x52, 0x14, 0xa2, 0x3c, 0x68, 0x23, 0xa3, 0xf1, 0x04, 0xb6, 0x6c, 0xf8, 0x99, 0x64,
	0xe4, 0x4f, 0x33, 0xc0, 0x0c, 0x62, 0xc7, 0x1d, 0x8f, 0x4e, 0xee, 0x16, 0xa2, 0xdc, 0x6f, 0x03,
	0x52, 0xd7, 0x70, 0xdc, 0x0c, 0xd4, 0x31, 0x4d, 0x59, 0x2d, 0xbd, 0x8e, 0xe4, 0x18, 0xcf, 0x61,
	0x33, 0xf5, 0xf9, 0xd0, 0xa4, 0x4e, 0xe6, 0xde, 0xca, 0x2b, 0x3c, 0xa1, 0x2e, 0x00, 0x97, 0x2e,
	0xd7, 0xdb, 0x17, 0x47, 0xeb, 0x25, 0xaa, 0x80, 0xf4, 0x8e, 0x78, 0x19, 0xbc, 0x56, 0xd4, 0x70,
	0xf8, 0xa5, 0x08, 0x21, 0xff, 0x75, 0xd7, 0xef, 0x02, 0x92, 0x09, 0xde, 0xd3, 0xf0, 0x66, 0x9e,
	0x08, 0x1b, 0x80, 0xef, 0x2d, 0x28, 0x83, 0xe1, 0xd7, 0xa3, 0xf2, 0xb3, 0x3f, 0x98, 0xb9, 0x53,
	0xed, 0xe0, 0x0d, 0xec, 0x85, 0x21, 0x78, 0x1a, 0x74, 0x3f, 0xa7, 0xe7, 0xd9, 0xfa, 0xfc, 0xe9,
	0x7d, 0x8c, 0x3d, 0x71, 0xf5, 0x11, 0x00, 0x00, 0xff, 0xff, 0x2a, 0x62, 0xa4, 0xb2, 0xaa, 0x01,
	0x00, 0x00,
}
