package rule

import (
	"errors"
	"fmt"

	"github.com/mauleyzaola/maupod/src/server/pkg/pb"
)

func FileStoreValidate(store *pb.FileStore) error {
	if store == nil {
		return errors.New("fs is nil")
	}
	switch store.Type {
	case pb.FileStore_FILE_SYSTEM:
	case pb.FileStore_S3:
	case pb.FileStore_IMAGE:
	default:
		return fmt.Errorf("unsupported file store: %s", store.Type)
	}
	return nil
}
