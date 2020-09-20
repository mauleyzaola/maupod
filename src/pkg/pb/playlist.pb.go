// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.13.0
// source: playlist.proto

package pb

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

type PlayList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *PlayList) Reset() {
	*x = PlayList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_playlist_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PlayList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlayList) ProtoMessage() {}

func (x *PlayList) ProtoReflect() protoreflect.Message {
	mi := &file_playlist_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlayList.ProtoReflect.Descriptor instead.
func (*PlayList) Descriptor() ([]byte, []int) {
	return file_playlist_proto_rawDescGZIP(), []int{0}
}

func (x *PlayList) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *PlayList) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type PlaylistItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       string    `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Playlist *PlayList `protobuf:"bytes,2,opt,name=playlist,proto3" json:"playlist,omitempty"`
	Position int32     `protobuf:"varint,3,opt,name=position,proto3" json:"position,omitempty"`
	Media    *Media    `protobuf:"bytes,4,opt,name=media,proto3" json:"media,omitempty"`
}

func (x *PlaylistItem) Reset() {
	*x = PlaylistItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_playlist_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PlaylistItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlaylistItem) ProtoMessage() {}

func (x *PlaylistItem) ProtoReflect() protoreflect.Message {
	mi := &file_playlist_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlaylistItem.ProtoReflect.Descriptor instead.
func (*PlaylistItem) Descriptor() ([]byte, []int) {
	return file_playlist_proto_rawDescGZIP(), []int{1}
}

func (x *PlaylistItem) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *PlaylistItem) GetPlaylist() *PlayList {
	if x != nil {
		return x.Playlist
	}
	return nil
}

func (x *PlaylistItem) GetPosition() int32 {
	if x != nil {
		return x.Position
	}
	return 0
}

func (x *PlaylistItem) GetMedia() *Media {
	if x != nil {
		return x.Media
	}
	return nil
}

var File_playlist_proto protoreflect.FileDescriptor

var file_playlist_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x70, 0x6c, 0x61, 0x79, 0x6c, 0x69, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x02, 0x70, 0x62, 0x1a, 0x0b, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x2e, 0x0a, 0x08, 0x50, 0x6c, 0x61, 0x79, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x22, 0x85, 0x01, 0x0a, 0x0c, 0x50, 0x6c, 0x61, 0x79, 0x6c, 0x69, 0x73, 0x74, 0x49, 0x74,
	0x65, 0x6d, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x28, 0x0a, 0x08, 0x70, 0x6c, 0x61, 0x79, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x70, 0x62, 0x2e, 0x50, 0x6c, 0x61, 0x79, 0x4c, 0x69,
	0x73, 0x74, 0x52, 0x08, 0x70, 0x6c, 0x61, 0x79, 0x6c, 0x69, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08,
	0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08,
	0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1f, 0x0a, 0x05, 0x6d, 0x65, 0x64, 0x69,
	0x61, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x70, 0x62, 0x2e, 0x4d, 0x65, 0x64,
	0x69, 0x61, 0x52, 0x05, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x42, 0x26, 0x5a, 0x24, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6d, 0x61, 0x75, 0x6c, 0x65, 0x79, 0x7a, 0x61,
	0x6f, 0x6c, 0x61, 0x2f, 0x6d, 0x61, 0x75, 0x70, 0x6f, 0x64, 0x2f, 0x73, 0x72, 0x63, 0x2f, 0x70,
	0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_playlist_proto_rawDescOnce sync.Once
	file_playlist_proto_rawDescData = file_playlist_proto_rawDesc
)

func file_playlist_proto_rawDescGZIP() []byte {
	file_playlist_proto_rawDescOnce.Do(func() {
		file_playlist_proto_rawDescData = protoimpl.X.CompressGZIP(file_playlist_proto_rawDescData)
	})
	return file_playlist_proto_rawDescData
}

var file_playlist_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_playlist_proto_goTypes = []interface{}{
	(*PlayList)(nil),     // 0: pb.PlayList
	(*PlaylistItem)(nil), // 1: pb.PlaylistItem
	(*Media)(nil),        // 2: pb.Media
}
var file_playlist_proto_depIdxs = []int32{
	0, // 0: pb.PlaylistItem.playlist:type_name -> pb.PlayList
	2, // 1: pb.PlaylistItem.media:type_name -> pb.Media
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_playlist_proto_init() }
func file_playlist_proto_init() {
	if File_playlist_proto != nil {
		return
	}
	file_media_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_playlist_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PlayList); i {
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
		file_playlist_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PlaylistItem); i {
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
			RawDescriptor: file_playlist_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_playlist_proto_goTypes,
		DependencyIndexes: file_playlist_proto_depIdxs,
		MessageInfos:      file_playlist_proto_msgTypes,
	}.Build()
	File_playlist_proto = out.File
	file_playlist_proto_rawDesc = nil
	file_playlist_proto_goTypes = nil
	file_playlist_proto_depIdxs = nil
}
