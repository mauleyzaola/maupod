// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.12.3
// source: media.proto

package pb

import (
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

// Media is a direct map between a mediainfo result and additional fields
type Media struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// uuid
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// the sha 256 to uniquely identify a file (dup cleanup)
	Sha string `protobuf:"bytes,2,opt,name=sha,proto3" json:"sha,omitempty"`
	// the full path to the file on the file store
	Location string `protobuf:"bytes,3,opt,name=location,proto3" json:"location,omitempty"`
	// date the media file was scanned
	LastScan              *timestamp.Timestamp `protobuf:"bytes,4,opt,name=last_scan,json=lastScan,proto3" json:"last_scan,omitempty"`
	FileExtension         string               `protobuf:"bytes,6,opt,name=file_extension,json=fileExtension,proto3" json:"file_extension,omitempty"`
	Format                string               `protobuf:"bytes,7,opt,name=format,proto3" json:"format,omitempty"`
	FileSize              int64                `protobuf:"varint,8,opt,name=file_size,json=fileSize,proto3" json:"file_size,omitempty"`
	Duration              float64              `protobuf:"fixed64,9,opt,name=duration,proto3" json:"duration,omitempty"`
	OverallBitRateMode    string               `protobuf:"bytes,10,opt,name=overall_bit_rate_mode,json=overallBitRateMode,proto3" json:"overall_bit_rate_mode,omitempty"`
	OverallBitRate        int64                `protobuf:"varint,11,opt,name=overall_bit_rate,json=overallBitRate,proto3" json:"overall_bit_rate,omitempty"`
	StreamSize            int64                `protobuf:"varint,12,opt,name=stream_size,json=streamSize,proto3" json:"stream_size,omitempty"`
	Album                 string               `protobuf:"bytes,13,opt,name=album,proto3" json:"album,omitempty"`
	Title                 string               `protobuf:"bytes,14,opt,name=title,proto3" json:"title,omitempty"`
	Track                 string               `protobuf:"bytes,15,opt,name=track,proto3" json:"track,omitempty"`
	TrackPosition         int64                `protobuf:"varint,16,opt,name=track_position,json=trackPosition,proto3" json:"track_position,omitempty"`
	Performer             string               `protobuf:"bytes,17,opt,name=performer,proto3" json:"performer,omitempty"`
	Genre                 string               `protobuf:"bytes,18,opt,name=genre,proto3" json:"genre,omitempty"`
	RecordedDate          int64                `protobuf:"varint,19,opt,name=recorded_date,json=recordedDate,proto3" json:"recorded_date,omitempty"`
	Comment               string               `protobuf:"bytes,21,opt,name=comment,proto3" json:"comment,omitempty"`
	Channels              string               `protobuf:"bytes,22,opt,name=channels,proto3" json:"channels,omitempty"`
	ChannelPositions      string               `protobuf:"bytes,23,opt,name=channel_positions,json=channelPositions,proto3" json:"channel_positions,omitempty"`
	ChannelLayout         string               `protobuf:"bytes,24,opt,name=channel_layout,json=channelLayout,proto3" json:"channel_layout,omitempty"`
	SamplingRate          int64                `protobuf:"varint,25,opt,name=sampling_rate,json=samplingRate,proto3" json:"sampling_rate,omitempty"`
	SamplingCount         int64                `protobuf:"varint,26,opt,name=sampling_count,json=samplingCount,proto3" json:"sampling_count,omitempty"`
	BitDepth              int64                `protobuf:"varint,27,opt,name=bit_depth,json=bitDepth,proto3" json:"bit_depth,omitempty"`
	CompressionMode       string               `protobuf:"bytes,28,opt,name=compression_mode,json=compressionMode,proto3" json:"compression_mode,omitempty"`
	EncodedLibraryName    string               `protobuf:"bytes,30,opt,name=encoded_library_name,json=encodedLibraryName,proto3" json:"encoded_library_name,omitempty"`
	EncodedLibraryVersion string               `protobuf:"bytes,31,opt,name=encoded_library_version,json=encodedLibraryVersion,proto3" json:"encoded_library_version,omitempty"`
	BitRateMode           string               `protobuf:"bytes,32,opt,name=bit_rate_mode,json=bitRateMode,proto3" json:"bit_rate_mode,omitempty"`
	BitRate               int64                `protobuf:"varint,33,opt,name=bit_rate,json=bitRate,proto3" json:"bit_rate,omitempty"`
	TrackNameTotal        int64                `protobuf:"varint,34,opt,name=track_name_total,json=trackNameTotal,proto3" json:"track_name_total,omitempty"`
	AlbumPerformer        string               `protobuf:"bytes,35,opt,name=album_performer,json=albumPerformer,proto3" json:"album_performer,omitempty"`
	AudioCount            int64                `protobuf:"varint,36,opt,name=audio_count,json=audioCount,proto3" json:"audio_count,omitempty"`
	BitDepthString        string               `protobuf:"bytes,37,opt,name=bit_depth_string,json=bitDepthString,proto3" json:"bit_depth_string,omitempty"`
	CommercialName        string               `protobuf:"bytes,38,opt,name=commercial_name,json=commercialName,proto3" json:"commercial_name,omitempty"`
	CompleteName          string               `protobuf:"bytes,39,opt,name=complete_name,json=completeName,proto3" json:"complete_name,omitempty"`
	CountOfAudioStreams   int64                `protobuf:"varint,40,opt,name=count_of_audio_streams,json=countOfAudioStreams,proto3" json:"count_of_audio_streams,omitempty"`
	EncodedLibraryDate    string               `protobuf:"bytes,41,opt,name=encoded_library_date,json=encodedLibraryDate,proto3" json:"encoded_library_date,omitempty"`
	FileName              string               `protobuf:"bytes,42,opt,name=file_name,json=fileName,proto3" json:"file_name,omitempty"`
	FolderName            string               `protobuf:"bytes,43,opt,name=folder_name,json=folderName,proto3" json:"folder_name,omitempty"`
	FormatInfo            string               `protobuf:"bytes,44,opt,name=format_info,json=formatInfo,proto3" json:"format_info,omitempty"`
	FormatUrl             string               `protobuf:"bytes,45,opt,name=format_url,json=formatUrl,proto3" json:"format_url,omitempty"`
	InternetMediaType     string               `protobuf:"bytes,46,opt,name=internet_media_type,json=internetMediaType,proto3" json:"internet_media_type,omitempty"`
	KindOfStream          string               `protobuf:"bytes,47,opt,name=kind_of_stream,json=kindOfStream,proto3" json:"kind_of_stream,omitempty"`
	Part                  int64                `protobuf:"varint,48,opt,name=part,proto3" json:"part,omitempty"`
	PartTotal             int64                `protobuf:"varint,49,opt,name=part_total,json=partTotal,proto3" json:"part_total,omitempty"`
	StreamIdentifier      int64                `protobuf:"varint,50,opt,name=stream_identifier,json=streamIdentifier,proto3" json:"stream_identifier,omitempty"`
	WritingLibrary        string               `protobuf:"bytes,51,opt,name=writing_library,json=writingLibrary,proto3" json:"writing_library,omitempty"`
	ModifiedDate          *timestamp.Timestamp `protobuf:"bytes,52,opt,name=modified_date,json=modifiedDate,proto3" json:"modified_date,omitempty"`
	Composer              string               `protobuf:"bytes,53,opt,name=composer,proto3" json:"composer,omitempty"`
	LastImageScan         *timestamp.Timestamp `protobuf:"bytes,55,opt,name=last_image_scan,json=lastImageScan,proto3" json:"last_image_scan,omitempty"`
	ImageLocation         string               `protobuf:"bytes,56,opt,name=image_location,json=imageLocation,proto3" json:"image_location,omitempty"`
	AlbumIdentifier       string               `protobuf:"bytes,57,opt,name=album_identifier,json=albumIdentifier,proto3" json:"album_identifier,omitempty"`
	IsCompilation         bool                 `protobuf:"varint,58,opt,name=is_compilation,json=isCompilation,proto3" json:"is_compilation,omitempty"`
}

