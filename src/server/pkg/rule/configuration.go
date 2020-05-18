package rule

import (
	"path/filepath"
	"strings"

	"github.com/mauleyzaola/maupod/src/server/pkg/pb"
	"github.com/spf13/viper"
)

func ConfigurationValidate(c *pb.Configuration) error {
	for _, v := range c.Stores {
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
