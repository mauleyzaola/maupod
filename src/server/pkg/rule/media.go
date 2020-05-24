package rule

import (
	"errors"
	"math"
	"os"
	"path/filepath"
	"time"

	"github.com/mauleyzaola/maupod/src/server/pkg/helpers"
	"github.com/mauleyzaola/maupod/src/server/pkg/media"
	"github.com/mauleyzaola/maupod/src/server/pkg/pb"
)

func NewMediaFile(info *media.MediaInfo, filename string, scanDate time.Time, fileInfo os.FileInfo) *pb.Media {
	media := info.ToProto()
	media.Id = helpers.NewUUID()
	//media.Sha = "" // needs to be defined in another process
	media.LastScan = helpers.TimeToTs(&scanDate)
	media.ModifiedDate = helpers.TimeToTs2(fileInfo.ModTime())
	media.Location = filename
	media.FileExtension = filepath.Ext(fileInfo.Name())

	return media
}

// FileInfo returns the information from the file system about a media item
func (m *Media) FileInfo() (os.FileInfo, error) {
	if m.Location == "" {
		return nil, errors.New("missing location")
	}
	return os.Stat(m.Location)
}

func (m *Media) FileIsValidExtension(c *pb.Configuration) bool {
	return FileIsValidExtension(c, m.Location)
}

// Needs update compares the file system modified date vs database value
func (m *Media) NeedsUpdate() bool {
	info, err := m.FileInfo()
	if err != nil {
		return false
	}
	lastUpdateDb, lastUpdateFileSystem := m.ModifiedDate.Seconds, info.ModTime().Unix()
	diffSeconds := math.Abs(float64(lastUpdateFileSystem - lastUpdateDb))
	// less than 1 second should not be considered as a change
	return diffSeconds > 1
}

// Needs update compares the file system modified date vs database value
func (m *Media) NeedsImageUpdate() bool {
	info, err := m.FileInfo()
	if err != nil {
		return false
	}
	if m.LastImageScan == nil {
		return true
	}
	lastUpdateDb, lastUpdateFileSystem := m.LastImageScan.Seconds, info.ModTime().Unix()
	diffSeconds := math.Abs(float64(lastUpdateFileSystem - lastUpdateDb))
	// less than 1 second should not be considered as a change
	return diffSeconds > 1
}

func (m *Media) HasImage() bool {
	return m.ShaImage != ""
}

func ImageFileName(m *pb.Media, store *pb.FileStore) (string, error) {
	if m == nil {
		return "", errors.New("missing parameter: m")
	}
	if store == nil {
		return "", errors.New("missing parameter: store")
	}
	if m.ShaImage == "" {
		return "", errors.New("missing sha image")
	}
	// TODO: allow to configure the image format, png is assumed for now
	filename := m.ShaImage + ".png"
	return filepath.Join(store.Location, filename), nil
}
