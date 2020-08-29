// artworks contains the atomic methods for dealing with artwork related stuff
// all paths received in the functions should be absolute, unless a media object is received
// in which case we would infer it based on environment variables for artwork and media audio files location
//
package artworks

import (
	"bytes"
	"errors"
	"fmt"
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
	var validImageFiles = helpers.StringSlice([]string{pngExt, jpgExt})
	var ext string
	var result []string
	for _, v := range files {
		ext = filepath.Ext(v.Name())
		ext = strings.ToLower(ext)
		if !validImageFiles.ContainsAny(ext) {
			continue
		}
		result = append(result, filepath.Join(dir, v.Name()))
	}

	return result, nil
}

// FindFirstTrackSubdirectories will return all the sibling directories it can find from the root
// all of which should contain at least one track with a valid extension, based in provided configuration
func FindFirstTrackSubdirectories(config *pb.Configuration, root string) ([]string, error) {
	var dirFirstTrackMap = make(map[string]string)
	var files []string

	fn := func(name string, isDir bool) (stop bool) {
		var dir = filepath.Dir(name)
		if !rules.FileIsValidExtension(config, name) {
			return false
		}
		if _, ok := dirFirstTrackMap[dir]; ok {
			return false
		}
		dirFirstTrackMap[dir] = name
		files = append(files, name)
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

// extractArtworkFromAudioFile will try to extract the artwork image from an audio file
// and store in a temp location which is returned as filename if success
func ExtractArtworkFromAudioFile(nc *nats.Conn, config *pb.Configuration, audioFile string) (filename string, err error) {
	filename = filepath.Join(os.TempDir(), helpers.NewUUID()+".png")
	if err = images.ExtractImageFromMedia(audioFile, filename); err != nil {
		filename = ""
		return
	}
	return
}

// ExtractWithinAudioFile will try to extract the image from the audio file
func ExtractWithinAudioFile(nc *nats.Conn, config *pb.Configuration, media *pb.Media) error {
	// TODO: extract this logic to a function and reuse in other similar message receivers
	albumTracks, err := LookupAlbumTracks(nc, config, media)
	if err != nil {
		log.Println(err)
		return err
	}

	if len(albumTracks) == 0 {
		err = errors.New("[ERROR] album has no tracks")
		log.Println(err)
		return err
	}

	var artworkUpdate = struct {
		ImageLocation string
	}{}
	var artworkExists bool
	// search for any image location on any track of the same album
	for _, track := range albumTracks {
		if artworkExists = ArtworkFileExist(config, track); artworkExists {
			// assign to media in case we need to process the artwork below
			media = track
			artworkUpdate.ImageLocation = track.ImageLocation
			break
		}
	}

	// if artwork was found for any track, update the other tracks with the same artwork info and exit
	lastImageScanDate := helpers.TimeToTs(helpers.Now())
	if artworkExists {
		for _, track := range albumTracks {
			track.ImageLocation = artworkUpdate.ImageLocation
			track.LastImageScan = lastImageScanDate
			if err = PublishSaveArtworkTrack(nc, track); err != nil {
				log.Println(err)
				return err
			}
		}
		log.Println(err)
	}
	// END: extract this logic to a function and reuse in other similar message receivers

	// no other tracks were found with artwork, then scan audio file and search for images
	var mediaFullPath = paths.MediaFullPathAudioFile(media.Location)
	var artworkFullPath = ArtworkFullPath(config, media)
	if err = images.ExtractImageFromMedia(mediaFullPath, artworkFullPath); err != nil {
		log.Println(err)
		return err
	}
	// check for image size to be ok
	if err = IsArtworkValidSize(nc, artworkFullPath, int(config.ArtworkBigSize)); err != nil {
		log.Println(err)
		return err
	}

	// if all went fine, create the artwork image
	width := int(config.ArtworkBigSize)
	height := width
	if err = ArtworkResizeFile(artworkFullPath, artworkFullPath, width, height); err != nil {
		log.Println(err)
		return err
	}

	// if we got this far, assign the artwork value to the track
	media.ImageLocation = rules.ArtworkFileName(media)

	// update image in db for each track
	log.Println("[INFO] update image data for all album tracks")
	for _, track := range albumTracks {
		track.ImageLocation = media.ImageLocation
		track.LastImageScan = lastImageScanDate
		if err = PublishSaveArtworkTrack(nc, track); err != nil {
			log.Println(err)
			return err
		}
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

	defer func() {
		lastImageScanDate := helpers.TimeToTs(helpers.Now())
		media.LastImageScan = lastImageScanDate
		if media.ImageLocation != "" {
			var albumTracks []*pb.Media
			if albumTracks, err = LookupAlbumTracks(nc, config, media); err != nil {
				log.Println(err)
				return
			}
			for _, track := range albumTracks {
				track.ImageLocation = media.ImageLocation
				track.LastImageScan = lastImageScanDate
				if err = PublishSaveArtworkTrack(nc, track); err != nil {
					log.Println(err)
					return
				}
			}
		} else {
			if err = PublishSaveArtworkTrack(nc, media); err != nil {
				log.Println(err)
				return
			}
		}
	}()

	// check if artwork already exist for the same album in the /artwork directory
	if ArtworkFileExist(config, media) {
		log.Printf("track: %s album: %s already exists artwork\n", media.Track, media.Album)
		media.ImageLocation = rules.ArtworkFileName(media)
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

	// if we got this far, assign the artwork value to the track
	media.ImageLocation = rules.ArtworkFileName(media)
	return nil
}
