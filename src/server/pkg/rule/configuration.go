package rule

import (
	"path/filepath"
	"strings"

	"github.com/mauleyzaola/maupod/src/server/pkg/pb"
	"github.com/spf13/viper"
)

func ConfigurationValidate(c *pb.Configuration) error {
	for _, v := range c.MediaStores {
		if err := FileStoreValidate(v); err != nil {
			return err
		}
	}
	return nil
}

func ConfigurationParse() (*pb.Configuration, error) {
	var c pb.Configuration
	if err := viper.Unmarshal(&c); err != nil {
		return nil, err
	}
	return &c, nil
}

func FileIsValidExtension(c *pb.Configuration, filename string) bool {
	ext := strings.ToLower(filepath.Ext(filename))
	for _, v := range c.Extensions {
		if strings.ToLower(v) == ext {
			return true
		}
	}
	return false
}

func ConfigurationFirstImageStore(c *pb.Configuration) *pb.FileStore {
	// TODO: implement
	return nil
}

func ConfigurationFileSystemStores(c *pb.Configuration) []*pb.FileStore {
	var roots []*pb.FileStore
	for _, v := range c.MediaStores {
		switch v.Type {
		case pb.FileStore_FILE_SYSTEM:
			roots = append(roots, v)
		default:
			continue
		}
	}
	if len(roots) == 0 {
		// if no store is available in yaml file, lookup in the environment
		if val := viper.GetString("MEDIA_STORE"); val != "" {
			roots = append(roots, &pb.FileStore{
				Name:     "store",
				Type:     pb.FileStore_FILE_SYSTEM,
				Location: val,
			})
		}
	}
	return roots
}
