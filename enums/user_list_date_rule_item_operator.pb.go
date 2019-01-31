// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v0/enums/user_list_date_rule_item_operator.proto

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

// Enum describing possible user list date rule item operators.
type UserListDateRuleItemOperatorEnum_UserListDateRuleItemOperator int32

const (
	// Not specified.
	UserListDateRuleItemOperatorEnum_UNSPECIFIED UserListDateRuleItemOperatorEnum_UserListDateRuleItemOperator = 0
	// Used for return value only. Represents value unknown in this version.
	UserListDateRuleItemOperatorEnum_UNKNOWN UserListDateRuleItemOperatorEnum_UserListDateRuleItemOperator = 1
	// Equals.
	UserListDateRuleItemOperatorEnum_EQUALS UserListDateRuleItemOperatorEnum_UserListDateRuleItemOperator = 2
	// Not Equals.
	UserListDateRuleItemOperatorEnum_NOT_EQUALS UserListDateRuleItemOperatorEnum_UserListDateRuleItemOperator = 3
	// Before.
	UserListDateRuleItemOperatorEnum_BEFORE UserListDateRuleItemOperatorEnum_UserListDateRuleItemOperator = 4
	// After.
	UserListDateRuleItemOperatorEnum_AFTER UserListDateRuleItemOperatorEnum_UserListDateRuleItemOperator = 5
)

var UserListDateRuleItemOperatorEnum_UserListDateRuleItemOperator_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "UNKNOWN",
	2: "EQUALS",
	3: "NOT_EQUALS",
	4: "BEFORE",
	5: "AFTER",
}

var UserListDateRuleItemOperatorEnum_UserListDateRuleItemOperator_value = map[string]int32{
	"UNSPECIFIED": 0,
	"UNKNOWN":     1,
	"EQUALS":      2,
	"NOT_EQUALS":  3,
	"BEFORE":      4,
	"AFTER":       5,
}

func (x UserListDateRuleItemOperatorEnum_UserListDateRuleItemOperator) String() string {
	return proto.EnumName(UserListDateRuleItemOperatorEnum_UserListDateRuleItemOperator_name, int32(x))
}

func (UserListDateRuleItemOperatorEnum_UserListDateRuleItemOperator) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_38697bca8dea58a6, []int{0, 0}
}

// Supported rule operator for date type.
type UserListDateRuleItemOperatorEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserListDateRuleItemOperatorEnum) Reset()         { *m = UserListDateRuleItemOperatorEnum{} }
func (m *UserListDateRuleItemOperatorEnum) String() string { return proto.CompactTextString(m) }
func (*UserListDateRuleItemOperatorEnum) ProtoMessage()    {}
func (*UserListDateRuleItemOperatorEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_38697bca8dea58a6, []int{0}
}

func (m *UserListDateRuleItemOperatorEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserListDateRuleItemOperatorEnum.Unmarshal(m, b)
}
func (m *UserListDateRuleItemOperatorEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserListDateRuleItemOperatorEnum.Marshal(b, m, deterministic)
}
func (m *UserListDateRuleItemOperatorEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserListDateRuleItemOperatorEnum.Merge(m, src)
}
func (m *UserListDateRuleItemOperatorEnum) XXX_Size() int {
	return xxx_messageInfo_UserListDateRuleItemOperatorEnum.Size(m)
}
func (m *UserListDateRuleItemOperatorEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_UserListDateRuleItemOperatorEnum.DiscardUnknown(m)
}

var xxx_messageInfo_UserListDateRuleItemOperatorEnum proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("google.ads.googleads.v0.enums.UserListDateRuleItemOperatorEnum_UserListDateRuleItemOperator", UserListDateRuleItemOperatorEnum_UserListDateRuleItemOperator_name, UserListDateRuleItemOperatorEnum_UserListDateRuleItemOperator_value)
	proto.RegisterType((*UserListDateRuleItemOperatorEnum)(nil), "google.ads.googleads.v0.enums.UserListDateRuleItemOperatorEnum")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v0/enums/user_list_date_rule_item_operator.proto", fileDescriptor_38697bca8dea58a6)
}

