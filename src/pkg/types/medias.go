package types

import "github.com/mauleyzaola/maupod/src/pkg/pb"

type Medias []*pb.Media

func (me Medias) InsertAt(m *pb.Media, index int) Medias {
	var result Medias
	if index > len(me) {
		return me
	}
	first := me[0:index]
	last := me[index:]
	result = append(first, m)
	result = append(result, last...)
	return result
}
