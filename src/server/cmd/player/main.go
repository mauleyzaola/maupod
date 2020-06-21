package main

import (
	"log"
	"os"
	"time"

	"github.com/mauleyzaola/maupod/src/server/cmd/player/pkg"
	"github.com/spf13/viper"
)

func main() {
	if err := run(); err != nil {
		log.Println("error: ", err)
		os.Exit(1)
	}
	os.Exit(0)
}

func init() {
	//viper.AddConfigPath(".")
	//viper.SetConfigType("yaml")
	//viper.SetConfigName(".maupod")

	//_ = viper.ReadInConfig()
	viper.AutomaticEnv()
}

func run() error {
	const songWarpigs = "/Users/mau/Downloads/music/Black Sabbath/1970 Paranoid (Black Box Remaster)/01 War Pigs , Luke's Wall.flac"
	const songGrave = "/Users/mau/Downloads/music/Black Sabbath/1971 Master Of Reality (Black Box Remaster)/04 Children Of The Grave.flac"
	wr, err := pkg.NewMPV(songWarpigs)
	if err != nil {
		return err
	}
	wr.Play(songWarpigs)
	wr.PauseToggle()
	time.Sleep(time.Second * 15)
	return wr.Terminate()
}
