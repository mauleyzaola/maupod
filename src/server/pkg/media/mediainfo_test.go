package media

import (
	"testing"

	"github.com/mauleyzaola/maupod/src/server/pkg/pb"
	"github.com/stretchr/testify/assert"
)

func TestMediaInfo_ToProto(t *testing.T) {
	type fields struct {
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
	tests := []struct {
		name   string
		fields fields
		want   *pb.Media
	}{
		{
			name: "check fields matches",
			fields: fields{
				Album:                 "1",
				AlbumPerformer:        "2",
				AudioCount:            3,
				AudioFormatList:       "4",
				BitDepth:              4,
				BitDepthString:        "5",
				BitRate:               7,
				BitRateMode:           "8",
				Channels:              "9",
				ChannelsLayout:        "8",
				ChannelsPosition:      "11",
				Comment:               "11",
				CommercialName:        "13",
				CompleteName:          "13",
				Compression:           "15",
				CountOfAudioStreams:   16,
				Duration:              17,
				EncodedLibraryDate:    "18",
				EncodedLibraryName:    "19",
				EncodedLibraryVersion: "20",
				FileExtension:         "21",
				FileName:              "22",
				FileSize:              23,
				FolderName:            "24",
				Format:                "25",
				FormatInfo:            "26",
				FormatURL:             "27",
				Genre:                 "28",
				InternetMediaType:     "29",
				KindOfStream:          "30",
				OverallBitRate:        31,
				OverallBitRateMode:    "32",
				Part:                  33,
				PartTotal:             34,
				Performer:             "35",
				RecordedDate:          36,
				SamplesCount:          37,
				SamplingRate:          38,
				StreamIdentifier:      39,
				StreamSize:            40,
				Title:                 "41",
				TrackName:             "42",
				TrackNamePosition:     43,
				TrackNameTotal:        44,
				WritingLibrary:        "45",
			},
			want: &pb.Media{
				Album:                 "1",
				AlbumPerformer:        "2",
				AudioCount:            3,
				BitDepth:              4,
				BitDepthString:        "5",
				BitRate:               7,
				BitRateMode:           "8",
				ChannelLayout:         "8",
				ChannelPositions:      "11",
				Channels:              "9",
				Comment:               "11",
				CommercialName:        "13",
				CompleteName:          "13",
				CompressionMode:       "15",
				CountOfAudioStreams:   16,
				Duration:              17,
				EncodedLibraryDate:    "18",
				EncodedLibraryName:    "19",
				EncodedLibraryVersion: "20",
				FileExtension:         "21",
				FileName:              "22",
				FileSize:              23,
				FolderName:            "24",
				Format:                "25",
				FormatInfo:            "26",
				FormatUrl:             "27",
				Genre:                 "28",
				Id:                    "29",
				InternetMediaType:     "29",
				KindOfStream:          "30",
				LastScan:              nil,
				Location:              "33",
				ModifiedDate:          nil,
				OverallBitRate:        31,
				OverallBitRateMode:    "32",
				Part:                  33,
				PartTotal:             34,
				Performer:             "35",
				RecordedDate:          36,
				SamplingCount:         37,
				SamplingRate:          38,
				Sha:                   "39",
				StreamIdentifier:      39,
				StreamSize:            40,
				Title:                 "41",
				Track:                 "42",
				TrackNameTotal:        44,
				TrackPosition:         43,
				WritingLibrary:        "45",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &MediaInfo{
				AlbumPerformer:        tt.fields.AlbumPerformer,
				Album:                 tt.fields.Album,
				AudioCount:            tt.fields.AudioCount,
				AudioFormatList:       tt.fields.AudioFormatList,
				BitDepth:              tt.fields.BitDepth,
				BitDepthString:        tt.fields.BitDepthString,
				BitRate:               tt.fields.BitRate,
				BitRateMode:           tt.fields.BitRateMode,
				Channels:              tt.fields.Channels,
				ChannelsLayout:        tt.fields.ChannelsLayout,
				ChannelsPosition:      tt.fields.ChannelsPosition,
				Comment:               tt.fields.Comment,
				CommercialName:        tt.fields.CommercialName,
				CompleteName:          tt.fields.CompleteName,
				Compression:           tt.fields.Compression,
				CountOfAudioStreams:   tt.fields.CountOfAudioStreams,
				Duration:              tt.fields.Duration,
				EncodedLibraryDate:    tt.fields.EncodedLibraryDate,
				EncodedLibraryName:    tt.fields.EncodedLibraryName,
				EncodedLibraryVersion: tt.fields.EncodedLibraryVersion,
				FileExtension:         tt.fields.FileExtension,
				FileName:              tt.fields.FileName,
				FileSize:              tt.fields.FileSize,
				FolderName:            tt.fields.FolderName,
				FormatInfo:            tt.fields.FormatInfo,
				Format:                tt.fields.Format,
				FormatURL:             tt.fields.FormatURL,
				Genre:                 tt.fields.Genre,
				InternetMediaType:     tt.fields.InternetMediaType,
				KindOfStream:          tt.fields.KindOfStream,
				OverallBitRate:        tt.fields.OverallBitRate,
				OverallBitRateMode:    tt.fields.OverallBitRateMode,
				Part:                  tt.fields.Part,
				PartTotal:             tt.fields.PartTotal,
				Performer:             tt.fields.Performer,
				RecordedDate:          tt.fields.RecordedDate,
				SamplesCount:          tt.fields.SamplesCount,
				SamplingRate:          tt.fields.SamplingRate,
				StreamIdentifier:      tt.fields.StreamIdentifier,
				StreamSize:            tt.fields.StreamSize,
				Title:                 tt.fields.Title,
				TrackNamePosition:     tt.fields.TrackNamePosition,
				TrackName:             tt.fields.TrackName,
				TrackNameTotal:        tt.fields.TrackNameTotal,
				WritingLibrary:        tt.fields.WritingLibrary,
			}
			g := m.ToProto()
			w := tt.want
			assert.EqualValues(t, g.AlbumPerformer, w.AlbumPerformer, "AlbumPerformer")
			assert.EqualValues(t, g.Album, w.Album, "Album")
			assert.EqualValues(t, g.AudioCount, w.AudioCount, "AudioCount")
			assert.EqualValues(t, g.BitDepth, w.BitDepth, "BitDepth")
			assert.EqualValues(t, g.BitDepthString, w.BitDepthString, "BitDepthString")
			assert.EqualValues(t, g.BitRate, w.BitRate, "BitRate")
			assert.EqualValues(t, g.BitRateMode, w.BitRateMode, "BitRateMode")
			assert.EqualValues(t, g.Channels, w.Channels, "Channels")
			assert.EqualValues(t, g.ChannelLayout, w.ChannelLayout, "ChannelsLayout")
			assert.EqualValues(t, g.ChannelPositions, w.ChannelPositions, "ChannelsPosition")
			assert.EqualValues(t, g.Comment, w.Comment, "Comment")
			assert.EqualValues(t, g.CommercialName, w.CommercialName, "CommercialName")
			assert.EqualValues(t, g.CompleteName, w.CompleteName, "CompleteName")
			assert.EqualValues(t, g.CompressionMode, w.CompressionMode, "CompressionMode")
			assert.EqualValues(t, g.CountOfAudioStreams, w.CountOfAudioStreams, "CountOfAudioStreams")
			assert.EqualValues(t, g.Duration, w.Duration, "Duration")
			assert.EqualValues(t, g.EncodedLibraryDate, w.EncodedLibraryDate, "EncodedLibraryDate")
			assert.EqualValues(t, g.EncodedLibraryName, w.EncodedLibraryName, "EncodedLibraryName")
			assert.EqualValues(t, g.EncodedLibraryVersion, w.EncodedLibraryVersion, "EncodedLibraryVersion")
			assert.EqualValues(t, g.FileExtension, w.FileExtension, "FileExtension")
			assert.EqualValues(t, g.FileName, w.FileName, "FileName")
			assert.EqualValues(t, g.FileSize, w.FileSize, "FileSize")
			assert.EqualValues(t, g.FolderName, w.FolderName, "FolderName")
			assert.EqualValues(t, g.FormatInfo, w.FormatInfo, "FormatInfo")
			assert.EqualValues(t, g.Format, w.Format, "Format")
			assert.EqualValues(t, g.FormatUrl, w.FormatUrl, "FormatURL")
			assert.EqualValues(t, g.Genre, w.Genre, "Genre")
			assert.EqualValues(t, g.InternetMediaType, w.InternetMediaType, "InternetMediaType")
			assert.EqualValues(t, g.KindOfStream, w.KindOfStream, "KindOfStream")
			assert.EqualValues(t, g.OverallBitRate, w.OverallBitRate, "OverallBitRate")
			assert.EqualValues(t, g.OverallBitRateMode, w.OverallBitRateMode, "OverallBitRateMode")
			assert.EqualValues(t, g.Part, w.Part, "Part")
			assert.EqualValues(t, g.PartTotal, w.PartTotal, "PartTotal")
			assert.EqualValues(t, g.Performer, w.Performer, "Performer")
			assert.EqualValues(t, g.RecordedDate, w.RecordedDate, "RecordedDate")
			assert.EqualValues(t, g.SamplingCount, w.SamplingCount, "SamplesCount")
			assert.EqualValues(t, g.SamplingRate, w.SamplingRate, "SamplingRate")
			assert.EqualValues(t, g.StreamIdentifier, w.StreamIdentifier, "StreamIdentifier")
			assert.EqualValues(t, g.StreamSize, w.StreamSize, "StreamSize")
			assert.EqualValues(t, g.Title, w.Title, "Title")
			assert.EqualValues(t, g.TrackPosition, w.TrackPosition, "TrackNamePosition")
			assert.EqualValues(t, g.Track, w.Track, "TrackName")
			assert.EqualValues(t, g.TrackNameTotal, w.TrackNameTotal, "TrackNameTotal")
			assert.EqualValues(t, g.WritingLibrary, w.WritingLibrary, "WritingLibrary")
		})
	}
}
