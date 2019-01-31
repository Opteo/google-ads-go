// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v0/services/ad_parameter_service.proto

package services

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/golang/protobuf/ptypes/wrappers"
	resources "github.com/kritzware/google-ads-go/resources"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	status "google.golang.org/genproto/googleapis/rpc/status"
	field_mask "google.golang.org/genproto/protobuf/field_mask"
	grpc "google.golang.org/grpc"
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

// Request message for [AdParameterService.GetAdParameter][google.ads.googleads.v0.services.AdParameterService.GetAdParameter]
type GetAdParameterRequest struct {
	// The resource name of the ad parameter to fetch.
	ResourceName         string   `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetAdParameterRequest) Reset()         { *m = GetAdParameterRequest{} }
func (m *GetAdParameterRequest) String() string { return proto.CompactTextString(m) }
func (*GetAdParameterRequest) ProtoMessage()    {}
func (*GetAdParameterRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d5f66eb03fa8bcff, []int{0}
}

func (m *GetAdParameterRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetAdParameterRequest.Unmarshal(m, b)
}
func (m *GetAdParameterRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetAdParameterRequest.Marshal(b, m, deterministic)
}
func (m *GetAdParameterRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetAdParameterRequest.Merge(m, src)
}
func (m *GetAdParameterRequest) XXX_Size() int {
	return xxx_messageInfo_GetAdParameterRequest.Size(m)
}
func (m *GetAdParameterRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetAdParameterRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetAdParameterRequest proto.InternalMessageInfo

func (m *GetAdParameterRequest) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

// Request message for [AdParameterService.MutateAdParameters][google.ads.googleads.v0.services.AdParameterService.MutateAdParameters]
type MutateAdParametersRequest struct {
	// The ID of the customer whose ad parameters are being modified.
	CustomerId string `protobuf:"bytes,1,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	// The list of operations to perform on individual ad parameters.
	Operations []*AdParameterOperation `protobuf:"bytes,2,rep,name=operations,proto3" json:"operations,omitempty"`
	// If true, successful operations will be carried out and invalid
	// operations will return errors. If false, all operations will be carried
	// out in one transaction if and only if they are all valid.
	// Default is false.
	PartialFailure bool `protobuf:"varint,3,opt,name=partial_failure,json=partialFailure,proto3" json:"partial_failure,omitempty"`
	// If true, the request is validated but not executed. Only errors are
	// returned, not results.
	ValidateOnly         bool     `protobuf:"varint,4,opt,name=validate_only,json=validateOnly,proto3" json:"validate_only,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MutateAdParametersRequest) Reset()         { *m = MutateAdParametersRequest{} }
func (m *MutateAdParametersRequest) String() string { return proto.CompactTextString(m) }
func (*MutateAdParametersRequest) ProtoMessage()    {}
func (*MutateAdParametersRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d5f66eb03fa8bcff, []int{1}
}

func (m *MutateAdParametersRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MutateAdParametersRequest.Unmarshal(m, b)
}
func (m *MutateAdParametersRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MutateAdParametersRequest.Marshal(b, m, deterministic)
}
func (m *MutateAdParametersRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MutateAdParametersRequest.Merge(m, src)
}
func (m *MutateAdParametersRequest) XXX_Size() int {
	return xxx_messageInfo_MutateAdParametersRequest.Size(m)
}
func (m *MutateAdParametersRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_MutateAdParametersRequest.DiscardUnknown(m)
}

var xxx_messageInfo_MutateAdParametersRequest proto.InternalMessageInfo

func (m *MutateAdParametersRequest) GetCustomerId() string {
	if m != nil {
		return m.CustomerId
	}
	return ""
}

func (m *MutateAdParametersRequest) GetOperations() []*AdParameterOperation {
	if m != nil {
		return m.Operations
	}
	return nil
}

func (m *MutateAdParametersRequest) GetPartialFailure() bool {
	if m != nil {
		return m.PartialFailure
	}
	return false
}

func (m *MutateAdParametersRequest) GetValidateOnly() bool {
	if m != nil {
		return m.ValidateOnly
	}
	return false
}

// A single operation (create, update, remove) on ad parameter.
type AdParameterOperation struct {
	// FieldMask that determines which resource fields are modified in an update.
	UpdateMask *field_mask.FieldMask `protobuf:"bytes,4,opt,name=update_mask,json=updateMask,proto3" json:"update_mask,omitempty"`
	// The mutate operation.
	//
	// Types that are valid to be assigned to Operation:
	//	*AdParameterOperation_Create
	//	*AdParameterOperation_Update
	//	*AdParameterOperation_Remove
	Operation            isAdParameterOperation_Operation `protobuf_oneof:"operation"`
	XXX_NoUnkeyedLiteral struct{}                         `json:"-"`
	XXX_unrecognized     []byte                           `json:"-"`
	XXX_sizecache        int32                            `json:"-"`
}

func (m *AdParameterOperation) Reset()         { *m = AdParameterOperation{} }
func (m *AdParameterOperation) String() string { return proto.CompactTextString(m) }
func (*AdParameterOperation) ProtoMessage()    {}
func (*AdParameterOperation) Descriptor() ([]byte, []int) {
	return fileDescriptor_d5f66eb03fa8bcff, []int{2}
}

func (m *AdParameterOperation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AdParameterOperation.Unmarshal(m, b)
}
func (m *AdParameterOperation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AdParameterOperation.Marshal(b, m, deterministic)
}
func (m *AdParameterOperation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AdParameterOperation.Merge(m, src)
}
func (m *AdParameterOperation) XXX_Size() int {
	return xxx_messageInfo_AdParameterOperation.Size(m)
}
func (m *AdParameterOperation) XXX_DiscardUnknown() {
	xxx_messageInfo_AdParameterOperation.DiscardUnknown(m)
}

var xxx_messageInfo_AdParameterOperation proto.InternalMessageInfo

func (m *AdParameterOperation) GetUpdateMask() *field_mask.FieldMask {
	if m != nil {
		return m.UpdateMask
	}
	return nil
}

type isAdParameterOperation_Operation interface {
	isAdParameterOperation_Operation()
}

type AdParameterOperation_Create struct {
	Create *resources.AdParameter `protobuf:"bytes,1,opt,name=create,proto3,oneof"`
}

type AdParameterOperation_Update struct {
	Update *resources.AdParameter `protobuf:"bytes,2,opt,name=update,proto3,oneof"`
}

type AdParameterOperation_Remove struct {
	Remove string `protobuf:"bytes,3,opt,name=remove,proto3,oneof"`
}

func (*AdParameterOperation_Create) isAdParameterOperation_Operation() {}

func (*AdParameterOperation_Update) isAdParameterOperation_Operation() {}

func (*AdParameterOperation_Remove) isAdParameterOperation_Operation() {}

func (m *AdParameterOperation) GetOperation() isAdParameterOperation_Operation {
	if m != nil {
		return m.Operation
	}
	return nil
}

func (m *AdParameterOperation) GetCreate() *resources.AdParameter {
	if x, ok := m.GetOperation().(*AdParameterOperation_Create); ok {
		return x.Create
	}
	return nil
}

func (m *AdParameterOperation) GetUpdate() *resources.AdParameter {
	if x, ok := m.GetOperation().(*AdParameterOperation_Update); ok {
		return x.Update
	}
	return nil
}

func (m *AdParameterOperation) GetRemove() string {
	if x, ok := m.GetOperation().(*AdParameterOperation_Remove); ok {
		return x.Remove
	}
	return ""
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*AdParameterOperation) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*AdParameterOperation_Create)(nil),
		(*AdParameterOperation_Update)(nil),
		(*AdParameterOperation_Remove)(nil),
	}
}

// Response message for an ad parameter mutate.
type MutateAdParametersResponse struct {
	// Errors that pertain to operation failures in the partial failure mode.
	// Returned only when partial_failure = true and all errors occur inside the
	// operations. If any errors occur outside the operations (e.g. auth errors),
	// we return an RPC level error.
	PartialFailureError *status.Status `protobuf:"bytes,3,opt,name=partial_failure_error,json=partialFailureError,proto3" json:"partial_failure_error,omitempty"`
	// All results for the mutate.
	Results              []*MutateAdParameterResult `protobuf:"bytes,2,rep,name=results,proto3" json:"results,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                   `json:"-"`
	XXX_unrecognized     []byte                     `json:"-"`
	XXX_sizecache        int32                      `json:"-"`
}

