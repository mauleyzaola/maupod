package main

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"time"

	"github.com/mauleyzaola/maupod/src/pkg/broker"

	"github.com/mauleyzaola/maupod/src/cmd/artwork/pkg/artworks"
	"github.com/mauleyzaola/maupod/src/pkg/helpers"
	"github.com/mauleyzaola/maupod/src/pkg/pb"
	"github.com/mauleyzaola/maupod/src/pkg/rules"
	"github.com/nats-io/nats.go"
)

func (m *MsgHandler) handlerArtworkDownload(msg *nats.Msg) {
	var err error
	var input pb.ArtworkDownloadInput
	var output pb.ArtworkDownloadOutput

	defer func() {
		if msg.Reply == "" {
			return
		}
		var data []byte
		data, err = helpers.ProtoMarshal(&output)
		if err != nil {
			log.Println(err)
			return
		}
		if err = msg.Respond(data); err != nil {
			log.Println(err)
			return
		}
		return
	}()

	if err = helpers.ProtoUnmarshal(msg.Data, &input); err != nil {
		output.Error = err.Error()
		return
	}
	if input.Uri == "" {
		output.Error = "missing artwork uri"
		return
	}
	if input.AlbumIdentifier == "" {
		output.Error = "missing album identifier"
		return
	}

	//  download the image from the uri
	tmpFileName := filepath.Join(os.TempDir(), helpers.NewUUID())
	if err = downloadFile(input.Uri, tmpFileName); err != nil {
		log.Println(err)
		output.Error = err.Error()
		return
	}
	data, err := ioutil.ReadFile(tmpFileName)
	if err != nil {
		log.Println(err)
		output.Error = err.Error()
		return
	}
	defer func() {
		if err = os.Remove(tmpFileName); err != nil {
			log.Println(err)
		}
	}()
	var nc = m.base.NATS()

	// check file is valid and store image location
	var baseMedia = &pb.Media{AlbumIdentifier: input.AlbumIdentifier}

	// send message for each media in this album
	medium, err := broker.RequestMediaInfoScanFromDB(nc, &pb.MediaInfoInput{
		Media: baseMedia,
	}, rules.Timeout(m.config))
	if err != nil {
		log.Println(err)
		output.Error = err.Error()
		return
	}
	log.Printf("[DEBUG] found %d tracks in album\n", len(medium.Medias))
	if len(medium.Medias) == 0 {
		err = errors.New("no tracks found in album")
		log.Println(err)
		output.Error = err.Error()
		return
	}
	var newImageLocation = rules.ArtworkFileName(baseMedia)
	for i, media := range medium.Medias {
		// no need to write the same file for each track
		if i != 0 {
			continue
		}
		if err = artworks.UpdateArtworkCoverFile(nc, m.config, media, data, input.Force); err != nil {
			log.Println(err)
			output.Error = err.Error()
			return
		}
	}
	// send message for updating db
	var now = time.Now()
	for _, media := range medium.Medias {
		media.ImageLocation = newImageLocation
		media.LastImageScan = helpers.TimeToTs2(now)
		if err = broker.PublishMediaArtworkUpdate(nc, media); err != nil {
			log.Println(err)
			output.Error = err.Error()
			return
		}
	}
}

func downloadFile(uri, destination string) error {
	log.Println("querying provider for image at uri: ", uri)
	const programName = "wget"
	if !helpers.ProgramExists(programName) {
		return fmt.Errorf("could not find program: %s in path", programName)
	}
	var p = []string{
		uri,
		"-O",
	}
	p = append(p, destination)
	cmd := exec.Command(programName, p...)
	output := &bytes.Buffer{}
	errOutput := &bytes.Buffer{}
	cmd.Stdout = output
	cmd.Stderr = errOutput
	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("%s %s : %v", output.String(), errOutput.String(), err)
	}
	return nil
}
