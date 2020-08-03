package paths

import (
	"os"
	"path/filepath"
	"strings"
)

// RootDirectory returns the root directory for calculating the location on the media files
func RootDirectory() string {
	return os.Getenv("MEDIA_STORE")
}

// FullPath returns the full path to a media file location, based on the root specified in the environment
func FullPath(location string) string {
	return filepath.Join(RootDirectory(), location)
}

func LocationPath(fullPath string) string {
	return strings.TrimPrefix(fullPath, RootDirectory())
}
