// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v0/resources/campaign_criterion.proto

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

// A campaign criterion.
type CampaignCriterion struct {
	// The resource name of the campaign criterion.
	// Campaign criterion resource names have the form:
	//
	// `customers/{customer_id}/campaignCriteria/{campaign_id}_{criterion_id}`
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// The campaign to which the criterion belongs.
	Campaign *wrappers.StringValue `protobuf:"bytes,4,opt,name=campaign,proto3" json:"campaign,omitempty"`
	// The ID of the criterion.
	//
	// This field is ignored during mutate.
	CriterionId *wrappers.Int64Value `protobuf:"bytes,5,opt,name=criterion_id,json=criterionId,proto3" json:"criterion_id,omitempty"`
	// The modifier for the bids when the criterion matches. The modifier must be
	// in the range: 0.1 - 10.0. Most targetable criteria types support modifiers.
	// Use 0 to opt out of a Device type.
	BidModifier *wrappers.FloatValue `protobuf:"bytes,14,opt,name=bid_modifier,json=bidModifier,proto3" json:"bid_modifier,omitempty"`
	// Whether to target (`false`) or exclude (`true`) the criterion.
	Negative *wrappers.BoolValue `protobuf:"bytes,7,opt,name=negative,proto3" json:"negative,omitempty"`
	// The type of the criterion.
	Type enums.CriterionTypeEnum_CriterionType `protobuf:"varint,6,opt,name=type,proto3,enum=google.ads.googleads.v0.enums.CriterionTypeEnum_CriterionType" json:"type,omitempty"`
	// The campaign criterion.
	//
	// Exactly one must be set.
	//
	// Types that are valid to be assigned to Criterion:
	//	*CampaignCriterion_Keyword
	//	*CampaignCriterion_Placement
	//	*CampaignCriterion_MobileAppCategory
	//	*CampaignCriterion_Location
	//	*CampaignCriterion_Device
	//	*CampaignCriterion_AdSchedule
	//	*CampaignCriterion_AgeRange
	//	*CampaignCriterion_Gender
	//	*CampaignCriterion_IncomeRange
	//	*CampaignCriterion_ParentalStatus
	//	*CampaignCriterion_UserList
	//	*CampaignCriterion_YoutubeVideo
	//	*CampaignCriterion_YoutubeChannel
	//	*CampaignCriterion_Proximity
	//	*CampaignCriterion_Topic
	//	*CampaignCriterion_ListingScope
	//	*CampaignCriterion_Language
	//	*CampaignCriterion_IpBlock
	//	*CampaignCriterion_ContentLabel
	//	*CampaignCriterion_Carrier
	//	*CampaignCriterion_UserInterest
	//	*CampaignCriterion_Webpage
	//	*CampaignCriterion_OperatingSystemVersion
	Criterion            isCampaignCriterion_Criterion `protobuf_oneof:"criterion"`
	XXX_NoUnkeyedLiteral struct{}                      `json:"-"`
	XXX_unrecognized     []byte                        `json:"-"`
	XXX_sizecache        int32                         `json:"-"`
}

func (m *CampaignCriterion) Reset()         { *m = CampaignCriterion{} }
func (m *CampaignCriterion) String() string { return proto.CompactTextString(m) }
func (*CampaignCriterion) ProtoMessage()    {}
func (*CampaignCriterion) Descriptor() ([]byte, []int) {
	return fileDescriptor_889de17d165924ac, []int{0}
}

func (m *CampaignCriterion) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CampaignCriterion.Unmarshal(m, b)
}
func (m *CampaignCriterion) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CampaignCriterion.Marshal(b, m, deterministic)
}
func (m *CampaignCriterion) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CampaignCriterion.Merge(m, src)
}
func (m *CampaignCriterion) XXX_Size() int {
	return xxx_messageInfo_CampaignCriterion.Size(m)
}
func (m *CampaignCriterion) XXX_DiscardUnknown() {
	xxx_messageInfo_CampaignCriterion.DiscardUnknown(m)
}

var xxx_messageInfo_CampaignCriterion proto.InternalMessageInfo

