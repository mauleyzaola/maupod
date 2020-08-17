package images

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/mauleyzaola/maupod/src/pkg/helpers"
)

func TestImageResize(t *testing.T) {
	type args struct {
		source string
		target string
		width  int
		height int
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "600x600 to 500x500",
			args: args{
				source: "./test_data/600x600.png",
				target: filepath.Join(os.TempDir(), helpers.NewUUID()),
				width:  200,
				height: 200,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ImageResize(tt.args.source, tt.args.target, tt.args.width, tt.args.height); (err != nil) != tt.wantErr {
				t.Errorf("ImageResize() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
