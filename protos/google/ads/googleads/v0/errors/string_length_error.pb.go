// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v0/errors/string_length_error.proto

package google_ads_googleads_v0_errors

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
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Enum describing possible string length errors.
type StringLengthErrorEnum_StringLengthError int32

const (
	// Enum unspecified.
	StringLengthErrorEnum_UNSPECIFIED StringLengthErrorEnum_StringLengthError = 0
	// The received error code is not known in this version.
	StringLengthErrorEnum_UNKNOWN StringLengthErrorEnum_StringLengthError = 1
	// Too short.
	StringLengthErrorEnum_TOO_SHORT StringLengthErrorEnum_StringLengthError = 2
	// Too long.
	StringLengthErrorEnum_TOO_LONG StringLengthErrorEnum_StringLengthError = 3
)

var StringLengthErrorEnum_StringLengthError_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "UNKNOWN",
	2: "TOO_SHORT",
	3: "TOO_LONG",
}

var StringLengthErrorEnum_StringLengthError_value = map[string]int32{
	"UNSPECIFIED": 0,
	"UNKNOWN":     1,
	"TOO_SHORT":   2,
	"TOO_LONG":    3,
}

func (x StringLengthErrorEnum_StringLengthError) String() string {
	return proto.EnumName(StringLengthErrorEnum_StringLengthError_name, int32(x))
}

func (StringLengthErrorEnum_StringLengthError) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_f1bc703f5ad048f9, []int{0, 0}
}

// Container for enum describing possible string length errors.
type StringLengthErrorEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StringLengthErrorEnum) Reset()         { *m = StringLengthErrorEnum{} }
func (m *StringLengthErrorEnum) String() string { return proto.CompactTextString(m) }
func (*StringLengthErrorEnum) ProtoMessage()    {}
func (*StringLengthErrorEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_f1bc703f5ad048f9, []int{0}
}
func (m *StringLengthErrorEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StringLengthErrorEnum.Unmarshal(m, b)
}
func (m *StringLengthErrorEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StringLengthErrorEnum.Marshal(b, m, deterministic)
}
func (m *StringLengthErrorEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StringLengthErrorEnum.Merge(m, src)
}
func (m *StringLengthErrorEnum) XXX_Size() int {
	return xxx_messageInfo_StringLengthErrorEnum.Size(m)
}
func (m *StringLengthErrorEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_StringLengthErrorEnum.DiscardUnknown(m)
}

var xxx_messageInfo_StringLengthErrorEnum proto.InternalMessageInfo

func init() {
	proto.RegisterType((*StringLengthErrorEnum)(nil), "google.ads.googleads.v0.errors.StringLengthErrorEnum")
	proto.RegisterEnum("google.ads.googleads.v0.errors.StringLengthErrorEnum_StringLengthError", StringLengthErrorEnum_StringLengthError_name, StringLengthErrorEnum_StringLengthError_value)
}

func init() {
	proto.RegisterFile("google/ads/googleads/v0/errors/string_length_error.proto", fileDescriptor_f1bc703f5ad048f9)
}

var fileDescriptor_f1bc703f5ad048f9 = []byte{
	// 241 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xb2, 0x48, 0xcf, 0xcf, 0x4f,
	0xcf, 0x49, 0xd5, 0x4f, 0x4c, 0x29, 0xd6, 0x87, 0x30, 0x41, 0xac, 0x32, 0x03, 0xfd, 0xd4, 0xa2,
	0xa2, 0xfc, 0xa2, 0x62, 0xfd, 0xe2, 0x92, 0xa2, 0xcc, 0xbc, 0xf4, 0xf8, 0x9c, 0xd4, 0xbc, 0xf4,
	0x92, 0x8c, 0x78, 0xb0, 0xa0, 0x5e, 0x41, 0x51, 0x7e, 0x49, 0xbe, 0x90, 0x1c, 0x44, 0xb9, 0x5e,
	0x62, 0x4a, 0xb1, 0x1e, 0x5c, 0xa7, 0x5e, 0x99, 0x81, 0x1e, 0x44, 0xa7, 0x52, 0x3a, 0x97, 0x68,
	0x30, 0x58, 0xb3, 0x0f, 0x58, 0xaf, 0x2b, 0x48, 0xd4, 0x35, 0xaf, 0x34, 0x57, 0xc9, 0x8f, 0x4b,
	0x10, 0x43, 0x42, 0x88, 0x9f, 0x8b, 0x3b, 0xd4, 0x2f, 0x38, 0xc0, 0xd5, 0xd9, 0xd3, 0xcd, 0xd3,
	0xd5, 0x45, 0x80, 0x41, 0x88, 0x9b, 0x8b, 0x3d, 0xd4, 0xcf, 0xdb, 0xcf, 0x3f, 0xdc, 0x4f, 0x80,
	0x51, 0x88, 0x97, 0x8b, 0x33, 0xc4, 0xdf, 0x3f, 0x3e, 0xd8, 0xc3, 0x3f, 0x28, 0x44, 0x80, 0x49,
	0x88, 0x87, 0x8b, 0x03, 0xc4, 0xf5, 0xf1, 0xf7, 0x73, 0x17, 0x60, 0x76, 0x5a, 0xce, 0xc8, 0xa5,
	0x94, 0x9c, 0x9f, 0xab, 0x87, 0xdf, 0x3d, 0x4e, 0x62, 0x18, 0x96, 0x06, 0x80, 0xfc, 0x11, 0xc0,
	0xb8, 0x88, 0x89, 0xd9, 0xdd, 0xd1, 0x71, 0x15, 0x93, 0x9c, 0x3b, 0xc4, 0x00, 0xc7, 0x94, 0x62,
	0x3d, 0x08, 0x13, 0xc4, 0x0a, 0x33, 0xd0, 0x03, 0x2b, 0x2e, 0x3e, 0x05, 0x53, 0x10, 0xe3, 0x98,
	0x52, 0x1c, 0x03, 0x57, 0x10, 0x13, 0x66, 0x10, 0x03, 0x51, 0xf0, 0x88, 0x90, 0x82, 0x24, 0x36,
	0x70, 0xc8, 0x19, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0x4a, 0x76, 0x29, 0xff, 0x75, 0x01, 0x00,
	0x00,
}