func (m *CampaignCriterion) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func (m *CampaignCriterion) GetCampaign() *wrappers.StringValue {
	if m != nil {
		return m.Campaign
	}
	return nil
}

func (m *CampaignCriterion) GetCriterionId() *wrappers.Int64Value {
	if m != nil {
		return m.CriterionId
	}
	return nil
}

func (m *CampaignCriterion) GetBidModifier() *wrappers.FloatValue {
	if m != nil {
		return m.BidModifier
	}
	return nil
}

func (m *CampaignCriterion) GetNegative() *wrappers.BoolValue {
	if m != nil {
		return m.Negative
	}
	return nil
}

func (m *CampaignCriterion) GetType() enums.CriterionTypeEnum_CriterionType {
	if m != nil {
		return m.Type
	}
	return enums.CriterionTypeEnum_UNSPECIFIED
}

type isCampaignCriterion_Criterion interface {
	isCampaignCriterion_Criterion()
}

type CampaignCriterion_Keyword struct {
	Keyword *common.KeywordInfo `protobuf:"bytes,8,opt,name=keyword,proto3,oneof"`
}

type CampaignCriterion_Placement struct {
	Placement *common.PlacementInfo `protobuf:"bytes,9,opt,name=placement,proto3,oneof"`
}

type CampaignCriterion_MobileAppCategory struct {
	MobileAppCategory *common.MobileAppCategoryInfo `protobuf:"bytes,10,opt,name=mobile_app_category,json=mobileAppCategory,proto3,oneof"`
}

type CampaignCriterion_Location struct {
	Location *common.LocationInfo `protobuf:"bytes,12,opt,name=location,proto3,oneof"`
}

type CampaignCriterion_Device struct {
	Device *common.DeviceInfo `protobuf:"bytes,13,opt,name=device,proto3,oneof"`
}

type CampaignCriterion_AdSchedule struct {
	AdSchedule *common.AdScheduleInfo `protobuf:"bytes,15,opt,name=ad_schedule,json=adSchedule,proto3,oneof"`
}

type CampaignCriterion_AgeRange struct {
	AgeRange *common.AgeRangeInfo `protobuf:"bytes,16,opt,name=age_range,json=ageRange,proto3,oneof"`
}

type CampaignCriterion_Gender struct {
	Gender *common.GenderInfo `protobuf:"bytes,17,opt,name=gender,proto3,oneof"`
}

type CampaignCriterion_IncomeRange struct {
	IncomeRange *common.IncomeRangeInfo `protobuf:"bytes,18,opt,name=income_range,json=incomeRange,proto3,oneof"`
}

type CampaignCriterion_ParentalStatus struct {
	ParentalStatus *common.ParentalStatusInfo `protobuf:"bytes,19,opt,name=parental_status,json=parentalStatus,proto3,oneof"`
}

type CampaignCriterion_UserList struct {
	UserList *common.UserListInfo `protobuf:"bytes,22,opt,name=user_list,json=userList,proto3,oneof"`
}

type CampaignCriterion_YoutubeVideo struct {
	YoutubeVideo *common.YouTubeVideoInfo `protobuf:"bytes,20,opt,name=youtube_video,json=youtubeVideo,proto3,oneof"`
}

type CampaignCriterion_YoutubeChannel struct {
	YoutubeChannel *common.YouTubeChannelInfo `protobuf:"bytes,21,opt,name=youtube_channel,json=youtubeChannel,proto3,oneof"`
}

type CampaignCriterion_Proximity struct {
	Proximity *common.ProximityInfo `protobuf:"bytes,23,opt,name=proximity,proto3,oneof"`
}

type CampaignCriterion_Topic struct {
	Topic *common.TopicInfo `protobuf:"bytes,24,opt,name=topic,proto3,oneof"`
}

type CampaignCriterion_ListingScope struct {
	ListingScope *common.ListingScopeInfo `protobuf:"bytes,25,opt,name=listing_scope,json=listingScope,proto3,oneof"`
}

type CampaignCriterion_Language struct {
	Language *common.LanguageInfo `protobuf:"bytes,26,opt,name=language,proto3,oneof"`
}

