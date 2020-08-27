package main

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	. "github.com/mauleyzaola/maupod/src/cmd/artwork/pkg/artwork_helpers"
	"github.com/mauleyzaola/maupod/src/pkg/broker"
	"github.com/mauleyzaola/maupod/src/pkg/helpers"
	"github.com/mauleyzaola/maupod/src/pkg/images"
	"github.com/mauleyzaola/maupod/src/pkg/paths"
	"github.com/mauleyzaola/maupod/src/pkg/pb"
	"github.com/mauleyzaola/maupod/src/pkg/rules"
	"github.com/nats-io/nats.go"
)

func (m *MsgHandler) handlerArtworkExtractDirectories(msg *nats.Msg) {
	var err error
	var input pb.ArtworkExtractInput
	if err = helpers.ProtoUnmarshal(msg.Data, &input); err != nil {
		log.Println(err)
		return
	}

	trackFiles, err := FindFirstTrackSubdirectories(m.config, input.Root)
	if err != nil {
		log.Println(err)
		return
	}
	log.Printf("found %d directories with valid files\n", len(trackFiles))

	for _, trackFile := range trackFiles {
		fileInput := &pb.ArtworkExtractInput{
			Root: trackFile,
		}
		if err = broker.PublishBroker(m.base.NATS(), pb.Message_MESSAGE_MEDIA_EXTRACT_ARTWORK_FROM_FILE, fileInput); err != nil {
			log.Println(err)
			return
		}
	}
}

// handlerArtworkExtractWithinAudioFiles will try to look up into audio files content for images
func (m *MsgHandler) handlerArtworkExtractWithinAudioFiles(msg *nats.Msg) {
	var err error
	var input pb.ArtworkExtractInput
	if err = helpers.ProtoUnmarshal(msg.Data, &input); err != nil {
		log.Println(err)
		return
	}

	if input.Media == nil {
		input.Media = &pb.Media{}
	}

	var media = input.Media
	if input.Root != "" {
		media.Location = input.Root
	}

	albumTracks, err := LookupAlbumTracks(m.base.NATS(), m.config, media)
	if err != nil {
		log.Println(err)
		return
	}

	if len(albumTracks) == 0 {
		log.Println("[WARNING] album has no tracks")
		return
	}

	// TODO: extract this logic to a function and reuse in other similar message receivers
	var artworkUpdate = struct {
		ImageLocation string
	}{}
	var artworkExists bool
	// search for any image location on any track of the same album
	for _, track := range albumTracks {
		if artworkExists = ArtworkFileExist(track); artworkExists {
			// assign to media in case we need to process the artwork below
			media = track
			artworkUpdate.ImageLocation = track.ImageLocation
			break
		}
	}

	// if artwork was found for any track, update the other tracks with the same artwork info and exit
	if artworkExists {
		for _, track := range albumTracks {
			track.ImageLocation = artworkUpdate.ImageLocation
			if err = PublishSaveArtworkTrack(m.base.NATS(), track); err != nil {
				log.Println(err)
				return
			}
		}
		log.Println(err)
	}
	// END: extract this logic to a function and reuse in other similar message receivers

	// no other tracks were found with artwork, then scan audio file and search for images
	var mediaFullPath = paths.FullPath(media.Location)
	var artworkFullPath = artworkFullPath(m.config, media)
	if err = images.ExtractImageFromMedia(mediaFullPath, artworkFullPath); err != nil {
		log.Println(err)
		return
	}
	// check for image size to be ok
	if err = IsArtworkValidSize(m.base.NATS(), artworkFullPath, int(m.config.ArtworkBigSize)); err != nil {
		log.Println(err)
		return
	}

	// if all went fine, create the artwork image
	if err = ArtworkResizeFile(artworkFullPath, artworkFullPath, int(m.config.ArtworkBigSize)); err != nil {
		log.Println(err)
		return
	}

	// if we got this far, assign the artwork value to the track
	media.ImageLocation = rules.ArtworkFileName(media)

	// TODO: update image in db for each track
	// TODO: assign image location to the rest of the tracks
}

// handlerArtworkExtract this will only look for image files in the same directory of the audio files
// no scanning of audio files content should be done
func (m *MsgHandler) handlerArtworkExtract(msg *nats.Msg) {
	var err error
	var input pb.ArtworkExtractInput
	if err = helpers.ProtoUnmarshal(msg.Data, &input); err != nil {
		log.Println(err)
		return
	}

	//we need to update the database when this function exits, one way or another
	defer func() {
		var payload []byte
		if err != nil {
			return
		}
		if payload, err = helpers.ProtoMarshal(&pb.ArtworkUpdateInput{Media: input.Media}); err != nil {
			log.Println(err)
			return
		}
		if err = m.base.NATS().Publish(strconv.Itoa(int(pb.Message_MESSAGE_MEDIA_UPDATE_ARTWORK)), payload); err != nil {
			log.Println(err)
			return
		}
	}()

	input.Media.LastImageScan = helpers.TimeToTs(helpers.Now())
	if input.Media.AlbumIdentifier == "" {
		// the album identifier is a 1:1 match with the image file, if not present, cannot process artwork
		return
	}

	// check if artwork already exist for the same album
	if artworkFileExist(m.config, input.Media) {
		log.Printf("track: %s album: %s already exists artwork\n", input.Media.Track, input.Media.Album)
		input.Media.ImageLocation = rules.ArtworkFileName(input.Media)
		return
	}

	var artworkPath = artworkFullPath(m.config, input.Media)
	var coverLocation string
	// check for the image in the same directory of the audio file
	if coverLocation = findArtworkSameDirectory(paths.FullPath(input.Media.Location)); coverLocation == "" {
		return
	}

	// check shape and size are valid
	if err = imageValidSize(m.base.NATS(), coverLocation, int(m.config.ArtworkBigSize)); err != nil {
		log.Println(err)
		return
	}

	if err = imageWriteArtwork(coverLocation, artworkPath, int(m.config.ArtworkBigSize)); err != nil {
		log.Println(err)
		return
	}

	// if we got this far, assign the artwork value to the track
	input.Media.ImageLocation = rules.ArtworkFileName(input.Media)
	return
}

func artworkFileExist(config *pb.Configuration, media *pb.Media) bool {
	var artworkPath = artworkFullPath(config, media)
	_, err := os.Stat(artworkPath)
	return err == nil
}

func artworkFullPath(config *pb.Configuration, media *pb.Media) string {
	var dir = config.ArtworkStore.Location
	var imageLocation = rules.ArtworkFileName(media)
	return filepath.Join(dir, imageLocation)
}

func findArtworkSameDirectory(location string) string {
	const (
		pngExt = ".png"
		jpgExt = ".jpg"
	)

	dir := filepath.Dir(location)
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return ""
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
		return ""
	}

	return filepath.Join(dir, matchedFile.Name())
}
