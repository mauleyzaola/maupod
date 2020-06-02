package broker

import (
	"errors"
	"time"

	"github.com/mauleyzaola/maupod/src/server/pkg/pb"
	"github.com/mauleyzaola/maupod/src/server/pkg/types"
	"github.com/nats-io/nats.go"
)

func RequestScanAudioFile(nc *nats.Conn, logger types.Logger, filename string, timeout time.Duration) (*pb.Media, error) {
	input := &pb.MediaInfoInput{FileName: filename}
	output, err := mediaInfoRequest(nc, input, timeout)
	if err != nil {
		// TODO: send the files with errors to another listener and store in db
		logger.Error(err)
		return nil, err
	}
	if output.Response == nil {
		return nil, errors.New("missing response")
	}
	if !output.Response.Ok {
		return nil, errors.New(output.Response.Error)
	}

	return input.Media, nil
}
