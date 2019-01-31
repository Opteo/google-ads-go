// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v0/resources/shared_criterion.proto

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

// A criterion belonging to a shared set.
type SharedCriterion struct {
	// The resource name of the shared criterion.
	// Shared set resource names have the form:
	//
	// `customers/{customer_id}/sharedCriteria/{shared_set_id}_{criterion_id}`
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// The shared set to which the shared criterion belongs.
	SharedSet *wrappers.StringValue `protobuf:"bytes,2,opt,name=shared_set,json=sharedSet,proto3" json:"shared_set,omitempty"`
	// The ID of the criterion.
	//
	// This field is ignored for mutates.
	CriterionId *wrappers.Int64Value `protobuf:"bytes,26,opt,name=criterion_id,json=criterionId,proto3" json:"criterion_id,omitempty"`
	// The type of the criterion.
	Type enums.CriterionTypeEnum_CriterionType `protobuf:"varint,4,opt,name=type,proto3,enum=google.ads.googleads.v0.enums.CriterionTypeEnum_CriterionType" json:"type,omitempty"`
	// The criterion.
	//
	// Exactly one must be set.
	//
	// Types that are valid to be assigned to Criterion:
	//	*SharedCriterion_Keyword
	//	*SharedCriterion_YoutubeVideo
	//	*SharedCriterion_YoutubeChannel
	//	*SharedCriterion_Placement
	//	*SharedCriterion_MobileAppCategory
	Criterion            isSharedCriterion_Criterion `protobuf_oneof:"criterion"`
	XXX_NoUnkeyedLiteral struct{}                    `json:"-"`
	XXX_unrecognized     []byte                      `json:"-"`
	XXX_sizecache        int32                       `json:"-"`
}

func (m *SharedCriterion) Reset()         { *m = SharedCriterion{} }
func (m *SharedCriterion) String() string { return proto.CompactTextString(m) }
func (*SharedCriterion) ProtoMessage()    {}
func (*SharedCriterion) Descriptor() ([]byte, []int) {
	return fileDescriptor_d7037cca319430ac, []int{0}
}

func (m *SharedCriterion) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SharedCriterion.Unmarshal(m, b)
}
func (m *SharedCriterion) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SharedCriterion.Marshal(b, m, deterministic)
}
func (m *SharedCriterion) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SharedCriterion.Merge(m, src)
}
func (m *SharedCriterion) XXX_Size() int {
	return xxx_messageInfo_SharedCriterion.Size(m)
}
func (m *SharedCriterion) XXX_DiscardUnknown() {
	xxx_messageInfo_SharedCriterion.DiscardUnknown(m)
}

var xxx_messageInfo_SharedCriterion proto.InternalMessageInfo

func (m *SharedCriterion) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func (m *SharedCriterion) GetSharedSet() *wrappers.StringValue {
	if m != nil {
		return m.SharedSet
	}
	return nil
}

func (m *SharedCriterion) GetCriterionId() *wrappers.Int64Value {
	if m != nil {
		return m.CriterionId
	}
	return nil
}

func (m *SharedCriterion) GetType() enums.CriterionTypeEnum_CriterionType {
	if m != nil {
		return m.Type
	}
	return enums.CriterionTypeEnum_UNSPECIFIED
}

type isSharedCriterion_Criterion interface {
	isSharedCriterion_Criterion()
}

type SharedCriterion_Keyword struct {
	Keyword *common.KeywordInfo `protobuf:"bytes,3,opt,name=keyword,proto3,oneof"`
}

type SharedCriterion_YoutubeVideo struct {
	YoutubeVideo *common.YouTubeVideoInfo `protobuf:"bytes,5,opt,name=youtube_video,json=youtubeVideo,proto3,oneof"`
}

type SharedCriterion_YoutubeChannel struct {
	YoutubeChannel *common.YouTubeChannelInfo `protobuf:"bytes,6,opt,name=youtube_channel,json=youtubeChannel,proto3,oneof"`
}

type SharedCriterion_Placement struct {
	Placement *common.PlacementInfo `protobuf:"bytes,7,opt,name=placement,proto3,oneof"`
}

type SharedCriterion_MobileAppCategory struct {
	MobileAppCategory *common.MobileAppCategoryInfo `protobuf:"bytes,8,opt,name=mobile_app_category,json=mobileAppCategory,proto3,oneof"`
}

func (*SharedCriterion_Keyword) isSharedCriterion_Criterion() {}

func (*SharedCriterion_YoutubeVideo) isSharedCriterion_Criterion() {}

func (*SharedCriterion_YoutubeChannel) isSharedCriterion_Criterion() {}

func (*SharedCriterion_Placement) isSharedCriterion_Criterion() {}

func (*SharedCriterion_MobileAppCategory) isSharedCriterion_Criterion() {}

