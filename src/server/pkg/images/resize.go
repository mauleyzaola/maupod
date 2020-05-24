package images

import (
	"errors"
	"io"

	"github.com/disintegration/imaging"
)

func Size(r io.Reader) (x, y int, err error) {
	if r == nil {
		err = errors.New("missing parameter: r")
		return
	}
	img, err := imaging.Decode(r)
	if err != nil {
		return
	}
	bounds := img.Bounds()
	return bounds.Dx(), bounds.Dy(), nil
}

func ImageResize(r io.Reader, filename string, width, height int) error {
	if r == nil {
		return errors.New("missing parameter: r")
	}
	if filename == "" {
		return errors.New("missing parameter: filename")
	}
	img, err := imaging.Decode(r)
	if err != nil {
		return err
	}
	newImg := imaging.Resize(img, width, height, imaging.Lanczos)
	return imaging.Save(newImg, filename)
}
