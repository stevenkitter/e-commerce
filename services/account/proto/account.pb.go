// Code generated by protoc-gen-go. DO NOT EDIT.
// source: services/account/proto/account.proto

package proto

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	proto1 "github.com/stevenkitter/skills/common/proto"
	grpc "google.golang.org/grpc"
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

type SignUpReq struct {
	Email                string   `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	Password             string   `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SignUpReq) Reset()         { *m = SignUpReq{} }
func (m *SignUpReq) String() string { return proto.CompactTextString(m) }
func (*SignUpReq) ProtoMessage()    {}
func (*SignUpReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_6c302e52287ea4fa, []int{0}
}

func (m *SignUpReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SignUpReq.Unmarshal(m, b)
}
func (m *SignUpReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SignUpReq.Marshal(b, m, deterministic)
}
func (m *SignUpReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SignUpReq.Merge(m, src)
}
func (m *SignUpReq) XXX_Size() int {
	return xxx_messageInfo_SignUpReq.Size(m)
}
func (m *SignUpReq) XXX_DiscardUnknown() {
	xxx_messageInfo_SignUpReq.DiscardUnknown(m)
}

var xxx_messageInfo_SignUpReq proto.InternalMessageInfo

func (m *SignUpReq) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *SignUpReq) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type LoginReq struct {
	Email                string   `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	Password             string   `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LoginReq) Reset()         { *m = LoginReq{} }
func (m *LoginReq) String() string { return proto.CompactTextString(m) }
func (*LoginReq) ProtoMessage()    {}
func (*LoginReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_6c302e52287ea4fa, []int{1}
}

func (m *LoginReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoginReq.Unmarshal(m, b)
}
func (m *LoginReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoginReq.Marshal(b, m, deterministic)
}
func (m *LoginReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoginReq.Merge(m, src)
}
func (m *LoginReq) XXX_Size() int {
	return xxx_messageInfo_LoginReq.Size(m)
}
func (m *LoginReq) XXX_DiscardUnknown() {
	xxx_messageInfo_LoginReq.DiscardUnknown(m)
}

var xxx_messageInfo_LoginReq proto.InternalMessageInfo

func (m *LoginReq) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *LoginReq) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type VerifyEmailReq struct {
	Email                string   `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VerifyEmailReq) Reset()         { *m = VerifyEmailReq{} }
func (m *VerifyEmailReq) String() string { return proto.CompactTextString(m) }
func (*VerifyEmailReq) ProtoMessage()    {}
func (*VerifyEmailReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_6c302e52287ea4fa, []int{2}
}

func (m *VerifyEmailReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VerifyEmailReq.Unmarshal(m, b)
}
func (m *VerifyEmailReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VerifyEmailReq.Marshal(b, m, deterministic)
}
func (m *VerifyEmailReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VerifyEmailReq.Merge(m, src)
}
func (m *VerifyEmailReq) XXX_Size() int {
	return xxx_messageInfo_VerifyEmailReq.Size(m)
}
func (m *VerifyEmailReq) XXX_DiscardUnknown() {
	xxx_messageInfo_VerifyEmailReq.DiscardUnknown(m)
}

var xxx_messageInfo_VerifyEmailReq proto.InternalMessageInfo

func (m *VerifyEmailReq) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

type ResetPasswordReq struct {
	Token                string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	Password             string   `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ResetPasswordReq) Reset()         { *m = ResetPasswordReq{} }
func (m *ResetPasswordReq) String() string { return proto.CompactTextString(m) }
func (*ResetPasswordReq) ProtoMessage()    {}
func (*ResetPasswordReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_6c302e52287ea4fa, []int{3}
}

func (m *ResetPasswordReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ResetPasswordReq.Unmarshal(m, b)
}
func (m *ResetPasswordReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ResetPasswordReq.Marshal(b, m, deterministic)
}
func (m *ResetPasswordReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResetPasswordReq.Merge(m, src)
}
func (m *ResetPasswordReq) XXX_Size() int {
	return xxx_messageInfo_ResetPasswordReq.Size(m)
}
func (m *ResetPasswordReq) XXX_DiscardUnknown() {
	xxx_messageInfo_ResetPasswordReq.DiscardUnknown(m)
}

var xxx_messageInfo_ResetPasswordReq proto.InternalMessageInfo

func (m *ResetPasswordReq) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *ResetPasswordReq) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func init() {
	proto.RegisterType((*SignUpReq)(nil), "proto.SignUpReq")
	proto.RegisterType((*LoginReq)(nil), "proto.LoginReq")
	proto.RegisterType((*VerifyEmailReq)(nil), "proto.VerifyEmailReq")
	proto.RegisterType((*ResetPasswordReq)(nil), "proto.ResetPasswordReq")
}

