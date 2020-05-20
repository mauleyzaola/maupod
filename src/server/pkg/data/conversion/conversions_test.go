package conversion

import (
	"reflect"
	"testing"
	"time"

	"github.com/mauleyzaola/maupod/src/server/pkg/data/orm"
	"github.com/mauleyzaola/maupod/src/server/pkg/helpers"
	"github.com/mauleyzaola/maupod/src/server/pkg/pb"
	"github.com/stretchr/testify/assert"
)

func TestMediaToORM(t *testing.T) {
	type args struct {
		v *pb.Media
	}
	tests := []struct {
		name string
		args args
		want *orm.Medium
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MediaToORM(tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MediaToORM() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMediaFromORM(t *testing.T) {
	var year2000 = time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC)
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
			args: args{v: &orm.Medium{
				ID:                    "1",
				Sha:                   "2",
				Location:              "3",
				FileExtension:         "4",
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
				EncodedLibrary:        "26",
				EncodedLibraryName:    "27",
				EncodedLibraryVersion: "28",
				BitRateMode:           "29",
				BitRate:               30,
				ModifiedDate:          year2000,
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
			},
			},
			want: &pb.Media{
				Id:                    "1",
				Sha:                   "2",
				Location:              "3",
				FileExtension:         "4",
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
				EncodedLibrary:        "26",
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
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := MediaFromORM(tt.args.v)
			w := tt.want

			assert.EqualValues(t, w.Id, g.Id)
			assert.EqualValues(t, w.Sha, g.Sha)
			assert.EqualValues(t, w.LastScan, g.LastScan)
		})
	}
}

/*
   string file_extension = 6;
   string format = 7;
   int64 file_size = 8;
   double duration = 9;
   string overall_bit_rate_mode = 10;
   int64 overall_bit_rate = 11;
   int64 stream_size = 12;
   string album = 13;
   string title = 14;
   string track = 15;
   int64 track_position = 16;
   string performer = 17;
   string genre = 18;
   int64 recorded_date = 19;
   string comment = 21;
   string channels = 22;
   string channel_positions = 23;
   string channel_layout = 24;
   int64 sampling_rate = 25;
   int64 sampling_count = 26;
   int64 bit_depth = 27;
   string compression_mode = 28;
   string encoded_library = 29;
   string encoded_library_name = 30;
   string encoded_library_version = 31;
   string bit_rate_mode = 32;
   int64 bit_rate = 33;
   int64 track_name_total = 34;
   string album_performer = 35;
   int64 audio_count = 36;
   string bit_depth_string = 37;
   string commercial_name = 38;
   string complete_name = 39;
   int64 count_of_audio_streams = 40;
   string encoded_library_date = 41;
   string file_name = 42;
   string folder_name = 43;
   string format_info = 44;
   string format_url = 45;
   string internet_media_type = 46;
   string kind_of_stream = 47;
   int64 part = 48;
   int64 part_total = 49;
   int64 stream_identifier = 50;
   string writing_library = 51;
   google.protobuf.Timestamp modified_date = 52;

*/
