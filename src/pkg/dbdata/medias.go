package dbdata

import "github.com/mauleyzaola/maupod/src/server/pkg/pb"

type Medias []*pb.Media

// ToMap returns a map which key is the location of the audio file
func (m Medias) ToMap() map[string]*pb.Media {
	res := make(map[string]*pb.Media)
	for _, v := range m {
		res[v.Location] = v
	}
	return res
}
