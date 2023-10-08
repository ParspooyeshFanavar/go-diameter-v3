// Code generated by protoc-gen-go. DO NOT EDIT.
// source: s6a.proto

/*
Package protos is a generated protocol buffer package.

It is generated from these files:

	s6a.proto

It has these top-level messages:

	AuthenticationInformationRequest
	AuthenticationInformationAnswer
	UpdateLocationRequest
	UpdateLocationAnswer
	CancelLocationRequest
	CancelLocationAnswer
	PurgeUERequest
	PurgeUEAnswer
*/
package protos

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
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

// ErrorCode reflects Experimental-Result values which are 3GPP failures
// to be processed by EPC. Diameter Base Protocol errors are reflected in gRPC status code
type ErrorCode int32

const (
	ErrorCode_UNDEFINED ErrorCode = 0
	// Default success code
	ErrorCode_MULTI_ROUND_AUTH        ErrorCode = 1001
	ErrorCode_SUCCESS                 ErrorCode = 2001
	ErrorCode_LIMITED_SUCCESS         ErrorCode = 2002
	ErrorCode_COMMAND_UNSUPORTED      ErrorCode = 3001
	ErrorCode_UNABLE_TO_DELIVER       ErrorCode = 3002
	ErrorCode_REALM_NOT_SERVED        ErrorCode = 3003
	ErrorCode_TOO_BUSY                ErrorCode = 3004
	ErrorCode_LOOP_DETECTED           ErrorCode = 3005
	ErrorCode_REDIRECT_INDICATION     ErrorCode = 3006
	ErrorCode_APPLICATION_UNSUPPORTED ErrorCode = 3007
	ErrorCode_INVALIDH_DR_BITS        ErrorCode = 3008
	ErrorCode_INVALID_AVP_BITS        ErrorCode = 3009
	ErrorCode_UNKNOWN_PEER            ErrorCode = 3010
	ErrorCode_AUTHENTICATION_REJECTED ErrorCode = 4001
	ErrorCode_OUT_OF_SPACE            ErrorCode = 4002
	ErrorCode_ELECTION_LOST           ErrorCode = 4003
	// Permanent Failures 7.4.3
	ErrorCode_USER_UNKNOWN             ErrorCode = 5001
	ErrorCode_UNKNOWN_EPS_SUBSCRIPTION ErrorCode = 5420
	ErrorCode_RAT_NOT_ALLOWED          ErrorCode = 5421
	ErrorCode_ROAMING_NOT_ALLOWED      ErrorCode = 5004
	ErrorCode_EQUIPMENT_UNKNOWN        ErrorCode = 5422
	ErrorCode_UNKOWN_SERVING_NODE      ErrorCode = 5423
	// Transient Failures 7.4.4
	ErrorCode_AUTHENTICATION_DATA_UNAVAILABLE ErrorCode = 4181
)

var ErrorCode_name = map[int32]string{
	0:    "UNDEFINED",
	1001: "MULTI_ROUND_AUTH",
	2001: "SUCCESS",
	2002: "LIMITED_SUCCESS",
	3001: "COMMAND_UNSUPORTED",
	3002: "UNABLE_TO_DELIVER",
	3003: "REALM_NOT_SERVED",
	3004: "TOO_BUSY",
	3005: "LOOP_DETECTED",
	3006: "REDIRECT_INDICATION",
	3007: "APPLICATION_UNSUPPORTED",
	3008: "INVALIDH_DR_BITS",
	3009: "INVALID_AVP_BITS",
	3010: "UNKNOWN_PEER",
	4001: "AUTHENTICATION_REJECTED",
	4002: "OUT_OF_SPACE",
	4003: "ELECTION_LOST",
	5001: "USER_UNKNOWN",
	5420: "UNKNOWN_EPS_SUBSCRIPTION",
	5421: "RAT_NOT_ALLOWED",
	5004: "ROAMING_NOT_ALLOWED",
	5422: "EQUIPMENT_UNKNOWN",
	5423: "UNKOWN_SERVING_NODE",
	4181: "AUTHENTICATION_DATA_UNAVAILABLE",
}
var ErrorCode_value = map[string]int32{
	"UNDEFINED":                       0,
	"MULTI_ROUND_AUTH":                1001,
	"SUCCESS":                         2001,
	"LIMITED_SUCCESS":                 2002,
	"COMMAND_UNSUPORTED":              3001,
	"UNABLE_TO_DELIVER":               3002,
	"REALM_NOT_SERVED":                3003,
	"TOO_BUSY":                        3004,
	"LOOP_DETECTED":                   3005,
	"REDIRECT_INDICATION":             3006,
	"APPLICATION_UNSUPPORTED":         3007,
	"INVALIDH_DR_BITS":                3008,
	"INVALID_AVP_BITS":                3009,
	"UNKNOWN_PEER":                    3010,
	"AUTHENTICATION_REJECTED":         4001,
	"OUT_OF_SPACE":                    4002,
	"ELECTION_LOST":                   4003,
	"USER_UNKNOWN":                    5001,
	"UNKNOWN_EPS_SUBSCRIPTION":        5420,
	"RAT_NOT_ALLOWED":                 5421,
	"ROAMING_NOT_ALLOWED":             5004,
	"EQUIPMENT_UNKNOWN":               5422,
	"UNKOWN_SERVING_NODE":             5423,
	"AUTHENTICATION_DATA_UNAVAILABLE": 4181,
}

func (x ErrorCode) String() string {
	return proto.EnumName(ErrorCode_name, int32(x))
}
func (ErrorCode) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type CancelLocationRequest_CancellationType int32

const (
	CancelLocationRequest_MME_UPDATE_PROCEDURE     CancelLocationRequest_CancellationType = 0
	CancelLocationRequest_SGSN_UPDATE_PROCEDURE    CancelLocationRequest_CancellationType = 1
	CancelLocationRequest_SUBSCRIPTION_WITHDRAWAL  CancelLocationRequest_CancellationType = 2
	CancelLocationRequest_UPDATE_PROCEDURE_IWF     CancelLocationRequest_CancellationType = 3
	CancelLocationRequest_INITIAL_ATTACH_PROCEDURE CancelLocationRequest_CancellationType = 4
)

