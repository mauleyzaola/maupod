package conversion

import (
	"reflect"
	"strings"
	"testing"
	"time"

	"github.com/mauleyzaola/maupod/src/server/pkg/dbdata/orm"
	"github.com/mauleyzaola/maupod/src/server/pkg/helpers"
	"github.com/mauleyzaola/maupod/src/server/pkg/pb"
	"github.com/stretchr/testify/assert"
	"github.com/volatiletech/null/v8"
)

var year2000 = time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC)
var year2001 = year2000.AddDate(1, 0, 0)
var sampleMedium = &orm.Medium{
	ID:                    "1",
	Sha:                   "2",
	Location:              "3",
	FileExtension:         ".FLAC",
	Format:                "5",
	FileSize:              6,
	Duration:              7,
	OverallBitRateMode:    "8",
	OverallBitRate:        9,
	StreamSize:            10,
	Album:                 "11",
	Title:                 "12",
	Track:                 "13",
	TrackPosition:         14,
	Performer:             "15",
	LastScan:              year2000,
	Genre:                 "16",
	RecordedDate:          17,
	Comment:               "18",
	Channels:              "19",
	ChannelPositions:      "20",
	ChannelLayout:         "21",
	SamplingRate:          22,
	SamplingCount:         23,
	BitDepth:              24,
	CompressionMode:       "25",
	EncodedLibraryName:    "27",
	EncodedLibraryVersion: "28",
	BitRateMode:           "29",
	BitRate:               30,
	ModifiedDate:          year2001,
	TrackNameTotal:        31,
	AlbumPerformer:        "32",
	AudioCount:            33,
	BitDepthString:        "34",
	CommercialName:        "35",
	CompleteName:          "36",
	CountOfAudioStreams:   37,
	EncodedLibraryDate:    "38",
	FileName:              "39",
	FolderName:            "40",
	FormatInfo:            "41",
	FormatURL:             "42",
	InternetMediaType:     "43",
	KindOfStream:          "44",
	Part:                  45,
	PartTotal:             46,
	StreamIdentifier:      47,
	WritingLibrary:        "48",
	Composer:              "49",
	ShaImage:              "50",
	LastImageScan:         null.TimeFrom(year2000),
	ImageLocation:         "51",
	AlbumIdentifier:       "52",
	IsCompilation:         true,
}
var sampleMedia = &pb.Media{
	Id:                    "1",
	Sha:                   "2",
	Location:              "3",
	FileExtension:         ".FLAC",
	Format:                "5",
	FileSize:              6,
	Duration:              7,
	OverallBitRateMode:    "8",
	OverallBitRate:        9,
	StreamSize:            10,
	Album:                 "11",
	Title:                 "12",
	Track:                 "13",
	TrackPosition:         14,
	Performer:             "15",
	LastScan:              helpers.TimeToTs2(year2000),
	Genre:                 "16",
	RecordedDate:          17,
	Comment:               "18",
	Channels:              "19",
	ChannelPositions:      "20",
	ChannelLayout:         "21",
	SamplingRate:          22,
	SamplingCount:         23,
	BitDepth:              24,
	CompressionMode:       "25",
	EncodedLibraryName:    "27",
	EncodedLibraryVersion: "28",
	BitRateMode:           "29",
	BitRate:               30,
	TrackNameTotal:        31,
	AlbumPerformer:        "32",
	AudioCount:            33,
	BitDepthString:        "34",
	CommercialName:        "35",
	CompleteName:          "36",
	CountOfAudioStreams:   37,
	EncodedLibraryDate:    "38",
	FileName:              "39",
	FolderName:            "40",
	FormatInfo:            "41",
	FormatUrl:             "42",
	InternetMediaType:     "43",
	KindOfStream:          "44",
	Part:                  45,
	PartTotal:             46,
	StreamIdentifier:      47,
	WritingLibrary:        "48",
	ModifiedDate:          helpers.TimeToTs2(year2001),
	Composer:              "49",
	ShaImage:              "50",
	LastImageScan:         helpers.TimeToTs(&year2000),
	ImageLocation:         "51",
	AlbumIdentifier:       "52",
	IsCompilation:         true,
}

