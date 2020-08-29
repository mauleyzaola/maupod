package main

import (
	"io/ioutil"
	"log"
	"path/filepath"

	"github.com/mauleyzaola/maupod/src/pkg/helpers"
	"github.com/mauleyzaola/maupod/src/pkg/paths"
	"github.com/mauleyzaola/maupod/src/pkg/pb"
	"github.com/mauleyzaola/maupod/src/pkg/rules"
	"github.com/nats-io/nats.go"
)

func (m *MsgHandler) handlerReadDirectory(msg *nats.Msg) {
	var input pb.DirectoryReadInput
	var err error
	var output pb.DirectoryReadOutput

	defer func() {
		data, err := helpers.ProtoMarshal(&output)
		if err != nil {
			log.Println(err)
			return
		}
		if err = msg.Respond(data); err != nil {
			log.Println(err)
		}
	}()

	if err = helpers.ProtoUnmarshal(msg.Data, &input); err != nil {
		log.Println(err)
		output.Error = err.Error()
		return
	}
	var root = paths.MediaFullPathAudioFile(input.Root)
	infos, err := ioutil.ReadDir(root)
	if err != nil {
		output.Error = err.Error()
		return
	}
	for _, info := range infos {
		if !info.IsDir() && !rules.FileIsValidExtension(m.config, info.Name()) {
			continue
		}
		var file = &pb.FileItem{
			Location: paths.LocationPath(filepath.Join(root, info.Name())),
			IsDir:    info.IsDir(),
			Size:     info.Size(),
			Name:     info.Name(),
			Id:       helpers.NewUUID(),
		}
		output.Files = append(output.Files, file)
	}
	return
}
