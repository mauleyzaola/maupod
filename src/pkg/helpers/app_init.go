package helpers

import (
	"log"

	"github.com/spf13/viper"
)

func AppInit() {
	viper.AddConfigPath(".")
	viper.SetConfigType("yaml")
	viper.SetConfigName(".maupod")

	_ = viper.ReadInConfig()
	viper.AutomaticEnv()

	log.SetFlags(log.Lshortfile | log.Ltime | log.Ldate)
}