var fileDescriptor_38697bca8dea58a6 = []byte{
	// 329 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x90, 0xd1, 0x6a, 0xb3, 0x30,
	0x1c, 0xc5, 0x3f, 0xed, 0xd7, 0x8e, 0xa5, 0xb0, 0x89, 0xd7, 0x2b, 0xac, 0x7d, 0x80, 0x28, 0xec,
	0x2e, 0xbb, 0x8a, 0x6b, 0x5a, 0xca, 0x8a, 0x76, 0x6d, 0xed, 0x60, 0x08, 0xe2, 0x66, 0x10, 0x41,
	0x4d, 0xc9, 0x3f, 0xf6, 0x45, 0xf6, 0x06, 0xbb, 0xdc, 0xa3, 0xec, 0x51, 0xf6, 0x00, 0xbb, 0x1e,
	0xc6, 0xb6, 0x77, 0xf3, 0x46, 0x8e, 0x39, 0x27, 0x3f, 0x4e, 0x0e, 0x62, 0x99, 0x10, 0x59, 0xc1,
	0x9d, 0x24, 0x05, 0xa7, 0x95, 0x8d, 0x3a, 0xb8, 0x0e, 0xaf, 0xea, 0x12, 0x9c, 0x1a, 0xb8, 0x8c,
	0x8b, 0x1c, 0x54, 0x9c, 0x26, 0x8a, 0xc7, 0xb2, 0x2e, 0x78, 0x9c, 0x2b, 0x5e, 0xc6, 0x62, 0xcf,
	0x65, 0xa2, 0x84, 0xc4, 0x7b, 0x29, 0x94, 0xb0, 0x47, 0xed, 0x5d, 0x9c, 0xa4, 0x80, 0xcf, 0x18,
	0x7c, 0x70, 0xb1, 0xc6, 0x4c, 0xde, 0x0d, 0x74, 0x1b, 0x02, 0x97, 0xcb, 0x1c, 0xd4, 0x34, 0x51,
	0x7c, 0x5d, 0x17, 0x7c, 0xa1, 0x78, 0x19, 0x1c, 0x29, 0xac, 0xaa, 0xcb, 0x89, 0x40, 0x37, 0x5d,
	0x19, 0xfb, 0x1a, 0x0d, 0x43, 0x7f, 0xb3, 0x62, 0x0f, 0x8b, 0xd9, 0x82, 0x4d, 0xad, 0x7f, 0xf6,
	0x10, 0x5d, 0x84, 0xfe, 0xa3, 0x1f, 0x3c, 0xfb, 0x96, 0x61, 0x23, 0x34, 0x60, 0x4f, 0x21, 0x5d,
	0x6e, 0x2c, 0xd3, 0xbe, 0x42, 0xc8, 0x0f, 0xb6, 0xf1, 0xf1, 0xbf, 0xd7, 0x78, 0x1e, 0x9b, 0x05,
	0x6b, 0x66, 0xfd, 0xb7, 0x2f, 0x51, 0x9f, 0xce, 0xb6, 0x6c, 0x6d, 0xf5, 0xbd, 0x1f, 0x03, 0x8d,
	0xdf, 0x44, 0x89, 0x3b, 0xbb, 0x7b, 0xe3, 0xae, 0x52, 0xab, 0xe6, 0xf5, 0x2b, 0xe3, 0xc5, 0x3b,
	0x32, 0x32, 0x51, 0x24, 0x55, 0x86, 0x85, 0xcc, 0x9c, 0x8c, 0x57, 0x7a, 0x9b, 0xd3, 0xac, 0xfb,
	0x1c, 0xfe, 0x58, 0xf9, 0x5e, 0x7f, 0x3f, 0xcc, 0xde, 0x9c, 0xd2, 0x4f, 0x73, 0x34, 0x6f, 0x51,
	0x34, 0x05, 0xdc, 0xca, 0x46, 0xed, 0x5c, 0xdc, 0x8c, 0x04, 0x5f, 0x27, 0x3f, 0xa2, 0x29, 0x44,
	0x67, 0x3f, 0xda, 0xb9, 0x91, 0xf6, 0xbf, 0xcd, 0x71, 0x7b, 0x48, 0x08, 0x4d, 0x81, 0x90, 0x73,
	0x82, 0x90, 0x9d, 0x4b, 0x88, 0xce, 0xbc, 0x0e, 0x74, 0xb1, 0xbb, 0xdf, 0x00, 0x00, 0x00, 0xff,
	0xff, 0x56, 0xf3, 0x37, 0xcb, 0xfd, 0x01, 0x00, 0x00,
}
