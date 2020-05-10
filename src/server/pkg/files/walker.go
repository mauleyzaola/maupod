package files

import (
	"io/ioutil"
	"path/filepath"
)

type WalkerFunc func(name string, isDir bool) (stop bool)

func WalkFiles(root string, fn WalkerFunc) error {
	files, err := ioutil.ReadDir(root)
	if err != nil {
		return err
	}

	for _, file := range files {
		isDir := file.IsDir()
		name := filepath.Join(root, file.Name())
		if fn != nil {
			if fn(name, isDir) {
				return nil
			}
		}
		if isDir {
			if err = WalkFiles(name, fn); err != nil {
				return err
			}
		}
	}

	return nil
}
