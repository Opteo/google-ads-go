// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v0/resources/campaign_criterion.proto

package google_ads_googleads_v0_resources

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	common "google/ads/googleads/v0/common"
	enums "google/ads/googleads/v0/enums"
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

// A campaign criterion.
type CampaignCriterion struct {
	// The resource name of the campaign criterion.
	// Campaign criterion resource names have the form:
	//
	// `customers/{customer_id}/campaignCriteria/{campaign_id}_{criterion_id}
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// The campaign to which the criterion belongs.
	//
	// This field must not be used in WHERE clauses.
	Campaign *wrappers.StringValue `protobuf:"bytes,4,opt,name=campaign,proto3" json:"campaign,omitempty"`
	// The ID of the criterion.
	//
	// This field is ignored during mutate.
	CriterionId *wrappers.Int64Value `protobuf:"bytes,5,opt,name=criterion_id,json=criterionId,proto3" json:"criterion_id,omitempty"`
	// The modifier for the bids when the criterion matches.
	// Allowable modifier values depend on the criterion:
	//  - 0.1 - 10.0: Location
	//  - 0.1 - 4.0: Platform (mobile). Use 0 to opt out of mobile.
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
	//	*CampaignCriterion_Location
	//	*CampaignCriterion_Platform
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

type CampaignCriterion_Location struct {
	Location *common.LocationInfo `protobuf:"bytes,12,opt,name=location,proto3,oneof"`
}

type CampaignCriterion_Platform struct {
	Platform *common.PlatformInfo `protobuf:"bytes,13,opt,name=platform,proto3,oneof"`
}

func (*CampaignCriterion_Keyword) isCampaignCriterion_Criterion() {}

func (*CampaignCriterion_Location) isCampaignCriterion_Criterion() {}

func (*CampaignCriterion_Platform) isCampaignCriterion_Criterion() {}

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

func (m *CampaignCriterion) GetLocation() *common.LocationInfo {
	if x, ok := m.GetCriterion().(*CampaignCriterion_Location); ok {
		return x.Location
	}
	return nil
}

func (m *CampaignCriterion) GetPlatform() *common.PlatformInfo {
	if x, ok := m.GetCriterion().(*CampaignCriterion_Platform); ok {
		return x.Platform
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*CampaignCriterion) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _CampaignCriterion_OneofMarshaler, _CampaignCriterion_OneofUnmarshaler, _CampaignCriterion_OneofSizer, []interface{}{
		(*CampaignCriterion_Keyword)(nil),
		(*CampaignCriterion_Location)(nil),
		(*CampaignCriterion_Platform)(nil),
	}
}

func _CampaignCriterion_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*CampaignCriterion)
	// criterion
	switch x := m.Criterion.(type) {
	case *CampaignCriterion_Keyword:
		b.EncodeVarint(8<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Keyword); err != nil {
			return err
		}
	case *CampaignCriterion_Location:
		b.EncodeVarint(12<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Location); err != nil {
			return err
		}
	case *CampaignCriterion_Platform:
		b.EncodeVarint(13<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Platform); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("CampaignCriterion.Criterion has unexpected type %T", x)
	}
	return nil
}

func _CampaignCriterion_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*CampaignCriterion)
	switch tag {
	case 8: // criterion.keyword
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(common.KeywordInfo)
		err := b.DecodeMessage(msg)
		m.Criterion = &CampaignCriterion_Keyword{msg}
		return true, err
	case 12: // criterion.location
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(common.LocationInfo)
		err := b.DecodeMessage(msg)
		m.Criterion = &CampaignCriterion_Location{msg}
		return true, err
	case 13: // criterion.platform
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(common.PlatformInfo)
		err := b.DecodeMessage(msg)
		m.Criterion = &CampaignCriterion_Platform{msg}
		return true, err
	default:
		return false, nil
	}
}

