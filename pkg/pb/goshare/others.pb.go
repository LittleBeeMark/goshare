// Code generated by protoc-gen-go. DO NOT EDIT.
// source: goshare/others.proto

package goshare

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

type NetInAmountDetail struct {
	Amount               float64  `protobuf:"fixed64,1,opt,name=amount,proto3" json:"amount"`
	Percentage           float64  `protobuf:"fixed64,2,opt,name=percentage,proto3" json:"percentage"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NetInAmountDetail) Reset()         { *m = NetInAmountDetail{} }
func (m *NetInAmountDetail) String() string { return proto.CompactTextString(m) }
func (*NetInAmountDetail) ProtoMessage()    {}
func (*NetInAmountDetail) Descriptor() ([]byte, []int) {
	return fileDescriptor_0408abfcc2f0ac69, []int{0}
}

func (m *NetInAmountDetail) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NetInAmountDetail.Unmarshal(m, b)
}
func (m *NetInAmountDetail) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NetInAmountDetail.Marshal(b, m, deterministic)
}
func (m *NetInAmountDetail) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NetInAmountDetail.Merge(m, src)
}
func (m *NetInAmountDetail) XXX_Size() int {
	return xxx_messageInfo_NetInAmountDetail.Size(m)
}
func (m *NetInAmountDetail) XXX_DiscardUnknown() {
	xxx_messageInfo_NetInAmountDetail.DiscardUnknown(m)
}

var xxx_messageInfo_NetInAmountDetail proto.InternalMessageInfo

func (m *NetInAmountDetail) GetAmount() float64 {
	if m != nil {
		return m.Amount
	}
	return 0
}

func (m *NetInAmountDetail) GetPercentage() float64 {
	if m != nil {
		return m.Percentage
	}
	return 0
}

type RealtimeMoneyTrendItem struct {
	Symbol               string             `protobuf:"bytes,1,opt,name=symbol,proto3" json:"symbol"`
	Name                 string             `protobuf:"bytes,2,opt,name=name,proto3" json:"name"`
	Price                float64            `protobuf:"fixed64,3,opt,name=price,proto3" json:"price"`
	UpdownPercentage     float64            `protobuf:"fixed64,4,opt,name=updown_percentage,json=updownPercentage,proto3" json:"updownPercentage"`
	Time                 int64              `protobuf:"varint,6,opt,name=time,proto3" json:"time"`
	SuperSuperBigOrder   *NetInAmountDetail `protobuf:"bytes,7,opt,name=super_super_big_order,json=superSuperBigOrder,proto3" json:"superSuperBigOrder"`
	SuperBigOrder        *NetInAmountDetail `protobuf:"bytes,8,opt,name=super_big_order,json=superBigOrder,proto3" json:"superBigOrder"`
	BigOrder             *NetInAmountDetail `protobuf:"bytes,9,opt,name=big_order,json=bigOrder,proto3" json:"bigOrder"`
	MiddleOrder          *NetInAmountDetail `protobuf:"bytes,10,opt,name=middle_order,json=middleOrder,proto3" json:"middleOrder"`
	SmallOrder           *NetInAmountDetail `protobuf:"bytes,11,opt,name=small_order,json=smallOrder,proto3" json:"smallOrder"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *RealtimeMoneyTrendItem) Reset()         { *m = RealtimeMoneyTrendItem{} }
func (m *RealtimeMoneyTrendItem) String() string { return proto.CompactTextString(m) }
func (*RealtimeMoneyTrendItem) ProtoMessage()    {}
func (*RealtimeMoneyTrendItem) Descriptor() ([]byte, []int) {
	return fileDescriptor_0408abfcc2f0ac69, []int{1}
}

func (m *RealtimeMoneyTrendItem) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RealtimeMoneyTrendItem.Unmarshal(m, b)
}
func (m *RealtimeMoneyTrendItem) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RealtimeMoneyTrendItem.Marshal(b, m, deterministic)
}
func (m *RealtimeMoneyTrendItem) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RealtimeMoneyTrendItem.Merge(m, src)
}
func (m *RealtimeMoneyTrendItem) XXX_Size() int {
	return xxx_messageInfo_RealtimeMoneyTrendItem.Size(m)
}
func (m *RealtimeMoneyTrendItem) XXX_DiscardUnknown() {
	xxx_messageInfo_RealtimeMoneyTrendItem.DiscardUnknown(m)
}

