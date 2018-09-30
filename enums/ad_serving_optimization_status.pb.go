// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/kritzware/google-ads-go/enums/ad_serving_optimization_status.proto

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
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Enum describing possible serving statuses.
type AdServingOptimizationStatusEnum_AdServingOptimizationStatus int32

const (
	// No value has been specified.
	AdServingOptimizationStatusEnum_UNSPECIFIED AdServingOptimizationStatusEnum_AdServingOptimizationStatus = 0
	// The received value is not known in this version.
	//
	// This is a response-only value.
	AdServingOptimizationStatusEnum_UNKNOWN AdServingOptimizationStatusEnum_AdServingOptimizationStatus = 1
	// Ad serving is optimized based on CTR for the campaign.
	AdServingOptimizationStatusEnum_OPTIMIZE AdServingOptimizationStatusEnum_AdServingOptimizationStatus = 2
	// Ad serving is optimized based on CTR * Conversion for the campaign. If
	// the campaign is not in the conversion optimizer bidding strategy, it will
	// default to OPTIMIZED.
	AdServingOptimizationStatusEnum_CONVERSION_OPTIMIZE AdServingOptimizationStatusEnum_AdServingOptimizationStatus = 3
	// Ads are rotated evenly for 90 days, then optimized for clicks.
	AdServingOptimizationStatusEnum_ROTATE AdServingOptimizationStatusEnum_AdServingOptimizationStatus = 4
	// Show lower performing ads more evenly with higher performing ads, and do
	// not optimize.
	AdServingOptimizationStatusEnum_ROTATE_INDEFINITELY AdServingOptimizationStatusEnum_AdServingOptimizationStatus = 5
	// Ad serving optimization status is not available.
	AdServingOptimizationStatusEnum_UNAVAILABLE AdServingOptimizationStatusEnum_AdServingOptimizationStatus = 6
)

var AdServingOptimizationStatusEnum_AdServingOptimizationStatus_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "UNKNOWN",
	2: "OPTIMIZE",
	3: "CONVERSION_OPTIMIZE",
	4: "ROTATE",
	5: "ROTATE_INDEFINITELY",
	6: "UNAVAILABLE",
}

var AdServingOptimizationStatusEnum_AdServingOptimizationStatus_value = map[string]int32{
	"UNSPECIFIED":         0,
	"UNKNOWN":             1,
	"OPTIMIZE":            2,
	"CONVERSION_OPTIMIZE": 3,
	"ROTATE":              4,
	"ROTATE_INDEFINITELY": 5,
	"UNAVAILABLE":         6,
}

func (x AdServingOptimizationStatusEnum_AdServingOptimizationStatus) String() string {
	return proto.EnumName(AdServingOptimizationStatusEnum_AdServingOptimizationStatus_name, int32(x))
}

func (AdServingOptimizationStatusEnum_AdServingOptimizationStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_3e2bd4deca74525f, []int{0, 0}
}

// Possible ad serving statuses of a campaign.
type AdServingOptimizationStatusEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AdServingOptimizationStatusEnum) Reset()         { *m = AdServingOptimizationStatusEnum{} }
func (m *AdServingOptimizationStatusEnum) String() string { return proto.CompactTextString(m) }
func (*AdServingOptimizationStatusEnum) ProtoMessage()    {}
func (*AdServingOptimizationStatusEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_3e2bd4deca74525f, []int{0}
}
func (m *AdServingOptimizationStatusEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AdServingOptimizationStatusEnum.Unmarshal(m, b)
}
func (m *AdServingOptimizationStatusEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AdServingOptimizationStatusEnum.Marshal(b, m, deterministic)
}
func (m *AdServingOptimizationStatusEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AdServingOptimizationStatusEnum.Merge(m, src)
}
func (m *AdServingOptimizationStatusEnum) XXX_Size() int {
	return xxx_messageInfo_AdServingOptimizationStatusEnum.Size(m)
}
func (m *AdServingOptimizationStatusEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_AdServingOptimizationStatusEnum.DiscardUnknown(m)
}

var xxx_messageInfo_AdServingOptimizationStatusEnum proto.InternalMessageInfo

func init() {
	proto.RegisterType((*AdServingOptimizationStatusEnum)(nil), "google.ads.googleads.v0.enums.AdServingOptimizationStatusEnum")
	proto.RegisterEnum("google.ads.googleads.v0.enums.AdServingOptimizationStatusEnum_AdServingOptimizationStatus", AdServingOptimizationStatusEnum_AdServingOptimizationStatus_name, AdServingOptimizationStatusEnum_AdServingOptimizationStatus_value)
}

