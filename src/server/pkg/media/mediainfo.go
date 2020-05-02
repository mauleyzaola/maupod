package media

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os/exec"
	"strconv"
	"strings"
	"time"

	"github.com/mauleyzaola/maupod/src/server/pkg/domain"
)

const (
	mediaInfoDateFormat = "UTC 2006-01-02 15:04:05"
)

// Mediainfo returns a MediaInfo slice
// params can be either one file, many files or even many paths, not necessary
// pointing to specific audio files, but directories that contain audio files within
func Mediainfo(params ...string) ([]MediaInfo, error) {
	const mediaInfoProgram = "mediainfo"
	if !programExists(mediaInfoProgram) {
		return nil, fmt.Errorf("could not find program: %s in path", mediaInfoProgram)
	}
	var p = []string{
		"--Output=JSON",
	}
	p = append(p, params...)
	cmd := exec.Command(mediaInfoProgram, p...)
	output := &bytes.Buffer{}
	errOutput := &bytes.Buffer{}
	cmd.Stdout = output
	cmd.Stderr = errOutput
	err := cmd.Run()
	if err != nil {
		return nil, err
	}

	// mediainfo when using JSON format, can return one object or an array, need to check both options
	var payload MediaInfo
	var payloads []MediaInfo
	if err = json.Unmarshal(output.Bytes(), &payload); err == nil {
		return []MediaInfo{payload}, nil
	}
	if err = json.Unmarshal(output.Bytes(), &payloads); err != nil {
		return nil, err
	}

	return payloads, nil
}

func programExists(programName string) bool {
	_, err := exec.LookPath(programName)
	return err == nil
}

func (m *MediaInfo) findTrack(trackType string) *Track {
	for i := range m.Media.Track {
		track := &m.Media.Track[i]
		if strings.ToLower(track.Type) == strings.ToLower(trackType) {
			return track
		}
	}
	return nil
}

func (m *MediaInfo) ToDomain() *domain.Media {
	res := &domain.Media{}

	if a := m.findTrack("audio"); a != nil {
		res.Format = a.Format
		res.Duration, _ = strconv.ParseFloat(a.Duration, 10)
		res.BitRate, _ = strconv.ParseInt(a.BitRate, 10, 64)
		res.BitRateMode = a.BitRateMode
		res.Channels = a.Channels
		res.ChannelPositions = a.ChannelPositions
		res.ChannelLayout = a.ChannelLayout
		res.SamplingRate, _ = strconv.ParseInt(a.SamplingRate, 10, 64)
		res.SamplingCount, _ = strconv.ParseInt(a.SamplingCount, 10, 64)
		res.BitDepth, _ = strconv.ParseInt(a.BitDepth, 10, 64)
		res.CompressionMode = a.CompressionMode
		res.StreamSize, _ = strconv.ParseInt(a.StreamSize, 10, 64)
		res.EncodedLibrary = a.EncodedLibrary
		res.EncodedLibraryName = a.EncodedLibraryName
		res.EncodedLibraryVersion = a.EncodedLibraryVersion
	}

	if g := m.findTrack("general"); g != nil {
		res.FileExtension = g.FileExtension
		if res.Format == "" {
			res.Format = g.Format
		}
		res.FileSize, _ = strconv.ParseInt(g.FileSize, 10, 64)
		if res.Duration == 0 {
			res.Duration, _ = strconv.ParseFloat(g.Duration, 64)
		}
		res.OverallBitRate, _ = strconv.ParseInt(g.OverallBitRate, 10, 64)
		res.OverallBitRateMode = g.OverallBitRateMode
		res.StreamSize, _ = strconv.ParseInt(g.StreamSize, 10, 64)
		res.Album = g.Album
		res.Track = g.Track
		res.Title = g.Title
		res.TrackPosition, _ = strconv.ParseInt(g.TrackPosition, 10, 64)
		res.Performer = g.Performer
		res.Genre = g.Genre
		res.RecordedDate, _ = strconv.ParseInt(g.RecordedDate, 10, 64)
		res.Comment = g.Comment
		res.FileModifiedDate, _ = time.Parse(mediaInfoDateFormat, g.FileModifiedDate)
	}

	return res
}
