package information

import (
	"bytes"

	"github.com/mauleyzaola/maupod/src/pkg/pb"
)

func MediaFromRaw(raw string) (*pb.Media, error) {
	m, err := ParseMediaInfo(bytes.NewBufferString(raw))
	if err != nil {
		return nil, err
	}
	return m.ToProto(), nil
}

// MediaInfo represents the output from `mediainfo` program
type MediaInfo struct {
	AlbumPerformer        string
	Album                 string
	AudioCount            int64
	AudioFormatList       string
	BitDepth              int64
	BitDepthString        string
	BitRate               int64
	BitRateMode           string
	Channels              string
	ChannelsLayout        string
	ChannelsPosition      string
	Comment               string
	CommercialName        string
	CompleteName          string
	Composer              string
	Compression           string
	CountOfAudioStreams   int64
	Duration              float64
	EncodedLibraryDate    string
	EncodedLibraryName    string
	EncodedLibraryVersion string
	FileExtension         string
	FileName              string
	FileSize              int64
	FolderName            string
	FormatInfo            string
	Format                string
	FormatURL             string
	Genre                 string
	InternetMediaType     string
	KindOfStream          string
	OverallBitRate        int64
	OverallBitRateMode    string
	Part                  int64
	PartTotal             int64
	Performer             string
	RecordedDate          int64
	SamplesCount          int64
	SamplingRate          int64
	StreamIdentifier      int64
	StreamSize            int64
	Title                 string
	TrackNamePosition     int64
	TrackName             string
	TrackNameTotal        int64
	WritingLibrary        string
}

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
