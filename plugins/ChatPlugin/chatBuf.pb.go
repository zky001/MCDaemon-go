// Code generated by protoc-gen-go. DO NOT EDIT.
// source: chatBuf.proto

package ChatPlugin

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
	ServerName           *string  `protobuf:"bytes,1,req,name=ServerName" json:"ServerName,omitempty"`
	Player               *string  `protobuf:"bytes,2,req,name=Player" json:"Player,omitempty"`
	Message              *string  `protobuf:"bytes,3,req,name=Message" json:"Message,omitempty"`
	ServerNameColor      *string  `protobuf:"bytes,4,opt,name=ServerNameColor,def=white" json:"ServerNameColor,omitempty"`
	PlayerColor          *string  `protobuf:"bytes,5,opt,name=PlayerColor,def=white" json:"PlayerColor,omitempty"`
	MessageColor         *string  `protobuf:"bytes,6,opt,name=MessageColor,def=white" json:"MessageColor,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Message) Reset()         { *m = Message{} }
func (m *Message) String() string { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()    {}
func (*Message) Descriptor() ([]byte, []int) {
	return fileDescriptor_c340e74c0f2a7ab8, []int{0}
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

const Default_Message_ServerNameColor string = "white"
const Default_Message_PlayerColor string = "white"
const Default_Message_MessageColor string = "white"

func (m *Message) GetServerName() string {
	if m != nil && m.ServerName != nil {
		return *m.ServerName
	}
	return ""
}

func (m *Message) GetPlayer() string {
	if m != nil && m.Player != nil {
		return *m.Player
	}
	return ""
}

func (m *Message) GetMessage() string {
	if m != nil && m.Message != nil {
		return *m.Message
	}
	return ""
}

func (m *Message) GetServerNameColor() string {
	if m != nil && m.ServerNameColor != nil {
		return *m.ServerNameColor
	}
	return Default_Message_ServerNameColor
}

func (m *Message) GetPlayerColor() string {
	if m != nil && m.PlayerColor != nil {
		return *m.PlayerColor
	}
	return Default_Message_PlayerColor
}

func (m *Message) GetMessageColor() string {
	if m != nil && m.MessageColor != nil {
		return *m.MessageColor
	}
	return Default_Message_MessageColor
}

func init() {
	proto.RegisterType((*Message)(nil), "ChatPlugin.Message")
}

func init() { proto.RegisterFile("chatBuf.proto", fileDescriptor_c340e74c0f2a7ab8) }

var fileDescriptor_c340e74c0f2a7ab8 = []byte{
	// 164 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4d, 0xce, 0x48, 0x2c,
	0x71, 0x2a, 0x4d, 0xd3, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x72, 0xce, 0x48, 0x2c, 0x09,
	0xc8, 0x29, 0x4d, 0xcf, 0xcc, 0x53, 0x7a, 0xc0, 0xc8, 0xc5, 0xee, 0x9b, 0x5a, 0x5c, 0x9c, 0x98,
	0x9e, 0x2a, 0x24, 0xc7, 0xc5, 0x15, 0x9c, 0x5a, 0x54, 0x96, 0x5a, 0xe4, 0x97, 0x98, 0x9b, 0x2a,
	0xc1, 0xa8, 0xc0, 0xa4, 0xc1, 0x19, 0x84, 0x24, 0x22, 0x24, 0xc6, 0xc5, 0x16, 0x90, 0x93, 0x58,
	0x99, 0x5a, 0x24, 0xc1, 0x04, 0x96, 0x83, 0xf2, 0x84, 0x24, 0xe0, 0x46, 0x48, 0x30, 0x83, 0x25,
	0xe0, 0x26, 0xea, 0x73, 0xf1, 0x23, 0xf4, 0x3b, 0xe7, 0xe7, 0xe4, 0x17, 0x49, 0xb0, 0x28, 0x30,
	0x6a, 0x70, 0x5a, 0xb1, 0x96, 0x67, 0x64, 0x96, 0xa4, 0x06, 0xa1, 0xcb, 0x0a, 0xa9, 0x73, 0x71,
	0x43, 0x0c, 0x85, 0x28, 0x66, 0x45, 0x56, 0x8c, 0x2c, 0x23, 0xa4, 0xc9, 0xc5, 0x03, 0xb5, 0x04,
	0xa2, 0x92, 0x0d, 0x59, 0x25, 0x8a, 0x14, 0x20, 0x00, 0x00, 0xff, 0xff, 0xa2, 0xcf, 0x04, 0xc7,
	0xfe, 0x00, 0x00, 0x00,
}
