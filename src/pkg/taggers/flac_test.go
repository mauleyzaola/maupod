package taggers

import (
	"testing"

	"github.com/mauleyzaola/maupod/src/pkg/helpers"

	"github.com/mauleyzaola/maupod/src/pkg/information"
	"github.com/mauleyzaola/maupod/src/protos"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func readMediaInfo(filename string) (*protos.Media, error) {
	raw, err := information.MediaInfoFromFile(filename)
	if err != nil {
		return nil, err
	}
	mediaInfo, err := information.ParseMediaInfo(raw)
	if err != nil {
		return nil, err
	}
	return mediaInfo.ToProto(), nil
}

func compareMedia(t *testing.T, m1, m2 *protos.Media) {
	assert.EqualValues(t, m1.Album, m2.Album)
	assert.EqualValues(t, m1.TrackPosition, m2.TrackPosition)
	assert.EqualValues(t, m1.Performer, m2.Performer)
	assert.EqualValues(t, m1.Genre, m2.Genre)
	assert.EqualValues(t, m1.RecordedDate, m2.RecordedDate)
	assert.EqualValues(t, m1.Comment, m2.Comment)
	assert.EqualValues(t, m1.TrackNameTotal, m2.TrackNameTotal)
}

func TestFLACTagger(t *testing.T) {
	if !helpers.ProgramExists(flacTagProgram) {
		t.Skipf("cannot find program: %s", flacTagProgram)
	}
	if !helpers.ProgramExists(information.MediaInfoProgram) {
		t.Skipf("cannot find program: %s", information.MediaInfoProgram)
	}

	type args struct {
		filename    string
		taggedMedia *protos.Media
	}
	tests := []struct {
		name         string
		args         args
		wantErr      bool
		removedMedia *protos.Media
	}{
		{
			args: args{
				filename: "./test_data/sample.flac",
				taggedMedia: &protos.Media{
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
			removedMedia: &protos.Media{},
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
