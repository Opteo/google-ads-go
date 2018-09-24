// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v0/errors/authentication_error.proto

package google_ads_googleads_v0_errors

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
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Enum describing possible authentication errors.
type AuthenticationErrorEnum_AuthenticationError int32

const (
	// Enum unspecified.
	AuthenticationErrorEnum_UNSPECIFIED AuthenticationErrorEnum_AuthenticationError = 0
	// The received error code is not known in this version.
	AuthenticationErrorEnum_UNKNOWN AuthenticationErrorEnum_AuthenticationError = 1
	// Authentication of the request failed.
	AuthenticationErrorEnum_AUTHENTICATION_ERROR AuthenticationErrorEnum_AuthenticationError = 2
	// Client customer Id is not a number.
	AuthenticationErrorEnum_CLIENT_CUSTOMER_ID_INVALID AuthenticationErrorEnum_AuthenticationError = 5
	// No customer found for the customer id provided in the header.
	AuthenticationErrorEnum_CUSTOMER_NOT_FOUND AuthenticationErrorEnum_AuthenticationError = 8
	// Client's Google Account is deleted.
	AuthenticationErrorEnum_GOOGLE_ACCOUNT_DELETED AuthenticationErrorEnum_AuthenticationError = 9
	// Google account login token in the cookie is invalid.
	AuthenticationErrorEnum_GOOGLE_ACCOUNT_COOKIE_INVALID AuthenticationErrorEnum_AuthenticationError = 10
	// A problem occurred during Google account authentication.
	AuthenticationErrorEnum_FAILED_TO_AUTHENTICATE_GOOGLE_ACCOUNT AuthenticationErrorEnum_AuthenticationError = 11
	// The user in the google account login token does not match the UserId in
	// the cookie.
	AuthenticationErrorEnum_GOOGLE_ACCOUNT_USER_AND_ADS_USER_MISMATCH AuthenticationErrorEnum_AuthenticationError = 12
	// Login cookie is required for authentication.
	AuthenticationErrorEnum_LOGIN_COOKIE_REQUIRED AuthenticationErrorEnum_AuthenticationError = 13
	// User in the cookie is not a valid Ads user.
	AuthenticationErrorEnum_NOT_ADS_USER AuthenticationErrorEnum_AuthenticationError = 14
	// Oauth token in the header is not valid.
	AuthenticationErrorEnum_OAUTH_TOKEN_INVALID AuthenticationErrorEnum_AuthenticationError = 15
	// Oauth token in the header has expired.
	AuthenticationErrorEnum_OAUTH_TOKEN_EXPIRED AuthenticationErrorEnum_AuthenticationError = 16
	// Oauth token in the header has been disabled.
	AuthenticationErrorEnum_OAUTH_TOKEN_DISABLED AuthenticationErrorEnum_AuthenticationError = 17
	// Oauth token in the header has been revoked.
	AuthenticationErrorEnum_OAUTH_TOKEN_REVOKED AuthenticationErrorEnum_AuthenticationError = 18
	// Oauth token HTTP header is malformed.
	AuthenticationErrorEnum_OAUTH_TOKEN_HEADER_INVALID AuthenticationErrorEnum_AuthenticationError = 19
	// Login cookie is not valid.
	AuthenticationErrorEnum_LOGIN_COOKIE_INVALID AuthenticationErrorEnum_AuthenticationError = 20
	// Failed to decrypt the login cookie.
	AuthenticationErrorEnum_FAILED_TO_RETRIEVE_LOGIN_COOKIE AuthenticationErrorEnum_AuthenticationError = 21
	// User Id in the header is not a valid id.
	AuthenticationErrorEnum_USER_ID_INVALID AuthenticationErrorEnum_AuthenticationError = 22
)

var AuthenticationErrorEnum_AuthenticationError_name = map[int32]string{
	0:  "UNSPECIFIED",
	1:  "UNKNOWN",
	2:  "AUTHENTICATION_ERROR",
	5:  "CLIENT_CUSTOMER_ID_INVALID",
	8:  "CUSTOMER_NOT_FOUND",
	9:  "GOOGLE_ACCOUNT_DELETED",
	10: "GOOGLE_ACCOUNT_COOKIE_INVALID",
	11: "FAILED_TO_AUTHENTICATE_GOOGLE_ACCOUNT",
	12: "GOOGLE_ACCOUNT_USER_AND_ADS_USER_MISMATCH",
	13: "LOGIN_COOKIE_REQUIRED",
	14: "NOT_ADS_USER",
	15: "OAUTH_TOKEN_INVALID",
	16: "OAUTH_TOKEN_EXPIRED",
	17: "OAUTH_TOKEN_DISABLED",
	18: "OAUTH_TOKEN_REVOKED",
	19: "OAUTH_TOKEN_HEADER_INVALID",
	20: "LOGIN_COOKIE_INVALID",
	21: "FAILED_TO_RETRIEVE_LOGIN_COOKIE",
	22: "USER_ID_INVALID",
}

