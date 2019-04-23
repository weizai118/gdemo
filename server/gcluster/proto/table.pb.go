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
	return fileDescriptor_table_4fbf04d5de59202e, []int{0}
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
	return fileDescriptor_table_4fbf04d5de59202e, []int{1}
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
	return fileDescriptor_table_4fbf04d5de59202e, []int{2}
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
	return fileDescriptor_table_4fbf04d5de59202e, []int{3}
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
	return fileDescriptor_table_4fbf04d5de59202e, []int{4}
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
	return fileDescriptor_table_4fbf04d5de59202e, []int{5}
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
	proto.RegisterType((*Equip)(nil), "Equip")
	proto.RegisterType((*EquipArray)(nil), "EquipArray")
	proto.RegisterType((*Hero)(nil), "Hero")
	proto.RegisterType((*HeroArray)(nil), "HeroArray")
}

func init() { proto.RegisterFile("table.proto", fileDescriptor_table_4fbf04d5de59202e) }

var fileDescriptor_table_4fbf04d5de59202e = []byte{
	// 257 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x92, 0x41, 0x4b, 0xc3, 0x40,
	0x10, 0x85, 0xc9, 0x26, 0x1b, 0xed, 0x54, 0x44, 0x16, 0x0f, 0x0b, 0x8a, 0x84, 0x9c, 0x72, 0xca,
	0x41, 0xaf, 0x22, 0x44, 0x11, 0x13, 0x04, 0x91, 0x5c, 0x3c, 0xa7, 0x64, 0x41, 0x69, 0xcb, 0x8e,
	0xdb, 0x8d, 0xa5, 0xff, 0x5e, 0x66, 0x76, 0x3d, 0xf4, 0x50, 0xf0, 0xd4, 0xdb, 0x9b, 0xc7, 0x7b,
	0x33, 0xdf, 0xc2, 0xc2, 0xdc, 0x0f, 0x8b, 0x95, 0xa9, 0xd1, 0x59, 0x6f, 0xcb, 0x0f, 0x38, 0x79,
	0xfa, 0x1c, 0xd0, 0x1b, 0xa7, 0xce, 0x41, 0x74, 0xa3, 0x4e, 0x8a, 0xa4, 0x92, 0xbd, 0xe8, 0x46,
	0xa5, 0x20, 0x7b, 0x1b, 0xd6, 0x46, 0x8b, 0x22, 0xa9, 0x66, 0x3d, 0x6b, 0xf2, 0x5e, 0xec, 0x6a,
	0xd4, 0x29, 0xa7, 0x58, 0xab, 0x4b, 0x90, 0xef, 0x76, 0x6b, 0x9c, 0xce, 0xd8, 0x0c, 0x43, 0xf9,
	0x08, 0x67, 0x71, 0x71, 0xe3, 0xdc, 0xb0, 0xa3, 0xe6, 0xab, 0xd9, 0x6d, 0x74, 0x52, 0xa4, 0xd4,
	0x24, 0xad, 0x6e, 0x40, 0x76, 0xde, 0xac, 0x37, 0x5a, 0x14, 0x69, 0x35, 0xbf, 0x3d, 0xad, 0x63,
	0xa3, 0x0f, 0x76, 0xb9, 0x05, 0xf9, 0xfc, 0x3d, 0x7d, 0xe1, 0xbf, 0xd0, 0x2e, 0x20, 0x6d, 0xfc,
	0x32, 0x92, 0x91, 0xa4, 0x56, 0x8b, 0x91, 0x4a, 0xb4, 0x48, 0xa0, 0x8d, 0x5f, 0x4e, 0xa8, 0x65,
	0x00, 0xe5, 0x81, 0x76, 0xb5, 0x38, 0xa1, 0xce, 0xc3, 0x93, 0x48, 0x97, 0x0f, 0x00, 0x7c, 0xf8,
	0x30, 0xfa, 0xf5, 0x3e, 0x7a, 0x5e, 0x73, 0xfe, 0x0f, 0xfc, 0x07, 0xb2, 0xd6, 0x38, 0x7b, 0x74,
	0xee, 0x7b, 0x98, 0xd1, 0xdd, 0xc3, 0xd8, 0x57, 0xfb, 0xd8, 0xb2, 0xa6, 0x78, 0xa4, 0x5e, 0xe4,
	0xfc, 0x25, 0xee, 0x7e, 0x03, 0x00, 0x00, 0xff, 0xff, 0x81, 0x4f, 0x5f, 0x92, 0x21, 0x02, 0x00,
	0x00,
}
