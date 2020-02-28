// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/logtest/logtest.proto

package ps_clientele_svc_logtest

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

type Message struct {
	Say                  string   `protobuf:"bytes,1,opt,name=say,proto3" json:"say,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Message) Reset()         { *m = Message{} }
func (m *Message) String() string { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()    {}
func (*Message) Descriptor() ([]byte, []int) {
	return fileDescriptor_21e94befb1c8af61, []int{0}
}

func (m *Message) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Message.Unmarshal(m, b)
}
func (m *Message) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Message.Marshal(b, m, deterministic)
}
func (m *Message) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Message.Merge(m, src)
}
func (m *Message) XXX_Size() int {
	return xxx_messageInfo_Message.Size(m)
}
func (m *Message) XXX_DiscardUnknown() {
	xxx_messageInfo_Message.DiscardUnknown(m)
}

var xxx_messageInfo_Message proto.InternalMessageInfo

func (m *Message) GetSay() string {
	if m != nil {
		return m.Say
	}
	return ""
}

type Request struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_21e94befb1c8af61, []int{1}
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

func (m *Request) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type Response struct {
	Msg                  string   `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_21e94befb1c8af61, []int{2}
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

type StreamingRequest struct {
	Count                int64    `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StreamingRequest) Reset()         { *m = StreamingRequest{} }
func (m *StreamingRequest) String() string { return proto.CompactTextString(m) }
func (*StreamingRequest) ProtoMessage()    {}
func (*StreamingRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_21e94befb1c8af61, []int{3}
}

func (m *StreamingRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StreamingRequest.Unmarshal(m, b)
}
func (m *StreamingRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StreamingRequest.Marshal(b, m, deterministic)
}
func (m *StreamingRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StreamingRequest.Merge(m, src)
}
func (m *StreamingRequest) XXX_Size() int {
	return xxx_messageInfo_StreamingRequest.Size(m)
}
func (m *StreamingRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_StreamingRequest.DiscardUnknown(m)
}

var xxx_messageInfo_StreamingRequest proto.InternalMessageInfo

func (m *StreamingRequest) GetCount() int64 {
	if m != nil {
		return m.Count
	}
	return 0
}

type StreamingResponse struct {
	Count                int64    `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StreamingResponse) Reset()         { *m = StreamingResponse{} }
func (m *StreamingResponse) String() string { return proto.CompactTextString(m) }
func (*StreamingResponse) ProtoMessage()    {}
func (*StreamingResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_21e94befb1c8af61, []int{4}
}

func (m *StreamingResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StreamingResponse.Unmarshal(m, b)
}
func (m *StreamingResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StreamingResponse.Marshal(b, m, deterministic)
}
func (m *StreamingResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StreamingResponse.Merge(m, src)
}
func (m *StreamingResponse) XXX_Size() int {
	return xxx_messageInfo_StreamingResponse.Size(m)
}
func (m *StreamingResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_StreamingResponse.DiscardUnknown(m)
}

var xxx_messageInfo_StreamingResponse proto.InternalMessageInfo

func (m *StreamingResponse) GetCount() int64 {
	if m != nil {
		return m.Count
	}
	return 0
}

type Ping struct {
	Stroke               int64    `protobuf:"varint,1,opt,name=stroke,proto3" json:"stroke,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Ping) Reset()         { *m = Ping{} }
func (m *Ping) String() string { return proto.CompactTextString(m) }
func (*Ping) ProtoMessage()    {}
func (*Ping) Descriptor() ([]byte, []int) {
	return fileDescriptor_21e94befb1c8af61, []int{5}
}

func (m *Ping) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Ping.Unmarshal(m, b)
}
func (m *Ping) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Ping.Marshal(b, m, deterministic)
}
func (m *Ping) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ping.Merge(m, src)
}
func (m *Ping) XXX_Size() int {
	return xxx_messageInfo_Ping.Size(m)
}
func (m *Ping) XXX_DiscardUnknown() {
	xxx_messageInfo_Ping.DiscardUnknown(m)
}

var xxx_messageInfo_Ping proto.InternalMessageInfo

func (m *Ping) GetStroke() int64 {
	if m != nil {
		return m.Stroke
	}
	return 0
}

type Pong struct {
	Stroke               int64    `protobuf:"varint,1,opt,name=stroke,proto3" json:"stroke,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Pong) Reset()         { *m = Pong{} }
func (m *Pong) String() string { return proto.CompactTextString(m) }
func (*Pong) ProtoMessage()    {}
func (*Pong) Descriptor() ([]byte, []int) {
	return fileDescriptor_21e94befb1c8af61, []int{6}
}

func (m *Pong) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Pong.Unmarshal(m, b)
}
func (m *Pong) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Pong.Marshal(b, m, deterministic)
}
func (m *Pong) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Pong.Merge(m, src)
}
func (m *Pong) XXX_Size() int {
	return xxx_messageInfo_Pong.Size(m)
}
func (m *Pong) XXX_DiscardUnknown() {
	xxx_messageInfo_Pong.DiscardUnknown(m)
}

