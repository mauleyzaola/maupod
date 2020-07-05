package types

import (
	"reflect"
	"testing"

	"github.com/mauleyzaola/maupod/src/pkg/pb"
)

func TestMedias_InsertAt(t *testing.T) {
	type args struct {
		m     *pb.Media
		index int
	}
	tests := []struct {
		name string
		me   Medias
		args args
		want Medias
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.me.InsertAt(tt.args.m, tt.args.index); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("InsertAt() = %v, want %v", got, tt.want)
			}
		})
	}
}
