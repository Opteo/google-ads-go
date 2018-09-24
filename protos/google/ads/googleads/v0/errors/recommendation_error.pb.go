// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v0/errors/recommendation_error.proto

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

// Enum describing possible errors from applying a recommendation.
type RecommendationErrorEnum_RecommendationError int32

const (
	// Enum unspecified.
	RecommendationErrorEnum_UNSPECIFIED RecommendationErrorEnum_RecommendationError = 0
	// The received error code is not known in this version.
	RecommendationErrorEnum_UNKNOWN RecommendationErrorEnum_RecommendationError = 1
	// The specified budget amount is too low e.g. lower than minimum currency
	// unit or lower than ad group minimum cost-per-click.
	RecommendationErrorEnum_BUDGET_AMOUNT_TOO_SMALL RecommendationErrorEnum_RecommendationError = 2
	// The specified budget amount is too large.
	RecommendationErrorEnum_BUDGET_AMOUNT_TOO_LARGE RecommendationErrorEnum_RecommendationError = 3
	// The specified budget amount is not a valid amount. e.g. not a multiple
	// of minimum currency unit.
	RecommendationErrorEnum_INVALID_BUDGET_AMOUNT RecommendationErrorEnum_RecommendationError = 4
	// The specified keyword or ad violates ad policy.
	RecommendationErrorEnum_POLICY_ERROR RecommendationErrorEnum_RecommendationError = 5
	// The specified bid amount is not valid. e.g. too many fractional digits,
	// or negative amount.
	RecommendationErrorEnum_INVALID_BID_AMOUNT RecommendationErrorEnum_RecommendationError = 6
	// The number of keywords in ad group have reached the maximum allowed.
	RecommendationErrorEnum_ADGROUP_KEYWORD_LIMIT RecommendationErrorEnum_RecommendationError = 7
)

var RecommendationErrorEnum_RecommendationError_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "UNKNOWN",
	2: "BUDGET_AMOUNT_TOO_SMALL",
	3: "BUDGET_AMOUNT_TOO_LARGE",
	4: "INVALID_BUDGET_AMOUNT",
	5: "POLICY_ERROR",
	6: "INVALID_BID_AMOUNT",
	7: "ADGROUP_KEYWORD_LIMIT",
}

var RecommendationErrorEnum_RecommendationError_value = map[string]int32{
	"UNSPECIFIED":             0,
	"UNKNOWN":                 1,
	"BUDGET_AMOUNT_TOO_SMALL": 2,
	"BUDGET_AMOUNT_TOO_LARGE": 3,
	"INVALID_BUDGET_AMOUNT":   4,
	"POLICY_ERROR":            5,
	"INVALID_BID_AMOUNT":      6,
	"ADGROUP_KEYWORD_LIMIT":   7,
}

func (x RecommendationErrorEnum_RecommendationError) String() string {
	return proto.EnumName(RecommendationErrorEnum_RecommendationError_name, int32(x))
}

func (RecommendationErrorEnum_RecommendationError) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_d96255e6afd54850, []int{0, 0}
}

// Container for enum describing possible errors from applying a recommendation.
type RecommendationErrorEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RecommendationErrorEnum) Reset()         { *m = RecommendationErrorEnum{} }
func (m *RecommendationErrorEnum) String() string { return proto.CompactTextString(m) }
func (*RecommendationErrorEnum) ProtoMessage()    {}
func (*RecommendationErrorEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_d96255e6afd54850, []int{0}
}
func (m *RecommendationErrorEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RecommendationErrorEnum.Unmarshal(m, b)
}
func (m *RecommendationErrorEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RecommendationErrorEnum.Marshal(b, m, deterministic)
}
func (m *RecommendationErrorEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RecommendationErrorEnum.Merge(m, src)
}
func (m *RecommendationErrorEnum) XXX_Size() int {
	return xxx_messageInfo_RecommendationErrorEnum.Size(m)
}
func (m *RecommendationErrorEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_RecommendationErrorEnum.DiscardUnknown(m)
}

var xxx_messageInfo_RecommendationErrorEnum proto.InternalMessageInfo

func init() {
	proto.RegisterType((*RecommendationErrorEnum)(nil), "google.ads.googleads.v0.errors.RecommendationErrorEnum")
	proto.RegisterEnum("google.ads.googleads.v0.errors.RecommendationErrorEnum_RecommendationError", RecommendationErrorEnum_RecommendationError_name, RecommendationErrorEnum_RecommendationError_value)
}

