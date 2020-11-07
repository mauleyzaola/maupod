package rules

import (
	"errors"
	"path/filepath"
	"strings"
	"time"

	"github.com/mauleyzaola/maupod/src/protos"
	"github.com/spf13/viper"
)

func ConfigurationValidate(c *protos.Configuration) error {
	for _, v := range c.MediaStores {
		if err := FileStoreValidate(v); err != nil {
			return err
		}
	}
	return nil
}

func ConfigurationParse() (*protos.Configuration, error) {
	var c protos.Configuration
	if err := viper.Unmarshal(&c); err != nil {
		return nil, err
	}
	// check artwork size is not stupid
	if c.ArtworkBigSize < c.ArtworkSmallSize {
		return nil, errors.New("ArtworkBigSize cannot be smaller than ArtworkSmallSize")
	}
	return &c, nil
}

func FileIsValidExtension(c *protos.Configuration, filename string) bool {
	ext := strings.ToLower(filepath.Ext(filename))
	for _, v := range c.Extensions {
		if strings.ToLower(v) == ext {
			return true
		}
	}
	return false
}

func Timeout(c *protos.Configuration) time.Duration {
	return time.Second * time.Duration(c.Delay)
}
