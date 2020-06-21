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
	const songWarpigs = "/media/mau/music-library/music/Black Sabbath/1970 Paranoid (Black Box Remaster)/01 War Pigs , Luke's Wall.flac"
	const songGrave = "/media/mau/music-library/music/Black Sabbath/1971 Master Of Reality (Black Box Remaster)/04 Children Of The Grave.flac"

	mpvProcess, err := pkg.NewMPVProcess(songWarpigs)
	if err != nil {
		return err
	}

	// sample workflow pause/play
	ipc, err := pkg.NewIPC(mpvProcess)
	if err != nil {
		return err
	}
	if err = ipc.Load(songGrave); err != nil {
		return err
	}
	if err = ipc.Play(); err != nil {
		return err
	}
	time.Sleep(time.Second * 3)
	if err = ipc.Load(songWarpigs); err != nil {
		return err
	}
	if err = ipc.Play(); err != nil {
		return err
	}
	time.Sleep(time.Second * 3)
	return ipc.Terminate()

}
