package taggers

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/mauleyzaola/maupod/src/protos"
)

type Tagger interface {
	ProgramExist() bool
	RemoveAll(filename string) error
	Tag(media *protos.Media, filename string) error
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