func _CampaignCriterion_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*CampaignCriterion)
	// criterion
	switch x := m.Criterion.(type) {
	case *CampaignCriterion_Keyword:
		s := proto.Size(x.Keyword)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *CampaignCriterion_Location:
		s := proto.Size(x.Location)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *CampaignCriterion_Platform:
		s := proto.Size(x.Platform)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

func init() {
	proto.RegisterType((*CampaignCriterion)(nil), "google.ads.googleads.v0.resources.CampaignCriterion")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v0/resources/campaign_criterion.proto", fileDescriptor_889de17d165924ac)
}

var fileDescriptor_889de17d165924ac = []byte{
	// 478 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x93, 0xdd, 0x8a, 0xd3, 0x40,
	0x14, 0xc7, 0xcd, 0x5a, 0x77, 0xbb, 0xd3, 0xee, 0x82, 0xb9, 0x90, 0x50, 0x45, 0xba, 0x8a, 0x50,
	0x50, 0x27, 0xa5, 0xca, 0x22, 0x5e, 0x2c, 0xb4, 0x8b, 0xd6, 0xfa, 0x45, 0x89, 0xd2, 0xab, 0x42,
	0x99, 0x66, 0xa6, 0x61, 0x30, 0x33, 0x27, 0x4c, 0x92, 0x2e, 0x7d, 0x1d, 0x2f, 0xbd, 0xf5, 0x2d,
	0x7c, 0x0c, 0xaf, 0x7d, 0x08, 0x49, 0xe6, 0x43, 0xa4, 0x84, 0xdd, 0xbb, 0x93, 0xc9, 0xff, 0xf7,
	0xcb, 0x99, 0x39, 0x13, 0xf4, 0x3a, 0x01, 0x48, 0x52, 0x16, 0x12, 0x9a, 0x87, 0xba, 0xac, 0xaa,
	0xed, 0x30, 0x54, 0x2c, 0x87, 0x52, 0xc5, 0x2c, 0x0f, 0x63, 0x22, 0x32, 0xc2, 0x13, 0xb9, 0x8a,
	0x15, 0x2f, 0x98, 0xe2, 0x20, 0x71, 0xa6, 0xa0, 0x00, 0xff, 0x4c, 0x03, 0x98, 0xd0, 0x1c, 0x3b,
	0x16, 0x6f, 0x87, 0xd8, 0xb1, 0xbd, 0xe7, 0x4d, 0xfa, 0x18, 0x84, 0x00, 0x19, 0x1a, 0x25, 0xd1,
	0xc6, 0xde, 0xa8, 0x29, 0xce, 0x64, 0x29, 0xf2, 0xd0, 0x35, 0xb0, 0x2a, 0x76, 0x19, 0x33, 0xcc,
	0x43, 0xc3, 0xd4, 0x4f, 0xeb, 0x72, 0x13, 0x5e, 0x29, 0x92, 0x65, 0x4c, 0xe5, 0xfa, 0xfd, 0xa3,
	0x3f, 0x2d, 0x74, 0xf7, 0xd2, 0x6c, 0xe1, 0xd2, 0x0a, 0xfc, 0xc7, 0xe8, 0xc4, 0x76, 0xb9, 0x92,
	0x44, 0xb0, 0xc0, 0xeb, 0x7b, 0x83, 0xe3, 0xa8, 0x6b, 0x17, 0x3f, 0x13, 0xc1, 0xfc, 0x57, 0xa8,
	0x6d, 0x37, 0x1f, 0xb4, 0xfa, 0xde, 0xa0, 0x33, 0x7a, 0x60, 0x36, 0x8a, 0xed, 0xd7, 0xf0, 0x97,
	0x42, 0x71, 0x99, 0x2c, 0x48, 0x5a, 0xb2, 0xc8, 0xa5, 0xfd, 0x0b, 0xd4, 0xfd, 0xd7, 0x2c, 0xa7,
	0xc1, 0x9d, 0x9a, 0xbe, 0xbf, 0x47, 0xcf, 0x64, 0x71, 0xfe, 0x52, 0xc3, 0x1d, 0x07, 0xcc, 0x68,
	0xc5, 0xaf, 0x39, 0x5d, 0x09, 0xa0, 0x7c, 0xc3, 0x99, 0x0a, 0x4e, 0x1b, 0xf8, 0xb7, 0x29, 0x90,
	0xc2, 0xf0, 0x6b, 0x4e, 0x3f, 0x99, 0xbc, 0x7f, 0x8e, 0xda, 0x92, 0x25, 0xa4, 0xe0, 0x5b, 0x16,
	0x1c, 0xd5, 0x6c, 0x6f, 0x8f, 0x9d, 0x00, 0xa4, 0xa6, 0x6f, 0x9b, 0xf5, 0x23, 0xd4, 0xaa, 0x8e,
	0x36, 0x38, 0xec, 0x7b, 0x83, 0xd3, 0xd1, 0x05, 0x6e, 0x9a, 0x70, 0x3d, 0x0f, 0xec, 0x8e, 0xf3,
	0xeb, 0x2e, 0x63, 0x6f, 0x64, 0x29, 0xfe, 0x5f, 0x89, 0x6a, 0x97, 0x3f, 0x45, 0x47, 0xdf, 0xd8,
	0xee, 0x0a, 0x14, 0x0d, 0xda, 0x75, 0x2b, 0x4f, 0x1b, 0xb5, 0xfa, 0x56, 0xe0, 0x0f, 0x3a, 0x3e,
	0x93, 0x1b, 0x78, 0x77, 0x2b, 0xb2, 0xb4, 0xff, 0x1e, 0xb5, 0x53, 0x88, 0x49, 0xc1, 0x41, 0x06,
	0xdd, 0xda, 0xf4, 0xec, 0x3a, 0xd3, 0x47, 0x93, 0x37, 0x2a, 0xc7, 0x57, 0xae, 0x2c, 0x25, 0xc5,
	0x06, 0x94, 0x08, 0x4e, 0x6e, 0xe6, 0x9a, 0x9b, 0xbc, 0x75, 0x59, 0x7e, 0xd2, 0x41, 0xc7, 0x6e,
	0x76, 0x93, 0x9f, 0x1e, 0x7a, 0x12, 0x83, 0xc0, 0xd7, 0xfe, 0x1b, 0x93, 0x7b, 0x7b, 0xb7, 0x72,
	0x5e, 0x8d, 0x66, 0xee, 0x7d, 0x3f, 0xb8, 0x3d, 0x1d, 0x8f, 0x7f, 0x1c, 0x9c, 0x4d, 0xb5, 0x63,
	0x4c, 0x73, 0xac, 0xcb, 0xaa, 0x5a, 0x0c, 0x71, 0x64, 0x1d, 0xbf, 0x6c, 0x66, 0x39, 0xa6, 0xf9,
	0xd2, 0x65, 0x96, 0x8b, 0xe1, 0xd2, 0x65, 0x7e, 0xdf, 0x20, 0xb3, 0x3e, 0xac, 0x6f, 0xc5, 0x8b,
	0xbf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x65, 0xc2, 0x1d, 0x31, 0x0f, 0x04, 0x00, 0x00,
}