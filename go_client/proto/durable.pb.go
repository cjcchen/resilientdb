// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.6.1
// source: durable.proto

package XXXX

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type RocksDBInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EnableRocksdb           bool   `protobuf:"varint,1,opt,name=enable_rocksdb,json=enableRocksdb,proto3" json:"enable_rocksdb,omitempty"`
	NumThreads              uint32 `protobuf:"varint,2,opt,name=num_threads,json=numThreads,proto3" json:"num_threads,omitempty"`
	WriteBufferSizeMb       uint32 `protobuf:"varint,3,opt,name=write_buffer_size_mb,json=writeBufferSizeMb,proto3" json:"write_buffer_size_mb,omitempty"`
	WriteBatchSize          uint32 `protobuf:"varint,4,opt,name=write_batch_size,json=writeBatchSize,proto3" json:"write_batch_size,omitempty"`
	Path                    string `protobuf:"bytes,5,opt,name=path,proto3" json:"path,omitempty"`
	GenerateUniquePathnames bool   `protobuf:"varint,6,opt,name=generate_unique_pathnames,json=generateUniquePathnames,proto3" json:"generate_unique_pathnames,omitempty"`
}

func (x *RocksDBInfo) Reset() {
	*x = RocksDBInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_durable_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RocksDBInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RocksDBInfo) ProtoMessage() {}

func (x *RocksDBInfo) ProtoReflect() protoreflect.Message {
	mi := &file_durable_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RocksDBInfo.ProtoReflect.Descriptor instead.
func (*RocksDBInfo) Descriptor() ([]byte, []int) {
	return file_durable_proto_rawDescGZIP(), []int{0}
}

func (x *RocksDBInfo) GetEnableRocksdb() bool {
	if x != nil {
		return x.EnableRocksdb
	}
	return false
}

func (x *RocksDBInfo) GetNumThreads() uint32 {
	if x != nil {
		return x.NumThreads
	}
	return 0
}

func (x *RocksDBInfo) GetWriteBufferSizeMb() uint32 {
	if x != nil {
		return x.WriteBufferSizeMb
	}
	return 0
}

func (x *RocksDBInfo) GetWriteBatchSize() uint32 {
	if x != nil {
		return x.WriteBatchSize
	}
	return 0
}

func (x *RocksDBInfo) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *RocksDBInfo) GetGenerateUniquePathnames() bool {
	if x != nil {
		return x.GenerateUniquePathnames
	}
	return false
}

type LevelDBInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EnableLeveldb           bool   `protobuf:"varint,1,opt,name=enable_leveldb,json=enableLeveldb,proto3" json:"enable_leveldb,omitempty"`
	WriteBufferSizeMb       uint32 `protobuf:"varint,2,opt,name=write_buffer_size_mb,json=writeBufferSizeMb,proto3" json:"write_buffer_size_mb,omitempty"`
	WriteBatchSize          uint32 `protobuf:"varint,3,opt,name=write_batch_size,json=writeBatchSize,proto3" json:"write_batch_size,omitempty"`
	Path                    string `protobuf:"bytes,4,opt,name=path,proto3" json:"path,omitempty"`
	GenerateUniquePathnames bool   `protobuf:"varint,5,opt,name=generate_unique_pathnames,json=generateUniquePathnames,proto3" json:"generate_unique_pathnames,omitempty"`
}

func (x *LevelDBInfo) Reset() {
	*x = LevelDBInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_durable_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LevelDBInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LevelDBInfo) ProtoMessage() {}

func (x *LevelDBInfo) ProtoReflect() protoreflect.Message {
	mi := &file_durable_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LevelDBInfo.ProtoReflect.Descriptor instead.
func (*LevelDBInfo) Descriptor() ([]byte, []int) {
	return file_durable_proto_rawDescGZIP(), []int{1}
}

func (x *LevelDBInfo) GetEnableLeveldb() bool {
	if x != nil {
		return x.EnableLeveldb
	}
	return false
}

func (x *LevelDBInfo) GetWriteBufferSizeMb() uint32 {
	if x != nil {
		return x.WriteBufferSizeMb
	}
	return 0
}

func (x *LevelDBInfo) GetWriteBatchSize() uint32 {
	if x != nil {
		return x.WriteBatchSize
	}
	return 0
}

func (x *LevelDBInfo) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *LevelDBInfo) GetGenerateUniquePathnames() bool {
	if x != nil {
		return x.GenerateUniquePathnames
	}
	return false
}

var File_durable_proto protoreflect.FileDescriptor