func init() {
	proto.RegisterFile("google/ads/googleads/v0/errors/recommendation_error.proto", fileDescriptor_d96255e6afd54850)
}

var fileDescriptor_d96255e6afd54850 = []byte{
	// 318 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x91, 0xd1, 0x4a, 0xf3, 0x30,
	0x18, 0x86, 0xff, 0x6e, 0xbf, 0x1b, 0x64, 0x82, 0x21, 0xa2, 0x53, 0x84, 0x1d, 0xec, 0x02, 0xd2,
	0x82, 0x47, 0x1e, 0x66, 0x6b, 0x2c, 0x61, 0x5d, 0x52, 0xb2, 0x76, 0x63, 0x50, 0x08, 0x73, 0x2d,
	0x43, 0xb0, 0x8b, 0x34, 0x73, 0x17, 0xe4, 0x99, 0x5e, 0x8a, 0xe0, 0x4d, 0x78, 0xe6, 0x5d, 0x48,
	0x17, 0x36, 0x18, 0x14, 0x3d, 0x7b, 0xc8, 0xfb, 0xbc, 0x1f, 0xe4, 0xfb, 0xc0, 0xdd, 0x4a, 0xeb,
	0xd5, 0x53, 0xee, 0x2e, 0x32, 0xe3, 0x5a, 0xac, 0x68, 0xeb, 0xb9, 0x79, 0x59, 0xea, 0xd2, 0xb8,
	0x65, 0xbe, 0xd4, 0x45, 0x91, 0xaf, 0xb3, 0xc5, 0xe6, 0x51, 0xaf, 0xd5, 0xee, 0x15, 0x3f, 0x97,
	0x7a, 0xa3, 0x51, 0xcf, 0xfa, 0x78, 0x91, 0x19, 0x7c, 0xa8, 0xe2, 0xad, 0x87, 0x6d, 0xb5, 0xff,
	0xed, 0x80, 0xae, 0x3c, 0xaa, 0xd3, 0x2a, 0xa0, 0xeb, 0x97, 0xa2, 0xff, 0xe9, 0x80, 0xf3, 0x9a,
	0x0c, 0x9d, 0x81, 0x4e, 0xc2, 0x27, 0x11, 0x1d, 0xb2, 0x7b, 0x46, 0x7d, 0xf8, 0x0f, 0x75, 0x40,
	0x3b, 0xe1, 0x23, 0x2e, 0x66, 0x1c, 0x3a, 0xe8, 0x06, 0x74, 0x07, 0x89, 0x1f, 0xd0, 0x58, 0x91,
	0xb1, 0x48, 0x78, 0xac, 0x62, 0x21, 0xd4, 0x64, 0x4c, 0xc2, 0x10, 0x36, 0xea, 0xc3, 0x90, 0xc8,
	0x80, 0xc2, 0x26, 0xba, 0x06, 0x17, 0x8c, 0x4f, 0x49, 0xc8, 0x7c, 0x75, 0x24, 0xc1, 0xff, 0x08,
	0x82, 0xd3, 0x48, 0x84, 0x6c, 0x38, 0x57, 0x54, 0x4a, 0x21, 0xe1, 0x09, 0xba, 0x04, 0xe8, 0x20,
	0x33, 0x7f, 0x6f, 0xb6, 0xaa, 0x21, 0xc4, 0x0f, 0xa4, 0x48, 0x22, 0x35, 0xa2, 0xf3, 0x99, 0x90,
	0xbe, 0x0a, 0xd9, 0x98, 0xc5, 0xb0, 0x3d, 0x78, 0x73, 0x40, 0x7f, 0xa9, 0x0b, 0xfc, 0xfb, 0x4a,
	0x06, 0x57, 0x35, 0x7f, 0x8e, 0xaa, 0x65, 0x46, 0xce, 0x6b, 0xa3, 0x19, 0x10, 0xf2, 0xde, 0xe8,
	0x05, 0x76, 0x04, 0xc9, 0x0c, 0xb6, 0x58, 0xd1, 0xd4, 0xc3, 0x3b, 0xd9, 0x7c, 0xec, 0x85, 0x94,
	0x64, 0x26, 0x3d, 0x08, 0xe9, 0xd4, 0x4b, 0xad, 0xf0, 0xf5, 0x97, 0xf0, 0xd0, 0xda, 0x9d, 0xef,
	0xf6, 0x27, 0x00, 0x00, 0xff, 0xff, 0xd0, 0x8c, 0xee, 0xf2, 0xfb, 0x01, 0x00, 0x00,
}