// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pow.proto

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

type PoWRequest int32

const (
	PoWRequest_NONE                 PoWRequest = 0
	PoWRequest_TYPE_COMMITTED_BLOCK PoWRequest = 101
	PoWRequest_TYPE_SHIFT_MSG       PoWRequest = 102
	PoWRequest_TYPE_MINING_RESULTS  PoWRequest = 103
)

var PoWRequest_name = map[int32]string{
	0:   "NONE",
	101: "TYPE_COMMITTED_BLOCK",
	102: "TYPE_SHIFT_MSG",
	103: "TYPE_MINING_RESULTS",
}

var PoWRequest_value = map[string]int32{
	"NONE":                 0,
	"TYPE_COMMITTED_BLOCK": 101,
	"TYPE_SHIFT_MSG":       102,
	"TYPE_MINING_RESULTS":  103,
}

func (x PoWRequest) String() string {
	return proto.EnumName(PoWRequest_name, int32(x))
}

func (PoWRequest) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_ebb06f9e481a48a8, []int{0}
}

type ClientTransactions struct {
	TransactionData      []byte   `protobuf:"bytes,1,opt,name=transaction_data,json=transactionData,proto3" json:"transaction_data,omitempty"`
	Seq                  uint64   `protobuf:"varint,2,opt,name=seq,proto3" json:"seq,omitempty"`
	CreateTime           uint64   `protobuf:"varint,3,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ClientTransactions) Reset()         { *m = ClientTransactions{} }
func (m *ClientTransactions) String() string { return proto.CompactTextString(m) }
func (*ClientTransactions) ProtoMessage()    {}
func (*ClientTransactions) Descriptor() ([]byte, []int) {
	return fileDescriptor_ebb06f9e481a48a8, []int{0}
}

func (m *ClientTransactions) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClientTransactions.Unmarshal(m, b)
}
func (m *ClientTransactions) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClientTransactions.Marshal(b, m, deterministic)
}
func (m *ClientTransactions) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClientTransactions.Merge(m, src)
}
func (m *ClientTransactions) XXX_Size() int {
	return xxx_messageInfo_ClientTransactions.Size(m)
}
func (m *ClientTransactions) XXX_DiscardUnknown() {
	xxx_messageInfo_ClientTransactions.DiscardUnknown(m)
}

var xxx_messageInfo_ClientTransactions proto.InternalMessageInfo

func (m *ClientTransactions) GetTransactionData() []byte {
	if m != nil {
		return m.TransactionData
	}
	return nil
}

func (m *ClientTransactions) GetSeq() uint64 {
	if m != nil {
		return m.Seq
	}
	return 0
}

func (m *ClientTransactions) GetCreateTime() uint64 {
	if m != nil {
		return m.CreateTime
	}
	return 0
}

type BatchClientTransactions struct {
	Transactions         []*ClientTransactions `protobuf:"bytes,1,rep,name=transactions,proto3" json:"transactions,omitempty"`
	MinSeq               uint64                `protobuf:"varint,2,opt,name=min_seq,json=minSeq,proto3" json:"min_seq,omitempty"`
	MaxSeq               uint64                `protobuf:"varint,3,opt,name=max_seq,json=maxSeq,proto3" json:"max_seq,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *BatchClientTransactions) Reset()         { *m = BatchClientTransactions{} }
func (m *BatchClientTransactions) String() string { return proto.CompactTextString(m) }
func (*BatchClientTransactions) ProtoMessage()    {}
func (*BatchClientTransactions) Descriptor() ([]byte, []int) {
	return fileDescriptor_ebb06f9e481a48a8, []int{1}
}

func (m *BatchClientTransactions) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BatchClientTransactions.Unmarshal(m, b)
}
func (m *BatchClientTransactions) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BatchClientTransactions.Marshal(b, m, deterministic)
}
func (m *BatchClientTransactions) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BatchClientTransactions.Merge(m, src)
}
func (m *BatchClientTransactions) XXX_Size() int {
	return xxx_messageInfo_BatchClientTransactions.Size(m)
}
func (m *BatchClientTransactions) XXX_DiscardUnknown() {
	xxx_messageInfo_BatchClientTransactions.DiscardUnknown(m)
}

var xxx_messageInfo_BatchClientTransactions proto.InternalMessageInfo

func (m *BatchClientTransactions) GetTransactions() []*ClientTransactions {
	if m != nil {
		return m.Transactions
	}
	return nil
}

