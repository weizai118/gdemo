// Code generated by protoc-gen-go. DO NOT EDIT.
// source: cluster.proto

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

type ECMsgIds int32

const (
	ECMsgIds_ESS_None                  ECMsgIds = 0
	ECMsgIds_ESS_Tick                  ECMsgIds = 1
	ECMsgIds_ESS_ClusterRegister       ECMsgIds = 10
	ECMsgIds_ESS_KickPlayer            ECMsgIds = 11
	ECMsgIds_ESS_Player2Player         ECMsgIds = 12
	ECMsgIds_ESS_P2PRedirect           ECMsgIds = 13
	ECMsgIds_ESS_ForwardPlayerMsg      ECMsgIds = 14
	ECMsgIds_ESS_ForwardPlayerGuideMsg ECMsgIds = 15
)

var ECMsgIds_name = map[int32]string{
	0:  "ESS_None",
	1:  "ESS_Tick",
	10: "ESS_ClusterRegister",
	11: "ESS_KickPlayer",
	12: "ESS_Player2Player",
	13: "ESS_P2PRedirect",
	14: "ESS_ForwardPlayerMsg",
	15: "ESS_ForwardPlayerGuideMsg",
}
var ECMsgIds_value = map[string]int32{
	"ESS_None":                  0,
	"ESS_Tick":                  1,
	"ESS_ClusterRegister":       10,
	"ESS_KickPlayer":            11,
	"ESS_Player2Player":         12,
	"ESS_P2PRedirect":           13,
	"ESS_ForwardPlayerMsg":      14,
	"ESS_ForwardPlayerGuideMsg": 15,
}

func (x ECMsgIds) String() string {
	return proto.EnumName(ECMsgIds_name, int32(x))
}
func (ECMsgIds) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_cluster_2eeb6c3d028c3ce9, []int{0}
}

type SS_Tick struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SS_Tick) Reset()         { *m = SS_Tick{} }
func (m *SS_Tick) String() string { return proto.CompactTextString(m) }
func (*SS_Tick) ProtoMessage()    {}
func (*SS_Tick) Descriptor() ([]byte, []int) {
	return fileDescriptor_cluster_2eeb6c3d028c3ce9, []int{0}
}
func (m *SS_Tick) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SS_Tick.Unmarshal(m, b)
}
func (m *SS_Tick) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SS_Tick.Marshal(b, m, deterministic)
}
func (dst *SS_Tick) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SS_Tick.Merge(dst, src)
}
func (m *SS_Tick) XXX_Size() int {
	return xxx_messageInfo_SS_Tick.Size(m)
}
func (m *SS_Tick) XXX_DiscardUnknown() {
	xxx_messageInfo_SS_Tick.DiscardUnknown(m)
}

var xxx_messageInfo_SS_Tick proto.InternalMessageInfo

