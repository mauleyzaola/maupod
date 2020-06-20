package main

import (
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
	// mpv --no-video . --input-unix-socket=/tmp/mpv_socket
	conn := mpvipc.NewConnection("/tmp/mpv_socket")
	err := conn.Open()
	if err != nil {
		return err
	}
	defer conn.Close()

	//events, stopListening := conn.NewEventListener()

	//if err = conn.Set("playlist_next", ""); err != nil {
	//	return err
	//}

	var props = []string{}

	for _, prop := range props {
		value, err := conn.Get(prop)
		if err != nil {
			return err
		}
		log.Printf("prop: %s value: %v", prop, value)
	}

	const filePath = "/media/mau/music-library/music/10,000 Maniacs/10,000 Maniacs - Original Album Series/CD3/01 - Eat For Two.flac"

	//err = conn.Set("audio-device", "coreaudio/AppleUSBAudioEngine:Logitech USB Headset:Logitech USB Headset:14600000:2")
	//err = conn.Set("audio-device", "coreaudio/AppleGFXHDAEngineOutputDP:0:{6D9E-7721-0002D07E}")
	//err = conn.Set("external-file", "/media/mau/music-library/music/10,000 Maniacs/10,000 Maniacs - Original Album Series/CD5/02 - Eat For Two.flac")
	_, err = conn.Call("start", "00:10")
	if err != nil {
		return err
	}

	//_, err = conn.Call("observe_property", 42, "volume")
	//if err != nil {
	//	fmt.Print(err)
	//}

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
	// consider now only the default sound interface, but should support more in the future

	//go func() {
	//	conn.WaitUntilClosed()
	//	stopListening <- struct{}{}
	//}()

	//for event := range events {
	//	if event.ID == 42 {
	//		log.Printf("volume now is %f", event.Data.(float64))
	//	} else {
	//		log.Printf("received event: %s id: %v", event.Name, event.Text)
	//	}
	//}

	//log.Printf("mpv closed socket")

	return nil
}
