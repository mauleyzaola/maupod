package images

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExtractImageFromMedia(t *testing.T) {
	type args struct {
		w        *bytes.Buffer
		filename string
	}
	tests := []struct {
		name       string
		args       args
		wantErr    bool
		wantLenght int
	}{
		{
			name: "invalid file",
			args: args{
				w:        nil,
				filename: "sha la la",
			},
			wantErr: true,
		},
		{
			name: "600x600",
			args: args{
				w:        &bytes.Buffer{},
				filename: "./test_data/sample-600x600.m4a",
			},
			wantErr:    false,
			wantLenght: 167765,
		},
		{
			name: "1425x1425",
			args: args{
				w:        &bytes.Buffer{},
				filename: "./test_data/sample-1425x1425.flac",
			},
			wantErr:    false,
			wantLenght: 3633010,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &bytes.Buffer{}
			err := ExtractImageFromMedia(w, tt.args.filename)
			if (err != nil) != tt.wantErr {
				t.Errorf("ExtractImageFromMedia() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.wantLenght != 0 {
				assert.EqualValues(t, tt.wantLenght, w.Len())
			}
		})
	}
}
