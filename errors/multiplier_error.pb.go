// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v0/errors/multiplier_error.proto

package errors

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

// Enum describing possible multiplier errors.
type MultiplierErrorEnum_MultiplierError int32

const (
	// Enum unspecified.
	MultiplierErrorEnum_UNSPECIFIED MultiplierErrorEnum_MultiplierError = 0
	// The received error code is not known in this version.
	MultiplierErrorEnum_UNKNOWN MultiplierErrorEnum_MultiplierError = 1
	// Multiplier value is too high
	MultiplierErrorEnum_MULTIPLIER_TOO_HIGH MultiplierErrorEnum_MultiplierError = 2
	// Multiplier value is too low
	MultiplierErrorEnum_MULTIPLIER_TOO_LOW MultiplierErrorEnum_MultiplierError = 3
	// Too many fractional digits
	MultiplierErrorEnum_TOO_MANY_FRACTIONAL_DIGITS MultiplierErrorEnum_MultiplierError = 4
	// A multiplier cannot be set for this bidding strategy
	MultiplierErrorEnum_MULTIPLIER_NOT_ALLOWED_FOR_BIDDING_STRATEGY MultiplierErrorEnum_MultiplierError = 5
	// A multiplier cannot be set when there is no base bid (e.g., content max
	// cpc)
	MultiplierErrorEnum_MULTIPLIER_NOT_ALLOWED_WHEN_BASE_BID_IS_MISSING MultiplierErrorEnum_MultiplierError = 6
	// A bid multiplier must be specified
	MultiplierErrorEnum_NO_MULTIPLIER_SPECIFIED MultiplierErrorEnum_MultiplierError = 7
	// Multiplier causes bid to exceed daily budget
	MultiplierErrorEnum_MULTIPLIER_CAUSES_BID_TO_EXCEED_DAILY_BUDGET MultiplierErrorEnum_MultiplierError = 8
	// Multiplier causes bid to exceed monthly budget
	MultiplierErrorEnum_MULTIPLIER_CAUSES_BID_TO_EXCEED_MONTHLY_BUDGET MultiplierErrorEnum_MultiplierError = 9
	// Multiplier causes bid to exceed custom budget
	MultiplierErrorEnum_MULTIPLIER_CAUSES_BID_TO_EXCEED_CUSTOM_BUDGET MultiplierErrorEnum_MultiplierError = 10
	// Multiplier causes bid to exceed maximum allowed bid
	MultiplierErrorEnum_MULTIPLIER_CAUSES_BID_TO_EXCEED_MAX_ALLOWED_BID MultiplierErrorEnum_MultiplierError = 11
	// Multiplier causes bid to become less than the minimum bid allowed
	MultiplierErrorEnum_BID_LESS_THAN_MIN_ALLOWED_BID_WITH_MULTIPLIER MultiplierErrorEnum_MultiplierError = 12
	// Multiplier type (cpc vs. cpm) needs to match campaign's bidding strategy
	MultiplierErrorEnum_MULTIPLIER_AND_BIDDING_STRATEGY_TYPE_MISMATCH MultiplierErrorEnum_MultiplierError = 13
)

var MultiplierErrorEnum_MultiplierError_name = map[int32]string{
	0:  "UNSPECIFIED",
	1:  "UNKNOWN",
	2:  "MULTIPLIER_TOO_HIGH",
	3:  "MULTIPLIER_TOO_LOW",
	4:  "TOO_MANY_FRACTIONAL_DIGITS",
	5:  "MULTIPLIER_NOT_ALLOWED_FOR_BIDDING_STRATEGY",
	6:  "MULTIPLIER_NOT_ALLOWED_WHEN_BASE_BID_IS_MISSING",
	7:  "NO_MULTIPLIER_SPECIFIED",
	8:  "MULTIPLIER_CAUSES_BID_TO_EXCEED_DAILY_BUDGET",
	9:  "MULTIPLIER_CAUSES_BID_TO_EXCEED_MONTHLY_BUDGET",
	10: "MULTIPLIER_CAUSES_BID_TO_EXCEED_CUSTOM_BUDGET",
	11: "MULTIPLIER_CAUSES_BID_TO_EXCEED_MAX_ALLOWED_BID",
	12: "BID_LESS_THAN_MIN_ALLOWED_BID_WITH_MULTIPLIER",
	13: "MULTIPLIER_AND_BIDDING_STRATEGY_TYPE_MISMATCH",
}