var file_durable_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x64, 0x75, 0x72, 0x61, 0x62, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x05, 0x72, 0x65, 0x73, 0x64, 0x62, 0x22, 0x80, 0x02, 0x0a, 0x0b, 0x52, 0x6f, 0x63, 0x6b, 0x73,
	0x44, 0x42, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x25, 0x0a, 0x0e, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65,
	0x5f, 0x72, 0x6f, 0x63, 0x6b, 0x73, 0x64, 0x62, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0d,
	0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x52, 0x6f, 0x63, 0x6b, 0x73, 0x64, 0x62, 0x12, 0x1f, 0x0a,
	0x0b, 0x6e, 0x75, 0x6d, 0x5f, 0x74, 0x68, 0x72, 0x65, 0x61, 0x64, 0x73, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x0a, 0x6e, 0x75, 0x6d, 0x54, 0x68, 0x72, 0x65, 0x61, 0x64, 0x73, 0x12, 0x2f,
	0x0a, 0x14, 0x77, 0x72, 0x69, 0x74, 0x65, 0x5f, 0x62, 0x75, 0x66, 0x66, 0x65, 0x72, 0x5f, 0x73,
	0x69, 0x7a, 0x65, 0x5f, 0x6d, 0x62, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x11, 0x77, 0x72,
	0x69, 0x74, 0x65, 0x42, 0x75, 0x66, 0x66, 0x65, 0x72, 0x53, 0x69, 0x7a, 0x65, 0x4d, 0x62, 0x12,
	0x28, 0x0a, 0x10, 0x77, 0x72, 0x69, 0x74, 0x65, 0x5f, 0x62, 0x61, 0x74, 0x63, 0x68, 0x5f, 0x73,
	0x69, 0x7a, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0e, 0x77, 0x72, 0x69, 0x74, 0x65,
	0x42, 0x61, 0x74, 0x63, 0x68, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x74,
	0x68, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x12, 0x3a, 0x0a,
	0x19, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x5f, 0x75, 0x6e, 0x69, 0x71, 0x75, 0x65,
	0x5f, 0x70, 0x61, 0x74, 0x68, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x17, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x55, 0x6e, 0x69, 0x71, 0x75, 0x65,
	0x50, 0x61, 0x74, 0x68, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x22, 0xdf, 0x01, 0x0a, 0x0b, 0x4c, 0x65,
	0x76, 0x65, 0x6c, 0x44, 0x42, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x25, 0x0a, 0x0e, 0x65, 0x6e, 0x61,
	0x62, 0x6c, 0x65, 0x5f, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x64, 0x62, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x0d, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x64, 0x62,
	0x12, 0x2f, 0x0a, 0x14, 0x77, 0x72, 0x69, 0x74, 0x65, 0x5f, 0x62, 0x75, 0x66, 0x66, 0x65, 0x72,
	0x5f, 0x73, 0x69, 0x7a, 0x65, 0x5f, 0x6d, 0x62, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x11,
	0x77, 0x72, 0x69, 0x74, 0x65, 0x42, 0x75, 0x66, 0x66, 0x65, 0x72, 0x53, 0x69, 0x7a, 0x65, 0x4d,
	0x62, 0x12, 0x28, 0x0a, 0x10, 0x77, 0x72, 0x69, 0x74, 0x65, 0x5f, 0x62, 0x61, 0x74, 0x63, 0x68,
	0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0e, 0x77, 0x72, 0x69,
	0x74, 0x65, 0x42, 0x61, 0x74, 0x63, 0x68, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x70,
	0x61, 0x74, 0x68, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x12,
	0x3a, 0x0a, 0x19, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x5f, 0x75, 0x6e, 0x69, 0x71,
	0x75, 0x65, 0x5f, 0x70, 0x61, 0x74, 0x68, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x17, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x55, 0x6e, 0x69, 0x71,
	0x75, 0x65, 0x50, 0x61, 0x74, 0x68, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x42, 0x0a, 0x5a, 0x08, 0x2e,
	0x2f, 0x3b, 0x72, 0x65, 0x73, 0x64, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_durable_proto_rawDescOnce sync.Once
	file_durable_proto_rawDescData = file_durable_proto_rawDesc
)

func file_durable_proto_rawDescGZIP() []byte {
	file_durable_proto_rawDescOnce.Do(func() {
		file_durable_proto_rawDescData = protoimpl.X.CompressGZIP(file_durable_proto_rawDescData)
	})
	return file_durable_proto_rawDescData
}

var file_durable_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_durable_proto_goTypes = []interface{}{
	(*RocksDBInfo)(nil), // 0: XXXX.RocksDBInfo
	(*LevelDBInfo)(nil), // 1: XXXX.LevelDBInfo
}
var file_durable_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_durable_proto_init() }
func file_durable_proto_init() {
	if File_durable_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_durable_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RocksDBInfo); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_durable_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LevelDBInfo); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_durable_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_durable_proto_goTypes,
		DependencyIndexes: file_durable_proto_depIdxs,
		MessageInfos:      file_durable_proto_msgTypes,
	}.Build()
	File_durable_proto = out.File
	file_durable_proto_rawDesc = nil
	file_durable_proto_goTypes = nil
	file_durable_proto_depIdxs = nil
}
