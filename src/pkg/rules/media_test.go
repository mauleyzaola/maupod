package rules

import (
	"reflect"
	"testing"
	"time"

	"github.com/mauleyzaola/maupod/src/pkg/pb"
)

func TestMediaPercentToSeconds(t *testing.T) {
	type args struct {
		m       *pb.Media
		percent float64
	}
	duration1 := time.Millisecond * time.Duration(26250)
	tests := []struct {
		name    string
		args    args
		want    *time.Duration
		wantErr bool
	}{
		{
			name: "04:22",
			args: args{
				m:       &pb.Media{Duration: 262506},
				percent: 10,
			},
			wantErr: false,
			want:    &duration1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := MediaPercentToSeconds(tt.args.m, tt.args.percent)
			if (err != nil) != tt.wantErr {
				t.Errorf("MediaPercentToSeconds() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MediaPercentToSeconds() got = %v, want %v", got, tt.want)
			}
		})
	}
}
