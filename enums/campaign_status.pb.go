// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v0/enums/campaign_status.proto

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

// Possible statuses of a campaign.
type CampaignStatusEnum_CampaignStatus int32

const (
	// Not specified.
	CampaignStatusEnum_UNSPECIFIED CampaignStatusEnum_CampaignStatus = 0
	// Used for return value only. Represents value unknown in this version.
	CampaignStatusEnum_UNKNOWN CampaignStatusEnum_CampaignStatus = 1
	// Campaign is currently serving ads depending on budget information.
	CampaignStatusEnum_ENABLED CampaignStatusEnum_CampaignStatus = 2
	// Campaign has been paused by the user.
	CampaignStatusEnum_PAUSED CampaignStatusEnum_CampaignStatus = 3
	// Campaign has been removed.
	CampaignStatusEnum_REMOVED CampaignStatusEnum_CampaignStatus = 4
)

var CampaignStatusEnum_CampaignStatus_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "UNKNOWN",
	2: "ENABLED",
	3: "PAUSED",
	4: "REMOVED",
}

var CampaignStatusEnum_CampaignStatus_value = map[string]int32{
	"UNSPECIFIED": 0,
	"UNKNOWN":     1,
	"ENABLED":     2,
	"PAUSED":      3,
	"REMOVED":     4,
}

func (x CampaignStatusEnum_CampaignStatus) String() string {
	return proto.EnumName(CampaignStatusEnum_CampaignStatus_name, int32(x))
}

func (CampaignStatusEnum_CampaignStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_04d12fec335234d4, []int{0, 0}
}

// Container for enum describing possible statuses of a campaign.
type CampaignStatusEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CampaignStatusEnum) Reset()         { *m = CampaignStatusEnum{} }
func (m *CampaignStatusEnum) String() string { return proto.CompactTextString(m) }
func (*CampaignStatusEnum) ProtoMessage()    {}
func (*CampaignStatusEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_04d12fec335234d4, []int{0}
}

func (m *CampaignStatusEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CampaignStatusEnum.Unmarshal(m, b)
}
func (m *CampaignStatusEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CampaignStatusEnum.Marshal(b, m, deterministic)
}
func (m *CampaignStatusEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CampaignStatusEnum.Merge(m, src)
}
func (m *CampaignStatusEnum) XXX_Size() int {
	return xxx_messageInfo_CampaignStatusEnum.Size(m)
}
func (m *CampaignStatusEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_CampaignStatusEnum.DiscardUnknown(m)
}

var xxx_messageInfo_CampaignStatusEnum proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("google.ads.googleads.v0.enums.CampaignStatusEnum_CampaignStatus", CampaignStatusEnum_CampaignStatus_name, CampaignStatusEnum_CampaignStatus_value)
	proto.RegisterType((*CampaignStatusEnum)(nil), "google.ads.googleads.v0.enums.CampaignStatusEnum")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v0/enums/campaign_status.proto", fileDescriptor_04d12fec335234d4)
}

var fileDescriptor_04d12fec335234d4 = []byte{
	// 291 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x90, 0xd1, 0x4a, 0xc3, 0x30,
	0x14, 0x86, 0x6d, 0x27, 0x13, 0x32, 0xd0, 0x52, 0xaf, 0x77, 0xb1, 0x3d, 0x40, 0x5a, 0xd8, 0x5d,
	0xbc, 0x4a, 0xd7, 0x38, 0x86, 0xda, 0x15, 0x6b, 0x2b, 0x48, 0x41, 0xe2, 0x5a, 0xc2, 0x64, 0x6d,
	0xca, 0x4e, 0xbb, 0x07, 0xf2, 0xd2, 0x47, 0xf1, 0x49, 0xc4, 0xa7, 0x90, 0x24, 0x5b, 0x61, 0x17,
	0x7a, 0x13, 0xfe, 0x9c, 0xff, 0x7c, 0xc9, 0xf9, 0x0f, 0x9a, 0x09, 0x29, 0xc5, 0xb6, 0xf4, 0x78,
	0x01, 0x9e, 0x91, 0x4a, 0xed, 0x7d, 0xaf, 0xac, 0xbb, 0x0a, 0xbc, 0x35, 0xaf, 0x1a, 0xbe, 0x11,
	0xf5, 0x2b, 0xb4, 0xbc, 0xed, 0x00, 0x37, 0x3b, 0xd9, 0x4a, 0x77, 0x6c, 0x3a, 0x31, 0x2f, 0x00,
	0xf7, 0x10, 0xde, 0xfb, 0x58, 0x43, 0xd3, 0x77, 0xe4, 0xce, 0x0f, 0x5c, 0xa2, 0x31, 0x56, 0x77,
	0xd5, 0xf4, 0x09, 0x5d, 0x9e, 0x56, 0xdd, 0x2b, 0x34, 0x4a, 0xa3, 0x24, 0x66, 0xf3, 0xe5, 0xed,
	0x92, 0x85, 0xce, 0x99, 0x3b, 0x42, 0x17, 0x69, 0x74, 0x17, 0xad, 0x9e, 0x23, 0xc7, 0x52, 0x17,
	0x16, 0xd1, 0xe0, 0x9e, 0x85, 0x8e, 0xed, 0x22, 0x34, 0x8c, 0x69, 0x9a, 0xb0, 0xd0, 0x19, 0x28,
	0xe3, 0x91, 0x3d, 0xac, 0x32, 0x16, 0x3a, 0xe7, 0xc1, 0xb7, 0x85, 0x26, 0x6b, 0x59, 0xe1, 0x7f,
	0x27, 0x0a, 0xae, 0x4f, 0x7f, 0x8e, 0x55, 0x8a, 0xd8, 0x7a, 0x09, 0x0e, 0x94, 0x90, 0x5b, 0x5e,
	0x0b, 0x2c, 0x77, 0xc2, 0x13, 0x65, 0xad, 0x33, 0x1e, 0x97, 0xd1, 0x6c, 0xe0, 0x8f, 0xdd, 0xdc,
	0xe8, 0xf3, 0xc3, 0x1e, 0x2c, 0x28, 0xfd, 0xb4, 0xc7, 0x0b, 0xf3, 0x14, 0x2d, 0x00, 0x1b, 0xa9,
	0x54, 0xe6, 0x63, 0x95, 0x1d, 0xbe, 0x8e, 0x7e, 0x4e, 0x0b, 0xc8, 0x7b, 0x3f, 0xcf, 0xfc, 0x5c,
	0xfb, 0x3f, 0xf6, 0xc4, 0x14, 0x09, 0xa1, 0x05, 0x10, 0xd2, 0x77, 0x10, 0x92, 0xf9, 0x84, 0xe8,
	0x9e, 0xb7, 0xa1, 0x1e, 0x6c, 0xf6, 0x1b, 0x00, 0x00, 0xff, 0xff, 0xde, 0x45, 0x74, 0x45, 0xb3,
	0x01, 0x00, 0x00,
}
