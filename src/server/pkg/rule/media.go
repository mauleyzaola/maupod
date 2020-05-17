package rule

import (
	"os"
	"path/filepath"
	"time"

	"github.com/mauleyzaola/maupod/src/server/pkg/domain"
	"github.com/mauleyzaola/maupod/src/server/pkg/helpers"
	"github.com/mauleyzaola/maupod/src/server/pkg/media"
)

func NewMediaFile(info *media.MediaInfo, filename string, scanDate time.Time, fileInfo os.FileInfo) *domain.Media {
	media := info.ToDomain()
	media.ID = helpers.NewUUID()
	//media.Sha = "" // needs to be defined in another process
	media.LastScan = scanDate
	media.ModifiedDate = fileInfo.ModTime()
	media.Location = filename
	media.FileExtension = filepath.Ext(fileInfo.Name())

	return media
}