var CancelLocationRequest_CancellationType_name = map[int32]string{
	0: "MME_UPDATE_PROCEDURE",
	1: "SGSN_UPDATE_PROCEDURE",
	2: "SUBSCRIPTION_WITHDRAWAL",
	3: "UPDATE_PROCEDURE_IWF",
	4: "INITIAL_ATTACH_PROCEDURE",
}
var CancelLocationRequest_CancellationType_value = map[string]int32{
	"MME_UPDATE_PROCEDURE":     0,
	"SGSN_UPDATE_PROCEDURE":    1,
	"SUBSCRIPTION_WITHDRAWAL":  2,
	"UPDATE_PROCEDURE_IWF":     3,
	"INITIAL_ATTACH_PROCEDURE": 4,
}

func (x CancelLocationRequest_CancellationType) String() string {
	return proto.EnumName(CancelLocationRequest_CancellationType_name, int32(x))
}
func (CancelLocationRequest_CancellationType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor0, []int{4, 0}
}

// Authentication Information Request (Section 7.2.5)
type AuthenticationInformationRequest struct {
	// Subscriber identifier
	UserName string `protobuf:"bytes,1,opt,name=user_name,json=userName" json:"user_name,omitempty"`
	// Visted site identifier
	VisitedPlmn []byte `protobuf:"bytes,2,opt,name=visited_plmn,json=visitedPlmn,proto3" json:"visited_plmn,omitempty"`
	// Number of vectors to request in response
	NumRequestedEutranVectors uint32 `protobuf:"varint,3,opt,name=num_requested_eutran_vectors,json=numRequestedEutranVectors" json:"num_requested_eutran_vectors,omitempty"`
	// Indicates to the HSS the values are requested for immediate attach
	ImmediateResponsePreferred bool `protobuf:"varint,4,opt,name=immediate_response_preferred,json=immediateResponsePreferred" json:"immediate_response_preferred,omitempty"`
	// Concatenation of RAND and AUTS in the case of a resync attach case
	ResyncInfo []byte `protobuf:"bytes,5,opt,name=resync_info,json=resyncInfo,proto3" json:"resync_info,omitempty"`
}

func (m *AuthenticationInformationRequest) Reset()         { *m = AuthenticationInformationRequest{} }
func (m *AuthenticationInformationRequest) String() string { return proto.CompactTextString(m) }
func (*AuthenticationInformationRequest) ProtoMessage()    {}
func (*AuthenticationInformationRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor0, []int{0}
}

func (m *AuthenticationInformationRequest) GetUserName() string {
	if m != nil {
		return m.UserName
	}
	return ""
}

func (m *AuthenticationInformationRequest) GetVisitedPlmn() []byte {
	if m != nil {
		return m.VisitedPlmn
	}
	return nil
}

func (m *AuthenticationInformationRequest) GetNumRequestedEutranVectors() uint32 {
	if m != nil {
		return m.NumRequestedEutranVectors
	}
	return 0
}

func (m *AuthenticationInformationRequest) GetImmediateResponsePreferred() bool {
	if m != nil {
		return m.ImmediateResponsePreferred
	}
	return false
}

func (m *AuthenticationInformationRequest) GetResyncInfo() []byte {
	if m != nil {
		return m.ResyncInfo
	}
	return nil
}

// Authentication Information Answer (Section 7.2.6)
type AuthenticationInformationAnswer struct {
	// EPC error code on failure
	ErrorCode ErrorCode `protobuf:"varint,1,opt,name=error_code,json=errorCode,enum=ErrorCode" json:"error_code,omitempty"`
	// Authentication vectors matching the requested number
	EutranVectors []*AuthenticationInformationAnswer_EUTRANVector `protobuf:"bytes,2,rep,name=eutran_vectors,json=eutranVectors" json:"eutran_vectors,omitempty"`
}

func (m *AuthenticationInformationAnswer) Reset()         { *m = AuthenticationInformationAnswer{} }
func (m *AuthenticationInformationAnswer) String() string { return proto.CompactTextString(m) }
func (*AuthenticationInformationAnswer) ProtoMessage()    {}
func (*AuthenticationInformationAnswer) Descriptor() ([]byte, []int) {
	return fileDescriptor0, []int{1}
}

func (m *AuthenticationInformationAnswer) GetErrorCode() ErrorCode {
	if m != nil {
		return m.ErrorCode
	}
	return ErrorCode_UNDEFINED
}

func (m *AuthenticationInformationAnswer) GetEutranVectors() []*AuthenticationInformationAnswer_EUTRANVector {
	if m != nil {
		return m.EutranVectors
	}
	return nil
}

// For details about fields read 3GPP 33.401
type AuthenticationInformationAnswer_EUTRANVector struct {
	Rand  []byte `protobuf:"bytes,1,opt,name=rand,proto3" json:"rand,omitempty"`
	Xres  []byte `protobuf:"bytes,2,opt,name=xres,proto3" json:"xres,omitempty"`
	Autn  []byte `protobuf:"bytes,3,opt,name=autn,proto3" json:"autn,omitempty"`
	Kasme []byte `protobuf:"bytes,4,opt,name=kasme,proto3" json:"kasme,omitempty"`
}

func (m *AuthenticationInformationAnswer_EUTRANVector) Reset() {
	*m = AuthenticationInformationAnswer_EUTRANVector{}
}
func (m *AuthenticationInformationAnswer_EUTRANVector) String() string {
	return proto.CompactTextString(m)
}
func (*AuthenticationInformationAnswer_EUTRANVector) ProtoMessage() {}
func (*AuthenticationInformationAnswer_EUTRANVector) Descriptor() ([]byte, []int) {
	return fileDescriptor0, []int{1, 0}
}

func (m *AuthenticationInformationAnswer_EUTRANVector) GetRand() []byte {
	if m != nil {
		return m.Rand
	}
	return nil
}

func (m *AuthenticationInformationAnswer_EUTRANVector) GetXres() []byte {
	if m != nil {
		return m.Xres
	}
	return nil
}

func (m *AuthenticationInformationAnswer_EUTRANVector) GetAutn() []byte {
	if m != nil {
		return m.Autn
	}
	return nil
}

func (m *AuthenticationInformationAnswer_EUTRANVector) GetKasme() []byte {
	if m != nil {
		return m.Kasme
	}
	return nil
}

