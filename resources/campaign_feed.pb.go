// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v0/resources/campaign_feed.proto

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

// A campaign feed.
type CampaignFeed struct {
	// The resource name of the campaign feed.
	// Campaign feed resource names have the form:
	//
	// `customers/{customer_id}/campaignFeeds/{campaign_id}_{feed_id}
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// The feed to which the CampaignFeed belongs.
	Feed *wrappers.StringValue `protobuf:"bytes,2,opt,name=feed,proto3" json:"feed,omitempty"`
	// The campaign to which the CampaignFeed belongs.
	Campaign *wrappers.StringValue `protobuf:"bytes,3,opt,name=campaign,proto3" json:"campaign,omitempty"`
	// Indicates which placeholder types the feed may populate under the connected
	// campaign. Required.
	PlaceholderTypes []enums.PlaceholderTypeEnum_PlaceholderType `protobuf:"varint,4,rep,packed,name=placeholder_types,json=placeholderTypes,proto3,enum=google.ads.googleads.v0.enums.PlaceholderTypeEnum_PlaceholderType" json:"placeholder_types,omitempty"`
	// Matching function associated with the CampaignFeed.
	// The matching function is used to filter the set of feed items selected.
	// Required.
	MatchingFunction *common.MatchingFunction `protobuf:"bytes,5,opt,name=matching_function,json=matchingFunction,proto3" json:"matching_function,omitempty"`
	// Status of the campaign feed.
	// This field is read-only.
	Status               enums.FeedLinkStatusEnum_FeedLinkStatus `protobuf:"varint,6,opt,name=status,proto3,enum=google.ads.googleads.v0.enums.FeedLinkStatusEnum_FeedLinkStatus" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                `json:"-"`
	XXX_unrecognized     []byte                                  `json:"-"`
	XXX_sizecache        int32                                   `json:"-"`
}

func (m *CampaignFeed) Reset()         { *m = CampaignFeed{} }
func (m *CampaignFeed) String() string { return proto.CompactTextString(m) }
func (*CampaignFeed) ProtoMessage()    {}
func (*CampaignFeed) Descriptor() ([]byte, []int) {
	return fileDescriptor_18c98ca1b87e62d4, []int{0}
}

func (m *CampaignFeed) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CampaignFeed.Unmarshal(m, b)
}
func (m *CampaignFeed) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CampaignFeed.Marshal(b, m, deterministic)
}
func (m *CampaignFeed) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CampaignFeed.Merge(m, src)
}
func (m *CampaignFeed) XXX_Size() int {
	return xxx_messageInfo_CampaignFeed.Size(m)
}
func (m *CampaignFeed) XXX_DiscardUnknown() {
	xxx_messageInfo_CampaignFeed.DiscardUnknown(m)
}

var xxx_messageInfo_CampaignFeed proto.InternalMessageInfo

func (m *CampaignFeed) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func (m *CampaignFeed) GetFeed() *wrappers.StringValue {
	if m != nil {
		return m.Feed
	}
	return nil
}

func (m *CampaignFeed) GetCampaign() *wrappers.StringValue {
	if m != nil {
		return m.Campaign
	}
	return nil
}

func (m *CampaignFeed) GetPlaceholderTypes() []enums.PlaceholderTypeEnum_PlaceholderType {
	if m != nil {
		return m.PlaceholderTypes
	}
	return nil
}

func (m *CampaignFeed) GetMatchingFunction() *common.MatchingFunction {
	if m != nil {
		return m.MatchingFunction
	}
	return nil
}

func (m *CampaignFeed) GetStatus() enums.FeedLinkStatusEnum_FeedLinkStatus {
	if m != nil {
		return m.Status
	}
	return enums.FeedLinkStatusEnum_UNSPECIFIED
}

func init() {
	proto.RegisterType((*CampaignFeed)(nil), "google.ads.googleads.v0.resources.CampaignFeed")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v0/resources/campaign_feed.proto", fileDescriptor_18c98ca1b87e62d4)
}

var fileDescriptor_18c98ca1b87e62d4 = []byte{
	// 470 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x93, 0xdd, 0x6a, 0xd4, 0x40,
	0x14, 0xc7, 0xc9, 0x6e, 0x5d, 0x74, 0xac, 0xa5, 0x9b, 0xab, 0x50, 0x44, 0xb6, 0x4a, 0x61, 0xaf,
	0x26, 0x61, 0xfd, 0x40, 0xe2, 0x8d, 0x59, 0xb1, 0x05, 0x51, 0x59, 0x52, 0x59, 0x44, 0x56, 0xc2,
	0x34, 0x39, 0x3b, 0x0d, 0xcd, 0x7c, 0x90, 0x49, 0x2a, 0x7d, 0x1d, 0x2f, 0x7d, 0x13, 0x7d, 0x14,
	0x1f, 0x42, 0x24, 0x99, 0xc9, 0x68, 0x57, 0x62, 0xf7, 0xee, 0xe4, 0xe4, 0xff, 0x3b, 0x73, 0xce,
	0xf9, 0xcf, 0xa0, 0xa7, 0x54, 0x08, 0x5a, 0x80, 0x4f, 0x32, 0xe5, 0xeb, 0xb0, 0x89, 0x2e, 0x03,
	0xbf, 0x04, 0x25, 0xea, 0x32, 0x05, 0xe5, 0xa7, 0x84, 0x49, 0x92, 0x53, 0x9e, 0xac, 0x01, 0x32,
	0x2c, 0x4b, 0x51, 0x09, 0xf7, 0x50, 0x6b, 0x31, 0xc9, 0x14, 0xb6, 0x18, 0xbe, 0x0c, 0xb0, 0xc5,
	0x0e, 0x9e, 0xf5, 0x55, 0x4e, 0x05, 0x63, 0x82, 0xfb, 0x8c, 0x54, 0xe9, 0x79, 0xce, 0x69, 0xb2,
	0xae, 0x79, 0x5a, 0xe5, 0x82, 0xeb, 0xd2, 0x07, 0x4f, 0xfa, 0x38, 0xe0, 0x35, 0x53, 0x7e, 0xd3,
	0x44, 0x52, 0xe4, 0xfc, 0x22, 0x51, 0x15, 0xa9, 0x6a, 0xb5, 0x1d, 0x25, 0x0b, 0x92, 0xc2, 0xb9,
	0x28, 0x32, 0x28, 0x93, 0xea, 0x4a, 0x82, 0xa1, 0x1e, 0x18, 0xaa, 0xfd, 0x3a, 0xab, 0xd7, 0xfe,
	0x97, 0x92, 0x48, 0x09, 0xa5, 0xa9, 0xfa, 0xf0, 0xfb, 0x10, 0xed, 0xbe, 0x32, 0xe3, 0x1f, 0x03,
	0x64, 0xee, 0x23, 0x74, 0xaf, 0x9b, 0x30, 0xe1, 0x84, 0x81, 0xe7, 0x4c, 0x9c, 0xe9, 0x9d, 0x78,
	0xb7, 0x4b, 0xbe, 0x27, 0x0c, 0xdc, 0x00, 0xed, 0x34, 0x5d, 0x7a, 0x83, 0x89, 0x33, 0xbd, 0x3b,
	0xbb, 0x6f, 0x16, 0x84, 0xbb, 0x43, 0xf0, 0x69, 0x55, 0xe6, 0x9c, 0x2e, 0x49, 0x51, 0x43, 0xdc,
	0x2a, 0xdd, 0xe7, 0xe8, 0x76, 0xb7, 0x65, 0x6f, 0xb8, 0x05, 0x65, 0xd5, 0xae, 0x40, 0xe3, 0xcd,
	0xd9, 0x94, 0xb7, 0x33, 0x19, 0x4e, 0xf7, 0x66, 0x73, 0xdc, 0x67, 0x52, 0xbb, 0x13, 0xbc, 0xf8,
	0xc3, 0x7d, 0xb8, 0x92, 0xf0, 0x9a, 0xd7, 0x6c, 0x33, 0x17, 0xef, 0xcb, 0xeb, 0x09, 0xe5, 0x7e,
	0x46, 0xe3, 0x7f, 0x9c, 0xf3, 0x6e, 0xb5, 0x3d, 0x07, 0xbd, 0x07, 0x6a, 0xcb, 0xf1, 0x3b, 0x03,
	0x1e, 0x1b, 0x2e, 0xde, 0x67, 0x1b, 0x19, 0xf7, 0x23, 0x1a, 0x69, 0x5f, 0xbd, 0xd1, 0xc4, 0x99,
	0xee, 0xcd, 0x5e, 0xde, 0x30, 0x44, 0xe3, 0xca, 0xdb, 0x9c, 0x5f, 0x9c, 0xb6, 0x50, 0x3b, 0xc3,
	0xf5, 0x54, 0x6c, 0xea, 0xcd, 0x7f, 0x39, 0xe8, 0x28, 0x15, 0x0c, 0xdf, 0x78, 0x73, 0xe7, 0xe3,
	0xbf, 0x2d, 0x5f, 0x34, 0xfb, 0x5f, 0x38, 0x9f, 0xde, 0x18, 0x8e, 0x8a, 0x82, 0x70, 0x8a, 0x45,
	0x49, 0x7d, 0x0a, 0xbc, 0x75, 0xa7, 0xbb, 0x70, 0x32, 0x57, 0xff, 0x79, 0x47, 0x2f, 0x6c, 0xf4,
	0x75, 0x30, 0x3c, 0x89, 0xa2, 0x6f, 0x83, 0xc3, 0x13, 0x5d, 0x32, 0xca, 0x14, 0xd6, 0x61, 0x13,
	0x2d, 0x03, 0x1c, 0x77, 0xca, 0x1f, 0x9d, 0x66, 0x15, 0x65, 0x6a, 0x65, 0x35, 0xab, 0x65, 0xb0,
	0xb2, 0x9a, 0x9f, 0x83, 0x23, 0xfd, 0x23, 0x0c, 0xa3, 0x4c, 0x85, 0xa1, 0x55, 0x85, 0xe1, 0x32,
	0x08, 0x43, 0xab, 0x3b, 0x1b, 0xb5, 0xcd, 0x3e, 0xfe, 0x1d, 0x00, 0x00, 0xff, 0xff, 0xbd, 0x78,
	0xa4, 0x6f, 0xf3, 0x03, 0x00, 0x00,
}
