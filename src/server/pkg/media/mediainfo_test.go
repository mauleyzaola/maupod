package media

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/mauleyzaola/maupod/src/server/pkg/domain"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMediainfo(t *testing.T) {
	type args struct {
		filename []string
	}

	// sanity check if audio files exist
	audioFile1 := filepath.Join(os.Getenv("HOME"), "Downloads", "1.flac")
	_, err := os.Stat(audioFile1)
	if err != nil {
		t.Skip(err)
	}
	audioFile2 := filepath.Join(os.Getenv("HOME"), "Downloads", "2.mp3")
	if _, err = os.Stat(audioFile2); err != nil {
		t.Skip(err)
	}

	tests := []struct {
		name       string
		args       args
		wantErr    bool
		wantLength int
	}{
		{
			name:       "one file parsing",
			wantErr:    false,
			args:       args{filename: []string{audioFile1}},
			wantLength: 1,
		},
		{
			name:       "two file parsing",
			wantErr:    false,
			args:       args{filename: []string{audioFile1, audioFile2}},
			wantLength: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Mediainfo(tt.args.filename...)
			if (err != nil) != tt.wantErr {
				t.Errorf("Mediainfo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if length := len(got); tt.wantLength != length {
				t.Errorf("Mediainfo() length = %v, wantLength %v", length, tt.wantLength)
				return
			}
		})
	}
}

func TestMediaInfo_ToDomain(t *testing.T) {
	data, err := ioutil.ReadFile("./test_data/one.json")
	require.NoError(t, err)
	require.NotNil(t, data, "JSON data should be available in file")

	var mediaInfo MediaInfo
	err = json.Unmarshal(data, &mediaInfo)
	require.NoError(t, err, "mediaInfo should have been deserialized")

	type fields struct {
		Media Media
	}
	tests := []struct {
		name   string
		fields fields
		want   *domain.Media
	}{
		{
			name: "one song test conversion",
			fields: fields{
				Media: mediaInfo.Media,
			},
			want: &domain.Media{
				ID:                    "",
				Location:              "",
				FileExtension:         "flac",
				Format:                "FLAC",
				FileSize:              24953007,
				Duration:              238,
				OverallBitRateMode:    "VBR",
				OverallBitRate:        838757,
				StreamSize:            0,
				Album:                 "Gold - Greatest Hits",
				Title:                 "One Of Us",
				Track:                 "One Of Us",
				TrackPosition:         16,
				Performer:             "ABBA",
				Genre:                 "Dance-pop",
				RecordedDate:          1993,
				FileModifiedDate:      time.Date(2020, 04, 27, 1, 4, 9, 0, time.UTC),
				Comment:               "Music For All The World",
				Channels:              "2",
				ChannelPositions:      "Front: L R",
				ChannelLayout:         "L R",
				SamplingRate:          44100,
				SamplingCount:         10495800,
				BitDepth:              16,
				CompressionMode:       "Lossless",
				EncodedLibrary:        "reference libFLAC 1.2.1 20070917",
				EncodedLibraryName:    "libFLAC",
				EncodedLibraryVersion: "1.2.1",
				BitRateMode:           "VBR",
				BitRate:               838611,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &MediaInfo{
				Media: tt.fields.Media,
			}
			got := m.ToDomain()
			assert.EqualValues(t, tt.want, got, "conversion should succeed")
		})
	}
}
