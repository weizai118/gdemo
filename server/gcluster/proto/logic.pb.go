// Code generated by protoc-gen-go. DO NOT EDIT.
// source: logic.proto

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

// 错误码
type E_Code int32

const (
	E_Code_E_ERROR                 E_Code = 0
	E_Code_E_OK                    E_Code = 1
	E_Code_E_RELOGIN               E_Code = 2
	E_Code_E_NONE_EXIST            E_Code = 3
	E_Code_E_UNKNOWN               E_Code = 4
	E_Code_E_SERVER_INTERNAL_ERROR E_Code = 5
	E_Code_E_INVALID_PARAM         E_Code = 6
	E_Code_E_INVALID_OPT           E_Code = 7
	E_Code_E_LIMIT_GOLD            E_Code = 20
	E_Code_E_LIMIT_DIAMOND         E_Code = 21
	E_Code_E_LIMIT_TIMES           E_Code = 22
	E_Code_E_LIMIT_CHAPTER_LEVEL   E_Code = 23
	E_Code_E_LIMIT_PLAYER_LEVEL    E_Code = 24
)

var E_Code_name = map[int32]string{
	0:  "E_ERROR",
	1:  "E_OK",
	2:  "E_RELOGIN",
	3:  "E_NONE_EXIST",
	4:  "E_UNKNOWN",
	5:  "E_SERVER_INTERNAL_ERROR",
	6:  "E_INVALID_PARAM",
	7:  "E_INVALID_OPT",
	20: "E_LIMIT_GOLD",
	21: "E_LIMIT_DIAMOND",
	22: "E_LIMIT_TIMES",
	23: "E_LIMIT_CHAPTER_LEVEL",
	24: "E_LIMIT_PLAYER_LEVEL",
}
var E_Code_value = map[string]int32{
	"E_ERROR":                 0,
	"E_OK":                    1,
	"E_RELOGIN":               2,
	"E_NONE_EXIST":            3,
	"E_UNKNOWN":               4,
	"E_SERVER_INTERNAL_ERROR": 5,
	"E_INVALID_PARAM":         6,
	"E_INVALID_OPT":           7,
	"E_LIMIT_GOLD":            20,
	"E_LIMIT_DIAMOND":         21,
	"E_LIMIT_TIMES":           22,
	"E_LIMIT_CHAPTER_LEVEL":   23,
	"E_LIMIT_PLAYER_LEVEL":    24,
}

func (x E_Code) String() string {
	return proto.EnumName(E_Code_name, int32(x))
}
func (E_Code) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_logic_444f2d1a507ba940, []int{0}
}

// 前100号的消息留给系统用
type EMsgIds int32

const (
	EMsgIds_ECS_None EMsgIds = 0
	// 心跳
	EMsgIds_ECS_Tick             EMsgIds = 101
	EMsgIds_ESC_Tick             EMsgIds = 102
	EMsgIds_ECS_Login            EMsgIds = 103
	EMsgIds_ESC_Login            EMsgIds = 104
	EMsgIds_ECS_ShopAddDiamond   EMsgIds = 105
	EMsgIds_ESC_ShopAddDiamond   EMsgIds = 106
	EMsgIds_ECS_ShopBuyGold      EMsgIds = 107
	EMsgIds_ESC_ShopBuyGold      EMsgIds = 108
	EMsgIds_ECS_HeroDraw         EMsgIds = 109
	EMsgIds_ESC_HeroDraw         EMsgIds = 110
	EMsgIds_ECS_HeroUplevel      EMsgIds = 111
	EMsgIds_ESC_HeroUpLevel      EMsgIds = 112
	EMsgIds_ECS_HeroEquip        EMsgIds = 113
	EMsgIds_ESC_HeroEquip        EMsgIds = 114
	EMsgIds_ECS_ChapterGetPrize  EMsgIds = 115
	EMsgIds_ESC_ChapterGetPrize  EMsgIds = 116
	EMsgIds_ECS_ChapterChallenge EMsgIds = 117
	EMsgIds_ESC_ChapterChallenge EMsgIds = 118
)

var EMsgIds_name = map[int32]string{
	0:   "ECS_None",
	101: "ECS_Tick",
	102: "ESC_Tick",
	103: "ECS_Login",
	104: "ESC_Login",
	105: "ECS_ShopAddDiamond",
	106: "ESC_ShopAddDiamond",
	107: "ECS_ShopBuyGold",
	108: "ESC_ShopBuyGold",
	109: "ECS_HeroDraw",
	110: "ESC_HeroDraw",
	111: "ECS_HeroUplevel",
	112: "ESC_HeroUpLevel",
	113: "ECS_HeroEquip",
	114: "ESC_HeroEquip",
	115: "ECS_ChapterGetPrize",
	116: "ESC_ChapterGetPrize",
	117: "ECS_ChapterChallenge",
	118: "ESC_ChapterChallenge",
}
var EMsgIds_value = map[string]int32{
	"ECS_None":             0,
	"ECS_Tick":             101,
	"ESC_Tick":             102,
	"ECS_Login":            103,
	"ESC_Login":            104,
	"ECS_ShopAddDiamond":   105,
	"ESC_ShopAddDiamond":   106,
	"ECS_ShopBuyGold":      107,
	"ESC_ShopBuyGold":      108,
	"ECS_HeroDraw":         109,
	"ESC_HeroDraw":         110,
	"ECS_HeroUplevel":      111,
	"ESC_HeroUpLevel":      112,
	"ECS_HeroEquip":        113,
	"ESC_HeroEquip":        114,
	"ECS_ChapterGetPrize":  115,
	"ESC_ChapterGetPrize":  116,
	"ECS_ChapterChallenge": 117,
	"ESC_ChapterChallenge": 118,
}