type SS_ClusterRegister struct {
	SignCode             string   `protobuf:"bytes,1,opt,name=signCode" json:"signCode,omitempty"`
	Uuid                 string   `protobuf:"bytes,2,opt,name=uuid" json:"uuid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SS_ClusterRegister) Reset()         { *m = SS_ClusterRegister{} }
func (m *SS_ClusterRegister) String() string { return proto.CompactTextString(m) }
func (*SS_ClusterRegister) ProtoMessage()    {}
func (*SS_ClusterRegister) Descriptor() ([]byte, []int) {
	return fileDescriptor_cluster_2eeb6c3d028c3ce9, []int{1}
}
func (m *SS_ClusterRegister) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SS_ClusterRegister.Unmarshal(m, b)
}
func (m *SS_ClusterRegister) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SS_ClusterRegister.Marshal(b, m, deterministic)
}
func (dst *SS_ClusterRegister) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SS_ClusterRegister.Merge(dst, src)
}
func (m *SS_ClusterRegister) XXX_Size() int {
	return xxx_messageInfo_SS_ClusterRegister.Size(m)
}
func (m *SS_ClusterRegister) XXX_DiscardUnknown() {
	xxx_messageInfo_SS_ClusterRegister.DiscardUnknown(m)
}

var xxx_messageInfo_SS_ClusterRegister proto.InternalMessageInfo

func (m *SS_ClusterRegister) GetSignCode() string {
	if m != nil {
		return m.SignCode
	}
	return ""
}

func (m *SS_ClusterRegister) GetUuid() string {
	if m != nil {
		return m.Uuid
	}
	return ""
}

type SS_KickPlayer struct {
	PlayerId             int64    `protobuf:"varint,1,opt,name=playerId" json:"playerId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SS_KickPlayer) Reset()         { *m = SS_KickPlayer{} }
func (m *SS_KickPlayer) String() string { return proto.CompactTextString(m) }
func (*SS_KickPlayer) ProtoMessage()    {}
func (*SS_KickPlayer) Descriptor() ([]byte, []int) {
	return fileDescriptor_cluster_2eeb6c3d028c3ce9, []int{2}
}
func (m *SS_KickPlayer) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SS_KickPlayer.Unmarshal(m, b)
}
func (m *SS_KickPlayer) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SS_KickPlayer.Marshal(b, m, deterministic)
}
func (dst *SS_KickPlayer) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SS_KickPlayer.Merge(dst, src)
}
func (m *SS_KickPlayer) XXX_Size() int {
	return xxx_messageInfo_SS_KickPlayer.Size(m)
}
func (m *SS_KickPlayer) XXX_DiscardUnknown() {
	xxx_messageInfo_SS_KickPlayer.DiscardUnknown(m)
}

var xxx_messageInfo_SS_KickPlayer proto.InternalMessageInfo

func (m *SS_KickPlayer) GetPlayerId() int64 {
	if m != nil {
		return m.PlayerId
	}
	return 0
}

// 玩家之间的消息
type SS_Player2Player struct {
	FromPlayerId         int64    `protobuf:"varint,1,opt,name=fromPlayerId" json:"fromPlayerId,omitempty"`
	ToPlayerId           int64    `protobuf:"varint,2,opt,name=toPlayerId" json:"toPlayerId,omitempty"`
	Data                 []byte   `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SS_Player2Player) Reset()         { *m = SS_Player2Player{} }
func (m *SS_Player2Player) String() string { return proto.CompactTextString(m) }
func (*SS_Player2Player) ProtoMessage()    {}
func (*SS_Player2Player) Descriptor() ([]byte, []int) {
	return fileDescriptor_cluster_2eeb6c3d028c3ce9, []int{3}
}
func (m *SS_Player2Player) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SS_Player2Player.Unmarshal(m, b)
}
func (m *SS_Player2Player) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SS_Player2Player.Marshal(b, m, deterministic)
}
func (dst *SS_Player2Player) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SS_Player2Player.Merge(dst, src)
}
func (m *SS_Player2Player) XXX_Size() int {
	return xxx_messageInfo_SS_Player2Player.Size(m)
}
func (m *SS_Player2Player) XXX_DiscardUnknown() {
	xxx_messageInfo_SS_Player2Player.DiscardUnknown(m)
}

var xxx_messageInfo_SS_Player2Player proto.InternalMessageInfo

func (m *SS_Player2Player) GetFromPlayerId() int64 {
	if m != nil {
		return m.FromPlayerId
	}
	return 0
}

func (m *SS_Player2Player) GetToPlayerId() int64 {
	if m != nil {
		return m.ToPlayerId
	}
	return 0
}

func (m *SS_Player2Player) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

// 玩家之间的消息重定向
type SS_P2PRedirect struct {
	FromPlayerId         int64    `protobuf:"varint,1,opt,name=fromPlayerId" json:"fromPlayerId,omitempty"`
	ToPlayerId           int64    `protobuf:"varint,2,opt,name=toPlayerId" json:"toPlayerId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SS_P2PRedirect) Reset()         { *m = SS_P2PRedirect{} }
func (m *SS_P2PRedirect) String() string { return proto.CompactTextString(m) }
func (*SS_P2PRedirect) ProtoMessage()    {}
func (*SS_P2PRedirect) Descriptor() ([]byte, []int) {
	return fileDescriptor_cluster_2eeb6c3d028c3ce9, []int{4}
}
func (m *SS_P2PRedirect) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SS_P2PRedirect.Unmarshal(m, b)
}
func (m *SS_P2PRedirect) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SS_P2PRedirect.Marshal(b, m, deterministic)
}
func (dst *SS_P2PRedirect) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SS_P2PRedirect.Merge(dst, src)
}
func (m *SS_P2PRedirect) XXX_Size() int {
	return xxx_messageInfo_SS_P2PRedirect.Size(m)
}
func (m *SS_P2PRedirect) XXX_DiscardUnknown() {
	xxx_messageInfo_SS_P2PRedirect.DiscardUnknown(m)
}

