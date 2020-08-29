package main

import (
	"log"
	"strconv"

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
		if artworkExists = ArtworkFileExist(m.config, track); artworkExists {
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
	var artworkFullPath = ArtworkFullPath(m.config, media)
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

	// update image in db for each track
	for _, track := range albumTracks {
		track.ImageLocation = media.ImageLocation
		if err = PublishSaveArtworkTrack(m.base.NATS(), track); err != nil {
			log.Println(err)
			return
		}
	}
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
	if ArtworkFileExist(m.config, input.Media) {
		log.Printf("track: %s album: %s already exists artwork\n", input.Media.Track, input.Media.Album)
		input.Media.ImageLocation = rules.ArtworkFileName(input.Media)
		return
	}

	var artworkPath = ArtworkFullPath(m.config, input.Media)
	var coverLocation string
	// check for the image in the same directory of the audio file
	coverFile, err := FindArtworkInDirectory(input.Media)
	if err != nil {
		log.Println(err)
		return
	}
	coverLocation = *coverFile

	// check shape and size are valid
	minWidth := int(m.config.ArtworkBigSize)
	if err = IsArtworkValidSize(m.base.NATS(), coverLocation, minWidth); err != nil {
		log.Println(err)
		return
	}

	if err = ArtworkResizeFile(coverLocation, artworkPath, minWidth); err != nil {
		log.Println(err)
		return
	}

	// if we got this far, assign the artwork value to the track
	input.Media.ImageLocation = rules.ArtworkFileName(input.Media)
	// TODO: write to db the changes
	return
}
