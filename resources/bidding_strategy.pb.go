// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v0/resources/bidding_strategy.proto

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

// A bidding strategy.
type BiddingStrategy struct {
	// The resource name of the bidding strategy.
	// Bidding strategy resource names have the form:
	//
	// `customers/{customer_id}/biddingStrategies/{bidding_strategy_id}`
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// The ID of the bidding strategy.
	Id *wrappers.Int64Value `protobuf:"bytes,3,opt,name=id,proto3" json:"id,omitempty"`
	// The name of the bidding strategy.
	// All bidding strategies within an account must be named distinctly.
	//
	// The length of this string should be between 1 and 255, inclusive,
	// in UTF-8 bytes, (trimmed).
	Name *wrappers.StringValue `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	// The type of the bidding strategy.
	// Create a bidding strategy by setting the bidding scheme.
	//
	// This field is read-only.
	Type enums.BiddingStrategyTypeEnum_BiddingStrategyType `protobuf:"varint,5,opt,name=type,proto3,enum=google.ads.googleads.v0.enums.BiddingStrategyTypeEnum_BiddingStrategyType" json:"type,omitempty"`
	// The bidding scheme.
	//
	// Only one can be set.
	//
	// Types that are valid to be assigned to Scheme:
	//	*BiddingStrategy_EnhancedCpc
	//	*BiddingStrategy_PageOnePromoted
	//	*BiddingStrategy_TargetCpa
	//	*BiddingStrategy_TargetOutrankShare
	//	*BiddingStrategy_TargetRoas
	//	*BiddingStrategy_TargetSpend
	Scheme               isBiddingStrategy_Scheme `protobuf_oneof:"scheme"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *BiddingStrategy) Reset()         { *m = BiddingStrategy{} }
func (m *BiddingStrategy) String() string { return proto.CompactTextString(m) }
func (*BiddingStrategy) ProtoMessage()    {}
func (*BiddingStrategy) Descriptor() ([]byte, []int) {
	return fileDescriptor_954492da6a94ea83, []int{0}
}

func (m *BiddingStrategy) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BiddingStrategy.Unmarshal(m, b)
}
func (m *BiddingStrategy) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BiddingStrategy.Marshal(b, m, deterministic)
}
func (m *BiddingStrategy) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BiddingStrategy.Merge(m, src)
}
func (m *BiddingStrategy) XXX_Size() int {
	return xxx_messageInfo_BiddingStrategy.Size(m)
}
func (m *BiddingStrategy) XXX_DiscardUnknown() {
	xxx_messageInfo_BiddingStrategy.DiscardUnknown(m)
}

var xxx_messageInfo_BiddingStrategy proto.InternalMessageInfo

func (m *BiddingStrategy) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func (m *BiddingStrategy) GetId() *wrappers.Int64Value {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *BiddingStrategy) GetName() *wrappers.StringValue {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *BiddingStrategy) GetType() enums.BiddingStrategyTypeEnum_BiddingStrategyType {
	if m != nil {
		return m.Type
	}
	return enums.BiddingStrategyTypeEnum_UNSPECIFIED
}

type isBiddingStrategy_Scheme interface {
	isBiddingStrategy_Scheme()
}

type BiddingStrategy_EnhancedCpc struct {
	EnhancedCpc *common.EnhancedCpc `protobuf:"bytes,7,opt,name=enhanced_cpc,json=enhancedCpc,proto3,oneof"`
}

type BiddingStrategy_PageOnePromoted struct {
	PageOnePromoted *common.PageOnePromoted `protobuf:"bytes,8,opt,name=page_one_promoted,json=pageOnePromoted,proto3,oneof"`
}

type BiddingStrategy_TargetCpa struct {
	TargetCpa *common.TargetCpa `protobuf:"bytes,9,opt,name=target_cpa,json=targetCpa,proto3,oneof"`
}

type BiddingStrategy_TargetOutrankShare struct {
	TargetOutrankShare *common.TargetOutrankShare `protobuf:"bytes,10,opt,name=target_outrank_share,json=targetOutrankShare,proto3,oneof"`
}

type BiddingStrategy_TargetRoas struct {
	TargetRoas *common.TargetRoas `protobuf:"bytes,11,opt,name=target_roas,json=targetRoas,proto3,oneof"`
}

type BiddingStrategy_TargetSpend struct {
	TargetSpend *common.TargetSpend `protobuf:"bytes,12,opt,name=target_spend,json=targetSpend,proto3,oneof"`
}

func (*BiddingStrategy_EnhancedCpc) isBiddingStrategy_Scheme() {}

func (*BiddingStrategy_PageOnePromoted) isBiddingStrategy_Scheme() {}

func (*BiddingStrategy_TargetCpa) isBiddingStrategy_Scheme() {}

func (*BiddingStrategy_TargetOutrankShare) isBiddingStrategy_Scheme() {}

func (*BiddingStrategy_TargetRoas) isBiddingStrategy_Scheme() {}

func (*BiddingStrategy_TargetSpend) isBiddingStrategy_Scheme() {}

