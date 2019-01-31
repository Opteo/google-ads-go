// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v0/enums/ad_group_ad_rotation_mode.proto

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

// The possible ad rotation modes of an ad group.
type AdGroupAdRotationModeEnum_AdGroupAdRotationMode int32

const (
	// The ad rotation mode has not been specified.
	AdGroupAdRotationModeEnum_UNSPECIFIED AdGroupAdRotationModeEnum_AdGroupAdRotationMode = 0
	// The received value is not known in this version.
	//
	// This is a response-only value.
	AdGroupAdRotationModeEnum_UNKNOWN AdGroupAdRotationModeEnum_AdGroupAdRotationMode = 1
	// Optimize ad group ads based on clicks or conversions.
	AdGroupAdRotationModeEnum_OPTIMIZE AdGroupAdRotationModeEnum_AdGroupAdRotationMode = 2
	// Rotate evenly forever.
	AdGroupAdRotationModeEnum_ROTATE_FOREVER AdGroupAdRotationModeEnum_AdGroupAdRotationMode = 3
)

var AdGroupAdRotationModeEnum_AdGroupAdRotationMode_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "UNKNOWN",
	2: "OPTIMIZE",
	3: "ROTATE_FOREVER",
}

var AdGroupAdRotationModeEnum_AdGroupAdRotationMode_value = map[string]int32{
	"UNSPECIFIED":    0,
	"UNKNOWN":        1,
	"OPTIMIZE":       2,
	"ROTATE_FOREVER": 3,
}

func (x AdGroupAdRotationModeEnum_AdGroupAdRotationMode) String() string {
	return proto.EnumName(AdGroupAdRotationModeEnum_AdGroupAdRotationMode_name, int32(x))
}

func (AdGroupAdRotationModeEnum_AdGroupAdRotationMode) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_9d8204fe8a2690e1, []int{0, 0}
}

// Container for enum describing possible ad rotation modes of ads within an
// ad group.
type AdGroupAdRotationModeEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AdGroupAdRotationModeEnum) Reset()         { *m = AdGroupAdRotationModeEnum{} }
func (m *AdGroupAdRotationModeEnum) String() string { return proto.CompactTextString(m) }
func (*AdGroupAdRotationModeEnum) ProtoMessage()    {}
func (*AdGroupAdRotationModeEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_9d8204fe8a2690e1, []int{0}
}

func (m *AdGroupAdRotationModeEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AdGroupAdRotationModeEnum.Unmarshal(m, b)
}
func (m *AdGroupAdRotationModeEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AdGroupAdRotationModeEnum.Marshal(b, m, deterministic)
}
func (m *AdGroupAdRotationModeEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AdGroupAdRotationModeEnum.Merge(m, src)
}
func (m *AdGroupAdRotationModeEnum) XXX_Size() int {
	return xxx_messageInfo_AdGroupAdRotationModeEnum.Size(m)
}
func (m *AdGroupAdRotationModeEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_AdGroupAdRotationModeEnum.DiscardUnknown(m)
}

var xxx_messageInfo_AdGroupAdRotationModeEnum proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("google.ads.googleads.v0.enums.AdGroupAdRotationModeEnum_AdGroupAdRotationMode", AdGroupAdRotationModeEnum_AdGroupAdRotationMode_name, AdGroupAdRotationModeEnum_AdGroupAdRotationMode_value)
	proto.RegisterType((*AdGroupAdRotationModeEnum)(nil), "google.ads.googleads.v0.enums.AdGroupAdRotationModeEnum")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v0/enums/ad_group_ad_rotation_mode.proto", fileDescriptor_9d8204fe8a2690e1)
}

