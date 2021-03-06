// Code generated by protoc-gen-go. DO NOT EDIT.
// source: battle.proto

package netproto

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

// 协议id定义
type EBMsgIds int32

const (
	EBMsgIds_EBCS_None EBMsgIds = 0
	EBMsgIds_EBCS_Tick EBMsgIds = 1
	EBMsgIds_EBSC_Tick EBMsgIds = 2
)

var EBMsgIds_name = map[int32]string{
	0: "EBCS_None",
	1: "EBCS_Tick",
	2: "EBSC_Tick",
}
var EBMsgIds_value = map[string]int32{
	"EBCS_None": 0,
	"EBCS_Tick": 1,
	"EBSC_Tick": 2,
}

func (x EBMsgIds) String() string {
	return proto.EnumName(EBMsgIds_name, int32(x))
}
func (EBMsgIds) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_battle_90f3ea8b54fb0778, []int{0}
}

// 心跳
type BCS_Tick struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BCS_Tick) Reset()         { *m = BCS_Tick{} }
func (m *BCS_Tick) String() string { return proto.CompactTextString(m) }
func (*BCS_Tick) ProtoMessage()    {}
func (*BCS_Tick) Descriptor() ([]byte, []int) {
	return fileDescriptor_battle_90f3ea8b54fb0778, []int{0}
}
func (m *BCS_Tick) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BCS_Tick.Unmarshal(m, b)
}
func (m *BCS_Tick) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BCS_Tick.Marshal(b, m, deterministic)
}
func (dst *BCS_Tick) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BCS_Tick.Merge(dst, src)
}
func (m *BCS_Tick) XXX_Size() int {
	return xxx_messageInfo_BCS_Tick.Size(m)
}
func (m *BCS_Tick) XXX_DiscardUnknown() {
	xxx_messageInfo_BCS_Tick.DiscardUnknown(m)
}

var xxx_messageInfo_BCS_Tick proto.InternalMessageInfo

type BSC_Tick struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BSC_Tick) Reset()         { *m = BSC_Tick{} }
func (m *BSC_Tick) String() string { return proto.CompactTextString(m) }
func (*BSC_Tick) ProtoMessage()    {}
func (*BSC_Tick) Descriptor() ([]byte, []int) {
	return fileDescriptor_battle_90f3ea8b54fb0778, []int{1}
}
func (m *BSC_Tick) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BSC_Tick.Unmarshal(m, b)
}
func (m *BSC_Tick) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BSC_Tick.Marshal(b, m, deterministic)
}
func (dst *BSC_Tick) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BSC_Tick.Merge(dst, src)
}
func (m *BSC_Tick) XXX_Size() int {
	return xxx_messageInfo_BSC_Tick.Size(m)
}
func (m *BSC_Tick) XXX_DiscardUnknown() {
	xxx_messageInfo_BSC_Tick.DiscardUnknown(m)
}

var xxx_messageInfo_BSC_Tick proto.InternalMessageInfo

type BCS_None struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BCS_None) Reset()         { *m = BCS_None{} }
func (m *BCS_None) String() string { return proto.CompactTextString(m) }
func (*BCS_None) ProtoMessage()    {}
func (*BCS_None) Descriptor() ([]byte, []int) {
	return fileDescriptor_battle_90f3ea8b54fb0778, []int{2}
}
func (m *BCS_None) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BCS_None.Unmarshal(m, b)
}
func (m *BCS_None) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BCS_None.Marshal(b, m, deterministic)
}
func (dst *BCS_None) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BCS_None.Merge(dst, src)
}
func (m *BCS_None) XXX_Size() int {
	return xxx_messageInfo_BCS_None.Size(m)
}
func (m *BCS_None) XXX_DiscardUnknown() {
	xxx_messageInfo_BCS_None.DiscardUnknown(m)
}

var xxx_messageInfo_BCS_None proto.InternalMessageInfo

func init() {
	proto.RegisterType((*BCS_Tick)(nil), "netproto.BCS_Tick")
	proto.RegisterType((*BSC_Tick)(nil), "netproto.BSC_Tick")
	proto.RegisterType((*BCS_None)(nil), "netproto.BCS_None")
	proto.RegisterEnum("netproto.EBMsgIds", EBMsgIds_name, EBMsgIds_value)
}

func init() { proto.RegisterFile("battle.proto", fileDescriptor_battle_90f3ea8b54fb0778) }

var fileDescriptor_battle_90f3ea8b54fb0778 = []byte{
	// 107 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x49, 0x4a, 0x2c, 0x29,
	0xc9, 0x49, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0xc8, 0x4b, 0x2d, 0x01, 0xb3, 0x94,
	0xb8, 0xb8, 0x38, 0x9c, 0x9c, 0x83, 0xe3, 0x43, 0x32, 0x93, 0xb3, 0xc1, 0xec, 0x60, 0x67, 0x04,
	0xdb, 0x39, 0x38, 0xde, 0x2f, 0x3f, 0x2f, 0x55, 0xcb, 0x9c, 0x8b, 0xc3, 0xd5, 0xc9, 0xb7, 0x38,
	0xdd, 0x33, 0xa5, 0x58, 0x88, 0x97, 0x8b, 0xd3, 0x15, 0x26, 0x21, 0xc0, 0x00, 0xe7, 0x82, 0xf4,
	0x08, 0x30, 0x42, 0xb8, 0x50, 0x23, 0x04, 0x98, 0x92, 0xd8, 0xc0, 0x76, 0x18, 0x03, 0x02, 0x00,
	0x00, 0xff, 0xff, 0x18, 0x9a, 0x1b, 0x5a, 0x7d, 0x00, 0x00, 0x00,
}
