package information

import (
	"bytes"
	"io"
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestInfoString_Split(t *testing.T) {
	tests := []struct {
		name      string
		l         InfoString
		wantKey   string
		wantValue string
	}{
		{
			name:      "simple Split",
			l:         `Count                                    : 331`,
			wantKey:   "Count",
			wantValue: "331",
		},
		{
			name:      "Split double colon",
			l:         `Format/Url                               : https://xiph.org/flac/`,
			wantKey:   "Format/Url",
			wantValue: "https://xiph.org/flac/",
		},
		{
			name:      "Split missing value",
			l:         `General`,
			wantKey:   "",
			wantValue: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotKey, gotValue := tt.l.Split()
			if gotKey != tt.wantKey {
				t.Errorf("Split() gotKey = %v, want %v", gotKey, tt.wantKey)
			}
			if gotValue != tt.wantValue {
				t.Errorf("Split() gotValue = %v, want %v", gotValue, tt.wantValue)
			}
		})
	}
}

func TestToInfoData(t *testing.T) {
	file, err := os.Open("./test_data/mediainfo1.txt")
	require.NoError(t, err, "file should be present")
	defer func() {
		err = file.Close()
		assert.NoError(t, err, "file should be closed")
	}()

	var want = map[string][]string{
		"Album": {
			`Right In The Night (Fall In Love With Music) [JAM 659855 6]`,
		},
		"Bit depth": {
			`24 bits`,
			`24`,
		},
		"Commercial name": {
			`FLAC`,
			`FLAC`,
		},
		"Count": {
			`280`,
			`331`,
		},
		"Duration": {
			`00:06:07.200`,
			`00:06:07.200`,
			`00:06:07.200`,
			`00:06:07.200`,
			`367200`,
			`367200`,
			`6 min 7 s`,
			`6 min 7 s`,
			`6 min 7 s`,
			`6 min 7 s`,
			`6 min 7 s 200 ms`,
			`6 min 7 s 200 ms`,
		},
	}
	got := toInfoData(file)
	for k, v := range want {
		val, ok := got[k]
		assert.True(t, ok, k+" is missing")
		assert.ElementsMatch(t, v, val, k+" elements should match")
	}
}