func (m *BiddingStrategy) GetScheme() isBiddingStrategy_Scheme {
	if m != nil {
		return m.Scheme
	}
	return nil
}

func (m *BiddingStrategy) GetEnhancedCpc() *common.EnhancedCpc {
	if x, ok := m.GetScheme().(*BiddingStrategy_EnhancedCpc); ok {
		return x.EnhancedCpc
	}
	return nil
}

func (m *BiddingStrategy) GetPageOnePromoted() *common.PageOnePromoted {
	if x, ok := m.GetScheme().(*BiddingStrategy_PageOnePromoted); ok {
		return x.PageOnePromoted
	}
	return nil
}

func (m *BiddingStrategy) GetTargetCpa() *common.TargetCpa {
	if x, ok := m.GetScheme().(*BiddingStrategy_TargetCpa); ok {
		return x.TargetCpa
	}
	return nil
}

func (m *BiddingStrategy) GetTargetOutrankShare() *common.TargetOutrankShare {
	if x, ok := m.GetScheme().(*BiddingStrategy_TargetOutrankShare); ok {
		return x.TargetOutrankShare
	}
	return nil
}

func (m *BiddingStrategy) GetTargetRoas() *common.TargetRoas {
	if x, ok := m.GetScheme().(*BiddingStrategy_TargetRoas); ok {
		return x.TargetRoas
	}
	return nil
}

func (m *BiddingStrategy) GetTargetSpend() *common.TargetSpend {
	if x, ok := m.GetScheme().(*BiddingStrategy_TargetSpend); ok {
		return x.TargetSpend
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*BiddingStrategy) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*BiddingStrategy_EnhancedCpc)(nil),
		(*BiddingStrategy_PageOnePromoted)(nil),
		(*BiddingStrategy_TargetCpa)(nil),
		(*BiddingStrategy_TargetOutrankShare)(nil),
		(*BiddingStrategy_TargetRoas)(nil),
		(*BiddingStrategy_TargetSpend)(nil),
	}
}

func init() {
	proto.RegisterType((*BiddingStrategy)(nil), "google.ads.googleads.v0.resources.BiddingStrategy")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v0/resources/bidding_strategy.proto", fileDescriptor_954492da6a94ea83)
}

