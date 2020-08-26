package artwork_helpers

import (
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"

	"github.com/mauleyzaola/maupod/src/pkg/broker"
	"github.com/nats-io/nats.go"

	"github.com/mauleyzaola/maupod/src/pkg/helpers"
	"github.com/mauleyzaola/maupod/src/pkg/paths"
	"github.com/mauleyzaola/maupod/src/pkg/rules"

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

// FindFirstTrackSubdirectories will return all the sibling directories it can find from the root
// all of which should contain at least one track with a valid extension, based in provided configuration
func FindFirstTrackSubdirectories(config *pb.Configuration, root string) ([]string, error) {
	var dirFirstTrackMap = make(map[string]string)
	var files []string

	fn := func(name string, isDir bool) (stop bool) {
		var location = paths.LocationPath(name)
		var dir = filepath.Dir(location)
		if !rules.FileIsValidExtension(config, location) {
			return false
		}
		if _, ok := dirFirstTrackMap[dir]; ok {
			return false
		}
		dirFirstTrackMap[dir] = location
		files = append(files, location)
		return false
	}
	if err := helpers.WalkFiles(root, fn); err != nil {
		log.Println(err)
		return nil, err
	}
	return files, nil
}

// LookupAlbumTracks will return all the media tracks from the same album
func LookupAlbumTracks(nc *nats.Conn, config *pb.Configuration, media *pb.Media) ([]*pb.Media, error) {
	var mediaInfoInput = &pb.MediaInfoInput{
		FileName: media.Location,
		Media:    media,
	}
	mediaInfoOutput, err := broker.RequestMediaInfoScanFromDB(nc, mediaInfoInput, rules.Timeout(config))
	if err != nil {
		log.Println(err)
		return nil, err
	}
	medias := mediaInfoOutput.Medias
	if len(medias) == 0 {
		return nil, errors.New("no media found with provided filters")
	} else if len(medias) != 1 {
		return nil, fmt.Errorf("more than one media found: %d with provided filters", len(medias))
	}
	media = medias[0]
	mediaInfoInput = &pb.MediaInfoInput{
		Media: &pb.Media{AlbumIdentifier: media.AlbumIdentifier},
	}
	mediaInfoOutput, err = broker.RequestMediaInfoScanFromDB(nc, mediaInfoInput, rules.Timeout(config))
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return mediaInfoOutput.Medias, nil
}

func LookupEmbeddedArtwork(media *pb.Media) (ReadDestroyer, error) {
	panic("not implemented")
}

func SaveArtwork() {
	panic("not implemented")
}

func PublishSaveArtworkTrack(nc *nats.Conn, media *pb.Media) error {
	return broker.PublishMediaArtworkUpdate(nc, media)
}