func (m *BatchClientTransactions) GetMinSeq() uint64 {
	if m != nil {
		return m.MinSeq
	}
	return 0
}

func (m *BatchClientTransactions) GetMaxSeq() uint64 {
	if m != nil {
		return m.MaxSeq
	}
	return 0
}

type SliceInfo struct {
	Height               uint64   `protobuf:"varint,1,opt,name=height,proto3" json:"height,omitempty"`
	ShiftIdx             int32    `protobuf:"varint,2,opt,name=shift_idx,json=shiftIdx,proto3" json:"shift_idx,omitempty"`
	Sender               int32    `protobuf:"varint,3,opt,name=sender,proto3" json:"sender,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SliceInfo) Reset()         { *m = SliceInfo{} }
func (m *SliceInfo) String() string { return proto.CompactTextString(m) }
func (*SliceInfo) ProtoMessage()    {}
func (*SliceInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_ebb06f9e481a48a8, []int{2}
}

func (m *SliceInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SliceInfo.Unmarshal(m, b)
}
func (m *SliceInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SliceInfo.Marshal(b, m, deterministic)
}
func (m *SliceInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SliceInfo.Merge(m, src)
}
func (m *SliceInfo) XXX_Size() int {
	return xxx_messageInfo_SliceInfo.Size(m)
}
func (m *SliceInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_SliceInfo.DiscardUnknown(m)
}

var xxx_messageInfo_SliceInfo proto.InternalMessageInfo

func (m *SliceInfo) GetHeight() uint64 {
	if m != nil {
		return m.Height
	}
	return 0
}

func (m *SliceInfo) GetShiftIdx() int32 {
	if m != nil {
		return m.ShiftIdx
	}
	return 0
}

func (m *SliceInfo) GetSender() int32 {
	if m != nil {
		return m.Sender
	}
	return 0
}

// 256 bits hash value
type HashValue struct {
	Bits                 []uint64 `protobuf:"varint,1,rep,packed,name=bits,proto3" json:"bits,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HashValue) Reset()         { *m = HashValue{} }
func (m *HashValue) String() string { return proto.CompactTextString(m) }
func (*HashValue) ProtoMessage()    {}
func (*HashValue) Descriptor() ([]byte, []int) {
	return fileDescriptor_ebb06f9e481a48a8, []int{3}
}

func (m *HashValue) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HashValue.Unmarshal(m, b)
}
func (m *HashValue) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HashValue.Marshal(b, m, deterministic)
}
func (m *HashValue) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HashValue.Merge(m, src)
}
func (m *HashValue) XXX_Size() int {
	return xxx_messageInfo_HashValue.Size(m)
}
func (m *HashValue) XXX_DiscardUnknown() {
	xxx_messageInfo_HashValue.DiscardUnknown(m)
}

var xxx_messageInfo_HashValue proto.InternalMessageInfo

func (m *HashValue) GetBits() []uint64 {
	if m != nil {
		return m.Bits
	}
	return nil
}

