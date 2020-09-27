package taggers

import (
	"fmt"
	"strconv"

	"github.com/mauleyzaola/maupod/src/pkg/helpers"

	"github.com/mauleyzaola/maupod/src/pkg/pb"
)

const flacTagProgram = "metaflac"

type FLACTagger struct{}

func (t *FLACTagger) Tag(media *pb.Media, filename string) error {
	var params []string
	addParam := func(name, value string) {
		params = append(params, fmt.Sprintf("--remove-tag=%s", name))
		params = append(params, fmt.Sprintf("--set-tag=%s=%s", name, value))
	}
	if media.Album != "" {
		addParam("ALBUM", media.Album)
	}
	if media.Performer != "" {
		addParam("ARTIST", media.Performer)
	}
	if media.Track != "" {
		addParam("TITLE", media.Track)
	}
	if media.TrackPosition != 0 {
		addParam("TRACKNUMBER", strconv.Itoa(int(media.TrackPosition)))
	}
	if media.RecordedDate != 0 {
		addParam("DATE", strconv.Itoa(int(media.RecordedDate)))
	}
	if media.Genre != "" {
		addParam("GENRE", media.Genre)
	}
	if media.Comment != "" {
		addParam("COMMENT", media.Comment)
	}
	if media.TrackNameTotal != 0 {
		addParam("TOTALTRACKS", strconv.Itoa(int(media.TrackNameTotal)))
	}
	if err := run(flacTagProgram, filename, params...); err != nil {
		return err
	}
	return nil
}

func (t *FLACTagger) RemoveAll(filename string) error {
	if err := run(flacTagProgram, filename, "--remove-all"); err != nil {
		return err
	}
	return nil
}

func (t *FLACTagger) ProgramExist() bool {
	return helpers.ProgramExists(flacTagProgram)
}
