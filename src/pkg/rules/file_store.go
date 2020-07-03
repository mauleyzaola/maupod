package rules

import (
	"errors"
	"fmt"

	"github.com/mauleyzaola/maupod/src/pkg/pb"
)

func FileStoreValidate(store *pb.FileStore) error {
	if store == nil {
		return errors.New("fs is nil")
	}
	switch store.Type {
	case pb.FileStore_FILE_SYSTEM:
	case pb.FileStore_S3:
	default:
		return fmt.Errorf("unsupported file store: %s", store.Type)
	}
	return nil
}
