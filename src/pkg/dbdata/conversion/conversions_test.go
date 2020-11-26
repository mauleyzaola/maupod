package conversion

import (
	"reflect"
	"testing"
	"time"

	"github.com/mauleyzaola/maupod/src/pkg/dbdata/orm"
	"github.com/mauleyzaola/maupod/src/pkg/helpers"
	"github.com/mauleyzaola/maupod/src/protos"
	"github.com/stretchr/testify/assert"
	"github.com/volatiletech/null/v8"
)

var year2000 = time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC)
var year2001 = year2000.AddDate(1, 0, 0)
var sampleMedium = orm.Medium{
	ID:                    "1",
	Sha:                   "2",
	Location:              "3",
	FileExtension:         ".FLAC",
	Format:                "5",
	FileSize:              6,
	Duration:              265906,
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
	LastImageScan:         null.TimeFrom(year2000),
	ImageLocation:         "51",
	AlbumIdentifier:       "52",
}
var sampleMedia = protos.Media{
	Id:                    "1",
	Sha:                   "2",
	Location:              "3",
	FileExtension:         ".FLAC",
	Format:                "5",
	FileSize:              6,
	Duration:              265906,
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
	LastImageScan:         helpers.TimeToTs(&year2000),
	ImageLocation:         "51",
	AlbumIdentifier:       "52",
	Seconds:               265,
}

//Object Playlist
var samplePlaylistORM = &orm.Playlist{
	ID:   "1",
	Name: "Playlist Rock",
}
var samplePlaylistPB = &protos.PlayList{
	Id:   "1",
	Name: "Playlist Rock",
}

//Object PlaylistItem
var samplePlaylistItemORM = &orm.PlaylistItem{
	ID:         "1",
	PlaylistID: "1",
	Position:   1,
	MediaID:    "12345678",
}
var samplePlaylistItemPB = &protos.PlaylistItem{
	Id:       "1",
	Playlist: &protos.PlayList{Id: "1"},
	Position: 1,
	Media:    &protos.Media{Id: "12345678"},
}

