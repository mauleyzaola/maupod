package dbdata

import (
	"github.com/mauleyzaola/maupod/src/protos"
)

type Medias []*protos.Media

// ToMap returns a map which key is the location of the audio file
func (m Medias) ToMap() map[string]*protos.Media {
	res := make(map[string]*protos.Media)
	for _, v := range m {
		res[v.Location] = v
	}
	return res
}