func (x EMsgIds) String() string {
	return proto.EnumName(EMsgIds_name, int32(x))
}
func (EMsgIds) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_logic_444f2d1a507ba940, []int{1}
}

// 心跳
type CS_Tick struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CS_Tick) Reset()         { *m = CS_Tick{} }
func (m *CS_Tick) String() string { return proto.CompactTextString(m) }
func (*CS_Tick) ProtoMessage()    {}
func (*CS_Tick) Descriptor() ([]byte, []int) {
	return fileDescriptor_logic_444f2d1a507ba940, []int{0}
}
func (m *CS_Tick) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CS_Tick.Unmarshal(m, b)
}
func (m *CS_Tick) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CS_Tick.Marshal(b, m, deterministic)
}
func (dst *CS_Tick) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CS_Tick.Merge(dst, src)
}
func (m *CS_Tick) XXX_Size() int {
	return xxx_messageInfo_CS_Tick.Size(m)
}
func (m *CS_Tick) XXX_DiscardUnknown() {
	xxx_messageInfo_CS_Tick.DiscardUnknown(m)
}

var xxx_messageInfo_CS_Tick proto.InternalMessageInfo

type SC_Tick struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SC_Tick) Reset()         { *m = SC_Tick{} }
func (m *SC_Tick) String() string { return proto.CompactTextString(m) }
func (*SC_Tick) ProtoMessage()    {}
func (*SC_Tick) Descriptor() ([]byte, []int) {
	return fileDescriptor_logic_444f2d1a507ba940, []int{1}
}
func (m *SC_Tick) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SC_Tick.Unmarshal(m, b)
}
func (m *SC_Tick) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SC_Tick.Marshal(b, m, deterministic)
}
func (dst *SC_Tick) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SC_Tick.Merge(dst, src)
}
func (m *SC_Tick) XXX_Size() int {
	return xxx_messageInfo_SC_Tick.Size(m)
}
func (m *SC_Tick) XXX_DiscardUnknown() {
	xxx_messageInfo_SC_Tick.DiscardUnknown(m)
}

var xxx_messageInfo_SC_Tick proto.InternalMessageInfo

// 登录
type CS_Login struct {
	Token                string   `protobuf:"bytes,1,opt,name=token" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CS_Login) Reset()         { *m = CS_Login{} }
func (m *CS_Login) String() string { return proto.CompactTextString(m) }
func (*CS_Login) ProtoMessage()    {}
func (*CS_Login) Descriptor() ([]byte, []int) {
	return fileDescriptor_logic_444f2d1a507ba940, []int{2}
}
func (m *CS_Login) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CS_Login.Unmarshal(m, b)
}
func (m *CS_Login) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CS_Login.Marshal(b, m, deterministic)
}
func (dst *CS_Login) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CS_Login.Merge(dst, src)
}
func (m *CS_Login) XXX_Size() int {
	return xxx_messageInfo_CS_Login.Size(m)
}
func (m *CS_Login) XXX_DiscardUnknown() {
	xxx_messageInfo_CS_Login.DiscardUnknown(m)
}

var xxx_messageInfo_CS_Login proto.InternalMessageInfo

func (m *CS_Login) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type SC_Login struct {
	Code                 E_Code   `protobuf:"varint,1,opt,name=code,enum=netproto.E_Code" json:"code,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SC_Login) Reset()         { *m = SC_Login{} }
func (m *SC_Login) String() string { return proto.CompactTextString(m) }
func (*SC_Login) ProtoMessage()    {}
func (*SC_Login) Descriptor() ([]byte, []int) {
	return fileDescriptor_logic_444f2d1a507ba940, []int{3}
}
func (m *SC_Login) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SC_Login.Unmarshal(m, b)
}
func (m *SC_Login) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SC_Login.Marshal(b, m, deterministic)
}
func (dst *SC_Login) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SC_Login.Merge(dst, src)
}
func (m *SC_Login) XXX_Size() int {
	return xxx_messageInfo_SC_Login.Size(m)
}
func (m *SC_Login) XXX_DiscardUnknown() {
	xxx_messageInfo_SC_Login.DiscardUnknown(m)
}

var xxx_messageInfo_SC_Login proto.InternalMessageInfo

func (m *SC_Login) GetCode() E_Code {
	if m != nil {
		return m.Code
	}
	return E_Code_E_ERROR
}