var xxx_messageInfo_SS_P2PRedirect proto.InternalMessageInfo

func (m *SS_P2PRedirect) GetFromPlayerId() int64 {
	if m != nil {
		return m.FromPlayerId
	}
	return 0
}

func (m *SS_P2PRedirect) GetToPlayerId() int64 {
	if m != nil {
		return m.ToPlayerId
	}
	return 0
}

type SS_ForwardPlayerMsg struct {
	PlayerId             int64    `protobuf:"varint,1,opt,name=playerId" json:"playerId,omitempty"`
	MsgId                int32    `protobuf:"varint,2,opt,name=msgId" json:"msgId,omitempty"`
	Msgs                 []byte   `protobuf:"bytes,3,opt,name=msgs,proto3" json:"msgs,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SS_ForwardPlayerMsg) Reset()         { *m = SS_ForwardPlayerMsg{} }
func (m *SS_ForwardPlayerMsg) String() string { return proto.CompactTextString(m) }
func (*SS_ForwardPlayerMsg) ProtoMessage()    {}
func (*SS_ForwardPlayerMsg) Descriptor() ([]byte, []int) {
	return fileDescriptor_cluster_2eeb6c3d028c3ce9, []int{5}
}
func (m *SS_ForwardPlayerMsg) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SS_ForwardPlayerMsg.Unmarshal(m, b)
}
func (m *SS_ForwardPlayerMsg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SS_ForwardPlayerMsg.Marshal(b, m, deterministic)
}
func (dst *SS_ForwardPlayerMsg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SS_ForwardPlayerMsg.Merge(dst, src)
}
func (m *SS_ForwardPlayerMsg) XXX_Size() int {
	return xxx_messageInfo_SS_ForwardPlayerMsg.Size(m)
}
func (m *SS_ForwardPlayerMsg) XXX_DiscardUnknown() {
	xxx_messageInfo_SS_ForwardPlayerMsg.DiscardUnknown(m)
}

var xxx_messageInfo_SS_ForwardPlayerMsg proto.InternalMessageInfo

func (m *SS_ForwardPlayerMsg) GetPlayerId() int64 {
	if m != nil {
		return m.PlayerId
	}
	return 0
}

func (m *SS_ForwardPlayerMsg) GetMsgId() int32 {
	if m != nil {
		return m.MsgId
	}
	return 0
}

func (m *SS_ForwardPlayerMsg) GetMsgs() []byte {
	if m != nil {
		return m.Msgs
	}
	return nil
}

type SS_ForwardPlayerGuideMsg struct {
	PlayerId             int64    `protobuf:"varint,1,opt,name=playerId" json:"playerId,omitempty"`
	GuideId              int64    `protobuf:"varint,2,opt,name=guideId" json:"guideId,omitempty"`
	MsgId                int32    `protobuf:"varint,3,opt,name=msgId" json:"msgId,omitempty"`
	Msgs                 []byte   `protobuf:"bytes,4,opt,name=msgs,proto3" json:"msgs,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SS_ForwardPlayerGuideMsg) Reset()         { *m = SS_ForwardPlayerGuideMsg{} }
func (m *SS_ForwardPlayerGuideMsg) String() string { return proto.CompactTextString(m) }
func (*SS_ForwardPlayerGuideMsg) ProtoMessage()    {}
func (*SS_ForwardPlayerGuideMsg) Descriptor() ([]byte, []int) {
	return fileDescriptor_cluster_2eeb6c3d028c3ce9, []int{6}
}
func (m *SS_ForwardPlayerGuideMsg) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SS_ForwardPlayerGuideMsg.Unmarshal(m, b)
}
func (m *SS_ForwardPlayerGuideMsg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SS_ForwardPlayerGuideMsg.Marshal(b, m, deterministic)
}
func (dst *SS_ForwardPlayerGuideMsg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SS_ForwardPlayerGuideMsg.Merge(dst, src)
}
func (m *SS_ForwardPlayerGuideMsg) XXX_Size() int {
	return xxx_messageInfo_SS_ForwardPlayerGuideMsg.Size(m)
}
func (m *SS_ForwardPlayerGuideMsg) XXX_DiscardUnknown() {
	xxx_messageInfo_SS_ForwardPlayerGuideMsg.DiscardUnknown(m)
}