func (m *MutateAdParametersResponse) Reset()         { *m = MutateAdParametersResponse{} }
func (m *MutateAdParametersResponse) String() string { return proto.CompactTextString(m) }
func (*MutateAdParametersResponse) ProtoMessage()    {}
func (*MutateAdParametersResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_d5f66eb03fa8bcff, []int{3}
}

func (m *MutateAdParametersResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MutateAdParametersResponse.Unmarshal(m, b)
}
func (m *MutateAdParametersResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MutateAdParametersResponse.Marshal(b, m, deterministic)
}
func (m *MutateAdParametersResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MutateAdParametersResponse.Merge(m, src)
}
func (m *MutateAdParametersResponse) XXX_Size() int {
	return xxx_messageInfo_MutateAdParametersResponse.Size(m)
}
func (m *MutateAdParametersResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MutateAdParametersResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MutateAdParametersResponse proto.InternalMessageInfo

func (m *MutateAdParametersResponse) GetPartialFailureError() *status.Status {
	if m != nil {
		return m.PartialFailureError
	}
	return nil
}

func (m *MutateAdParametersResponse) GetResults() []*MutateAdParameterResult {
	if m != nil {
		return m.Results
	}
	return nil
}

// The result for the ad parameter mutate.
type MutateAdParameterResult struct {
	// The resource name returned for successful operations.
	ResourceName         string   `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MutateAdParameterResult) Reset()         { *m = MutateAdParameterResult{} }
func (m *MutateAdParameterResult) String() string { return proto.CompactTextString(m) }
func (*MutateAdParameterResult) ProtoMessage()    {}
func (*MutateAdParameterResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_d5f66eb03fa8bcff, []int{4}
}

func (m *MutateAdParameterResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MutateAdParameterResult.Unmarshal(m, b)
}
func (m *MutateAdParameterResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MutateAdParameterResult.Marshal(b, m, deterministic)
}
func (m *MutateAdParameterResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MutateAdParameterResult.Merge(m, src)
}
func (m *MutateAdParameterResult) XXX_Size() int {
	return xxx_messageInfo_MutateAdParameterResult.Size(m)
}
func (m *MutateAdParameterResult) XXX_DiscardUnknown() {
	xxx_messageInfo_MutateAdParameterResult.DiscardUnknown(m)
}

var xxx_messageInfo_MutateAdParameterResult proto.InternalMessageInfo

func (m *MutateAdParameterResult) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func init() {
	proto.RegisterType((*GetAdParameterRequest)(nil), "google.ads.googleads.v0.services.GetAdParameterRequest")
	proto.RegisterType((*MutateAdParametersRequest)(nil), "google.ads.googleads.v0.services.MutateAdParametersRequest")
	proto.RegisterType((*AdParameterOperation)(nil), "google.ads.googleads.v0.services.AdParameterOperation")
	proto.RegisterType((*MutateAdParametersResponse)(nil), "google.ads.googleads.v0.services.MutateAdParametersResponse")
	proto.RegisterType((*MutateAdParameterResult)(nil), "google.ads.googleads.v0.services.MutateAdParameterResult")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v0/services/ad_parameter_service.proto", fileDescriptor_d5f66eb03fa8bcff)
}

var fileDescriptor_d5f66eb03fa8bcff = []byte{
	// 710 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x55, 0x4f, 0x6b, 0xd4, 0x4e,
	0x18, 0xfe, 0x25, 0xfb, 0xa3, 0xda, 0x49, 0xad, 0x30, 0x5a, 0xba, 0x2e, 0xa2, 0x4b, 0x2c, 0x58,
	0xf6, 0x30, 0x59, 0x57, 0xa9, 0x98, 0x6d, 0x85, 0x2d, 0xd8, 0xd6, 0x43, 0x6d, 0x49, 0x61, 0x0f,
	0xb2, 0x10, 0xa6, 0x9b, 0xe9, 0x12, 0x9a, 0x64, 0xe2, 0xcc, 0x64, 0xa5, 0x94, 0x5e, 0xfc, 0x00,
	0x5e, 0xbc, 0x78, 0xf6, 0xe8, 0xcd, 0xa3, 0x5f, 0x41, 0xf0, 0xe4, 0x37, 0x10, 0x4f, 0x7e, 0x08,
	0x91, 0xc9, 0x64, 0xd6, 0x6c, 0xdb, 0x65, 0x75, 0x6f, 0x6f, 0xde, 0x79, 0x9e, 0x67, 0xde, 0xbf,
	0x13, 0xd0, 0x1e, 0x50, 0x3a, 0x88, 0x88, 0x83, 0x03, 0xee, 0x28, 0x53, 0x5a, 0xc3, 0xa6, 0xc3,
	0x09, 0x1b, 0x86, 0x7d, 0xc2, 0x1d, 0x1c, 0xf8, 0x29, 0x66, 0x38, 0x26, 0x82, 0x30, 0xbf, 0xf0,
	0xa2, 0x94, 0x51, 0x41, 0x61, 0x5d, 0x31, 0x10, 0x0e, 0x38, 0x1a, 0x91, 0xd1, 0xb0, 0x89, 0x34,
	0xb9, 0xf6, 0x68, 0x92, 0x3c, 0x23, 0x9c, 0x66, 0xec, 0xbc, 0xbe, 0xd2, 0xad, 0xdd, 0xd6, 0xac,
	0x34, 0x74, 0x70, 0x92, 0x50, 0x81, 0x45, 0x48, 0x13, 0x5e, 0x9c, 0x16, 0xb7, 0x3a, 0xf9, 0xd7,
	0x61, 0x76, 0xe4, 0x1c, 0x85, 0x24, 0x0a, 0xfc, 0x18, 0xf3, 0xe3, 0x02, 0x71, 0xe7, 0x3c, 0xe2,
	0x35, 0xc3, 0x69, 0x4a, 0x98, 0x56, 0x58, 0x2e, 0xce, 0x59, 0xda, 0x77, 0xb8, 0xc0, 0x22, 0x2b,
	0x0e, 0xec, 0x75, 0xb0, 0xb4, 0x4d, 0x44, 0x27, 0xd8, 0xd7, 0x01, 0x79, 0xe4, 0x55, 0x46, 0xb8,
	0x80, 0xf7, 0xc0, 0x35, 0x1d, 0xb1, 0x9f, 0xe0, 0x98, 0x54, 0x8d, 0xba, 0xb1, 0x3a, 0xef, 0x2d,
	0x68, 0xe7, 0x0b, 0x1c, 0x13, 0xfb, 0xbb, 0x01, 0x6e, 0xed, 0x66, 0x02, 0x0b, 0x52, 0x52, 0xe0,
	0x5a, 0xe2, 0x2e, 0xb0, 0xfa, 0x19, 0x17, 0x34, 0x26, 0xcc, 0x0f, 0x83, 0x42, 0x00, 0x68, 0xd7,
	0xf3, 0x00, 0x76, 0x01, 0xa0, 0x29, 0x61, 0x2a, 0xd7, 0xaa, 0x59, 0xaf, 0xac, 0x5a, 0xad, 0x35,
	0x34, 0xad, 0xc4, 0xa8, 0x74, 0xd7, 0x9e, 0xa6, 0x7b, 0x25, 0x25, 0x78, 0x1f, 0x5c, 0x4f, 0x31,
	0x13, 0x21, 0x8e, 0xfc, 0x23, 0x1c, 0x46, 0x19, 0x23, 0xd5, 0x4a, 0xdd, 0x58, 0xbd, 0xea, 0x2d,
	0x16, 0xee, 0x2d, 0xe5, 0x95, 0x49, 0x0e, 0x71, 0x14, 0x06, 0x58, 0x10, 0x9f, 0x26, 0xd1, 0x49,
	0xf5, 0xff, 0x1c, 0xb6, 0xa0, 0x9d, 0x7b, 0x49, 0x74, 0x62, 0xbf, 0x35, 0xc1, 0xcd, 0xcb, 0xae,
	0x84, 0x6d, 0x60, 0x65, 0x69, 0xce, 0x95, 0x9d, 0xc8, 0xb9, 0x56, 0xab, 0xa6, 0xe3, 0xd7, 0xad,
	0x40, 0x5b, 0xb2, 0x59, 0xbb, 0x98, 0x1f, 0x7b, 0x40, 0xc1, 0xa5, 0x0d, 0x77, 0xc0, 0x5c, 0x9f,
	0x11, 0x2c, 0x54, 0x61, 0xad, 0x16, 0x9a, 0x98, 0xf7, 0x68, 0x70, 0xca, 0x89, 0xef, 0xfc, 0xe7,
	0x15, 0x7c, 0xa9, 0xa4, 0x74, 0xab, 0xe6, 0xac, 0x4a, 0x8a, 0x0f, 0xab, 0x60, 0x8e, 0x91, 0x98,
	0x0e, 0x55, 0xb9, 0xe6, 0xe5, 0x89, 0xfa, 0xde, 0xb4, 0xc0, 0xfc, 0xa8, 0xbe, 0xf6, 0x67, 0x03,
	0xd4, 0x2e, 0xeb, 0x3a, 0x4f, 0x69, 0xc2, 0x09, 0xdc, 0x02, 0x4b, 0xe7, 0xaa, 0xef, 0x13, 0xc6,
	0x28, 0xcb, 0x45, 0xad, 0x16, 0xd4, 0xe1, 0xb1, 0xb4, 0x8f, 0x0e, 0xf2, 0x59, 0xf4, 0x6e, 0x8c,
	0xf7, 0xe5, 0x99, 0x84, 0xc3, 0x03, 0x70, 0x85, 0x11, 0x9e, 0x45, 0x42, 0x8f, 0xc6, 0x93, 0xe9,
	0xa3, 0x71, 0x21, 0x2c, 0x2f, 0x57, 0xf0, 0xb4, 0x92, 0xfd, 0x14, 0x2c, 0x4f, 0xc0, 0xfc, 0xd5,
	0xc4, 0xb7, 0xde, 0x57, 0x00, 0x2c, 0x51, 0x0f, 0xd4, 0xc5, 0xf0, 0x93, 0x01, 0x16, 0xc7, 0xf7,
	0x08, 0x3e, 0x9e, 0x1e, 0xed, 0xa5, 0x9b, 0x57, 0xfb, 0xc7, 0xfe, 0xd9, 0x6b, 0x6f, 0xbe, 0xfd,
	0x78, 0x67, 0x36, 0x21, 0x92, 0xaf, 0xcc, 0xe9, 0x58, 0x0a, 0x1b, 0x7a, 0xd9, 0xb8, 0xd3, 0x70,
	0x70, 0xa9, 0x59, 0x4e, 0xe3, 0x0c, 0x7e, 0x35, 0x00, 0xbc, 0xd8, 0x46, 0xd8, 0x9e, 0xa1, 0xca,
	0x7a, 0xe5, 0x6b, 0xeb, 0xb3, 0x91, 0xd5, 0xe4, 0xd8, 0xeb, 0x79, 0x26, 0x6b, 0xf6, 0x03, 0x99,
	0xc9, 0x9f, 0xd0, 0x4f, 0x4b, 0xaf, 0xc8, 0x46, 0xe3, 0x6c, 0x2c, 0x11, 0x37, 0xce, 0xe5, 0x5c,
	0xa3, 0xb1, 0xf9, 0xcb, 0x00, 0x2b, 0x7d, 0x1a, 0x4f, 0x8d, 0x60, 0x73, 0xf9, 0x62, 0x03, 0xf7,
	0xe5, 0xb2, 0xee, 0x1b, 0x2f, 0x77, 0x0a, 0xf2, 0x80, 0x46, 0x38, 0x19, 0x20, 0xca, 0x06, 0xce,
	0x80, 0x24, 0xf9, 0x2a, 0xeb, 0xd7, 0x3c, 0x0d, 0xf9, 0xe4, 0x7f, 0x47, 0x5b, 0x1b, 0x1f, 0xcc,
	0xca, 0x76, 0xa7, 0xf3, 0xd1, 0xac, 0x6f, 0x2b, 0xc1, 0x4e, 0xc0, 0x91, 0x32, 0xa5, 0xd5, 0x6d,
	0xa2, 0xe2, 0x62, 0xfe, 0x45, 0x43, 0x7a, 0x9d, 0x80, 0xf7, 0x46, 0x90, 0x5e, 0xb7, 0xd9, 0xd3,
	0x90, 0x9f, 0xe6, 0x8a, 0xf2, 0xbb, 0x6e, 0x27, 0xe0, 0xae, 0x3b, 0x02, 0xb9, 0x6e, 0xb7, 0xe9,
	0xba, 0x1a, 0x76, 0x38, 0x97, 0xc7, 0xf9, 0xf0, 0x77, 0x00, 0x00, 0x00, 0xff, 0xff, 0xdd, 0x1f,
	0x90, 0x74, 0xe2, 0x06, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// AdParameterServiceClient is the client API for AdParameterService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AdParameterServiceClient interface {
	// Returns the requested ad parameter in full detail.
	GetAdParameter(ctx context.Context, in *GetAdParameterRequest, opts ...grpc.CallOption) (*resources.AdParameter, error)
	// Creates, updates, or removes ad parameters. Operation statuses are
	// returned.
	MutateAdParameters(ctx context.Context, in *MutateAdParametersRequest, opts ...grpc.CallOption) (*MutateAdParametersResponse, error)
}

type adParameterServiceClient struct {
	cc *grpc.ClientConn
}

func NewAdParameterServiceClient(cc *grpc.ClientConn) AdParameterServiceClient {
	return &adParameterServiceClient{cc}
}

func (c *adParameterServiceClient) GetAdParameter(ctx context.Context, in *GetAdParameterRequest, opts ...grpc.CallOption) (*resources.AdParameter, error) {
	out := new(resources.AdParameter)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v0.services.AdParameterService/GetAdParameter", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adParameterServiceClient) MutateAdParameters(ctx context.Context, in *MutateAdParametersRequest, opts ...grpc.CallOption) (*MutateAdParametersResponse, error) {
	out := new(MutateAdParametersResponse)
	err := c.cc.Invoke(ctx, "/google.ads.googleads.v0.services.AdParameterService/MutateAdParameters", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AdParameterServiceServer is the server API for AdParameterService service.
type AdParameterServiceServer interface {
	// Returns the requested ad parameter in full detail.
	GetAdParameter(context.Context, *GetAdParameterRequest) (*resources.AdParameter, error)
	// Creates, updates, or removes ad parameters. Operation statuses are
	// returned.
	MutateAdParameters(context.Context, *MutateAdParametersRequest) (*MutateAdParametersResponse, error)
}

func RegisterAdParameterServiceServer(s *grpc.Server, srv AdParameterServiceServer) {
	s.RegisterService(&_AdParameterService_serviceDesc, srv)
}

func _AdParameterService_GetAdParameter_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAdParameterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdParameterServiceServer).GetAdParameter(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v0.services.AdParameterService/GetAdParameter",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdParameterServiceServer).GetAdParameter(ctx, req.(*GetAdParameterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdParameterService_MutateAdParameters_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MutateAdParametersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdParameterServiceServer).MutateAdParameters(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.ads.googleads.v0.services.AdParameterService/MutateAdParameters",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdParameterServiceServer).MutateAdParameters(ctx, req.(*MutateAdParametersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _AdParameterService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.ads.googleads.v0.services.AdParameterService",
	HandlerType: (*AdParameterServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAdParameter",
			Handler:    _AdParameterService_GetAdParameter_Handler,
		},
		{
			MethodName: "MutateAdParameters",
			Handler:    _AdParameterService_MutateAdParameters_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/ads/googleads/v0/services/ad_parameter_service.proto",
}
