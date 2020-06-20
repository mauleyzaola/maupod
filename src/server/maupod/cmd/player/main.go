package main

import (
	"log"
	"os"
	"time"

	"github.com/mauleyzaola/maupod/src/server/maupod/cmd/player/pkg"

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
	const songPath = "/media/mau/music-library/music/Adriana Calcanhotto/Maré/05 Mulher Sem Razão.m4a"
	const songPath2 = "/media/mau/music-library/music/Gino Vannelli/Gino Vannelli Live/06 Hurts to Be In Love (Live).m4a"

	process, err := pkg.MPVStart(songPath)
	if err != nil {
		return err
	}
	// need to have a startup time
	time.Sleep(time.Second)

	log.Println("pid: ", process.Pid)
	//time.Sleep(time.Second * 5)
	//log.Println("trying to kill process")
	//if err = process.Kill(); err != nil {
	//	return err
	//}
	//return nil

	// mpv --no-video . --input-unix-socket=/tmp/mpv_socket
	conn := mpvipc.NewConnection("/tmp/mpv_socket")
	err = conn.Open()
	if err != nil {
		return err
	}
	defer func() { _ = conn.Close() }()

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

	//err = conn.Set("audio-device", "coreaudio/AppleUSBAudioEngine:Logitech USB Headset:Logitech USB Headset:14600000:2")
	//err = conn.Set("audio-device", "coreaudio/AppleGFXHDAEngineOutputDP:0:{6D9E-7721-0002D07E}")
	//err = conn.Set("external-file", "/media/mau/music-library/music/10,000 Maniacs/10,000 Maniacs - Original Album Series/CD5/02 - Eat For Two.flac")
	//_, err = conn.Call("loadfile", songPath)
	//if err != nil {
	//	return err
	//}

	play := func(p string) {
		log.Println("loading file: ", p)
		conn.Call("loadfile", p)
		conn.Set("pause", false)
		time.Sleep(time.Second * 5)
	}

	play(songPath)
	play(songPath2)
	play(songPath2)
	play(songPath)

	log.Println("killing mpv process")
	process.Kill()
	return nil

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
