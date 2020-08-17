package images

import (
	"bufio"
	"io"
	"log"
	"strconv"

	"github.com/anthonynsimon/bild/imgio"
	"github.com/anthonynsimon/bild/transform"
	"github.com/mauleyzaola/maupod/src/pkg/information"
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
	img,err:=imgio.Open(source)
	if err!=nil{
		return err
	}
	resized:=transform.Resize(img,width,height,transform.Linear)
	if err=imgio.Save(target,resized,imgio.PNGEncoder());err!=nil{
		log.Print(err)
		return err
	}
	return nil
}
