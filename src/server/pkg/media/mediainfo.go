package media

import (
	"github.com/mauleyzaola/maupod/src/server/pkg/pb"
)

func (m *MediaInfo) ToProto() *pb.Media {
	res := &pb.Media{
		Format:              m.Format,
		FileSize:            m.FileSize,
		Duration:            m.Duration,
		OverallBitRate:      m.OverallBitRate,
		OverallBitRateMode:  m.OverallBitRateMode,
		StreamSize:          m.StreamSize,
		Album:               m.Album,
		Title:               m.Title,
		Track:               m.TrackName,
		TrackPosition:       m.TrackNamePosition,
		Performer:           m.Performer,
		Genre:               m.Genre,
		RecordedDate:        m.RecordedDate,
		Comment:             m.Comment,
		Channels:            m.Channels,
		ChannelLayout:       m.ChannelsLayout,
		ChannelPositions:    m.ChannelsPosition,
		SamplingRate:        m.SamplingRate,
		SamplingCount:       m.SamplesCount,
		BitDepth:            m.BitDepth,
		BitRateMode:         m.BitRateMode,
		CompressionMode:     m.Compression,
		EncodedLibrary:      m.EncodedLibraryVersion,
		EncodedLibraryName:  m.EncodedLibraryName,
		BitRate:             m.BitRate,
		TrackNameTotal:      m.TrackNameTotal,
		AlbumPerformer:      m.AlbumPerformer,
		AudioCount:          m.AudioCount,
		BitDepthString:      m.BitDepthString,
		CommercialName:      m.CommercialName,
		CompleteName:        m.CompleteName,
		CountOfAudioStreams: m.CountOfAudioStreams,
		EncodedLibraryDate:  m.EncodedLibraryDate,
		FileExtension:       m.FileExtension,
		FileName:            m.FileName,
		FolderName:          m.FolderName,
		FormatInfo:          m.FormatInfo,
		FormatUrl:           m.FormatURL,
		InternetMediaType:   m.InternetMediaType,
		KindOfStream:        m.KindOfStream,
		Part:                m.Part,
		PartTotal:           m.PartTotal,
		StreamIdentifier:    m.StreamIdentifier,
		WritingLibrary:      m.WritingLibrary,
	}

	return res
}
