package helpers

import (
	"os"
	"path/filepath"
)

func AppName() string {
	return filepath.Base(os.Args[0])
}
