package media

import (
	"os"
	"path/filepath"
	"testing"
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