// 冲钻石
type CS_ShopAddDiamond struct {
	Idx                  int32    `protobuf:"varint,1,opt,name=idx" json:"idx,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CS_ShopAddDiamond) Reset()         { *m = CS_ShopAddDiamond{} }
func (m *CS_ShopAddDiamond) String() string { return proto.CompactTextString(m) }
func (*CS_ShopAddDiamond) ProtoMessage()    {}
func (*CS_ShopAddDiamond) Descriptor() ([]byte, []int) {
	return fileDescriptor_logic_444f2d1a507ba940, []int{4}
}
func (m *CS_ShopAddDiamond) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CS_ShopAddDiamond.Unmarshal(m, b)
}
func (m *CS_ShopAddDiamond) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CS_ShopAddDiamond.Marshal(b, m, deterministic)
}
func (dst *CS_ShopAddDiamond) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CS_ShopAddDiamond.Merge(dst, src)
}
func (m *CS_ShopAddDiamond) XXX_Size() int {
	return xxx_messageInfo_CS_ShopAddDiamond.Size(m)
}
func (m *CS_ShopAddDiamond) XXX_DiscardUnknown() {
	xxx_messageInfo_CS_ShopAddDiamond.DiscardUnknown(m)
}

var xxx_messageInfo_CS_ShopAddDiamond proto.InternalMessageInfo

func (m *CS_ShopAddDiamond) GetIdx() int32 {
	if m != nil {
		return m.Idx
	}
	return 0
}

type SC_ShopAddDiamond struct {
	Code                 E_Code   `protobuf:"varint,1,opt,name=code,enum=netproto.E_Code" json:"code,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SC_ShopAddDiamond) Reset()         { *m = SC_ShopAddDiamond{} }
func (m *SC_ShopAddDiamond) String() string { return proto.CompactTextString(m) }
func (*SC_ShopAddDiamond) ProtoMessage()    {}
func (*SC_ShopAddDiamond) Descriptor() ([]byte, []int) {
	return fileDescriptor_logic_444f2d1a507ba940, []int{5}
}
func (m *SC_ShopAddDiamond) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SC_ShopAddDiamond.Unmarshal(m, b)
}
func (m *SC_ShopAddDiamond) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SC_ShopAddDiamond.Marshal(b, m, deterministic)
}
func (dst *SC_ShopAddDiamond) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SC_ShopAddDiamond.Merge(dst, src)
}
func (m *SC_ShopAddDiamond) XXX_Size() int {
	return xxx_messageInfo_SC_ShopAddDiamond.Size(m)
}
func (m *SC_ShopAddDiamond) XXX_DiscardUnknown() {
	xxx_messageInfo_SC_ShopAddDiamond.DiscardUnknown(m)
}

var xxx_messageInfo_SC_ShopAddDiamond proto.InternalMessageInfo

func (m *SC_ShopAddDiamond) GetCode() E_Code {
	if m != nil {
		return m.Code
	}
	return E_Code_E_ERROR
}

// 购买金币
type CS_ShopBuyGold struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CS_ShopBuyGold) Reset()         { *m = CS_ShopBuyGold{} }
func (m *CS_ShopBuyGold) String() string { return proto.CompactTextString(m) }
func (*CS_ShopBuyGold) ProtoMessage()    {}
func (*CS_ShopBuyGold) Descriptor() ([]byte, []int) {
	return fileDescriptor_logic_444f2d1a507ba940, []int{6}
}
func (m *CS_ShopBuyGold) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CS_ShopBuyGold.Unmarshal(m, b)
}
func (m *CS_ShopBuyGold) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CS_ShopBuyGold.Marshal(b, m, deterministic)
}
func (dst *CS_ShopBuyGold) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CS_ShopBuyGold.Merge(dst, src)
}
func (m *CS_ShopBuyGold) XXX_Size() int {
	return xxx_messageInfo_CS_ShopBuyGold.Size(m)
}
func (m *CS_ShopBuyGold) XXX_DiscardUnknown() {
	xxx_messageInfo_CS_ShopBuyGold.DiscardUnknown(m)
}

var xxx_messageInfo_CS_ShopBuyGold proto.InternalMessageInfo

type SC_ShopBuyGold struct {
	Code                 E_Code   `protobuf:"varint,1,opt,name=code,enum=netproto.E_Code" json:"code,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SC_ShopBuyGold) Reset()         { *m = SC_ShopBuyGold{} }
func (m *SC_ShopBuyGold) String() string { return proto.CompactTextString(m) }
func (*SC_ShopBuyGold) ProtoMessage()    {}
func (*SC_ShopBuyGold) Descriptor() ([]byte, []int) {
	return fileDescriptor_logic_444f2d1a507ba940, []int{7}
}
func (m *SC_ShopBuyGold) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SC_ShopBuyGold.Unmarshal(m, b)
}
func (m *SC_ShopBuyGold) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SC_ShopBuyGold.Marshal(b, m, deterministic)
}
func (dst *SC_ShopBuyGold) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SC_ShopBuyGold.Merge(dst, src)
}
func (m *SC_ShopBuyGold) XXX_Size() int {
	return xxx_messageInfo_SC_ShopBuyGold.Size(m)
}
func (m *SC_ShopBuyGold) XXX_DiscardUnknown() {
	xxx_messageInfo_SC_ShopBuyGold.DiscardUnknown(m)
}

var xxx_messageInfo_SC_ShopBuyGold proto.InternalMessageInfo

func (m *SC_ShopBuyGold) GetCode() E_Code {
	if m != nil {
		return m.Code
	}
	return E_Code_E_ERROR
}

// 抽英雄
type CS_HeroDraw struct {
	Num                  int32    `protobuf:"varint,1,opt,name=num" json:"num,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CS_HeroDraw) Reset()         { *m = CS_HeroDraw{} }
func (m *CS_HeroDraw) String() string { return proto.CompactTextString(m) }
func (*CS_HeroDraw) ProtoMessage()    {}
func (*CS_HeroDraw) Descriptor() ([]byte, []int) {
	return fileDescriptor_logic_444f2d1a507ba940, []int{8}
}
func (m *CS_HeroDraw) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CS_HeroDraw.Unmarshal(m, b)
}
func (m *CS_HeroDraw) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CS_HeroDraw.Marshal(b, m, deterministic)
}
func (dst *CS_HeroDraw) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CS_HeroDraw.Merge(dst, src)
}
func (m *CS_HeroDraw) XXX_Size() int {
	return xxx_messageInfo_CS_HeroDraw.Size(m)
}
func (m *CS_HeroDraw) XXX_DiscardUnknown() {
	xxx_messageInfo_CS_HeroDraw.DiscardUnknown(m)
}