var MultiplierErrorEnum_MultiplierError_value = map[string]int32{
	"UNSPECIFIED":                0,
	"UNKNOWN":                    1,
	"MULTIPLIER_TOO_HIGH":        2,
	"MULTIPLIER_TOO_LOW":         3,
	"TOO_MANY_FRACTIONAL_DIGITS": 4,
	"MULTIPLIER_NOT_ALLOWED_FOR_BIDDING_STRATEGY":     5,
	"MULTIPLIER_NOT_ALLOWED_WHEN_BASE_BID_IS_MISSING": 6,
	"NO_MULTIPLIER_SPECIFIED":                         7,
	"MULTIPLIER_CAUSES_BID_TO_EXCEED_DAILY_BUDGET":    8,
	"MULTIPLIER_CAUSES_BID_TO_EXCEED_MONTHLY_BUDGET":  9,
	"MULTIPLIER_CAUSES_BID_TO_EXCEED_CUSTOM_BUDGET":   10,
	"MULTIPLIER_CAUSES_BID_TO_EXCEED_MAX_ALLOWED_BID": 11,
	"BID_LESS_THAN_MIN_ALLOWED_BID_WITH_MULTIPLIER":   12,
	"MULTIPLIER_AND_BIDDING_STRATEGY_TYPE_MISMATCH":   13,
}

func (x MultiplierErrorEnum_MultiplierError) String() string {
	return proto.EnumName(MultiplierErrorEnum_MultiplierError_name, int32(x))
}

func (MultiplierErrorEnum_MultiplierError) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_bb6c3c8aa121bbde, []int{0, 0}
}

// Container for enum describing possible multiplier errors.
type MultiplierErrorEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MultiplierErrorEnum) Reset()         { *m = MultiplierErrorEnum{} }
func (m *MultiplierErrorEnum) String() string { return proto.CompactTextString(m) }
func (*MultiplierErrorEnum) ProtoMessage()    {}
func (*MultiplierErrorEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_bb6c3c8aa121bbde, []int{0}
}

func (m *MultiplierErrorEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MultiplierErrorEnum.Unmarshal(m, b)
}
func (m *MultiplierErrorEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MultiplierErrorEnum.Marshal(b, m, deterministic)
}
func (m *MultiplierErrorEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MultiplierErrorEnum.Merge(m, src)
}
func (m *MultiplierErrorEnum) XXX_Size() int {
	return xxx_messageInfo_MultiplierErrorEnum.Size(m)
}
func (m *MultiplierErrorEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_MultiplierErrorEnum.DiscardUnknown(m)
}

var xxx_messageInfo_MultiplierErrorEnum proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("google.ads.googleads.v0.errors.MultiplierErrorEnum_MultiplierError", MultiplierErrorEnum_MultiplierError_name, MultiplierErrorEnum_MultiplierError_value)
	proto.RegisterType((*MultiplierErrorEnum)(nil), "google.ads.googleads.v0.errors.MultiplierErrorEnum")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v0/errors/multiplier_error.proto", fileDescriptor_bb6c3c8aa121bbde)
}