var fileDescriptor_954492da6a94ea83 = []byte{
	// 574 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x94, 0xd1, 0x4e, 0xd4, 0x4c,
	0x14, 0xc7, 0x77, 0x0b, 0x1f, 0x1f, 0xcc, 0xa2, 0xc4, 0x86, 0x8b, 0x06, 0x8d, 0x01, 0x0d, 0x09,
	0x8a, 0x99, 0x36, 0x68, 0x8c, 0xd6, 0xab, 0x5d, 0x42, 0x58, 0x49, 0x94, 0x4d, 0x97, 0xec, 0x85,
	0x59, 0x6d, 0x86, 0xce, 0x61, 0x68, 0xa4, 0x33, 0x93, 0x99, 0x29, 0x86, 0x4b, 0x2f, 0x7c, 0x11,
	0x2f, 0x7d, 0x14, 0x1f, 0xc5, 0xa7, 0x30, 0x9d, 0x69, 0x4b, 0x04, 0x57, 0xb8, 0x3b, 0xe7, 0xf4,
	0xff, 0xff, 0xcd, 0x39, 0x3d, 0xd3, 0xa2, 0x57, 0x4c, 0x08, 0x76, 0x06, 0x21, 0xa1, 0x3a, 0x74,
	0x61, 0x15, 0x9d, 0x47, 0xa1, 0x02, 0x2d, 0x4a, 0x95, 0x81, 0x0e, 0x8f, 0x73, 0x4a, 0x73, 0xce,
	0x52, 0x6d, 0x14, 0x31, 0xc0, 0x2e, 0xb0, 0x54, 0xc2, 0x08, 0x7f, 0xc3, 0xc9, 0x31, 0xa1, 0x1a,
	0xb7, 0x4e, 0x7c, 0x1e, 0xe1, 0xd6, 0xb9, 0xf6, 0x6c, 0x16, 0x3c, 0x13, 0x45, 0x21, 0x78, 0x43,
	0x76, 0xc0, 0xb5, 0xd7, 0xb3, 0xd4, 0xc0, 0xcb, 0xe2, 0x7a, 0x1b, 0xa9, 0xb9, 0x90, 0x50, 0x5b,
	0x1f, 0xd6, 0x56, 0x9b, 0x1d, 0x97, 0x27, 0xe1, 0x17, 0x45, 0xa4, 0x04, 0xa5, 0xdd, 0xf3, 0x47,
	0xdf, 0x16, 0xd0, 0xca, 0xc0, 0xf9, 0xc7, 0xb5, 0xdd, 0x7f, 0x8c, 0xee, 0x34, 0x9d, 0xa6, 0x9c,
	0x14, 0x10, 0x74, 0xd7, 0xbb, 0x5b, 0x4b, 0xc9, 0x72, 0x53, 0x7c, 0x4f, 0x0a, 0xf0, 0xb7, 0x91,
	0x97, 0xd3, 0x60, 0x6e, 0xbd, 0xbb, 0xd5, 0xdb, 0xb9, 0x5f, 0x8f, 0x89, 0x9b, 0x53, 0xf0, 0x5b,
	0x6e, 0x5e, 0xbe, 0x98, 0x90, 0xb3, 0x12, 0x12, 0x2f, 0xa7, 0x7e, 0x84, 0xe6, 0x2d, 0x68, 0xde,
	0xca, 0x1f, 0x5c, 0x93, 0x8f, 0x8d, 0xca, 0x39, 0x73, 0x7a, 0xab, 0xf4, 0x3f, 0xa1, 0xf9, 0x6a,
	0x8a, 0xe0, 0xbf, 0xf5, 0xee, 0xd6, 0xdd, 0x9d, 0x03, 0x3c, 0xeb, 0x95, 0xda, 0x37, 0x80, 0xaf,
	0x4c, 0x70, 0x74, 0x21, 0x61, 0x8f, 0x97, 0xc5, 0xdf, 0xea, 0x89, 0xe5, 0xfa, 0x23, 0xb4, 0x0c,
	0xfc, 0x94, 0xf0, 0x0c, 0x68, 0x9a, 0xc9, 0x2c, 0xf8, 0xdf, 0x76, 0xb6, 0x3d, 0xf3, 0x1c, 0xb7,
	0x17, 0xbc, 0x57, 0x7b, 0x76, 0x65, 0x36, 0xec, 0x24, 0x3d, 0xb8, 0x4c, 0xfd, 0x8f, 0xe8, 0x9e,
	0x24, 0x0c, 0x52, 0xc1, 0x21, 0x95, 0x4a, 0x14, 0xc2, 0x00, 0x0d, 0x16, 0x2d, 0x36, 0xbc, 0x09,
	0x3b, 0x22, 0x0c, 0x0e, 0x39, 0x8c, 0x6a, 0xdb, 0xb0, 0x93, 0xac, 0xc8, 0x3f, 0x4b, 0xfe, 0x01,
	0x42, 0x86, 0x28, 0x06, 0x26, 0xcd, 0x24, 0x09, 0x96, 0x2c, 0xf7, 0xc9, 0x4d, 0xdc, 0x23, 0xeb,
	0xd8, 0x95, 0x64, 0xd8, 0x49, 0x96, 0x4c, 0x93, 0xf8, 0x27, 0x68, 0xb5, 0x66, 0x89, 0xd2, 0x28,
	0xc2, 0x3f, 0xa7, 0xfa, 0x94, 0x28, 0x08, 0x90, 0xa5, 0xee, 0xdc, 0x8e, 0x7a, 0xe8, 0xac, 0xe3,
	0xca, 0x39, 0xec, 0x24, 0xbe, 0xb9, 0x56, 0xf5, 0xdf, 0xa1, 0x5e, 0x7d, 0x8e, 0x12, 0x44, 0x07,
	0x3d, 0x8b, 0x7f, 0x7a, 0x3b, 0x7c, 0x22, 0x88, 0x1e, 0x76, 0x92, 0x7a, 0xe8, 0x2a, 0xab, 0x76,
	0x56, 0xe3, 0xb4, 0x04, 0x4e, 0x83, 0xe5, 0xdb, 0xed, 0xcc, 0xf1, 0xc6, 0x95, 0xa5, 0xda, 0x99,
	0xb9, 0x4c, 0x07, 0x8b, 0x68, 0x41, 0x67, 0xa7, 0x50, 0xc0, 0xe0, 0xab, 0x87, 0x36, 0x33, 0x51,
	0xe0, 0x1b, 0x3f, 0xdd, 0xc1, 0xea, 0x95, 0x4b, 0x35, 0xaa, 0x2e, 0xf1, 0xa8, 0xfb, 0xe1, 0xa0,
	0xb6, 0x32, 0x71, 0x46, 0x38, 0xc3, 0x42, 0xb1, 0x90, 0x01, 0xb7, 0x57, 0xbc, 0xf9, 0x68, 0x65,
	0xae, 0xff, 0xf1, 0x3b, 0x79, 0xd3, 0x46, 0xdf, 0xbd, 0xb9, 0xfd, 0x7e, 0xff, 0x87, 0xb7, 0xb1,
	0xef, 0x90, 0x7d, 0xaa, 0xb1, 0x0b, 0xab, 0x68, 0x12, 0xe1, 0xa4, 0x51, 0xfe, 0x6c, 0x34, 0xd3,
	0x3e, 0xd5, 0xd3, 0x56, 0x33, 0x9d, 0x44, 0xd3, 0x56, 0xf3, 0xcb, 0xdb, 0x74, 0x0f, 0xe2, 0xb8,
	0x4f, 0x75, 0x1c, 0xb7, 0xaa, 0x38, 0x9e, 0x44, 0x71, 0xdc, 0xea, 0x8e, 0x17, 0x6c, 0xb3, 0xcf,
	0x7f, 0x07, 0x00, 0x00, 0xff, 0xff, 0x21, 0x93, 0x56, 0xa6, 0xfa, 0x04, 0x00, 0x00,
}
