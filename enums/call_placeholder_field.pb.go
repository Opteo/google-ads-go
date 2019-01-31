// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v0/enums/call_placeholder_field.proto

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

// Possible values for Call placeholder fields.
type CallPlaceholderFieldEnum_CallPlaceholderField int32

const (
	// Not specified.
	CallPlaceholderFieldEnum_UNSPECIFIED CallPlaceholderFieldEnum_CallPlaceholderField = 0
	// Used for return value only. Represents value unknown in this version.
	CallPlaceholderFieldEnum_UNKNOWN CallPlaceholderFieldEnum_CallPlaceholderField = 1
	// Data Type: STRING. The advertiser's phone number to append to the ad.
	CallPlaceholderFieldEnum_PHONE_NUMBER CallPlaceholderFieldEnum_CallPlaceholderField = 2
	// Data Type: STRING. Uppercase two-letter country code of the advertiser's
	// phone number.
	CallPlaceholderFieldEnum_COUNTRY_CODE CallPlaceholderFieldEnum_CallPlaceholderField = 3
	// Data Type: BOOLEAN. Indicates whether call tracking is enabled. Default:
	// true.
	CallPlaceholderFieldEnum_TRACKED CallPlaceholderFieldEnum_CallPlaceholderField = 4
	// Data Type: INT64. The ID of an AdCallMetricsConversion object. This
	// object contains the phoneCallDurationfield which is the minimum duration
	// (in seconds) of a call to be considered a conversion.
	CallPlaceholderFieldEnum_CONVERSION_TYPE_ID CallPlaceholderFieldEnum_CallPlaceholderField = 5
	// Data Type: STRING. Indicates whether this call extension uses its own
	// call conversion setting or follows the account level setting.
	// Valid values are: USE_ACCOUNT_LEVEL_CALL_CONVERSION_ACTION and
	// USE_RESOURCE_LEVEL_CALL_CONVERSION_ACTION.
	CallPlaceholderFieldEnum_CONVERSION_REPORTING_STATE CallPlaceholderFieldEnum_CallPlaceholderField = 6
)

var CallPlaceholderFieldEnum_CallPlaceholderField_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "UNKNOWN",
	2: "PHONE_NUMBER",
	3: "COUNTRY_CODE",
	4: "TRACKED",
	5: "CONVERSION_TYPE_ID",
	6: "CONVERSION_REPORTING_STATE",
}

var CallPlaceholderFieldEnum_CallPlaceholderField_value = map[string]int32{
	"UNSPECIFIED":                0,
	"UNKNOWN":                    1,
	"PHONE_NUMBER":               2,
	"COUNTRY_CODE":               3,
	"TRACKED":                    4,
	"CONVERSION_TYPE_ID":         5,
	"CONVERSION_REPORTING_STATE": 6,
}

func (x CallPlaceholderFieldEnum_CallPlaceholderField) String() string {
	return proto.EnumName(CallPlaceholderFieldEnum_CallPlaceholderField_name, int32(x))
}

func (CallPlaceholderFieldEnum_CallPlaceholderField) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_1d28e3249364eca5, []int{0, 0}
}

// Values for Call placeholder fields.
type CallPlaceholderFieldEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CallPlaceholderFieldEnum) Reset()         { *m = CallPlaceholderFieldEnum{} }
func (m *CallPlaceholderFieldEnum) String() string { return proto.CompactTextString(m) }
func (*CallPlaceholderFieldEnum) ProtoMessage()    {}
func (*CallPlaceholderFieldEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_1d28e3249364eca5, []int{0}
}

func (m *CallPlaceholderFieldEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CallPlaceholderFieldEnum.Unmarshal(m, b)
}
func (m *CallPlaceholderFieldEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CallPlaceholderFieldEnum.Marshal(b, m, deterministic)
}
func (m *CallPlaceholderFieldEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CallPlaceholderFieldEnum.Merge(m, src)
}
func (m *CallPlaceholderFieldEnum) XXX_Size() int {
	return xxx_messageInfo_CallPlaceholderFieldEnum.Size(m)
}
func (m *CallPlaceholderFieldEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_CallPlaceholderFieldEnum.DiscardUnknown(m)
}

var xxx_messageInfo_CallPlaceholderFieldEnum proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("google.ads.googleads.v0.enums.CallPlaceholderFieldEnum_CallPlaceholderField", CallPlaceholderFieldEnum_CallPlaceholderField_name, CallPlaceholderFieldEnum_CallPlaceholderField_value)
	proto.RegisterType((*CallPlaceholderFieldEnum)(nil), "google.ads.googleads.v0.enums.CallPlaceholderFieldEnum")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v0/enums/call_placeholder_field.proto", fileDescriptor_1d28e3249364eca5)
}