func (x *Media) Reset() {
	*x = Media{}
	if protoimpl.UnsafeEnabled {
		mi := &file_media_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Media) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Media) ProtoMessage() {}

func (x *Media) ProtoReflect() protoreflect.Message {
	mi := &file_media_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Media.ProtoReflect.Descriptor instead.
func (*Media) Descriptor() ([]byte, []int) {
	return file_media_proto_rawDescGZIP(), []int{0}
}

func (x *Media) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Media) GetSha() string {
	if x != nil {
		return x.Sha
	}
	return ""
}

func (x *Media) GetLocation() string {
	if x != nil {
		return x.Location
	}
	return ""
}

func (x *Media) GetLastScan() *timestamp.Timestamp {
	if x != nil {
		return x.LastScan
	}
	return nil
}

func (x *Media) GetFileExtension() string {
	if x != nil {
		return x.FileExtension
	}
	return ""
}

func (x *Media) GetFormat() string {
	if x != nil {
		return x.Format
	}
	return ""
}

func (x *Media) GetFileSize() int64 {
	if x != nil {
		return x.FileSize
	}
	return 0
}

func (x *Media) GetDuration() float64 {
	if x != nil {
		return x.Duration
	}
	return 0
}

func (x *Media) GetOverallBitRateMode() string {
	if x != nil {
		return x.OverallBitRateMode
	}
	return ""
}

