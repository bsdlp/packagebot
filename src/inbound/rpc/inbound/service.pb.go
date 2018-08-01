// Code generated by protoc-gen-go. DO NOT EDIT.
// source: service.proto

package inbound

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/golang/protobuf/ptypes/empty"
import timestamp "github.com/golang/protobuf/ptypes/timestamp"

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
	DiscordUser          string               `protobuf:"bytes,1,opt,name=discord_user,json=discordUser,proto3" json:"discord_user,omitempty"`
	DiscordMessageId     string               `protobuf:"bytes,2,opt,name=discord_message_id,json=discordMessageId,proto3" json:"discord_message_id,omitempty"`
	DiscordChannelId     string               `protobuf:"bytes,3,opt,name=discord_channel_id,json=discordChannelId,proto3" json:"discord_channel_id,omitempty"`
	RawMessage           string               `protobuf:"bytes,4,opt,name=raw_message,json=rawMessage,proto3" json:"raw_message,omitempty"`
	MessageTimestamp     *timestamp.Timestamp `protobuf:"bytes,5,opt,name=message_timestamp,json=messageTimestamp,proto3" json:"message_timestamp,omitempty"`
	ReceivedTimestamp    *timestamp.Timestamp `protobuf:"bytes,6,opt,name=received_timestamp,json=receivedTimestamp,proto3" json:"received_timestamp,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_service_70a6757e934895b0, []int{0}
}
func (m *Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Request.Unmarshal(m, b)
}
func (m *Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Request.Marshal(b, m, deterministic)
}
func (dst *Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request.Merge(dst, src)
}
func (m *Request) XXX_Size() int {
	return xxx_messageInfo_Request.Size(m)
}
func (m *Request) XXX_DiscardUnknown() {
	xxx_messageInfo_Request.DiscardUnknown(m)
}

var xxx_messageInfo_Request proto.InternalMessageInfo

func (m *Request) GetDiscordUser() string {
	if m != nil {
		return m.DiscordUser
	}
	return ""
}

func (m *Request) GetDiscordMessageId() string {
	if m != nil {
		return m.DiscordMessageId
	}
	return ""
}

func (m *Request) GetDiscordChannelId() string {
	if m != nil {
		return m.DiscordChannelId
	}
	return ""
}

func (m *Request) GetRawMessage() string {
	if m != nil {
		return m.RawMessage
	}
	return ""
}

func (m *Request) GetMessageTimestamp() *timestamp.Timestamp {
	if m != nil {
		return m.MessageTimestamp
	}
	return nil
}

func (m *Request) GetReceivedTimestamp() *timestamp.Timestamp {
	if m != nil {
		return m.ReceivedTimestamp
	}
	return nil
}

func init() {
	proto.RegisterType((*Request)(nil), "bsdlp.packagebot.inbound.Request")
}

func init() { proto.RegisterFile("service.proto", fileDescriptor_service_70a6757e934895b0) }

var fileDescriptor_service_70a6757e934895b0 = []byte{
	// 298 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x91, 0xd1, 0x4a, 0xf3, 0x30,
	0x1c, 0xc5, 0xe9, 0xbe, 0xcf, 0x95, 0xa5, 0x2a, 0x5b, 0x2e, 0xa4, 0xd4, 0x8b, 0x4d, 0xaf, 0x76,
	0x21, 0x19, 0xcc, 0x37, 0x50, 0x44, 0x8a, 0x08, 0x52, 0xd4, 0x0b, 0x6f, 0x46, 0x9a, 0xfc, 0xad,
	0xc1, 0xb6, 0xa9, 0x49, 0xba, 0xe1, 0xc3, 0xf9, 0x6e, 0xb2, 0x26, 0x29, 0x53, 0x11, 0x6f, 0xff,
	0xe7, 0x77, 0x4e, 0x38, 0x39, 0xe8, 0x40, 0x83, 0x5a, 0x0b, 0x06, 0xa4, 0x51, 0xd2, 0x48, 0x1c,
	0xe7, 0x9a, 0x97, 0x0d, 0x69, 0x28, 0x7b, 0xa5, 0x05, 0xe4, 0xd2, 0x10, 0x51, 0xe7, 0xb2, 0xad,
	0x79, 0x72, 0x5c, 0x48, 0x59, 0x94, 0xb0, 0xe8, 0xb8, 0xbc, 0x7d, 0x5e, 0x40, 0xd5, 0x98, 0x77,
	0x6b, 0x4b, 0xa6, 0xdf, 0x45, 0x23, 0x2a, 0xd0, 0x86, 0x56, 0x8d, 0x05, 0x4e, 0x3f, 0x06, 0x28,
	0xcc, 0xe0, 0xad, 0x05, 0x6d, 0xf0, 0x09, 0xda, 0xe7, 0x42, 0x33, 0xa9, 0xf8, 0xaa, 0xd5, 0xa0,
	0xe2, 0x60, 0x16, 0xcc, 0x47, 0x59, 0xe4, 0x6e, 0x0f, 0x1a, 0x14, 0x3e, 0x43, 0xd8, 0x23, 0x15,
	0x68, 0x4d, 0x0b, 0x58, 0x09, 0x1e, 0x0f, 0x3a, 0x70, 0xec, 0x94, 0x5b, 0x2b, 0xa4, 0x7c, 0x97,
	0x66, 0x2f, 0xb4, 0xae, 0xa1, 0xdc, 0xd2, 0xff, 0xbe, 0xd0, 0x97, 0x56, 0x48, 0x39, 0x9e, 0xa2,
	0x48, 0xd1, 0x8d, 0xcf, 0x8d, 0xff, 0x77, 0x18, 0x52, 0x74, 0xe3, 0x02, 0xf1, 0x35, 0x9a, 0xf8,
	0x47, 0xfb, 0x1a, 0xf1, 0xde, 0x2c, 0x98, 0x47, 0xcb, 0x84, 0xd8, 0xa2, 0xc4, 0x17, 0x25, 0xf7,
	0x9e, 0xc8, 0xc6, 0xce, 0xd4, 0x5f, 0x70, 0x8a, 0xb0, 0x02, 0x06, 0x62, 0x0d, 0x7c, 0x27, 0x69,
	0xf8, 0x67, 0xd2, 0xc4, 0xbb, 0xfa, 0xd3, 0xf2, 0x11, 0x85, 0xa9, 0x1d, 0x02, 0xdf, 0xa0, 0xc3,
	0x3b, 0x25, 0x19, 0x68, 0xdd, 0x7f, 0x28, 0xf9, 0x6d, 0x35, 0xe2, 0x90, 0xe4, 0xe8, 0xc7, 0x73,
	0x57, 0xdb, 0xf9, 0x2e, 0x46, 0x4f, 0xa1, 0x43, 0xf3, 0x61, 0x27, 0x9d, 0x7f, 0x06, 0x00, 0x00,
	0xff, 0xff, 0x95, 0x87, 0x8c, 0x4d, 0x12, 0x02, 0x00, 0x00,
}