func TestMediaToORM(t *testing.T) {
	type args struct {
		v *pb.Media
	}
	tests := []struct {
		name string
		args args
		want *orm.Medium
	}{
		{
			name: "check all fields are present",
			args: args{v: sampleMedia},
			want: sampleMedium,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := MediaToORM(tt.args.v)
			w := tt.want

			assert.EqualValues(t, w.ID, g.ID)
			assert.EqualValues(t, w.Sha, g.Sha)
			assert.EqualValues(t, w.LastScan, g.LastScan)
			assert.EqualValues(t, strings.ToLower(w.FileExtension), g.FileExtension)
			assert.EqualValues(t, w.Format, g.Format)
			assert.EqualValues(t, w.FileSize, g.FileSize)
			assert.EqualValues(t, w.Duration, g.Duration)
			assert.EqualValues(t, w.OverallBitRateMode, g.OverallBitRateMode)
			assert.EqualValues(t, w.OverallBitRate, g.OverallBitRate)
			assert.EqualValues(t, w.StreamSize, g.StreamSize)
			assert.EqualValues(t, w.Album, g.Album)
			assert.EqualValues(t, w.Title, g.Title)
			assert.EqualValues(t, w.Track, g.Track)
			assert.EqualValues(t, w.TrackPosition, g.TrackPosition)
			assert.EqualValues(t, w.Performer, g.Performer)
			assert.EqualValues(t, w.Genre, g.Genre)
			assert.EqualValues(t, w.RecordedDate, g.RecordedDate)
			assert.EqualValues(t, w.Comment, g.Comment)
			assert.EqualValues(t, w.Channels, g.Channels)
			assert.EqualValues(t, w.ChannelPositions, g.ChannelPositions)
			assert.EqualValues(t, w.ChannelLayout, g.ChannelLayout)
			assert.EqualValues(t, w.SamplingRate, g.SamplingRate)
			assert.EqualValues(t, w.SamplingCount, g.SamplingCount)
			assert.EqualValues(t, w.BitDepth, g.BitDepth)
			assert.EqualValues(t, w.CompressionMode, g.CompressionMode)
			assert.EqualValues(t, w.EncodedLibraryName, g.EncodedLibraryName)
			assert.EqualValues(t, w.EncodedLibraryVersion, g.EncodedLibraryVersion)
			assert.EqualValues(t, w.BitRateMode, g.BitRateMode)
			assert.EqualValues(t, w.BitRate, g.BitRate)
			assert.EqualValues(t, w.TrackNameTotal, g.TrackNameTotal)
			assert.EqualValues(t, w.AlbumPerformer, g.AlbumPerformer)
			assert.EqualValues(t, w.AudioCount, g.AudioCount)
			assert.EqualValues(t, w.BitDepthString, g.BitDepthString)
			assert.EqualValues(t, w.CommercialName, g.CommercialName)
			assert.EqualValues(t, w.CompleteName, g.CompleteName)
			assert.EqualValues(t, w.CountOfAudioStreams, g.CountOfAudioStreams)
			assert.EqualValues(t, w.EncodedLibraryDate, g.EncodedLibraryDate)
			assert.EqualValues(t, w.FileName, g.FileName)
			assert.EqualValues(t, w.FolderName, g.FolderName)
			assert.EqualValues(t, w.FormatInfo, g.FormatInfo)
			assert.EqualValues(t, w.FormatURL, g.FormatURL)
			assert.EqualValues(t, w.InternetMediaType, g.InternetMediaType)
			assert.EqualValues(t, w.KindOfStream, g.KindOfStream)
			assert.EqualValues(t, w.Part, g.Part)
			assert.EqualValues(t, w.PartTotal, g.PartTotal)
			assert.EqualValues(t, w.StreamIdentifier, g.StreamIdentifier)
			assert.EqualValues(t, w.WritingLibrary, g.WritingLibrary)
			assert.EqualValues(t, w.ModifiedDate, g.ModifiedDate)
			assert.EqualValues(t, w.Composer, g.Composer)
			assert.EqualValues(t, w.ShaImage, g.ShaImage)
			assert.EqualValues(t, w.LastImageScan, g.LastImageScan)
			assert.EqualValues(t, w.ImageLocation, g.ImageLocation)
			assert.EqualValues(t, w.AlbumIdentifier, g.AlbumIdentifier)
			assert.EqualValues(t, w.IsCompilation, g.IsCompilation)
		})
	}
}

