// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/mail.proto

package educonn_mail

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type EmailRequest struct {
	From                 string   `protobuf:"bytes,1,opt,name=from,proto3" json:"from,omitempty"`
	To                   string   `protobuf:"bytes,2,opt,name=to,proto3" json:"to,omitempty"`
	Subject              string   `protobuf:"bytes,3,opt,name=subject,proto3" json:"subject,omitempty"`
	Message              string   `protobuf:"bytes,4,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EmailRequest) Reset()         { *m = EmailRequest{} }
func (m *EmailRequest) String() string { return proto.CompactTextString(m) }
func (*EmailRequest) ProtoMessage()    {}
func (*EmailRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_mail_b32833916ea8d983, []int{0}
}
func (m *EmailRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EmailRequest.Unmarshal(m, b)
}
func (m *EmailRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EmailRequest.Marshal(b, m, deterministic)
}
func (dst *EmailRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EmailRequest.Merge(dst, src)
}
func (m *EmailRequest) XXX_Size() int {
	return xxx_messageInfo_EmailRequest.Size(m)
}
func (m *EmailRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_EmailRequest.DiscardUnknown(m)
}

var xxx_messageInfo_EmailRequest proto.InternalMessageInfo

func (m *EmailRequest) GetFrom() string {
	if m != nil {
		return m.From
	}
	return ""
}

func (m *EmailRequest) GetTo() string {
	if m != nil {
		return m.To
	}
	return ""
}

func (m *EmailRequest) GetSubject() string {
	if m != nil {
		return m.Subject
	}
	return ""
}

func (m *EmailRequest) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type Response struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_mail_b32833916ea8d983, []int{1}
}
func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (dst *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(dst, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func init() {
	proto.RegisterType((*EmailRequest)(nil), "educonn.mail.EmailRequest")
	proto.RegisterType((*Response)(nil), "educonn.mail.Response")
}

func init() { proto.RegisterFile("proto/mail.proto", fileDescriptor_mail_b32833916ea8d983) }

var fileDescriptor_mail_b32833916ea8d983 = []byte{
	// 178 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x28, 0x28, 0xca, 0x2f,
	0xc9, 0xd7, 0xcf, 0x4d, 0xcc, 0xcc, 0xd1, 0x03, 0x33, 0x85, 0x78, 0x52, 0x53, 0x4a, 0x93, 0xf3,
	0xf3, 0xf2, 0xf4, 0x40, 0x62, 0x4a, 0x69, 0x5c, 0x3c, 0xae, 0x20, 0x46, 0x50, 0x6a, 0x61, 0x69,
	0x6a, 0x71, 0x89, 0x90, 0x10, 0x17, 0x4b, 0x5a, 0x51, 0x7e, 0xae, 0x04, 0xa3, 0x02, 0xa3, 0x06,
	0x67, 0x10, 0x98, 0x2d, 0xc4, 0xc7, 0xc5, 0x54, 0x92, 0x2f, 0xc1, 0x04, 0x16, 0x61, 0x2a, 0xc9,
	0x17, 0x92, 0xe0, 0x62, 0x2f, 0x2e, 0x4d, 0xca, 0x4a, 0x4d, 0x2e, 0x91, 0x60, 0x06, 0x0b, 0xc2,
	0xb8, 0x20, 0x99, 0xdc, 0xd4, 0xe2, 0xe2, 0xc4, 0xf4, 0x54, 0x09, 0x16, 0x88, 0x0c, 0x94, 0xab,
	0xc4, 0xc5, 0xc5, 0x11, 0x94, 0x5a, 0x5c, 0x90, 0x9f, 0x57, 0x9c, 0x6a, 0xe4, 0x0f, 0xb5, 0x33,
	0x38, 0xb5, 0xa8, 0x2c, 0x33, 0x39, 0x55, 0xc8, 0x9e, 0x8b, 0x33, 0x38, 0x35, 0x2f, 0x05, 0x2c,
	0x26, 0x24, 0xa5, 0x87, 0xec, 0x3e, 0x3d, 0x64, 0xc7, 0x49, 0x89, 0xa1, 0xca, 0xc1, 0x0c, 0x4c,
	0x62, 0x03, 0xfb, 0xcc, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0xa4, 0x50, 0x38, 0x75, 0xed, 0x00,
	0x00, 0x00,
}
