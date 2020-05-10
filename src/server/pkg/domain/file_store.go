package domain

import (
	"errors"
	"fmt"
)

type FileStore struct {
	Name     string `json:"name"`
	Type     string `json:"type"`
	Location string `json:"location"`
}

const (
	FileTypeFileSystem = "file-system"
	FileTypeS3         = "s3"
)

func (fs *FileStore) Validate() error {
	if fs == nil {
		return errors.New("fs is nil")
	}
	switch fs.Type {
	case FileTypeFileSystem:
	case FileTypeS3:
	default:
		return fmt.Errorf("unsupported file store: %s", fs.Type)
	}
	return nil
}
