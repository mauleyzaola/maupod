package rule

import (
	"os"
	"path/filepath"
	"time"

	"github.com/mauleyzaola/maupod/src/server/pkg/helpers"
	"github.com/mauleyzaola/maupod/src/server/pkg/media"
	"github.com/mauleyzaola/maupod/src/server/pkg/pb"
)

func NewMediaFile(info *media.MediaInfo, filename string, scanDate time.Time, fileInfo os.FileInfo) *pb.Media {
	media := info.ToProto()
	media.Id = helpers.NewUUID()
	//media.Sha = "" // needs to be defined in another process
	media.LastScan = helpers.TimeToTs(&scanDate)
	media.ModifiedDate = helpers.TimeToTs2(fileInfo.ModTime())
	media.Location = filename
	media.FileExtension = filepath.Ext(fileInfo.Name())

	return media
}