func (x *Media) GetOverallBitRate() int64 {
	if x != nil {
		return x.OverallBitRate
	}
	return 0
}

func (x *Media) GetStreamSize() int64 {
	if x != nil {
		return x.StreamSize
	}
	return 0
}

func (x *Media) GetAlbum() string {
	if x != nil {
		return x.Album
	}
	return ""
}

func (x *Media) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Media) GetTrack() string {
	if x != nil {
		return x.Track
	}
	return ""
}

func (x *Media) GetTrackPosition() int64 {
	if x != nil {
		return x.TrackPosition
	}
	return 0
}

func (x *Media) GetPerformer() string {
	if x != nil {
		return x.Performer
	}
	return ""
}

func (x *Media) GetGenre() string {
	if x != nil {
		return x.Genre
	}
	return ""
}

func (x *Media) GetRecordedDate() int64 {
	if x != nil {
		return x.RecordedDate
	}
	return 0
}

func (x *Media) GetComment() string {
	if x != nil {
		return x.Comment
	}
	return ""
}

func (x *Media) GetChannels() string {
	if x != nil {
		return x.Channels
	}
	return ""
}

func (x *Media) GetChannelPositions() string {
	if x != nil {
		return x.ChannelPositions
	}
	return ""
}

func (x *Media) GetChannelLayout() string {
	if x != nil {
		return x.ChannelLayout
	}
	return ""
}

func (x *Media) GetSamplingRate() int64 {
	if x != nil {
		return x.SamplingRate
	}
	return 0
}

func (x *Media) GetSamplingCount() int64 {
	if x != nil {
		return x.SamplingCount
	}
	return 0
}

func (x *Media) GetBitDepth() int64 {
	if x != nil {
		return x.BitDepth
	}
	return 0
}

func (x *Media) GetCompressionMode() string {
	if x != nil {
		return x.CompressionMode
	}
	return ""
}

func (x *Media) GetEncodedLibraryName() string {
	if x != nil {
		return x.EncodedLibraryName
	}
	return ""
}

func (x *Media) GetEncodedLibraryVersion() string {
	if x != nil {
		return x.EncodedLibraryVersion
	}
	return ""
}

func (x *Media) GetBitRateMode() string {
	if x != nil {
		return x.BitRateMode
	}
	return ""
}

func (x *Media) GetBitRate() int64 {
	if x != nil {
		return x.BitRate
	}
	return 0
}

func (x *Media) GetTrackNameTotal() int64 {
	if x != nil {
		return x.TrackNameTotal
	}
	return 0
}

func (x *Media) GetAlbumPerformer() string {
	if x != nil {
		return x.AlbumPerformer
	}
	return ""
}

func (x *Media) GetAudioCount() int64 {
	if x != nil {
		return x.AudioCount
	}
	return 0
}

