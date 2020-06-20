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
	const songAdriana = "/media/mau/music-library/music/Adriana Calcanhotto/Maré/05 Mulher Sem Razão.m4a"
	const songGino = "/media/mau/music-library/music/Gino Vannelli/Gino Vannelli Live/06 Hurts to Be In Love (Live).m4a"

	process, err := pkg.MPVStart(songAdriana)
	if err != nil {
		return err
	}
	// need to have a startup time
	time.Sleep(time.Second)

	log.Println("pid: ", process.Pid)

	conn := mpvipc.NewConnection("/tmp/mpv_socket")
	err = conn.Open()
	if err != nil {
		return err
	}
	defer func() { _ = conn.Close() }()

	var props = []string{}

	for _, prop := range props {
		value, err := conn.Get(prop)
		if err != nil {
			return err
		}
		log.Printf("prop: %s value: %v", prop, value)
	}

	play := func(p string) {
		log.Println("loading file: ", p)
		conn.Call("loadfile", p)
		conn.Set("pause", false)
	}

	seek := func(offset int) {
		log.Println("offset to: ", offset)
		conn.Call("seek", offset, "exact")
	}

	sleep := func(secs int) {
		log.Println("sleeping secs: ", secs)
		time.Sleep(time.Second * time.Duration(secs))
	}

	play(songAdriana)
	sleep(5)
	seek(150)
	sleep(5)
	play(songGino)
	sleep(2)
	seek(200)
	sleep(5)
	seek(-75)
	sleep(3)

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