func TestMediaParser(t *testing.T) {
	type args struct {
		r io.Reader
	}
	readFile := func(name string) io.Reader {
		data, err := ioutil.ReadFile(name)
		if err != nil {
			t.Error(err)
			return nil
		}
		return bytes.NewBuffer(data)
	}
	tests := []struct {
		name    string
		args    args
		want    *MediaInfo
		wantErr bool
	}{
		{
			name:    "empty",
			args:    args{r: nil},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "mediainfo1.txt",
			args:    args{readFile("./test_data/mediainfo1.txt")},
			wantErr: false,
			want: &MediaInfo{
				AlbumPerformer:        "Jam & Spoon",
				Album:                 "Right In The Night (Fall In Love With Music) [JAM 659855 6]",
				AudioCount:            331,
				AudioFormatList:       "FLAC",
				BitDepth:              24,
				BitDepthString:        "24 bits",
				BitRate:               3000111,
				BitRateMode:           "VBR",
				Channels:              "2",
				ChannelsLayout:        "L R",
				ChannelsPosition:      "Front: L R",
				Comment:               "",
				CommercialName:        "FLAC",
				CompleteName:          "/media/mau/music-library/music/Jam & Spoon/1993 - Right In The Night (Fall In Love With Music) [JAM 659855 6]/01 -  Right in the night (Fall in love with music) (feat. Plavka).flac",
				Compression:           "Lossless",
				CountOfAudioStreams:   1,
				Duration:              367200,
				EncodedLibraryDate:    "UTC 2019-08-04",
				EncodedLibraryName:    "libFLAC",
				EncodedLibraryVersion: "1.3.3",
				FileExtension:         "flac",
				FileName:              "01 -  Right in the night (Fall in love with music) (feat. Plavka).flac",
				FileSize:              137714396,
				FolderName:            "/media/mau/music-library/music/Jam & Spoon/1993 - Right In The Night (Fall In Love With Music) [JAM 659855 6]",
				FormatInfo:            "Free Lossless Audio Codec",
				Format:                "FLAC",
				FormatURL:             "https://xiph.org/flac/",
				Genre:                 "House, Euro House, Trance",
				InternetMediaType:     "audio/x-flac",
				KindOfStream:          "Audio",
				OverallBitRate:        3000314,
				OverallBitRateMode:    "VBR",
				Part:                  1,
				PartTotal:             1,
				Performer:             "Jam & Spoon",
				RecordedDate:          1993,
				SamplesCount:          35251200,
				SamplingRate:          96000,
				StreamIdentifier:      0,
				StreamSize:            137705073,
				Title:                 "Right in the night (Fall in love with music) (feat. Plavka)",
				TrackNamePosition:     1,
				TrackName:             "Right in the night (Fall in love with music) (feat. Plavka)",
				TrackNameTotal:        3,
				WritingLibrary:        "reference libFLAC 1.3.3 20190804",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g, err := ParseMediaInfo(tt.args.r)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseMediaInfo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if w := tt.want; w != nil {
				assert.EqualValues(t, w.AlbumPerformer, g.AlbumPerformer)
				assert.EqualValues(t, w.Album, g.Album)
				assert.EqualValues(t, w.AudioCount, g.AudioCount)
				assert.EqualValues(t, w.AudioFormatList, g.AudioFormatList)
				assert.EqualValues(t, w.BitDepth, g.BitDepth)
				assert.EqualValues(t, w.BitDepthString, g.BitDepthString)
				assert.EqualValues(t, w.BitRate, g.BitRate)
				assert.EqualValues(t, w.BitRateMode, g.BitRateMode)
				assert.EqualValues(t, w.Channels, g.Channels)
				assert.EqualValues(t, w.ChannelsLayout, g.ChannelsLayout)
				assert.EqualValues(t, w.ChannelsPosition, g.ChannelsPosition)
				assert.EqualValues(t, w.Comment, g.Comment)
				assert.EqualValues(t, w.CommercialName, g.CommercialName)
				assert.EqualValues(t, w.CompleteName, g.CompleteName)
				assert.EqualValues(t, w.Compression, g.Compression)
				assert.EqualValues(t, w.CountOfAudioStreams, g.CountOfAudioStreams)
				assert.EqualValues(t, w.Duration, g.Duration)
				assert.EqualValues(t, w.EncodedLibraryDate, g.EncodedLibraryDate)
				assert.EqualValues(t, w.EncodedLibraryName, g.EncodedLibraryName)
				assert.EqualValues(t, w.EncodedLibraryVersion, g.EncodedLibraryVersion)
				assert.EqualValues(t, w.FileExtension, g.FileExtension)
				assert.EqualValues(t, w.FileName, g.FileName)
				assert.EqualValues(t, w.FileSize, g.FileSize)
				assert.EqualValues(t, w.FolderName, g.FolderName)
				assert.EqualValues(t, w.FormatInfo, g.FormatInfo)
				assert.EqualValues(t, w.Format, g.Format)
				assert.EqualValues(t, w.FormatURL, g.FormatURL)
				assert.EqualValues(t, w.Genre, g.Genre)
				assert.EqualValues(t, w.InternetMediaType, g.InternetMediaType)
				assert.EqualValues(t, w.KindOfStream, g.KindOfStream)
				assert.EqualValues(t, w.OverallBitRate, g.OverallBitRate)
				assert.EqualValues(t, w.OverallBitRateMode, g.OverallBitRateMode)
				assert.EqualValues(t, w.Part, g.Part)
				assert.EqualValues(t, w.PartTotal, g.PartTotal)
				assert.EqualValues(t, w.Performer, g.Performer)
				assert.EqualValues(t, w.RecordedDate, g.RecordedDate)
				assert.EqualValues(t, w.SamplesCount, g.SamplesCount)
				assert.EqualValues(t, w.SamplingRate, g.SamplingRate)
				assert.EqualValues(t, w.StreamIdentifier, g.StreamIdentifier)
				assert.EqualValues(t, w.StreamSize, g.StreamSize)
				assert.EqualValues(t, w.Title, g.Title)
				assert.EqualValues(t, w.TrackNamePosition, g.TrackNamePosition)
				assert.EqualValues(t, w.TrackName, g.TrackName)
				assert.EqualValues(t, w.TrackNameTotal, g.TrackNameTotal)
				assert.EqualValues(t, w.WritingLibrary, g.WritingLibrary)
			}
		})
	}
}
