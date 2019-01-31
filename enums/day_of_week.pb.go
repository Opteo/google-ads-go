// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v0/enums/day_of_week.proto

package enums

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

// Enumerates days of the week, e.g., "Monday".
type DayOfWeekEnum_DayOfWeek int32

const (
	// Not specified.
	DayOfWeekEnum_UNSPECIFIED DayOfWeekEnum_DayOfWeek = 0
	// The value is unknown in this version.
	DayOfWeekEnum_UNKNOWN DayOfWeekEnum_DayOfWeek = 1
	// Monday.
	DayOfWeekEnum_MONDAY DayOfWeekEnum_DayOfWeek = 2
	// Tuesday.
	DayOfWeekEnum_TUESDAY DayOfWeekEnum_DayOfWeek = 3
	// Wednesday.
	DayOfWeekEnum_WEDNESDAY DayOfWeekEnum_DayOfWeek = 4
	// Thursday.
	DayOfWeekEnum_THURSDAY DayOfWeekEnum_DayOfWeek = 5
	// Friday.
	DayOfWeekEnum_FRIDAY DayOfWeekEnum_DayOfWeek = 6
	// Saturday.
	DayOfWeekEnum_SATURDAY DayOfWeekEnum_DayOfWeek = 7
	// Sunday.
	DayOfWeekEnum_SUNDAY DayOfWeekEnum_DayOfWeek = 8
)

var DayOfWeekEnum_DayOfWeek_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "UNKNOWN",
	2: "MONDAY",
	3: "TUESDAY",
	4: "WEDNESDAY",
	5: "THURSDAY",
	6: "FRIDAY",
	7: "SATURDAY",
	8: "SUNDAY",
}

var DayOfWeekEnum_DayOfWeek_value = map[string]int32{
	"UNSPECIFIED": 0,
	"UNKNOWN":     1,
	"MONDAY":      2,
	"TUESDAY":     3,
	"WEDNESDAY":   4,
	"THURSDAY":    5,
	"FRIDAY":      6,
	"SATURDAY":    7,
	"SUNDAY":      8,
}

func (x DayOfWeekEnum_DayOfWeek) String() string {
	return proto.EnumName(DayOfWeekEnum_DayOfWeek_name, int32(x))
}

func (DayOfWeekEnum_DayOfWeek) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_e8f57aab9f85f172, []int{0, 0}
}

// Container for enumeration of days of the week, e.g., "Monday".
type DayOfWeekEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DayOfWeekEnum) Reset()         { *m = DayOfWeekEnum{} }
func (m *DayOfWeekEnum) String() string { return proto.CompactTextString(m) }
func (*DayOfWeekEnum) ProtoMessage()    {}
func (*DayOfWeekEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_e8f57aab9f85f172, []int{0}
}

func (m *DayOfWeekEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DayOfWeekEnum.Unmarshal(m, b)
}
func (m *DayOfWeekEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DayOfWeekEnum.Marshal(b, m, deterministic)
}
func (m *DayOfWeekEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DayOfWeekEnum.Merge(m, src)
}
func (m *DayOfWeekEnum) XXX_Size() int {
	return xxx_messageInfo_DayOfWeekEnum.Size(m)
}
func (m *DayOfWeekEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_DayOfWeekEnum.DiscardUnknown(m)
}

var xxx_messageInfo_DayOfWeekEnum proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("google.ads.googleads.v0.enums.DayOfWeekEnum_DayOfWeek", DayOfWeekEnum_DayOfWeek_name, DayOfWeekEnum_DayOfWeek_value)
	proto.RegisterType((*DayOfWeekEnum)(nil), "google.ads.googleads.v0.enums.DayOfWeekEnum")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v0/enums/day_of_week.proto", fileDescriptor_e8f57aab9f85f172)
}