// Update Location Request (Section 7.2.3)
type UpdateLocationRequest struct {
	// Subscriber identifier
	UserName string `protobuf:"bytes,1,opt,name=user_name,json=userName" json:"user_name,omitempty"`
	// Visited site identifier
	VisitedPlmn []byte `protobuf:"bytes,2,opt,name=visited_plmn,json=visitedPlmn,proto3" json:"visited_plmn,omitempty"`
	// Selective unrolling of ULR-Flags 29.272 Table 7.3.7/1
	// Skip subscription data in response
	SkipSubscriberData bool `protobuf:"varint,3,opt,name=skip_subscriber_data,json=skipSubscriberData" json:"skip_subscriber_data,omitempty"`
	// Send Cancel Location to other EPCs serving the UE
	InitialAttach bool `protobuf:"varint,4,opt,name=initial_attach,json=initialAttach" json:"initial_attach,omitempty"`
}

func (m *UpdateLocationRequest) Reset()                    { *m = UpdateLocationRequest{} }
func (m *UpdateLocationRequest) String() string            { return proto.CompactTextString(m) }
func (*UpdateLocationRequest) ProtoMessage()               {}
func (*UpdateLocationRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *UpdateLocationRequest) GetUserName() string {
	if m != nil {
		return m.UserName
	}
	return ""
}

func (m *UpdateLocationRequest) GetVisitedPlmn() []byte {
	if m != nil {
		return m.VisitedPlmn
	}
	return nil
}

func (m *UpdateLocationRequest) GetSkipSubscriberData() bool {
	if m != nil {
		return m.SkipSubscriberData
	}
	return false
}

func (m *UpdateLocationRequest) GetInitialAttach() bool {
	if m != nil {
		return m.InitialAttach
	}
	return false
}

// Update Location Answer (Section 7.2.4)
type UpdateLocationAnswer struct {
	// EPC error code on failure
	ErrorCode ErrorCode `protobuf:"varint,1,opt,name=error_code,json=errorCode,enum=ErrorCode" json:"error_code,omitempty"`
	// Identifier of the default APN
	DefaultContextId uint32 `protobuf:"varint,2,opt,name=default_context_id,json=defaultContextId" json:"default_context_id,omitempty"`
	// Subscriber authorized aggregate bitrate
	TotalAmbr *UpdateLocationAnswer_AggregatedMaximumBitrate `protobuf:"bytes,3,opt,name=total_ambr,json=totalAmbr" json:"total_ambr,omitempty"`
	// Indicates to wipe other stored APNs
	AllApnsIncluded bool `protobuf:"varint,4,opt,name=all_apns_included,json=allApnsIncluded" json:"all_apns_included,omitempty"`
	// APN configurations
	Apn []*UpdateLocationAnswer_APNConfiguration `protobuf:"bytes,5,rep,name=apn" json:"apn,omitempty"`
}

func (m *UpdateLocationAnswer) Reset()                    { *m = UpdateLocationAnswer{} }
func (m *UpdateLocationAnswer) String() string            { return proto.CompactTextString(m) }
func (*UpdateLocationAnswer) ProtoMessage()               {}
func (*UpdateLocationAnswer) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *UpdateLocationAnswer) GetErrorCode() ErrorCode {
	if m != nil {
		return m.ErrorCode
	}
	return ErrorCode_UNDEFINED
}

func (m *UpdateLocationAnswer) GetDefaultContextId() uint32 {
	if m != nil {
		return m.DefaultContextId
	}
	return 0
}

func (m *UpdateLocationAnswer) GetTotalAmbr() *UpdateLocationAnswer_AggregatedMaximumBitrate {
	if m != nil {
		return m.TotalAmbr
	}
	return nil
}

func (m *UpdateLocationAnswer) GetAllApnsIncluded() bool {
	if m != nil {
		return m.AllApnsIncluded
	}
	return false
}

func (m *UpdateLocationAnswer) GetApn() []*UpdateLocationAnswer_APNConfiguration {
	if m != nil {
		return m.Apn
	}
	return nil
}

type UpdateLocationAnswer_APNConfiguration struct {
	// APN identifier
	ContextId uint32 `protobuf:"varint,1,opt,name=context_id,json=contextId" json:"context_id,omitempty"`
	// Contains either the APN Name or wildcard "*"
	ServiceSelection string `protobuf:"bytes,2,opt,name=service_selection,json=serviceSelection" json:"service_selection,omitempty"`
	// APN QoS profile
	QosProfile *UpdateLocationAnswer_APNConfiguration_QoSProfile `protobuf:"bytes,3,opt,name=qos_profile,json=qosProfile" json:"qos_profile,omitempty"`
	// APN authorized bitrate
	Ambr *UpdateLocationAnswer_AggregatedMaximumBitrate `protobuf:"bytes,4,opt,name=ambr" json:"ambr,omitempty"`
}

func (m *UpdateLocationAnswer_APNConfiguration) Reset()         { *m = UpdateLocationAnswer_APNConfiguration{} }
func (m *UpdateLocationAnswer_APNConfiguration) String() string { return proto.CompactTextString(m) }
func (*UpdateLocationAnswer_APNConfiguration) ProtoMessage()    {}
func (*UpdateLocationAnswer_APNConfiguration) Descriptor() ([]byte, []int) {
	return fileDescriptor0, []int{3, 0}
}

func (m *UpdateLocationAnswer_APNConfiguration) GetContextId() uint32 {
	if m != nil {
		return m.ContextId
	}
	return 0
}

func (m *UpdateLocationAnswer_APNConfiguration) GetServiceSelection() string {
	if m != nil {
		return m.ServiceSelection
	}
	return ""
}

func (m *UpdateLocationAnswer_APNConfiguration) GetQosProfile() *UpdateLocationAnswer_APNConfiguration_QoSProfile {
	if m != nil {
		return m.QosProfile
	}
	return nil
}

func (m *UpdateLocationAnswer_APNConfiguration) GetAmbr() *UpdateLocationAnswer_AggregatedMaximumBitrate {
	if m != nil {
		return m.Ambr
	}
	return nil
}

// For details about values see 29.212
type UpdateLocationAnswer_APNConfiguration_QoSProfile struct {
	ClassId                 int32  `protobuf:"varint,1,opt,name=class_id,json=classId" json:"class_id,omitempty"`
	PriorityLevel           uint32 `protobuf:"varint,2,opt,name=priority_level,json=priorityLevel" json:"priority_level,omitempty"`
	PreemptionCapability    bool   `protobuf:"varint,3,opt,name=preemption_capability,json=preemptionCapability" json:"preemption_capability,omitempty"`
	PreemptionVulnerability bool   `protobuf:"varint,4,opt,name=preemption_vulnerability,json=preemptionVulnerability" json:"preemption_vulnerability,omitempty"`
}

