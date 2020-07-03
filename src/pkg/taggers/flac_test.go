package taggers

import (
	"testing"

	"github.com/mauleyzaola/maupod/src/pkg/information"
	"github.com/mauleyzaola/maupod/src/pkg/pb"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func readMediaInfo(filename string) (*pb.Media, error) {
	return information.MediaFromFile(filename)
}

func compareMedia(t *testing.T, m1, m2 *pb.Media) {
	assert.EqualValues(t, m1.Album, m2.Album)
	assert.EqualValues(t, m1.TrackPosition, m2.TrackPosition)
	assert.EqualValues(t, m1.Performer, m2.Performer)
	assert.EqualValues(t, m1.Genre, m2.Genre)
	assert.EqualValues(t, m1.RecordedDate, m2.RecordedDate)
	assert.EqualValues(t, m1.Comment, m2.Comment)
	assert.EqualValues(t, m1.TrackNameTotal, m2.TrackNameTotal)
}

func TestFLACTagger(t *testing.T) {
	type args struct {
		filename    string
		taggedMedia *pb.Media
	}
	tests := []struct {
		name         string
		args         args
		wantErr      bool
		removedMedia *pb.Media
	}{
		{
			args: args{
				filename: "./test_data/sample.flac",
				taggedMedia: &pb.Media{
					Album:          "Abbey Road",
					Track:          "Here comes the sun",
					TrackPosition:  7,
					Performer:      "The Beatles",
					Genre:          "Rock",
					RecordedDate:   1969,
					Comment:        "Beer is good",
					TrackNameTotal: 12,
				},
			},
			wantErr:      false,
			removedMedia: &pb.Media{},
		},
		{
			args: args{
				filename: "./test_data/wrong.flac",
			},
			wantErr:      true,
			removedMedia: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t1 *testing.T) {
			tagger, err := TaggerFactory(tt.args.filename)
			require.NoError(t, err)
			if tt.wantErr {
				err := tagger.RemoveAll(tt.args.filename)
				assert.Error(t, err)
				return
			}
			err = tagger.RemoveAll(tt.args.filename)
			require.NoError(t, err)
			err = tagger.Tag(tt.args.taggedMedia, tt.args.filename)
			require.NoError(t, err)

			newMedia, err := readMediaInfo(tt.args.filename)
			require.NoError(t, err)
			compareMedia(t, tt.args.taggedMedia, newMedia)

			if err = tagger.RemoveAll(tt.args.filename); (err != nil) != tt.wantErr {
				t1.Errorf("RemoveTag() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			got, err := readMediaInfo(tt.args.filename)
			require.NoError(t, err)
			compareMedia(t, tt.removedMedia, got)
		})
	}
}