var xxx_messageInfo_SS_ForwardPlayerGuideMsg proto.InternalMessageInfo

func (m *SS_ForwardPlayerGuideMsg) GetPlayerId() int64 {
	if m != nil {
		return m.PlayerId
	}
	return 0
}

func (m *SS_ForwardPlayerGuideMsg) GetGuideId() int64 {
	if m != nil {
		return m.GuideId
	}
	return 0
}

func (m *SS_ForwardPlayerGuideMsg) GetMsgId() int32 {
	if m != nil {
		return m.MsgId
	}
	return 0
}

func (m *SS_ForwardPlayerGuideMsg) GetMsgs() []byte {
	if m != nil {
		return m.Msgs
	}
	return nil
}

type SS_None struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SS_None) Reset()         { *m = SS_None{} }
func (m *SS_None) String() string { return proto.CompactTextString(m) }
func (*SS_None) ProtoMessage()    {}
func (*SS_None) Descriptor() ([]byte, []int) {
	return fileDescriptor_cluster_2eeb6c3d028c3ce9, []int{7}
}
func (m *SS_None) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SS_None.Unmarshal(m, b)
}
func (m *SS_None) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SS_None.Marshal(b, m, deterministic)
}
func (dst *SS_None) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SS_None.Merge(dst, src)
}
func (m *SS_None) XXX_Size() int {
	return xxx_messageInfo_SS_None.Size(m)
}
func (m *SS_None) XXX_DiscardUnknown() {
	xxx_messageInfo_SS_None.DiscardUnknown(m)
}

var xxx_messageInfo_SS_None proto.InternalMessageInfo

func init() {
	proto.RegisterType((*SS_Tick)(nil), "netproto.SS_Tick")
	proto.RegisterType((*SS_ClusterRegister)(nil), "netproto.SS_ClusterRegister")
	proto.RegisterType((*SS_KickPlayer)(nil), "netproto.SS_KickPlayer")
	proto.RegisterType((*SS_Player2Player)(nil), "netproto.SS_Player2Player")
	proto.RegisterType((*SS_P2PRedirect)(nil), "netproto.SS_P2PRedirect")
	proto.RegisterType((*SS_ForwardPlayerMsg)(nil), "netproto.SS_ForwardPlayerMsg")
	proto.RegisterType((*SS_ForwardPlayerGuideMsg)(nil), "netproto.SS_ForwardPlayerGuideMsg")
	proto.RegisterType((*SS_None)(nil), "netproto.SS_None")
	proto.RegisterEnum("netproto.ECMsgIds", ECMsgIds_name, ECMsgIds_value)
}

func init() { proto.RegisterFile("cluster.proto", fileDescriptor_cluster_2eeb6c3d028c3ce9) }

