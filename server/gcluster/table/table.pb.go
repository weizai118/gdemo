// Code generated by protoc-gen-go. DO NOT EDIT.
// source: table.proto

package table

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

type Chapter struct {
	Id                   int32    `protobuf:"varint,1,opt,name=Id" json:"Id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=Name" json:"Name,omitempty"`
	Gold                 int32    `protobuf:"varint,3,opt,name=Gold" json:"Gold,omitempty"`
	Power                int32    `protobuf:"varint,4,opt,name=Power" json:"Power,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Chapter) Reset()         { *m = Chapter{} }
func (m *Chapter) String() string { return proto.CompactTextString(m) }
func (*Chapter) ProtoMessage()    {}
func (*Chapter) Descriptor() ([]byte, []int) {
	return fileDescriptor_table_b59d63efbe51f792, []int{0}
}
func (m *Chapter) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Chapter.Unmarshal(m, b)
}
func (m *Chapter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Chapter.Marshal(b, m, deterministic)
}
func (dst *Chapter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Chapter.Merge(dst, src)
}
func (m *Chapter) XXX_Size() int {
	return xxx_messageInfo_Chapter.Size(m)
}
func (m *Chapter) XXX_DiscardUnknown() {
	xxx_messageInfo_Chapter.DiscardUnknown(m)
}

var xxx_messageInfo_Chapter proto.InternalMessageInfo

func (m *Chapter) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Chapter) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Chapter) GetGold() int32 {
	if m != nil {
		return m.Gold
	}
	return 0
}

func (m *Chapter) GetPower() int32 {
	if m != nil {
		return m.Power
	}
	return 0
}

type ChapterArray struct {
	Keys                 []int32    `protobuf:"varint,1,rep,packed,name=Keys" json:"Keys,omitempty"`
	Items                []*Chapter `protobuf:"bytes,2,rep,name=Items" json:"Items,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *ChapterArray) Reset()         { *m = ChapterArray{} }
func (m *ChapterArray) String() string { return proto.CompactTextString(m) }
func (*ChapterArray) ProtoMessage()    {}
func (*ChapterArray) Descriptor() ([]byte, []int) {
	return fileDescriptor_table_b59d63efbe51f792, []int{1}
}
func (m *ChapterArray) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ChapterArray.Unmarshal(m, b)
}
func (m *ChapterArray) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ChapterArray.Marshal(b, m, deterministic)
}
func (dst *ChapterArray) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChapterArray.Merge(dst, src)
}
func (m *ChapterArray) XXX_Size() int {
	return xxx_messageInfo_ChapterArray.Size(m)
}
func (m *ChapterArray) XXX_DiscardUnknown() {
	xxx_messageInfo_ChapterArray.DiscardUnknown(m)
}

var xxx_messageInfo_ChapterArray proto.InternalMessageInfo

func (m *ChapterArray) GetKeys() []int32 {
	if m != nil {
		return m.Keys
	}
	return nil
}

func (m *ChapterArray) GetItems() []*Chapter {
	if m != nil {
		return m.Items
	}
	return nil
}

type Common struct {
	Id                   int32    `protobuf:"varint,1,opt,name=Id" json:"Id,omitempty"`
	Val                  int32    `protobuf:"varint,2,opt,name=Val" json:"Val,omitempty"`
	Des                  string   `protobuf:"bytes,3,opt,name=Des" json:"Des,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Common) Reset()         { *m = Common{} }
func (m *Common) String() string { return proto.CompactTextString(m) }
func (*Common) ProtoMessage()    {}
func (*Common) Descriptor() ([]byte, []int) {
	return fileDescriptor_table_b59d63efbe51f792, []int{2}
}
func (m *Common) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Common.Unmarshal(m, b)
}
func (m *Common) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Common.Marshal(b, m, deterministic)
}
func (dst *Common) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Common.Merge(dst, src)
}
func (m *Common) XXX_Size() int {
	return xxx_messageInfo_Common.Size(m)
}
func (m *Common) XXX_DiscardUnknown() {
	xxx_messageInfo_Common.DiscardUnknown(m)
}

var xxx_messageInfo_Common proto.InternalMessageInfo

func (m *Common) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Common) GetVal() int32 {
	if m != nil {
		return m.Val
	}
	return 0
}

func (m *Common) GetDes() string {
	if m != nil {
		return m.Des
	}
	return ""
}

type CommonArray struct {
	Keys                 []int32   `protobuf:"varint,1,rep,packed,name=Keys" json:"Keys,omitempty"`
	Items                []*Common `protobuf:"bytes,2,rep,name=Items" json:"Items,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *CommonArray) Reset()         { *m = CommonArray{} }
func (m *CommonArray) String() string { return proto.CompactTextString(m) }
func (*CommonArray) ProtoMessage()    {}
func (*CommonArray) Descriptor() ([]byte, []int) {
	return fileDescriptor_table_b59d63efbe51f792, []int{3}
}
func (m *CommonArray) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CommonArray.Unmarshal(m, b)
}
func (m *CommonArray) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CommonArray.Marshal(b, m, deterministic)
}
func (dst *CommonArray) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CommonArray.Merge(dst, src)
}
func (m *CommonArray) XXX_Size() int {
	return xxx_messageInfo_CommonArray.Size(m)
}
func (m *CommonArray) XXX_DiscardUnknown() {
	xxx_messageInfo_CommonArray.DiscardUnknown(m)
}