func init() {
	proto.RegisterFile("github.com/kritzware/google-ads-go/enums/ad_serving_optimization_status.proto", fileDescriptor_3e2bd4deca74525f)
}

var fileDescriptor_3e2bd4deca74525f = []byte{
	// 306 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x91, 0x51, 0x4a, 0xc3, 0x30,
	0x18, 0x80, 0xed, 0xa6, 0x53, 0x32, 0xc1, 0x10, 0x1f, 0x7c, 0x90, 0xa1, 0xee, 0x00, 0x69, 0xc1,
	0x13, 0xa4, 0x5b, 0x36, 0x82, 0x33, 0x29, 0x6b, 0x57, 0x51, 0x0a, 0xa5, 0x9a, 0x31, 0x0a, 0xb6,
	0x19, 0x4b, 0xdb, 0x07, 0x4f, 0xe2, 0xb3, 0x8f, 0x82, 0x37, 0xf0, 0x04, 0x1e, 0xc3, 0x93, 0x48,
	0x1b, 0x2c, 0x3e, 0xd5, 0xb7, 0x0f, 0xbe, 0xbf, 0x5f, 0x93, 0x3f, 0xc0, 0xdd, 0x28, 0xb5, 0x79,
	0x5e, 0xdb, 0x89, 0xd4, 0xb6, 0xc1, 0x9a, 0x2a, 0xc7, 0x5e, 0xe7, 0x65, 0xa6, 0xed, 0x44, 0xc6,
	0x7a, 0xbd, 0xab, 0xd2, 0x7c, 0x13, 0xab, 0x6d, 0x91, 0x66, 0xe9, 0x4b, 0x52, 0xa4, 0x2a, 0x8f,
	0x75, 0x91, 0x14, 0xa5, 0xc6, 0xdb, 0x9d, 0x2a, 0x14, 0x1a, 0x99, 0x0f, 0x71, 0x22, 0x35, 0x6e,
	0x1b, 0xb8, 0x72, 0x70, 0xd3, 0x18, 0x7f, 0x5a, 0xe0, 0x82, 0x48, 0xdf, 0x64, 0xc4, 0x9f, 0x8a,
	0xdf, 0x44, 0x68, 0x5e, 0x66, 0xe3, 0x57, 0x0b, 0x9c, 0x77, 0xcc, 0xa0, 0x13, 0x30, 0x5c, 0x71,
	0xdf, 0xa3, 0x13, 0x36, 0x63, 0x74, 0x0a, 0xf7, 0xd0, 0x10, 0x1c, 0xae, 0xf8, 0x0d, 0x17, 0x77,
	0x1c, 0x5a, 0xe8, 0x18, 0x1c, 0x09, 0x2f, 0x60, 0xb7, 0xec, 0x81, 0xc2, 0x1e, 0x3a, 0x03, 0xa7,
	0x13, 0xc1, 0x43, 0xba, 0xf4, 0x99, 0xe0, 0x71, 0x2b, 0xfa, 0x08, 0x80, 0xc1, 0x52, 0x04, 0x24,
	0xa0, 0x70, 0xbf, 0x1e, 0x32, 0x1c, 0x33, 0x3e, 0xa5, 0x33, 0xc6, 0x59, 0x40, 0x17, 0xf7, 0xf0,
	0xc0, 0xfc, 0x89, 0x84, 0x84, 0x2d, 0x88, 0xbb, 0xa0, 0x70, 0xe0, 0x7e, 0x58, 0xe0, 0xea, 0x49,
	0x65, 0xb8, 0xf3, 0x92, 0xee, 0x65, 0xc7, 0xe9, 0xbd, 0x7a, 0x4b, 0x9e, 0xf5, 0xd6, 0xeb, 0xcf,
	0x09, 0x79, 0xef, 0x8d, 0xe6, 0xa6, 0x44, 0xa4, 0xc6, 0x06, 0x6b, 0x0a, 0x1d, 0x5c, 0x6f, 0x42,
	0x7f, 0xfd, 0xfa, 0x88, 0x48, 0x1d, 0xb5, 0x3e, 0x0a, 0x9d, 0xa8, 0xf1, 0xdf, 0xff, 0xf8, 0xc7,
	0x41, 0xf3, 0x28, 0xd7, 0x3f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x6b, 0x3c, 0xbd, 0x68, 0xda, 0x01,
	0x00, 0x00,
}