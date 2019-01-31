// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v0/enums/campaign_serving_status.proto

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

// Possible serving statuses of a campaign.
type CampaignServingStatusEnum_CampaignServingStatus int32

const (
	// No value has been specified.
	CampaignServingStatusEnum_UNSPECIFIED CampaignServingStatusEnum_CampaignServingStatus = 0
	// The received value is not known in this version.
	//
	// This is a response-only value.
	CampaignServingStatusEnum_UNKNOWN CampaignServingStatusEnum_CampaignServingStatus = 1
	// Serving.
	CampaignServingStatusEnum_SERVING CampaignServingStatusEnum_CampaignServingStatus = 2
	// None.
	CampaignServingStatusEnum_NONE CampaignServingStatusEnum_CampaignServingStatus = 3
	// Ended.
	CampaignServingStatusEnum_ENDED CampaignServingStatusEnum_CampaignServingStatus = 4
	// Pending.
	CampaignServingStatusEnum_PENDING CampaignServingStatusEnum_CampaignServingStatus = 5
	// Suspended.
	CampaignServingStatusEnum_SUSPENDED CampaignServingStatusEnum_CampaignServingStatus = 6
)

var CampaignServingStatusEnum_CampaignServingStatus_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "UNKNOWN",
	2: "SERVING",
	3: "NONE",
	4: "ENDED",
	5: "PENDING",
	6: "SUSPENDED",
}

var CampaignServingStatusEnum_CampaignServingStatus_value = map[string]int32{
	"UNSPECIFIED": 0,
	"UNKNOWN":     1,
	"SERVING":     2,
	"NONE":        3,
	"ENDED":       4,
	"PENDING":     5,
	"SUSPENDED":   6,
}

func (x CampaignServingStatusEnum_CampaignServingStatus) String() string {
	return proto.EnumName(CampaignServingStatusEnum_CampaignServingStatus_name, int32(x))
}

func (CampaignServingStatusEnum_CampaignServingStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_782a10ac4e85e3ac, []int{0, 0}
}

// Message describing Campaign serving statuses.
type CampaignServingStatusEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CampaignServingStatusEnum) Reset()         { *m = CampaignServingStatusEnum{} }
func (m *CampaignServingStatusEnum) String() string { return proto.CompactTextString(m) }
func (*CampaignServingStatusEnum) ProtoMessage()    {}
func (*CampaignServingStatusEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_782a10ac4e85e3ac, []int{0}
}

func (m *CampaignServingStatusEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CampaignServingStatusEnum.Unmarshal(m, b)
}
func (m *CampaignServingStatusEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CampaignServingStatusEnum.Marshal(b, m, deterministic)
}
func (m *CampaignServingStatusEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CampaignServingStatusEnum.Merge(m, src)
}
func (m *CampaignServingStatusEnum) XXX_Size() int {
	return xxx_messageInfo_CampaignServingStatusEnum.Size(m)
}
func (m *CampaignServingStatusEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_CampaignServingStatusEnum.DiscardUnknown(m)
}

var xxx_messageInfo_CampaignServingStatusEnum proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("google.ads.googleads.v0.enums.CampaignServingStatusEnum_CampaignServingStatus", CampaignServingStatusEnum_CampaignServingStatus_name, CampaignServingStatusEnum_CampaignServingStatus_value)
	proto.RegisterType((*CampaignServingStatusEnum)(nil), "google.ads.googleads.v0.enums.CampaignServingStatusEnum")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v0/enums/campaign_serving_status.proto", fileDescriptor_782a10ac4e85e3ac)
}