var xxx_messageInfo_CommonArray proto.InternalMessageInfo

func (m *CommonArray) GetKeys() []int32 {
	if m != nil {
		return m.Keys
	}
	return nil
}

func (m *CommonArray) GetItems() []*Common {
	if m != nil {
		return m.Items
	}
	return nil
}

type Equip struct {
	Id                   int32    `protobuf:"varint,1,opt,name=Id" json:"Id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=Name" json:"Name,omitempty"`
	Atk                  int32    `protobuf:"varint,3,opt,name=Atk" json:"Atk,omitempty"`
	Hp                   int32    `protobuf:"varint,4,opt,name=Hp" json:"Hp,omitempty"`
	Atkup                int32    `protobuf:"varint,5,opt,name=Atkup" json:"Atkup,omitempty"`
	Hpup                 int32    `protobuf:"varint,6,opt,name=Hpup" json:"Hpup,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Equip) Reset()         { *m = Equip{} }
func (m *Equip) String() string { return proto.CompactTextString(m) }
func (*Equip) ProtoMessage()    {}
func (*Equip) Descriptor() ([]byte, []int) {
	return fileDescriptor_table_b59d63efbe51f792, []int{4}
}
func (m *Equip) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Equip.Unmarshal(m, b)
}
func (m *Equip) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Equip.Marshal(b, m, deterministic)
}
func (dst *Equip) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Equip.Merge(dst, src)
}
func (m *Equip) XXX_Size() int {
	return xxx_messageInfo_Equip.Size(m)
}
func (m *Equip) XXX_DiscardUnknown() {
	xxx_messageInfo_Equip.DiscardUnknown(m)
}

var xxx_messageInfo_Equip proto.InternalMessageInfo

func (m *Equip) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Equip) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Equip) GetAtk() int32 {
	if m != nil {
		return m.Atk
	}
	return 0
}

func (m *Equip) GetHp() int32 {
	if m != nil {
		return m.Hp
	}
	return 0
}

func (m *Equip) GetAtkup() int32 {
	if m != nil {
		return m.Atkup
	}
	return 0
}

func (m *Equip) GetHpup() int32 {
	if m != nil {
		return m.Hpup
	}
	return 0
}

type EquipArray struct {
	Keys                 []int32  `protobuf:"varint,1,rep,packed,name=Keys" json:"Keys,omitempty"`
	Items                []*Equip `protobuf:"bytes,2,rep,name=Items" json:"Items,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EquipArray) Reset()         { *m = EquipArray{} }
func (m *EquipArray) String() string { return proto.CompactTextString(m) }
func (*EquipArray) ProtoMessage()    {}
func (*EquipArray) Descriptor() ([]byte, []int) {
	return fileDescriptor_table_b59d63efbe51f792, []int{5}
}
func (m *EquipArray) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EquipArray.Unmarshal(m, b)
}
func (m *EquipArray) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EquipArray.Marshal(b, m, deterministic)
}
func (dst *EquipArray) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EquipArray.Merge(dst, src)
}
func (m *EquipArray) XXX_Size() int {
	return xxx_messageInfo_EquipArray.Size(m)
}
func (m *EquipArray) XXX_DiscardUnknown() {
	xxx_messageInfo_EquipArray.DiscardUnknown(m)
}