type BlockHeader struct {
	Height               uint64     `protobuf:"varint,1,opt,name=height,proto3" json:"height,omitempty"`
	PreHash              *HashValue `protobuf:"bytes,2,opt,name=pre_hash,json=preHash,proto3" json:"pre_hash,omitempty"`
	MerkleHash           *HashValue `protobuf:"bytes,3,opt,name=merkle_hash,json=merkleHash,proto3" json:"merkle_hash,omitempty"`
	Nonce                uint64     `protobuf:"varint,4,opt,name=nonce,proto3" json:"nonce,omitempty"`
	MinSeq               uint64     `protobuf:"varint,5,opt,name=min_seq,json=minSeq,proto3" json:"min_seq,omitempty"`
	MaxSeq               uint64     `protobuf:"varint,6,opt,name=max_seq,json=maxSeq,proto3" json:"max_seq,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *BlockHeader) Reset()         { *m = BlockHeader{} }
func (m *BlockHeader) String() string { return proto.CompactTextString(m) }
func (*BlockHeader) ProtoMessage()    {}
func (*BlockHeader) Descriptor() ([]byte, []int) {
	return fileDescriptor_ebb06f9e481a48a8, []int{4}
}

func (m *BlockHeader) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BlockHeader.Unmarshal(m, b)
}
func (m *BlockHeader) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BlockHeader.Marshal(b, m, deterministic)
}
func (m *BlockHeader) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BlockHeader.Merge(m, src)
}
func (m *BlockHeader) XXX_Size() int {
	return xxx_messageInfo_BlockHeader.Size(m)
}
func (m *BlockHeader) XXX_DiscardUnknown() {
	xxx_messageInfo_BlockHeader.DiscardUnknown(m)
}

var xxx_messageInfo_BlockHeader proto.InternalMessageInfo

func (m *BlockHeader) GetHeight() uint64 {
	if m != nil {
		return m.Height
	}
	return 0
}

func (m *BlockHeader) GetPreHash() *HashValue {
	if m != nil {
		return m.PreHash
	}
	return nil
}

func (m *BlockHeader) GetMerkleHash() *HashValue {
	if m != nil {
		return m.MerkleHash
	}
	return nil
}

func (m *BlockHeader) GetNonce() uint64 {
	if m != nil {
		return m.Nonce
	}
	return 0
}

func (m *BlockHeader) GetMinSeq() uint64 {
	if m != nil {
		return m.MinSeq
	}
	return 0
}

func (m *BlockHeader) GetMaxSeq() uint64 {
	if m != nil {
		return m.MaxSeq
	}
	return 0
}

type Block struct {
	Header               *BlockHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	TransactionData      []byte       `protobuf:"bytes,2,opt,name=transaction_data,json=transactionData,proto3" json:"transaction_data,omitempty"`
	Hash                 *HashValue   `protobuf:"bytes,3,opt,name=hash,proto3" json:"hash,omitempty"`
	Miner                uint64       `protobuf:"varint,6,opt,name=miner,proto3" json:"miner,omitempty"`
	BlockTime            uint64       `protobuf:"varint,7,opt,name=block_time,json=blockTime,proto3" json:"block_time,omitempty"`
	MiningTime           uint64       `protobuf:"varint,8,opt,name=mining_time,json=miningTime,proto3" json:"mining_time,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *Block) Reset()         { *m = Block{} }
func (m *Block) String() string { return proto.CompactTextString(m) }
func (*Block) ProtoMessage()    {}
func (*Block) Descriptor() ([]byte, []int) {
	return fileDescriptor_ebb06f9e481a48a8, []int{5}
}

func (m *Block) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Block.Unmarshal(m, b)
}
func (m *Block) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Block.Marshal(b, m, deterministic)
}
func (m *Block) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Block.Merge(m, src)
}
func (m *Block) XXX_Size() int {
	return xxx_messageInfo_Block.Size(m)
}
func (m *Block) XXX_DiscardUnknown() {
	xxx_messageInfo_Block.DiscardUnknown(m)
}

var xxx_messageInfo_Block proto.InternalMessageInfo

func (m *Block) GetHeader() *BlockHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *Block) GetTransactionData() []byte {
	if m != nil {
		return m.TransactionData
	}
	return nil
}

func (m *Block) GetHash() *HashValue {
	if m != nil {
		return m.Hash
	}
	return nil
}

func (m *Block) GetMiner() uint64 {
	if m != nil {
		return m.Miner
	}
	return 0
}

func (m *Block) GetBlockTime() uint64 {
	if m != nil {
		return m.BlockTime
	}
	return 0
}

func (m *Block) GetMiningTime() uint64 {
	if m != nil {
		return m.MiningTime
	}
	return 0
}

