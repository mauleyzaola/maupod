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
	const songNowhereFast = "/media/mau/music-library/music/Compilations/Streets of Fire/01 Nowhere Fast.m4a"

	mpvProcess, err := pkg.NewMpvProcessor(songWarpigs)
	if err != nil {
		return err
	}

	// sample workflow pause/play
	ipc, err := pkg.NewIPC(mpvProcess)
	if err != nil {
		return err
	}
	if err = ipc.Load(songWarpigs); err != nil {
		return err
	}
	if err = ipc.Play(); err != nil {
		return err
	}
	time.Sleep(time.Second * 3)
	if err = ipc.Load(songNowhereFast); err != nil {
		return err
	}
	if err = ipc.Play(); err != nil {
		return err
	}
	time.Sleep(time.Second * 3)
	if err = ipc.Seek(50); err != nil {
		return err
	}
	if err = ipc.Volume(120); err != nil {
		return err
	}
	time.Sleep(time.Second * 3)
	if err = ipc.Volume(70); err != nil {
		return err
	}
	if err = ipc.SeekExact(200); err != nil {
		return err
	}
	time.Sleep(time.Second * 3)
	return ipc.Terminate()

}
