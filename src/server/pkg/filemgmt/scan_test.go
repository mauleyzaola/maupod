package filemgmt

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/mauleyzaola/maupod/src/server/pkg/helpers"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestScanFiles(t *testing.T) {
	var mode = os.ModePerm
	// create file structure
	uid := helpers.NewUUID()
	assert.NotEmpty(t, uid)

	root := filepath.Join(os.TempDir(), uid)
	err := os.MkdirAll(root, mode)
	require.NoError(t, err)

	defer func() {
		err = os.RemoveAll(root)
		assert.NoError(t, err)
	}()

	rush := filepath.Join(root, "rush")
	err = os.MkdirAll(rush, mode)
	require.NoError(t, err)

	donna := filepath.Join(root, "donna")
	err = os.MkdirAll(donna, mode)
	require.NoError(t, err)

	const (
		tomSayer     = "Tom Sawyer.mp3"
		subdivisions = "Subdivisions.flac"
		lastDance    = "Last Dance.mp3"
		springAffair = "Spring Affair.flac"
	)

	err = ioutil.WriteFile(filepath.Join(rush, tomSayer), nil, mode)
	require.NoError(t, err)
	err = ioutil.WriteFile(filepath.Join(rush, subdivisions), nil, mode)
	require.NoError(t, err)
	err = ioutil.WriteFile(filepath.Join(rush, "Rush in Rio.mp4"), nil, mode)
	require.NoError(t, err)

	err = ioutil.WriteFile(filepath.Join(donna, "The Collection.avi"), nil, mode)
	require.NoError(t, err)
	err = ioutil.WriteFile(filepath.Join(donna, lastDance), nil, mode)
	require.NoError(t, err)
	err = ioutil.WriteFile(filepath.Join(donna, springAffair), nil, mode)
	require.NoError(t, err)

	type args struct {
		extensions  []string
		directories []string
	}
	tests := []struct {
		name    string
		args    args
		want    []string
		wantErr bool
	}{
		{
			name: "lookup mp3 files",
			args: args{
				extensions:  []string{".mp3"},
				directories: []string{root},
			},
			wantErr: false,
			want: []string{
				filepath.Join(rush, tomSayer),
				filepath.Join(donna, lastDance),
			},
		},
		{
			name: "lookup audio files with duplicates",
			args: args{
				extensions:  []string{".mp3", ".FLAC"},
				directories: []string{root},
			},
			wantErr: false,
			want: []string{
				filepath.Join(rush, tomSayer),
				filepath.Join(donna, lastDance),
				filepath.Join(rush, subdivisions),
				filepath.Join(donna, springAffair),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ScanFiles(tt.args.extensions, tt.args.directories...)
			if (err != nil) != tt.wantErr {
				t.Errorf("ScanFiles() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			assert.ElementsMatch(t, got, tt.want)
		})
	}
}