func TestMediaToORM(t *testing.T) {
	type args struct {
		v *protos.Media
	}
	tests := []struct {
		name string
		args args
		want *orm.Medium
	}{
		{
			name: "check all fields are present",
			args: args{
				v: &protos.Media{
					Album:                 "11",
					AlbumIdentifier:       "52",
					AlbumPerformer:        "32",
					AudioCount:            33,
					BitDepth:              24,
					BitDepthString:        "34",
					BitRateMode:           "29",
					BitRate:               30,
					Channels:              "19",
					ChannelPositions:      "20",
					ChannelLayout:         "21",
					Comment:               "18",
					CommercialName:        "35",
					CompleteName:          "36",
					Composer:              "49",
					CompressionMode:       "25",
					CountOfAudioStreams:   37,
					Duration:              265906,
					EncodedLibraryName:    "27",
					EncodedLibraryDate:    "38",
					EncodedLibraryVersion: "28",
					FileName:              "39",
					FolderName:            "40",
					FormatInfo:            "41",
					FormatUrl:             "42",
					FileExtension:         ".FLAC",
					FileSize:              6,
					Format:                "FLAC",
					Genre:                 "16",
					Id:                    "1",
					ImageLocation:         "51",
					InternetMediaType:     "43",
					KindOfStream:          "44",
					LastImageScan:         helpers.TimeToTs2(year2000),
					LastScan:              helpers.TimeToTs2(year2000),
					Location:              "3",
					ModifiedDate:          helpers.TimeToTs2(year2001),
					OverallBitRate:        9,
					OverallBitRateMode:    "8",
					Part:                  45,
					PartTotal:             46,
					Performer:             "15",
					RecordedDate:          1789,
					SamplingRate:          22,
					SamplingCount:         23,
					Seconds:               265,
					Sha:                   "2",
					StreamIdentifier:      47,
					StreamSize:            10,
					Title:                 "12",
					Track:                 "13",
					TrackNameTotal:        31,
					TrackPosition:         14,
					WritingLibrary:        "48",
				},
			},
			want: &orm.Medium{
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
				FormatURL:             "42",
				InternetMediaType:     "43",
				KindOfStream:          "44",
				Part:                  45,
				PartTotal:             46,
				StreamIdentifier:      47,
				WritingLibrary:        "48",
				Composer:              "49",
				ImageLocation:         "51",
				AlbumIdentifier:       "52",
				Album:                 "11",
				Channels:              "19",
				ChannelPositions:      "20",
				ChannelLayout:         "21",
				Comment:               "18",
				Duration:              265906,
				FileExtension:         ".flac",
				FileSize:              6,
				Format:                "FLAC",
				Genre:                 "16",
				ID:                    "1",
				LastImageScan:         null.TimeFrom(year2000),
				LastScan:              year2000,
				Location:              "3",
				ModifiedDate:          year2001,
				OverallBitRate:        9,
				OverallBitRateMode:    "8",
				Performer:             "15",
				RecordedDate:          1789,
				SamplingRate:          22,
				SamplingCount:         23,
				Sha:                   "2",
				StreamSize:            10,
				Title:                 "12",
				Track:                 "13",
				TrackPosition:         14,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := MediaToORM(tt.args.v)
			w := tt.want
			assert.EqualValues(t, w, g)
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
		want *protos.Media
	}{
		{
			name: "check all fields are present",
			args: args{
				v: &orm.Medium{
					Duration: 265906,
				},
			},
			want: &protos.Media{
				Duration: 265906,
				Seconds:  265,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := MediaFromORM(tt.args.v)
			w := tt.want
			assert.EqualValues(t, w.Duration, g.Duration)
			assert.EqualValues(t, w.Seconds, g.Seconds)
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
		want []*protos.Media
	}{
		{
			name: "check mapping matches one item",
			args: args{a: []*orm.Medium{&sampleMedium}},
			want: []*protos.Media{&sampleMedia},
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

//Function PlayList
func TestPlaylistToORM(t *testing.T) {
	type args struct {
		v *protos.PlayList
	}
	tests := []struct {
		name string
		args args
		want *orm.Playlist
	}{
		{
			name: "check all fields are present",
			args: args{v: samplePlaylistPB},
			want: samplePlaylistORM,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := PlaylistToORM(tt.args.v)
			w := tt.want

			assert.EqualValues(t, w.ID, g.ID)
			assert.EqualValues(t, w.Name, g.Name)
		})
	}
}
func TestPlaylistFromORM(t *testing.T) {
	type args struct {
		v *orm.Playlist
	}
	tests := []struct {
		name string
		args args
		want *protos.PlayList
	}{
		{
			name: "check all fields are present",
			args: args{v: samplePlaylistORM},
			want: samplePlaylistPB,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := PlaylistFromORM(tt.args.v)
			w := tt.want

			assert.EqualValues(t, w.Id, g.Id)
			assert.EqualValues(t, w.Name, g.Name)
		})
	}
}
func TestPlayListsFromORM(t *testing.T) {
	type args struct {
		a []*orm.Playlist
	}
	tests := []struct {
		name string
		args args
		want []*protos.PlayList
	}{
		{
			name: "check mapping matches one item",
			args: args{a: []*orm.Playlist{samplePlaylistORM}},
			want: []*protos.PlayList{samplePlaylistPB},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PlaylistsFromORM(tt.args.a...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PlaylistsFromORM() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPlaylistItemToORM(t *testing.T) {
	type args struct {
		v *protos.PlaylistItem
	}
	tests := []struct {
		name string
		args args
		want *orm.PlaylistItem
	}{
		{
			name: "check all fields are present",
			args: args{v: samplePlaylistItemPB},
			want: samplePlaylistItemORM,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := PlaylistItemToORM(tt.args.v)
			w := tt.want

			assert.EqualValues(t, w.ID, g.ID)
			assert.EqualValues(t, w.PlaylistID, g.PlaylistID)
			assert.EqualValues(t, w.Position, g.Position)
			assert.EqualValues(t, w.MediaID, g.MediaID)

		})
	}
}
func TestPlaylistItemFromORM(t *testing.T) {
	type args struct {
		v *orm.PlaylistItem
	}
	tests := []struct {
		name string
		args args
		want *protos.PlaylistItem
	}{
		{
			name: "check all fields are present",
			args: args{v: samplePlaylistItemORM},
			want: samplePlaylistItemPB,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := PlaylistItemFromORM(tt.args.v)
			w := tt.want

			assert.EqualValues(t, w.Id, g.Id)
			assert.EqualValues(t, w.Playlist, g.Playlist)
			assert.EqualValues(t, w.Position, g.Position)
			assert.EqualValues(t, w.Media, g.Media)

		})

	}
}

func TestPlaylistItemsFromORM(t *testing.T) {
	type args struct {
		a []*orm.PlaylistItem
	}
	tests := []struct {
		name string
		args args
		want []*protos.PlaylistItem
	}{
		{
			name: "check mapping matches one item",
			args: args{a: []*orm.PlaylistItem{samplePlaylistItemORM}},
			want: []*protos.PlaylistItem{samplePlaylistItemPB},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PlaylistItemsFromORM(tt.args.a...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PlaylistItemsFromORM() = %v, want %v", got, tt.want)
			}
		})
	}
}
