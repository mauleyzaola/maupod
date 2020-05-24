package images

import (
	"bytes"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/mauleyzaola/maupod/src/server/pkg/helpers"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func createReader(filename string) io.Reader {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil
	}
	return bytes.NewBuffer(data)
}

func TestSize(t *testing.T) {
	type args struct {
		r io.Reader
	}
	tests := []struct {
		name    string
		args    args
		wantX   int
		wantY   int
		wantErr bool
	}{
		{
			name:    "missing reader",
			wantErr: true,
		},
		{
			name:    "600x600",
			args:    args{r: createReader("./test_data/600x600.png")},
			wantX:   600,
			wantY:   600,
			wantErr: false,
		},
		{
			name:    "1425x1425",
			args:    args{r: createReader("./test_data/1425x1425.png")},
			wantX:   1425,
			wantY:   1425,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotX, gotY, err := Size(tt.args.r)
			if (err != nil) != tt.wantErr {
				t.Errorf("Size() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotX != tt.wantX {
				t.Errorf("Size() gotX = %v, want %v", gotX, tt.wantX)
			}
			if gotY != tt.wantY {
				t.Errorf("Size() gotY = %v, want %v", gotY, tt.wantY)
			}
		})
	}
}

func TestImageResize(t *testing.T) {
	type args struct {
		r        io.Reader
		filename string
		width    int
		height   int
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
		wantX   int
		wantY   int
	}{
		{
			name:    "missing paramters",
			wantErr: true,
		},
		{
			name: "600x600 to 300x300",
			args: args{
				r:        createReader("./test_data/600x600.png"),
				filename: filepath.Join(os.TempDir(), helpers.NewUUID()+".png"),
				width:    300,
				height:   300,
			},
			wantX:   300,
			wantY:   300,
			wantErr: false,
		},
		{
			name: "600x600 to 300x300 automatically",
			args: args{
				r:        createReader("./test_data/600x600.png"),
				filename: filepath.Join(os.TempDir(), helpers.NewUUID()+".png"),
				width:    300,
				height:   0,
			},
			wantX:   300,
			wantY:   300,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ImageResize(tt.args.r, tt.args.filename, tt.args.width, tt.args.height)
			if (err != nil) != tt.wantErr {
				t.Errorf("ImageResize() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.args.filename != "" {
				if tt.wantX != 0 || tt.wantY != 0 {
					r := createReader(tt.args.filename)
					x, y, err := Size(r)
					require.NoError(t, err)
					assert.EqualValues(t, tt.wantX, x)
					assert.EqualValues(t, tt.wantY, y)
				}
			}
		})
	}
}
