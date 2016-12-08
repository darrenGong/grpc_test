// Code generated by protoc-gen-go.
// source: hello.proto
// DO NOT EDIT!

/*
Package helloworld is a generated protocol buffer package.

It is generated from these files:
	hello.proto

It has these top-level messages:
	ResponseCode
	HelloRequest
	HelloResponse
*/
package helloworld

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type ResponseCode struct {
	ErrCode    int32  `protobuf:"varint,10,opt,name=errCode" json:"errCode,omitempty"`
	ErrMessage string `protobuf:"bytes,20,opt,name=errMessage" json:"errMessage,omitempty"`
}

func (m *ResponseCode) Reset()                    { *m = ResponseCode{} }
func (m *ResponseCode) String() string            { return proto.CompactTextString(m) }
func (*ResponseCode) ProtoMessage()               {}
func (*ResponseCode) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *ResponseCode) GetErrCode() int32 {
	if m != nil {
		return m.ErrCode
	}
	return 0
}

func (m *ResponseCode) GetErrMessage() string {
	if m != nil {
		return m.ErrMessage
	}
	return ""
}

type HelloRequest struct {
	Name string `protobuf:"bytes,10,opt,name=name" json:"name,omitempty"`
}

func (m *HelloRequest) Reset()                    { *m = HelloRequest{} }
func (m *HelloRequest) String() string            { return proto.CompactTextString(m) }
func (*HelloRequest) ProtoMessage()               {}
func (*HelloRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *HelloRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type HelloResponse struct {
	Rc      *ResponseCode `protobuf:"bytes,10,opt,name=rc" json:"rc,omitempty"`
	Message string        `protobuf:"bytes,20,opt,name=message" json:"message,omitempty"`
}

func (m *HelloResponse) Reset()                    { *m = HelloResponse{} }
func (m *HelloResponse) String() string            { return proto.CompactTextString(m) }
func (*HelloResponse) ProtoMessage()               {}
func (*HelloResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *HelloResponse) GetRc() *ResponseCode {
	if m != nil {
		return m.Rc
	}
	return nil
}

func (m *HelloResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*ResponseCode)(nil), "helloworld.ResponseCode")
	proto.RegisterType((*HelloRequest)(nil), "helloworld.HelloRequest")
	proto.RegisterType((*HelloResponse)(nil), "helloworld.HelloResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Greeter service

type GreeterClient interface {
	SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloResponse, error)
}

type greeterClient struct {
	cc *grpc.ClientConn
}

func NewGreeterClient(cc *grpc.ClientConn) GreeterClient {
	return &greeterClient{cc}
}

func (c *greeterClient) SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloResponse, error) {
	out := new(HelloResponse)
	err := grpc.Invoke(ctx, "/helloworld.Greeter/SayHello", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Greeter service

type GreeterServer interface {
	SayHello(context.Context, *HelloRequest) (*HelloResponse, error)
}

func RegisterGreeterServer(s *grpc.Server, srv GreeterServer) {
	s.RegisterService(&_Greeter_serviceDesc, srv)
}

func _Greeter_SayHello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GreeterServer).SayHello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/helloworld.Greeter/SayHello",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GreeterServer).SayHello(ctx, req.(*HelloRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Greeter_serviceDesc = grpc.ServiceDesc{
	ServiceName: "helloworld.Greeter",
	HandlerType: (*GreeterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SayHello",
			Handler:    _Greeter_SayHello_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "hello.proto",
}

func init() { proto.RegisterFile("hello.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 205 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x6c, 0x90, 0xb1, 0x6e, 0x83, 0x30,
	0x10, 0x86, 0x0b, 0x6a, 0x4b, 0x39, 0xe8, 0x72, 0xea, 0xe0, 0x76, 0xa8, 0x90, 0x27, 0x26, 0x06,
	0xf2, 0x04, 0x51, 0x86, 0x30, 0x24, 0x8b, 0x79, 0x02, 0x02, 0xa7, 0x64, 0x00, 0x4c, 0xce, 0x44,
	0x51, 0xde, 0x3e, 0xc2, 0x04, 0x05, 0xa4, 0x6c, 0xf7, 0xdb, 0xa7, 0xff, 0xfb, 0x74, 0x10, 0x9c,
	0xa8, 0xae, 0x75, 0xd2, 0xb1, 0xee, 0x35, 0x82, 0x0d, 0x57, 0xcd, 0x75, 0x25, 0x33, 0x08, 0x15,
	0x99, 0x4e, 0xb7, 0x86, 0x36, 0xba, 0x22, 0x14, 0xe0, 0x11, 0xf3, 0x30, 0x0a, 0x88, 0x9c, 0xf8,
	0x43, 0x4d, 0x11, 0xff, 0x01, 0x88, 0x79, 0x4f, 0xc6, 0x14, 0x47, 0x12, 0x3f, 0x91, 0x13, 0xfb,
	0x6a, 0xf6, 0x22, 0x25, 0x84, 0xd9, 0xd0, 0xab, 0xe8, 0x7c, 0x21, 0xd3, 0x23, 0xc2, 0x7b, 0x5b,
	0x34, 0x63, 0x8d, 0xaf, 0xec, 0x2c, 0x73, 0xf8, 0x7e, 0xec, 0x8c, 0x48, 0x8c, 0xc1, 0xe5, 0xd2,
	0xae, 0x04, 0xa9, 0x48, 0x9e, 0x5e, 0xc9, 0x5c, 0x4a, 0xb9, 0x5c, 0x0e, 0x62, 0xcd, 0x82, 0x3d,
	0xc5, 0x74, 0x07, 0xde, 0x96, 0x89, 0x7a, 0x62, 0x5c, 0xc3, 0x57, 0x5e, 0xdc, 0x2c, 0x02, 0x17,
	0x75, 0x73, 0xb3, 0xbf, 0xdf, 0x17, 0x3f, 0x23, 0x4d, 0xbe, 0x1d, 0x3e, 0xed, 0x8d, 0x56, 0xf7,
	0x00, 0x00, 0x00, 0xff, 0xff, 0x0d, 0x51, 0x1b, 0xdc, 0x32, 0x01, 0x00, 0x00,
}
