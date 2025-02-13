// Code generated by protoc-gen-go. DO NOT EDIT.
// source: XXXX.proto

package XXXX

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

type Request_Type int32

const (
	Request_TYPE_NONE               Request_Type = 0
	Request_TYPE_HEART_BEAT         Request_Type = 1
	Request_TYPE_CLIENT_REQUEST     Request_Type = 2
	Request_TYPE_PRE_PREPARE        Request_Type = 3
	Request_TYPE_PREPARE            Request_Type = 4
	Request_TYPE_COMMIT             Request_Type = 5
	Request_TYPE_CLIENT_CERT        Request_Type = 6
	Request_TYPE_RESPONSE           Request_Type = 7
	Request_TYPE_RECOVERY_DATA      Request_Type = 8
	Request_TYPE_RECOVERY_DATA_RESP Request_Type = 9
	Request_TYPE_CHECKPOINT         Request_Type = 10
	Request_TYPE_QUERY              Request_Type = 11
	Request_TYPE_REPLICA_STATE      Request_Type = 12
	Request_TYPE_NEW_TXNS           Request_Type = 14
	//with batch transactions.
	Request_TYPE_GEO_REQUEST       Request_Type = 15
	Request_TYPE_VIEWCHANGE        Request_Type = 16
	Request_TYPE_NEWVIEW           Request_Type = 17
	Request_TYPE_CUSTOM_QUERY      Request_Type = 18
	Request_TYPE_GEO_MINING_RESULT Request_Type = 19
	Request_NUM_OF_TYPE            Request_Type = 20
)

var Request_Type_name = map[int32]string{
	0:  "TYPE_NONE",
	1:  "TYPE_HEART_BEAT",
	2:  "TYPE_CLIENT_REQUEST",
	3:  "TYPE_PRE_PREPARE",
	4:  "TYPE_PREPARE",
	5:  "TYPE_COMMIT",
	6:  "TYPE_CLIENT_CERT",
	7:  "TYPE_RESPONSE",
	8:  "TYPE_RECOVERY_DATA",
	9:  "TYPE_RECOVERY_DATA_RESP",
	10: "TYPE_CHECKPOINT",
	11: "TYPE_QUERY",
	12: "TYPE_REPLICA_STATE",
	14: "TYPE_NEW_TXNS",
	15: "TYPE_GEO_REQUEST",
	16: "TYPE_VIEWCHANGE",
	17: "TYPE_NEWVIEW",
	18: "TYPE_CUSTOM_QUERY",
	19: "TYPE_GEO_MINING_RESULT",
	20: "NUM_OF_TYPE",
}

var Request_Type_value = map[string]int32{
	"TYPE_NONE":               0,
	"TYPE_HEART_BEAT":         1,
	"TYPE_CLIENT_REQUEST":     2,
	"TYPE_PRE_PREPARE":        3,
	"TYPE_PREPARE":            4,
	"TYPE_COMMIT":             5,
	"TYPE_CLIENT_CERT":        6,
	"TYPE_RESPONSE":           7,
	"TYPE_RECOVERY_DATA":      8,
	"TYPE_RECOVERY_DATA_RESP": 9,
	"TYPE_CHECKPOINT":         10,
	"TYPE_QUERY":              11,
	"TYPE_REPLICA_STATE":      12,
	"TYPE_NEW_TXNS":           14,
	"TYPE_GEO_REQUEST":        15,
	"TYPE_VIEWCHANGE":         16,
	"TYPE_NEWVIEW":            17,
	"TYPE_CUSTOM_QUERY":       18,
	"TYPE_GEO_MINING_RESULT":  19,
	"NUM_OF_TYPE":             20,
}

func (x Request_Type) String() string {
	return proto.EnumName(Request_Type_name, int32(x))
}

func (Request_Type) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_0a81bd5ef5ef03a2, []int{1, 0}
}