func (m *SharedCriterion) GetCriterion() isSharedCriterion_Criterion {
	if m != nil {
		return m.Criterion
	}
	return nil
}

func (m *SharedCriterion) GetKeyword() *common.KeywordInfo {
	if x, ok := m.GetCriterion().(*SharedCriterion_Keyword); ok {
		return x.Keyword
	}
	return nil
}

func (m *SharedCriterion) GetYoutubeVideo() *common.YouTubeVideoInfo {
	if x, ok := m.GetCriterion().(*SharedCriterion_YoutubeVideo); ok {
		return x.YoutubeVideo
	}
	return nil
}

func (m *SharedCriterion) GetYoutubeChannel() *common.YouTubeChannelInfo {
	if x, ok := m.GetCriterion().(*SharedCriterion_YoutubeChannel); ok {
		return x.YoutubeChannel
	}
	return nil
}

func (m *SharedCriterion) GetPlacement() *common.PlacementInfo {
	if x, ok := m.GetCriterion().(*SharedCriterion_Placement); ok {
		return x.Placement
	}
	return nil
}

func (m *SharedCriterion) GetMobileAppCategory() *common.MobileAppCategoryInfo {
	if x, ok := m.GetCriterion().(*SharedCriterion_MobileAppCategory); ok {
		return x.MobileAppCategory
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*SharedCriterion) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*SharedCriterion_Keyword)(nil),
		(*SharedCriterion_YoutubeVideo)(nil),
		(*SharedCriterion_YoutubeChannel)(nil),
		(*SharedCriterion_Placement)(nil),
		(*SharedCriterion_MobileAppCategory)(nil),
	}
}

func init() {
	proto.RegisterType((*SharedCriterion)(nil), "google.ads.googleads.v0.resources.SharedCriterion")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v0/resources/shared_criterion.proto", fileDescriptor_d7037cca319430ac)
}

