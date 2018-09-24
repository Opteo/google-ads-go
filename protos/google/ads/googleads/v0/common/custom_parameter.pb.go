// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v0/common/custom_parameter.proto

package google_ads_googleads_v0_common

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
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

// A mapping that can be used by custom parameter tags in a
// `tracking_url_template`, `final_urls`, or `mobile_final_urls`.
type CustomParameter struct {
	// The key matching the parameter tag name.
	Key *wrappers.StringValue `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	// The value to be substituted.
	Value                *wrappers.StringValue `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *CustomParameter) Reset()         { *m = CustomParameter{} }
func (m *CustomParameter) String() string { return proto.CompactTextString(m) }
func (*CustomParameter) ProtoMessage()    {}
func (*CustomParameter) Descriptor() ([]byte, []int) {
	return fileDescriptor_77e235fc7c2795f5, []int{0}
}
func (m *CustomParameter) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CustomParameter.Unmarshal(m, b)
}
func (m *CustomParameter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CustomParameter.Marshal(b, m, deterministic)
}
func (m *CustomParameter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CustomParameter.Merge(m, src)
}
func (m *CustomParameter) XXX_Size() int {
	return xxx_messageInfo_CustomParameter.Size(m)
}
func (m *CustomParameter) XXX_DiscardUnknown() {
	xxx_messageInfo_CustomParameter.DiscardUnknown(m)
}

var xxx_messageInfo_CustomParameter proto.InternalMessageInfo

func (m *CustomParameter) GetKey() *wrappers.StringValue {
	if m != nil {
		return m.Key
	}
	return nil
}

func (m *CustomParameter) GetValue() *wrappers.StringValue {
	if m != nil {
		return m.Value
	}
	return nil
}

func init() {
	proto.RegisterType((*CustomParameter)(nil), "google.ads.googleads.v0.common.CustomParameter")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v0/common/custom_parameter.proto", fileDescriptor_77e235fc7c2795f5)
}

var fileDescriptor_77e235fc7c2795f5 = []byte{
	// 240 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x32, 0x4d, 0xcf, 0xcf, 0x4f,
	0xcf, 0x49, 0xd5, 0x4f, 0x4c, 0x29, 0xd6, 0x87, 0x30, 0x41, 0xac, 0x32, 0x03, 0xfd, 0xe4, 0xfc,
	0xdc, 0xdc, 0xfc, 0x3c, 0xfd, 0xe4, 0xd2, 0xe2, 0x92, 0xfc, 0xdc, 0xf8, 0x82, 0xc4, 0xa2, 0xc4,
	0xdc, 0xd4, 0x92, 0xd4, 0x22, 0xbd, 0x82, 0xa2, 0xfc, 0x92, 0x7c, 0x21, 0x39, 0x88, 0x5a, 0xbd,
	0xc4, 0x94, 0x62, 0x3d, 0xb8, 0x36, 0xbd, 0x32, 0x03, 0x3d, 0x88, 0x36, 0x29, 0xa8, 0xbc, 0x3e,
	0x58, 0x75, 0x52, 0x69, 0x9a, 0x7e, 0x79, 0x51, 0x62, 0x41, 0x41, 0x6a, 0x51, 0x31, 0x44, 0xbf,
	0x52, 0x29, 0x17, 0xbf, 0x33, 0xd8, 0xe4, 0x00, 0x98, 0xc1, 0x42, 0x7a, 0x5c, 0xcc, 0xd9, 0xa9,
	0x95, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0xdc, 0x46, 0x32, 0x50, 0x53, 0xf5, 0x60, 0x06, 0xe8, 0x05,
	0x97, 0x14, 0x65, 0xe6, 0xa5, 0x87, 0x25, 0xe6, 0x94, 0xa6, 0x06, 0x81, 0x14, 0x0a, 0x19, 0x71,
	0xb1, 0x96, 0x81, 0x78, 0x12, 0x4c, 0x44, 0xe8, 0x80, 0x28, 0x75, 0x5a, 0xca, 0xc8, 0xa5, 0x94,
	0x9c, 0x9f, 0xab, 0x87, 0xdf, 0xf5, 0x4e, 0x22, 0x68, 0x6e, 0x0b, 0x00, 0x19, 0x19, 0xc0, 0xb8,
	0x88, 0x89, 0xd9, 0xdd, 0xd1, 0x71, 0x15, 0x93, 0x9c, 0x3b, 0x44, 0xbb, 0x63, 0x4a, 0xb1, 0x1e,
	0x84, 0x09, 0x62, 0x85, 0x19, 0xe8, 0x39, 0x83, 0xb5, 0x9f, 0x82, 0x29, 0x88, 0x71, 0x4c, 0x29,
	0x8e, 0x81, 0x2b, 0x88, 0x09, 0x33, 0x88, 0x81, 0x28, 0x78, 0x44, 0x48, 0x41, 0x12, 0x1b, 0xd8,
	0x13, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0x28, 0xa6, 0x08, 0x0b, 0x9e, 0x01, 0x00, 0x00,
}