var fileDescriptor_bb6c3c8aa121bbde = []byte{
	// 505 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x93, 0xdf, 0x6e, 0xd3, 0x3e,
	0x14, 0xc7, 0x7f, 0xed, 0xfa, 0xdb, 0xc0, 0x05, 0xcd, 0xf2, 0x10, 0x93, 0x40, 0xea, 0x45, 0x6f,
	0x81, 0x24, 0x30, 0x71, 0x13, 0xae, 0x9c, 0xd8, 0x4d, 0x2c, 0x12, 0xbb, 0xaa, 0x9d, 0x76, 0x45,
	0x95, 0x8e, 0x0a, 0xad, 0xa2, 0x4a, 0x6d, 0x53, 0x25, 0xdb, 0x1e, 0x88, 0x4b, 0x5e, 0x80, 0x77,
	0xe0, 0x86, 0xf7, 0xe0, 0x86, 0x57, 0x40, 0x69, 0x68, 0x57, 0x05, 0xc6, 0xae, 0x72, 0x72, 0xfc,
	0xfd, 0xf8, 0xfc, 0xd1, 0xd7, 0xe8, 0x6d, 0x9a, 0x65, 0xe9, 0x72, 0x6e, 0x4f, 0x67, 0x85, 0x5d,
	0x85, 0x65, 0x74, 0xe3, 0xd8, 0xf3, 0x3c, 0xcf, 0xf2, 0xc2, 0x5e, 0x5d, 0x2f, 0xaf, 0x16, 0x9b,
	0xe5, 0x62, 0x9e, 0xc3, 0x36, 0x63, 0x6d, 0xf2, 0xec, 0x2a, 0x23, 0x9d, 0x4a, 0x6b, 0x4d, 0x67,
	0x85, 0xb5, 0xc7, 0xac, 0x1b, 0xc7, 0xaa, 0xb0, 0xee, 0xf7, 0x16, 0x3a, 0x8b, 0xf7, 0x28, 0x2f,
	0x93, 0x7c, 0x7d, 0xbd, 0xea, 0x7e, 0x6d, 0xa1, 0xd3, 0x5a, 0x9e, 0x9c, 0xa2, 0x76, 0x22, 0x75,
	0x9f, 0xfb, 0xa2, 0x27, 0x38, 0xc3, 0xff, 0x91, 0x36, 0x3a, 0x49, 0xe4, 0x7b, 0xa9, 0x46, 0x12,
	0x37, 0xc8, 0x39, 0x3a, 0x8b, 0x93, 0xc8, 0x88, 0x7e, 0x24, 0xf8, 0x00, 0x8c, 0x52, 0x10, 0x8a,
	0x20, 0xc4, 0x4d, 0xf2, 0x14, 0x91, 0xda, 0x41, 0xa4, 0x46, 0xf8, 0x88, 0x74, 0xd0, 0xb3, 0xf2,
	0x27, 0xa6, 0x72, 0x0c, 0xbd, 0x01, 0xf5, 0x8d, 0x50, 0x92, 0x46, 0xc0, 0x44, 0x20, 0x8c, 0xc6,
	0x2d, 0x62, 0xa3, 0x17, 0x07, 0x9c, 0x54, 0x06, 0x68, 0x14, 0xa9, 0x11, 0x67, 0xd0, 0x53, 0x03,
	0xf0, 0x04, 0x63, 0x42, 0x06, 0xa0, 0xcd, 0x80, 0x1a, 0x1e, 0x8c, 0xf1, 0xff, 0xe4, 0x02, 0xd9,
	0x77, 0x00, 0xa3, 0x90, 0x4b, 0xf0, 0xa8, 0xe6, 0x25, 0x06, 0x42, 0x43, 0x2c, 0xb4, 0x16, 0x32,
	0xc0, 0xc7, 0xe4, 0x39, 0x3a, 0x97, 0x0a, 0x0e, 0xb8, 0xdb, 0x01, 0x4f, 0x88, 0x83, 0x5e, 0x1e,
	0x9c, 0xf8, 0x34, 0xd1, 0x5c, 0x6f, 0xaf, 0x30, 0x0a, 0xf8, 0xa5, 0xcf, 0x39, 0x03, 0x46, 0x45,
	0x34, 0x06, 0x2f, 0x61, 0x01, 0x37, 0xf8, 0x01, 0x79, 0x83, 0xac, 0xfb, 0x88, 0x58, 0x49, 0x13,
	0xde, 0x32, 0x0f, 0xc9, 0x6b, 0xf4, 0xea, 0x3e, 0xc6, 0x4f, 0xb4, 0x51, 0xf1, 0x0e, 0x41, 0xb5,
	0x51, 0xff, 0x5e, 0x86, 0x5e, 0xee, 0xe7, 0xf7, 0x04, 0xc3, 0xed, 0xb2, 0x4e, 0x29, 0x89, 0xb8,
	0xd6, 0x60, 0x42, 0x2a, 0x21, 0x16, 0xf2, 0x50, 0x02, 0x23, 0x61, 0xc2, 0x83, 0x55, 0xe0, 0x47,
	0xb5, 0xd6, 0xa8, 0x64, 0x7f, 0xec, 0x1d, 0xcc, 0xb8, 0xcf, 0xcb, 0x7d, 0xc6, 0xd4, 0xf8, 0x21,
	0x7e, 0xec, 0xfd, 0x6c, 0xa0, 0xee, 0xa7, 0x6c, 0x65, 0xfd, 0xdb, 0x78, 0xde, 0x93, 0x9a, 0xbb,
	0xfa, 0xa5, 0x5d, 0xfb, 0x8d, 0x0f, 0xec, 0x37, 0x97, 0x66, 0xcb, 0xe9, 0x3a, 0xb5, 0xb2, 0x3c,
	0xb5, 0xd3, 0xf9, 0x7a, 0x6b, 0xe6, 0x9d, 0xef, 0x37, 0x8b, 0xe2, 0xae, 0x67, 0xf0, 0xae, 0xfa,
	0x7c, 0x6e, 0x1e, 0x05, 0x94, 0x7e, 0x69, 0x76, 0x82, 0xea, 0x32, 0x3a, 0x2b, 0xac, 0x2a, 0x2c,
	0xa3, 0xa1, 0x63, 0x6d, 0x4b, 0x16, 0xdf, 0x76, 0x82, 0x09, 0x9d, 0x15, 0x93, 0xbd, 0x60, 0x32,
	0x74, 0x26, 0x95, 0xe0, 0x47, 0xb3, 0x5b, 0x65, 0x5d, 0x97, 0xce, 0x0a, 0xd7, 0xdd, 0x4b, 0x5c,
	0x77, 0xe8, 0xb8, 0x6e, 0x25, 0xfa, 0x78, 0xbc, 0xed, 0xee, 0xe2, 0x57, 0x00, 0x00, 0x00, 0xff,
	0xff, 0x72, 0x8f, 0xb6, 0xc9, 0xa3, 0x03, 0x00, 0x00,
}
