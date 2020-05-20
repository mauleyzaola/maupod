package media

import (
	"github.com/mauleyzaola/maupod/src/server/pkg/pb"
)

func (m *MediaInfo) ToProto() *pb.Media {
	res := &pb.Media{
		Format:             m.Format,
		FileSize:           m.FileSize,
		Duration:           m.Duration,
		OverallBitRate:     m.OverallBitRate,
		OverallBitRateMode: m.OverallBitRateMode,
		StreamSize:         m.StreamSize,
		Album:              m.Album,
		Title:              m.Title,
		Track:              m.TrackName,
		TrackPosition:      m.TrackNamePosition,
		Performer:          m.Performer,
		Genre:              m.Genre,
		RecordedDate:       m.RecordedDate,
		Comment:            m.Comment,
		Channels:           m.Channels,
		ChannelLayout:      m.ChannelsLayout,
		ChannelPositions:   m.ChannelsPosition,
		SamplingRate:       m.SamplingRate,
		SamplingCount:      m.SamplesCount,
		BitDepth:           m.BitDepth,
		BitRateMode:        m.BitRateMode,
		CompressionMode:    m.Compression,
		EncodedLibrary:     m.EncodedLibraryVersion,
		EncodedLibraryName: m.EncodedLibraryName,
		BitRate:            m.BitRate,
	}

	return res
}

/*
TODO: include these fields in proto.Media
TrackNameTotal        int64
AlbumPerformer        string
AudioCount            int64
BitDepthString        string
CommercialName        string
CompleteName          string
CountOfAudioStreams   int64
EncodedLibraryDate    string
FileExtension         string
FileName              string
FolderName            string
FormatInfo            string
Format                string
FormatURL             string
InternetMediaType     string
KindOfStream          string
Part                  int64
PartTotal             int64
StreamIdentifier      int64
WritingLibrary        string
*/