var xxx_messageInfo_CS_HeroDraw proto.InternalMessageInfo

func (m *CS_HeroDraw) GetNum() int32 {
	if m != nil {
		return m.Num
	}
	return 0
}

type SC_HeroDraw struct {
	Code                 E_Code   `protobuf:"varint,1,opt,name=code,enum=netproto.E_Code" json:"code,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SC_HeroDraw) Reset()         { *m = SC_HeroDraw{} }
func (m *SC_HeroDraw) String() string { return proto.CompactTextString(m) }
func (*SC_HeroDraw) ProtoMessage()    {}
func (*SC_HeroDraw) Descriptor() ([]byte, []int) {
	return fileDescriptor_logic_444f2d1a507ba940, []int{9}
}
func (m *SC_HeroDraw) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SC_HeroDraw.Unmarshal(m, b)
}
func (m *SC_HeroDraw) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SC_HeroDraw.Marshal(b, m, deterministic)
}
func (dst *SC_HeroDraw) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SC_HeroDraw.Merge(dst, src)
}
func (m *SC_HeroDraw) XXX_Size() int {
	return xxx_messageInfo_SC_HeroDraw.Size(m)
}
func (m *SC_HeroDraw) XXX_DiscardUnknown() {
	xxx_messageInfo_SC_HeroDraw.DiscardUnknown(m)
}

var xxx_messageInfo_SC_HeroDraw proto.InternalMessageInfo

func (m *SC_HeroDraw) GetCode() E_Code {
	if m != nil {
		return m.Code
	}
	return E_Code_E_ERROR
}

// 升级英雄等级
type CS_HeroUplevel struct {
	HeroId               int32    `protobuf:"varint,1,opt,name=heroId" json:"heroId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CS_HeroUplevel) Reset()         { *m = CS_HeroUplevel{} }
func (m *CS_HeroUplevel) String() string { return proto.CompactTextString(m) }
func (*CS_HeroUplevel) ProtoMessage()    {}
func (*CS_HeroUplevel) Descriptor() ([]byte, []int) {
	return fileDescriptor_logic_444f2d1a507ba940, []int{10}
}
func (m *CS_HeroUplevel) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CS_HeroUplevel.Unmarshal(m, b)
}
func (m *CS_HeroUplevel) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CS_HeroUplevel.Marshal(b, m, deterministic)
}
func (dst *CS_HeroUplevel) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CS_HeroUplevel.Merge(dst, src)
}
func (m *CS_HeroUplevel) XXX_Size() int {
	return xxx_messageInfo_CS_HeroUplevel.Size(m)
}
func (m *CS_HeroUplevel) XXX_DiscardUnknown() {
	xxx_messageInfo_CS_HeroUplevel.DiscardUnknown(m)
}

var xxx_messageInfo_CS_HeroUplevel proto.InternalMessageInfo

func (m *CS_HeroUplevel) GetHeroId() int32 {
	if m != nil {
		return m.HeroId
	}
	return 0
}

