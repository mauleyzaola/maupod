package taggers

import (
	"fmt"
	"strconv"

	"github.com/mauleyzaola/maupod/src/pkg/helpers"

	"github.com/mauleyzaola/maupod/src/protos"
)

const mp3TagProgram = "id3v2"

type MP3Tagger struct{}

func (t *MP3Tagger) Tag(media *protos.Media, filename string) error {
	var params []string
	addParam := func(name, value string) {
		params = append(params, fmt.Sprintf("%s=%s", name, value))
	}
	if media.Album != "" {
		addParam("--album", media.Album)
	}
	if media.Performer != "" {
		addParam("--artist", media.Performer)
	}
	if media.Track != "" {
		addParam("--song", media.Track)
	}
	if media.TrackPosition != 0 || media.TrackNameTotal != 0 {
		addParam("--track", fmt.Sprintf("%v/%v", media.TrackPosition, media.TrackNameTotal))
	}
	if media.RecordedDate != 0 {
		addParam("--year", strconv.Itoa(int(media.RecordedDate)))
	}
	if media.Genre != "" {
		addParam("--genre", media.Genre)
	}
	if media.Comment != "" {
		addParam("--comment", fmt.Sprintf(":%s", media.Comment))
	}

	if err := run(mp3TagProgram, filename, params...); err != nil {
		return err
	}
	return nil
}

func (t *MP3Tagger) RemoveAll(filename string) error {
	if err := run(mp3TagProgram, filename, "-D"); err != nil {
		return err
	}
	return nil
}

func (t *MP3Tagger) ProgramExist() bool {
	return helpers.ProgramExists(mp3TagProgram)
}
