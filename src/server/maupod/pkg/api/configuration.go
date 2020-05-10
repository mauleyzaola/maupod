package api

import (
	"github.com/mauleyzaola/maupod/src/server/pkg/domain"
	"github.com/spf13/viper"
)

func ParseConfiguration() (*domain.Configuration, error) {
	var c domain.Configuration
	if err := viper.Unmarshal(&c); err != nil {
		return nil, err
	}
	return &c, nil
}
