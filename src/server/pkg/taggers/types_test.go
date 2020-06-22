package taggers

import (
	"reflect"
	"testing"
)

func TestTaggerFactory(t *testing.T) {
	type args struct {
		filename string
	}
	tests := []struct {
		name    string
		args    args
		want    Tagger
		wantErr bool
	}{
		{
			args:    args{filename: "something.m4a"},
			wantErr: true,
			want:    nil,
		},
		{
			args:    args{filename: "something.mp3"},
			wantErr: false,
			want:    &MP3Tagger{},
		},
		{
			args:    args{filename: "something.flac"},
			wantErr: false,
			want:    &FLACTagger{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := TaggerFactory(tt.args.filename)
			if (err != nil) != tt.wantErr {
				t.Errorf("TaggerFactory() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TaggerFactory() got = %v, want %v", got, tt.want)
			}
		})
	}
}