var xxx_messageInfo_EquipArray proto.InternalMessageInfo

func (m *EquipArray) GetKeys() []int32 {
	if m != nil {
		return m.Keys
	}
	return nil
}

func (m *EquipArray) GetItems() []*Equip {
	if m != nil {
		return m.Items
	}
	return nil
}

type Hero struct {
	Id                   int32    `protobuf:"varint,1,opt,name=Id" json:"Id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=Name" json:"Name,omitempty"`
	Atk                  int32    `protobuf:"varint,3,opt,name=Atk" json:"Atk,omitempty"`
	Hp                   int32    `protobuf:"varint,4,opt,name=Hp" json:"Hp,omitempty"`
	Atkup                int32    `protobuf:"varint,5,opt,name=Atkup" json:"Atkup,omitempty"`
	Hpup                 int32    `protobuf:"varint,6,opt,name=Hpup" json:"Hpup,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Hero) Reset()         { *m = Hero{} }
func (m *Hero) String() string { return proto.CompactTextString(m) }
func (*Hero) ProtoMessage()    {}
func (*Hero) Descriptor() ([]byte, []int) {
	return fileDescriptor_table_b59d63efbe51f792, []int{6}
}
func (m *Hero) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Hero.Unmarshal(m, b)
}
func (m *Hero) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Hero.Marshal(b, m, deterministic)
}
func (dst *Hero) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Hero.Merge(dst, src)
}
func (m *Hero) XXX_Size() int {
	return xxx_messageInfo_Hero.Size(m)
}
func (m *Hero) XXX_DiscardUnknown() {
	xxx_messageInfo_Hero.DiscardUnknown(m)
}

var xxx_messageInfo_Hero proto.InternalMessageInfo

func (m *Hero) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Hero) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Hero) GetAtk() int32 {
	if m != nil {
		return m.Atk
	}
	return 0
}

func (m *Hero) GetHp() int32 {
	if m != nil {
		return m.Hp
	}
	return 0
}

func (m *Hero) GetAtkup() int32 {
	if m != nil {
		return m.Atkup
	}
	return 0
}

func (m *Hero) GetHpup() int32 {
	if m != nil {
		return m.Hpup
	}
	return 0
}

type HeroArray struct {
	Keys                 []int32  `protobuf:"varint,1,rep,packed,name=Keys" json:"Keys,omitempty"`
	Items                []*Hero  `protobuf:"bytes,2,rep,name=Items" json:"Items,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HeroArray) Reset()         { *m = HeroArray{} }
func (m *HeroArray) String() string { return proto.CompactTextString(m) }
func (*HeroArray) ProtoMessage()    {}
func (*HeroArray) Descriptor() ([]byte, []int) {
	return fileDescriptor_table_b59d63efbe51f792, []int{7}
}
func (m *HeroArray) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HeroArray.Unmarshal(m, b)
}
func (m *HeroArray) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HeroArray.Marshal(b, m, deterministic)
}
func (dst *HeroArray) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HeroArray.Merge(dst, src)
}
func (m *HeroArray) XXX_Size() int {
	return xxx_messageInfo_HeroArray.Size(m)
}
func (m *HeroArray) XXX_DiscardUnknown() {
	xxx_messageInfo_HeroArray.DiscardUnknown(m)
}

var xxx_messageInfo_HeroArray proto.InternalMessageInfo

func (m *HeroArray) GetKeys() []int32 {
	if m != nil {
		return m.Keys
	}
	return nil
}

func (m *HeroArray) GetItems() []*Hero {
	if m != nil {
		return m.Items
	}
	return nil
}

func init() {
	proto.RegisterType((*Chapter)(nil), "Chapter")
	proto.RegisterType((*ChapterArray)(nil), "ChapterArray")
	proto.RegisterType((*Common)(nil), "Common")
	proto.RegisterType((*CommonArray)(nil), "CommonArray")
	proto.RegisterType((*Equip)(nil), "Equip")
	proto.RegisterType((*EquipArray)(nil), "EquipArray")
	proto.RegisterType((*Hero)(nil), "Hero")
	proto.RegisterType((*HeroArray)(nil), "HeroArray")
}

