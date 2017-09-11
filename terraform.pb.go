// Code generated by protoc-gen-go. DO NOT EDIT.
// source: tfgrpc/terraform.proto

/*
Package tfgrpc is a generated protocol buffer package.

It is generated from these files:
	tfgrpc/terraform.proto

It has these top-level messages:
	Body
	Id
	Arg
	Output
*/
package tfgrpc

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

type Body struct {
	Zipfile []byte `protobuf:"bytes,1,opt,name=zipfile,proto3" json:"zipfile,omitempty"`
}

func (m *Body) Reset()                    { *m = Body{} }
func (m *Body) String() string            { return proto.CompactTextString(m) }
func (*Body) ProtoMessage()               {}
func (*Body) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Body) GetZipfile() []byte {
	if m != nil {
		return m.Zipfile
	}
	return nil
}

type Id struct {
	Tmpdir string `protobuf:"bytes,1,opt,name=tmpdir" json:"tmpdir,omitempty"`
}

func (m *Id) Reset()                    { *m = Id{} }
func (m *Id) String() string            { return proto.CompactTextString(m) }
func (*Id) ProtoMessage()               {}
func (*Id) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Id) GetTmpdir() string {
	if m != nil {
		return m.Tmpdir
	}
	return ""
}

// The request message containing the user's name.
type Arg struct {
	WorkingDir string   `protobuf:"bytes,1,opt,name=workingDir" json:"workingDir,omitempty"`
	Args       []string `protobuf:"bytes,2,rep,name=args" json:"args,omitempty"`
}

func (m *Arg) Reset()                    { *m = Arg{} }
func (m *Arg) String() string            { return proto.CompactTextString(m) }
func (*Arg) ProtoMessage()               {}
func (*Arg) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Arg) GetWorkingDir() string {
	if m != nil {
		return m.WorkingDir
	}
	return ""
}

func (m *Arg) GetArgs() []string {
	if m != nil {
		return m.Args
	}
	return nil
}

// The response message containing the greetings
type Output struct {
	Retcode int32  `protobuf:"varint,1,opt,name=retcode" json:"retcode,omitempty"`
	Stdout  []byte `protobuf:"bytes,2,opt,name=stdout,proto3" json:"stdout,omitempty"`
	Stderr  []byte `protobuf:"bytes,3,opt,name=stderr,proto3" json:"stderr,omitempty"`
}

func (m *Output) Reset()                    { *m = Output{} }
func (m *Output) String() string            { return proto.CompactTextString(m) }
func (*Output) ProtoMessage()               {}
func (*Output) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *Output) GetRetcode() int32 {
	if m != nil {
		return m.Retcode
	}
	return 0
}

func (m *Output) GetStdout() []byte {
	if m != nil {
		return m.Stdout
	}
	return nil
}

func (m *Output) GetStderr() []byte {
	if m != nil {
		return m.Stderr
	}
	return nil
}

func init() {
	proto.RegisterType((*Body)(nil), "tfgrpc.Body")
	proto.RegisterType((*Id)(nil), "tfgrpc.Id")
	proto.RegisterType((*Arg)(nil), "tfgrpc.Arg")
	proto.RegisterType((*Output)(nil), "tfgrpc.Output")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Terraform service

type TerraformClient interface {
	Init(ctx context.Context, in *Arg, opts ...grpc.CallOption) (*Output, error)
	Plan(ctx context.Context, in *Arg, opts ...grpc.CallOption) (*Output, error)
	Apply(ctx context.Context, in *Arg, opts ...grpc.CallOption) (*Output, error)
	Push(ctx context.Context, opts ...grpc.CallOption) (Terraform_PushClient, error)
}

type terraformClient struct {
	cc *grpc.ClientConn
}

func NewTerraformClient(cc *grpc.ClientConn) TerraformClient {
	return &terraformClient{cc}
}

func (c *terraformClient) Init(ctx context.Context, in *Arg, opts ...grpc.CallOption) (*Output, error) {
	out := new(Output)
	err := grpc.Invoke(ctx, "/tfgrpc.Terraform/Init", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *terraformClient) Plan(ctx context.Context, in *Arg, opts ...grpc.CallOption) (*Output, error) {
	out := new(Output)
	err := grpc.Invoke(ctx, "/tfgrpc.Terraform/Plan", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *terraformClient) Apply(ctx context.Context, in *Arg, opts ...grpc.CallOption) (*Output, error) {
	out := new(Output)
	err := grpc.Invoke(ctx, "/tfgrpc.Terraform/Apply", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *terraformClient) Push(ctx context.Context, opts ...grpc.CallOption) (Terraform_PushClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_Terraform_serviceDesc.Streams[0], c.cc, "/tfgrpc.Terraform/Push", opts...)
	if err != nil {
		return nil, err
	}
	x := &terraformPushClient{stream}
	return x, nil
}

type Terraform_PushClient interface {
	Send(*Body) error
	CloseAndRecv() (*Id, error)
	grpc.ClientStream
}

type terraformPushClient struct {
	grpc.ClientStream
}

func (x *terraformPushClient) Send(m *Body) error {
	return x.ClientStream.SendMsg(m)
}

func (x *terraformPushClient) CloseAndRecv() (*Id, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(Id)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for Terraform service

type TerraformServer interface {
	Init(context.Context, *Arg) (*Output, error)
	Plan(context.Context, *Arg) (*Output, error)
	Apply(context.Context, *Arg) (*Output, error)
	Push(Terraform_PushServer) error
}

func RegisterTerraformServer(s *grpc.Server, srv TerraformServer) {
	s.RegisterService(&_Terraform_serviceDesc, srv)
}

func _Terraform_Init_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Arg)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TerraformServer).Init(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tfgrpc.Terraform/Init",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TerraformServer).Init(ctx, req.(*Arg))
	}
	return interceptor(ctx, in, info, handler)
}

func _Terraform_Plan_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Arg)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TerraformServer).Plan(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tfgrpc.Terraform/Plan",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TerraformServer).Plan(ctx, req.(*Arg))
	}
	return interceptor(ctx, in, info, handler)
}