type SC_HeroUpLevel struct {
	Code                 E_Code   `protobuf:"varint,1,opt,name=code,enum=netproto.E_Code" json:"code,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SC_HeroUpLevel) Reset()         { *m = SC_HeroUpLevel{} }
func (m *SC_HeroUpLevel) String() string { return proto.CompactTextString(m) }
func (*SC_HeroUpLevel) ProtoMessage()    {}
func (*SC_HeroUpLevel) Descriptor() ([]byte, []int) {
	return fileDescriptor_logic_444f2d1a507ba940, []int{11}
}
func (m *SC_HeroUpLevel) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SC_HeroUpLevel.Unmarshal(m, b)
}
func (m *SC_HeroUpLevel) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SC_HeroUpLevel.Marshal(b, m, deterministic)
}
func (dst *SC_HeroUpLevel) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SC_HeroUpLevel.Merge(dst, src)
}
func (m *SC_HeroUpLevel) XXX_Size() int {
	return xxx_messageInfo_SC_HeroUpLevel.Size(m)
}
func (m *SC_HeroUpLevel) XXX_DiscardUnknown() {
	xxx_messageInfo_SC_HeroUpLevel.DiscardUnknown(m)
}

var xxx_messageInfo_SC_HeroUpLevel proto.InternalMessageInfo

func (m *SC_HeroUpLevel) GetCode() E_Code {
	if m != nil {
		return m.Code
	}
	return E_Code_E_ERROR
}

// 装备英雄
type CS_HeroEquip struct {
	HeroId               int32    `protobuf:"varint,1,opt,name=heroId" json:"heroId,omitempty"`
	EquipId              int32    `protobuf:"varint,2,opt,name=equipId" json:"equipId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CS_HeroEquip) Reset()         { *m = CS_HeroEquip{} }
func (m *CS_HeroEquip) String() string { return proto.CompactTextString(m) }
func (*CS_HeroEquip) ProtoMessage()    {}
func (*CS_HeroEquip) Descriptor() ([]byte, []int) {
	return fileDescriptor_logic_444f2d1a507ba940, []int{12}
}
func (m *CS_HeroEquip) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CS_HeroEquip.Unmarshal(m, b)
}
func (m *CS_HeroEquip) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CS_HeroEquip.Marshal(b, m, deterministic)
}
func (dst *CS_HeroEquip) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CS_HeroEquip.Merge(dst, src)
}
func (m *CS_HeroEquip) XXX_Size() int {
	return xxx_messageInfo_CS_HeroEquip.Size(m)
}
func (m *CS_HeroEquip) XXX_DiscardUnknown() {
	xxx_messageInfo_CS_HeroEquip.DiscardUnknown(m)
}

var xxx_messageInfo_CS_HeroEquip proto.InternalMessageInfo

func (m *CS_HeroEquip) GetHeroId() int32 {
	if m != nil {
		return m.HeroId
	}
	return 0
}

func (m *CS_HeroEquip) GetEquipId() int32 {
	if m != nil {
		return m.EquipId
	}
	return 0
}

type SC_HeroEquip struct {
	Code                 E_Code   `protobuf:"varint,1,opt,name=code,enum=netproto.E_Code" json:"code,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SC_HeroEquip) Reset()         { *m = SC_HeroEquip{} }
func (m *SC_HeroEquip) String() string { return proto.CompactTextString(m) }
func (*SC_HeroEquip) ProtoMessage()    {}
func (*SC_HeroEquip) Descriptor() ([]byte, []int) {
	return fileDescriptor_logic_444f2d1a507ba940, []int{13}
}
func (m *SC_HeroEquip) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SC_HeroEquip.Unmarshal(m, b)
}
func (m *SC_HeroEquip) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SC_HeroEquip.Marshal(b, m, deterministic)
}
func (dst *SC_HeroEquip) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SC_HeroEquip.Merge(dst, src)
}
func (m *SC_HeroEquip) XXX_Size() int {
	return xxx_messageInfo_SC_HeroEquip.Size(m)
}
func (m *SC_HeroEquip) XXX_DiscardUnknown() {
	xxx_messageInfo_SC_HeroEquip.DiscardUnknown(m)
}

var xxx_messageInfo_SC_HeroEquip proto.InternalMessageInfo

func (m *SC_HeroEquip) GetCode() E_Code {
	if m != nil {
		return m.Code
	}
	return E_Code_E_ERROR
}

// 领取挂机奖励
type CS_ChapterGetPrize struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CS_ChapterGetPrize) Reset()         { *m = CS_ChapterGetPrize{} }
func (m *CS_ChapterGetPrize) String() string { return proto.CompactTextString(m) }
func (*CS_ChapterGetPrize) ProtoMessage()    {}
func (*CS_ChapterGetPrize) Descriptor() ([]byte, []int) {
	return fileDescriptor_logic_444f2d1a507ba940, []int{14}
}
func (m *CS_ChapterGetPrize) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CS_ChapterGetPrize.Unmarshal(m, b)
}
func (m *CS_ChapterGetPrize) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CS_ChapterGetPrize.Marshal(b, m, deterministic)
}
func (dst *CS_ChapterGetPrize) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CS_ChapterGetPrize.Merge(dst, src)
}
func (m *CS_ChapterGetPrize) XXX_Size() int {
	return xxx_messageInfo_CS_ChapterGetPrize.Size(m)
}
func (m *CS_ChapterGetPrize) XXX_DiscardUnknown() {
	xxx_messageInfo_CS_ChapterGetPrize.DiscardUnknown(m)
}

var xxx_messageInfo_CS_ChapterGetPrize proto.InternalMessageInfo

type SC_ChapterGetPrize struct {
	Code                 E_Code   `protobuf:"varint,1,opt,name=code,enum=netproto.E_Code" json:"code,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SC_ChapterGetPrize) Reset()         { *m = SC_ChapterGetPrize{} }
func (m *SC_ChapterGetPrize) String() string { return proto.CompactTextString(m) }
func (*SC_ChapterGetPrize) ProtoMessage()    {}
func (*SC_ChapterGetPrize) Descriptor() ([]byte, []int) {
	return fileDescriptor_logic_444f2d1a507ba940, []int{15}
}
func (m *SC_ChapterGetPrize) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SC_ChapterGetPrize.Unmarshal(m, b)
}
func (m *SC_ChapterGetPrize) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SC_ChapterGetPrize.Marshal(b, m, deterministic)
}
func (dst *SC_ChapterGetPrize) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SC_ChapterGetPrize.Merge(dst, src)
}
func (m *SC_ChapterGetPrize) XXX_Size() int {
	return xxx_messageInfo_SC_ChapterGetPrize.Size(m)
}
func (m *SC_ChapterGetPrize) XXX_DiscardUnknown() {
	xxx_messageInfo_SC_ChapterGetPrize.DiscardUnknown(m)
}

