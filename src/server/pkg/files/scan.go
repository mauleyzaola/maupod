package files

import (
	"errors"
	"path/filepath"
	"strings"

	"github.com/mauleyzaola/maupod/src/server/pkg/helpers"
)

func ScanFiles(extensions []string, directories ...string) ([]string, error) {
	var lowerExt []string
	for _, v := range extensions {
		lowerExt = append(lowerExt, strings.ToLower(v))
	}
	validExt := helpers.StringSlice(lowerExt).ToMap()
	if len(validExt) == 0 {
		return nil, errors.New("missing parameter: extensions")
	}
	var result []string
	keys := make(map[string]struct{})
	var fn WalkerFunc = func(name string, isDir bool) {
		// ignore directories
		if isDir {
			return
		}
		// consider only the allowed extensions
		ext := strings.ToLower(filepath.Ext(name))
		if _, ok := validExt[ext]; !ok {
			return
		}
		// ignore duplicate files
		if _, ok := keys[name]; ok {
			return
		}
		keys[name] = struct{}{}
		result = append(result, name)
	}
	for _, dir := range directories {
		if err := WalkFiles(dir, fn); err != nil {
			return nil, err
		}
	}
	return result, nil
}
