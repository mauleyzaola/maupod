package rules

import (
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/mauleyzaola/maupod/src/pkg/helpers"
	"github.com/mauleyzaola/maupod/src/pkg/paths"

	"github.com/mauleyzaola/maupod/src/protos"
)

// FileInfo returns the information from the file system about a media item
func FileInfo(m *protos.Media) (os.FileInfo, error) {
	if m.Location == "" {
		return nil, errors.New("missing location")
	}
	var fullPath = paths.MediaFullPathAudioFile(m.Location)
	return os.Stat(fullPath)
}

// Needs update compares the file system modified date vs database value
func NeedsMediaUpdate(m *protos.Media) bool {
	info, err := FileInfo(m)
	if err != nil {
		log.Println("[ERROR] ", err)
		return false
	}
	if m.LastScan == nil {
		return true
	}

	var lastScan = helpers.TsToTime(m.LastScan)
	return lastScan.Before(info.ModTime())
}

func MediaCheckMinimalData(m *protos.Media) error {
	if m.Album == "" {
		return errors.New("media missing: album")
	}
	if m.Track == "" {
		return errors.New("media missing: track")
	}
	if m.Performer == "" {
		return errors.New("media missing: performer")
	}
	return nil
}

// MediaPercentToSeconds converts the percent of the track to seconds played
func MediaPercentToSeconds(m *protos.Media, percent float64) (*time.Duration, error) {
	if percent < 0 || percent > 100 {
		return nil, fmt.Errorf("percent out of range: %v", percent)
	}
	if m.Duration == 0 {
		return nil, errors.New("missing duration, cannot calculate percent")
	}
	var percentPlayed = m.Duration / percent
	var duration = time.Millisecond * time.Duration(percentPlayed)
	return &duration, nil
}

func MediaTotalSeconds(m *protos.Media) (*time.Duration, error) {
	if m.Duration == 0 {
		return nil, errors.New("missing duration, cannot calculate percent")
	}
	var duration = time.Millisecond * time.Duration(m.Duration)
	return &duration, nil
}