func (x *Media) GetBitDepthString() string {
	if x != nil {
		return x.BitDepthString
	}
	return ""
}

func (x *Media) GetCommercialName() string {
	if x != nil {
		return x.CommercialName
	}
	return ""
}

func (x *Media) GetCompleteName() string {
	if x != nil {
		return x.CompleteName
	}
	return ""
}

func (x *Media) GetCountOfAudioStreams() int64 {
	if x != nil {
		return x.CountOfAudioStreams
	}
	return 0
}

func (x *Media) GetEncodedLibraryDate() string {
	if x != nil {
		return x.EncodedLibraryDate
	}
	return ""
}

func (x *Media) GetFileName() string {
	if x != nil {
		return x.FileName
	}
	return ""
}

func (x *Media) GetFolderName() string {
	if x != nil {
		return x.FolderName
	}
	return ""
}

func (x *Media) GetFormatInfo() string {
	if x != nil {
		return x.FormatInfo
	}
	return ""
}

func (x *Media) GetFormatUrl() string {
	if x != nil {
		return x.FormatUrl
	}
	return ""
}

func (x *Media) GetInternetMediaType() string {
	if x != nil {
		return x.InternetMediaType
	}
	return ""
}

func (x *Media) GetKindOfStream() string {
	if x != nil {
		return x.KindOfStream
	}
	return ""
}

func (x *Media) GetPart() int64 {
	if x != nil {
		return x.Part
	}
	return 0
}

func (x *Media) GetPartTotal() int64 {
	if x != nil {
		return x.PartTotal
	}
	return 0
}

func (x *Media) GetStreamIdentifier() int64 {
	if x != nil {
		return x.StreamIdentifier
	}
	return 0
}

func (x *Media) GetWritingLibrary() string {
	if x != nil {
		return x.WritingLibrary
	}
	return ""
}

func (x *Media) GetModifiedDate() *timestamp.Timestamp {
	if x != nil {
		return x.ModifiedDate
	}
	return nil
}

func (x *Media) GetComposer() string {
	if x != nil {
		return x.Composer
	}
	return ""
}

func (x *Media) GetLastImageScan() *timestamp.Timestamp {
	if x != nil {
		return x.LastImageScan
	}
	return nil
}

func (x *Media) GetImageLocation() string {
	if x != nil {
		return x.ImageLocation
	}
	return ""
}

func (x *Media) GetAlbumIdentifier() string {
	if x != nil {
		return x.AlbumIdentifier
	}
	return ""
}

func (x *Media) GetIsCompilation() bool {
	if x != nil {
		return x.IsCompilation
	}
	return false
}

var File_media_proto protoreflect.FileDescriptor

