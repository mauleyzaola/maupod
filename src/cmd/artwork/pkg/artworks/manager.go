// artworks contains the atomic methods for dealing with artwork related stuff
// all paths received in the functions should be absolute, unless a media object is received
// in which case we would infer it based on environment variables for artwork and media audio files location
//
package artworks

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

func ArtworkPathFromEnvironment() string {
	return os.Getenv("MAUPOD_ARTWORK")
}

// IsArtworkValidSize will return true if image complies with requested image features
// filename should be an existent absolute path to the image
// for the time being: at least 500x500 pixeles and to be square (same width and height)
func IsArtworkValidSize(nc *nats.Conn, filename string, minWidth int) error {
	// TODO: pass timeout as parameter
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

// ArtworkFullPath returns the absolute path on this micro service for an artwork file, based on a media object
func ArtworkFullPath(config *pb.Configuration, media *pb.Media) string {
	filename := filepath.Join(config.ArtworkStore.Location, rules.ArtworkFileName(media))
	return paths.LocationPath(filename)
}

// ArtworkFileExist will check if the artwork file already exists
func ArtworkFileExist(config *pb.Configuration, media *pb.Media) bool {
	_, err := os.Stat(ArtworkFullPath(config, media))
	return err == nil
}

// FindArtworkFilesInDirectory will return all the files that are candidates for artwork files in an audio directory
func FindArtworkFilesInDirectory(dir string) ([]string, error) {
	const (
		pngExt = ".png"
		jpgExt = ".jpg"
	)

	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil, err
	}
	var keys = make(map[string]struct{})
	keys[pngExt] = struct{}{}
	keys[jpgExt] = struct{}{}
	var ext string
	var result []string
	for _, v := range files {
		ext = filepath.Ext(v.Name())
		ext = strings.ToLower(ext)
		if _, ok := keys[ext]; !ok {
			continue
		}
		result = append(result, filepath.Join(dir, v.Name()))
	}

	return result, nil
}

// PublishSaveArtworkTrack will send a message to update the db with the artwork related info in the media
func PublishSaveArtworkTrack(nc *nats.Conn, media *pb.Media) error {
	return broker.PublishMediaArtworkUpdate(nc, media)
}

// ArtworkResizeFile will read the source file and write the target file as a square of imageSize
func ArtworkResizeFile(source, target string, width, height int) error {
	err := images.ImageResize(source, target, width, height)
	if err != nil {
		return err
	}
	return nil
}

// ExtractWithinAudioFile will try to extract the image from the audio file
func ExtractWithinAudioFile(nc *nats.Conn, config *pb.Configuration, media *pb.Media) error {
	var err error
	// search for any image location on any track of the same album
	artworkExists := ArtworkFileExist(config, media)

	// if artwork was found, return
	if artworkExists {
		log.Println("[INFO] artwork already exists")
		return nil
	}

	width := int(config.ArtworkBigSize)
	height := width
	var artworkFullPath = ArtworkFullPath(config, media)
	if err = images.ExtractImageFromMedia(paths.MediaFullPathAudioFile(media.Location), artworkFullPath); err != nil {
		log.Println(err)
		return err
	}
	if err = IsArtworkValidSize(nc, artworkFullPath, width); err != nil {
		log.Println("[ERROR] image is not valid size: ", err)
		if err = os.Remove(artworkFullPath); err != nil {
			log.Println(err)
		}
		return err
	}

	// if all went fine, create the artwork image
	if err = ArtworkResizeFile(artworkFullPath, artworkFullPath, width, height); err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func ExtractFromCoverFile(nc *nats.Conn, config *pb.Configuration, media *pb.Media) error {
	var err error

	media.LastImageScan = helpers.TimeToTs(helpers.Now())
	if media.AlbumIdentifier == "" {
		// the album identifier is a 1:1 match with the image file, if not present, cannot process artwork
		err = errors.New("[ERROR] no album identifier detected in media")
		return err
	}

	// check if artwork already exist for the same album in the /artwork directory
	if ArtworkFileExist(config, media) {
		log.Printf("track: %s album: %s already exists artwork\n", media.Track, media.Album)
		return nil
	}

	var artworkPath = ArtworkFullPath(config, media)
	var coverLocation string

	// check for the image in the same directory of the audio file
	dir := filepath.Dir(paths.MediaFullPathAudioFile(media.Location))
	coverFiles, err := FindArtworkFilesInDirectory(dir)
	if err != nil {
		log.Println(err)
		return err
	}

	// check shape and size are valid for each of the valid artwork files
	width := int(config.ArtworkBigSize)
	height := width
	for _, v := range coverFiles {
		if err = IsArtworkValidSize(nc, v, width); err != nil {
			log.Println(err)
			continue
		}
		coverLocation = v
		log.Println("[INFO] found valid artwork file: ", coverLocation)
		break
	}

	if coverLocation == "" {
		err = errors.New("could not find cover location")
		log.Println(err)
		return err
	}
	if err = ArtworkResizeFile(coverLocation, artworkPath, width, height); err != nil {
		log.Println(err)
		return err
	}

	return nil
}

// TODO: use this function for the other artwork methods as well
// UpdateArtworkCoverFile will try to update the artwork for one media file, if force is true it will overwrite the current artwork
func UpdateArtworkCoverFile(nc *nats.Conn, config *pb.Configuration, media *pb.Media, artworkData []byte, force bool) error {
	var err error

	if len(artworkData) == 0 {
		err = errors.New("artwork data is empty")
		log.Println(err)
		return err
	}

	media.LastImageScan = helpers.TimeToTs(helpers.Now())
	if media.AlbumIdentifier == "" {
		// the album identifier is a 1:1 match with the image file, if not present, cannot process artwork
		err = errors.New("[ERROR] no album identifier detected in media")
		return err
	}

	var artworkPath = ArtworkFullPath(config, media)
	var tmpArtworkFile = filepath.Join(config.ArtworkStore.Location, helpers.NewUUID()+".png")

	if !force {
		if _, err = os.Stat(artworkPath); err == nil {
			// artwork already exists, do nothing
			log.Println("artwork already exists")
			return nil
		}
	}

	if err = ioutil.WriteFile(tmpArtworkFile, artworkData, os.ModePerm); err != nil {
		log.Println(err)
		return err
	}
	defer func() {
		_ = os.Remove(tmpArtworkFile)
	}()

	// check shape and size are valid for each of the valid artwork files
	width := int(config.ArtworkBigSize)
	height := width
	if err = IsArtworkValidSize(nc, tmpArtworkFile, width); err != nil {
		log.Println(err)
		return err
	}

	artworkFile, err := os.OpenFile(artworkPath, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, os.ModePerm)
	if err != nil {
		log.Println(err)
		if err = os.Remove(artworkPath); err != nil {
			log.Println(err)
		}
		return err
	}

	log.Println("[DEBUG] storing file at path: ", artworkPath)
	if _, err = io.Copy(artworkFile, bytes.NewReader(artworkData)); err != nil {
		log.Println(err)
		return err
	}
	if err = artworkFile.Close(); err != nil {
		log.Println(err)
		return err
	}

	if err = ArtworkResizeFile(tmpArtworkFile, artworkPath, width, height); err != nil {
		log.Println(err)
		return err
	}

	return nil
}
