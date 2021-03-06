// Code generated by protoc-gen-go. DO NOT EDIT.
// source: clients.proto

package proto

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	math "math"
	_ "www.velocidex.com/golang/velociraptor/actions/proto"
	_ "www.velocidex.com/golang/velociraptor/proto"
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

type ApiClient_IPAddressClass int32

const (
	ApiClient_UNKNOWN  ApiClient_IPAddressClass = 0
	ApiClient_INTERNAL ApiClient_IPAddressClass = 1
	ApiClient_EXTERNAL ApiClient_IPAddressClass = 2
	ApiClient_VPN      ApiClient_IPAddressClass = 3
)

var ApiClient_IPAddressClass_name = map[int32]string{
	0: "UNKNOWN",
	1: "INTERNAL",
	2: "EXTERNAL",
	3: "VPN",
}

var ApiClient_IPAddressClass_value = map[string]int32{
	"UNKNOWN":  0,
	"INTERNAL": 1,
	"EXTERNAL": 2,
	"VPN":      3,
}

func (x ApiClient_IPAddressClass) String() string {
	return proto.EnumName(ApiClient_IPAddressClass_name, int32(x))
}

func (ApiClient_IPAddressClass) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_6c7b36ecb5ad4a28, []int{2, 0}
}

type SearchClientsRequest_QueryType int32

const (
	SearchClientsRequest_VALUE SearchClientsRequest_QueryType = 0
	SearchClientsRequest_KEY   SearchClientsRequest_QueryType = 1
)

var SearchClientsRequest_QueryType_name = map[int32]string{
	0: "VALUE",
	1: "KEY",
}

var SearchClientsRequest_QueryType_value = map[string]int32{
	"VALUE": 0,
	"KEY":   1,
}

func (x SearchClientsRequest_QueryType) String() string {
	return proto.EnumName(SearchClientsRequest_QueryType_name, int32(x))
}

func (SearchClientsRequest_QueryType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_6c7b36ecb5ad4a28, []int{3, 0}
}

// GRR uses an int for client_version which is difficult to use
// semantic versioning. We use a string instead which represents the
// commit number.
type AgentInformation struct {
	Version              string   `protobuf:"bytes,1,opt,name=version,proto3" json:"version,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	BuildTime            string   `protobuf:"bytes,3,opt,name=build_time,json=buildTime,proto3" json:"build_time,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AgentInformation) Reset()         { *m = AgentInformation{} }
func (m *AgentInformation) String() string { return proto.CompactTextString(m) }
func (*AgentInformation) ProtoMessage()    {}
func (*AgentInformation) Descriptor() ([]byte, []int) {
	return fileDescriptor_6c7b36ecb5ad4a28, []int{0}
}

func (m *AgentInformation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AgentInformation.Unmarshal(m, b)
}
func (m *AgentInformation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AgentInformation.Marshal(b, m, deterministic)
}
func (m *AgentInformation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AgentInformation.Merge(m, src)
}
func (m *AgentInformation) XXX_Size() int {
	return xxx_messageInfo_AgentInformation.Size(m)
}
func (m *AgentInformation) XXX_DiscardUnknown() {
	xxx_messageInfo_AgentInformation.DiscardUnknown(m)
}

var xxx_messageInfo_AgentInformation proto.InternalMessageInfo

func (m *AgentInformation) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *AgentInformation) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *AgentInformation) GetBuildTime() string {
	if m != nil {
		return m.BuildTime
	}
	return ""
}

