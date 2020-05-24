/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
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

// albumimageCmd represents the albumimage command
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

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// albumimageCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// albumimageCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
