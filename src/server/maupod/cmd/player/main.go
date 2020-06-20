package main

import (
	"fmt"
	"log"
	"os"

	"github.com/DexterLB/mpvipc"
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
	conn := mpvipc.NewConnection("/tmp/mpvsocket")
	err := conn.Open()
	if err != nil {
		return err
	}
	defer conn.Close()

	events, stopListening := conn.NewEventListener()

	path, err := conn.Get("path")
	if err != nil {
		return err
	}
	log.Printf("current file playing: %s", path)

	//err = conn.Set("pause", true)
	//if err != nil {
	//	return err
	//}
	//log.Printf("paused!")

	_, err = conn.Call("observe_property", 42, "volume")
	if err != nil {
		fmt.Print(err)
	}

	//ticker := time.NewTicker(time.Millisecond * 1000)
	//go func() {
	//	for {
	//		select {
	//		case t := <-ticker.C:
	//			v, err := conn.Get("playback-time")
	//			if err != nil {
	//				panic(err)
	//			}
	//			fmt.Println("playback-time: ", v, t)
	//		}
	//	}
	//}()

	// TODO: implement state of paused/playing
	// consider now only the default sound interface, but shold support more in the future

	go func() {
		conn.WaitUntilClosed()
		stopListening <- struct{}{}
	}()

	for event := range events {
		if event.ID == 42 {
			log.Printf("volume now is %f", event.Data.(float64))
		} else {
			log.Printf("received event: %s id: %v", event.Name, event.Text)
		}
	}

	log.Printf("mpv closed socket")
	return nil
}
