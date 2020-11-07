package rules

import (
	"errors"
	"fmt"

	"github.com/mauleyzaola/maupod/src/protos"
)

func FileStoreValidate(store *protos.FileStore) error {
	if store == nil {
		return errors.New("fs is nil")
	}
	switch store.Type {
	case protos.FileStore_FILE_SYSTEM:
	case protos.FileStore_S3:
	default:
		return fmt.Errorf("unsupported file store: %s", store.Type)
	}
	return nil
}
