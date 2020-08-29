package artwork_helpers

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/mauleyzaola/maupod/src/pkg/images"

	"github.com/mauleyzaola/maupod/src/pkg/broker"
	"github.com/nats-io/nats.go"

	"github.com/mauleyzaola/maupod/src/pkg/helpers"
	"github.com/mauleyzaola/maupod/src/pkg/paths"
	"github.com/mauleyzaola/maupod/src/pkg/rules"

	"github.com/mauleyzaola/maupod/src/pkg/pb"
)

var ErrNoFileFound = errors.New("no file found")

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
// filename should be an existent absolute path to the image
// for the time being: at least 500x500 pixeles and to be square (same width and height)
func IsArtworkValidSize(nc *nats.Conn, filename string, minWidth int) error {
	output, err := broker.RequestMediaInfoScan(nc, filename, time.Second*3)
	if err != nil {
		return err
	}
	x, y, err := images.Size(bytes.NewBufferString(output.Raw))
	if err != nil {
		return err
	}

	// check symmetry
	if x != y {
		return fmt.Errorf("invalid image shape: %dx%d", x, y)
	}

	// check min width
	if x < minWidth {
		return fmt.Errorf("image too small: %dx%d", x, y)
	}
	return nil
}

// ArtworkFullPath returns the absolute path on this micro service for an artwork file
func ArtworkFullPath(config *pb.Configuration, media *pb.Media) string {
	filename := filepath.Join(config.ArtworkStore.Location, rules.ArtworkFileName(media))
	return paths.LocationPath(filename)
}

// ArtworkSave will save the artwork from a reader, based on the media file provided
func ArtworkSave(media *pb.Media, reader io.Reader) error {
	panic("not implemented")
}

// ArtworkFileExist will check if the artwork file already exists
func ArtworkFileExist(config *pb.Configuration, media *pb.Media) bool {
	filename := filepath.Join(config.ArtworkStore.Location, rules.ArtworkFileName(media))
	_, err := os.Stat(filename)
	return err == nil
}

// FindArtworkInDirectory will return the first matching filename which is a valid artwork file
func FindArtworkInDirectory(media *pb.Media) (*string, error) {
	const (
		pngExt = ".png"
		jpgExt = ".jpg"
	)

	dir := filepath.Dir(paths.FullPath(media.Location))
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil, err
	}
	var validImageFiles = helpers.StringSlice([]string{pngExt, jpgExt})
	var matchedFile os.FileInfo
	var ext string
	for _, v := range files {
		ext = filepath.Ext(v.Name())
		ext = strings.ToLower(ext)
		if !validImageFiles.ContainsAny(ext) {
			continue
		}
		matchedFile = v
		break
	}
	if matchedFile == nil {
		return nil, ErrNoFileFound
	}

	var result = filepath.Join(dir, matchedFile.Name())
	return &result, nil
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

func ArtworkResizeFile(source, target string, imageSize int) error {
	var bigSize = imageSize
	err := images.ImageResize(source, target, bigSize, bigSize)
	if err != nil {
		return err
	}
	return nil
}
