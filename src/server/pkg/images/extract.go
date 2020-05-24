package images

import (
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/mauleyzaola/maupod/src/server/pkg/helpers"
)

func ExtractImageFromMedia(w io.Writer, filename string) error {
	const ffmpeg = "ffmpeg"
	if !helpers.ProgramExists(ffmpeg) {
		return fmt.Errorf("could not find program: %s in path", ffmpeg)
	}
	outputFile := filepath.Join(os.TempDir(), helpers.NewUUID()+".png")
	var p = []string{
		"-i",
		filename,
		outputFile,
	}
	cmd := exec.Command(ffmpeg, p...)
	data, err := cmd.CombinedOutput()
	if err != nil {
		return errors.New(string(data) + " : " + err.Error())
	}

	file, err := os.Open(outputFile)
	if err != nil {
		return err
	}
	defer func() {
		if err = file.Close(); err != nil {
			log.Println(err)
		}
		if err = os.Remove(outputFile); err != nil {
			log.Println(err)
		}
	}()
	info, err := file.Stat()
	if err != nil {
		return err
	}
	if info.Size() == 0 {
		return errors.New("file is empty")
	}
	if _, err = io.Copy(w, file); err != nil {
		return err
	}
	return nil
}
