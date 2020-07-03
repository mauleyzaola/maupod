package rules

import "github.com/mauleyzaola/maupod/src/pkg/pb"

func ArtworkFileName(media *pb.Media) string {
	if media.ShaImage == "" {
		return ""
	}
	return media.ShaImage + ".png"
}
