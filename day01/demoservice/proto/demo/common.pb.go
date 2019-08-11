// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/demo/common.proto

package go_micro_srv_demo

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

// enum
type RespType int32

const (
	RespType_NONE    RespType = 0
	RespType_ASCEND  RespType = 1
	RespType_DESCEND RespType = 2
)

var RespType_name = map[int32]string{
	0: "NONE",
	1: "ASCEND",
	2: "DESCEND",
}

var RespType_value = map[string]int32{
	"NONE":    0,
	"ASCEND":  1,
	"DESCEND": 2,
}

func (x RespType) String() string {
	return proto.EnumName(RespType_name, int32(x))
}

func (RespType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_444744b1a245f759, []int{0}
}

type Request struct {
	Msg                  string   `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_444744b1a245f759, []int{0}
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

func (m *Request) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

type Pair struct {
	Key                  int32    `protobuf:"varint,1,opt,name=key,proto3" json:"key,omitempty"`
	Values               string   `protobuf:"bytes,2,opt,name=values,proto3" json:"values,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Pair) Reset()         { *m = Pair{} }
func (m *Pair) String() string { return proto.CompactTextString(m) }
func (*Pair) ProtoMessage()    {}
func (*Pair) Descriptor() ([]byte, []int) {
	return fileDescriptor_444744b1a245f759, []int{1}
}

func (m *Pair) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Pair.Unmarshal(m, b)
}
func (m *Pair) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Pair.Marshal(b, m, deterministic)
}
func (m *Pair) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Pair.Merge(m, src)
}
func (m *Pair) XXX_Size() int {
	return xxx_messageInfo_Pair.Size(m)
}
func (m *Pair) XXX_DiscardUnknown() {
	xxx_messageInfo_Pair.DiscardUnknown(m)
}

var xxx_messageInfo_Pair proto.InternalMessageInfo

func (m *Pair) GetKey() int32 {
	if m != nil {
		return m.Key
	}
	return 0
}

func (m *Pair) GetValues() string {
	if m != nil {
		return m.Values
	}
	return ""
}

type Response struct {
	Msg string `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
	// 数组
	Values []string `protobuf:"bytes,2,rep,name=values,proto3" json:"values,omitempty"`
	// map
	Header               map[string]*Pair `protobuf:"bytes,3,rep,name=header,proto3" json:"header,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Type                 RespType         `protobuf:"varint,4,opt,name=type,proto3,enum=go.micro.srv.demo.RespType" json:"type,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_444744b1a245f759, []int{2}
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

func (m *Response) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func (m *Response) GetValues() []string {
	if m != nil {
		return m.Values
	}
	return nil
}

func (m *Response) GetHeader() map[string]*Pair {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *Response) GetType() RespType {
	if m != nil {
		return m.Type
	}
	return RespType_NONE
}

func init() {
	proto.RegisterEnum("go.micro.srv.demo.RespType", RespType_name, RespType_value)
	proto.RegisterType((*Request)(nil), "go.micro.srv.demo.Request")
	proto.RegisterType((*Pair)(nil), "go.micro.srv.demo.Pair")
	proto.RegisterType((*Response)(nil), "go.micro.srv.demo.Response")
	proto.RegisterMapType((map[string]*Pair)(nil), "go.micro.srv.demo.Response.HeaderEntry")
}

func init() { proto.RegisterFile("proto/demo/common.proto", fileDescriptor_444744b1a245f759) }

var fileDescriptor_444744b1a245f759 = []byte{
	// 302 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x51, 0x41, 0x4f, 0xf2, 0x40,
	0x10, 0xfd, 0x96, 0x96, 0x02, 0xd3, 0xe4, 0x0b, 0xee, 0x41, 0x1a, 0xb8, 0x90, 0x5e, 0x6c, 0x4c,
	0xd8, 0x9a, 0x7a, 0x31, 0x5e, 0xd4, 0x48, 0x23, 0xa7, 0x6a, 0x16, 0xff, 0x40, 0x85, 0x09, 0x12,
	0xdb, 0x6e, 0xdd, 0x2d, 0x24, 0xfd, 0xeb, 0x9e, 0xcc, 0x6e, 0x31, 0x62, 0x40, 0x6f, 0x33, 0x3b,
	0xef, 0xbd, 0x7d, 0x6f, 0x06, 0x06, 0xa5, 0x14, 0x95, 0x08, 0x97, 0x98, 0x8b, 0x70, 0x21, 0xf2,
	0x5c, 0x14, 0xcc, 0xbc, 0xd0, 0x93, 0x95, 0x60, 0xf9, 0x7a, 0x21, 0x05, 0x53, 0x72, 0xcb, 0xf4,
	0xdc, 0x1f, 0x41, 0x87, 0xe3, 0xfb, 0x06, 0x55, 0x45, 0xfb, 0x60, 0xe5, 0x6a, 0xe5, 0x91, 0x31,
	0x09, 0x7a, 0x5c, 0x97, 0xfe, 0x05, 0xd8, 0x4f, 0xe9, 0x5a, 0xea, 0xc9, 0x1b, 0xd6, 0x66, 0xd2,
	0xe6, 0xba, 0xa4, 0xa7, 0xe0, 0x6c, 0xd3, 0x6c, 0x83, 0xca, 0x6b, 0x19, 0xf8, 0xae, 0xf3, 0x3f,
	0x08, 0x74, 0x39, 0xaa, 0x52, 0x14, 0x0a, 0x0f, 0x05, 0x7f, 0xd0, 0xac, 0x6f, 0x1a, 0xbd, 0x01,
	0xe7, 0x15, 0xd3, 0x25, 0x4a, 0xcf, 0x1a, 0x5b, 0x81, 0x1b, 0x9d, 0xb1, 0x03, 0xa7, 0xec, 0x4b,
	0x96, 0xcd, 0x0c, 0x32, 0x2e, 0x2a, 0x59, 0xf3, 0x1d, 0x8d, 0x86, 0x60, 0x57, 0x75, 0x89, 0x9e,
	0x3d, 0x26, 0xc1, 0xff, 0x68, 0xf4, 0x0b, 0xfd, 0xb9, 0x2e, 0x91, 0x1b, 0xe0, 0x90, 0x83, 0xbb,
	0xa7, 0xb3, 0x9f, 0xb0, 0xd7, 0x24, 0x9c, 0x40, 0xdb, 0x98, 0x33, 0x01, 0xdd, 0x68, 0x70, 0x44,
	0x52, 0xef, 0x86, 0x37, 0xa8, 0xeb, 0xd6, 0x15, 0x39, 0x9f, 0x34, 0xd9, 0xf5, 0x2f, 0xb4, 0x0b,
	0x76, 0xf2, 0x98, 0xc4, 0xfd, 0x7f, 0x14, 0xc0, 0xb9, 0x9b, 0xdf, 0xc7, 0xc9, 0xb4, 0x4f, 0xa8,
	0x0b, 0x9d, 0x69, 0xdc, 0x34, 0xad, 0xe8, 0x01, 0xac, 0x79, 0x5a, 0xd3, 0x5b, 0x68, 0xcf, 0x30,
	0xcb, 0x04, 0x1d, 0x1e, 0x75, 0x6d, 0x6e, 0x33, 0x1c, 0xfd, 0xb1, 0x90, 0x17, 0xc7, 0x5c, 0xf7,
	0xf2, 0x33, 0x00, 0x00, 0xff, 0xff, 0x5f, 0xc0, 0x2b, 0x2c, 0xf8, 0x01, 0x00, 0x00,
}