func (m *UpdateLocationAnswer_APNConfiguration_QoSProfile) Reset() {
	*m = UpdateLocationAnswer_APNConfiguration_QoSProfile{}
}
func (m *UpdateLocationAnswer_APNConfiguration_QoSProfile) String() string {
	return proto.CompactTextString(m)
}
func (*UpdateLocationAnswer_APNConfiguration_QoSProfile) ProtoMessage() {}
func (*UpdateLocationAnswer_APNConfiguration_QoSProfile) Descriptor() ([]byte, []int) {
	return fileDescriptor0, []int{3, 0, 0}
}

func (m *UpdateLocationAnswer_APNConfiguration_QoSProfile) GetClassId() int32 {
	if m != nil {
		return m.ClassId
	}
	return 0
}

func (m *UpdateLocationAnswer_APNConfiguration_QoSProfile) GetPriorityLevel() uint32 {
	if m != nil {
		return m.PriorityLevel
	}
	return 0
}

func (m *UpdateLocationAnswer_APNConfiguration_QoSProfile) GetPreemptionCapability() bool {
	if m != nil {
		return m.PreemptionCapability
	}
	return false
}

func (m *UpdateLocationAnswer_APNConfiguration_QoSProfile) GetPreemptionVulnerability() bool {
	if m != nil {
		return m.PreemptionVulnerability
	}
	return false
}

type UpdateLocationAnswer_AggregatedMaximumBitrate struct {
	// Maximum uplink bitrate
	MaxBandwidthUl uint32 `protobuf:"varint,1,opt,name=max_bandwidth_ul,json=maxBandwidthUl" json:"max_bandwidth_ul,omitempty"`
	// Maximum downlink bitrate
	MaxBandwidthDl uint32 `protobuf:"varint,2,opt,name=max_bandwidth_dl,json=maxBandwidthDl" json:"max_bandwidth_dl,omitempty"`
}

func (m *UpdateLocationAnswer_AggregatedMaximumBitrate) Reset() {
	*m = UpdateLocationAnswer_AggregatedMaximumBitrate{}
}
func (m *UpdateLocationAnswer_AggregatedMaximumBitrate) String() string {
	return proto.CompactTextString(m)
}
func (*UpdateLocationAnswer_AggregatedMaximumBitrate) ProtoMessage() {}
func (*UpdateLocationAnswer_AggregatedMaximumBitrate) Descriptor() ([]byte, []int) {
	return fileDescriptor0, []int{3, 1}
}

func (m *UpdateLocationAnswer_AggregatedMaximumBitrate) GetMaxBandwidthUl() uint32 {
	if m != nil {
		return m.MaxBandwidthUl
	}
	return 0
}

func (m *UpdateLocationAnswer_AggregatedMaximumBitrate) GetMaxBandwidthDl() uint32 {
	if m != nil {
		return m.MaxBandwidthDl
	}
	return 0
}

// Cancel Location Request (Section 7.2.7)
type CancelLocationRequest struct {
	// Subscriber identifier
	UserName string `protobuf:"bytes,1,opt,name=user_name,json=userName" json:"user_name,omitempty"`
	// Type of cancellation 7.3.24
	CancellationType CancelLocationRequest_CancellationType `protobuf:"varint,2,opt,name=cancellation_type,json=cancellationType,enum=CancelLocationRequest_CancellationType" json:"cancellation_type,omitempty"`
}

func (m *CancelLocationRequest) Reset()                    { *m = CancelLocationRequest{} }
func (m *CancelLocationRequest) String() string            { return proto.CompactTextString(m) }
func (*CancelLocationRequest) ProtoMessage()               {}
func (*CancelLocationRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *CancelLocationRequest) GetUserName() string {
	if m != nil {
		return m.UserName
	}
	return ""
}

func (m *CancelLocationRequest) GetCancellationType() CancelLocationRequest_CancellationType {
	if m != nil {
		return m.CancellationType
	}
	return CancelLocationRequest_MME_UPDATE_PROCEDURE
}

// Cancel Location Answer (Section 7.2.8)
type CancelLocationAnswer struct {
	// EPC error code on failure
	ErrorCode ErrorCode `protobuf:"varint,1,opt,name=error_code,json=errorCode,enum=ErrorCode" json:"error_code,omitempty"`
}

func (m *CancelLocationAnswer) Reset()                    { *m = CancelLocationAnswer{} }
func (m *CancelLocationAnswer) String() string            { return proto.CompactTextString(m) }
func (*CancelLocationAnswer) ProtoMessage()               {}
func (*CancelLocationAnswer) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *CancelLocationAnswer) GetErrorCode() ErrorCode {
	if m != nil {
		return m.ErrorCode
	}
	return ErrorCode_UNDEFINED
}

// Purge UE Request (Section 7.2.14)
type PurgeUERequest struct {
	// Subscriber identifier
	UserName string `protobuf:"bytes,1,opt,name=user_name,json=userName" json:"user_name,omitempty"`
}

func (m *PurgeUERequest) Reset()                    { *m = PurgeUERequest{} }
func (m *PurgeUERequest) String() string            { return proto.CompactTextString(m) }
func (*PurgeUERequest) ProtoMessage()               {}
func (*PurgeUERequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *PurgeUERequest) GetUserName() string {
	if m != nil {
		return m.UserName
	}
	return ""
}

// Purge UE Answer (Section 7.2.15)
type PurgeUEAnswer struct {
	// EPC error code on failure
	ErrorCode ErrorCode `protobuf:"varint,1,opt,name=error_code,json=errorCode,enum=ErrorCode" json:"error_code,omitempty"`
}

func (m *PurgeUEAnswer) Reset()                    { *m = PurgeUEAnswer{} }
func (m *PurgeUEAnswer) String() string            { return proto.CompactTextString(m) }
func (*PurgeUEAnswer) ProtoMessage()               {}
func (*PurgeUEAnswer) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *PurgeUEAnswer) GetErrorCode() ErrorCode {
	if m != nil {
		return m.ErrorCode
	}
	return ErrorCode_UNDEFINED
}

