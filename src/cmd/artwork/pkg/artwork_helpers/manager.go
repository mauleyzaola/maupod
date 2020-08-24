package artwork_helpers

import (
	"io"
	"os"

	"github.com/mauleyzaola/maupod/src/pkg/pb"
)

type ReadDestroyer interface {
	io.ReadCloser
	Destroy() error
}

func ArtworkPathFromEnvironment() string {
	return os.Getenv("MAUPOD_ARTWORK")
}

func MediaFullPath(media *pb.Media) string {
	panic("not implemented")
}

// SearchCoverFile will try to find a matching image in the directory
// based on a set of provided patterns
// will return the first match
// dir should be an absolute path to the audio files directory
func SearchCoverFile(dir string, patterns []string) {
	//filepath.Match()
	panic("not implemented")
}

// ExtractArtworkFromAudioFile will look up into the audio file for artwork
// image will be stored to a temporary location, and returned to the caller
func ExtractArtworkFromAudioFile(media *pb.Media) (io.ReadCloser, error) {
	panic("not implemented")
}

// IsArtworkValidSize will return true if image complies with requested image features
// for the time being: at least 500x500 pixeles and to be square (same width and height)
func IsArtworkValidSize(config *pb.Configuration, reader io.Reader) (ok bool, err error) {
	panic("not implemented")
}

// ArtworkFullPath returns the absolute path on this micro service for an artwork file
func ArtworkFullPath(media *pb.Media) string {
	panic("not implemented")
}

// ArtworkSave will save the artwork from a reader, based on the media file provided
func ArtworkSave(media *pb.Media, reader io.Reader) error {
	panic("not implemented")
}

// ArtworkFileExist will check if the artwork file already exists
func ArtworkFileExist(config *pb.Configuration, media *pb.Media) bool {
	panic("not implemented")
}

// FindArtworkInDirectory will return true if the artwork file for the media, exists
func FindArtworkInDirectory(media *pb.Media) bool {
	panic("not implemented")
}

// FindFirstTrackSubdirectories will return all the slibing directories it can find from the root
// all of which should contain at least one track
func FindFirstTrackSubdirectories(config *pb.Configuration, root string) ([]string, error) {
	panic("not implemented")
}

// LookupTracks will return all the media tracks from the same album
func LookupTracks(media *pb.Media) ([]*pb.Media, error) {
	panic("not implemented")
}

func LookupEmbeddedArtwork(media *pb.Media) (ReadDestroyer, error) {
	panic("not implemented")
}

func SaveArtwork() {
	panic("not implemented")
}