var AuthenticationErrorEnum_AuthenticationError_value = map[string]int32{
	"UNSPECIFIED":                               0,
	"UNKNOWN":                                   1,
	"AUTHENTICATION_ERROR":                      2,
	"CLIENT_CUSTOMER_ID_INVALID":                5,
	"CUSTOMER_NOT_FOUND":                        8,
	"GOOGLE_ACCOUNT_DELETED":                    9,
	"GOOGLE_ACCOUNT_COOKIE_INVALID":             10,
	"FAILED_TO_AUTHENTICATE_GOOGLE_ACCOUNT":     11,
	"GOOGLE_ACCOUNT_USER_AND_ADS_USER_MISMATCH": 12,
	"LOGIN_COOKIE_REQUIRED":                     13,
	"NOT_ADS_USER":                              14,
	"OAUTH_TOKEN_INVALID":                       15,
	"OAUTH_TOKEN_EXPIRED":                       16,
	"OAUTH_TOKEN_DISABLED":                      17,
	"OAUTH_TOKEN_REVOKED":                       18,
	"OAUTH_TOKEN_HEADER_INVALID":                19,
	"LOGIN_COOKIE_INVALID":                      20,
	"FAILED_TO_RETRIEVE_LOGIN_COOKIE":           21,
	"USER_ID_INVALID":                           22,
}

func (x AuthenticationErrorEnum_AuthenticationError) String() string {
	return proto.EnumName(AuthenticationErrorEnum_AuthenticationError_name, int32(x))
}

func (AuthenticationErrorEnum_AuthenticationError) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_bf995c8bcd802bf7, []int{0, 0}
}

// Container for enum describing possible authentication errors.
type AuthenticationErrorEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AuthenticationErrorEnum) Reset()         { *m = AuthenticationErrorEnum{} }
func (m *AuthenticationErrorEnum) String() string { return proto.CompactTextString(m) }
func (*AuthenticationErrorEnum) ProtoMessage()    {}
func (*AuthenticationErrorEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_bf995c8bcd802bf7, []int{0}
}
func (m *AuthenticationErrorEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AuthenticationErrorEnum.Unmarshal(m, b)
}
func (m *AuthenticationErrorEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AuthenticationErrorEnum.Marshal(b, m, deterministic)
}
func (m *AuthenticationErrorEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuthenticationErrorEnum.Merge(m, src)
}
func (m *AuthenticationErrorEnum) XXX_Size() int {
	return xxx_messageInfo_AuthenticationErrorEnum.Size(m)
}
func (m *AuthenticationErrorEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_AuthenticationErrorEnum.DiscardUnknown(m)
}

var xxx_messageInfo_AuthenticationErrorEnum proto.InternalMessageInfo

func init() {
	proto.RegisterType((*AuthenticationErrorEnum)(nil), "google.ads.googleads.v0.errors.AuthenticationErrorEnum")
	proto.RegisterEnum("google.ads.googleads.v0.errors.AuthenticationErrorEnum_AuthenticationError", AuthenticationErrorEnum_AuthenticationError_name, AuthenticationErrorEnum_AuthenticationError_value)
}

func init() {
	proto.RegisterFile("google/ads/googleads/v0/errors/authentication_error.proto", fileDescriptor_bf995c8bcd802bf7)
}

