package domain

import (
	"path/filepath"
	"strings"
	"time"

	"github.com/spf13/viper"
)

type Configuration struct {
	Stores     []FileStore
	PgConn     string
	Retries    int
	Delay      time.Duration
	Port       string
	Extensions []string
}

func (c *Configuration) Validate() error {
	for _, v := range c.Stores {
		if err := v.Validate(); err != nil {
			return err
		}
	}
	return nil
}

func ParseConfiguration() (*Configuration, error) {
	var c Configuration
	if err := viper.Unmarshal(&c); err != nil {
		return nil, err
	}
	return &c, nil
}

func (c *Configuration) FileIsValidExtension(filename string) bool {
	ext := strings.ToLower(filepath.Ext(filename))
	for _, v := range c.Extensions {
		if strings.ToLower(v) == ext {
			return true
		}
	}
	return false
}