var fileDescriptor_1d28e3249364eca5 = []byte{
	// 356 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x90, 0xcf, 0x4a, 0xf3, 0x40,
	0x14, 0xc5, 0xbf, 0xa4, 0x9f, 0x15, 0xa6, 0x82, 0x43, 0x10, 0x51, 0xa1, 0x42, 0xfb, 0x00, 0x93,
	0x80, 0xbb, 0x71, 0x95, 0x3f, 0xd3, 0x1a, 0x8a, 0x93, 0x90, 0x26, 0x91, 0x4a, 0x20, 0xc4, 0x26,
	0xc6, 0xc2, 0x34, 0x53, 0x3a, 0xb6, 0x6f, 0xe3, 0xc6, 0xa5, 0x8f, 0xe0, 0x23, 0xf8, 0x28, 0x2e,
	0x7c, 0x06, 0x99, 0xc4, 0x56, 0x17, 0xd5, 0xcd, 0x70, 0x38, 0xe7, 0xfe, 0x86, 0x7b, 0x0f, 0xc0,
	0x25, 0xe7, 0x25, 0x2b, 0xf4, 0x2c, 0x17, 0x7a, 0x23, 0xa5, 0x5a, 0x1b, 0x7a, 0x51, 0xad, 0xe6,
	0x42, 0x9f, 0x66, 0x8c, 0xa5, 0x0b, 0x96, 0x4d, 0x8b, 0x07, 0xce, 0xf2, 0x62, 0x99, 0xde, 0xcf,
	0x0a, 0x96, 0xa3, 0xc5, 0x92, 0x3f, 0x72, 0xad, 0xdb, 0x00, 0x28, 0xcb, 0x05, 0xda, 0xb2, 0x68,
	0x6d, 0xa0, 0x9a, 0xed, 0xbf, 0x2a, 0xe0, 0xc4, 0xce, 0x18, 0xf3, 0xbf, 0xf1, 0x81, 0xa4, 0x49,
	0xb5, 0x9a, 0xf7, 0x9f, 0x14, 0x70, 0xb4, 0x2b, 0xd4, 0x0e, 0x41, 0x27, 0xa2, 0x63, 0x9f, 0xd8,
	0xee, 0xc0, 0x25, 0x0e, 0xfc, 0xa7, 0x75, 0xc0, 0x7e, 0x44, 0x47, 0xd4, 0xbb, 0xa1, 0x50, 0xd1,
	0x20, 0x38, 0xf0, 0xaf, 0x3c, 0x4a, 0x52, 0x1a, 0x5d, 0x5b, 0x24, 0x80, 0xaa, 0x74, 0x6c, 0x2f,
	0xa2, 0x61, 0x30, 0x49, 0x6d, 0xcf, 0x21, 0xb0, 0x25, 0x81, 0x30, 0x30, 0xed, 0x11, 0x71, 0xe0,
	0x7f, 0xed, 0x18, 0x68, 0xb6, 0x47, 0x63, 0x12, 0x8c, 0x5d, 0x8f, 0xa6, 0xe1, 0xc4, 0x27, 0xa9,
	0xeb, 0xc0, 0x3d, 0xed, 0x1c, 0x9c, 0xfd, 0xf0, 0x03, 0xe2, 0x7b, 0x41, 0xe8, 0xd2, 0x61, 0x3a,
	0x0e, 0xcd, 0x90, 0xc0, 0xb6, 0xf5, 0xa1, 0x80, 0xde, 0x94, 0xcf, 0xd1, 0x9f, 0x27, 0x5a, 0xa7,
	0xbb, 0x4e, 0xf0, 0x65, 0x39, 0xbe, 0x72, 0x6b, 0x7d, 0xb1, 0x25, 0x67, 0x59, 0x55, 0x22, 0xbe,
	0x2c, 0xf5, 0xb2, 0xa8, 0xea, 0xea, 0x36, 0x55, 0x2f, 0x66, 0xe2, 0x97, 0xe6, 0x2f, 0xeb, 0xf7,
	0x59, 0x6d, 0x0d, 0x4d, 0xf3, 0x45, 0xed, 0x0e, 0x9b, 0xaf, 0xcc, 0x5c, 0xa0, 0x46, 0x4a, 0x15,
	0x1b, 0x48, 0x76, 0x29, 0xde, 0x36, 0x79, 0x62, 0xe6, 0x22, 0xd9, 0xe6, 0x49, 0x6c, 0x24, 0x75,
	0xfe, 0xae, 0xf6, 0x1a, 0x13, 0x63, 0x33, 0x17, 0x18, 0x6f, 0x27, 0x30, 0x8e, 0x0d, 0x8c, 0xeb,
	0x99, 0xbb, 0x76, 0xbd, 0xd8, 0xc5, 0x67, 0x00, 0x00, 0x00, 0xff, 0xff, 0x11, 0xf5, 0xfd, 0x3d,
	0x11, 0x02, 0x00, 0x00,
}
