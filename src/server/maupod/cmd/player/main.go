package main

import (
	"bytes"
	"encoding/json"
	"log"
	"os"
	"os/exec"

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
	// https://github.com/mpv-player/mpv/blob/master/DOCS/man/ipc.rst
	// https://mpv.io/manual/master/#list-of-events
	// https://github.com/DexterLB/mpvipc/blob/master/mpvipc.go
	const command = `{ "command": ["get_property", "playback-time"] }`

	cmd := exec.Command("./socket.sh")
	cmd.Stdin = bytes.NewBufferString(command)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return err
	}

	var input = struct {
		Data  float64
		Error string
	}{}

	if err = json.Unmarshal(output, &input); err != nil {
		return err
	}

	log.Println("data: ", input.Data)
	log.Println("error: ", input.Error)

	return nil
}