var fileDescriptor_9d8204fe8a2690e1 = []byte{
	// 308 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x90, 0xcf, 0x4a, 0xf3, 0x40,
	0x14, 0xc5, 0xbf, 0xa6, 0xf0, 0x29, 0x53, 0xd1, 0x30, 0xe0, 0x42, 0xa1, 0x8b, 0xf6, 0x01, 0x26,
	0x01, 0x77, 0x23, 0x2e, 0xa6, 0x3a, 0x2d, 0x41, 0x9a, 0x84, 0xd8, 0xa6, 0x50, 0x02, 0x61, 0x74,
	0xc2, 0x50, 0x68, 0x72, 0x4b, 0x26, 0xe9, 0x03, 0xb9, 0xf4, 0x51, 0x7c, 0x14, 0x37, 0xbe, 0x82,
	0x64, 0xd2, 0x66, 0x55, 0xdd, 0x0c, 0x87, 0x39, 0xf7, 0x77, 0xff, 0x1c, 0xf4, 0xa0, 0x00, 0xd4,
	0x36, 0x73, 0x84, 0xd4, 0x4e, 0x2b, 0x1b, 0xb5, 0x77, 0x9d, 0xac, 0xa8, 0x73, 0xed, 0x08, 0x99,
	0xaa, 0x12, 0xea, 0x5d, 0x2a, 0x64, 0x5a, 0x42, 0x25, 0xaa, 0x0d, 0x14, 0x69, 0x0e, 0x32, 0x23,
	0xbb, 0x12, 0x2a, 0xc0, 0xc3, 0x96, 0x21, 0x42, 0x6a, 0xd2, 0xe1, 0x64, 0xef, 0x12, 0x83, 0x8f,
	0x2b, 0x74, 0xc3, 0xe4, 0xac, 0x69, 0xc0, 0x64, 0x74, 0xc0, 0xe7, 0x20, 0x33, 0x5e, 0xd4, 0xf9,
	0x78, 0x85, 0xae, 0x4f, 0x9a, 0xf8, 0x0a, 0x0d, 0x96, 0xfe, 0x4b, 0xc8, 0x1f, 0xbd, 0xa9, 0xc7,
	0x9f, 0xec, 0x7f, 0x78, 0x80, 0xce, 0x96, 0xfe, 0xb3, 0x1f, 0xac, 0x7c, 0xbb, 0x87, 0x2f, 0xd0,
	0x79, 0x10, 0x2e, 0xbc, 0xb9, 0xb7, 0xe6, 0xb6, 0x85, 0x31, 0xba, 0x8c, 0x82, 0x05, 0x5b, 0xf0,
	0x74, 0x1a, 0x44, 0x3c, 0xe6, 0x91, 0xdd, 0x9f, 0x7c, 0xf7, 0xd0, 0xe8, 0x0d, 0x72, 0xf2, 0xe7,
	0x6e, 0x93, 0xdb, 0x93, 0xc3, 0xc3, 0xe6, 0xac, 0xb0, 0xb7, 0x9e, 0x1c, 0x60, 0x05, 0x5b, 0x51,
	0x28, 0x02, 0xa5, 0x72, 0x54, 0x56, 0x98, 0xa3, 0x8f, 0x39, 0xed, 0x36, 0xfa, 0x97, 0xd8, 0xee,
	0xcd, 0xfb, 0x6e, 0xf5, 0x67, 0x8c, 0x7d, 0x58, 0xc3, 0x59, 0xdb, 0x8a, 0x49, 0x4d, 0x5a, 0xd9,
	0xa8, 0xd8, 0x25, 0x4d, 0x0a, 0xfa, 0xf3, 0xe8, 0x27, 0x4c, 0xea, 0xa4, 0xf3, 0x93, 0xd8, 0x4d,
	0x8c, 0xff, 0x65, 0x8d, 0xda, 0x4f, 0x4a, 0x99, 0xd4, 0x94, 0x76, 0x15, 0x94, 0xc6, 0x2e, 0xa5,
	0xa6, 0xe6, 0xf5, 0xbf, 0x59, 0xec, 0xee, 0x27, 0x00, 0x00, 0xff, 0xff, 0xb4, 0xbf, 0xfb, 0x61,
	0xce, 0x01, 0x00, 0x00,
}