var xxx_messageInfo_Pong proto.InternalMessageInfo

func (m *Pong) GetStroke() int64 {
	if m != nil {
		return m.Stroke
	}
	return 0
}

func init() {
	proto.RegisterType((*Message)(nil), "ps.clientele.svc.logtest.Message")
	proto.RegisterType((*Request)(nil), "ps.clientele.svc.logtest.Request")
	proto.RegisterType((*Response)(nil), "ps.clientele.svc.logtest.Response")
	proto.RegisterType((*StreamingRequest)(nil), "ps.clientele.svc.logtest.StreamingRequest")
	proto.RegisterType((*StreamingResponse)(nil), "ps.clientele.svc.logtest.StreamingResponse")
	proto.RegisterType((*Ping)(nil), "ps.clientele.svc.logtest.Ping")
	proto.RegisterType((*Pong)(nil), "ps.clientele.svc.logtest.Pong")
}

func init() { proto.RegisterFile("proto/logtest/logtest.proto", fileDescriptor_21e94befb1c8af61) }

var fileDescriptor_21e94befb1c8af61 = []byte{
	// 274 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0xe1, 0x4a, 0xc3, 0x30,
	0x10, 0xc7, 0x57, 0x37, 0xdb, 0x79, 0x9f, 0xe6, 0x21, 0x32, 0x3a, 0x1d, 0x9a, 0x4f, 0x55, 0x21,
	0x0e, 0x7d, 0x04, 0xbf, 0x2a, 0x8e, 0xfa, 0x04, 0xb5, 0x1c, 0xa1, 0x98, 0x26, 0xb5, 0x97, 0x09,
	0x3e, 0x87, 0x2f, 0x2c, 0x4d, 0x53, 0x10, 0xb1, 0xc3, 0x4f, 0xb9, 0xcb, 0xff, 0x77, 0x97, 0xbb,
	0x3f, 0x81, 0x55, 0xd3, 0x5a, 0x67, 0x6f, 0xb5, 0x55, 0x8e, 0xd8, 0x0d, 0xa7, 0xf4, 0xb7, 0xb8,
	0x6c, 0x58, 0x96, 0xba, 0x22, 0xe3, 0x48, 0x93, 0xe4, 0x8f, 0x52, 0x06, 0x5d, 0xac, 0x20, 0x79,
	0x22, 0xe6, 0x42, 0x11, 0x2e, 0x60, 0xca, 0xc5, 0xe7, 0x32, 0xba, 0x88, 0xb2, 0xa3, 0xbc, 0x0b,
	0xc5, 0x39, 0x24, 0x39, 0xbd, 0xef, 0x88, 0x1d, 0x22, 0xcc, 0x4c, 0x51, 0x53, 0x50, 0x7d, 0x2c,
	0xce, 0x60, 0x9e, 0x13, 0x37, 0xd6, 0xb0, 0x2f, 0xae, 0x59, 0x0d, 0xc5, 0x35, 0x2b, 0x91, 0xc1,
	0xe2, 0xc5, 0xb5, 0x54, 0xd4, 0x95, 0x51, 0x43, 0x97, 0x13, 0x38, 0x2c, 0xed, 0xce, 0x38, 0xcf,
	0x4d, 0xf3, 0x3e, 0x11, 0x57, 0x70, 0xfc, 0x83, 0x0c, 0x0d, 0xff, 0x46, 0xd7, 0x30, 0xdb, 0x56,
	0x46, 0xe1, 0x29, 0xc4, 0xec, 0x5a, 0xfb, 0x46, 0x41, 0x0e, 0x99, 0xd7, 0xed, 0xb8, 0x7e, 0xf7,
	0x75, 0x00, 0xc9, 0x63, 0xbf, 0x3a, 0x3e, 0xc3, 0xec, 0xa1, 0xd0, 0x1a, 0x2f, 0xe5, 0x98, 0x3b,
	0x32, 0xcc, 0x9d, 0x8a, 0x7d, 0x48, 0x3f, 0xb0, 0x98, 0x20, 0x41, 0xdc, 0xef, 0x81, 0xd7, 0xe3,
	0xfc, 0x6f, 0x4f, 0xd2, 0x9b, 0x7f, 0xb1, 0xc3, 0x23, 0x9b, 0x08, 0xb7, 0x30, 0xef, 0x3c, 0xf0,
	0x7b, 0xae, 0xc7, 0x8b, 0x3b, 0x26, 0xdd, 0xa7, 0x5b, 0xa3, 0xc4, 0x24, 0x8b, 0x36, 0xd1, 0x6b,
	0xec, 0x7f, 0xc9, 0xfd, 0x77, 0x00, 0x00, 0x00, 0xff, 0xff, 0x4e, 0x47, 0x3a, 0x9a, 0x44, 0x02,
	0x00, 0x00,
}