// Network message used to deliver Requests between replicas and client.
type XDBMessage struct {
	Data                 []byte   `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *XDBMessage) Reset()         { *m = XDBMessage{} }
func (m *XDBMessage) String() string { return proto.CompactTextString(m) }
func (*XDBMessage) ProtoMessage()    {}
func (*XDBMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_0a81bd5ef5ef03a2, []int{0}
}

func (m *XDBMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_XDBMessage.Unmarshal(m, b)
}
func (m *XDBMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_XDBMessage.Marshal(b, m, deterministic)
}
func (m *XDBMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_XDBMessage.Merge(m, src)
}
func (m *XDBMessage) XXX_Size() int {
	return xxx_messageInfo_XDBMessage.Size(m)
}
func (m *XDBMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_XDBMessage.DiscardUnknown(m)
}

var xxx_messageInfo_XDBMessage proto.InternalMessageInfo

func (m *XDBMessage) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

// The request message containing requested numbers
type Request struct {
	Type            int32  `protobuf:"varint,1,opt,name=type,proto3" json:"type,omitempty"`
	Data            []byte `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	CurrentView     uint64 `protobuf:"varint,4,opt,name=current_view,json=currentView,proto3" json:"current_view,omitempty"`
	Seq             uint64 `protobuf:"varint,5,opt,name=seq,proto3" json:"seq,omitempty"`
	Hash            []byte `protobuf:"bytes,6,opt,name=hash,proto3" json:"hash,omitempty"`
	SenderId        int32  `protobuf:"varint,7,opt,name=sender_id,json=senderId,proto3" json:"sender_id,omitempty"`
	ProxyId         int64  `protobuf:"varint,8,opt,name=proxy_id,json=proxyId,proto3" json:"proxy_id,omitempty"`
	IsSystemRequest bool   `protobuf:"varint,9,opt,name=is_system_request,json=isSystemRequest,proto3" json:"is_system_request,omitempty"`
	// request, like CMD:ADDREPLICA.
	CurrentExecutedSeq   uint64   `protobuf:"varint,10,opt,name=current_executed_seq,json=currentExecutedSeq,proto3" json:"current_executed_seq,omitempty"`
	NeedResponse         bool     `protobuf:"varint,11,opt,name=need_response,json=needResponse,proto3" json:"need_response,omitempty"`
	Ret                  int32    `protobuf:"varint,12,opt,name=ret,proto3" json:"ret,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_0a81bd5ef5ef03a2, []int{1}
}

func (m *Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Request.Unmarshal(m, b)
}
func (m *Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Request.Marshal(b, m, deterministic)
}
func (m *Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request.Merge(m, src)
}
func (m *Request) XXX_Size() int {
	return xxx_messageInfo_Request.Size(m)
}
func (m *Request) XXX_DiscardUnknown() {
	xxx_messageInfo_Request.DiscardUnknown(m)
}

var xxx_messageInfo_Request proto.InternalMessageInfo

func (m *Request) GetType() int32 {
	if m != nil {
		return m.Type
	}
	return 0
}

func (m *Request) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *Request) GetCurrentView() uint64 {
	if m != nil {
		return m.CurrentView
	}
	return 0
}

func (m *Request) GetSeq() uint64 {
	if m != nil {
		return m.Seq
	}
	return 0
}

func (m *Request) GetHash() []byte {
	if m != nil {
		return m.Hash
	}
	return nil
}

func (m *Request) GetSenderId() int32 {
	if m != nil {
		return m.SenderId
	}
	return 0
}

func (m *Request) GetProxyId() int64 {
	if m != nil {
		return m.ProxyId
	}
	return 0
}

func (m *Request) GetIsSystemRequest() bool {
	if m != nil {
		return m.IsSystemRequest
	}
	return false
}

func (m *Request) GetCurrentExecutedSeq() uint64 {
	if m != nil {
		return m.CurrentExecutedSeq
	}
	return 0
}

func (m *Request) GetNeedResponse() bool {
	if m != nil {
		return m.NeedResponse
	}
	return false
}

func (m *Request) GetRet() int32 {
	if m != nil {
		return m.Ret
	}
	return 0
}

type BatchClientRequest struct {
	ClientRequests       []*BatchClientRequest_ClientRequest `protobuf:"bytes,1,rep,name=client_requests,json=clientRequests,proto3" json:"client_requests,omitempty"`
	Createtime           uint64                              `protobuf:"varint,2,opt,name=createtime,proto3" json:"createtime,omitempty"`
	LocalId              uint64                              `protobuf:"varint,3,opt,name=local_id,json=localId,proto3" json:"local_id,omitempty"`
	Seq                  uint64                              `protobuf:"varint,4,opt,name=seq,proto3" json:"seq,omitempty"`
	Hash                 []byte                              `protobuf:"bytes,6,opt,name=hash,proto3" json:"hash,omitempty"`
	ProxyId              int32                               `protobuf:"varint,7,opt,name=proxy_id,json=proxyId,proto3" json:"proxy_id,omitempty"`
	ExData               []byte                              `protobuf:"bytes,8,opt,name=ex_data,json=exData,proto3" json:"ex_data,omitempty"`
	SystemData           bool                                `protobuf:"varint,9,opt,name=system_data,json=systemData,proto3" json:"system_data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                            `json:"-"`
	XXX_unrecognized     []byte                              `json:"-"`
	XXX_sizecache        int32                               `json:"-"`
}

