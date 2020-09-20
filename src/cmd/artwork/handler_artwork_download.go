package main

import (
	"io/ioutil"
	"log"
	"net/http"
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
	client := &http.Client{
		Timeout: rules.Timeout(m.config),
	}
	response, err := client.Get(input.Uri)
	if err != nil {
		output.Error = err.Error()
		return
	}
	log.Println("querying provider for image at uri: ", input.Uri)
	defer func() {
		if err = response.Body.Close(); err != nil {
			log.Println(err)
			return
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
		output.Error = err.Error()
		return
	}
	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
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
		if err = artworks.UpdateArtworkCoverFile(nc, m.config, media, data); err != nil {
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
			output.Error = err.Error()
			return
		}
	}
}
