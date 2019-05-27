// Code generated by protoc-gen-go. DO NOT EDIT.
// source: auth.proto

package example

import (
	_ "auth"
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Request struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_8bbd6f3875b0e874, []int{0}
}

func (m *Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Request.Unmarshal(m, b)
}
func (m *Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Request.Marshal(b, m, deterministic)
}
func (m *Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request.Merge(m, src)
}
func (m *Request) XXX_Size() int {
	return xxx_messageInfo_Request.Size(m)
}
func (m *Request) XXX_DiscardUnknown() {
	xxx_messageInfo_Request.DiscardUnknown(m)
}

var xxx_messageInfo_Request proto.InternalMessageInfo

type Response struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_8bbd6f3875b0e874, []int{1}
}

func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (m *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(m, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Request)(nil), "example.Request")
	proto.RegisterType((*Response)(nil), "example.Response")
}

func init() { proto.RegisterFile("auth.proto", fileDescriptor_8bbd6f3875b0e874) }

var fileDescriptor_8bbd6f3875b0e874 = []byte{
	// 202 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4a, 0x2c, 0x2d, 0xc9,
	0xd0, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x4f, 0xad, 0x48, 0xcc, 0x2d, 0xc8, 0x49, 0x95,
	0x12, 0x03, 0x09, 0xea, 0x27, 0xe6, 0xe5, 0xe5, 0x97, 0x24, 0x96, 0x64, 0xe6, 0xe7, 0x15, 0x43,
	0x14, 0x48, 0xc9, 0xa4, 0xe7, 0xe7, 0xa7, 0xe7, 0xa4, 0xea, 0x27, 0x16, 0x64, 0x62, 0xca, 0x2a,
	0x71, 0x72, 0xb1, 0x07, 0xa5, 0x16, 0x96, 0xa6, 0x16, 0x97, 0x28, 0x71, 0x71, 0x71, 0x04, 0xa5,
	0x16, 0x17, 0xe4, 0xe7, 0x15, 0xa7, 0x1a, 0x39, 0x73, 0xb1, 0xbb, 0x42, 0xcc, 0x15, 0x32, 0xe6,
	0x62, 0x2d, 0x49, 0x2d, 0x2e, 0x31, 0x14, 0x12, 0xd0, 0x83, 0x5a, 0xa5, 0x07, 0xd5, 0x21, 0x25,
	0x88, 0x24, 0x02, 0xd1, 0xa8, 0xc4, 0x3a, 0xc9, 0x91, 0x89, 0x23, 0x45, 0x8a, 0x79, 0x92, 0x23,
	0x83, 0xd1, 0x24, 0x46, 0x2e, 0x0e, 0xa8, 0x29, 0x46, 0x42, 0x9e, 0x10, 0x63, 0x8c, 0x88, 0x33,
	0x46, 0x7a, 0x92, 0x23, 0x33, 0xc7, 0x09, 0xc6, 0xa6, 0xcb, 0x4f, 0x26, 0x33, 0xf1, 0x0b, 0xf1,
	0xea, 0x43, 0x15, 0xe8, 0x83, 0x8c, 0x10, 0x32, 0x81, 0x18, 0x65, 0x4c, 0x9c, 0x51, 0x6c, 0x20,
	0xa3, 0x26, 0x30, 0x4b, 0x81, 0x5d, 0xc6, 0x94, 0xc4, 0x06, 0xf6, 0xb7, 0x31, 0x20, 0x00, 0x00,
	0xff, 0xff, 0xd7, 0xc3, 0xb4, 0x2c, 0x44, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ExampleClient is the client API for Example service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ExampleClient interface {
	// comment
	Test1(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
}

type exampleClient struct {
	cc *grpc.ClientConn
}

func NewExampleClient(cc *grpc.ClientConn) ExampleClient {
	return &exampleClient{cc}
}

func (c *exampleClient) Test1(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/example.Example/test1", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ExampleServer is the server API for Example service.
type ExampleServer interface {
	// comment
	Test1(context.Context, *Request) (*Response, error)
}

// UnimplementedExampleServer can be embedded to have forward compatible implementations.
type UnimplementedExampleServer struct {
}

func (*UnimplementedExampleServer) Test1(ctx context.Context, req *Request) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Test1 not implemented")
}

func RegisterExampleServer(s *grpc.Server, srv ExampleServer) {
	s.RegisterService(&_Example_serviceDesc, srv)
}

func _Example_Test1_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExampleServer).Test1(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/example.Example/Test1",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExampleServer).Test1(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

var _Example_serviceDesc = grpc.ServiceDesc{
	ServiceName: "example.Example",
	HandlerType: (*ExampleServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "test1",
			Handler:    _Example_Test1_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "auth.proto",
}

// Example2Client is the client API for Example2 service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type Example2Client interface {
	// comment
	Test2(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
	Test3(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
}

type example2Client struct {
	cc *grpc.ClientConn
}

func NewExample2Client(cc *grpc.ClientConn) Example2Client {
	return &example2Client{cc}
}

func (c *example2Client) Test2(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/example.Example2/test2", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *example2Client) Test3(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/example.Example2/test3", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Example2Server is the server API for Example2 service.
type Example2Server interface {
	// comment
	Test2(context.Context, *Request) (*Response, error)
	Test3(context.Context, *Request) (*Response, error)
}

// UnimplementedExample2Server can be embedded to have forward compatible implementations.
type UnimplementedExample2Server struct {
}

func (*UnimplementedExample2Server) Test2(ctx context.Context, req *Request) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Test2 not implemented")
}
func (*UnimplementedExample2Server) Test3(ctx context.Context, req *Request) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Test3 not implemented")
}

func RegisterExample2Server(s *grpc.Server, srv Example2Server) {
	s.RegisterService(&_Example2_serviceDesc, srv)
}

func _Example2_Test2_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(Example2Server).Test2(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/example.Example2/Test2",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(Example2Server).Test2(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _Example2_Test3_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(Example2Server).Test3(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/example.Example2/Test3",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(Example2Server).Test3(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

var _Example2_serviceDesc = grpc.ServiceDesc{
	ServiceName: "example.Example2",
	HandlerType: (*Example2Server)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "test2",
			Handler:    _Example2_Test2_Handler,
		},
		{
			MethodName: "test3",
			Handler:    _Example2_Test3_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "auth.proto",
}