func (m *BatchClientRequest) Reset()         { *m = BatchClientRequest{} }
func (m *BatchClientRequest) String() string { return proto.CompactTextString(m) }
func (*BatchClientRequest) ProtoMessage()    {}
func (*BatchClientRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_0a81bd5ef5ef03a2, []int{2}
}

func (m *BatchClientRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BatchClientRequest.Unmarshal(m, b)
}
func (m *BatchClientRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BatchClientRequest.Marshal(b, m, deterministic)
}
func (m *BatchClientRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BatchClientRequest.Merge(m, src)
}
func (m *BatchClientRequest) XXX_Size() int {
	return xxx_messageInfo_BatchClientRequest.Size(m)
}
func (m *BatchClientRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_BatchClientRequest.DiscardUnknown(m)
}

var xxx_messageInfo_BatchClientRequest proto.InternalMessageInfo

func (m *BatchClientRequest) GetClientRequests() []*BatchClientRequest_ClientRequest {
	if m != nil {
		return m.ClientRequests
	}
	return nil
}

func (m *BatchClientRequest) GetCreatetime() uint64 {
	if m != nil {
		return m.Createtime
	}
	return 0
}

func (m *BatchClientRequest) GetLocalId() uint64 {
	if m != nil {
		return m.LocalId
	}
	return 0
}

func (m *BatchClientRequest) GetSeq() uint64 {
	if m != nil {
		return m.Seq
	}
	return 0
}

func (m *BatchClientRequest) GetHash() []byte {
	if m != nil {
		return m.Hash
	}
	return nil
}

func (m *BatchClientRequest) GetProxyId() int32 {
	if m != nil {
		return m.ProxyId
	}
	return 0
}

func (m *BatchClientRequest) GetExData() []byte {
	if m != nil {
		return m.ExData
	}
	return nil
}

func (m *BatchClientRequest) GetSystemData() bool {
	if m != nil {
		return m.SystemData
	}
	return false
}

