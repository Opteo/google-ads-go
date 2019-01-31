// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v0/enums/feed_link_status.proto

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

// Possible statuses of a feed link.
type FeedLinkStatusEnum_FeedLinkStatus int32

const (
	// Not specified.
	FeedLinkStatusEnum_UNSPECIFIED FeedLinkStatusEnum_FeedLinkStatus = 0
	// Used for return value only. Represents value unknown in this version.
	FeedLinkStatusEnum_UNKNOWN FeedLinkStatusEnum_FeedLinkStatus = 1
	// Feed link is enabled.
	FeedLinkStatusEnum_ENABLED FeedLinkStatusEnum_FeedLinkStatus = 2
	// Feed link has been removed.
	FeedLinkStatusEnum_REMOVED FeedLinkStatusEnum_FeedLinkStatus = 3
)

var FeedLinkStatusEnum_FeedLinkStatus_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "UNKNOWN",
	2: "ENABLED",
	3: "REMOVED",
}

var FeedLinkStatusEnum_FeedLinkStatus_value = map[string]int32{
	"UNSPECIFIED": 0,
	"UNKNOWN":     1,
	"ENABLED":     2,
	"REMOVED":     3,
}

func (x FeedLinkStatusEnum_FeedLinkStatus) String() string {
	return proto.EnumName(FeedLinkStatusEnum_FeedLinkStatus_name, int32(x))
}

func (FeedLinkStatusEnum_FeedLinkStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_220be290d1e4ef70, []int{0, 0}
}

// Container for an enum describing possible statuses of a feed link.
type FeedLinkStatusEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FeedLinkStatusEnum) Reset()         { *m = FeedLinkStatusEnum{} }
func (m *FeedLinkStatusEnum) String() string { return proto.CompactTextString(m) }
func (*FeedLinkStatusEnum) ProtoMessage()    {}
func (*FeedLinkStatusEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_220be290d1e4ef70, []int{0}
}

func (m *FeedLinkStatusEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FeedLinkStatusEnum.Unmarshal(m, b)
}
func (m *FeedLinkStatusEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FeedLinkStatusEnum.Marshal(b, m, deterministic)
}
func (m *FeedLinkStatusEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FeedLinkStatusEnum.Merge(m, src)
}
func (m *FeedLinkStatusEnum) XXX_Size() int {
	return xxx_messageInfo_FeedLinkStatusEnum.Size(m)
}
func (m *FeedLinkStatusEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_FeedLinkStatusEnum.DiscardUnknown(m)
}

var xxx_messageInfo_FeedLinkStatusEnum proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("google.ads.googleads.v0.enums.FeedLinkStatusEnum_FeedLinkStatus", FeedLinkStatusEnum_FeedLinkStatus_name, FeedLinkStatusEnum_FeedLinkStatus_value)
	proto.RegisterType((*FeedLinkStatusEnum)(nil), "google.ads.googleads.v0.enums.FeedLinkStatusEnum")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v0/enums/feed_link_status.proto", fileDescriptor_220be290d1e4ef70)
}

var fileDescriptor_220be290d1e4ef70 = []byte{
	// 286 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x32, 0x49, 0xcf, 0xcf, 0x4f,
	0xcf, 0x49, 0xd5, 0x4f, 0x4c, 0x29, 0xd6, 0x87, 0x30, 0x41, 0xac, 0x32, 0x03, 0xfd, 0xd4, 0xbc,
	0xd2, 0xdc, 0x62, 0xfd, 0xb4, 0xd4, 0xd4, 0x94, 0xf8, 0x9c, 0xcc, 0xbc, 0xec, 0xf8, 0xe2, 0x92,
	0xc4, 0x92, 0xd2, 0x62, 0xbd, 0x82, 0xa2, 0xfc, 0x92, 0x7c, 0x21, 0x59, 0x88, 0x52, 0xbd, 0xc4,
	0x94, 0x62, 0x3d, 0xb8, 0x2e, 0xbd, 0x32, 0x03, 0x3d, 0xb0, 0x2e, 0xa5, 0x38, 0x2e, 0x21, 0xb7,
	0xd4, 0xd4, 0x14, 0x9f, 0xcc, 0xbc, 0xec, 0x60, 0xb0, 0x36, 0xd7, 0xbc, 0xd2, 0x5c, 0x25, 0x0f,
	0x2e, 0x3e, 0x54, 0x51, 0x21, 0x7e, 0x2e, 0xee, 0x50, 0xbf, 0xe0, 0x00, 0x57, 0x67, 0x4f, 0x37,
	0x4f, 0x57, 0x17, 0x01, 0x06, 0x21, 0x6e, 0x2e, 0xf6, 0x50, 0x3f, 0x6f, 0x3f, 0xff, 0x70, 0x3f,
	0x01, 0x46, 0x10, 0xc7, 0xd5, 0xcf, 0xd1, 0xc9, 0xc7, 0xd5, 0x45, 0x80, 0x09, 0xc4, 0x09, 0x72,
	0xf5, 0xf5, 0x0f, 0x73, 0x75, 0x11, 0x60, 0x76, 0x7a, 0xc1, 0xc8, 0xa5, 0x98, 0x9c, 0x9f, 0xab,
	0x87, 0xd7, 0x15, 0x4e, 0xc2, 0xa8, 0xb6, 0x05, 0x80, 0x5c, 0x1e, 0xc0, 0x18, 0xe5, 0x04, 0xd5,
	0x95, 0x9e, 0x9f, 0x93, 0x98, 0x97, 0xae, 0x97, 0x5f, 0x94, 0xae, 0x9f, 0x9e, 0x9a, 0x07, 0xf6,
	0x17, 0x2c, 0x04, 0x0a, 0x32, 0x8b, 0x71, 0x04, 0x88, 0x35, 0x98, 0x5c, 0xc4, 0xc4, 0xec, 0xee,
	0xe8, 0xb8, 0x8a, 0x49, 0xd6, 0x1d, 0x62, 0x94, 0x63, 0x4a, 0xb1, 0x1e, 0x84, 0x09, 0x62, 0x85,
	0x19, 0xe8, 0x81, 0xfc, 0x5b, 0x7c, 0x0a, 0x26, 0x1f, 0xe3, 0x98, 0x52, 0x1c, 0x03, 0x97, 0x8f,
	0x09, 0x33, 0x88, 0x01, 0xcb, 0xbf, 0x62, 0x52, 0x84, 0x08, 0x5a, 0x59, 0x39, 0xa6, 0x14, 0x5b,
	0x59, 0xc1, 0x55, 0x58, 0x59, 0x85, 0x19, 0x58, 0x59, 0x81, 0xd5, 0x24, 0xb1, 0x81, 0x1d, 0x66,
	0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0x5f, 0x28, 0x89, 0xc3, 0xa8, 0x01, 0x00, 0x00,
}
