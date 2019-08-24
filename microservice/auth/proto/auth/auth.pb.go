// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/auth/auth.proto

package mu_micro_book_srv_auth

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type Error struct {
	Code                 int32    `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Detail               string   `protobuf:"bytes,2,opt,name=detail,proto3" json:"detail,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Error) Reset()         { *m = Error{} }
func (m *Error) String() string { return proto.CompactTextString(m) }
func (*Error) ProtoMessage()    {}
func (*Error) Descriptor() ([]byte, []int) {
	return fileDescriptor_82b5829f48cfb8e5, []int{0}
}

func (m *Error) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Error.Unmarshal(m, b)
}
func (m *Error) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Error.Marshal(b, m, deterministic)
}
func (m *Error) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Error.Merge(m, src)
}
func (m *Error) XXX_Size() int {
	return xxx_messageInfo_Error.Size(m)
}
func (m *Error) XXX_DiscardUnknown() {
	xxx_messageInfo_Error.DiscardUnknown(m)
}

var xxx_messageInfo_Error proto.InternalMessageInfo

func (m *Error) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *Error) GetDetail() string {
	if m != nil {
		return m.Detail
	}
	return ""
}

type Request struct {
	UserId               uint64   `protobuf:"varint,1,opt,name=userId,proto3" json:"userId,omitempty"`
	UserName             string   `protobuf:"bytes,2,opt,name=userName,proto3" json:"userName,omitempty"`
	Token                string   `protobuf:"bytes,3,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_82b5829f48cfb8e5, []int{1}
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

func (m *Request) GetUserId() uint64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *Request) GetUserName() string {
	if m != nil {
		return m.UserName
	}
	return ""
}

func (m *Request) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type Response struct {
	Success              bool     `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Error                *Error   `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	Token                string   `protobuf:"bytes,3,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_82b5829f48cfb8e5, []int{2}
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

func (m *Response) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func (m *Response) GetError() *Error {
	if m != nil {
		return m.Error
	}
	return nil
}

func (m *Response) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func init() {
	proto.RegisterType((*Error)(nil), "mu.micro.book.srv.auth.Error")
	proto.RegisterType((*Request)(nil), "mu.micro.book.srv.auth.Request")
	proto.RegisterType((*Response)(nil), "mu.micro.book.srv.auth.Response")
}

func init() { proto.RegisterFile("proto/auth/auth.proto", fileDescriptor_82b5829f48cfb8e5) }

var fileDescriptor_82b5829f48cfb8e5 = []byte{
	// 272 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x92, 0x41, 0x4b, 0xc3, 0x40,
	0x10, 0x85, 0x49, 0x6d, 0x9a, 0x38, 0x1e, 0x84, 0xa1, 0x96, 0x50, 0x10, 0x43, 0x4e, 0x3d, 0xad,
	0xd0, 0xfc, 0x02, 0x51, 0x11, 0x0f, 0x7a, 0xd8, 0x56, 0xc4, 0x63, 0xba, 0x19, 0x68, 0x48, 0xd3,
	0xa9, 0xbb, 0x9b, 0xfe, 0x11, 0xff, 0xb0, 0x64, 0x1a, 0x3d, 0xb5, 0xb7, 0x5e, 0xc2, 0x7c, 0xbc,
	0x37, 0x6f, 0xf2, 0x60, 0xe1, 0x66, 0x67, 0xd9, 0xf3, 0x7d, 0xd1, 0xfa, 0xb5, 0x7c, 0x94, 0x30,
	0x4e, 0x9a, 0x56, 0x35, 0x95, 0xb1, 0xac, 0x56, 0xcc, 0xb5, 0x72, 0x76, 0xaf, 0x3a, 0x35, 0xcb,
	0x21, 0x7c, 0xb6, 0x96, 0x2d, 0x22, 0x0c, 0x0d, 0x97, 0x94, 0x04, 0x69, 0x30, 0x0b, 0xb5, 0xcc,
	0x38, 0x81, 0x51, 0x49, 0xbe, 0xa8, 0x36, 0xc9, 0x20, 0x0d, 0x66, 0x97, 0xba, 0xa7, 0x6c, 0x01,
	0x91, 0xa6, 0xef, 0x96, 0x9c, 0xef, 0x2c, 0xad, 0x23, 0xfb, 0x5a, 0xca, 0xe2, 0x50, 0xf7, 0x84,
	0x53, 0x88, 0xbb, 0xe9, 0xbd, 0x68, 0xa8, 0x5f, 0xfe, 0x67, 0x1c, 0x43, 0xe8, 0xb9, 0xa6, 0x6d,
	0x72, 0x21, 0xc2, 0x01, 0x32, 0x86, 0x58, 0x93, 0xdb, 0xf1, 0xd6, 0x11, 0x26, 0x10, 0xb9, 0xd6,
	0x18, 0x72, 0x4e, 0x62, 0x63, 0xfd, 0x87, 0x98, 0x43, 0x48, 0xdd, 0xff, 0x4a, 0xe8, 0xd5, 0xfc,
	0x56, 0x1d, 0xef, 0xa5, 0xa4, 0x94, 0x3e, 0x78, 0x8f, 0x1f, 0x9c, 0xff, 0x0c, 0x20, 0x5a, 0x90,
	0xdd, 0x57, 0x86, 0x70, 0x09, 0xd7, 0x6f, 0x45, 0x4d, 0x0f, 0x72, 0x64, 0xd9, 0xc9, 0x78, 0x77,
	0x2a, 0xba, 0xaf, 0x3e, 0x4d, 0x4f, 0x1b, 0xfa, 0x1a, 0x9f, 0x80, 0x4f, 0xb4, 0xf9, 0x70, 0x64,
	0xcf, 0x1c, 0xfc, 0x05, 0xe3, 0x17, 0xf2, 0x8f, 0x85, 0x59, 0x53, 0x79, 0xde, 0xe8, 0xd5, 0x48,
	0xde, 0x4b, 0xfe, 0x1b, 0x00, 0x00, 0xff, 0xff, 0xfe, 0x67, 0xc5, 0x05, 0x48, 0x02, 0x00, 0x00,
}