func init() { proto.RegisterFile("table.proto", fileDescriptor_table_b59d63efbe51f792) }

var fileDescriptor_table_b59d63efbe51f792 = []byte{
	// 300 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x92, 0xcd, 0x4b, 0xc4, 0x30,
	0x10, 0xc5, 0xe9, 0x47, 0xba, 0x76, 0x2a, 0xb2, 0x14, 0x0f, 0x01, 0x3f, 0x28, 0x3d, 0xf5, 0xd4,
	0x83, 0x5e, 0x17, 0xb1, 0xae, 0x62, 0x8b, 0x20, 0xd2, 0x83, 0x9e, 0xbb, 0x34, 0xa0, 0x6c, 0x6b,
	0x62, 0x9a, 0xba, 0xec, 0x7f, 0x2f, 0x93, 0x64, 0x0f, 0x45, 0x0a, 0x9e, 0xbc, 0xbd, 0x3c, 0xe6,
	0xcd, 0xfc, 0x78, 0x04, 0x22, 0xd5, 0x6c, 0x3a, 0x96, 0x0b, 0xc9, 0x15, 0x4f, 0xdf, 0x60, 0xb1,
	0x7e, 0x6f, 0x84, 0x62, 0x32, 0x3e, 0x01, 0xb7, 0x6a, 0xa9, 0x93, 0x38, 0x19, 0xa9, 0xdd, 0xaa,
	0x8d, 0x63, 0xf0, 0x9f, 0x9b, 0x9e, 0x51, 0x37, 0x71, 0xb2, 0xb0, 0xd6, 0x1a, 0xbd, 0x47, 0xde,
	0xb5, 0xd4, 0xd3, 0x53, 0x5a, 0xc7, 0xa7, 0x40, 0x5e, 0xf8, 0x8e, 0x49, 0xea, 0x6b, 0xd3, 0x3c,
	0xd2, 0x3b, 0x38, 0xb6, 0x8b, 0x0b, 0x29, 0x9b, 0x3d, 0x26, 0x9f, 0xd8, 0x7e, 0xa0, 0x4e, 0xe2,
	0x61, 0x12, 0x75, 0x7c, 0x09, 0xa4, 0x52, 0xac, 0x1f, 0xa8, 0x9b, 0x78, 0x59, 0x74, 0x75, 0x94,
	0xdb, 0x44, 0x6d, 0xec, 0x74, 0x05, 0xc1, 0x9a, 0xf7, 0x3d, 0xff, 0xfc, 0xc5, 0xb6, 0x04, 0xef,
	0xb5, 0xe9, 0x34, 0x1a, 0xa9, 0x51, 0xa2, 0x73, 0xcf, 0x06, 0x0d, 0x16, 0xd6, 0x28, 0xd3, 0x5b,
	0x88, 0x4c, 0x7a, 0x1e, 0xe0, 0x62, 0x0a, 0xb0, 0xc8, 0x4d, 0xe0, 0x70, 0x7f, 0x07, 0xe4, 0xe1,
	0x6b, 0xfc, 0x10, 0x7f, 0xaa, 0x66, 0x09, 0x5e, 0xa1, 0xb6, 0xb6, 0x19, 0x94, 0x98, 0x2a, 0x85,
	0x6d, 0xc5, 0x2d, 0x05, 0x16, 0x55, 0xa8, 0xed, 0x28, 0x28, 0x31, 0x45, 0xe9, 0x07, 0xee, 0x2a,
	0xc5, 0x28, 0x68, 0x60, 0x2a, 0x45, 0x9d, 0xde, 0x00, 0xe8, 0xc3, 0xf3, 0xe4, 0xe7, 0x53, 0xf2,
	0x20, 0xd7, 0xf3, 0x07, 0xf0, 0x6f, 0xf0, 0x4b, 0x26, 0xf9, 0xbf, 0x73, 0xaf, 0x20, 0xc4, 0xbb,
	0xf3, 0xd8, 0x67, 0x53, 0x6c, 0x92, 0xe3, 0xb8, 0xa5, 0xde, 0x04, 0xfa, 0x4b, 0x5e, 0xff, 0x04,
	0x00, 0x00, 0xff, 0xff, 0x23, 0xe4, 0xe3, 0x83, 0xa1, 0x02, 0x00, 0x00,
}
