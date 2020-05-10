package cmd

import (
	"context"
	"database/sql"
	"log"
	"path/filepath"

	"github.com/mauleyzaola/maupod/src/server/pkg/data/psql"

	"github.com/mauleyzaola/maupod/src/server/pkg/data/orm"

	"github.com/mauleyzaola/maupod/src/server/pkg/media"

	"github.com/mauleyzaola/maupod/src/server/pkg/domain"
	"github.com/mauleyzaola/maupod/src/server/pkg/files"
	"github.com/mauleyzaola/maupod/src/server/pkg/helpers"
	"github.com/spf13/cobra"
)

// scannerCmd represents the restapi command
var scannerCmd = &cobra.Command{
	Use:   "audioscan",
	Short: "Scans for new audio files",
	Long:  "",
	RunE: func(cmd *cobra.Command, args []string) error {
		config, err := domain.ParseConfiguration()
		if err != nil {
			return err
		}

		if err = config.Validate(); err != nil {
			return err
		}

		var db *sql.DB
		var conn *sql.Tx
		if db, err = helpers.DbBootstrap(config); err != nil {
			return err
		}
		defer func() {
			if conn != nil {
				if err != nil {
					err = conn.Rollback()
				} else {
					err = conn.Commit()
				}
				if err != nil {
					log.Println(err)
				}
			}
			if err = db.Close(); err != nil {
				log.Println(err)
			}
		}()
		if conn, err = db.Begin(); err != nil {
			return err
		}

		// buffer all the existing hashes in db
		store := &psql.MediaStore{}
		log.Println("reading previous hashes from db")
		ctx := context.Background()
		hashes := make(map[string]struct{})
		rows, err := orm.Media().All(ctx, db)
		if err != nil {
			return err
		}
		for _, v := range rows {
			hashes[v.ID] = struct{}{}
		}
		log.Println("completed reading hashes")

		root := "/media/mau/music-library/music/"
		walker := func(filename string, isDir bool) bool {
			// debug
			if len(hashes) > 10 {
				return true
			}
			// debug

			var info *media.MediaInfo
			var ID string
			if isDir {
				return false
			}
			if !config.FileIsValidExtension(filename) {
				return false
			}

			if err = media.MediaInfoWithId(filename, func(mi *media.MediaInfo, id string) {
				info = mi
				ID = id
			}); err != nil {
				log.Println(err)
				return false
			}
			if _, ok := hashes[ID]; ok {
				return false
			}
			m := info.ToDomain()
			m.FileExtension = filepath.Ext(filename)
			m.ID = ID
			m.Location = m.ID + m.FileExtension
			if err = store.Insert(ctx, conn, m); err != nil {
				log.Println(err)
				return true
			}
			log.Println(filename)
			hashes[ID] = struct{}{}
			return false
		}
		if err = files.WalkFiles(root, walker); err != nil {
			return err
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(scannerCmd)
}
