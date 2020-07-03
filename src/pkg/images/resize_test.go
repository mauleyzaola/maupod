package images

import (
	"testing"
)

func TestSize(t *testing.T) {
	type args struct {
		filename string
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
			args:    args{filename: "./test_data/600x600.png"},
			wantX:   600,
			wantY:   600,
			wantErr: false,
		},
		{
			name:    "1425x1425",
			args:    args{filename: "./test_data/1425x1425.png"},
			wantX:   1425,
			wantY:   1425,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotX, gotY, err := Size(tt.args.filename)
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
