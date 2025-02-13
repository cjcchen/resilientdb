// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.6.1
// source: kv_server.proto

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

type KVRequest_CMD int32

const (
	KVRequest_NONE      KVRequest_CMD = 0
	KVRequest_SET       KVRequest_CMD = 1
	KVRequest_GET       KVRequest_CMD = 2
	KVRequest_GETVALUES KVRequest_CMD = 3
	KVRequest_GETRANGE  KVRequest_CMD = 4
)

// Enum value maps for KVRequest_CMD.
var (
	KVRequest_CMD_name = map[int32]string{
		0: "NONE",
		1: "SET",
		2: "GET",
		3: "GETVALUES",
		4: "GETRANGE",
	}
	KVRequest_CMD_value = map[string]int32{
		"NONE":      0,
		"SET":       1,
		"GET":       2,
		"GETVALUES": 3,
		"GETRANGE":  4,
	}
)

func (x KVRequest_CMD) Enum() *KVRequest_CMD {
	p := new(KVRequest_CMD)
	*p = x
	return p
}

func (x KVRequest_CMD) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (KVRequest_CMD) Descriptor() protoreflect.EnumDescriptor {
	return file_kv_server_proto_enumTypes[0].Descriptor()
}

func (KVRequest_CMD) Type() protoreflect.EnumType {
	return &file_kv_server_proto_enumTypes[0]
}

func (x KVRequest_CMD) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use KVRequest_CMD.Descriptor instead.
func (KVRequest_CMD) EnumDescriptor() ([]byte, []int) {
	return file_kv_server_proto_rawDescGZIP(), []int{0, 0}
}

type KVRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Cmd   KVRequest_CMD `protobuf:"varint,1,opt,name=cmd,proto3,enum=XXXX.KVRequest_CMD" json:"cmd,omitempty"`
	Key   string        `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
	Value []byte        `protobuf:"bytes,3,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *KVRequest) Reset() {
	*x = KVRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kv_server_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KVRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KVRequest) ProtoMessage() {}

func (x *KVRequest) ProtoReflect() protoreflect.Message {
	mi := &file_kv_server_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KVRequest.ProtoReflect.Descriptor instead.
func (*KVRequest) Descriptor() ([]byte, []int) {
	return file_kv_server_proto_rawDescGZIP(), []int{0}
}

func (x *KVRequest) GetCmd() KVRequest_CMD {
	if x != nil {
		return x.Cmd
	}
	return KVRequest_NONE
}

func (x *KVRequest) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *KVRequest) GetValue() []byte {
	if x != nil {
		return x.Value
	}
	return nil
}

type KVResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key   string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value []byte `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *KVResponse) Reset() {
	*x = KVResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kv_server_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KVResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KVResponse) ProtoMessage() {}

func (x *KVResponse) ProtoReflect() protoreflect.Message {
	mi := &file_kv_server_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KVResponse.ProtoReflect.Descriptor instead.
func (*KVResponse) Descriptor() ([]byte, []int) {
	return file_kv_server_proto_rawDescGZIP(), []int{1}
}

func (x *KVResponse) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *KVResponse) GetValue() []byte {
	if x != nil {
		return x.Value
	}
	return nil
}

var File_kv_server_proto protoreflect.FileDescriptor

var file_kv_server_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x6b, 0x76, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x05, 0x72, 0x65, 0x73, 0x64, 0x62, 0x22, 0x9b, 0x01, 0x0a, 0x09, 0x4b, 0x56, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x26, 0x0a, 0x03, 0x63, 0x6d, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x14, 0x2e, 0x72, 0x65, 0x73, 0x64, 0x62, 0x2e, 0x4b, 0x56, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x43, 0x4d, 0x44, 0x52, 0x03, 0x63, 0x6d, 0x64, 0x12, 0x10,
	0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79,
	0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x3e, 0x0a, 0x03, 0x43, 0x4d, 0x44, 0x12, 0x08, 0x0a,
	0x04, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00, 0x12, 0x07, 0x0a, 0x03, 0x53, 0x45, 0x54, 0x10, 0x01,
	0x12, 0x07, 0x0a, 0x03, 0x47, 0x45, 0x54, 0x10, 0x02, 0x12, 0x0d, 0x0a, 0x09, 0x47, 0x45, 0x54,
	0x56, 0x41, 0x4c, 0x55, 0x45, 0x53, 0x10, 0x03, 0x12, 0x0c, 0x0a, 0x08, 0x47, 0x45, 0x54, 0x52,
	0x41, 0x4e, 0x47, 0x45, 0x10, 0x04, 0x22, 0x34, 0x0a, 0x0a, 0x4b, 0x56, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x0a, 0x5a, 0x08,
	0x2e, 0x2f, 0x3b, 0x72, 0x65, 0x73, 0x64, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_kv_server_proto_rawDescOnce sync.Once
	file_kv_server_proto_rawDescData = file_kv_server_proto_rawDesc
)

func file_kv_server_proto_rawDescGZIP() []byte {
	file_kv_server_proto_rawDescOnce.Do(func() {
		file_kv_server_proto_rawDescData = protoimpl.X.CompressGZIP(file_kv_server_proto_rawDescData)
	})
	return file_kv_server_proto_rawDescData
}

var file_kv_server_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_kv_server_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_kv_server_proto_goTypes = []interface{}{
	(KVRequest_CMD)(0), // 0: XXXX.KVRequest.CMD
	(*KVRequest)(nil),  // 1: XXXX.KVRequest
	(*KVResponse)(nil), // 2: XXXX.KVResponse
}
var file_kv_server_proto_depIdxs = []int32{
	0, // 0: XXXX.KVRequest.cmd:type_name -> XXXX.KVRequest.CMD
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_kv_server_proto_init() }
func file_kv_server_proto_init() {
	if File_kv_server_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_kv_server_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KVRequest); i {
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
		file_kv_server_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KVResponse); i {
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
			RawDescriptor: file_kv_server_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_kv_server_proto_goTypes,
		DependencyIndexes: file_kv_server_proto_depIdxs,
		EnumInfos:         file_kv_server_proto_enumTypes,
		MessageInfos:      file_kv_server_proto_msgTypes,
	}.Build()
	File_kv_server_proto = out.File
	file_kv_server_proto_rawDesc = nil
	file_kv_server_proto_goTypes = nil
	file_kv_server_proto_depIdxs = nil
}
