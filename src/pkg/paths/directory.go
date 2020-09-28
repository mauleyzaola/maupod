package paths

import (
	"os"
	"path/filepath"
	"strings"
)

const MediaStoreEnvName = "MAUPOD_MEDIA_STORE"
const SyncPathEnvName = "MAUPOD_SYNC_PATH"

// RootDirectory returns the root directory for calculating the location on the media files
func RootDirectory() string {
	return os.Getenv(MediaStoreEnvName)
}

// MediaFullPathAudioFile returns the full path to a media file location, based on the root specified in the environment
func MediaFullPathAudioFile(location string) string {
	return filepath.Join(RootDirectory(), location)
}

func LocationPath(fullPath string) string {
	return strings.TrimPrefix(fullPath, RootDirectory())
}

// SyncFullPath the path to the sync directory within the docker image
func SyncFullPath(location string) string {
	return filepath.Join(SyncRootDirectory(), location)
}

func SyncRootDirectory() string {
	return filepath.Join("/", "sync")
}