type CampaignCriterion_IpBlock struct {
	IpBlock *common.IpBlockInfo `protobuf:"bytes,27,opt,name=ip_block,json=ipBlock,proto3,oneof"`
}

type CampaignCriterion_ContentLabel struct {
	ContentLabel *common.ContentLabelInfo `protobuf:"bytes,28,opt,name=content_label,json=contentLabel,proto3,oneof"`
}

type CampaignCriterion_Carrier struct {
	Carrier *common.CarrierInfo `protobuf:"bytes,29,opt,name=carrier,proto3,oneof"`
}

type CampaignCriterion_UserInterest struct {
	UserInterest *common.UserInterestInfo `protobuf:"bytes,30,opt,name=user_interest,json=userInterest,proto3,oneof"`
}

type CampaignCriterion_Webpage struct {
	Webpage *common.WebpageInfo `protobuf:"bytes,31,opt,name=webpage,proto3,oneof"`
}

type CampaignCriterion_OperatingSystemVersion struct {
	OperatingSystemVersion *common.OperatingSystemVersionInfo `protobuf:"bytes,32,opt,name=operating_system_version,json=operatingSystemVersion,proto3,oneof"`
}

func (*CampaignCriterion_Keyword) isCampaignCriterion_Criterion() {}

func (*CampaignCriterion_Placement) isCampaignCriterion_Criterion() {}

func (*CampaignCriterion_MobileAppCategory) isCampaignCriterion_Criterion() {}

func (*CampaignCriterion_Location) isCampaignCriterion_Criterion() {}

func (*CampaignCriterion_Device) isCampaignCriterion_Criterion() {}

func (*CampaignCriterion_AdSchedule) isCampaignCriterion_Criterion() {}

func (*CampaignCriterion_AgeRange) isCampaignCriterion_Criterion() {}

func (*CampaignCriterion_Gender) isCampaignCriterion_Criterion() {}

func (*CampaignCriterion_IncomeRange) isCampaignCriterion_Criterion() {}

func (*CampaignCriterion_ParentalStatus) isCampaignCriterion_Criterion() {}

func (*CampaignCriterion_UserList) isCampaignCriterion_Criterion() {}

func (*CampaignCriterion_YoutubeVideo) isCampaignCriterion_Criterion() {}

func (*CampaignCriterion_YoutubeChannel) isCampaignCriterion_Criterion() {}

func (*CampaignCriterion_Proximity) isCampaignCriterion_Criterion() {}

func (*CampaignCriterion_Topic) isCampaignCriterion_Criterion() {}

func (*CampaignCriterion_ListingScope) isCampaignCriterion_Criterion() {}

func (*CampaignCriterion_Language) isCampaignCriterion_Criterion() {}

func (*CampaignCriterion_IpBlock) isCampaignCriterion_Criterion() {}

func (*CampaignCriterion_ContentLabel) isCampaignCriterion_Criterion() {}

func (*CampaignCriterion_Carrier) isCampaignCriterion_Criterion() {}

func (*CampaignCriterion_UserInterest) isCampaignCriterion_Criterion() {}

func (*CampaignCriterion_Webpage) isCampaignCriterion_Criterion() {}

func (*CampaignCriterion_OperatingSystemVersion) isCampaignCriterion_Criterion() {}

func (m *CampaignCriterion) GetCriterion() isCampaignCriterion_Criterion {
	if m != nil {
		return m.Criterion
	}
	return nil
}

func (m *CampaignCriterion) GetKeyword() *common.KeywordInfo {
	if x, ok := m.GetCriterion().(*CampaignCriterion_Keyword); ok {
		return x.Keyword
	}
	return nil
}

func (m *CampaignCriterion) GetPlacement() *common.PlacementInfo {
	if x, ok := m.GetCriterion().(*CampaignCriterion_Placement); ok {
		return x.Placement
	}
	return nil
}

func (m *CampaignCriterion) GetMobileAppCategory() *common.MobileAppCategoryInfo {
	if x, ok := m.GetCriterion().(*CampaignCriterion_MobileAppCategory); ok {
		return x.MobileAppCategory
	}
	return nil
}

