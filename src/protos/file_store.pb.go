// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.13.0
// source: file_store.proto

package protos

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type FileStore_StoreType int32

const (
	FileStore_FILE_SYSTEM FileStore_StoreType = 0
	FileStore_S3          FileStore_StoreType = 1
)

// Enum value maps for FileStore_StoreType.
var (
	FileStore_StoreType_name = map[int32]string{
		0: "FILE_SYSTEM",
		1: "S3",
	}
	FileStore_StoreType_value = map[string]int32{
		"FILE_SYSTEM": 0,
		"S3":          1,
	}
)

func (x FileStore_StoreType) Enum() *FileStore_StoreType {
	p := new(FileStore_StoreType)
	*p = x
	return p
}

func (x FileStore_StoreType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (FileStore_StoreType) Descriptor() protoreflect.EnumDescriptor {
	return file_file_store_proto_enumTypes[0].Descriptor()
}

func (FileStore_StoreType) Type() protoreflect.EnumType {
	return &file_file_store_proto_enumTypes[0]
}

func (x FileStore_StoreType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use FileStore_StoreType.Descriptor instead.
func (FileStore_StoreType) EnumDescriptor() ([]byte, []int) {
	return file_file_store_proto_rawDescGZIP(), []int{0, 0}
}

type FileStore struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name     string              `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Type     FileStore_StoreType `protobuf:"varint,2,opt,name=type,proto3,enum=protos.FileStore_StoreType" json:"type,omitempty"`
	Location string              `protobuf:"bytes,3,opt,name=location,proto3" json:"location,omitempty"`
}

func (x *FileStore) Reset() {
	*x = FileStore{}
	if protoimpl.UnsafeEnabled {
		mi := &file_file_store_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FileStore) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileStore) ProtoMessage() {}

func (x *FileStore) ProtoReflect() protoreflect.Message {
	mi := &file_file_store_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FileStore.ProtoReflect.Descriptor instead.
func (*FileStore) Descriptor() ([]byte, []int) {
	return file_file_store_proto_rawDescGZIP(), []int{0}
}

func (x *FileStore) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *FileStore) GetType() FileStore_StoreType {
	if x != nil {
		return x.Type
	}
	return FileStore_FILE_SYSTEM
}

func (x *FileStore) GetLocation() string {
	if x != nil {
		return x.Location
	}
	return ""
}

var File_file_store_proto protoreflect.FileDescriptor

var file_file_store_proto_rawDesc = []byte{
	0x0a, 0x10, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x22, 0x92, 0x01, 0x0a, 0x09, 0x46,
	0x69, 0x6c, 0x65, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x2f, 0x0a, 0x04,
	0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1b, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x73, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x53, 0x74,
	0x6f, 0x72, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x1a, 0x0a,
	0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x24, 0x0a, 0x09, 0x53, 0x74, 0x6f,
	0x72, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0f, 0x0a, 0x0b, 0x46, 0x49, 0x4c, 0x45, 0x5f, 0x53,
	0x59, 0x53, 0x54, 0x45, 0x4d, 0x10, 0x00, 0x12, 0x06, 0x0a, 0x02, 0x53, 0x33, 0x10, 0x01, 0x42,
	0x2a, 0x5a, 0x28, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6d, 0x61,
	0x75, 0x6c, 0x65, 0x79, 0x7a, 0x61, 0x6f, 0x6c, 0x61, 0x2f, 0x6d, 0x61, 0x75, 0x70, 0x6f, 0x64,
	0x2f, 0x73, 0x72, 0x63, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_file_store_proto_rawDescOnce sync.Once
	file_file_store_proto_rawDescData = file_file_store_proto_rawDesc
)

func file_file_store_proto_rawDescGZIP() []byte {
	file_file_store_proto_rawDescOnce.Do(func() {
		file_file_store_proto_rawDescData = protoimpl.X.CompressGZIP(file_file_store_proto_rawDescData)
	})
	return file_file_store_proto_rawDescData
}

var file_file_store_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_file_store_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_file_store_proto_goTypes = []interface{}{
	(FileStore_StoreType)(0), // 0: protos.FileStore.StoreType
	(*FileStore)(nil),        // 1: protos.FileStore
}
var file_file_store_proto_depIdxs = []int32{
	0, // 0: protos.FileStore.type:type_name -> protos.FileStore.StoreType
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_file_store_proto_init() }
func file_file_store_proto_init() {
	if File_file_store_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_file_store_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FileStore); i {
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
			RawDescriptor: file_file_store_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_file_store_proto_goTypes,
		DependencyIndexes: file_file_store_proto_depIdxs,
		EnumInfos:         file_file_store_proto_enumTypes,
		MessageInfos:      file_file_store_proto_msgTypes,
	}.Build()
	File_file_store_proto = out.File
	file_file_store_proto_rawDesc = nil
	file_file_store_proto_goTypes = nil
	file_file_store_proto_depIdxs = nil
}
