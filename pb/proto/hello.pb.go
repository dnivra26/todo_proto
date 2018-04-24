// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/hello.proto

/*
Package proto is a generated protocol buffer package.

It is generated from these files:
	proto/hello.proto
	proto/todo.proto

It has these top-level messages:
	PingMessage
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

type PingMessage struct {
	Greeting string `protobuf:"bytes,1,opt,name=greeting" json:"greeting,omitempty"`
}

func (m *PingMessage) Reset()                    { *m = PingMessage{} }
func (m *PingMessage) String() string            { return proto1.CompactTextString(m) }
func (*PingMessage) ProtoMessage()               {}
func (*PingMessage) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *PingMessage) GetGreeting() string {
	if m != nil {
		return m.Greeting
	}
	return ""
}

func init() {
	proto1.RegisterType((*PingMessage)(nil), "proto.PingMessage")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Ping service

type PingClient interface {
	SayHello(ctx context.Context, in *PingMessage, opts ...grpc.CallOption) (*PingMessage, error)
}

type pingClient struct {
	cc *grpc.ClientConn
}

func NewPingClient(cc *grpc.ClientConn) PingClient {
	return &pingClient{cc}
}

func (c *pingClient) SayHello(ctx context.Context, in *PingMessage, opts ...grpc.CallOption) (*PingMessage, error) {
	out := new(PingMessage)
	err := grpc.Invoke(ctx, "/proto.Ping/SayHello", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Ping service

type PingServer interface {
	SayHello(context.Context, *PingMessage) (*PingMessage, error)
}

func RegisterPingServer(s *grpc.Server, srv PingServer) {
	s.RegisterService(&_Ping_serviceDesc, srv)
}

func _Ping_SayHello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PingMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PingServer).SayHello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Ping/SayHello",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PingServer).SayHello(ctx, req.(*PingMessage))
	}
	return interceptor(ctx, in, info, handler)
}

var _Ping_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Ping",
	HandlerType: (*PingServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SayHello",
			Handler:    _Ping_SayHello_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/hello.proto",
}

func init() { proto1.RegisterFile("proto/hello.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 113 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2c, 0x28, 0xca, 0x2f,
	0xc9, 0xd7, 0xcf, 0x48, 0xcd, 0xc9, 0xc9, 0xd7, 0x03, 0xb3, 0x85, 0x58, 0xc1, 0x94, 0x92, 0x26,
	0x17, 0x77, 0x40, 0x66, 0x5e, 0xba, 0x6f, 0x6a, 0x71, 0x71, 0x62, 0x7a, 0xaa, 0x90, 0x14, 0x17,
	0x47, 0x7a, 0x51, 0x6a, 0x6a, 0x49, 0x66, 0x5e, 0xba, 0x04, 0xa3, 0x02, 0xa3, 0x06, 0x67, 0x10,
	0x9c, 0x6f, 0x64, 0xc3, 0xc5, 0x02, 0x52, 0x2a, 0x64, 0xc2, 0xc5, 0x11, 0x9c, 0x58, 0xe9, 0x01,
	0x32, 0x4b, 0x48, 0x08, 0x62, 0x9a, 0x1e, 0x92, 0x19, 0x52, 0x58, 0xc4, 0x94, 0x18, 0x92, 0xd8,
	0xc0, 0x82, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0xc2, 0x45, 0x36, 0x63, 0x8b, 0x00, 0x00,
	0x00,
}