var fileDescriptor_782a10ac4e85e3ac = []byte{
	// 319 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x90, 0xdf, 0x6a, 0xc2, 0x30,
	0x18, 0xc5, 0xd7, 0xfa, 0x67, 0x33, 0x32, 0x56, 0x0a, 0xbb, 0xd8, 0xc0, 0x0b, 0x7d, 0x80, 0xb4,
	0xb0, 0xbb, 0x78, 0x55, 0x6d, 0x26, 0x32, 0x88, 0x65, 0xc5, 0x0e, 0x46, 0x41, 0x32, 0x5b, 0x82,
	0x60, 0x13, 0xf1, 0x53, 0x9f, 0x63, 0xcf, 0xb0, 0xcb, 0x3d, 0xca, 0x1e, 0x65, 0x37, 0x7b, 0x85,
	0x91, 0x44, 0x7b, 0xe5, 0x76, 0x13, 0x4e, 0x72, 0xce, 0xef, 0xe3, 0xcb, 0x41, 0x43, 0xa1, 0x94,
	0x58, 0x97, 0x01, 0x2f, 0x20, 0xb0, 0x52, 0xab, 0x43, 0x18, 0x94, 0x72, 0x5f, 0x41, 0xb0, 0xe4,
	0xd5, 0x86, 0xaf, 0x84, 0x5c, 0x40, 0xb9, 0x3d, 0xac, 0xa4, 0x58, 0xc0, 0x8e, 0xef, 0xf6, 0x80,
	0x37, 0x5b, 0xb5, 0x53, 0x7e, 0xcf, 0x12, 0x98, 0x17, 0x80, 0x6b, 0x18, 0x1f, 0x42, 0x6c, 0xe0,
	0xc1, 0xbb, 0x83, 0xee, 0xc6, 0xc7, 0x01, 0xa9, 0xe5, 0x53, 0x83, 0x53, 0xb9, 0xaf, 0x06, 0x80,
	0x6e, 0xcf, 0x9a, 0xfe, 0x0d, 0xea, 0xce, 0x59, 0x9a, 0xd0, 0xf1, 0xf4, 0x71, 0x4a, 0x63, 0xef,
	0xc2, 0xef, 0xa2, 0xcb, 0x39, 0x7b, 0x62, 0xb3, 0x17, 0xe6, 0x39, 0xfa, 0x92, 0xd2, 0xe7, 0x6c,
	0xca, 0x26, 0x9e, 0xeb, 0x5f, 0xa1, 0x26, 0x9b, 0x31, 0xea, 0x35, 0xfc, 0x0e, 0x6a, 0x51, 0x16,
	0xd3, 0xd8, 0x6b, 0xea, 0x44, 0x42, 0x59, 0xac, 0x13, 0x2d, 0xff, 0x1a, 0x75, 0xd2, 0x79, 0x9a,
	0x58, 0xaf, 0x3d, 0xfa, 0x71, 0x50, 0x7f, 0xa9, 0x2a, 0xfc, 0xef, 0xe2, 0xa3, 0xfb, 0xb3, 0x8b,
	0x25, 0xfa, 0xcf, 0x89, 0xf3, 0x3a, 0x3a, 0xc2, 0x42, 0xad, 0xb9, 0x14, 0x58, 0x6d, 0x45, 0x20,
	0x4a, 0x69, 0x1a, 0x39, 0x55, 0xb8, 0x59, 0xc1, 0x1f, 0x8d, 0x0e, 0xcd, 0xf9, 0xe1, 0x36, 0x26,
	0x51, 0xf4, 0xe9, 0xf6, 0x26, 0x76, 0x54, 0x54, 0x00, 0xb6, 0x52, 0xab, 0x2c, 0xc4, 0xba, 0x21,
	0xf8, 0x3a, 0xf9, 0x79, 0x54, 0x40, 0x5e, 0xfb, 0x79, 0x16, 0xe6, 0xc6, 0xff, 0x76, 0xfb, 0xf6,
	0x91, 0x90, 0xa8, 0x00, 0x42, 0xea, 0x04, 0x21, 0x59, 0x48, 0x88, 0xc9, 0xbc, 0xb5, 0xcd, 0x62,
	0x0f, 0xbf, 0x01, 0x00, 0x00, 0xff, 0xff, 0xcf, 0xd7, 0x58, 0x85, 0xe9, 0x01, 0x00, 0x00,
}
