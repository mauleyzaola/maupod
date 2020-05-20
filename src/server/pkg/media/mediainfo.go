package media

import (
	"github.com/mauleyzaola/maupod/src/server/pkg/pb"
)

func (m *MediaInfo) ToProto() *pb.Media {
	res := &pb.Media{
		FileExtension:         m.FileExtension,
		Format:                m.Format,
		FileSize:              m.FileSize,
		Duration:              m.Duration,
		OverallBitRateMode:    m.OverallBitRateMode,
		OverallBitRate:        m.OverallBitRate,
		StreamSize:            m.StreamSize,
		Album:                 m.Album,
		Title:                 m.Title,
		Track:                 m.TrackName,
		TrackPosition:         m.TrackNamePosition,
		Performer:             m.Performer,
		Genre:                 m.Genre,
		RecordedDate:          m.RecordedDate,
		Comment:               m.Comment,
		Channels:              m.Channels,
		ChannelPositions:      m.ChannelsPosition,
		ChannelLayout:         m.ChannelsLayout,
		SamplingRate:          m.SamplingRate,
		SamplingCount:         m.SamplesCount,
		BitDepth:              m.BitDepth,
		CompressionMode:       m.Compression,
		EncodedLibraryName:    m.EncodedLibraryName,
		EncodedLibraryVersion: m.EncodedLibraryVersion,
		BitRateMode:           m.BitRateMode,
		BitRate:               m.BitRate,
		TrackNameTotal:        m.TrackNameTotal,
		AlbumPerformer:        m.AlbumPerformer,
		AudioCount:            m.AudioCount,
		BitDepthString:        m.BitDepthString,
		CommercialName:        m.CommercialName,
		CompleteName:          m.CompleteName,
		CountOfAudioStreams:   m.CountOfAudioStreams,
		EncodedLibraryDate:    m.EncodedLibraryDate,
		FileName:              m.FileName,
		FolderName:            m.FolderName,
		FormatInfo:            m.FormatInfo,
		FormatUrl:             m.FormatURL,
		InternetMediaType:     m.InternetMediaType,
		KindOfStream:          m.KindOfStream,
		Part:                  m.Part,
		PartTotal:             m.PartTotal,
		StreamIdentifier:      m.StreamIdentifier,
		WritingLibrary:        m.WritingLibrary,
	}

	return res
}