var fileDescriptor_e8f57aab9f85f172 = []byte{
	// 324 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x50, 0x4f, 0x4e, 0x3a, 0x31,
	0x18, 0xfd, 0xcd, 0xf0, 0x93, 0x3f, 0x45, 0x94, 0x74, 0xcf, 0x02, 0x0e, 0xd0, 0x99, 0xc4, 0x5d,
	0x5d, 0x15, 0x67, 0x40, 0x62, 0x2c, 0x84, 0xa1, 0x10, 0xcd, 0x24, 0x64, 0xb4, 0xa5, 0x31, 0x30,
	0x53, 0x42, 0x05, 0xc3, 0x01, 0x3c, 0x83, 0x7b, 0x97, 0x1e, 0xc5, 0x6b, 0xb8, 0xf3, 0x14, 0xa6,
	0x1d, 0x98, 0x9d, 0x6e, 0x9a, 0xf7, 0xbe, 0xf7, 0xbd, 0x2f, 0xaf, 0x0f, 0x78, 0x52, 0x29, 0xb9,
	0x12, 0x5e, 0xc2, 0xf5, 0x01, 0x1a, 0xb4, 0xf3, 0x3d, 0x91, 0x6d, 0x53, 0xed, 0xf1, 0x64, 0x3f,
	0x57, 0x8b, 0xf9, 0x8b, 0x10, 0x4b, 0xb4, 0xde, 0xa8, 0x67, 0x05, 0x5b, 0xf9, 0x16, 0x4a, 0xb8,
	0x46, 0x85, 0x01, 0xed, 0x7c, 0x64, 0x0d, 0x9d, 0x37, 0x07, 0x34, 0x82, 0x64, 0x3f, 0x5c, 0xcc,
	0x84, 0x58, 0x86, 0xd9, 0x36, 0xed, 0xbc, 0x3a, 0xa0, 0x56, 0x4c, 0xe0, 0x39, 0xa8, 0x33, 0x1a,
	0x8d, 0xc2, 0xab, 0x41, 0x6f, 0x10, 0x06, 0xcd, 0x7f, 0xb0, 0x0e, 0x2a, 0x8c, 0xde, 0xd0, 0xe1,
	0x8c, 0x36, 0x1d, 0x08, 0x40, 0xf9, 0x76, 0x48, 0x03, 0x72, 0xd7, 0x74, 0x8d, 0x30, 0x61, 0x61,
	0x64, 0x48, 0x09, 0x36, 0x40, 0x6d, 0x16, 0x06, 0x34, 0xa7, 0xff, 0xe1, 0x29, 0xa8, 0x4e, 0xae,
	0xd9, 0xd8, 0xb2, 0x13, 0xe3, 0xea, 0x8d, 0x07, 0x06, 0x97, 0x8d, 0x12, 0x91, 0x09, 0x1b, 0x1b,
	0x56, 0x31, 0x4a, 0xc4, 0xec, 0xbd, 0x6a, 0xf7, 0xcb, 0x01, 0xed, 0x47, 0x95, 0xa2, 0x3f, 0xf3,
	0x77, 0xcf, 0x8a, 0xa8, 0x23, 0xf3, 0xdd, 0x91, 0x73, 0xdf, 0x3d, 0x18, 0xa4, 0x5a, 0x25, 0x99,
	0x44, 0x6a, 0x23, 0x3d, 0x29, 0x32, 0x5b, 0xc6, 0xb1, 0xb1, 0xf5, 0x93, 0xfe, 0xa5, 0xc0, 0x4b,
	0xfb, 0xbe, 0xbb, 0xa5, 0x3e, 0x21, 0x1f, 0x6e, 0xab, 0x9f, 0x9f, 0x22, 0x5c, 0xa3, 0x1c, 0x1a,
	0x34, 0xf5, 0x91, 0x29, 0x4a, 0x7f, 0x1e, 0xf5, 0x98, 0x70, 0x1d, 0x17, 0x7a, 0x3c, 0xf5, 0x63,
	0xab, 0x7f, 0xbb, 0xed, 0x7c, 0x88, 0x31, 0xe1, 0x1a, 0xe3, 0x62, 0x03, 0xe3, 0xa9, 0x8f, 0xb1,
	0xdd, 0x79, 0x28, 0xdb, 0x60, 0x17, 0x3f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x0a, 0x09, 0x94, 0x1f,
	0xd8, 0x01, 0x00, 0x00,
}
