package taggers

import (
	"testing"

	"github.com/mauleyzaola/maupod/src/pkg/helpers"

	"github.com/mauleyzaola/maupod/src/protos"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMP3Tagger(t *testing.T) {
	if !helpers.ProgramExists(mp3TagProgram) {
		t.Skipf("cannot find program: %s", mp3TagProgram)
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
				filename: "./test_data/sample.mp3",
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
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t1 *testing.T) {
			tagger, err := TaggerFactory(tt.args.filename)
			require.NoError(t, err)

			if !tagger.ProgramExist() {
				t.Skip("program not found in path")
			}

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