var file_media_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70,
	0x62, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0xc7, 0x0f, 0x0a, 0x05, 0x4d, 0x65, 0x64, 0x69, 0x61, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x10, 0x0a, 0x03,
	0x73, 0x68, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x73, 0x68, 0x61, 0x12, 0x1a,
	0x0a, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x37, 0x0a, 0x09, 0x6c, 0x61,
	0x73, 0x74, 0x5f, 0x73, 0x63, 0x61, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x53,
	0x63, 0x61, 0x6e, 0x12, 0x25, 0x0a, 0x0e, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x65, 0x78, 0x74, 0x65,
	0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x66, 0x69, 0x6c,
	0x65, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x66, 0x6f,
	0x72, 0x6d, 0x61, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x66, 0x6f, 0x72, 0x6d,
	0x61, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12,
	0x1a, 0x0a, 0x08, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x09, 0x20, 0x01, 0x28,
	0x01, 0x52, 0x08, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x31, 0x0a, 0x15, 0x6f,
	0x76, 0x65, 0x72, 0x61, 0x6c, 0x6c, 0x5f, 0x62, 0x69, 0x74, 0x5f, 0x72, 0x61, 0x74, 0x65, 0x5f,
	0x6d, 0x6f, 0x64, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x12, 0x6f, 0x76, 0x65, 0x72,
	0x61, 0x6c, 0x6c, 0x42, 0x69, 0x74, 0x52, 0x61, 0x74, 0x65, 0x4d, 0x6f, 0x64, 0x65, 0x12, 0x28,
	0x0a, 0x10, 0x6f, 0x76, 0x65, 0x72, 0x61, 0x6c, 0x6c, 0x5f, 0x62, 0x69, 0x74, 0x5f, 0x72, 0x61,
	0x74, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0e, 0x6f, 0x76, 0x65, 0x72, 0x61, 0x6c,
	0x6c, 0x42, 0x69, 0x74, 0x52, 0x61, 0x74, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x74, 0x72, 0x65,
	0x61, 0x6d, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x73,
	0x74, 0x72, 0x65, 0x61, 0x6d, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x61, 0x6c, 0x62,
	0x75, 0x6d, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x61, 0x6c, 0x62, 0x75, 0x6d, 0x12,
	0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x18, 0x0f,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x12, 0x25, 0x0a, 0x0e, 0x74,
	0x72, 0x61, 0x63, 0x6b, 0x5f, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x10, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x0d, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x65, 0x72, 0x66, 0x6f, 0x72, 0x6d, 0x65, 0x72, 0x18,
	0x11, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x65, 0x72, 0x66, 0x6f, 0x72, 0x6d, 0x65, 0x72,
	0x12, 0x14, 0x0a, 0x05, 0x67, 0x65, 0x6e, 0x72, 0x65, 0x18, 0x12, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x67, 0x65, 0x6e, 0x72, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64,
	0x65, 0x64, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x13, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x72,
	0x65, 0x63, 0x6f, 0x72, 0x64, 0x65, 0x64, 0x44, 0x61, 0x74, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63,
	0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x15, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f,
	0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c,
	0x73, 0x18, 0x16, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c,
	0x73, 0x12, 0x2b, 0x0a, 0x11, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x5f, 0x70, 0x6f, 0x73,
	0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x17, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x63, 0x68,
	0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x25,
	0x0a, 0x0e, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x5f, 0x6c, 0x61, 0x79, 0x6f, 0x75, 0x74,
	0x18, 0x18, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x4c,
	0x61, 0x79, 0x6f, 0x75, 0x74, 0x12, 0x23, 0x0a, 0x0d, 0x73, 0x61, 0x6d, 0x70, 0x6c, 0x69, 0x6e,
	0x67, 0x5f, 0x72, 0x61, 0x74, 0x65, 0x18, 0x19, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x73, 0x61,
	0x6d, 0x70, 0x6c, 0x69, 0x6e, 0x67, 0x52, 0x61, 0x74, 0x65, 0x12, 0x25, 0x0a, 0x0e, 0x73, 0x61,
	0x6d, 0x70, 0x6c, 0x69, 0x6e, 0x67, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x1a, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x0d, 0x73, 0x61, 0x6d, 0x70, 0x6c, 0x69, 0x6e, 0x67, 0x43, 0x6f, 0x75, 0x6e,
	0x74, 0x12, 0x1b, 0x0a, 0x09, 0x62, 0x69, 0x74, 0x5f, 0x64, 0x65, 0x70, 0x74, 0x68, 0x18, 0x1b,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x62, 0x69, 0x74, 0x44, 0x65, 0x70, 0x74, 0x68, 0x12, 0x29,
	0x0a, 0x10, 0x63, 0x6f, 0x6d, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x6d, 0x6f,
	0x64, 0x65, 0x18, 0x1c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x63, 0x6f, 0x6d, 0x70, 0x72, 0x65,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x4d, 0x6f, 0x64, 0x65, 0x12, 0x30, 0x0a, 0x14, 0x65, 0x6e, 0x63,
	0x6f, 0x64, 0x65, 0x64, 0x5f, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x79, 0x5f, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x12, 0x65, 0x6e, 0x63, 0x6f, 0x64, 0x65, 0x64,
	0x4c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x36, 0x0a, 0x17, 0x65,
	0x6e, 0x63, 0x6f, 0x64, 0x65, 0x64, 0x5f, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x79, 0x5f, 0x76,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x1f, 0x20, 0x01, 0x28, 0x09, 0x52, 0x15, 0x65, 0x6e,
	0x63, 0x6f, 0x64, 0x65, 0x64, 0x4c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x79, 0x56, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x12, 0x22, 0x0a, 0x0d, 0x62, 0x69, 0x74, 0x5f, 0x72, 0x61, 0x74, 0x65, 0x5f,
	0x6d, 0x6f, 0x64, 0x65, 0x18, 0x20, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x62, 0x69, 0x74, 0x52,
	0x61, 0x74, 0x65, 0x4d, 0x6f, 0x64, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x62, 0x69, 0x74, 0x5f, 0x72,
	0x61, 0x74, 0x65, 0x18, 0x21, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x62, 0x69, 0x74, 0x52, 0x61,
	0x74, 0x65, 0x12, 0x28, 0x0a, 0x10, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x5f, 0x6e, 0x61, 0x6d, 0x65,
	0x5f, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x22, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0e, 0x74, 0x72,
	0x61, 0x63, 0x6b, 0x4e, 0x61, 0x6d, 0x65, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x12, 0x27, 0x0a, 0x0f,
	0x61, 0x6c, 0x62, 0x75, 0x6d, 0x5f, 0x70, 0x65, 0x72, 0x66, 0x6f, 0x72, 0x6d, 0x65, 0x72, 0x18,
	0x23, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x61, 0x6c, 0x62, 0x75, 0x6d, 0x50, 0x65, 0x72, 0x66,
	0x6f, 0x72, 0x6d, 0x65, 0x72, 0x12, 0x1f, 0x0a, 0x0b, 0x61, 0x75, 0x64, 0x69, 0x6f, 0x5f, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x18, 0x24, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x61, 0x75, 0x64, 0x69,
	0x6f, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x28, 0x0a, 0x10, 0x62, 0x69, 0x74, 0x5f, 0x64, 0x65,
	0x70, 0x74, 0x68, 0x5f, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x18, 0x25, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0e, 0x62, 0x69, 0x74, 0x44, 0x65, 0x70, 0x74, 0x68, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67,
	0x12, 0x27, 0x0a, 0x0f, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x72, 0x63, 0x69, 0x61, 0x6c, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x26, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x63, 0x6f, 0x6d, 0x6d, 0x65,
	0x72, 0x63, 0x69, 0x61, 0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x63, 0x6f, 0x6d,
	0x70, 0x6c, 0x65, 0x74, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x27, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0c, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x33,
	0x0a, 0x16, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x6f, 0x66, 0x5f, 0x61, 0x75, 0x64, 0x69, 0x6f,
	0x5f, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x73, 0x18, 0x28, 0x20, 0x01, 0x28, 0x03, 0x52, 0x13,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x4f, 0x66, 0x41, 0x75, 0x64, 0x69, 0x6f, 0x53, 0x74, 0x72, 0x65,
	0x61, 0x6d, 0x73, 0x12, 0x30, 0x0a, 0x14, 0x65, 0x6e, 0x63, 0x6f, 0x64, 0x65, 0x64, 0x5f, 0x6c,
	0x69, 0x62, 0x72, 0x61, 0x72, 0x79, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x29, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x12, 0x65, 0x6e, 0x63, 0x6f, 0x64, 0x65, 0x64, 0x4c, 0x69, 0x62, 0x72, 0x61, 0x72,
	0x79, 0x44, 0x61, 0x74, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x2a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x66, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x5f, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x2b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x66, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x5f, 0x69, 0x6e,
	0x66, 0x6f, 0x18, 0x2c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74,
	0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1d, 0x0a, 0x0a, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x5f, 0x75,
	0x72, 0x6c, 0x18, 0x2d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74,
	0x55, 0x72, 0x6c, 0x12, 0x2e, 0x0a, 0x13, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x5f,
	0x6d, 0x65, 0x64, 0x69, 0x61, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x2e, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x11, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x4d, 0x65, 0x64, 0x69, 0x61, 0x54,
	0x79, 0x70, 0x65, 0x12, 0x24, 0x0a, 0x0e, 0x6b, 0x69, 0x6e, 0x64, 0x5f, 0x6f, 0x66, 0x5f, 0x73,
	0x74, 0x72, 0x65, 0x61, 0x6d, 0x18, 0x2f, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x6b, 0x69, 0x6e,
	0x64, 0x4f, 0x66, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x72,
	0x74, 0x18, 0x30, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x70, 0x61, 0x72, 0x74, 0x12, 0x1d, 0x0a,
	0x0a, 0x70, 0x61, 0x72, 0x74, 0x5f, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x31, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x09, 0x70, 0x61, 0x72, 0x74, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x12, 0x2b, 0x0a, 0x11,
	0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x5f, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65,
	0x72, 0x18, 0x32, 0x20, 0x01, 0x28, 0x03, 0x52, 0x10, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x49,
	0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x12, 0x27, 0x0a, 0x0f, 0x77, 0x72, 0x69,
	0x74, 0x69, 0x6e, 0x67, 0x5f, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x79, 0x18, 0x33, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0e, 0x77, 0x72, 0x69, 0x74, 0x69, 0x6e, 0x67, 0x4c, 0x69, 0x62, 0x72, 0x61,
	0x72, 0x79, 0x12, 0x3f, 0x0a, 0x0d, 0x6d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x64, 0x5f, 0x64,
	0x61, 0x74, 0x65, 0x18, 0x34, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0c, 0x6d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x64, 0x44,
	0x61, 0x74, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x73, 0x65, 0x72, 0x18,
	0x35, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x73, 0x65, 0x72, 0x12,
	0x42, 0x0a, 0x0f, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x63,
	0x61, 0x6e, 0x18, 0x37, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x52, 0x0d, 0x6c, 0x61, 0x73, 0x74, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x53,
	0x63, 0x61, 0x6e, 0x12, 0x25, 0x0a, 0x0e, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x5f, 0x6c, 0x6f, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x38, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x69, 0x6d, 0x61,
	0x67, 0x65, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x29, 0x0a, 0x10, 0x61, 0x6c,
	0x62, 0x75, 0x6d, 0x5f, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x18, 0x39,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x61, 0x6c, 0x62, 0x75, 0x6d, 0x49, 0x64, 0x65, 0x6e, 0x74,
	0x69, 0x66, 0x69, 0x65, 0x72, 0x12, 0x25, 0x0a, 0x0e, 0x69, 0x73, 0x5f, 0x63, 0x6f, 0x6d, 0x70,
	0x69, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x3a, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0d, 0x69,
	0x73, 0x43, 0x6f, 0x6d, 0x70, 0x69, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x2d, 0x5a, 0x2b,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6d, 0x61, 0x75, 0x6c, 0x65,
	0x79, 0x7a, 0x61, 0x6f, 0x6c, 0x61, 0x2f, 0x6d, 0x61, 0x75, 0x70, 0x6f, 0x64, 0x2f, 0x73, 0x72,
	0x63, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_media_proto_rawDescOnce sync.Once
	file_media_proto_rawDescData = file_media_proto_rawDesc
)

func file_media_proto_rawDescGZIP() []byte {
	file_media_proto_rawDescOnce.Do(func() {
		file_media_proto_rawDescData = protoimpl.X.CompressGZIP(file_media_proto_rawDescData)
	})
	return file_media_proto_rawDescData
}

var file_media_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_media_proto_goTypes = []interface{}{
	(*Media)(nil),               // 0: pb.Media
	(*timestamp.Timestamp)(nil), // 1: google.protobuf.Timestamp
}
var file_media_proto_depIdxs = []int32{
	1, // 0: pb.Media.last_scan:type_name -> google.protobuf.Timestamp
	1, // 1: pb.Media.modified_date:type_name -> google.protobuf.Timestamp
	1, // 2: pb.Media.last_image_scan:type_name -> google.protobuf.Timestamp
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_media_proto_init() }
func file_media_proto_init() {
	if File_media_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_media_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Media); i {
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
			RawDescriptor: file_media_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_media_proto_goTypes,
		DependencyIndexes: file_media_proto_depIdxs,
		MessageInfos:      file_media_proto_msgTypes,
	}.Build()
	File_media_proto = out.File
	file_media_proto_rawDesc = nil
	file_media_proto_goTypes = nil
	file_media_proto_depIdxs = nil
}