func (m *CampaignCriterion) GetLocation() *common.LocationInfo {
	if x, ok := m.GetCriterion().(*CampaignCriterion_Location); ok {
		return x.Location
	}
	return nil
}

func (m *CampaignCriterion) GetDevice() *common.DeviceInfo {
	if x, ok := m.GetCriterion().(*CampaignCriterion_Device); ok {
		return x.Device
	}
	return nil
}

func (m *CampaignCriterion) GetAdSchedule() *common.AdScheduleInfo {
	if x, ok := m.GetCriterion().(*CampaignCriterion_AdSchedule); ok {
		return x.AdSchedule
	}
	return nil
}

func (m *CampaignCriterion) GetAgeRange() *common.AgeRangeInfo {
	if x, ok := m.GetCriterion().(*CampaignCriterion_AgeRange); ok {
		return x.AgeRange
	}
	return nil
}

func (m *CampaignCriterion) GetGender() *common.GenderInfo {
	if x, ok := m.GetCriterion().(*CampaignCriterion_Gender); ok {
		return x.Gender
	}
	return nil
}

func (m *CampaignCriterion) GetIncomeRange() *common.IncomeRangeInfo {
	if x, ok := m.GetCriterion().(*CampaignCriterion_IncomeRange); ok {
		return x.IncomeRange
	}
	return nil
}

func (m *CampaignCriterion) GetParentalStatus() *common.ParentalStatusInfo {
	if x, ok := m.GetCriterion().(*CampaignCriterion_ParentalStatus); ok {
		return x.ParentalStatus
	}
	return nil
}

func (m *CampaignCriterion) GetUserList() *common.UserListInfo {
	if x, ok := m.GetCriterion().(*CampaignCriterion_UserList); ok {
		return x.UserList
	}
	return nil
}

func (m *CampaignCriterion) GetYoutubeVideo() *common.YouTubeVideoInfo {
	if x, ok := m.GetCriterion().(*CampaignCriterion_YoutubeVideo); ok {
		return x.YoutubeVideo
	}
	return nil
}

func (m *CampaignCriterion) GetYoutubeChannel() *common.YouTubeChannelInfo {
	if x, ok := m.GetCriterion().(*CampaignCriterion_YoutubeChannel); ok {
		return x.YoutubeChannel
	}
	return nil
}

func (m *CampaignCriterion) GetProximity() *common.ProximityInfo {
	if x, ok := m.GetCriterion().(*CampaignCriterion_Proximity); ok {
		return x.Proximity
	}
	return nil
}

func (m *CampaignCriterion) GetTopic() *common.TopicInfo {
	if x, ok := m.GetCriterion().(*CampaignCriterion_Topic); ok {
		return x.Topic
	}
	return nil
}

func (m *CampaignCriterion) GetListingScope() *common.ListingScopeInfo {
	if x, ok := m.GetCriterion().(*CampaignCriterion_ListingScope); ok {
		return x.ListingScope
	}
	return nil
}

func (m *CampaignCriterion) GetLanguage() *common.LanguageInfo {
	if x, ok := m.GetCriterion().(*CampaignCriterion_Language); ok {
		return x.Language
	}
	return nil
}

func (m *CampaignCriterion) GetIpBlock() *common.IpBlockInfo {
	if x, ok := m.GetCriterion().(*CampaignCriterion_IpBlock); ok {
		return x.IpBlock
	}
	return nil
}

func (m *CampaignCriterion) GetContentLabel() *common.ContentLabelInfo {
	if x, ok := m.GetCriterion().(*CampaignCriterion_ContentLabel); ok {
		return x.ContentLabel
	}
	return nil
}

func (m *CampaignCriterion) GetCarrier() *common.CarrierInfo {
	if x, ok := m.GetCriterion().(*CampaignCriterion_Carrier); ok {
		return x.Carrier
	}
	return nil
}

func (m *CampaignCriterion) GetUserInterest() *common.UserInterestInfo {
	if x, ok := m.GetCriterion().(*CampaignCriterion_UserInterest); ok {
		return x.UserInterest
	}
	return nil
}