func TestMediaFromORM(t *testing.T) {
	type args struct {
		v *orm.Medium
	}
	tests := []struct {
		name string
		args args
		want *pb.Media
	}{
		{
			name: "check all fields are present",
			args: args{v: sampleMedium},
			want: sampleMedia,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := MediaFromORM(tt.args.v)
			w := tt.want

			assert.EqualValues(t, w.Id, g.Id)
			assert.EqualValues(t, w.Sha, g.Sha)
			assert.EqualValues(t, w.LastScan, g.LastScan)
			assert.EqualValues(t, w.FileExtension, g.FileExtension)
			assert.EqualValues(t, w.Format, g.Format)
			assert.EqualValues(t, w.FileSize, g.FileSize)
			assert.EqualValues(t, w.Duration, g.Duration)
			assert.EqualValues(t, w.OverallBitRateMode, g.OverallBitRateMode)
			assert.EqualValues(t, w.OverallBitRate, g.OverallBitRate)
			assert.EqualValues(t, w.StreamSize, g.StreamSize)
			assert.EqualValues(t, w.Album, g.Album)
			assert.EqualValues(t, w.Title, g.Title)
			assert.EqualValues(t, w.Track, g.Track)
			assert.EqualValues(t, w.TrackPosition, g.TrackPosition)
			assert.EqualValues(t, w.Performer, g.Performer)
			assert.EqualValues(t, w.Genre, g.Genre)
			assert.EqualValues(t, w.RecordedDate, g.RecordedDate)
			assert.EqualValues(t, w.Comment, g.Comment)
			assert.EqualValues(t, w.Channels, g.Channels)
			assert.EqualValues(t, w.ChannelPositions, g.ChannelPositions)
			assert.EqualValues(t, w.ChannelLayout, g.ChannelLayout)
			assert.EqualValues(t, w.SamplingRate, g.SamplingRate)
			assert.EqualValues(t, w.SamplingCount, g.SamplingCount)
			assert.EqualValues(t, w.BitDepth, g.BitDepth)
			assert.EqualValues(t, w.CompressionMode, g.CompressionMode)
			assert.EqualValues(t, w.EncodedLibraryName, g.EncodedLibraryName)
			assert.EqualValues(t, w.EncodedLibraryVersion, g.EncodedLibraryVersion)
			assert.EqualValues(t, w.BitRateMode, g.BitRateMode)
			assert.EqualValues(t, w.BitRate, g.BitRate)
			assert.EqualValues(t, w.TrackNameTotal, g.TrackNameTotal)
			assert.EqualValues(t, w.AlbumPerformer, g.AlbumPerformer)
			assert.EqualValues(t, w.AudioCount, g.AudioCount)
			assert.EqualValues(t, w.BitDepthString, g.BitDepthString)
			assert.EqualValues(t, w.CommercialName, g.CommercialName)
			assert.EqualValues(t, w.CompleteName, g.CompleteName)
			assert.EqualValues(t, w.CountOfAudioStreams, g.CountOfAudioStreams)
			assert.EqualValues(t, w.EncodedLibraryDate, g.EncodedLibraryDate)
			assert.EqualValues(t, w.FileName, g.FileName)
			assert.EqualValues(t, w.FolderName, g.FolderName)
			assert.EqualValues(t, w.FormatInfo, g.FormatInfo)
			assert.EqualValues(t, w.FormatUrl, g.FormatUrl)
			assert.EqualValues(t, w.InternetMediaType, g.InternetMediaType)
			assert.EqualValues(t, w.KindOfStream, g.KindOfStream)
			assert.EqualValues(t, w.Part, g.Part)
			assert.EqualValues(t, w.PartTotal, g.PartTotal)
			assert.EqualValues(t, w.StreamIdentifier, g.StreamIdentifier)
			assert.EqualValues(t, w.WritingLibrary, g.WritingLibrary)
			assert.EqualValues(t, w.ModifiedDate, g.ModifiedDate)
			assert.EqualValues(t, w.Composer, g.Composer)
			assert.EqualValues(t, w.ShaImage, g.ShaImage)
			assert.EqualValues(t, w.LastImageScan.Nanos, g.LastImageScan.Nanos)
			assert.EqualValues(t, w.ImageLocation, g.ImageLocation)
			assert.EqualValues(t, w.AlbumIdentifier, g.AlbumIdentifier)
			assert.EqualValues(t, w.IsCompilation, g.IsCompilation)
		})
	}
}

func TestMediasFromORM(t *testing.T) {
	type args struct {
		a []*orm.Medium
	}
	tests := []struct {
		name string
		args args
		want []*pb.Media
	}{
		{
			name: "check mapping matches one item",
			args: args{a: []*orm.Medium{sampleMedium}},
			want: []*pb.Media{sampleMedia},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MediasFromORM(tt.args.a...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MediasFromORM() = %v, want %v", got, tt.want)
			}
		})
	}
}
