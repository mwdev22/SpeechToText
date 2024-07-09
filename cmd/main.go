package main

import (
	"fmt"
	"speech_to_text/cmd/audio"
	"speech_to_text/cmd/config"
)

func main() {
	config.GetEnv()
	fmt.Printf("App is starting")
	audio.Listen()
}