type BatchClientRequest_ClientRequest struct {
	Request              *Request `protobuf:"bytes,1,opt,name=request,proto3" json:"request,omitempty"`
	Id                   int32    `protobuf:"varint,3,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BatchClientRequest_ClientRequest) Reset()         { *m = BatchClientRequest_ClientRequest{} }
func (m *BatchClientRequest_ClientRequest) String() string { return proto.CompactTextString(m) }
func (*BatchClientRequest_ClientRequest) ProtoMessage()    {}
func (*BatchClientRequest_ClientRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_0a81bd5ef5ef03a2, []int{2, 0}
}

func (m *BatchClientRequest_ClientRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BatchClientRequest_ClientRequest.Unmarshal(m, b)
}
func (m *BatchClientRequest_ClientRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BatchClientRequest_ClientRequest.Marshal(b, m, deterministic)
}
func (m *BatchClientRequest_ClientRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BatchClientRequest_ClientRequest.Merge(m, src)
}
func (m *BatchClientRequest_ClientRequest) XXX_Size() int {
	return xxx_messageInfo_BatchClientRequest_ClientRequest.Size(m)
}
func (m *BatchClientRequest_ClientRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_BatchClientRequest_ClientRequest.DiscardUnknown(m)
}

var xxx_messageInfo_BatchClientRequest_ClientRequest proto.InternalMessageInfo

func (m *BatchClientRequest_ClientRequest) GetRequest() *Request {
	if m != nil {
		return m.Request
	}
	return nil
}

func (m *BatchClientRequest_ClientRequest) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

type CustomQueryRequest struct {
	RequestStr           []byte   `protobuf:"bytes,1,opt,name=request_str,json=requestStr,proto3" json:"request_str,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CustomQueryRequest) Reset()         { *m = CustomQueryRequest{} }
func (m *CustomQueryRequest) String() string { return proto.CompactTextString(m) }
func (*CustomQueryRequest) ProtoMessage()    {}
func (*CustomQueryRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_0a81bd5ef5ef03a2, []int{3}
}

func (m *CustomQueryRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CustomQueryRequest.Unmarshal(m, b)
}
func (m *CustomQueryRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CustomQueryRequest.Marshal(b, m, deterministic)
}
func (m *CustomQueryRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CustomQueryRequest.Merge(m, src)
}
func (m *CustomQueryRequest) XXX_Size() int {
	return xxx_messageInfo_CustomQueryRequest.Size(m)
}
func (m *CustomQueryRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CustomQueryRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CustomQueryRequest proto.InternalMessageInfo

func (m *CustomQueryRequest) GetRequestStr() []byte {
	if m != nil {
		return m.RequestStr
	}
	return nil
}

type CustomQueryResponse struct {
	RespStr              []byte   `protobuf:"bytes,1,opt,name=resp_str,json=respStr,proto3" json:"resp_str,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CustomQueryResponse) Reset()         { *m = CustomQueryResponse{} }
func (m *CustomQueryResponse) String() string { return proto.CompactTextString(m) }
func (*CustomQueryResponse) ProtoMessage()    {}
func (*CustomQueryResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_0a81bd5ef5ef03a2, []int{4}
}

func (m *CustomQueryResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CustomQueryResponse.Unmarshal(m, b)
}
func (m *CustomQueryResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CustomQueryResponse.Marshal(b, m, deterministic)
}
func (m *CustomQueryResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CustomQueryResponse.Merge(m, src)
}
func (m *CustomQueryResponse) XXX_Size() int {
	return xxx_messageInfo_CustomQueryResponse.Size(m)
}
func (m *CustomQueryResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CustomQueryResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CustomQueryResponse proto.InternalMessageInfo

func (m *CustomQueryResponse) GetRespStr() []byte {
	if m != nil {
		return m.RespStr
	}
	return nil
}

func init() {
	proto.RegisterEnum("XXXX.Request_Type", Request_Type_name, Request_Type_value)
	proto.RegisterType((*XDBMessage)(nil), "XXXX.XDBMessage")
	proto.RegisterType((*Request)(nil), "XXXX.Request")
	proto.RegisterType((*BatchClientRequest)(nil), "XXXX.BatchClientRequest")
	proto.RegisterType((*BatchClientRequest_ClientRequest)(nil), "XXXX.BatchClientRequest.ClientRequest")
	proto.RegisterType((*CustomQueryRequest)(nil), "XXXX.CustomQueryRequest")
	proto.RegisterType((*CustomQueryResponse)(nil), "XXXX.CustomQueryResponse")
}

func init() { proto.RegisterFile("XXXX.proto", fileDescriptor_0a81bd5ef5ef03a2) }

var fileDescriptor_0a81bd5ef5ef03a2 = []byte{
	// 725 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x54, 0xdb, 0x4e, 0xe3, 0x48,
	0x10, 0xdd, 0x5c, 0x9d, 0x94, 0x73, 0xe9, 0x74, 0x58, 0x30, 0x20, 0x2d, 0xd9, 0xec, 0xc3, 0x46,
	0xfb, 0x80, 0x10, 0xab, 0xfd, 0x00, 0x63, 0x7a, 0xc1, 0x5a, 0x62, 0x87, 0xb6, 0x03, 0xcb, 0x53,
	0x2b, 0xc4, 0xad, 0xc1, 0x12, 0x24, 0xc1, 0xdd, 0x19, 0x92, 0x6f, 0x9a, 0x0f, 0x9b, 0x2f, 0x98,
	0xf7, 0x51, 0xb7, 0x2f, 0x13, 0x84, 0x34, 0x0f, 0x96, 0xaa, 0xce, 0x71, 0x55, 0x9f, 0x6a, 0xf5,
	0x29, 0x30, 0x13, 0x2e, 0xa2, 0xc7, 0xd3, 0x55, 0xb2, 0x94, 0x4b, 0x5c, 0xd3, 0xc9, 0x70, 0x08,
	0x2d, 0xca, 0xc5, 0xe5, 0xc5, 0x98, 0x0b, 0x31, 0xfb, 0xc4, 0x31, 0x86, 0x6a, 0x34, 0x93, 0x33,
	0xab, 0x34, 0x28, 0x8d, 0x5a, 0x54, 0xc7, 0xc3, 0x6f, 0x35, 0x30, 0x28, 0x7f, 0x5d, 0x73, 0x21,
	0x15, 0x2f, 0xb7, 0x2b, 0xae, 0xf9, 0x1a, 0xd5, 0x71, 0x51, 0x53, 0xfe, 0x51, 0x83, 0x7f, 0x87,
	0xd6, 0x7c, 0x9d, 0x24, 0x7c, 0x21, 0xd9, 0xe7, 0x98, 0xbf, 0x59, 0xd5, 0x41, 0x69, 0x54, 0xa5,
	0x66, 0x86, 0xdd, 0xc5, 0xfc, 0x0d, 0x23, 0xa8, 0x08, 0xfe, 0x6a, 0xd5, 0x34, 0xa3, 0x42, 0xd5,
	0xe8, 0x69, 0x26, 0x9e, 0xac, 0x7a, 0xda, 0x48, 0xc5, 0xf8, 0x18, 0x9a, 0x82, 0x2f, 0x22, 0x9e,
	0xb0, 0x38, 0xb2, 0x0c, 0x7d, 0x6a, 0x23, 0x05, 0xdc, 0x08, 0x1f, 0x42, 0x63, 0x95, 0x2c, 0x37,
	0x5b, 0xc5, 0x35, 0x06, 0xa5, 0x51, 0x85, 0x1a, 0x3a, 0x77, 0x23, 0xfc, 0x17, 0xf4, 0x62, 0xc1,
	0xc4, 0x56, 0x48, 0xfe, 0xc2, 0x92, 0x54, 0xbd, 0xd5, 0x1c, 0x94, 0x46, 0x0d, 0xda, 0x8d, 0x45,
	0xa0, 0xf1, 0x7c, 0xa8, 0x33, 0xd8, 0xcb, 0xc5, 0xf2, 0x0d, 0x9f, 0xaf, 0x25, 0x8f, 0x98, 0x92,
	0x06, 0x5a, 0x1a, 0xce, 0x38, 0x92, 0x51, 0x01, 0x7f, 0xc5, 0x7f, 0x40, 0x7b, 0xc1, 0x79, 0xc4,
	0x12, 0x2e, 0x56, 0xcb, 0x85, 0xe0, 0x96, 0xa9, 0x3b, 0xb7, 0x14, 0x48, 0x33, 0x4c, 0x0d, 0x98,
	0x70, 0x69, 0xb5, 0xb4, 0x68, 0x15, 0x0e, 0xbf, 0x54, 0xa0, 0x1a, 0xaa, 0x2b, 0x6b, 0x43, 0x33,
	0x7c, 0x98, 0x10, 0xe6, 0xf9, 0x1e, 0x41, 0xbf, 0xe0, 0x3e, 0x74, 0x75, 0x7a, 0x4d, 0x6c, 0x1a,
	0xb2, 0x0b, 0x62, 0x87, 0xa8, 0x84, 0x0f, 0xa0, 0xaf, 0x41, 0xe7, 0xc6, 0x25, 0x5e, 0xc8, 0x28,
	0xb9, 0x9d, 0x92, 0x20, 0x44, 0x65, 0xbc, 0x07, 0x48, 0x13, 0x13, 0xaa, 0xbf, 0x89, 0x4d, 0x09,
	0xaa, 0x60, 0x04, 0xad, 0x1c, 0xd5, 0x48, 0x15, 0x77, 0xc1, 0x4c, 0x1b, 0xf8, 0xe3, 0xb1, 0x1b,
	0xa2, 0x5a, 0x51, 0x98, 0x75, 0x74, 0x08, 0x0d, 0x51, 0x1d, 0xf7, 0xa0, 0xad, 0x51, 0x4a, 0x82,
	0x89, 0xef, 0x05, 0x04, 0x19, 0x78, 0x1f, 0x70, 0x06, 0x39, 0xfe, 0x1d, 0xa1, 0x0f, 0xec, 0xd2,
	0x0e, 0x6d, 0xd4, 0xc0, 0xc7, 0x70, 0xf0, 0x11, 0xd7, 0x85, 0xa8, 0x59, 0x0c, 0xe1, 0x5c, 0x13,
	0xe7, 0xbf, 0x89, 0xef, 0x7a, 0x21, 0x02, 0xdc, 0x01, 0xd0, 0xe0, 0xed, 0x94, 0xd0, 0x07, 0x64,
	0xee, 0x74, 0x9e, 0xdc, 0xb8, 0x8e, 0xcd, 0x82, 0xd0, 0x0e, 0x09, 0x6a, 0x15, 0x22, 0x3c, 0x72,
	0xcf, 0xc2, 0xff, 0xbd, 0x00, 0x75, 0x0a, 0xb5, 0x57, 0xc4, 0x2f, 0x86, 0xef, 0x16, 0xa7, 0xdc,
	0xb9, 0xe4, 0xde, 0xb9, 0xb6, 0xbd, 0x2b, 0x82, 0x50, 0x31, 0xbb, 0x47, 0xee, 0x15, 0x8e, 0x7a,
	0xf8, 0x57, 0xe8, 0xa5, 0x62, 0xa6, 0x41, 0xe8, 0x8f, 0xb3, 0xe3, 0x31, 0x3e, 0x82, 0xfd, 0xa2,
	0xe7, 0xd8, 0xf5, 0x5c, 0xef, 0x4a, 0xa9, 0x9f, 0xde, 0x84, 0xa8, 0xaf, 0xae, 0xcb, 0x9b, 0x8e,
	0x99, 0xff, 0x2f, 0x53, 0xbf, 0xa0, 0xbd, 0xe1, 0xd7, 0x32, 0xe0, 0x8b, 0x99, 0x9c, 0x3f, 0x39,
	0xcf, 0x31, 0x5f, 0xc8, 0xfc, 0xb5, 0x4c, 0xa0, 0x3b, 0xd7, 0x40, 0xfe, 0xac, 0x84, 0x55, 0x1a,
	0x54, 0x46, 0xe6, 0xf9, 0x9f, 0xa7, 0xa9, 0xc1, 0x3e, 0xd6, 0x9c, 0xbe, 0xcb, 0x68, 0x67, 0xbe,
	0x9b, 0x0a, 0xfc, 0x1b, 0xc0, 0x3c, 0xe1, 0x33, 0xc9, 0x65, 0xfc, 0xc2, 0xb5, 0x8d, 0xaa, 0x74,
	0x07, 0x51, 0xcf, 0xfc, 0x79, 0x39, 0x9f, 0x3d, 0xab, 0x67, 0x5e, 0xd1, 0xac, 0xa1, 0x73, 0x37,
	0xca, 0x4d, 0x54, 0xfd, 0xb9, 0x89, 0x76, 0x7d, 0x92, 0x7a, 0xa8, 0xf0, 0xc9, 0x01, 0x18, 0x7c,
	0xc3, 0xb4, 0x7f, 0x1b, 0xba, 0xa2, 0xce, 0x37, 0x97, 0xca, 0xc1, 0x27, 0x60, 0x66, 0xee, 0xd1,
	0x64, 0x6a, 0x1d, 0x48, 0x21, 0xf5, 0xc3, 0x91, 0x0b, 0xed, 0xf7, 0x17, 0x33, 0x02, 0x23, 0x37,
	0x9a, 0x5a, 0x0f, 0xe6, 0x79, 0x27, 0xbb, 0x90, 0x7c, 0xee, 0x9c, 0xc6, 0x1d, 0x28, 0x67, 0xa3,
	0xd4, 0x68, 0x39, 0x8e, 0x86, 0xff, 0x00, 0x76, 0xd6, 0x42, 0x2e, 0x5f, 0x6e, 0xd7, 0x3c, 0xd9,
	0xe6, 0xfd, 0x4e, 0xd4, 0xc6, 0xd2, 0x21, 0x13, 0x32, 0xc9, 0x56, 0x12, 0x64, 0x50, 0x20, 0x93,
	0xe1, 0x19, 0xf4, 0xdf, 0x95, 0x65, 0xbe, 0x3b, 0x84, 0x86, 0xf2, 0xe5, 0x4e, 0x91, 0xa1, 0xf2,
	0x40, 0x26, 0x8f, 0x75, 0xbd, 0xfc, 0xfe, 0xfe, 0x1e, 0x00, 0x00, 0xff, 0xff, 0x2d, 0x5b, 0x29,
	0x68, 0x0b, 0x05, 0x00, 0x00,
}
