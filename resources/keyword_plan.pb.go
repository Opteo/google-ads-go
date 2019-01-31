// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v0/resources/keyword_plan.proto

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

// A Keyword Planner plan.
// Max number of saved keyword plans: 10000.
// It's possible to remove plans if limit is reached.
type KeywordPlan struct {
	// The resource name of the Keyword Planner plan.
	// KeywordPlan resource names have the form:
	//
	// `customers/{customer_id}/keywordPlans/{kp_plan_id}`
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// The ID of the keyword plan.
	Id *wrappers.Int64Value `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	// The name of the keyword plan.
	//
	// This field is required and should not be empty when creating new keyword
	// plans.
	Name *wrappers.StringValue `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	// The date period used for forecasting the plan.
	ForecastPeriod       *KeywordPlanForecastPeriod `protobuf:"bytes,4,opt,name=forecast_period,json=forecastPeriod,proto3" json:"forecast_period,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                   `json:"-"`
	XXX_unrecognized     []byte                     `json:"-"`
	XXX_sizecache        int32                      `json:"-"`
}

func (m *KeywordPlan) Reset()         { *m = KeywordPlan{} }
func (m *KeywordPlan) String() string { return proto.CompactTextString(m) }
func (*KeywordPlan) ProtoMessage()    {}
func (*KeywordPlan) Descriptor() ([]byte, []int) {
	return fileDescriptor_1c80fba176ceb86c, []int{0}
}

func (m *KeywordPlan) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KeywordPlan.Unmarshal(m, b)
}
func (m *KeywordPlan) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KeywordPlan.Marshal(b, m, deterministic)
}
func (m *KeywordPlan) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KeywordPlan.Merge(m, src)
}
func (m *KeywordPlan) XXX_Size() int {
	return xxx_messageInfo_KeywordPlan.Size(m)
}
func (m *KeywordPlan) XXX_DiscardUnknown() {
	xxx_messageInfo_KeywordPlan.DiscardUnknown(m)
}

var xxx_messageInfo_KeywordPlan proto.InternalMessageInfo

func (m *KeywordPlan) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func (m *KeywordPlan) GetId() *wrappers.Int64Value {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *KeywordPlan) GetName() *wrappers.StringValue {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *KeywordPlan) GetForecastPeriod() *KeywordPlanForecastPeriod {
	if m != nil {
		return m.ForecastPeriod
	}
	return nil
}

// The forecasting period associated with the keyword plan.
type KeywordPlanForecastPeriod struct {
	// Required. The date used for forecasting the Plan.
	//
	// Types that are valid to be assigned to Interval:
	//	*KeywordPlanForecastPeriod_DateInterval
	//	*KeywordPlanForecastPeriod_DateRange
	Interval             isKeywordPlanForecastPeriod_Interval `protobuf_oneof:"interval"`
	XXX_NoUnkeyedLiteral struct{}                             `json:"-"`
	XXX_unrecognized     []byte                               `json:"-"`
	XXX_sizecache        int32                                `json:"-"`
}

func (m *KeywordPlanForecastPeriod) Reset()         { *m = KeywordPlanForecastPeriod{} }
func (m *KeywordPlanForecastPeriod) String() string { return proto.CompactTextString(m) }
func (*KeywordPlanForecastPeriod) ProtoMessage()    {}
func (*KeywordPlanForecastPeriod) Descriptor() ([]byte, []int) {
	return fileDescriptor_1c80fba176ceb86c, []int{1}
}

func (m *KeywordPlanForecastPeriod) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KeywordPlanForecastPeriod.Unmarshal(m, b)
}
func (m *KeywordPlanForecastPeriod) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KeywordPlanForecastPeriod.Marshal(b, m, deterministic)
}
func (m *KeywordPlanForecastPeriod) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KeywordPlanForecastPeriod.Merge(m, src)
}
func (m *KeywordPlanForecastPeriod) XXX_Size() int {
	return xxx_messageInfo_KeywordPlanForecastPeriod.Size(m)
}
func (m *KeywordPlanForecastPeriod) XXX_DiscardUnknown() {
	xxx_messageInfo_KeywordPlanForecastPeriod.DiscardUnknown(m)
}

var xxx_messageInfo_KeywordPlanForecastPeriod proto.InternalMessageInfo

type isKeywordPlanForecastPeriod_Interval interface {
	isKeywordPlanForecastPeriod_Interval()
}

type KeywordPlanForecastPeriod_DateInterval struct {
	DateInterval enums.KeywordPlanForecastIntervalEnum_KeywordPlanForecastInterval `protobuf:"varint,1,opt,name=date_interval,json=dateInterval,proto3,enum=google.ads.googleads.v0.enums.KeywordPlanForecastIntervalEnum_KeywordPlanForecastInterval,oneof"`
}

type KeywordPlanForecastPeriod_DateRange struct {
	DateRange *common.DateRange `protobuf:"bytes,2,opt,name=date_range,json=dateRange,proto3,oneof"`
}

func (*KeywordPlanForecastPeriod_DateInterval) isKeywordPlanForecastPeriod_Interval() {}

func (*KeywordPlanForecastPeriod_DateRange) isKeywordPlanForecastPeriod_Interval() {}

func (m *KeywordPlanForecastPeriod) GetInterval() isKeywordPlanForecastPeriod_Interval {
	if m != nil {
		return m.Interval
	}
	return nil
}

func (m *KeywordPlanForecastPeriod) GetDateInterval() enums.KeywordPlanForecastIntervalEnum_KeywordPlanForecastInterval {
	if x, ok := m.GetInterval().(*KeywordPlanForecastPeriod_DateInterval); ok {
		return x.DateInterval
	}
	return enums.KeywordPlanForecastIntervalEnum_UNSPECIFIED
}

func (m *KeywordPlanForecastPeriod) GetDateRange() *common.DateRange {
	if x, ok := m.GetInterval().(*KeywordPlanForecastPeriod_DateRange); ok {
		return x.DateRange
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*KeywordPlanForecastPeriod) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*KeywordPlanForecastPeriod_DateInterval)(nil),
		(*KeywordPlanForecastPeriod_DateRange)(nil),
	}
}

func init() {
	proto.RegisterType((*KeywordPlan)(nil), "google.ads.googleads.v0.resources.KeywordPlan")
	proto.RegisterType((*KeywordPlanForecastPeriod)(nil), "google.ads.googleads.v0.resources.KeywordPlanForecastPeriod")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v0/resources/keyword_plan.proto", fileDescriptor_1c80fba176ceb86c)
}

var fileDescriptor_1c80fba176ceb86c = []byte{
	// 473 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x93, 0xdd, 0x6a, 0xd4, 0x40,
	0x14, 0xc7, 0x9b, 0xb4, 0x88, 0x9d, 0x7e, 0x28, 0xb9, 0x5a, 0xab, 0x48, 0x5b, 0x29, 0x54, 0x85,
	0x49, 0xa8, 0xc5, 0x8b, 0xe8, 0x4d, 0x16, 0xb5, 0x1f, 0x82, 0x2c, 0x11, 0xf6, 0xa2, 0x2c, 0x2c,
	0xd3, 0x9d, 0xb3, 0x43, 0x30, 0x99, 0x09, 0x33, 0xc9, 0x16, 0x2f, 0x7d, 0x15, 0x2f, 0x7d, 0x14,
	0x1f, 0xc5, 0x17, 0xd0, 0x1b, 0x41, 0x32, 0x5f, 0xb4, 0xd8, 0x74, 0xef, 0xce, 0xd9, 0xfc, 0xce,
	0xff, 0xfc, 0xcf, 0x39, 0xb3, 0xe8, 0x98, 0x09, 0xc1, 0x4a, 0x88, 0x09, 0x55, 0xb1, 0x09, 0xbb,
	0x68, 0x91, 0xc4, 0x12, 0x94, 0x68, 0xe5, 0x0c, 0x54, 0xfc, 0x05, 0xbe, 0x5e, 0x09, 0x49, 0xa7,
	0x75, 0x49, 0x38, 0xae, 0xa5, 0x68, 0x44, 0xb4, 0x67, 0x50, 0x4c, 0xa8, 0xc2, 0xbe, 0x0a, 0x2f,
	0x12, 0xec, 0xab, 0x76, 0x5e, 0xf4, 0x09, 0xcf, 0x44, 0x55, 0x09, 0x1e, 0x53, 0xd2, 0x80, 0x32,
	0x72, 0x3b, 0xc3, 0x3e, 0x16, 0x78, 0x5b, 0xdd, 0x34, 0x30, 0x9d, 0x0b, 0x09, 0x33, 0xa2, 0x9a,
	0x69, 0xc1, 0x1b, 0x90, 0x0b, 0x52, 0x5a, 0x8d, 0xa7, 0x56, 0x43, 0x67, 0x97, 0xed, 0x3c, 0xbe,
	0x92, 0xa4, 0xae, 0x41, 0xda, 0x1e, 0xfb, 0x7f, 0x02, 0xb4, 0xf1, 0xd1, 0x08, 0x8d, 0x4a, 0xc2,
	0xa3, 0x67, 0x68, 0xcb, 0x99, 0x9d, 0x72, 0x52, 0xc1, 0x20, 0xd8, 0x0d, 0x0e, 0xd7, 0xf3, 0x4d,
	0xf7, 0xe3, 0x27, 0x52, 0x41, 0xf4, 0x12, 0x85, 0x05, 0x1d, 0x84, 0xbb, 0xc1, 0xe1, 0xc6, 0xd1,
	0x63, 0x3b, 0x29, 0x76, 0x1d, 0xf0, 0x19, 0x6f, 0x5e, 0x1f, 0x8f, 0x49, 0xd9, 0x42, 0x1e, 0x16,
	0x34, 0x4a, 0xd0, 0x9a, 0x16, 0x5a, 0xd5, 0xf8, 0x93, 0xff, 0xf0, 0xcf, 0x8d, 0x2c, 0x38, 0x33,
	0xbc, 0x26, 0x23, 0x40, 0x0f, 0xfc, 0x38, 0x35, 0xc8, 0x42, 0xd0, 0xc1, 0x9a, 0x2e, 0x7e, 0x8b,
	0x97, 0x2e, 0x18, 0x5f, 0x1b, 0xe6, 0x83, 0x15, 0x19, 0x69, 0x8d, 0x7c, 0x7b, 0x7e, 0x23, 0xdf,
	0xff, 0x1d, 0xa0, 0x47, 0xbd, 0x74, 0xf4, 0x2d, 0x40, 0x5b, 0xdd, 0x31, 0xfc, 0x42, 0xf5, 0x26,
	0xb6, 0x8f, 0x2e, 0x7a, 0x3d, 0xe8, 0xab, 0xdc, 0xd6, 0xff, 0xcc, 0x2a, 0xbc, 0xe7, 0x6d, 0x75,
	0xd7, 0xf7, 0xd3, 0x95, 0x7c, 0xb3, 0x6b, 0xe9, 0xf2, 0xe8, 0x1c, 0x21, 0x6d, 0x41, 0x12, 0xce,
	0xc0, 0xee, 0xfb, 0x79, 0x6f, 0x7f, 0xf3, 0x82, 0xf0, 0x3b, 0xd2, 0x40, 0xde, 0x15, 0x9c, 0xae,
	0xe4, 0xeb, 0xd4, 0x25, 0x43, 0x84, 0xee, 0xbb, 0x49, 0x86, 0x7f, 0x03, 0x74, 0x30, 0x13, 0xd5,
	0xf2, 0x6d, 0x0e, 0x1f, 0x5e, 0xb3, 0x3b, 0xea, 0x2e, 0x36, 0x0a, 0x2e, 0xce, 0x6d, 0x19, 0x13,
	0x25, 0xe1, 0x0c, 0x0b, 0xc9, 0x62, 0x06, 0x5c, 0xdf, 0xd3, 0x3d, 0xd3, 0xba, 0x50, 0x77, 0xfc,
	0x75, 0xde, 0xf8, 0xe8, 0x7b, 0xb8, 0x7a, 0x92, 0x65, 0x3f, 0xc2, 0xbd, 0x13, 0x23, 0x99, 0x51,
	0x85, 0x4d, 0xd8, 0x45, 0xe3, 0x04, 0xe7, 0x8e, 0xfc, 0xe9, 0x98, 0x49, 0x46, 0xd5, 0xc4, 0x33,
	0x93, 0x71, 0x32, 0xf1, 0xcc, 0xaf, 0xf0, 0xc0, 0x7c, 0x48, 0xd3, 0x8c, 0xaa, 0x34, 0xf5, 0x54,
	0x9a, 0x8e, 0x93, 0x34, 0xf5, 0xdc, 0xe5, 0x3d, 0x6d, 0xf6, 0xd5, 0xbf, 0x00, 0x00, 0x00, 0xff,
	0xff, 0xc4, 0xd4, 0x18, 0x9c, 0xe6, 0x03, 0x00, 0x00,
}
