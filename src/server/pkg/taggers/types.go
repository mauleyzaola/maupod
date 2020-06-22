package taggers

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/mauleyzaola/maupod/src/server/pkg/pb"
)

type Tagger interface {
	Tag(media *pb.Media, filename string) error
	RemoveAll(filename string) error
}

func TaggerFactory(filename string) (Tagger, error) {
	ext := filepath.Ext(filename)
	switch strings.ToLower(ext) {
	case ".flac":
		return &FLACTagger{}, nil
	case ".mp3":
		return &MP3Tagger{}, nil
	default:
		return nil, fmt.Errorf("unsupported tagging for extension: %s", ext)
	}
}