var xxx_messageInfo_SC_ChapterGetPrize proto.InternalMessageInfo

func (m *SC_ChapterGetPrize) GetCode() E_Code {
	if m != nil {
		return m.Code
	}
	return E_Code_E_ERROR
}

// 关卡挑战
type CS_ChapterChallenge struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CS_ChapterChallenge) Reset()         { *m = CS_ChapterChallenge{} }
func (m *CS_ChapterChallenge) String() string { return proto.CompactTextString(m) }
func (*CS_ChapterChallenge) ProtoMessage()    {}
func (*CS_ChapterChallenge) Descriptor() ([]byte, []int) {
	return fileDescriptor_logic_444f2d1a507ba940, []int{16}
}
func (m *CS_ChapterChallenge) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CS_ChapterChallenge.Unmarshal(m, b)
}
func (m *CS_ChapterChallenge) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CS_ChapterChallenge.Marshal(b, m, deterministic)
}
func (dst *CS_ChapterChallenge) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CS_ChapterChallenge.Merge(dst, src)
}
func (m *CS_ChapterChallenge) XXX_Size() int {
	return xxx_messageInfo_CS_ChapterChallenge.Size(m)
}
func (m *CS_ChapterChallenge) XXX_DiscardUnknown() {
	xxx_messageInfo_CS_ChapterChallenge.DiscardUnknown(m)
}

var xxx_messageInfo_CS_ChapterChallenge proto.InternalMessageInfo