// Message to carry uname information.
type Uname struct {
	System               string   `protobuf:"bytes,1,opt,name=system,proto3" json:"system,omitempty"`
	Node                 string   `protobuf:"bytes,2,opt,name=node,proto3" json:"node,omitempty"`
	Release              string   `protobuf:"bytes,3,opt,name=release,proto3" json:"release,omitempty"`
	Version              string   `protobuf:"bytes,4,opt,name=version,proto3" json:"version,omitempty"`
	Machine              string   `protobuf:"bytes,5,opt,name=machine,proto3" json:"machine,omitempty"`
	Kernel               string   `protobuf:"bytes,6,opt,name=kernel,proto3" json:"kernel,omitempty"`
	Fqdn                 string   `protobuf:"bytes,7,opt,name=fqdn,proto3" json:"fqdn,omitempty"`
	InstallDate          uint64   `protobuf:"varint,8,opt,name=install_date,json=installDate,proto3" json:"install_date,omitempty"`
	LibcVer              string   `protobuf:"bytes,9,opt,name=libc_ver,json=libcVer,proto3" json:"libc_ver,omitempty"`
	Architecture         string   `protobuf:"bytes,10,opt,name=architecture,proto3" json:"architecture,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Uname) Reset()         { *m = Uname{} }
func (m *Uname) String() string { return proto.CompactTextString(m) }
func (*Uname) ProtoMessage()    {}
func (*Uname) Descriptor() ([]byte, []int) {
	return fileDescriptor_6c7b36ecb5ad4a28, []int{1}
}

func (m *Uname) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Uname.Unmarshal(m, b)
}
func (m *Uname) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Uname.Marshal(b, m, deterministic)
}
func (m *Uname) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Uname.Merge(m, src)
}
func (m *Uname) XXX_Size() int {
	return xxx_messageInfo_Uname.Size(m)
}
func (m *Uname) XXX_DiscardUnknown() {
	xxx_messageInfo_Uname.DiscardUnknown(m)
}

var xxx_messageInfo_Uname proto.InternalMessageInfo

func (m *Uname) GetSystem() string {
	if m != nil {
		return m.System
	}
	return ""
}

func (m *Uname) GetNode() string {
	if m != nil {
		return m.Node
	}
	return ""
}

func (m *Uname) GetRelease() string {
	if m != nil {
		return m.Release
	}
	return ""
}

func (m *Uname) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *Uname) GetMachine() string {
	if m != nil {
		return m.Machine
	}
	return ""
}

func (m *Uname) GetKernel() string {
	if m != nil {
		return m.Kernel
	}
	return ""
}

func (m *Uname) GetFqdn() string {
	if m != nil {
		return m.Fqdn
	}
	return ""
}

func (m *Uname) GetInstallDate() uint64 {
	if m != nil {
		return m.InstallDate
	}
	return 0
}

func (m *Uname) GetLibcVer() string {
	if m != nil {
		return m.LibcVer
	}
	return ""
}

func (m *Uname) GetArchitecture() string {
	if m != nil {
		return m.Architecture
	}
	return ""
}

// Describe a client. We fill in some metadata about the client but
// this is by no means exhaustive.
type ApiClient struct {
	ClientId              string                   `protobuf:"bytes,1,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
	AgentInformation      *AgentInformation        `protobuf:"bytes,2,opt,name=agent_information,json=agentInformation,proto3" json:"agent_information,omitempty"`
	OsInfo                *Uname                   `protobuf:"bytes,3,opt,name=os_info,json=osInfo,proto3" json:"os_info,omitempty"`
	FirstSeenAt           uint64                   `protobuf:"varint,6,opt,name=first_seen_at,json=firstSeenAt,proto3" json:"first_seen_at,omitempty"`
	LastSeenAt            uint64                   `protobuf:"varint,7,opt,name=last_seen_at,json=lastSeenAt,proto3" json:"last_seen_at,omitempty"`
	LastBootedAt          uint64                   `protobuf:"varint,8,opt,name=last_booted_at,json=lastBootedAt,proto3" json:"last_booted_at,omitempty"`
	LastClock             uint64                   `protobuf:"varint,9,opt,name=last_clock,json=lastClock,proto3" json:"last_clock,omitempty"`
	LastCrashAt           uint64                   `protobuf:"varint,10,opt,name=last_crash_at,json=lastCrashAt,proto3" json:"last_crash_at,omitempty"`
	LastIp                string                   `protobuf:"bytes,16,opt,name=last_ip,json=lastIp,proto3" json:"last_ip,omitempty"`
	LastInterrogateFlowId string                   `protobuf:"bytes,19,opt,name=last_interrogate_flow_id,json=lastInterrogateFlowId,proto3" json:"last_interrogate_flow_id,omitempty"`
	LastIpClass           ApiClient_IPAddressClass `protobuf:"varint,17,opt,name=last_ip_class,json=lastIpClass,proto3,enum=proto.ApiClient_IPAddressClass" json:"last_ip_class,omitempty"`
	Labels                []string                 `protobuf:"bytes,18,rep,name=labels,proto3" json:"labels,omitempty"`
	XXX_NoUnkeyedLiteral  struct{}                 `json:"-"`
	XXX_unrecognized      []byte                   `json:"-"`
	XXX_sizecache         int32                    `json:"-"`
}

func (m *ApiClient) Reset()         { *m = ApiClient{} }
func (m *ApiClient) String() string { return proto.CompactTextString(m) }
func (*ApiClient) ProtoMessage()    {}
func (*ApiClient) Descriptor() ([]byte, []int) {
	return fileDescriptor_6c7b36ecb5ad4a28, []int{2}
}

func (m *ApiClient) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ApiClient.Unmarshal(m, b)
}
func (m *ApiClient) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ApiClient.Marshal(b, m, deterministic)
}
func (m *ApiClient) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ApiClient.Merge(m, src)
}
func (m *ApiClient) XXX_Size() int {
	return xxx_messageInfo_ApiClient.Size(m)
}
func (m *ApiClient) XXX_DiscardUnknown() {
	xxx_messageInfo_ApiClient.DiscardUnknown(m)
}

var xxx_messageInfo_ApiClient proto.InternalMessageInfo

func (m *ApiClient) GetClientId() string {
	if m != nil {
		return m.ClientId
	}
	return ""
}

