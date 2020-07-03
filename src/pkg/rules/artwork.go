package rules

import "github.com/mauleyzaola/maupod/src/pkg/pb"

func ArtworkFileName(media *pb.Media) string {
	if media.AlbumIdentifier == "" {
		return ""
	}
	return media.AlbumIdentifier + ".png"
}