var fileDescriptor_cluster_2eeb6c3d028c3ce9 = []byte{
	// 362 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x52, 0x4d, 0x4f, 0xea, 0x40,
	0x14, 0x7d, 0xe5, 0xe3, 0x51, 0xee, 0x2b, 0x30, 0xef, 0xc2, 0xcb, 0xab, 0x26, 0x1a, 0x32, 0x2b,
	0xa2, 0x89, 0x0b, 0xfc, 0x09, 0x38, 0x1a, 0x63, 0x30, 0xa4, 0xc3, 0xce, 0x45, 0x53, 0xdb, 0xb1,
	0xa9, 0x7c, 0x94, 0x74, 0xda, 0x18, 0xfd, 0x65, 0xfe, 0x3c, 0x33, 0x53, 0xa8, 0x14, 0x51, 0x37,
	0xae, 0x7a, 0xce, 0xed, 0xe9, 0x3d, 0xa7, 0x33, 0x07, 0x5a, 0xfe, 0x3c, 0x93, 0xa9, 0x48, 0xce,
	0x56, 0x49, 0x9c, 0xc6, 0x68, 0x2e, 0x45, 0xaa, 0x11, 0x6d, 0x42, 0x83, 0x73, 0x77, 0x1a, 0xf9,
	0x33, 0x7a, 0x01, 0xc8, 0xb9, 0x3b, 0xca, 0x85, 0x8e, 0x08, 0x23, 0xf5, 0xc4, 0x43, 0x30, 0x65,
	0x14, 0x2e, 0x47, 0x71, 0x20, 0x6c, 0xa3, 0x6f, 0x0c, 0x9a, 0x4e, 0xc1, 0x11, 0xa1, 0x96, 0x65,
	0x51, 0x60, 0x57, 0xf4, 0x5c, 0x63, 0x7a, 0x0a, 0x2d, 0xce, 0xdd, 0x9b, 0xc8, 0x9f, 0x4d, 0xe6,
	0xde, 0x73, 0xbe, 0x60, 0xa5, 0xd1, 0x75, 0xa0, 0x17, 0x54, 0x9d, 0x82, 0xd3, 0x47, 0x20, 0x9c,
	0xbb, 0xb9, 0x70, 0xb8, 0xd6, 0x53, 0xb0, 0x1e, 0x92, 0x78, 0x31, 0x29, 0x7f, 0x53, 0x9a, 0xe1,
	0x31, 0x40, 0x1a, 0x17, 0x8a, 0x8a, 0x56, 0x6c, 0x4d, 0x54, 0xb0, 0xc0, 0x4b, 0x3d, 0xbb, 0xda,
	0x37, 0x06, 0x96, 0xa3, 0x31, 0x9d, 0x42, 0x5b, 0x79, 0x0d, 0x27, 0x8e, 0x08, 0xa2, 0x44, 0xf8,
	0xe9, 0x4f, 0x38, 0xd1, 0x3b, 0xe8, 0x72, 0xee, 0x5e, 0xc6, 0xc9, 0x93, 0x97, 0x04, 0xf9, 0x74,
	0x2c, 0xc3, 0xaf, 0x7e, 0x1a, 0x7b, 0x50, 0x5f, 0xc8, 0x70, 0xbd, 0xad, 0xee, 0xe4, 0x44, 0x45,
	0x5e, 0xc8, 0x50, 0x6e, 0x22, 0x2b, 0x4c, 0x5f, 0xc0, 0xde, 0x5d, 0x7e, 0x95, 0x45, 0x81, 0xf8,
	0xce, 0xc1, 0x86, 0x46, 0xa8, 0x74, 0x45, 0xe2, 0x0d, 0x7d, 0xf7, 0xae, 0xee, 0xf3, 0xae, 0x6d,
	0x79, 0xe7, 0xc5, 0xb8, 0x8d, 0x97, 0xe2, 0xe4, 0xd5, 0x00, 0x93, 0x8d, 0xc6, 0x4a, 0x2a, 0xd1,
	0x02, 0x93, 0xad, 0x5f, 0x90, 0x5f, 0x1b, 0xa6, 0xfa, 0x43, 0x0c, 0xfc, 0x0f, 0x5d, 0xf6, 0xb1,
	0x42, 0x04, 0x10, 0xa1, 0xcd, 0x4a, 0xad, 0x20, 0x7f, 0xf0, 0x1f, 0xfc, 0x65, 0xbb, 0x97, 0x4f,
	0x2c, 0xec, 0x42, 0x87, 0x95, 0xef, 0x89, 0xb4, 0xd0, 0x86, 0x1e, 0xdb, 0x73, 0xcc, 0xa4, 0x8d,
	0x47, 0x70, 0xc0, 0x3e, 0x3b, 0x23, 0xd2, 0xb9, 0xff, 0xad, 0x5b, 0x7e, 0xfe, 0x16, 0x00, 0x00,
	0xff, 0xff, 0xea, 0x51, 0x51, 0x65, 0x00, 0x03, 0x00, 0x00,
}