func init() {
	proto.RegisterType((*AuthenticationInformationRequest)(nil), "AuthenticationInformationRequest")
	proto.RegisterType((*AuthenticationInformationAnswer)(nil), "AuthenticationInformationAnswer")
	proto.RegisterType((*AuthenticationInformationAnswer_EUTRANVector)(nil), "AuthenticationInformationAnswer.EUTRANVector")
	proto.RegisterType((*UpdateLocationRequest)(nil), "UpdateLocationRequest")
	proto.RegisterType((*UpdateLocationAnswer)(nil), "UpdateLocationAnswer")
	proto.RegisterType((*UpdateLocationAnswer_APNConfiguration)(nil), "UpdateLocationAnswer.APNConfiguration")
	proto.RegisterType((*UpdateLocationAnswer_APNConfiguration_QoSProfile)(nil), "UpdateLocationAnswer.APNConfiguration.QoSProfile")
	proto.RegisterType((*UpdateLocationAnswer_AggregatedMaximumBitrate)(nil), "UpdateLocationAnswer.AggregatedMaximumBitrate")
	proto.RegisterType((*CancelLocationRequest)(nil), "CancelLocationRequest")
	proto.RegisterType((*CancelLocationAnswer)(nil), "CancelLocationAnswer")
	proto.RegisterType((*PurgeUERequest)(nil), "PurgeUERequest")
	proto.RegisterType((*PurgeUEAnswer)(nil), "PurgeUEAnswer")
	proto.RegisterEnum("ErrorCode", ErrorCode_name, ErrorCode_value)
	proto.RegisterEnum("CancelLocationRequest_CancellationType", CancelLocationRequest_CancellationType_name, CancelLocationRequest_CancellationType_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for S6AProxy service

type S6AProxyClient interface {
	// Authentication-Information (Code 318)
	AuthenticationInformation(ctx context.Context, in *AuthenticationInformationRequest, opts ...grpc.CallOption) (*AuthenticationInformationAnswer, error)
	// Update-Location (Code 316)
	UpdateLocation(ctx context.Context, in *UpdateLocationRequest, opts ...grpc.CallOption) (*UpdateLocationAnswer, error)
	// Cancel-Location (Code 317)
	CancelLocation(ctx context.Context, in *CancelLocationRequest, opts ...grpc.CallOption) (*CancelLocationAnswer, error)
	// Purge-UE (Code 321)
	PurgeUE(ctx context.Context, in *PurgeUERequest, opts ...grpc.CallOption) (*PurgeUEAnswer, error)
}

type s6AProxyClient struct {
	cc *grpc.ClientConn
}

func NewS6AProxyClient(cc *grpc.ClientConn) S6AProxyClient {
	return &s6AProxyClient{cc}
}

func (c *s6AProxyClient) AuthenticationInformation(ctx context.Context, in *AuthenticationInformationRequest, opts ...grpc.CallOption) (*AuthenticationInformationAnswer, error) {
	out := new(AuthenticationInformationAnswer)
	err := grpc.Invoke(ctx, "/S6aProxy/AuthenticationInformation", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *s6AProxyClient) UpdateLocation(ctx context.Context, in *UpdateLocationRequest, opts ...grpc.CallOption) (*UpdateLocationAnswer, error) {
	out := new(UpdateLocationAnswer)
	err := grpc.Invoke(ctx, "/S6aProxy/UpdateLocation", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *s6AProxyClient) CancelLocation(ctx context.Context, in *CancelLocationRequest, opts ...grpc.CallOption) (*CancelLocationAnswer, error) {
	out := new(CancelLocationAnswer)
	err := grpc.Invoke(ctx, "/S6aProxy/CancelLocation", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *s6AProxyClient) PurgeUE(ctx context.Context, in *PurgeUERequest, opts ...grpc.CallOption) (*PurgeUEAnswer, error) {
	out := new(PurgeUEAnswer)
	err := grpc.Invoke(ctx, "/S6aProxy/PurgeUE", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for S6AProxy service

type S6AProxyServer interface {
	// Authentication-Information (Code 318)
	AuthenticationInformation(context.Context, *AuthenticationInformationRequest) (*AuthenticationInformationAnswer, error)
	// Update-Location (Code 316)
	UpdateLocation(context.Context, *UpdateLocationRequest) (*UpdateLocationAnswer, error)
	// Cancel-Location (Code 317)
	CancelLocation(context.Context, *CancelLocationRequest) (*CancelLocationAnswer, error)
	// Purge-UE (Code 321)
	PurgeUE(context.Context, *PurgeUERequest) (*PurgeUEAnswer, error)
}

func RegisterS6AProxyServer(s *grpc.Server, srv S6AProxyServer) {
	s.RegisterService(&_S6AProxy_serviceDesc, srv)
}

func _S6AProxy_AuthenticationInformation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuthenticationInformationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(S6AProxyServer).AuthenticationInformation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/S6aProxy/AuthenticationInformation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(S6AProxyServer).AuthenticationInformation(ctx, req.(*AuthenticationInformationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _S6AProxy_UpdateLocation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateLocationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(S6AProxyServer).UpdateLocation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/S6aProxy/UpdateLocation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(S6AProxyServer).UpdateLocation(ctx, req.(*UpdateLocationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _S6AProxy_CancelLocation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CancelLocationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(S6AProxyServer).CancelLocation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/S6aProxy/CancelLocation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(S6AProxyServer).CancelLocation(ctx, req.(*CancelLocationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _S6AProxy_PurgeUE_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PurgeUERequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(S6AProxyServer).PurgeUE(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/S6aProxy/PurgeUE",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(S6AProxyServer).PurgeUE(ctx, req.(*PurgeUERequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _S6AProxy_serviceDesc = grpc.ServiceDesc{
	ServiceName: "S6aProxy",
	HandlerType: (*S6AProxyServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AuthenticationInformation",
			Handler:    _S6AProxy_AuthenticationInformation_Handler,
		},
		{
			MethodName: "UpdateLocation",
			Handler:    _S6AProxy_UpdateLocation_Handler,
		},
		{
			MethodName: "CancelLocation",
			Handler:    _S6AProxy_CancelLocation_Handler,
		},
		{
			MethodName: "PurgeUE",
			Handler:    _S6AProxy_PurgeUE_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "s6a.proto",
}

func init() { proto.RegisterFile("s6a.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 1398 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x56, 0xdb, 0x6e, 0xe3, 0xc6,
	0x19, 0xb6, 0x6c, 0x6f, 0xd6, 0xfa, 0x2d, 0xc9, 0xf4, 0xd4, 0x5e, 0xc9, 0x8a, 0x03, 0x2b, 0x42,
	0xdb, 0x75, 0xdd, 0x44, 0x68, 0x1d, 0x20, 0x68, 0x7b, 0xd3, 0x52, 0xe4, 0x6c, 0xcc, 0x86, 0x22,
	0x99, 0x21, 0x69, 0xa3, 0xbd, 0x99, 0x8e, 0xc8, 0xb1, 0x97, 0x08, 0x4f, 0x3b, 0xa4, 0x1c, 0xfb,
	0xbe, 0x28, 0x50, 0xa0, 0x8f, 0xd0, 0x9b, 0xa6, 0x57, 0x05, 0x7a, 0xba, 0xeb, 0x39, 0x3d, 0x3c,
	0x40, 0x81, 0x16, 0xe8, 0x43, 0xf4, 0x2d, 0x0a, 0x1e, 0xe4, 0x95, 0x1c, 0x1b, 0x1b, 0x17, 0xb9,
	0xf2, 0xe8, 0xfb, 0xfe, 0xe3, 0xf7, 0xff, 0x43, 0x0f, 0x34, 0xb3, 0x77, 0xd9, 0x28, 0x15, 0x49,
	0x9e, 0x0c, 0x7f, 0xb4, 0x0a, 0x03, 0x79, 0x96, 0x3f, 0xe7, 0x71, 0x1e, 0x78, 0x2c, 0x0f, 0x92,
	0x58, 0x8b, 0xcf, 0x13, 0x11, 0x95, 0x47, 0xc2, 0x5f, 0xcc, 0x78, 0x96, 0xa3, 0xd7, 0xa1, 0x39,
	0xcb, 0xb8, 0xa0, 0x31, 0x8b, 0x78, 0xaf, 0x31, 0x68, 0x1c, 0x36, 0xc9, 0x46, 0x01, 0x18, 0x2c,
	0xe2, 0xe8, 0x4d, 0x68, 0x5d, 0x06, 0x59, 0x90, 0x73, 0x9f, 0xa6, 0x61, 0x14, 0xf7, 0x56, 0x07,
	0x8d, 0xc3, 0x16, 0xd9, 0xac, 0x31, 0x2b, 0x8c, 0x62, 0xf4, 0x6d, 0xd8, 0x8f, 0x67, 0x11, 0x15,
	0x55, 0x38, 0xee, 0x53, 0x3e, 0xcb, 0x05, 0x8b, 0xe9, 0x25, 0xf7, 0xf2, 0x44, 0x64, 0xbd, 0xb5,
	0x41, 0xe3, 0xb0, 0x4d, 0xf6, 0xe2, 0x59, 0x44, 0xe6, 0x26, 0xb8, 0xb4, 0x38, 0xad, 0x0c, 0xd0,
	0x77, 0x60, 0x3f, 0x88, 0x22, 0xee, 0x07, 0x2c, 0xe7, 0x54, 0xf0, 0x2c, 0x4d, 0xe2, 0x8c, 0xd3,
	0x54, 0xf0, 0x73, 0x2e, 0x04, 0xf7, 0x7b, 0xeb, 0x83, 0xc6, 0xe1, 0x06, 0xe9, 0xdf, 0xd8, 0x90,
	0xda, 0xc4, 0x9a, 0x5b, 0xa0, 0x03, 0xd8, 0x14, 0x3c, 0xbb, 0x8e, 0x3d, 0x1a, 0xc4, 0xe7, 0x49,
	0xef, 0x51, 0x59, 0x24, 0x54, 0x50, 0xd1, 0xf1, 0xf0, 0x87, 0xab, 0x70, 0x70, 0xaf, 0x10, 0x72,
	0x9c, 0x7d, 0xc4, 0x05, 0xfa, 0x0a, 0x00, 0x17, 0x22, 0x11, 0xd4, 0x4b, 0xfc, 0x4a, 0x88, 0xce,
	0x31, 0x8c, 0x70, 0x01, 0x29, 0x89, 0xcf, 0x49, 0x93, 0xcf, 0x8f, 0xc8, 0x81, 0xce, 0xad, 0x26,
	0x57, 0x07, 0x6b, 0x87, 0x9b, 0xc7, 0x6f, 0x8f, 0x5e, 0x91, 0x64, 0x84, 0x5d, 0x87, 0xc8, 0x46,
	0xd5, 0x39, 0x69, 0xf3, 0x45, 0x1d, 0xfa, 0x3f, 0x80, 0xd6, 0x22, 0x8d, 0x10, 0xac, 0x0b, 0x16,
	0xfb, 0x65, 0x29, 0x2d, 0x52, 0x9e, 0x0b, 0xec, 0x4a, 0xf0, 0xac, 0x9e, 0x43, 0x79, 0x2e, 0x30,
	0x36, 0xcb, 0xe3, 0x52, 0xe8, 0x16, 0x29, 0xcf, 0x68, 0x07, 0x1e, 0x7d, 0xc8, 0xb2, 0x88, 0x97,
	0xe2, 0xb5, 0x48, 0xf5, 0x63, 0xf8, 0xdb, 0x06, 0xec, 0xba, 0xa9, 0xcf, 0x72, 0xae, 0x27, 0xde,
	0xe7, 0xba, 0x04, 0x5f, 0x83, 0x9d, 0xec, 0xc3, 0x20, 0xa5, 0xd9, 0x6c, 0x9a, 0x79, 0x22, 0x98,
	0x72, 0x41, 0x7d, 0x96, 0xb3, 0xb2, 0xa6, 0x0d, 0x82, 0x0a, 0xce, 0xbe, 0xa1, 0x54, 0x96, 0x33,
	0xf4, 0x25, 0xe8, 0x04, 0x71, 0x90, 0x07, 0x2c, 0xa4, 0x2c, 0xcf, 0x99, 0xf7, 0xbc, 0x9e, 0x73,
	0xbb, 0x46, 0xe5, 0x12, 0x1c, 0xfe, 0xf3, 0x35, 0xd8, 0x59, 0x2e, 0xf9, 0xe1, 0xe3, 0x7a, 0x0b,
	0x90, 0xcf, 0xcf, 0xd9, 0x2c, 0xcc, 0xa9, 0x97, 0xc4, 0x39, 0xbf, 0xca, 0x69, 0xe0, 0x97, 0x5d,
	0xb4, 0x89, 0x54, 0x33, 0x4a, 0x45, 0x68, 0x3e, 0x9a, 0x00, 0xe4, 0x49, 0x5e, 0x94, 0x15, 0x4d,
	0x45, 0xd9, 0xc0, 0xe6, 0xf1, 0x68, 0x74, 0x57, 0x0d, 0x23, 0xf9, 0xe2, 0x42, 0xf0, 0x0b, 0x96,
	0x73, 0x7f, 0xc2, 0xae, 0x82, 0x68, 0x16, 0x8d, 0x83, 0x5c, 0x14, 0x1b, 0xda, 0x2c, 0x23, 0xc8,
	0xd1, 0x54, 0xa0, 0x23, 0xd8, 0x66, 0x61, 0x48, 0x59, 0x1a, 0x67, 0x34, 0x88, 0xbd, 0x70, 0xe6,
	0xdf, 0xac, 0xf4, 0x16, 0x0b, 0x43, 0x39, 0x8d, 0x33, 0xad, 0x86, 0xd1, 0x37, 0x60, 0x8d, 0xa5,
	0x71, 0xef, 0x51, 0xb9, 0x4c, 0x5f, 0xbe, 0x27, 0xa7, 0x65, 0x28, 0x49, 0x7c, 0x1e, 0x5c, 0xcc,
	0x44, 0x35, 0xc1, 0xc2, 0xa5, 0xff, 0xbb, 0x35, 0x90, 0x6e, 0x33, 0xe8, 0x0d, 0x80, 0x85, 0x7e,
	0x1b, 0x65, 0xbf, 0x4d, 0xef, 0xa6, 0xd1, 0xaf, 0xc2, 0x76, 0xc6, 0xc5, 0x65, 0xe0, 0x71, 0x9a,
	0xf1, 0x90, 0x7b, 0x85, 0x4f, 0xa9, 0x4a, 0x93, 0x48, 0x35, 0x61, 0xcf, 0x71, 0x44, 0x60, 0xf3,
	0x45, 0x92, 0xd1, 0x54, 0x24, 0xe7, 0x41, 0xc8, 0x6b, 0x59, 0xbe, 0xfe, 0xd9, 0x4a, 0x1c, 0x7d,
	0x90, 0xd8, 0x56, 0xe5, 0x48, 0xe0, 0x45, 0x92, 0xd5, 0x67, 0x34, 0x86, 0xf5, 0x52, 0xe3, 0xf5,
	0xff, 0x4b, 0xe3, 0xd2, 0xb7, 0xff, 0x49, 0x03, 0xe0, 0x65, 0x78, 0xb4, 0x07, 0x1b, 0x5e, 0xc8,
	0xb2, 0x6c, 0xde, 0xf0, 0x23, 0xf2, 0xb8, 0xfc, 0xad, 0xf9, 0xc5, 0xc2, 0xa5, 0x22, 0x48, 0x44,
	0x90, 0x5f, 0xd3, 0x90, 0x5f, 0xf2, 0xb0, 0xde, 0x80, 0xf6, 0x1c, 0xd5, 0x0b, 0x10, 0xbd, 0x03,
	0xbb, 0xa9, 0xe0, 0x3c, 0x4a, 0x8b, 0x1a, 0xa8, 0xc7, 0x52, 0x36, 0x0d, 0xc2, 0x20, 0xbf, 0xae,
	0x57, 0x79, 0xe7, 0x25, 0xa9, 0xdc, 0x70, 0xe8, 0x9b, 0xd0, 0x5b, 0x70, 0xba, 0x9c, 0x85, 0x31,
	0x17, 0x73, 0xbf, 0x6a, 0xd6, 0xdd, 0x97, 0xfc, 0xe9, 0x22, 0xdd, 0x8f, 0xa1, 0x77, 0x5f, 0x8b,
	0xe8, 0x10, 0xa4, 0x88, 0x5d, 0xd1, 0x29, 0x8b, 0xfd, 0x8f, 0x02, 0x3f, 0x7f, 0x4e, 0x67, 0x61,
	0x3d, 0xc6, 0x4e, 0xc4, 0xae, 0xc6, 0x73, 0xd8, 0x0d, 0x3f, 0x6d, 0xe9, 0xcf, 0xdb, 0x5b, 0xb2,
	0x54, 0xc3, 0xe1, 0x2f, 0x56, 0x61, 0x57, 0x61, 0xb1, 0xc7, 0xc3, 0x07, 0x7d, 0x03, 0x1c, 0xd8,
	0xf6, 0x4a, 0xaf, 0xb0, 0xf4, 0xa1, 0xf9, 0x75, 0xca, 0xcb, 0x0c, 0x9d, 0xe3, 0xa7, 0xa3, 0x3b,
	0xe3, 0xd5, 0x68, 0x65, 0xef, 0x5c, 0xa7, 0x9c, 0x48, 0xde, 0x2d, 0x64, 0xf8, 0xd3, 0x06, 0x48,
	0xb7, 0xcd, 0x50, 0x0f, 0x76, 0x26, 0x13, 0x4c, 0x5d, 0x4b, 0x95, 0x1d, 0x4c, 0x2d, 0x62, 0x2a,
	0x58, 0x75, 0x09, 0x96, 0x56, 0xd0, 0x1e, 0xec, 0xda, 0xef, 0xd9, 0xc6, 0xa7, 0xa9, 0x06, 0x7a,
	0x1d, 0xba, 0xb6, 0x3b, 0xb6, 0x15, 0xa2, 0x59, 0x8e, 0x66, 0x1a, 0xf4, 0x4c, 0x73, 0x4e, 0x54,
	0x22, 0x9f, 0xc9, 0xba, 0xb4, 0x5a, 0x44, 0xbc, 0xed, 0x42, 0xb5, 0xb3, 0x67, 0xd2, 0x1a, 0xda,
	0x87, 0x9e, 0x66, 0x68, 0x8e, 0x26, 0xeb, 0x54, 0x76, 0x1c, 0x59, 0x39, 0x59, 0x08, 0xba, 0x3e,
	0x94, 0x61, 0x67, 0xb9, 0xb5, 0x07, 0x7f, 0x7b, 0x86, 0x6f, 0x43, 0xc7, 0x9a, 0x89, 0x0b, 0xee,
	0xe2, 0xcf, 0x22, 0xf3, 0xf0, 0x5b, 0xd0, 0xae, 0xcd, 0x1f, 0x9c, 0xea, 0xe8, 0xf7, 0xeb, 0xd0,
	0xbc, 0x21, 0x50, 0x1b, 0x9a, 0xae, 0xa1, 0xe2, 0x67, 0x9a, 0x81, 0x55, 0x69, 0x05, 0xed, 0x82,
	0x34, 0x71, 0x75, 0x47, 0xa3, 0xc4, 0x74, 0x0d, 0x95, 0xca, 0xae, 0x73, 0x22, 0xfd, 0xf7, 0x31,
	0x6a, 0xc1, 0x63, 0xdb, 0x55, 0x14, 0x6c, 0xdb, 0xd2, 0xbf, 0xb6, 0xd0, 0x0e, 0x6c, 0xe9, 0xda,
	0x44, 0x73, 0xb0, 0x4a, 0xe7, 0xe8, 0xbf, 0xb7, 0x50, 0x17, 0x90, 0x62, 0x4e, 0x26, 0xb2, 0xa1,
	0x52, 0xd7, 0xb0, 0x5d, 0xcb, 0x24, 0x0e, 0x56, 0xa5, 0x3f, 0x74, 0xd1, 0x13, 0xd8, 0x76, 0x0d,
	0x79, 0xac, 0x63, 0xea, 0x98, 0x54, 0xc5, 0xba, 0x76, 0x8a, 0x89, 0xf4, 0xc7, 0x6e, 0x91, 0x8b,
	0x60, 0x59, 0x9f, 0x50, 0xc3, 0x74, 0xa8, 0x8d, 0xc9, 0x29, 0x56, 0xa5, 0x3f, 0x75, 0x51, 0x1b,
	0x36, 0x1c, 0xd3, 0xa4, 0x63, 0xd7, 0xfe, 0x9e, 0xf4, 0xe7, 0x2e, 0x42, 0xd0, 0xd6, 0x4d, 0xd3,
	0xa2, 0x2a, 0x76, 0xb0, 0x52, 0x44, 0xfc, 0x4b, 0x17, 0xf5, 0xe0, 0x0b, 0x04, 0xab, 0x1a, 0xc1,
	0x8a, 0x43, 0x35, 0x43, 0xd5, 0x14, 0xb9, 0x18, 0xa6, 0xf4, 0x49, 0x17, 0xed, 0x43, 0x57, 0xb6,
	0x2c, 0xbd, 0x46, 0xaa, 0x42, 0xea, 0x4a, 0xfe, 0x5a, 0x66, 0xd4, 0x8c, 0x53, 0x59, 0xd7, 0xd4,
	0x13, 0xaa, 0x12, 0x3a, 0xd6, 0x1c, 0x5b, 0xfa, 0xdb, 0x22, 0x4c, 0xe5, 0x53, 0xab, 0x82, 0xff,
	0xde, 0x45, 0xdb, 0xd0, 0x72, 0x8d, 0xf7, 0x0d, 0xf3, 0xcc, 0xa0, 0x16, 0xc6, 0x44, 0xfa, 0x47,
	0x15, 0xde, 0x75, 0x4e, 0xb0, 0xe1, 0xcc, 0x33, 0x10, 0xfc, 0xdd, 0xaa, 0xac, 0x9f, 0x1d, 0x14,
	0x0e, 0xa6, 0xeb, 0x50, 0xf3, 0x19, 0xb5, 0x2d, 0x59, 0xc1, 0xd2, 0xc7, 0x07, 0x45, 0xf5, 0x58,
	0xc7, 0x4a, 0x69, 0xaa, 0x9b, 0xb6, 0x23, 0xfd, 0xbc, 0x34, 0x73, 0x6d, 0x4c, 0x68, 0x1d, 0x5c,
	0xfa, 0xf1, 0x53, 0xf4, 0x06, 0xf4, 0xe6, 0xa9, 0xb0, 0x65, 0xd3, 0xc5, 0x15, 0x95, 0x7e, 0x79,
	0x54, 0x08, 0x4e, 0x64, 0xa7, 0xd4, 0x49, 0xd6, 0x75, 0xf3, 0x0c, 0xab, 0xd2, 0xaf, 0x8e, 0x4a,
	0x15, 0x4c, 0x79, 0xa2, 0x19, 0xef, 0x2d, 0x31, 0x3f, 0x79, 0x5a, 0x28, 0x8e, 0x3f, 0x70, 0x35,
	0x6b, 0x82, 0x0d, 0xe7, 0x26, 0xcd, 0xaf, 0x4b, 0x0f, 0xd7, 0x78, 0xbf, 0xc8, 0x52, 0xc8, 0x5d,
	0x39, 0xaa, 0x58, 0xfa, 0xcd, 0x11, 0xfa, 0x22, 0x1c, 0xdc, 0x6a, 0x4c, 0x95, 0x1d, 0x99, 0xba,
	0x86, 0x7c, 0x2a, 0x6b, 0x7a, 0x31, 0x3c, 0xe9, 0x3f, 0x83, 0xe3, 0x8f, 0x57, 0x61, 0xc3, 0x7e,
	0x97, 0x59, 0x22, 0xb9, 0xba, 0x46, 0x53, 0xd8, 0xbb, 0xf7, 0x19, 0x83, 0xde, 0x1c, 0xbd, 0xea,
	0x41, 0xd9, 0x1f, 0xbc, 0xea, 0x15, 0x34, 0x5c, 0x41, 0x32, 0x74, 0x96, 0xbf, 0xf6, 0xe8, 0xc9,
	0xe8, 0xce, 0x97, 0x49, 0x7f, 0xf7, 0xce, 0x7f, 0x0b, 0x55, 0x88, 0xe5, 0xcb, 0x89, 0x9e, 0xdc,
	0xfd, 0x21, 0xea, 0xef, 0x8e, 0xee, 0xba, 0xc5, 0xc3, 0x15, 0xf4, 0x16, 0x3c, 0xae, 0x6f, 0x1b,
	0xda, 0x1a, 0x2d, 0x5f, 0xd3, 0x7e, 0x67, 0xb4, 0x74, 0x11, 0x87, 0x2b, 0xe3, 0x8d, 0xef, 0xbf,
	0x56, 0x3e, 0xab, 0xb3, 0x69, 0xf5, 0xf7, 0x9d, 0xff, 0x05, 0x00, 0x00, 0xff, 0xff, 0x8e, 0xab,
	0x09, 0x7c, 0x6b, 0x0b, 0x00, 0x00,
}