func _Terraform_Apply_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Arg)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TerraformServer).Apply(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tfgrpc.Terraform/Apply",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TerraformServer).Apply(ctx, req.(*Arg))
	}
	return interceptor(ctx, in, info, handler)
}

func _Terraform_Push_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(TerraformServer).Push(&terraformPushServer{stream})
}

type Terraform_PushServer interface {
	SendAndClose(*Id) error
	Recv() (*Body, error)
	grpc.ServerStream
}

type terraformPushServer struct {
	grpc.ServerStream
}

func (x *terraformPushServer) SendAndClose(m *Id) error {
	return x.ServerStream.SendMsg(m)
}

func (x *terraformPushServer) Recv() (*Body, error) {
	m := new(Body)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _Terraform_serviceDesc = grpc.ServiceDesc{
	ServiceName: "tfgrpc.Terraform",
	HandlerType: (*TerraformServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Init",
			Handler:    _Terraform_Init_Handler,
		},
		{
			MethodName: "Plan",
			Handler:    _Terraform_Plan_Handler,
		},
		{
			MethodName: "Apply",
			Handler:    _Terraform_Apply_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Push",
			Handler:       _Terraform_Push_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "tfgrpc/terraform.proto",
}

func init() { proto.RegisterFile("tfgrpc/terraform.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 262 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x91, 0xcf, 0x4a, 0xc4, 0x30,
	0x10, 0xc6, 0xb7, 0xdb, 0x6e, 0xa5, 0xe3, 0xe2, 0x21, 0x87, 0xa5, 0x2c, 0x22, 0x25, 0xa8, 0xf4,
	0x54, 0x41, 0x4f, 0x1e, 0x2b, 0x5e, 0x7a, 0x72, 0x29, 0xbe, 0x40, 0xdd, 0xa4, 0x31, 0xd8, 0x6d,
	0xc2, 0x74, 0x8a, 0xac, 0x2f, 0xe4, 0x6b, 0x4a, 0xff, 0x64, 0xf1, 0xb6, 0xb7, 0xfc, 0xbe, 0xf9,
	0xe0, 0x97, 0x49, 0x60, 0x43, 0xb5, 0x42, 0xbb, 0x7f, 0x20, 0x89, 0x58, 0xd5, 0x06, 0x0f, 0x99,
	0x45, 0x43, 0x86, 0x85, 0x53, 0xce, 0x13, 0x08, 0x5e, 0x8c, 0x38, 0xb2, 0x18, 0x2e, 0x7e, 0xb4,
	0xad, 0x75, 0x23, 0x63, 0x2f, 0xf1, 0xd2, 0x75, 0xe9, 0x90, 0x5f, 0xc3, 0xb2, 0x10, 0x6c, 0x03,
	0x21, 0x1d, 0xac, 0xd0, 0x38, 0x8e, 0xa3, 0x72, 0x26, 0xfe, 0x0c, 0x7e, 0x8e, 0x8a, 0xdd, 0x00,
	0x7c, 0x1b, 0xfc, 0xd2, 0xad, 0x7a, 0x3d, 0x55, 0xfe, 0x25, 0x8c, 0x41, 0x50, 0xa1, 0xea, 0xe2,
	0x65, 0xe2, 0xa7, 0x51, 0x39, 0x9e, 0x79, 0x09, 0xe1, 0x5b, 0x4f, 0xb6, 0xa7, 0x41, 0x8e, 0x92,
	0xf6, 0x46, 0x4c, 0xf2, 0x55, 0xe9, 0x70, 0xd0, 0x76, 0x24, 0x4c, 0x4f, 0xf1, 0x72, 0xbc, 0xd5,
	0x4c, 0x73, 0x2e, 0x11, 0x63, 0xff, 0x94, 0x4b, 0xc4, 0xc7, 0x5f, 0x0f, 0xa2, 0x77, 0xb7, 0x2a,
	0xbb, 0x83, 0xa0, 0x68, 0x35, 0xb1, 0xcb, 0x6c, 0xda, 0x36, 0xcb, 0x51, 0x6d, 0xaf, 0x1c, 0x4c,
	0x72, 0xbe, 0x18, 0x6a, 0xbb, 0xa6, 0x6a, 0xcf, 0xd5, 0xee, 0x61, 0x95, 0x5b, 0xdb, 0x1c, 0xcf,
	0xf5, 0x6e, 0x21, 0xd8, 0xf5, 0xdd, 0x27, 0x5b, 0xbb, 0xc9, 0xf0, 0xc0, 0x5b, 0x70, 0x54, 0x08,
	0xbe, 0x48, 0xbd, 0x8f, 0x70, 0xfc, 0x87, 0xa7, 0xbf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x29, 0x80,
	0x6d, 0x36, 0xa1, 0x01, 0x00, 0x00,
}
