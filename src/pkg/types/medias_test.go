package types

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/mauleyzaola/maupod/src/protos"
)

var (
	m1 = &protos.Media{
		Id: "1",
	}
	m2 = &protos.Media{
		Id: "2",
	}
	m3 = &protos.Media{
		Id: "3",
	}
)

func TestMedias_InsertAt(t *testing.T) {
	type args struct {
		media *protos.Media
		index int
	}
	tests := []struct {
		name    string
		m       Medias
		args    args
		want    Medias
		wantErr bool
	}{
		{
			name: "index out of bounds",
			args: args{
				index: 1,
			},
			m:       nil,
			wantErr: true,
		},
		{
			name: "index out of bounds",
			args: args{
				index: -1,
			},
			m:       nil,
			wantErr: true,
		},
		{
			name: "add to empty slice",
			args: args{
				media: m1,
				index: 0,
			},
			m:    nil,
			want: Medias{m1},
		},
		{
			name: "add to top of slice",
			args: args{
				media: m3,
				index: 0,
			},
			m:    Medias{m1, m2},
			want: Medias{m3, m1, m2},
		},
		{
			name: "add to bottom of slice",
			args: args{
				media: m3,
				index: 2,
			},
			m:    Medias{m1, m2},
			want: Medias{m1, m2, m3},
		},
		{
			name: "add to middle of slice",
			args: args{
				media: m2,
				index: 1,
			},
			m:    Medias{m1, m3},
			want: Medias{m1, m2, m3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.m.InsertAt(tt.args.media, tt.args.index)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.ElementsMatch(t, tt.want, got)
			}
		})
	}
}

func TestMedias_RemoveAt(t *testing.T) {
	type args struct {
		index int
	}
	tests := []struct {
		name    string
		m       Medias
		args    args
		want    Medias
		wantErr bool
	}{
		{
			name: "index out of bounds",
			args: args{
				index: 1,
			},
			m:       nil,
			wantErr: true,
		},
		{
			name: "index out of bounds",
			args: args{
				index: -1,
			},
			m:       nil,
			wantErr: true,
		},
		{
			name: "remove from empty slice",
			args: args{
				index: 0,
			},
			m:       nil,
			wantErr: true,
		},
		{
			name: "remove from top of slice",
			args: args{
				index: 0,
			},
			m:    Medias{m3, m1, m2},
			want: Medias{m1, m2},
		},
		{
			name: "remove from bottom of slice",
			args: args{
				index: 2,
			},
			m:    Medias{m1, m2, m3},
			want: Medias{m1, m2},
		},
		{
			name: "remove from middle of slice",
			args: args{
				index: 1,
			},
			m:    Medias{m1, m2, m3},
			want: Medias{m1, m3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.m.RemoveAt(tt.args.index)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.ElementsMatch(t, tt.want, got)
			}
		})
	}
}

func TestMedias_InsertTop(t *testing.T) {
	type args struct {
		media *protos.Media
	}
	tests := []struct {
		name string
		m    Medias
		args args
		want Medias
	}{
		{
			name: "one item",
			args: args{media: m1},
			m:    nil,
			want: Medias{m1},
		},
		{
			name: "two items",
			args: args{media: m1},
			m:    Medias{m2},
			want: Medias{m1, m2},
		},
		{
			name: "three items",
			args: args{media: m1},
			m:    Medias{m2, m3},
			want: Medias{m1, m2, m3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.m.InsertTop(tt.args.media)
			assert.ElementsMatch(t, tt.want, got)
		})
	}
}

func TestMedias_InsertBottom(t *testing.T) {
	type args struct {
		media *protos.Media
	}
	tests := []struct {
		name string
		m    Medias
		args args
		want Medias
	}{
		{
			name: "one item",
			args: args{media: m1},
			m:    nil,
			want: Medias{m1},
		},
		{
			name: "two items",
			args: args{media: m2},
			m:    Medias{m1},
			want: Medias{m1, m2},
		},
		{
			name: "three items",
			args: args{media: m3},
			m:    Medias{m1, m2},
			want: Medias{m1, m2, m3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.m.InsertBottom(tt.args.media)
			assert.ElementsMatch(t, tt.want, got)
		})
	}
}