var xxx_messageInfo_RealtimeMoneyTrendItem proto.InternalMessageInfo

func (m *RealtimeMoneyTrendItem) GetSymbol() string {
	if m != nil {
		return m.Symbol
	}
	return ""
}

func (m *RealtimeMoneyTrendItem) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *RealtimeMoneyTrendItem) GetPrice() float64 {
	if m != nil {
		return m.Price
	}
	return 0
}

func (m *RealtimeMoneyTrendItem) GetUpdownPercentage() float64 {
	if m != nil {
		return m.UpdownPercentage
	}
	return 0
}

func (m *RealtimeMoneyTrendItem) GetTime() int64 {
	if m != nil {
		return m.Time
	}
	return 0
}

func (m *RealtimeMoneyTrendItem) GetSuperSuperBigOrder() *NetInAmountDetail {
	if m != nil {
		return m.SuperSuperBigOrder
	}
	return nil
}

func (m *RealtimeMoneyTrendItem) GetSuperBigOrder() *NetInAmountDetail {
	if m != nil {
		return m.SuperBigOrder
	}
	return nil
}

func (m *RealtimeMoneyTrendItem) GetBigOrder() *NetInAmountDetail {
	if m != nil {
		return m.BigOrder
	}
	return nil
}

func (m *RealtimeMoneyTrendItem) GetMiddleOrder() *NetInAmountDetail {
	if m != nil {
		return m.MiddleOrder
	}
	return nil
}

func (m *RealtimeMoneyTrendItem) GetSmallOrder() *NetInAmountDetail {
	if m != nil {
		return m.SmallOrder
	}
	return nil
}

