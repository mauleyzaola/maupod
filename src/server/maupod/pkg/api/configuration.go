package api

import (
	"github.com/spf13/viper"
)

func ParseConfiguration() (*Configuration, error) {
	var c Configuration
	if err := viper.Unmarshal(&c); err != nil {
		return nil, err
	}
	return &c, nil
}