func (m *CampaignCriterion) GetWebpage() *common.WebpageInfo {
	if x, ok := m.GetCriterion().(*CampaignCriterion_Webpage); ok {
		return x.Webpage
	}
	return nil
}

func (m *CampaignCriterion) GetOperatingSystemVersion() *common.OperatingSystemVersionInfo {
	if x, ok := m.GetCriterion().(*CampaignCriterion_OperatingSystemVersion); ok {
		return x.OperatingSystemVersion
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*CampaignCriterion) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*CampaignCriterion_Keyword)(nil),
		(*CampaignCriterion_Placement)(nil),
		(*CampaignCriterion_MobileAppCategory)(nil),
		(*CampaignCriterion_Location)(nil),
		(*CampaignCriterion_Device)(nil),
		(*CampaignCriterion_AdSchedule)(nil),
		(*CampaignCriterion_AgeRange)(nil),
		(*CampaignCriterion_Gender)(nil),
		(*CampaignCriterion_IncomeRange)(nil),
		(*CampaignCriterion_ParentalStatus)(nil),
		(*CampaignCriterion_UserList)(nil),
		(*CampaignCriterion_YoutubeVideo)(nil),
		(*CampaignCriterion_YoutubeChannel)(nil),
		(*CampaignCriterion_Proximity)(nil),
		(*CampaignCriterion_Topic)(nil),
		(*CampaignCriterion_ListingScope)(nil),
		(*CampaignCriterion_Language)(nil),
		(*CampaignCriterion_IpBlock)(nil),
		(*CampaignCriterion_ContentLabel)(nil),
		(*CampaignCriterion_Carrier)(nil),
		(*CampaignCriterion_UserInterest)(nil),
		(*CampaignCriterion_Webpage)(nil),
		(*CampaignCriterion_OperatingSystemVersion)(nil),
	}
}

func init() {
	proto.RegisterType((*CampaignCriterion)(nil), "google.ads.googleads.v0.resources.CampaignCriterion")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v0/resources/campaign_criterion.proto", fileDescriptor_889de17d165924ac)
}