var fileDescriptor_d7037cca319430ac = []byte{
	// 559 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x94, 0xdf, 0x8a, 0xd3, 0x4e,
	0x14, 0xc7, 0x7f, 0xed, 0xfe, 0xfb, 0x75, 0xda, 0xdd, 0xc5, 0xe8, 0x45, 0xa8, 0x22, 0x5d, 0x65,
	0xa1, 0x20, 0x3b, 0x09, 0xf5, 0x0f, 0x92, 0x85, 0x85, 0xb4, 0x48, 0xad, 0xb2, 0x52, 0xd2, 0xa5,
	0xa2, 0x54, 0xc2, 0x34, 0x39, 0x9b, 0x0d, 0x26, 0x33, 0x61, 0x26, 0xe9, 0xd2, 0x4b, 0x5f, 0xc5,
	0x4b, 0x1f, 0x45, 0xf0, 0x45, 0x7c, 0x0a, 0xe9, 0xcc, 0x24, 0x8b, 0x2b, 0xb5, 0xde, 0x9d, 0xce,
	0x7c, 0x3f, 0x9f, 0x9c, 0x93, 0xce, 0x04, 0xbd, 0x8c, 0x18, 0x8b, 0x12, 0xb0, 0x48, 0x28, 0x2c,
	0x55, 0xae, 0xaa, 0x85, 0x6d, 0x71, 0x10, 0xac, 0xe0, 0x01, 0x08, 0x4b, 0x5c, 0x11, 0x0e, 0xa1,
	0x1f, 0xf0, 0x38, 0x07, 0x1e, 0x33, 0x8a, 0x33, 0xce, 0x72, 0x66, 0x1c, 0xa9, 0x38, 0x26, 0xa1,
	0xc0, 0x15, 0x89, 0x17, 0x36, 0xae, 0xc8, 0xf6, 0xc9, 0x3a, 0x79, 0xc0, 0xd2, 0x94, 0x51, 0x4b,
	0x2b, 0x89, 0x32, 0xb6, 0x7b, 0xeb, 0xe2, 0x40, 0x8b, 0x54, 0x58, 0x55, 0x03, 0x7e, 0xbe, 0xcc,
	0x40, 0x33, 0x0f, 0x35, 0x23, 0x7f, 0xcd, 0x8b, 0x4b, 0xeb, 0x9a, 0x93, 0x2c, 0x03, 0x2e, 0xd4,
	0xfe, 0xa3, 0x1f, 0x3b, 0xe8, 0x70, 0x22, 0x07, 0x18, 0x94, 0xb8, 0xf1, 0x18, 0xed, 0x97, 0x3d,
	0xfa, 0x94, 0xa4, 0x60, 0xd6, 0x3a, 0xb5, 0x6e, 0xc3, 0x6b, 0x95, 0x8b, 0xef, 0x48, 0x0a, 0xc6,
	0x29, 0x42, 0x7a, 0x70, 0x01, 0xb9, 0x59, 0xef, 0xd4, 0xba, 0xcd, 0xde, 0x03, 0x3d, 0x28, 0x2e,
	0x9f, 0x86, 0x27, 0x39, 0x8f, 0x69, 0x34, 0x25, 0x49, 0x01, 0x5e, 0x43, 0xe5, 0x27, 0x90, 0x1b,
	0x67, 0xa8, 0x75, 0xd3, 0x6d, 0x1c, 0x9a, 0x6d, 0x89, 0xdf, 0xff, 0x03, 0x1f, 0xd1, 0xfc, 0xc5,
	0x33, 0x45, 0x37, 0x2b, 0x60, 0x14, 0x1a, 0x1e, 0xda, 0x5e, 0xcd, 0x68, 0x6e, 0x77, 0x6a, 0xdd,
	0x83, 0xde, 0x19, 0x5e, 0xf7, 0xaa, 0xe5, 0x8b, 0xc1, 0xd5, 0x64, 0x17, 0xcb, 0x0c, 0x5e, 0xd1,
	0x22, 0xfd, 0x7d, 0xc5, 0x93, 0x2e, 0x63, 0x88, 0xf6, 0x3e, 0xc3, 0xf2, 0x9a, 0xf1, 0xd0, 0xdc,
	0x92, 0xed, 0x3c, 0x59, 0xab, 0x55, 0x7f, 0x0f, 0x7e, 0xab, 0xe2, 0x23, 0x7a, 0xc9, 0x5e, 0xff,
	0xe7, 0x95, 0xb4, 0xf1, 0x1e, 0xed, 0x2f, 0x59, 0x91, 0x17, 0x73, 0xf0, 0x17, 0x71, 0x08, 0xcc,
	0xdc, 0x91, 0x3a, 0x7b, 0x93, 0xee, 0x03, 0x2b, 0x2e, 0x8a, 0x39, 0x4c, 0x57, 0x8c, 0x76, 0xb6,
	0xb4, 0x48, 0xae, 0x19, 0x9f, 0xd0, 0x61, 0x29, 0x0e, 0xae, 0x08, 0xa5, 0x90, 0x98, 0xbb, 0x52,
	0xdd, 0xfb, 0x47, 0xf5, 0x40, 0x51, 0x5a, 0x7e, 0xa0, 0x65, 0x7a, 0xd5, 0x38, 0x47, 0x8d, 0x2c,
	0x21, 0x01, 0xa4, 0x40, 0x73, 0x73, 0x4f, 0x8a, 0x4f, 0x36, 0x89, 0xc7, 0x25, 0xa0, 0x9d, 0x37,
	0x06, 0x23, 0x42, 0x77, 0x53, 0x36, 0x8f, 0x13, 0xf0, 0x49, 0x96, 0xf9, 0x01, 0xc9, 0x21, 0x62,
	0x7c, 0x69, 0xfe, 0x2f, 0xc5, 0xcf, 0x37, 0x89, 0xcf, 0x25, 0xea, 0x66, 0xd9, 0x40, 0x83, 0xfa,
	0x01, 0x77, 0xd2, 0xdb, 0x1b, 0xfd, 0x26, 0x6a, 0x54, 0x67, 0xa3, 0xff, 0xa5, 0x8e, 0x8e, 0x03,
	0x96, 0xe2, 0x8d, 0x97, 0xaf, 0x7f, 0xef, 0xd6, 0xb1, 0x1f, 0xaf, 0x0e, 0xdd, 0xb8, 0xf6, 0xf1,
	0x8d, 0x46, 0x23, 0x96, 0x10, 0x1a, 0x61, 0xc6, 0x23, 0x2b, 0x02, 0x2a, 0x8f, 0x64, 0x79, 0xeb,
	0xb2, 0x58, 0xfc, 0xe5, 0x83, 0x70, 0x5a, 0x55, 0x5f, 0xeb, 0x5b, 0x43, 0xd7, 0xfd, 0x56, 0x3f,
	0x1a, 0x2a, 0xa5, 0x1b, 0x0a, 0xac, 0xca, 0x55, 0x35, 0xb5, 0xb1, 0x57, 0x26, 0xbf, 0x97, 0x99,
	0x99, 0x1b, 0x8a, 0x59, 0x95, 0x99, 0x4d, 0xed, 0x59, 0x95, 0xf9, 0x59, 0x3f, 0x56, 0x1b, 0x8e,
	0xe3, 0x86, 0xc2, 0x71, 0xaa, 0x94, 0xe3, 0x4c, 0x6d, 0xc7, 0xa9, 0x72, 0xf3, 0x5d, 0xd9, 0xec,
	0xd3, 0x5f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xbe, 0x05, 0xb8, 0xd3, 0xbc, 0x04, 0x00, 0x00,
}
