// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v0/resources/ad_group_feed.proto

package resources

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	common "github.com/kritzware/google-ads-go/common"
	enums "github.com/kritzware/google-ads-go/enums"
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

// An ad group feed.
type AdGroupFeed struct {
	// The resource name of the ad group feed.
	// Ad group feed resource names have the form:
	//
	// `customers/{customer_id}/adGroupFeeds/{ad_group_id}_{feed_id}
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// The feed being linked to the ad group.
	Feed *wrappers.StringValue `protobuf:"bytes,2,opt,name=feed,proto3" json:"feed,omitempty"`
	// The ad group being linked to the feed.
	AdGroup *wrappers.StringValue `protobuf:"bytes,3,opt,name=ad_group,json=adGroup,proto3" json:"ad_group,omitempty"`
	// Indicates which placeholder types the feed may populate under the connected
	// ad group. Required.
	PlaceholderTypes []enums.PlaceholderTypeEnum_PlaceholderType `protobuf:"varint,4,rep,packed,name=placeholder_types,json=placeholderTypes,proto3,enum=google.ads.googleads.v0.enums.PlaceholderTypeEnum_PlaceholderType" json:"placeholder_types,omitempty"`
	// Matching function associated with the AdGroupFeed.
	// The matching function is used to filter the set of feed items selected.
	// Required.
	MatchingFunction *common.MatchingFunction `protobuf:"bytes,5,opt,name=matching_function,json=matchingFunction,proto3" json:"matching_function,omitempty"`
	// Status of the ad group feed.
	// This field is read-only.
	Status               enums.FeedLinkStatusEnum_FeedLinkStatus `protobuf:"varint,6,opt,name=status,proto3,enum=google.ads.googleads.v0.enums.FeedLinkStatusEnum_FeedLinkStatus" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                `json:"-"`
	XXX_unrecognized     []byte                                  `json:"-"`
	XXX_sizecache        int32                                   `json:"-"`
}

func (m *AdGroupFeed) Reset()         { *m = AdGroupFeed{} }
func (m *AdGroupFeed) String() string { return proto.CompactTextString(m) }
func (*AdGroupFeed) ProtoMessage()    {}
func (*AdGroupFeed) Descriptor() ([]byte, []int) {
	return fileDescriptor_1f048fb4293eaa8a, []int{0}
}

func (m *AdGroupFeed) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AdGroupFeed.Unmarshal(m, b)
}
func (m *AdGroupFeed) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AdGroupFeed.Marshal(b, m, deterministic)
}
func (m *AdGroupFeed) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AdGroupFeed.Merge(m, src)
}
func (m *AdGroupFeed) XXX_Size() int {
	return xxx_messageInfo_AdGroupFeed.Size(m)
}
func (m *AdGroupFeed) XXX_DiscardUnknown() {
	xxx_messageInfo_AdGroupFeed.DiscardUnknown(m)
}

var xxx_messageInfo_AdGroupFeed proto.InternalMessageInfo

func (m *AdGroupFeed) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func (m *AdGroupFeed) GetFeed() *wrappers.StringValue {
	if m != nil {
		return m.Feed
	}
	return nil
}

func (m *AdGroupFeed) GetAdGroup() *wrappers.StringValue {
	if m != nil {
		return m.AdGroup
	}
	return nil
}

func (m *AdGroupFeed) GetPlaceholderTypes() []enums.PlaceholderTypeEnum_PlaceholderType {
	if m != nil {
		return m.PlaceholderTypes
	}
	return nil
}

func (m *AdGroupFeed) GetMatchingFunction() *common.MatchingFunction {
	if m != nil {
		return m.MatchingFunction
	}
	return nil
}

func (m *AdGroupFeed) GetStatus() enums.FeedLinkStatusEnum_FeedLinkStatus {
	if m != nil {
		return m.Status
	}
	return enums.FeedLinkStatusEnum_UNSPECIFIED
}

func init() {
	proto.RegisterType((*AdGroupFeed)(nil), "google.ads.googleads.v0.resources.AdGroupFeed")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v0/resources/ad_group_feed.proto", fileDescriptor_1f048fb4293eaa8a)
}

var fileDescriptor_1f048fb4293eaa8a = []byte{
	// 474 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x93, 0xdd, 0x6a, 0xd4, 0x40,
	0x14, 0xc7, 0xc9, 0x6e, 0x5d, 0x75, 0xaa, 0x65, 0x9b, 0xab, 0x50, 0x44, 0xb6, 0x4a, 0x61, 0xaf,
	0x26, 0x61, 0xfd, 0x82, 0x78, 0x63, 0x16, 0xec, 0x82, 0xa8, 0x2c, 0xa9, 0x2c, 0x22, 0x2b, 0x61,
	0x9a, 0x39, 0x3b, 0x0d, 0xcd, 0x7c, 0x30, 0x93, 0x54, 0xfa, 0x3a, 0x5e, 0xfa, 0x22, 0x82, 0x8f,
	0xe2, 0x3b, 0x08, 0x92, 0x4c, 0x12, 0xed, 0x4a, 0x6c, 0xef, 0x4e, 0x4e, 0xfe, 0xbf, 0x33, 0xe7,
	0x9c, 0xff, 0x0c, 0x7a, 0xc6, 0xa4, 0x64, 0x39, 0xf8, 0x84, 0x1a, 0xdf, 0x86, 0x55, 0x74, 0x11,
	0xf8, 0x1a, 0x8c, 0x2c, 0x75, 0x0a, 0xc6, 0x27, 0x34, 0x61, 0x5a, 0x96, 0x2a, 0xd9, 0x00, 0x50,
	0xac, 0xb4, 0x2c, 0xa4, 0x7b, 0x68, 0xb5, 0x98, 0x50, 0x83, 0x3b, 0x0c, 0x5f, 0x04, 0xb8, 0xc3,
	0x0e, 0x9e, 0xf7, 0x55, 0x4e, 0x25, 0xe7, 0x52, 0xf8, 0x9c, 0x14, 0xe9, 0x59, 0x26, 0x58, 0xb2,
	0x29, 0x45, 0x5a, 0x64, 0x52, 0xd8, 0xd2, 0x07, 0x4f, 0xfb, 0x38, 0x10, 0x25, 0x37, 0x7e, 0xd5,
	0x44, 0x92, 0x67, 0xe2, 0x3c, 0x31, 0x05, 0x29, 0x4a, 0x73, 0x33, 0x4a, 0xe5, 0x24, 0x85, 0x33,
	0x99, 0x53, 0xd0, 0x49, 0x71, 0xa9, 0xa0, 0xa1, 0x1e, 0x36, 0x54, 0xfd, 0x75, 0x5a, 0x6e, 0xfc,
	0x2f, 0x9a, 0x28, 0x05, 0xba, 0xa9, 0xfa, 0xe8, 0xfb, 0x10, 0xed, 0x46, 0x74, 0x51, 0x4d, 0x7f,
	0x0c, 0x40, 0xdd, 0xc7, 0xe8, 0x7e, 0x3b, 0x60, 0x22, 0x08, 0x07, 0xcf, 0x99, 0x38, 0xd3, 0xbb,
	0xf1, 0xbd, 0x36, 0xf9, 0x9e, 0x70, 0x70, 0x03, 0xb4, 0x53, 0x35, 0xe9, 0x0d, 0x26, 0xce, 0x74,
	0x77, 0xf6, 0xa0, 0xd9, 0x0f, 0x6e, 0xcf, 0xc0, 0x27, 0x85, 0xce, 0x04, 0x5b, 0x91, 0xbc, 0x84,
	0xb8, 0x56, 0xba, 0x2f, 0xd0, 0x9d, 0x76, 0xc9, 0xde, 0xf0, 0x06, 0xd4, 0x6d, 0x62, 0x7b, 0x72,
	0x25, 0xda, 0xdf, 0x9e, 0xcc, 0x78, 0x3b, 0x93, 0xe1, 0x74, 0x6f, 0x36, 0xc7, 0x7d, 0x16, 0xd5,
	0x1b, 0xc1, 0xcb, 0x3f, 0xdc, 0x87, 0x4b, 0x05, 0xaf, 0x45, 0xc9, 0xb7, 0x73, 0xf1, 0x58, 0x5d,
	0x4d, 0x18, 0xf7, 0x33, 0xda, 0xff, 0xc7, 0x37, 0xef, 0x56, 0xdd, 0x72, 0xd0, 0x7b, 0xa0, 0x35,
	0x1c, 0xbf, 0x6b, 0xc0, 0xe3, 0x86, 0x8b, 0xc7, 0x7c, 0x2b, 0xe3, 0x7e, 0x44, 0x23, 0xeb, 0xaa,
	0x37, 0x9a, 0x38, 0xd3, 0xbd, 0xd9, 0xab, 0x6b, 0x86, 0xa8, 0x4c, 0x79, 0x9b, 0x89, 0xf3, 0x93,
	0x1a, 0xaa, 0x67, 0xb8, 0x9a, 0x8a, 0x9b, 0x7a, 0xf3, 0x5f, 0x0e, 0x3a, 0x4a, 0x25, 0xc7, 0xd7,
	0xde, 0xdb, 0xf9, 0xf8, 0x2f, 0xc3, 0x97, 0xd5, 0xf6, 0x97, 0xce, 0xa7, 0x37, 0x0d, 0xc6, 0x64,
	0x4e, 0x04, 0xc3, 0x52, 0x33, 0x9f, 0x81, 0xa8, 0xbd, 0x69, 0x6f, 0x9b, 0xca, 0xcc, 0x7f, 0x1e,
	0xd1, 0xcb, 0x2e, 0xfa, 0x3a, 0x18, 0x2e, 0xa2, 0xe8, 0xdb, 0xe0, 0x70, 0x61, 0x4b, 0x46, 0xd4,
	0x60, 0x1b, 0x56, 0xd1, 0x2a, 0xc0, 0x71, 0xab, 0xfc, 0xd1, 0x6a, 0xd6, 0x11, 0x35, 0xeb, 0x4e,
	0xb3, 0x5e, 0x05, 0xeb, 0x4e, 0xf3, 0x73, 0x70, 0x64, 0x7f, 0x84, 0x61, 0x44, 0x4d, 0x18, 0x76,
	0xaa, 0x30, 0x5c, 0x05, 0x61, 0xd8, 0xe9, 0x4e, 0x47, 0x75, 0xb3, 0x4f, 0x7e, 0x07, 0x00, 0x00,
	0xff, 0xff, 0x69, 0xbe, 0xb6, 0x73, 0xf0, 0x03, 0x00, 0x00,
}