var fileDescriptor_889de17d165924ac = []byte{
	// 1007 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x96, 0xdd, 0x6e, 0x1c, 0x35,
	0x14, 0xc7, 0x49, 0xe8, 0xc7, 0xc6, 0xf9, 0x28, 0x99, 0x42, 0x30, 0x69, 0x29, 0x29, 0xa8, 0x52,
	0xf8, 0xe8, 0xec, 0x2a, 0x40, 0x85, 0x16, 0xa9, 0xd2, 0x66, 0x0b, 0xe9, 0xb6, 0x09, 0x84, 0x4d,
	0xd8, 0x0a, 0x14, 0x34, 0xf2, 0xce, 0x9c, 0x4c, 0xad, 0xce, 0xd8, 0x96, 0xed, 0xd9, 0xb0, 0xd7,
	0xbc, 0x09, 0x97, 0x3c, 0x0a, 0x8f, 0xc2, 0x2b, 0x70, 0x83, 0xfc, 0x35, 0x49, 0x09, 0x61, 0xe6,
	0xce, 0x3e, 0xf3, 0xff, 0xfd, 0xe7, 0x9c, 0x33, 0xb6, 0xc7, 0xa8, 0x9f, 0x73, 0x9e, 0x17, 0xd0,
	0x25, 0x99, 0xea, 0xba, 0xa1, 0x19, 0xcd, 0x7a, 0x5d, 0x09, 0x8a, 0x57, 0x32, 0x05, 0xd5, 0x4d,
	0x49, 0x29, 0x08, 0xcd, 0x59, 0x92, 0x4a, 0xaa, 0x41, 0x52, 0xce, 0x62, 0x21, 0xb9, 0xe6, 0xd1,
	0x7d, 0x07, 0xc4, 0x24, 0x53, 0x71, 0xcd, 0xc6, 0xb3, 0x5e, 0x5c, 0xb3, 0x9b, 0x0f, 0xaf, 0xb2,
	0x4f, 0x79, 0x59, 0x72, 0xd6, 0xf5, 0x96, 0xc4, 0x39, 0x6e, 0xee, 0x5c, 0x25, 0x07, 0x56, 0x95,
	0xaa, 0x5b, 0x27, 0x90, 0xe8, 0xb9, 0x00, 0xcf, 0xdc, 0xf3, 0x8c, 0x9d, 0x4d, 0xab, 0xd3, 0xee,
	0x99, 0x24, 0x42, 0x80, 0x54, 0xee, 0xf9, 0x87, 0x7f, 0xaf, 0xa3, 0xf5, 0xa1, 0x2f, 0x61, 0x18,
	0x0c, 0xa2, 0x8f, 0xd0, 0x6a, 0xc8, 0x32, 0x61, 0xa4, 0x04, 0xbc, 0xb0, 0xb5, 0xb0, 0xbd, 0x34,
	0x5e, 0x09, 0xc1, 0xef, 0x48, 0x09, 0xd1, 0x57, 0xa8, 0x13, 0x8a, 0xc7, 0xd7, 0xb6, 0x16, 0xb6,
	0x97, 0x77, 0xee, 0xfa, 0x42, 0xe3, 0xf0, 0xb6, 0xf8, 0x48, 0x4b, 0xca, 0xf2, 0x09, 0x29, 0x2a,
	0x18, 0xd7, 0xea, 0xe8, 0x31, 0x5a, 0x39, 0x4f, 0x96, 0x66, 0xf8, 0xba, 0xa5, 0xef, 0x5c, 0xa2,
	0x47, 0x4c, 0x3f, 0xfa, 0xc2, 0xc1, 0xcb, 0x35, 0x30, 0xca, 0x0c, 0x3f, 0xa5, 0x59, 0x52, 0xf2,
	0x8c, 0x9e, 0x52, 0x90, 0x78, 0xed, 0x0a, 0xfe, 0xdb, 0x82, 0x13, 0xed, 0xf9, 0x29, 0xcd, 0x0e,
	0xbc, 0x3e, 0x7a, 0x84, 0x3a, 0x0c, 0x72, 0xa2, 0xe9, 0x0c, 0xf0, 0x4d, 0xcb, 0x6e, 0x5e, 0x62,
	0x77, 0x39, 0x2f, 0x7c, 0xde, 0x41, 0x1b, 0x8d, 0xd1, 0x35, 0xd3, 0x5a, 0x7c, 0x63, 0x6b, 0x61,
	0x7b, 0x6d, 0xe7, 0x71, 0x7c, 0xd5, 0x17, 0xb6, 0xdf, 0x23, 0xae, 0xdb, 0x79, 0x3c, 0x17, 0xf0,
	0x0d, 0xab, 0xca, 0xd7, 0x23, 0x63, 0xeb, 0x15, 0xed, 0xa1, 0x9b, 0xaf, 0x60, 0x7e, 0xc6, 0x65,
	0x86, 0x3b, 0x36, 0x95, 0x4f, 0xaf, 0xb4, 0x75, 0xab, 0x22, 0x7e, 0xee, 0xe4, 0x23, 0x76, 0xca,
	0x9f, 0xbe, 0x31, 0x0e, 0x74, 0x74, 0x80, 0x96, 0x44, 0x41, 0x52, 0x28, 0x81, 0x69, 0xbc, 0x64,
	0xad, 0x1e, 0x36, 0x59, 0x1d, 0x06, 0xc0, 0x9b, 0x9d, 0x3b, 0x44, 0x39, 0xba, 0x5d, 0xf2, 0x29,
	0x2d, 0x20, 0x21, 0x42, 0x24, 0x29, 0xd1, 0x90, 0x73, 0x39, 0xc7, 0xc8, 0x1a, 0x7f, 0xd9, 0x64,
	0x7c, 0x60, 0xd1, 0x81, 0x10, 0x43, 0x0f, 0xfa, 0x17, 0xac, 0x97, 0xff, 0x7e, 0x10, 0x3d, 0x43,
	0x9d, 0x82, 0xa7, 0x44, 0x53, 0xce, 0xf0, 0x8a, 0x75, 0xff, 0xac, 0xc9, 0x7d, 0xdf, 0xeb, 0xbd,
	0x69, 0xcd, 0x47, 0x4f, 0xd0, 0x8d, 0x0c, 0x66, 0x34, 0x05, 0xbc, 0x6a, 0x9d, 0x3e, 0x69, 0x72,
	0x7a, 0x62, 0xd5, 0xde, 0xc7, 0xb3, 0xd1, 0x0f, 0x68, 0x99, 0x64, 0x89, 0x4a, 0x5f, 0x42, 0x56,
	0x15, 0x80, 0x6f, 0x59, 0xab, 0xb8, 0xc9, 0x6a, 0x90, 0x1d, 0x79, 0xc2, 0xdb, 0x21, 0x52, 0x47,
	0xa2, 0xe7, 0x68, 0x89, 0xe4, 0x90, 0x48, 0xc2, 0x72, 0xc0, 0x6f, 0xb5, 0xab, 0x72, 0x90, 0xc3,
	0xd8, 0xe8, 0x43, 0x95, 0xc4, 0xcf, 0x4d, 0x95, 0x39, 0xb0, 0x0c, 0x24, 0x5e, 0x6f, 0x57, 0xe5,
	0x9e, 0x55, 0x87, 0x2a, 0x1d, 0x1b, 0x1d, 0xa3, 0x15, 0xca, 0x52, 0x5e, 0x86, 0xac, 0x22, 0xeb,
	0xd5, 0x6d, 0xf2, 0x1a, 0x59, 0xe6, 0x62, 0x62, 0xcb, 0xf4, 0x3c, 0x14, 0xfd, 0x82, 0x6e, 0x09,
	0x22, 0x81, 0x69, 0x52, 0x24, 0x4a, 0x13, 0x5d, 0x29, 0x7c, 0xdb, 0x1a, 0xef, 0x34, 0xae, 0x45,
	0x8f, 0x1d, 0x59, 0xca, 0x7b, 0xaf, 0x89, 0xd7, 0xa2, 0xa6, 0x8f, 0x95, 0x02, 0x99, 0x14, 0x54,
	0x69, 0xbc, 0xd1, 0xae, 0x8f, 0x3f, 0x2a, 0x90, 0xfb, 0x54, 0x85, 0x35, 0xde, 0xa9, 0xfc, 0x3c,
	0x7a, 0x81, 0x56, 0xe7, 0xbc, 0xd2, 0xd5, 0x14, 0x92, 0x19, 0xcd, 0x80, 0xe3, 0xb7, 0xad, 0x61,
	0xaf, 0xc9, 0xf0, 0x27, 0x5e, 0x1d, 0x57, 0x53, 0x98, 0x18, 0xc6, 0x9b, 0xae, 0x78, 0x23, 0x1b,
	0x33, 0x4d, 0x08, 0xc6, 0xe9, 0x4b, 0xc2, 0x18, 0x14, 0xf8, 0x9d, 0x76, 0x4d, 0xf0, 0xd6, 0x43,
	0x47, 0x85, 0x26, 0x78, 0x33, 0x1f, 0xb5, 0x3b, 0x5d, 0xf2, 0x5f, 0x69, 0x49, 0xf5, 0x1c, 0xbf,
	0xdb, 0x72, 0xa7, 0x07, 0xa0, 0xde, 0xe9, 0x21, 0x10, 0x0d, 0xd0, 0x75, 0xcd, 0x05, 0x4d, 0x31,
	0xb6, 0x56, 0x1f, 0x37, 0x59, 0x1d, 0x1b, 0xb1, 0xb7, 0x71, 0xa4, 0xe9, 0xa4, 0xf9, 0x22, 0x94,
	0xe5, 0x89, 0x4a, 0xb9, 0x00, 0xfc, 0x5e, 0xbb, 0x4e, 0xee, 0x3b, 0xe8, 0xc8, 0x30, 0xa1, 0x93,
	0xc5, 0x85, 0x98, 0x3d, 0x1c, 0x08, 0xcb, 0x2b, 0x92, 0x03, 0xde, 0x6c, 0x79, 0x38, 0x78, 0x7d,
	0x7d, 0x38, 0xf8, 0x79, 0xf4, 0x14, 0x75, 0xa8, 0x48, 0xa6, 0x05, 0x4f, 0x5f, 0xe1, 0x3b, 0xed,
	0x8e, 0xda, 0x91, 0xd8, 0x35, 0xf2, 0x70, 0xd4, 0x52, 0x37, 0x35, 0xe5, 0xa6, 0x9c, 0x69, 0x60,
	0x3a, 0x29, 0xc8, 0x14, 0x0a, 0x7c, 0xb7, 0x5d, 0xb9, 0x43, 0x07, 0xed, 0x1b, 0x26, 0x94, 0x9b,
	0x5e, 0x88, 0x99, 0x9f, 0x41, 0x4a, 0xa4, 0x34, 0xff, 0xb4, 0xf7, 0xdb, 0x65, 0x38, 0x74, 0xf2,
	0x90, 0xa1, 0xa7, 0x4d, 0x86, 0x76, 0x9f, 0x50, 0xa6, 0x41, 0x82, 0xd2, 0xf8, 0x5e, 0xbb, 0x0c,
	0xcd, 0x5e, 0x19, 0x79, 0x26, 0x64, 0x58, 0x5d, 0x88, 0x99, 0x0c, 0xcf, 0x60, 0x2a, 0xcc, 0xf7,
	0xf8, 0xa0, 0x5d, 0x86, 0x2f, 0x9c, 0x3c, 0x64, 0xe8, 0xe9, 0x68, 0x86, 0x30, 0x17, 0x20, 0x89,
	0x5b, 0x34, 0x73, 0xa5, 0xa1, 0x4c, 0x66, 0x20, 0x95, 0xf9, 0x0d, 0x6c, 0x59, 0xe7, 0x7e, 0x93,
	0xf3, 0xf7, 0x81, 0x3f, 0xb2, 0xf8, 0xc4, 0xd1, 0xfe, 0x45, 0x1b, 0xfc, 0x3f, 0x9f, 0xee, 0x2e,
	0xa3, 0xa5, 0xfa, 0x2a, 0xb1, 0xfb, 0xdb, 0x22, 0x7a, 0x90, 0xf2, 0x32, 0x6e, 0xbc, 0xaa, 0xed,
	0x6e, 0x5c, 0xba, 0x24, 0x1d, 0x9a, 0x9b, 0xc2, 0xe1, 0xc2, 0xcf, 0xcf, 0x3c, 0x9c, 0x73, 0xb3,
	0xd2, 0x62, 0x2e, 0xf3, 0x6e, 0x0e, 0xcc, 0xde, 0x23, 0xc2, 0x2d, 0x4d, 0x50, 0xf5, 0x3f, 0x57,
	0xc8, 0xaf, 0xeb, 0xd1, 0xef, 0x8b, 0x6f, 0xee, 0x0d, 0x06, 0x7f, 0x2c, 0xde, 0xdf, 0x73, 0x96,
	0x83, 0x4c, 0xc5, 0x6e, 0x68, 0x46, 0x93, 0x5e, 0x3c, 0x0e, 0xca, 0x3f, 0x83, 0xe6, 0x64, 0x90,
	0xa9, 0x93, 0x5a, 0x73, 0x32, 0xe9, 0x9d, 0xd4, 0x9a, 0xbf, 0x16, 0x1f, 0xb8, 0x07, 0xfd, 0xfe,
	0x20, 0x53, 0xfd, 0x7e, 0xad, 0xea, 0xf7, 0x27, 0xbd, 0x7e, 0xbf, 0xd6, 0x4d, 0x6f, 0xd8, 0x64,
	0x3f, 0xff, 0x27, 0x00, 0x00, 0xff, 0xff, 0x21, 0x7d, 0xe9, 0xfd, 0xee, 0x0a, 0x00, 0x00,
}