type RealtimeMoneyTrendItemList struct {
	List                 []*RealtimeMoneyTrendItem `protobuf:"bytes,1,rep,name=list,proto3" json:"list"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *RealtimeMoneyTrendItemList) Reset()         { *m = RealtimeMoneyTrendItemList{} }
func (m *RealtimeMoneyTrendItemList) String() string { return proto.CompactTextString(m) }
func (*RealtimeMoneyTrendItemList) ProtoMessage()    {}
func (*RealtimeMoneyTrendItemList) Descriptor() ([]byte, []int) {
	return fileDescriptor_0408abfcc2f0ac69, []int{2}
}

func (m *RealtimeMoneyTrendItemList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RealtimeMoneyTrendItemList.Unmarshal(m, b)
}
func (m *RealtimeMoneyTrendItemList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RealtimeMoneyTrendItemList.Marshal(b, m, deterministic)
}
func (m *RealtimeMoneyTrendItemList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RealtimeMoneyTrendItemList.Merge(m, src)
}
func (m *RealtimeMoneyTrendItemList) XXX_Size() int {
	return xxx_messageInfo_RealtimeMoneyTrendItemList.Size(m)
}
func (m *RealtimeMoneyTrendItemList) XXX_DiscardUnknown() {
	xxx_messageInfo_RealtimeMoneyTrendItemList.DiscardUnknown(m)
}

var xxx_messageInfo_RealtimeMoneyTrendItemList proto.InternalMessageInfo

func (m *RealtimeMoneyTrendItemList) GetList() []*RealtimeMoneyTrendItem {
	if m != nil {
		return m.List
	}
	return nil
}

func init() {
	proto.RegisterType((*NetInAmountDetail)(nil), "goshare.NetInAmountDetail")
	proto.RegisterType((*RealtimeMoneyTrendItem)(nil), "goshare.RealtimeMoneyTrendItem")
	proto.RegisterType((*RealtimeMoneyTrendItemList)(nil), "goshare.RealtimeMoneyTrendItemList")
}

func init() { proto.RegisterFile("goshare/others.proto", fileDescriptor_0408abfcc2f0ac69) }

var fileDescriptor_0408abfcc2f0ac69 = []byte{
	// 372 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x92, 0x4d, 0xab, 0xd3, 0x40,
	0x14, 0x86, 0x89, 0xa9, 0x6d, 0x73, 0xa2, 0x68, 0x87, 0x5a, 0x42, 0x17, 0x5a, 0xba, 0x0a, 0x28,
	0x09, 0xb4, 0x0b, 0x17, 0xe2, 0xc2, 0xe2, 0xa6, 0x68, 0xfd, 0x88, 0xae, 0xdc, 0x94, 0x49, 0x73,
	0x48, 0x07, 0xe7, 0x23, 0xcc, 0x4c, 0x90, 0xfe, 0x5e, 0xff, 0x88, 0x64, 0x92, 0xd6, 0xde, 0x7b,
	0x0b, 0xd9, 0x0c, 0xf3, 0x9e, 0x39, 0xcf, 0x93, 0x61, 0x72, 0x60, 0x5a, 0x2a, 0x73, 0xa4, 0x1a,
	0x53, 0x65, 0x8f, 0xa8, 0x4d, 0x52, 0x69, 0x65, 0x15, 0x19, 0x75, 0xd5, 0xe5, 0x27, 0x98, 0x7c,
	0x41, 0xbb, 0x95, 0x1f, 0x84, 0xaa, 0xa5, 0xfd, 0x88, 0x96, 0x32, 0x4e, 0x66, 0x30, 0xa4, 0x2e,
	0x47, 0xde, 0xc2, 0x8b, 0xbd, 0xac, 0x4b, 0xe4, 0x25, 0x40, 0x85, 0xfa, 0x80, 0xd2, 0xd2, 0x12,
	0xa3, 0x47, 0xee, 0xec, 0xaa, 0xb2, 0xfc, 0xeb, 0xc3, 0x2c, 0x43, 0xca, 0x2d, 0x13, 0xb8, 0x53,
	0x12, 0x4f, 0x3f, 0x35, 0xca, 0x62, 0x6b, 0x51, 0x34, 0x4a, 0x73, 0x12, 0xb9, 0xe2, 0x4e, 0x19,
	0x64, 0x5d, 0x22, 0x04, 0x06, 0x92, 0x8a, 0x56, 0x16, 0x64, 0x6e, 0x4f, 0xa6, 0xf0, 0xb8, 0xd2,
	0xec, 0x80, 0x91, 0xef, 0xbe, 0xd0, 0x06, 0xf2, 0x1a, 0x26, 0x75, 0x55, 0xa8, 0x3f, 0x72, 0x7f,
	0x75, 0x87, 0x81, 0xeb, 0x78, 0xde, 0x1e, 0x7c, 0xbb, 0xd4, 0x1b, 0x6d, 0x73, 0x89, 0x68, 0xb8,
	0xf0, 0x62, 0x3f, 0x73, 0x7b, 0xb2, 0x83, 0x17, 0xa6, 0xae, 0x50, 0xef, 0xdb, 0x35, 0x67, 0xe5,
	0x5e, 0xe9, 0x02, 0x75, 0x34, 0x5a, 0x78, 0x71, 0xb8, 0x9a, 0x27, 0xdd, 0x9b, 0x24, 0x0f, 0x1e,
	0x24, 0x23, 0x0e, 0xf9, 0xd1, 0x2c, 0x1b, 0x56, 0x7e, 0x6d, 0x28, 0xb2, 0x81, 0x67, 0xf7, 0x45,
	0xe3, 0x5e, 0xd1, 0x53, 0x73, 0xc7, 0xf1, 0x16, 0x82, 0xff, 0x74, 0xd0, 0x4b, 0x8f, 0xf3, 0x33,
	0xf8, 0x1e, 0x9e, 0x08, 0x56, 0x14, 0x1c, 0x3b, 0x16, 0x7a, 0xd9, 0xb0, 0xed, 0x6f, 0xf1, 0x77,
	0x10, 0x1a, 0x41, 0x39, 0xef, 0xe8, 0xb0, 0x97, 0x06, 0xd7, 0xee, 0xe0, 0xe5, 0x77, 0x98, 0xdf,
	0xfe, 0xc9, 0x9f, 0x99, 0xb1, 0x64, 0x0d, 0x03, 0xce, 0x4c, 0x33, 0x39, 0x7e, 0x1c, 0xae, 0x5e,
	0x5d, 0x9c, 0xb7, 0x91, 0xcc, 0x35, 0x6f, 0x92, 0x5f, 0x6f, 0x4a, 0x66, 0x8f, 0x75, 0x9e, 0x1c,
	0x94, 0x48, 0x05, 0x93, 0xa8, 0x29, 0xd7, 0x68, 0xd2, 0xf3, 0xf0, 0x56, 0xbf, 0xcb, 0xb4, 0xca,
	0xcf, 0x31, 0x1f, 0xba, 0x29, 0x5e, 0xff, 0x0b, 0x00, 0x00, 0xff, 0xff, 0x10, 0x15, 0xba, 0x4e,
	0xdd, 0x02, 0x00, 0x00,
}