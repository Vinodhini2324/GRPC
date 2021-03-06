// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/micro/go-grpc/examples/greeter/function/proto/hello/hello.proto

/*
Package go_micro_srv_greeter is a generated protocol buffer package.
It is generated from these files:
	github.com/micro/go-grpc/examples/greeter/function/proto/hello/hello.proto
It has these top-level messages:
	Request
	Response
*/
package go_micro_srv_greeter

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

type Request struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *Request) Reset()                    { *m = Request{} }
func (m *Request) String() string            { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()               {}
func (*Request) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Request) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type Response struct {
	Msg string `protobuf:"bytes,1,opt,name=msg" json:"msg,omitempty"`
}

func (m *Response) Reset()                    { *m = Response{} }
func (m *Response) String() string            { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()               {}
func (*Response) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Response) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func init() {
	proto.RegisterType((*Request)(nil), "go.micro.srv.greeter.Request")
	proto.RegisterType((*Response)(nil), "go.micro.srv.greeter.Response")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Say service

type SayClient interface {
	Hello(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
}

type sayClient struct {
	cc *grpc.ClientConn
}

func NewSayClient(cc *grpc.ClientConn) SayClient {
	return &sayClient{cc}
}

func (c *sayClient) Hello(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := grpc.Invoke(ctx, "/go.micro.srv.greeter.Say/Hello", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Say service

type SayServer interface {
	Hello(context.Context, *Request) (*Response, error)
}

func RegisterSayServer(s *grpc.Server, srv SayServer) {
	s.RegisterService(&_Say_serviceDesc, srv)
}

func _Say_Hello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SayServer).Hello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/go.micro.srv.greeter.Say/Hello",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SayServer).Hello(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

var _Say_serviceDesc = grpc.ServiceDesc{
	ServiceName: "go.micro.srv.greeter.Say",
	HandlerType: (*SayServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Hello",
			Handler:    _Say_Hello_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "github.com/micro/go-grpc/examples/greeter/function/proto/hello/hello.proto",
}

func init() {
	proto.RegisterFile("github.com/micro/go-grpc/examples/greeter/function/proto/hello/hello.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 187 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x8e, 0x4d, 0x0a, 0xc2, 0x30,
	0x10, 0x85, 0x2d, 0xf5, 0x37, 0x2b, 0x09, 0x2e, 0x44, 0xac, 0x48, 0x57, 0x6e, 0x4c, 0x40, 0x2f,
	0x51, 0xdc, 0x08, 0xf5, 0x04, 0x6d, 0x18, 0xd3, 0x42, 0x93, 0x89, 0x49, 0x2a, 0x7a, 0x7b, 0x69,
	0xcc, 0x52, 0x37, 0xc3, 0x63, 0x3e, 0x66, 0xbe, 0x47, 0x2e, 0xb2, 0xf5, 0x4d, 0x5f, 0x33, 0x81,
	0x8a, 0xab, 0x56, 0x58, 0xe4, 0x12, 0x8f, 0xd2, 0x1a, 0xc1, 0xe1, 0x55, 0x29, 0xd3, 0x81, 0xe3,
	0xd2, 0x02, 0x78, 0xb0, 0xfc, 0xde, 0x6b, 0xe1, 0x5b, 0xd4, 0xdc, 0x58, 0xf4, 0xc8, 0x1b, 0xe8,
	0xba, 0x38, 0x59, 0xd8, 0xd0, 0x95, 0x44, 0x16, 0x7e, 0x30, 0x67, 0x9f, 0x2c, 0x9e, 0xe5, 0x19,
	0x99, 0x95, 0xf0, 0xe8, 0xc1, 0x79, 0x4a, 0xc9, 0x58, 0x57, 0x0a, 0xd6, 0xc9, 0x3e, 0x39, 0x2c,
	0xca, 0x90, 0xf3, 0x2d, 0x99, 0x97, 0xe0, 0x0c, 0x6a, 0x07, 0x74, 0x49, 0x52, 0xe5, 0x64, 0xc4,
	0x43, 0x3c, 0x5d, 0x49, 0x7a, 0xab, 0xde, 0xb4, 0x20, 0x93, 0x62, 0x10, 0xd1, 0x8c, 0xfd, 0x72,
	0xb0, 0x28, 0xd8, 0xec, 0xfe, 0xe1, 0xaf, 0x20, 0x1f, 0xd5, 0xd3, 0x50, 0xf5, 0xfc, 0x09, 0x00,
	0x00, 0xff, 0xff, 0x79, 0x62, 0x76, 0xe6, 0xf8, 0x00, 0x00, 0x00,
}