type SC_ChapterChallenge struct {
	Code                 E_Code   `protobuf:"varint,1,opt,name=code,enum=netproto.E_Code" json:"code,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SC_ChapterChallenge) Reset()         { *m = SC_ChapterChallenge{} }
func (m *SC_ChapterChallenge) String() string { return proto.CompactTextString(m) }
func (*SC_ChapterChallenge) ProtoMessage()    {}
func (*SC_ChapterChallenge) Descriptor() ([]byte, []int) {
	return fileDescriptor_logic_444f2d1a507ba940, []int{17}
}
func (m *SC_ChapterChallenge) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SC_ChapterChallenge.Unmarshal(m, b)
}
func (m *SC_ChapterChallenge) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SC_ChapterChallenge.Marshal(b, m, deterministic)
}
func (dst *SC_ChapterChallenge) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SC_ChapterChallenge.Merge(dst, src)
}
func (m *SC_ChapterChallenge) XXX_Size() int {
	return xxx_messageInfo_SC_ChapterChallenge.Size(m)
}
func (m *SC_ChapterChallenge) XXX_DiscardUnknown() {
	xxx_messageInfo_SC_ChapterChallenge.DiscardUnknown(m)
}

var xxx_messageInfo_SC_ChapterChallenge proto.InternalMessageInfo

func (m *SC_ChapterChallenge) GetCode() E_Code {
	if m != nil {
		return m.Code
	}
	return E_Code_E_ERROR
}

type CS_None struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CS_None) Reset()         { *m = CS_None{} }
func (m *CS_None) String() string { return proto.CompactTextString(m) }
func (*CS_None) ProtoMessage()    {}
func (*CS_None) Descriptor() ([]byte, []int) {
	return fileDescriptor_logic_444f2d1a507ba940, []int{18}
}
func (m *CS_None) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CS_None.Unmarshal(m, b)
}
func (m *CS_None) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CS_None.Marshal(b, m, deterministic)
}
func (dst *CS_None) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CS_None.Merge(dst, src)
}
func (m *CS_None) XXX_Size() int {
	return xxx_messageInfo_CS_None.Size(m)
}
func (m *CS_None) XXX_DiscardUnknown() {
	xxx_messageInfo_CS_None.DiscardUnknown(m)
}

var xxx_messageInfo_CS_None proto.InternalMessageInfo

func init() {
	proto.RegisterType((*CS_Tick)(nil), "netproto.CS_Tick")
	proto.RegisterType((*SC_Tick)(nil), "netproto.SC_Tick")
	proto.RegisterType((*CS_Login)(nil), "netproto.CS_Login")
	proto.RegisterType((*SC_Login)(nil), "netproto.SC_Login")
	proto.RegisterType((*CS_ShopAddDiamond)(nil), "netproto.CS_ShopAddDiamond")
	proto.RegisterType((*SC_ShopAddDiamond)(nil), "netproto.SC_ShopAddDiamond")
	proto.RegisterType((*CS_ShopBuyGold)(nil), "netproto.CS_ShopBuyGold")
	proto.RegisterType((*SC_ShopBuyGold)(nil), "netproto.SC_ShopBuyGold")
	proto.RegisterType((*CS_HeroDraw)(nil), "netproto.CS_HeroDraw")
	proto.RegisterType((*SC_HeroDraw)(nil), "netproto.SC_HeroDraw")
	proto.RegisterType((*CS_HeroUplevel)(nil), "netproto.CS_HeroUplevel")
	proto.RegisterType((*SC_HeroUpLevel)(nil), "netproto.SC_HeroUpLevel")
	proto.RegisterType((*CS_HeroEquip)(nil), "netproto.CS_HeroEquip")
	proto.RegisterType((*SC_HeroEquip)(nil), "netproto.SC_HeroEquip")
	proto.RegisterType((*CS_ChapterGetPrize)(nil), "netproto.CS_ChapterGetPrize")
	proto.RegisterType((*SC_ChapterGetPrize)(nil), "netproto.SC_ChapterGetPrize")
	proto.RegisterType((*CS_ChapterChallenge)(nil), "netproto.CS_ChapterChallenge")
	proto.RegisterType((*SC_ChapterChallenge)(nil), "netproto.SC_ChapterChallenge")
	proto.RegisterType((*CS_None)(nil), "netproto.CS_None")
	proto.RegisterEnum("netproto.E_Code", E_Code_name, E_Code_value)
	proto.RegisterEnum("netproto.EMsgIds", EMsgIds_name, EMsgIds_value)
}

func init() { proto.RegisterFile("logic.proto", fileDescriptor_logic_444f2d1a507ba940) }

var fileDescriptor_logic_444f2d1a507ba940 = []byte{
	// 643 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x54, 0xcd, 0x4f, 0xdb, 0x4e,
	0x10, 0xe5, 0x33, 0x09, 0x93, 0xc0, 0x6f, 0xd8, 0x04, 0xc8, 0x4f, 0x3d, 0x14, 0x59, 0xad, 0x84,
	0x38, 0xa0, 0xaa, 0x54, 0x95, 0xda, 0x5e, 0xea, 0x6e, 0x56, 0xc1, 0xc2, 0xb1, 0x23, 0xaf, 0xa1,
	0xed, 0x69, 0x95, 0xe2, 0x6d, 0xe2, 0xe2, 0x78, 0x83, 0x93, 0xd0, 0x8f, 0x6b, 0xd5, 0xff, 0xbb,
	0xca, 0xda, 0x26, 0x81, 0x50, 0xc9, 0xb7, 0x9d, 0x37, 0x6f, 0xdf, 0xec, 0x8e, 0xe6, 0x0d, 0x54,
	0x23, 0xd5, 0x0f, 0xaf, 0x4e, 0x46, 0x89, 0x9a, 0x28, 0x52, 0x89, 0xe5, 0x44, 0x9f, 0x8c, 0x2d,
	0x28, 0x53, 0x2e, 0xfc, 0xf0, 0xea, 0x7a, 0x76, 0xe4, 0x34, 0x3d, 0x1e, 0x42, 0x85, 0x72, 0x61,
	0xab, 0x7e, 0x18, 0x93, 0x06, 0x6c, 0x4e, 0xd4, 0xb5, 0x8c, 0x9b, 0xab, 0x87, 0xab, 0x47, 0x5b,
	0x5e, 0x1a, 0x18, 0x2f, 0xa0, 0xc2, 0x69, 0xc6, 0x78, 0x06, 0x1b, 0x57, 0x2a, 0x90, 0x9a, 0xb0,
	0xf3, 0x12, 0x4f, 0x72, 0xf1, 0x13, 0x26, 0xa8, 0x0a, 0xa4, 0xa7, 0xb3, 0xc6, 0x73, 0xd8, 0xa5,
	0x5c, 0xf0, 0x81, 0x1a, 0x99, 0x41, 0xd0, 0x0a, 0x7b, 0x43, 0x15, 0x07, 0x04, 0x61, 0x3d, 0x0c,
	0x7e, 0xe8, 0x9b, 0x9b, 0xde, 0xec, 0x68, 0xbc, 0x81, 0x5d, 0x4e, 0x1f, 0xd2, 0x8a, 0x55, 0x40,
	0xd8, 0xc9, 0x2a, 0x7c, 0x98, 0xfe, 0x6c, 0xab, 0x28, 0x30, 0x5e, 0xc3, 0x4e, 0x26, 0x96, 0x21,
	0x05, 0x95, 0x9e, 0x42, 0x95, 0x72, 0x71, 0x26, 0x13, 0xd5, 0x4a, 0x7a, 0xdf, 0x67, 0xaf, 0x8c,
	0xa7, 0xc3, 0xfc, 0x95, 0xf1, 0x74, 0x68, 0x9c, 0x42, 0x95, 0xd3, 0x39, 0xa1, 0x98, 0xea, 0x91,
	0x7e, 0xdf, 0xec, 0xd2, 0xc5, 0x28, 0x92, 0xb7, 0x32, 0x22, 0xfb, 0x50, 0x1a, 0xc8, 0x44, 0x59,
	0x41, 0xa6, 0x9d, 0x45, 0xd9, 0xbb, 0x53, 0xa6, 0xad, 0x99, 0xc5, 0x2a, 0xbc, 0x87, 0x5a, 0x56,
	0x81, 0xdd, 0x4c, 0xc3, 0xd1, 0xbf, 0xf4, 0x49, 0x13, 0xca, 0x72, 0x46, 0xb0, 0x82, 0xe6, 0x9a,
	0x4e, 0xe4, 0xa1, 0xf1, 0x0a, 0x6a, 0x59, 0xe5, 0x54, 0xa1, 0x58, 0xdd, 0x06, 0x10, 0xca, 0x05,
	0x1d, 0xf4, 0x46, 0x13, 0x99, 0xb4, 0xe5, 0xa4, 0x9b, 0x84, 0xbf, 0xa4, 0xf1, 0x16, 0x08, 0xa7,
	0x0f, 0xd1, 0x82, 0x8a, 0x7b, 0x50, 0x9f, 0x2b, 0xd2, 0x41, 0x2f, 0x8a, 0x64, 0xdc, 0x97, 0xc6,
	0x3b, 0xa8, 0xcf, 0x25, 0xef, 0xe0, 0x82, 0x9a, 0xe9, 0xac, 0x3b, 0x2a, 0x96, 0xc7, 0xbf, 0xd7,
	0xa0, 0x94, 0xe6, 0x48, 0x15, 0xca, 0x4c, 0x30, 0xcf, 0x73, 0x3d, 0x5c, 0x21, 0x15, 0xd8, 0x60,
	0xc2, 0x3d, 0xc7, 0x55, 0xb2, 0x0d, 0x5b, 0x4c, 0x78, 0xcc, 0x76, 0xdb, 0x96, 0x83, 0x6b, 0x04,
	0xa1, 0xc6, 0x84, 0xe3, 0x3a, 0x4c, 0xb0, 0x4f, 0x16, 0xf7, 0x71, 0x3d, 0x25, 0x5c, 0x38, 0xe7,
	0x8e, 0xfb, 0xd1, 0xc1, 0x0d, 0xf2, 0x04, 0x0e, 0x98, 0xe0, 0xcc, 0xbb, 0x64, 0x9e, 0xb0, 0x1c,
	0x9f, 0x79, 0x8e, 0x69, 0x67, 0xb2, 0x9b, 0xa4, 0x0e, 0xff, 0x31, 0x61, 0x39, 0x97, 0xa6, 0x6d,
	0xb5, 0x44, 0xd7, 0xf4, 0xcc, 0x0e, 0x96, 0xc8, 0x2e, 0x6c, 0xcf, 0x41, 0xb7, 0xeb, 0x63, 0x39,
	0xad, 0x62, 0x5b, 0x1d, 0xcb, 0x17, 0x6d, 0xd7, 0x6e, 0x61, 0x23, 0xbd, 0x99, 0x22, 0x2d, 0xcb,
	0xec, 0xb8, 0x4e, 0x0b, 0xf7, 0xd2, 0x9b, 0x29, 0xe8, 0x5b, 0x1d, 0xc6, 0x71, 0x9f, 0xfc, 0x0f,
	0x7b, 0x39, 0x44, 0xcf, 0xcc, 0xae, 0xcf, 0x3c, 0x61, 0xb3, 0x4b, 0x66, 0xe3, 0x01, 0x69, 0x42,
	0x23, 0x4f, 0x75, 0x6d, 0xf3, 0xf3, 0x5d, 0xa6, 0x79, 0xfc, 0x67, 0x1d, 0xca, 0xac, 0x33, 0xee,
	0x5b, 0xc1, 0x98, 0xd4, 0xa0, 0xc2, 0xb2, 0xee, 0xe0, 0x4a, 0x1e, 0xcd, 0x96, 0x01, 0x4a, 0x1d,
	0x65, 0xab, 0x01, 0xbf, 0xea, 0x8f, 0xe7, 0xdb, 0x01, 0xfb, 0x3a, 0xcc, 0x57, 0x01, 0x0e, 0xc8,
	0x3e, 0x10, 0xb6, 0xe4, 0x73, 0x0c, 0x35, 0xbe, 0x64, 0x6c, 0xfc, 0xa6, 0x3f, 0x78, 0xdf, 0xb5,
	0x78, 0xad, 0xc1, 0xfb, 0xc6, 0xc5, 0x48, 0x37, 0x67, 0xc1, 0x95, 0x38, 0xd4, 0xc8, 0x82, 0x0d,
	0x31, 0xce, 0xd5, 0x16, 0x3c, 0x86, 0x2a, 0x57, 0x5b, 0xb0, 0x13, 0x8e, 0x74, 0x0f, 0x17, 0xbd,
	0x82, 0x37, 0x1a, 0x5a, 0x1c, 0x7e, 0x4c, 0xc8, 0x01, 0xd4, 0xd9, 0xf2, 0x64, 0xe3, 0x58, 0x27,
	0x96, 0x87, 0x1b, 0x27, 0xba, 0xdb, 0x8f, 0x4c, 0x2e, 0x4e, 0x75, 0xe6, 0x91, 0xe1, 0xc5, 0xdb,
	0x2f, 0x25, 0x3d, 0xac, 0xa7, 0x7f, 0x03, 0x00, 0x00, 0xff, 0xff, 0xb6, 0x74, 0xbb, 0x0f, 0xa4,
	0x05, 0x00, 0x00,
}
