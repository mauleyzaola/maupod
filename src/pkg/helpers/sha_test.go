package helpers

import "testing"

func TestSHAFromFile(t *testing.T) {
	type args struct {
		filename string
	}
	tests := []struct {
		name       string
		args       args
		wantLength int
		wantErr    bool
	}{
		{
			name:       "sha legth text file",
			args:       args{filename: "../information/test_data/mediainfo1.txt"},
			wantErr:    false,
			wantLength: 64,
		},
		{
			name:       "sha legth audio file",
			args:       args{filename: "../information/test_data/sample1.m4a"},
			wantErr:    false,
			wantLength: 64,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := SHAFromFile(tt.args.filename)
			if (err != nil) != tt.wantErr {
				t.Errorf("SHAFromFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if val := len(got); val != tt.wantLength {
				t.Errorf("SHAFromFile() got = %v, want %v", val, tt.wantLength)
			}
		})
	}
}