func (m *ApiClient) GetAgentInformation() *AgentInformation {
	if m != nil {
		return m.AgentInformation
	}
	return nil
}

func (m *ApiClient) GetOsInfo() *Uname {
	if m != nil {
		return m.OsInfo
	}
	return nil
}

func (m *ApiClient) GetFirstSeenAt() uint64 {
	if m != nil {
		return m.FirstSeenAt
	}
	return 0
}

func (m *ApiClient) GetLastSeenAt() uint64 {
	if m != nil {
		return m.LastSeenAt
	}
	return 0
}

func (m *ApiClient) GetLastBootedAt() uint64 {
	if m != nil {
		return m.LastBootedAt
	}
	return 0
}

func (m *ApiClient) GetLastClock() uint64 {
	if m != nil {
		return m.LastClock
	}
	return 0
}

func (m *ApiClient) GetLastCrashAt() uint64 {
	if m != nil {
		return m.LastCrashAt
	}
	return 0
}

func (m *ApiClient) GetLastIp() string {
	if m != nil {
		return m.LastIp
	}
	return ""
}

func (m *ApiClient) GetLastInterrogateFlowId() string {
	if m != nil {
		return m.LastInterrogateFlowId
	}
	return ""
}

func (m *ApiClient) GetLastIpClass() ApiClient_IPAddressClass {
	if m != nil {
		return m.LastIpClass
	}
	return ApiClient_UNKNOWN
}

func (m *ApiClient) GetLabels() []string {
	if m != nil {
		return m.Labels
	}
	return nil
}

type SearchClientsRequest struct {
	Offset               uint64                         `protobuf:"varint,1,opt,name=offset,proto3" json:"offset,omitempty"`
	Limit                uint64                         `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
	Query                string                         `protobuf:"bytes,3,opt,name=query,proto3" json:"query,omitempty"`
	NameOnly             bool                           `protobuf:"varint,4,opt,name=name_only,json=nameOnly,proto3" json:"name_only,omitempty"`
	Type                 SearchClientsRequest_QueryType `protobuf:"varint,5,opt,name=type,proto3,enum=proto.SearchClientsRequest_QueryType" json:"type,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                       `json:"-"`
	XXX_unrecognized     []byte                         `json:"-"`
	XXX_sizecache        int32                          `json:"-"`
}

func (m *SearchClientsRequest) Reset()         { *m = SearchClientsRequest{} }
func (m *SearchClientsRequest) String() string { return proto.CompactTextString(m) }
func (*SearchClientsRequest) ProtoMessage()    {}
func (*SearchClientsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_6c7b36ecb5ad4a28, []int{3}
}

func (m *SearchClientsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SearchClientsRequest.Unmarshal(m, b)
}
func (m *SearchClientsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SearchClientsRequest.Marshal(b, m, deterministic)
}
func (m *SearchClientsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SearchClientsRequest.Merge(m, src)
}
func (m *SearchClientsRequest) XXX_Size() int {
	return xxx_messageInfo_SearchClientsRequest.Size(m)
}
func (m *SearchClientsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SearchClientsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SearchClientsRequest proto.InternalMessageInfo

func (m *SearchClientsRequest) GetOffset() uint64 {
	if m != nil {
		return m.Offset
	}
	return 0
}

func (m *SearchClientsRequest) GetLimit() uint64 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *SearchClientsRequest) GetQuery() string {
	if m != nil {
		return m.Query
	}
	return ""
}

func (m *SearchClientsRequest) GetNameOnly() bool {
	if m != nil {
		return m.NameOnly
	}
	return false
}

func (m *SearchClientsRequest) GetType() SearchClientsRequest_QueryType {
	if m != nil {
		return m.Type
	}
	return SearchClientsRequest_VALUE
}

