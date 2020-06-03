package rule

import (
	"errors"
	"os"

	"github.com/mauleyzaola/maupod/src/server/pkg/pb"
)

// FileInfo returns the information from the file system about a media item
func FileInfo(m *pb.Media) (os.FileInfo, error) {
	if m.Location == "" {
		return nil, errors.New("missing location")
	}
	return os.Stat(m.Location)
}

// Needs update compares the file system modified date vs database value
func NeedsMediaUpdate(m *pb.Media) bool {
	info, err := FileInfo(m)
	if err != nil {
		return false
	}
	if m.LastScan == nil {
		return true
	}
	diffSeconds := m.LastScan.Seconds - info.ModTime().Unix()
	return diffSeconds < 0
}

// Needs update compares the file system modified date vs database value
func NeedsImageUpdate(m *pb.Media) bool {
	if m.LastImageScan == nil {
		return true
	}
	if m.ModifiedDate == nil {
		return true
	}
	diffSeconds := m.LastImageScan.Seconds - m.ModifiedDate.Seconds
	return diffSeconds < 0
}

func MediaHasImage(m *pb.Media) bool {
	return m.ShaImage != "" && m.ImageLocation != ""
}