var fileDescriptor_bf995c8bcd802bf7 = []byte{
	// 478 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x53, 0x4f, 0x6f, 0xd3, 0x30,
	0x14, 0xa7, 0x65, 0xfc, 0x7b, 0x1d, 0xd4, 0xb8, 0x5b, 0x37, 0x26, 0x51, 0x44, 0x11, 0x87, 0x1d,
	0x48, 0x2b, 0x71, 0xe2, 0xe8, 0xc6, 0xaf, 0xad, 0xd5, 0xd4, 0x2e, 0x8e, 0x13, 0x38, 0x54, 0xb2,
	0xca, 0x5a, 0x8d, 0x49, 0xac, 0x41, 0x4d, 0xb7, 0x0f, 0x84, 0xb8, 0xc0, 0xa7, 0xe0, 0xcc, 0xc7,
	0xe0, 0x93, 0x20, 0x27, 0x34, 0xa4, 0xd5, 0x04, 0xb7, 0x97, 0xfc, 0xfe, 0xbc, 0xe7, 0x9f, 0x9f,
	0xe1, 0xcd, 0x79, 0x92, 0x9c, 0x7f, 0x5a, 0x74, 0x66, 0xf3, 0xb4, 0x93, 0x97, 0xae, 0xba, 0xee,
	0x76, 0x16, 0xab, 0x55, 0xb2, 0x4a, 0x3b, 0xb3, 0xab, 0xf5, 0xc7, 0xc5, 0x72, 0x7d, 0x71, 0x36,
	0x5b, 0x5f, 0x24, 0x4b, 0x9b, 0xfd, 0xf5, 0x3e, 0xaf, 0x92, 0x75, 0x42, 0x5b, 0x39, 0xdf, 0x9b,
	0xcd, 0x53, 0xaf, 0x90, 0x7a, 0xd7, 0x5d, 0x2f, 0x97, 0xb6, 0x7f, 0xec, 0xc1, 0x11, 0xdb, 0x92,
	0xa3, 0x03, 0x70, 0x79, 0x75, 0xd9, 0xfe, 0xba, 0x07, 0x8d, 0x1b, 0x30, 0x5a, 0x87, 0x5a, 0x24,
	0xc3, 0x09, 0xfa, 0xa2, 0x2f, 0x90, 0x93, 0x5b, 0xb4, 0x06, 0xf7, 0x22, 0x39, 0x92, 0xea, 0x9d,
	0x24, 0x15, 0x7a, 0x0c, 0x07, 0x2c, 0x32, 0x43, 0x94, 0x46, 0xf8, 0xcc, 0x08, 0x25, 0x2d, 0x6a,
	0xad, 0x34, 0xa9, 0xd2, 0x16, 0x9c, 0xf8, 0x81, 0x40, 0x69, 0xac, 0x1f, 0x85, 0x46, 0x8d, 0x51,
	0x5b, 0xc1, 0xad, 0x90, 0x31, 0x0b, 0x04, 0x27, 0x77, 0x68, 0x13, 0x68, 0x01, 0x48, 0x65, 0x6c,
	0x5f, 0x45, 0x92, 0x93, 0xfb, 0xf4, 0x04, 0x9a, 0x03, 0xa5, 0x06, 0x01, 0x5a, 0xe6, 0xfb, 0x2a,
	0x92, 0xc6, 0x72, 0x0c, 0xd0, 0x20, 0x27, 0x0f, 0xe8, 0x73, 0x78, 0xba, 0x83, 0xf9, 0x4a, 0x8d,
	0x04, 0x16, 0xb6, 0x40, 0x4f, 0xe1, 0x65, 0x9f, 0x89, 0x00, 0xb9, 0x35, 0xca, 0x96, 0x46, 0x43,
	0xbb, 0xad, 0x24, 0x35, 0xfa, 0x0a, 0x4e, 0x77, 0xdc, 0xa2, 0x10, 0xb5, 0x65, 0x92, 0x5b, 0xc6,
	0xc3, 0xfc, 0x63, 0x2c, 0xc2, 0x31, 0x33, 0xfe, 0x90, 0xec, 0xd3, 0x27, 0x70, 0x18, 0xa8, 0x81,
	0x90, 0x9b, 0x9e, 0x1a, 0xdf, 0x46, 0x42, 0x23, 0x27, 0x0f, 0x29, 0x81, 0x7d, 0x77, 0x84, 0x8d,
	0x8a, 0x3c, 0xa2, 0x47, 0xd0, 0x50, 0xae, 0xbb, 0x35, 0x6a, 0x84, 0xb2, 0x98, 0xaf, 0xbe, 0x0b,
	0xe0, 0xfb, 0x49, 0xe6, 0x41, 0x5c, 0x92, 0x65, 0x80, 0x8b, 0x90, 0xf5, 0x02, 0xe4, 0xe4, 0xf1,
	0xae, 0x44, 0x63, 0xac, 0x46, 0xc8, 0x09, 0x75, 0x11, 0x97, 0x81, 0x21, 0x32, 0xee, 0x52, 0xfe,
	0xd3, 0xab, 0xe1, 0x2c, 0xb7, 0x26, 0xde, 0x20, 0x07, 0xf4, 0x05, 0x3c, 0xfb, 0x9b, 0x92, 0x46,
	0xa3, 0x05, 0xc6, 0x68, 0xcb, 0x64, 0x72, 0x48, 0x1b, 0x50, 0xcf, 0x32, 0x28, 0x5d, 0x5b, 0xb3,
	0xf7, 0xad, 0x02, 0xed, 0xb3, 0xe4, 0xd2, 0xfb, 0xf7, 0xa6, 0xf5, 0x8e, 0x6f, 0x58, 0xa5, 0x89,
	0xdb, 0xd1, 0x49, 0xe5, 0x4b, 0xf5, 0xf6, 0x80, 0xb1, 0xef, 0xd5, 0xd6, 0x20, 0xb7, 0x60, 0xf3,
	0xd4, 0xcb, 0x4b, 0x57, 0xc5, 0x5d, 0x2f, 0x23, 0xa7, 0x3f, 0x37, 0x84, 0x29, 0x9b, 0xa7, 0xd3,
	0x82, 0x30, 0x8d, 0xbb, 0xd3, 0x9c, 0xf0, 0xeb, 0x7f, 0x84, 0x0f, 0x77, 0xb3, 0x57, 0xf1, 0xfa,
	0x77, 0x00, 0x00, 0x00, 0xff, 0xff, 0x98, 0x78, 0xab, 0x73, 0x52, 0x03, 0x00, 0x00,
}