func init() {
	proto.RegisterFile("services/account/proto/account.proto", fileDescriptor_6c302e52287ea4fa)
}

var fileDescriptor_6c302e52287ea4fa = []byte{
	// 241 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x29, 0x4e, 0x2d, 0x2a,
	0xcb, 0x4c, 0x4e, 0x2d, 0xd6, 0x4f, 0x4c, 0x4e, 0xce, 0x2f, 0xcd, 0x2b, 0xd1, 0x2f, 0x28, 0xca,
	0x2f, 0xc9, 0x87, 0xf1, 0xf4, 0xc0, 0x3c, 0x21, 0x56, 0x30, 0x25, 0x25, 0x99, 0x9c, 0x9f, 0x9b,
	0x9b, 0x9f, 0x07, 0x55, 0x02, 0xe1, 0x40, 0x54, 0x28, 0xd9, 0x72, 0x71, 0x06, 0x67, 0xa6, 0xe7,
	0x85, 0x16, 0x04, 0xa5, 0x16, 0x0a, 0x89, 0x70, 0xb1, 0xa6, 0xe6, 0x26, 0x66, 0xe6, 0x48, 0x30,
	0x2a, 0x30, 0x6a, 0x70, 0x06, 0x41, 0x38, 0x42, 0x52, 0x5c, 0x1c, 0x05, 0x89, 0xc5, 0xc5, 0xe5,
	0xf9, 0x45, 0x29, 0x12, 0x4c, 0x60, 0x09, 0x38, 0x5f, 0xc9, 0x86, 0x8b, 0xc3, 0x27, 0x3f, 0x3d,
	0x33, 0x8f, 0x3c, 0xdd, 0x6a, 0x5c, 0x7c, 0x61, 0xa9, 0x45, 0x99, 0x69, 0x95, 0xae, 0x20, 0xa5,
	0x38, 0xcd, 0x50, 0x72, 0xe1, 0x12, 0x08, 0x4a, 0x2d, 0x4e, 0x2d, 0x09, 0x80, 0x6a, 0x84, 0xaa,
	0x2c, 0xc9, 0xcf, 0x4e, 0xcd, 0x83, 0xa9, 0x04, 0x73, 0xf0, 0xd9, 0x66, 0x64, 0xc4, 0xc5, 0xee,
	0x08, 0x09, 0x1d, 0x21, 0x75, 0x2e, 0x36, 0x88, 0xaf, 0x85, 0x04, 0x20, 0xe1, 0xa0, 0x07, 0x0f,
	0x04, 0x29, 0x6e, 0xa8, 0x48, 0x50, 0x6a, 0x71, 0x81, 0x93, 0x79, 0x94, 0x69, 0x7a, 0x66, 0x49,
	0x46, 0x69, 0x92, 0x5e, 0x72, 0x7e, 0xae, 0x7e, 0x71, 0x49, 0x6a, 0x59, 0x6a, 0x5e, 0x76, 0x66,
	0x49, 0x49, 0x6a, 0x91, 0x7e, 0x71, 0x76, 0x66, 0x4e, 0x4e, 0xb1, 0x3e, 0xf6, 0x78, 0x48, 0x62,
	0x03, 0x53, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0x2d, 0x38, 0x23, 0x37, 0xa8, 0x01, 0x00,
	0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// AccountClient is the client API for Account service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AccountClient interface {
	SignUp(ctx context.Context, in *SignUpReq, opts ...grpc.CallOption) (*proto1.Resp, error)
}

type accountClient struct {
	cc *grpc.ClientConn
}

func NewAccountClient(cc *grpc.ClientConn) AccountClient {
	return &accountClient{cc}
}

func (c *accountClient) SignUp(ctx context.Context, in *SignUpReq, opts ...grpc.CallOption) (*proto1.Resp, error) {
	out := new(proto1.Resp)
	err := c.cc.Invoke(ctx, "/proto.Account/SignUp", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AccountServer is the server API for Account service.
type AccountServer interface {
	SignUp(context.Context, *SignUpReq) (*proto1.Resp, error)
}

func RegisterAccountServer(s *grpc.Server, srv AccountServer) {
	s.RegisterService(&_Account_serviceDesc, srv)
}

func _Account_SignUp_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SignUpReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServer).SignUp(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Account/SignUp",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServer).SignUp(ctx, req.(*SignUpReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _Account_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Account",
	HandlerType: (*AccountServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SignUp",
			Handler:    _Account_SignUp_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "services/account/proto/account.proto",
}
