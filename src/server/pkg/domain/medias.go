package domain

type Medias []*Media

// ToMap returns a map which key is the location of the audio file
func (m Medias) ToMap() map[string]*Media {
	res := make(map[string]*Media)
	for _, v := range m {
		res[v.Location] = v
	}
	return res
}
