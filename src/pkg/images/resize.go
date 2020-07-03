package images

import (
	"bufio"
	"io"
	"strconv"

	"github.com/mauleyzaola/maupod/src/pkg/information"
	"gopkg.in/gographics/imagick.v2/imagick"
)

func Size(r io.Reader) (x, y int, err error) {
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		line := information.InfoString(scanner.Text())
		key, value := line.Split()
		switch key {
		case "Width":
			if val, err := strconv.Atoi(value); err == nil {
				x = val
			}
		case "Height":
			if val, err := strconv.Atoi(value); err == nil {
				y = val
			}
		default:
			continue
		}
		if x != 0 && y != 0 {
			break
		}
	}
	return
}

func ImageResize(source, target string, width, height int) error {
	imagick.Initialize()
	defer imagick.Terminate()
	mw := imagick.NewMagickWand()
	if err := mw.ReadImage(source); err != nil {
		return err
	}
	if err := mw.ResizeImage(uint(width), uint(height), imagick.FILTER_LANCZOS, 1); err != nil {
		return err
	}
	if err := mw.WriteImage(target); err != nil {
		return err
	}
	return nil
}