type BlockMiningInfo struct {
	Header               *BlockHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	Hash                 *HashValue   `protobuf:"bytes,3,opt,name=hash,proto3" json:"hash,omitempty"`
	Miner                uint64       `protobuf:"varint,6,opt,name=miner,proto3" json:"miner,omitempty"`
	BlockTime            uint64       `protobuf:"varint,7,opt,name=block_time,json=blockTime,proto3" json:"block_time,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *BlockMiningInfo) Reset()         { *m = BlockMiningInfo{} }
func (m *BlockMiningInfo) String() string { return proto.CompactTextString(m) }
func (*BlockMiningInfo) ProtoMessage()    {}
func (*BlockMiningInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_ebb06f9e481a48a8, []int{6}
}

func (m *BlockMiningInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BlockMiningInfo.Unmarshal(m, b)
}
func (m *BlockMiningInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BlockMiningInfo.Marshal(b, m, deterministic)
}
func (m *BlockMiningInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BlockMiningInfo.Merge(m, src)
}
func (m *BlockMiningInfo) XXX_Size() int {
	return xxx_messageInfo_BlockMiningInfo.Size(m)
}
func (m *BlockMiningInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_BlockMiningInfo.DiscardUnknown(m)
}

var xxx_messageInfo_BlockMiningInfo proto.InternalMessageInfo

func (m *BlockMiningInfo) GetHeader() *BlockHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *BlockMiningInfo) GetHash() *HashValue {
	if m != nil {
		return m.Hash
	}
	return nil
}

func (m *BlockMiningInfo) GetMiner() uint64 {
	if m != nil {
		return m.Miner
	}
	return 0
}

func (m *BlockMiningInfo) GetBlockTime() uint64 {
	if m != nil {
		return m.BlockTime
	}
	return 0
}

type TxnQueryRequest struct {
	MinSeq               uint64   `protobuf:"varint,1,opt,name=min_seq,json=minSeq,proto3" json:"min_seq,omitempty"`
	MaxSeq               uint64   `protobuf:"varint,2,opt,name=max_seq,json=maxSeq,proto3" json:"max_seq,omitempty"`
	IsQueryResults       uint64   `protobuf:"varint,3,opt,name=is_query_results,json=isQueryResults,proto3" json:"is_query_results,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TxnQueryRequest) Reset()         { *m = TxnQueryRequest{} }
func (m *TxnQueryRequest) String() string { return proto.CompactTextString(m) }
func (*TxnQueryRequest) ProtoMessage()    {}
func (*TxnQueryRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ebb06f9e481a48a8, []int{7}
}

func (m *TxnQueryRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TxnQueryRequest.Unmarshal(m, b)
}
func (m *TxnQueryRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TxnQueryRequest.Marshal(b, m, deterministic)
}
func (m *TxnQueryRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TxnQueryRequest.Merge(m, src)
}
func (m *TxnQueryRequest) XXX_Size() int {
	return xxx_messageInfo_TxnQueryRequest.Size(m)
}
func (m *TxnQueryRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_TxnQueryRequest.DiscardUnknown(m)
}

var xxx_messageInfo_TxnQueryRequest proto.InternalMessageInfo

func (m *TxnQueryRequest) GetMinSeq() uint64 {
	if m != nil {
		return m.MinSeq
	}
	return 0
}

func (m *TxnQueryRequest) GetMaxSeq() uint64 {
	if m != nil {
		return m.MaxSeq
	}
	return 0
}

func (m *TxnQueryRequest) GetIsQueryResults() uint64 {
	if m != nil {
		return m.IsQueryResults
	}
	return 0
}

type TxnQueryResponse struct {
	Seq                  []uint64 `protobuf:"varint,1,rep,packed,name=seq,proto3" json:"seq,omitempty"`
	Data                 [][]byte `protobuf:"bytes,2,rep,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TxnQueryResponse) Reset()         { *m = TxnQueryResponse{} }
func (m *TxnQueryResponse) String() string { return proto.CompactTextString(m) }
func (*TxnQueryResponse) ProtoMessage()    {}
func (*TxnQueryResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_ebb06f9e481a48a8, []int{8}
}

func (m *TxnQueryResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TxnQueryResponse.Unmarshal(m, b)
}
func (m *TxnQueryResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TxnQueryResponse.Marshal(b, m, deterministic)
}
func (m *TxnQueryResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TxnQueryResponse.Merge(m, src)
}
func (m *TxnQueryResponse) XXX_Size() int {
	return xxx_messageInfo_TxnQueryResponse.Size(m)
}
func (m *TxnQueryResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_TxnQueryResponse.DiscardUnknown(m)
}

var xxx_messageInfo_TxnQueryResponse proto.InternalMessageInfo

func (m *TxnQueryResponse) GetSeq() []uint64 {
	if m != nil {
		return m.Seq
	}
	return nil
}

func (m *TxnQueryResponse) GetData() [][]byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterEnum("XXXX.PoWRequest", PoWRequest_name, PoWRequest_value)
	proto.RegisterType((*ClientTransactions)(nil), "XXXX.ClientTransactions")
	proto.RegisterType((*BatchClientTransactions)(nil), "XXXX.BatchClientTransactions")
	proto.RegisterType((*SliceInfo)(nil), "XXXX.SliceInfo")
	proto.RegisterType((*HashValue)(nil), "XXXX.HashValue")
	proto.RegisterType((*BlockHeader)(nil), "XXXX.BlockHeader")
	proto.RegisterType((*Block)(nil), "XXXX.Block")
	proto.RegisterType((*BlockMiningInfo)(nil), "XXXX.BlockMiningInfo")
	proto.RegisterType((*TxnQueryRequest)(nil), "XXXX.TxnQueryRequest")
	proto.RegisterType((*TxnQueryResponse)(nil), "XXXX.TxnQueryResponse")
}

func init() { proto.RegisterFile("pow.proto", fileDescriptor_ebb06f9e481a48a8) }

var fileDescriptor_ebb06f9e481a48a8 = []byte{
	// 602 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x54, 0xcd, 0x6e, 0xd3, 0x40,
	0x10, 0xc6, 0x8d, 0x93, 0x26, 0x93, 0xaa, 0xb5, 0x96, 0x8a, 0x06, 0x21, 0xd4, 0xca, 0xe2, 0x10,
	0x8a, 0x54, 0x89, 0x72, 0xe1, 0xc2, 0xa5, 0x6d, 0x68, 0x23, 0x9a, 0xb4, 0xac, 0xcd, 0xdf, 0x01,
	0x59, 0x9b, 0x64, 0x5a, 0xaf, 0x6a, 0xaf, 0x93, 0xdd, 0x8d, 0x08, 0xaf, 0xc0, 0x13, 0xf0, 0x52,
	0xbc, 0x00, 0x4f, 0x83, 0x3c, 0x76, 0xa9, 0x2b, 0x1a, 0x24, 0x0e, 0xdc, 0x76, 0xe6, 0x9b, 0x9d,
	0x6f, 0xbe, 0xdd, 0x6f, 0x17, 0x5a, 0xd3, 0xec, 0xcb, 0xde, 0x54, 0x67, 0x36, 0x63, 0x75, 0x8d,
	0x66, 0x32, 0xf2, 0x35, 0xb0, 0xc3, 0x44, 0xa2, 0xb2, 0xa1, 0x16, 0xca, 0x88, 0xb1, 0x95, 0x99,
	0x32, 0xec, 0x29, 0x78, 0xf6, 0x26, 0x8e, 0x26, 0xc2, 0x8a, 0x8e, 0xb3, 0xe3, 0x74, 0xd7, 0xf8,
	0x46, 0x25, 0x7f, 0x24, 0xac, 0x60, 0x1e, 0xd4, 0x0c, 0xce, 0x3a, 0x2b, 0x3b, 0x4e, 0xd7, 0xe5,
	0xf9, 0x92, 0x6d, 0x43, 0x7b, 0xac, 0x51, 0x58, 0x8c, 0xac, 0x4c, 0xb1, 0x53, 0x23, 0x04, 0x8a,
	0x54, 0x28, 0x53, 0xf4, 0xbf, 0x39, 0xb0, 0x75, 0x20, 0xec, 0x38, 0xbe, 0x83, 0xf9, 0x15, 0xac,
	0x55, 0x18, 0x4c, 0xc7, 0xd9, 0xa9, 0x75, 0xdb, 0xfb, 0x0f, 0xf7, 0x68, 0xda, 0xbd, 0x3f, 0x37,
	0xf0, 0x5b, 0xe5, 0x6c, 0x0b, 0x56, 0x53, 0xa9, 0xa2, 0x9b, 0x89, 0x1a, 0xa9, 0x54, 0x01, 0xce,
	0x08, 0x10, 0x0b, 0x02, 0x6a, 0x25, 0x20, 0x16, 0x01, 0xce, 0xfc, 0x8f, 0xd0, 0x0a, 0x12, 0x39,
	0xc6, 0xbe, 0xba, 0xc8, 0xd8, 0x03, 0x68, 0xc4, 0x28, 0x2f, 0x63, 0x4b, 0x6a, 0x5d, 0x5e, 0x46,
	0xec, 0x11, 0xb4, 0x4c, 0x2c, 0x2f, 0x6c, 0x24, 0x27, 0x0b, 0x6a, 0x5c, 0xe7, 0x4d, 0x4a, 0xf4,
	0x27, 0x8b, 0x7c, 0x93, 0x41, 0x35, 0x41, 0x4d, 0x9d, 0xeb, 0xbc, 0x8c, 0xfc, 0x6d, 0x68, 0x9d,
	0x08, 0x13, 0xbf, 0x17, 0xc9, 0x1c, 0x19, 0x03, 0x77, 0x24, 0x6d, 0xa1, 0xc7, 0xe5, 0xb4, 0xf6,
	0x7f, 0x38, 0xd0, 0x3e, 0x48, 0xb2, 0xf1, 0xd5, 0x09, 0x8a, 0x09, 0xea, 0xa5, 0xec, 0xcf, 0xa0,
	0x39, 0xd5, 0x18, 0xc5, 0xc2, 0xc4, 0x44, 0xde, 0xde, 0xf7, 0xca, 0xf3, 0xf8, 0xdd, 0x9f, 0xaf,
	0x4e, 0x35, 0xe6, 0x11, 0x7b, 0x0e, 0xed, 0x14, 0xf5, 0x55, 0x52, 0xd6, 0xd7, 0x96, 0xd4, 0x43,
	0x51, 0x44, 0x5b, 0x36, 0xa1, 0xae, 0x32, 0x35, 0xc6, 0x8e, 0x4b, 0xb4, 0x45, 0x50, 0x3d, 0xca,
	0xfa, 0xb2, 0xa3, 0x6c, 0xdc, 0x3a, 0xca, 0x9f, 0x0e, 0xd4, 0x49, 0x0f, 0xdb, 0xcd, 0x95, 0xe4,
	0x9a, 0x48, 0x49, 0x7b, 0x9f, 0x95, 0xfc, 0x15, 0xb5, 0xbc, 0xac, 0xb8, 0xd3, 0x6b, 0x2b, 0x77,
	0x7b, 0xed, 0x09, 0xb8, 0x7f, 0x15, 0x45, 0x68, 0x2e, 0x27, 0x95, 0x0a, 0x75, 0x39, 0x5d, 0x11,
	0xb0, 0xc7, 0x00, 0xa3, 0x9c, 0xbd, 0x30, 0xe5, 0x2a, 0x41, 0x2d, 0xca, 0xe4, 0x9e, 0xcc, 0x4d,
	0x9b, 0x4a, 0x25, 0xd5, 0x65, 0x81, 0x37, 0x0b, 0xd3, 0x16, 0x29, 0x32, 0xed, 0x77, 0x07, 0x36,
	0x68, 0xfc, 0x01, 0xe5, 0xc8, 0x2e, 0xff, 0x22, 0xf3, 0xff, 0xcd, 0xee, 0xa7, 0xb0, 0x11, 0x2e,
	0xd4, 0xdb, 0x39, 0xea, 0xaf, 0x1c, 0x67, 0x73, 0x34, 0xb6, 0x7a, 0x79, 0xce, 0xb2, 0xcb, 0x5b,
	0xa9, 0x5e, 0x1e, 0xeb, 0x82, 0x27, 0x4d, 0x34, 0xcb, 0x9b, 0x44, 0x1a, 0xcd, 0x3c, 0xb1, 0xa6,
	0x7c, 0x29, 0xeb, 0xd2, 0x94, 0xbd, 0x29, 0xeb, 0xbf, 0x04, 0xef, 0x86, 0xce, 0x4c, 0x33, 0x65,
	0xf0, 0xfa, 0x17, 0x28, 0xdc, 0x4d, 0xbf, 0x00, 0x03, 0xb7, 0xbc, 0xca, 0x5a, 0x77, 0x8d, 0xd3,
	0x7a, 0xf7, 0x33, 0xc0, 0x79, 0xf6, 0xe1, 0x7a, 0xc6, 0x26, 0xb8, 0xc3, 0xb3, 0x61, 0xcf, 0xbb,
	0xc7, 0x3a, 0xb0, 0x19, 0x7e, 0x3a, 0xef, 0x45, 0x87, 0x67, 0x83, 0x41, 0x3f, 0x0c, 0x7b, 0x47,
	0xd1, 0xc1, 0xe9, 0xd9, 0xe1, 0x1b, 0x2f, 0x7f, 0x36, 0xeb, 0x84, 0x04, 0x27, 0xfd, 0xd7, 0x61,
	0x34, 0x08, 0x8e, 0xbd, 0x0b, 0xb6, 0x05, 0xf7, 0x29, 0x37, 0xe8, 0x0f, 0xfb, 0xc3, 0xe3, 0x88,
	0xf7, 0x82, 0x77, 0xa7, 0x61, 0xe0, 0x5d, 0x8e, 0x1a, 0xf4, 0xb3, 0xbd, 0xf8, 0x15, 0x00, 0x00,
	0xff, 0xff, 0x72, 0x41, 0xb3, 0xc3, 0xe6, 0x04, 0x00, 0x00,
}