type SearchClientsResponse struct {
	Items                []*ApiClient `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
	Names                []string     `protobuf:"bytes,2,rep,name=names,proto3" json:"names,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *SearchClientsResponse) Reset()         { *m = SearchClientsResponse{} }
func (m *SearchClientsResponse) String() string { return proto.CompactTextString(m) }
func (*SearchClientsResponse) ProtoMessage()    {}
func (*SearchClientsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_6c7b36ecb5ad4a28, []int{4}
}

func (m *SearchClientsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SearchClientsResponse.Unmarshal(m, b)
}
func (m *SearchClientsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SearchClientsResponse.Marshal(b, m, deterministic)
}
func (m *SearchClientsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SearchClientsResponse.Merge(m, src)
}
func (m *SearchClientsResponse) XXX_Size() int {
	return xxx_messageInfo_SearchClientsResponse.Size(m)
}
func (m *SearchClientsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SearchClientsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SearchClientsResponse proto.InternalMessageInfo

func (m *SearchClientsResponse) GetItems() []*ApiClient {
	if m != nil {
		return m.Items
	}
	return nil
}

func (m *SearchClientsResponse) GetNames() []string {
	if m != nil {
		return m.Names
	}
	return nil
}

type GetClientRequest struct {
	ClientId             string   `protobuf:"bytes,1,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
	Lightweight          bool     `protobuf:"varint,2,opt,name=lightweight,proto3" json:"lightweight,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetClientRequest) Reset()         { *m = GetClientRequest{} }
func (m *GetClientRequest) String() string { return proto.CompactTextString(m) }
func (*GetClientRequest) ProtoMessage()    {}
func (*GetClientRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_6c7b36ecb5ad4a28, []int{5}
}

func (m *GetClientRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetClientRequest.Unmarshal(m, b)
}
func (m *GetClientRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetClientRequest.Marshal(b, m, deterministic)
}
func (m *GetClientRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetClientRequest.Merge(m, src)
}
func (m *GetClientRequest) XXX_Size() int {
	return xxx_messageInfo_GetClientRequest.Size(m)
}
func (m *GetClientRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetClientRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetClientRequest proto.InternalMessageInfo

func (m *GetClientRequest) GetClientId() string {
	if m != nil {
		return m.ClientId
	}
	return ""
}

func (m *GetClientRequest) GetLightweight() bool {
	if m != nil {
		return m.Lightweight
	}
	return false
}

type LabelClientsRequest struct {
	ClientIds            []string `protobuf:"bytes,1,rep,name=client_ids,json=clientIds,proto3" json:"client_ids,omitempty"`
	Labels               []string `protobuf:"bytes,2,rep,name=labels,proto3" json:"labels,omitempty"`
	Operation            string   `protobuf:"bytes,3,opt,name=operation,proto3" json:"operation,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LabelClientsRequest) Reset()         { *m = LabelClientsRequest{} }
func (m *LabelClientsRequest) String() string { return proto.CompactTextString(m) }
func (*LabelClientsRequest) ProtoMessage()    {}
func (*LabelClientsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_6c7b36ecb5ad4a28, []int{6}
}

func (m *LabelClientsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LabelClientsRequest.Unmarshal(m, b)
}
func (m *LabelClientsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LabelClientsRequest.Marshal(b, m, deterministic)
}
func (m *LabelClientsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LabelClientsRequest.Merge(m, src)
}
func (m *LabelClientsRequest) XXX_Size() int {
	return xxx_messageInfo_LabelClientsRequest.Size(m)
}
func (m *LabelClientsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_LabelClientsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_LabelClientsRequest proto.InternalMessageInfo

func (m *LabelClientsRequest) GetClientIds() []string {
	if m != nil {
		return m.ClientIds
	}
	return nil
}

func (m *LabelClientsRequest) GetLabels() []string {
	if m != nil {
		return m.Labels
	}
	return nil
}

func (m *LabelClientsRequest) GetOperation() string {
	if m != nil {
		return m.Operation
	}
	return ""
}

func init() {
	proto.RegisterEnum("proto.ApiClient_IPAddressClass", ApiClient_IPAddressClass_name, ApiClient_IPAddressClass_value)
	proto.RegisterEnum("proto.SearchClientsRequest_QueryType", SearchClientsRequest_QueryType_name, SearchClientsRequest_QueryType_value)
	proto.RegisterType((*AgentInformation)(nil), "proto.AgentInformation")
	proto.RegisterType((*Uname)(nil), "proto.Uname")
	proto.RegisterType((*ApiClient)(nil), "proto.ApiClient")
	proto.RegisterType((*SearchClientsRequest)(nil), "proto.SearchClientsRequest")
	proto.RegisterType((*SearchClientsResponse)(nil), "proto.SearchClientsResponse")
	proto.RegisterType((*GetClientRequest)(nil), "proto.GetClientRequest")
	proto.RegisterType((*LabelClientsRequest)(nil), "proto.LabelClientsRequest")
}

func init() { proto.RegisterFile("clients.proto", fileDescriptor_6c7b36ecb5ad4a28) }

var fileDescriptor_6c7b36ecb5ad4a28 = []byte{
	// 1461 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x96, 0xdd, 0x52, 0x23, 0xb9,
	0x15, 0xc7, 0xd7, 0x60, 0x6c, 0x2c, 0x03, 0xf1, 0x6a, 0x77, 0x76, 0xbb, 0x66, 0x99, 0x8c, 0xe2,
	0x0a, 0x59, 0x93, 0x61, 0x1a, 0x6c, 0x08, 0xf3, 0x51, 0x49, 0x6d, 0xd9, 0xc0, 0x4e, 0x39, 0xc3,
	0x00, 0x69, 0x98, 0x8f, 0xe4, 0x22, 0x2e, 0xb9, 0xfb, 0xd8, 0x56, 0xa6, 0x2d, 0x19, 0x49, 0xc6,
	0xe3, 0xaa, 0xb9, 0xc8, 0x4b, 0xa4, 0x2a, 0x77, 0xb9, 0xcb, 0x93, 0xcc, 0x1b, 0xe4, 0x0d, 0x92,
	0x97, 0xc8, 0x45, 0x52, 0x95, 0xd2, 0x69, 0x19, 0x30, 0x49, 0xaa, 0x52, 0xb9, 0x81, 0x96, 0xce,
	0x39, 0x3f, 0xfd, 0x75, 0x74, 0x74, 0x64, 0xb2, 0x1a, 0xa7, 0x02, 0xa4, 0x35, 0xe1, 0x48, 0x2b,
	0xab, 0xe8, 0x12, 0xfe, 0xbb, 0xff, 0x7c, 0x32, 0x99, 0x84, 0x57, 0x90, 0xaa, 0x58, 0x24, 0xf0,
	0x21, 0x8c, 0xd5, 0x70, 0xbb, 0xaf, 0x52, 0x2e, 0xfb, 0xdb, 0xd9, 0xa4, 0xe6, 0x23, 0xab, 0xf4,
	0x36, 0x3a, 0x6f, 0x1b, 0x18, 0x72, 0x69, 0x45, 0x9c, 0x21, 0xee, 0x37, 0xff, 0xb7, 0x58, 0x1e,
	0x5b, 0xa1, 0xa4, 0xf1, 0x0c, 0x3f, 0xf2, 0x88, 0x5f, 0xfc, 0x3f, 0x88, 0xab, 0xcb, 0xd4, 0x87,
	0xaf, 0xf7, 0x95, 0xea, 0xa7, 0xb0, 0xcd, 0x47, 0x62, 0x9b, 0x4b, 0xa9, 0x2c, 0xbf, 0x05, 0xaf,
	0x76, 0x48, 0xa5, 0xd9, 0x07, 0x69, 0xdb, 0xb2, 0xa7, 0xf4, 0x10, 0x4d, 0x34, 0x20, 0xc5, 0x2b,
	0xd0, 0x46, 0x28, 0x19, 0xe4, 0x58, 0xae, 0x56, 0x8a, 0x66, 0x43, 0x4a, 0x49, 0x5e, 0xf2, 0x21,
	0x04, 0x0b, 0x38, 0x8d, 0xdf, 0xf4, 0x01, 0x21, 0xdd, 0xb1, 0x48, 0x93, 0x8e, 0x15, 0x43, 0x08,
	0x16, 0xd1, 0x52, 0xc2, 0x99, 0x0b, 0x31, 0x84, 0xea, 0x3f, 0x0b, 0x64, 0xe9, 0x35, 0x3a, 0xbe,
	0x24, 0x05, 0x33, 0x35, 0x16, 0x86, 0x19, 0xb5, 0xb5, 0xfb, 0xd7, 0x7f, 0xfc, 0xed, 0x53, 0xee,
	0x31, 0x7d, 0x74, 0x31, 0x00, 0x96, 0x59, 0xd8, 0x28, 0xe5, 0xd6, 0x09, 0x61, 0xb5, 0xb7, 0x42,
	0x26, 0x6a, 0x62, 0x3e, 0x1e, 0x72, 0x3d, 0x11, 0xf2, 0xe3, 0xb1, 0x90, 0xe3, 0x0f, 0x9b, 0x61,
	0xe4, 0x11, 0xf4, 0x29, 0xc9, 0x4b, 0x95, 0x78, 0x25, 0xad, 0x1f, 0x23, 0xea, 0x87, 0x74, 0xdd,
	0xa1, 0x06, 0xca, 0x58, 0xb7, 0x20, 0x53, 0x3d, 0x66, 0x07, 0xc2, 0x78, 0x76, 0x18, 0x61, 0x04,
	0x3d, 0x23, 0x45, 0x0d, 0x29, 0x70, 0xe3, 0xc5, 0xb6, 0xf6, 0x31, 0x78, 0x87, 0x86, 0x2e, 0xf8,
	0xf4, 0x9c, 0x79, 0x2b, 0x13, 0x09, 0x48, 0x2b, 0x7a, 0x02, 0x34, 0x83, 0xb0, 0x1f, 0xb2, 0x27,
	0x5b, 0xec, 0xf4, 0xfc, 0xdd, 0x16, 0x4b, 0xa0, 0x2b, 0xb8, 0x0c, 0xa3, 0x19, 0x86, 0x5e, 0xdc,
	0xe4, 0x2b, 0x8f, 0xc4, 0xe7, 0x48, 0xdc, 0xa3, 0x0d, 0x4f, 0xf4, 0x56, 0xd6, 0x3e, 0xcc, 0x48,
	0xfb, 0x61, 0x3d, 0x7c, 0xb2, 0xbf, 0x53, 0x3f, 0x3f, 0xab, 0x6f, 0xb1, 0xfa, 0x4e, 0xf8, 0x2c,
	0x6c, 0x6c, 0xb1, 0xfa, 0x5e, 0xb8, 0xb3, 0x17, 0xde, 0xe4, 0xfa, 0x15, 0x29, 0x0e, 0x79, 0x3c,
	0x10, 0x12, 0x82, 0xa5, 0xff, 0x9a, 0x2f, 0xae, 0xe3, 0x81, 0xb0, 0x10, 0xdb, 0xb1, 0x86, 0x8c,
	0xdd, 0x7c, 0x75, 0xb8, 0xbf, 0xb7, 0xc5, 0x3e, 0x3c, 0xdd, 0xef, 0xec, 0x3b, 0x9c, 0x67, 0xd0,
	0xdf, 0x90, 0xc2, 0x7b, 0xd0, 0x12, 0xd2, 0xa0, 0x80, 0xb4, 0x16, 0xd2, 0x7e, 0x4e, 0x9f, 0x3b,
	0x5a, 0x66, 0xb9, 0xd6, 0x69, 0xac, 0x16, 0xb2, 0x3f, 0xaf, 0x75, 0x8b, 0xd5, 0x77, 0xc3, 0x7a,
	0xb8, 0xb3, 0xc5, 0x76, 0xc3, 0xfa, 0xcf, 0x1e, 0xeb, 0xb8, 0x11, 0x46, 0x9e, 0x48, 0x8f, 0x48,
	0xbe, 0x77, 0x99, 0xc8, 0xa0, 0x88, 0xe4, 0x3a, 0x92, 0x1f, 0xd1, 0xcd, 0x1b, 0x9d, 0xdf, 0x1a,
	0xd6, 0x1b, 0xa7, 0xe9, 0x94, 0x5d, 0x8e, 0x79, 0xea, 0x52, 0x9a, 0xb0, 0x44, 0x0d, 0xb9, 0x90,
	0xcc, 0x1d, 0x54, 0x18, 0x61, 0x38, 0x8d, 0xc8, 0x8a, 0x90, 0xc6, 0xf2, 0x34, 0xed, 0x24, 0xdc,
	0x42, 0xb0, 0xcc, 0x72, 0xb5, 0x7c, 0x6b, 0x1b, 0x71, 0x9b, 0xa4, 0x1c, 0x1d, 0x7e, 0x7f, 0xc8,
	0x2d, 0xb8, 0x32, 0xa3, 0xf7, 0xdf, 0x0e, 0x40, 0xce, 0x92, 0x30, 0xe1, 0x86, 0xf9, 0x40, 0x48,
	0xc2, 0xa8, 0xec, 0xbf, 0x9d, 0x33, 0x7d, 0x4a, 0x96, 0x53, 0xd1, 0x8d, 0x3b, 0x57, 0xa0, 0x83,
	0x12, 0xca, 0x7b, 0x80, 0xbc, 0xaf, 0xe9, 0x3d, 0x27, 0xef, 0x80, 0xa5, 0xa2, 0xab, 0xb9, 0x9e,
	0xce, 0xf6, 0x1e, 0x15, 0x9d, 0xfb, 0x1b, 0xd0, 0xf4, 0x53, 0x8e, 0xac, 0xdc, 0x4e, 0x6f, 0x40,
	0x30, 0xfc, 0x4f, 0x39, 0x8c, 0xff, 0x63, 0x8e, 0xfe, 0x21, 0xe7, 0x08, 0x73, 0x27, 0x30, 0xab,
	0xb8, 0xae, 0x90, 0x5c, 0x4f, 0x43, 0x56, 0x3b, 0x51, 0x16, 0xb2, 0xa9, 0x98, 0x4b, 0xd6, 0x05,
	0x96, 0x88, 0x5e, 0x0f, 0x34, 0x48, 0xcb, 0x7a, 0x5a, 0x0d, 0x99, 0x1d, 0x00, 0xf3, 0x27, 0x34,
	0x4f, 0x12, 0x12, 0x6d, 0xb1, 0x2b, 0x44, 0xd5, 0x63, 0x9c, 0xed, 0x36, 0x58, 0x57, 0x58, 0x4f,
	0x66, 0x7a, 0x2c, 0xa5, 0x3b, 0x22, 0x25, 0x19, 0x67, 0xfb, 0x7b, 0x68, 0xca, 0xb2, 0xb1, 0x19,
	0xcd, 0xa9, 0xae, 0xfe, 0xbd, 0x40, 0x4a, 0xcd, 0x91, 0x38, 0xc0, 0xc6, 0x46, 0xbf, 0x23, 0xa5,
	0xac, 0xc5, 0x75, 0x44, 0xe2, 0xaf, 0x61, 0x15, 0xf7, 0xb3, 0x4e, 0xca, 0xd7, 0x5e, 0xed, 0x84,
	0xae, 0xba, 0xad, 0x65, 0x9e, 0x4c, 0x24, 0xd1, 0x72, 0x3c, 0x33, 0x1c, 0x92, 0xcf, 0x79, 0x1f,
	0xe3, 0x6f, 0x1a, 0x06, 0x5e, 0xc2, 0x72, 0xe3, 0xeb, 0xac, 0xa5, 0x84, 0x77, 0xfb, 0x49, 0x54,
	0xe1, 0x77, 0x3b, 0xcc, 0x06, 0x29, 0x2a, 0x83, 0x08, 0xbc, 0x83, 0xe5, 0xc6, 0x8a, 0x8f, 0xc5,
	0x4e, 0x11, 0x15, 0x94, 0x71, 0xde, 0xf4, 0x94, 0xac, 0xf6, 0x84, 0x36, 0xb6, 0x63, 0x00, 0x64,
	0x87, 0x5b, 0x2c, 0xdd, 0x7c, 0xeb, 0x11, 0x2a, 0xde, 0x98, 0xaf, 0x88, 0xaf, 0xb0, 0x22, 0x7a,
	0xa9, 0x9a, 0x60, 0x3d, 0xc4, 0x1a, 0xb8, 0xc5, 0x6a, 0x40, 0xc2, 0x39, 0x80, 0x6c, 0x5a, 0xfa,
	0x8e, 0xac, 0xa4, 0xfc, 0x16, 0xaf, 0x88, 0x3c, 0xdf, 0x00, 0xe6, 0x79, 0x3f, 0x3a, 0xe6, 0xc6,
	0x32, 0xf7, 0xc9, 0x26, 0x8e, 0xec, 0x93, 0x11, 0x0f, 0x20, 0x7e, 0x0f, 0x09, 0x13, 0x32, 0x8c,
	0x88, 0x63, 0x79, 0xf2, 0x2f, 0xc9, 0x1a, 0x92, 0xbb, 0x4a, 0x59, 0x48, 0x1c, 0x3b, 0xab, 0x5e,
	0xdf, 0x99, 0xe6, 0xd9, 0x3f, 0x40, 0xb6, 0x73, 0xc5, 0x05, 0xc2, 0x08, 0x55, 0xb5, 0x30, 0xb4,
	0x69, 0xe9, 0x6f, 0x09, 0x92, 0x3b, 0x71, 0xaa, 0xe2, 0xf7, 0x58, 0xb5, 0xf9, 0xd6, 0x77, 0xc8,
	0x79, 0x36, 0xcf, 0xf9, 0xe9, 0x81, 0x17, 0xe5, 0x1c, 0x0d, 0xbb, 0xe2, 0xe9, 0x18, 0x58, 0x32,
	0xc6, 0xcb, 0x9b, 0x72, 0x0b, 0xc6, 0xeb, 0x75, 0x62, 0x4b, 0x0e, 0x79, 0xe0, 0x1c, 0x69, 0x9b,
	0xac, 0x66, 0x7c, 0xcd, 0xcd, 0xc0, 0x49, 0x25, 0xb8, 0xc4, 0x06, 0x2e, 0xf1, 0x70, 0x7e, 0x89,
	0x0a, 0x4a, 0x45, 0x4f, 0xaf, 0xb5, 0x8c, 0x20, 0x37, 0xd1, 0xb4, 0xb4, 0x49, 0x8a, 0x88, 0x12,
	0xa3, 0xa0, 0x82, 0xd5, 0x54, 0x43, 0x48, 0x95, 0x32, 0x57, 0x40, 0xce, 0xc4, 0x5c, 0xae, 0x99,
	0x86, 0xa1, 0xbb, 0x09, 0xcd, 0xb3, 0x36, 0xe3, 0x49, 0xa2, 0xc1, 0x98, 0xa8, 0xe0, 0xac, 0xed,
	0x11, 0x7d, 0x42, 0x82, 0x0c, 0x21, 0x2d, 0x68, 0xad, 0xfa, 0xdc, 0x42, 0xc7, 0x1d, 0xa3, 0xab,
	0xd0, 0x2f, 0xf0, 0x35, 0xb9, 0x87, 0x9e, 0x37, 0xe6, 0xef, 0x53, 0x35, 0x69, 0x27, 0xf4, 0xc0,
	0x6f, 0x43, 0x8c, 0x3a, 0x71, 0xca, 0x8d, 0x09, 0x3e, 0x67, 0xb9, 0xda, 0x5a, 0xe3, 0xe1, 0xac,
	0x0c, 0x67, 0xe5, 0x1c, 0xb6, 0xcf, 0x9a, 0xd9, 0x9a, 0x07, 0xce, 0x2d, 0xdb, 0x40, 0x7b, 0x84,
	0x03, 0xfa, 0x15, 0x29, 0xa4, 0xbc, 0x0b, 0xa9, 0x09, 0x28, 0x5b, 0xac, 0x95, 0x22, 0x3f, 0xaa,
	0xb6, 0xc8, 0xda, 0x7c, 0x18, 0x2d, 0x93, 0xe2, 0xeb, 0x93, 0x97, 0x27, 0xa7, 0x6f, 0x4f, 0x2a,
	0x9f, 0xd1, 0x15, 0xb2, 0xdc, 0x3e, 0xb9, 0x38, 0x8a, 0x4e, 0x9a, 0xc7, 0x95, 0x9c, 0x1b, 0x1d,
	0xbd, 0xf3, 0xa3, 0x05, 0x5a, 0x24, 0x8b, 0x6f, 0xce, 0x4e, 0x2a, 0x8b, 0xd5, 0xbf, 0xe4, 0xc8,
	0x97, 0xe7, 0xe0, 0x6e, 0x63, 0x26, 0xc4, 0x44, 0x70, 0x39, 0x06, 0x63, 0xdd, 0xa2, 0xaa, 0xd7,
	0x33, 0x60, 0xf1, 0x0a, 0xe6, 0x23, 0x3f, 0xa2, 0x5f, 0x92, 0xa5, 0x54, 0x0c, 0x85, 0xc5, 0x0b,
	0x95, 0x8f, 0xb2, 0x81, 0x9b, 0xbd, 0x1c, 0x83, 0x9e, 0xfa, 0xb7, 0x35, 0x1b, 0xd0, 0x6f, 0x48,
	0xc9, 0xdd, 0x95, 0x8e, 0x92, 0xe9, 0x14, 0x9f, 0x9d, 0xe5, 0x68, 0xd9, 0x4d, 0x9c, 0xca, 0x74,
	0x4a, 0x9f, 0x91, 0xbc, 0x9d, 0x8e, 0xb2, 0x87, 0x63, 0xad, 0xb1, 0xe1, 0x33, 0xf2, 0x9f, 0xb4,
	0x84, 0xbf, 0x72, 0xb4, 0x8b, 0xe9, 0x08, 0x22, 0x0c, 0xa9, 0x3e, 0x24, 0xa5, 0xeb, 0x29, 0x5a,
	0x22, 0x4b, 0x6f, 0x9a, 0xc7, 0xaf, 0x8f, 0x2a, 0x9f, 0xb9, 0x5d, 0xbd, 0x3c, 0xfa, 0x75, 0x25,
	0x57, 0xfd, 0x73, 0x8e, 0xdc, 0xbb, 0x43, 0x32, 0x23, 0x25, 0x0d, 0xd0, 0x9f, 0x90, 0x25, 0x61,
	0x61, 0x68, 0x82, 0x1c, 0x5b, 0xac, 0x95, 0x1b, 0x95, 0xbb, 0x07, 0x11, 0x65, 0x66, 0x0a, 0x64,
	0xc9, 0x29, 0x35, 0xc1, 0x82, 0x4b, 0x79, 0xeb, 0x14, 0x4b, 0xa6, 0x4d, 0x5f, 0xb4, 0x7b, 0xec,
	0x7a, 0x4b, 0xcc, 0xbd, 0xdb, 0x23, 0x88, 0xb3, 0xc7, 0xc2, 0x77, 0x43, 0x9d, 0x69, 0x66, 0x13,
	0x60, 0xe8, 0xa3, 0xc1, 0x8e, 0x75, 0x66, 0x42, 0x20, 0x1b, 0x80, 0x86, 0x30, 0xca, 0xe8, 0xd5,
	0xdf, 0xe7, 0x48, 0xe5, 0x05, 0x58, 0xbf, 0xb6, 0x4f, 0xfd, 0x37, 0xff, 0xd6, 0x00, 0x6f, 0x35,
	0xb7, 0x53, 0x52, 0x4e, 0x45, 0x7f, 0x60, 0x27, 0xe0, 0xfe, 0xe2, 0x29, 0x2c, 0xb7, 0x1e, 0xa3,
	0xbc, 0x6f, 0xe9, 0x46, 0xbb, 0xc7, 0x0c, 0xd8, 0x6c, 0xe5, 0x58, 0xa5, 0x29, 0xc4, 0x96, 0x19,
	0x35, 0x74, 0xbd, 0xfa, 0xba, 0xb3, 0xb9, 0xeb, 0x71, 0x43, 0xa8, 0xfe, 0x8e, 0x7c, 0x71, 0xec,
	0xea, 0xe9, 0xce, 0xf9, 0x3f, 0x20, 0xe4, 0x5a, 0x44, 0x96, 0xad, 0x52, 0x54, 0x9a, 0xa9, 0xb8,
	0x5d, 0x93, 0x0b, 0xb7, 0x6b, 0x92, 0xae, 0x93, 0x92, 0x1a, 0x81, 0xce, 0x7a, 0xae, 0xff, 0xa1,
	0x75, 0x3d, 0xd1, 0x2d, 0x60, 0xb6, 0x77, 0xff, 0x15, 0x00, 0x00, 0xff, 0xff, 0x0c, 0x85, 0xf0,
	0x53, 0xc4, 0x0a, 0x00, 0x00,
}
