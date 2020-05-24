package cmd

import (
	"database/sql"
	"errors"
	"log"
	"os"

	"github.com/mauleyzaola/maupod/src/server/pkg/data"
	"github.com/mauleyzaola/maupod/src/server/pkg/rule"

	"github.com/spf13/cobra"
)

var albumimageCmd = &cobra.Command{
	Use:   "albumimage",
	Short: "extract image from audio file",
	Long:  ``,
	RunE: func(cmd *cobra.Command, args []string) error {
		config, err := rule.ConfigurationParse()
		if err != nil {
			return err
		}

		if err = rule.ConfigurationValidate(config); err != nil {
			return err
		}

		// read first image store in the configuration and all the stores where to look up for audio files
		var imageStore = rule.ConfigurationFirstImageStore(config)
		if imageStore == nil {
			return errors.New("could not find any image store in configuration, exiting now")
		}

		var fileSystemStores = rule.ConfigurationFileSystemStores(config)
		var roots []string
		for _, v := range fileSystemStores {
			roots = append(roots, v.Location)
		}
		for _, root := range roots {
			if _, err = os.Stat(root); err != nil {
				return err
			}
		}

		//scanDate := time.Now()
		var db *sql.DB
		if db, err = data.DbBootstrap(config); err != nil {
			return err
		}
		defer func() {
			if err = db.Close(); err != nil {
				log.Println(err)
			}
		}()

		// create store directory if it doesn't exist
		if err = os.MkdirAll(imageStore.Location, os.ModePerm); err != nil {
			return err
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(albumimageCmd)
}
