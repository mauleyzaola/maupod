package types

import (
	"errors"

	"github.com/mauleyzaola/maupod/src/pkg/pb"
)

type Medias []*pb.Media

func (m Medias) InsertAt(media *pb.Media, index int) (Medias, error) {
	var res Medias
	if index < 0 || index > len(m) {
		return nil, errors.New("index out of bounds")
	}
	if len(m) == index {
		return append(m, media), nil
	}
	res = append(m[:index+1], m[index:]...)
	res[index] = media
	return res, nil
}

func (m Medias) InsertTop(media *pb.Media) Medias {
	res, _ := m.InsertAt(media, 0)
	return res
}

func (m Medias) InsertBottom(media *pb.Media) Medias {
	var index = len(m) - 1
	if index == -1 {
		index++
	}
	res, _ := m.InsertAt(media, index)
	return res
}

func (m Medias) RemoveAt(index int) (Medias, error) {
	if index < 0 || index > len(m)-1 {
		return nil, errors.New("index out of bounds")
	}

	return append(m[:index], m[index+1:]...), nil
}
