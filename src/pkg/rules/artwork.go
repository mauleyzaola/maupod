package rules

import "github.com/mauleyzaola/maupod/src/protos"

func ArtworkFileName(media *protos.Media) string {
	if media.AlbumIdentifier == "" {
		return ""
	}
	return media.AlbumIdentifier + ".png"
}
