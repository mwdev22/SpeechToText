package main

import (
	"fmt"
	"speech_to_text/cmd/audio"
)

func main() {
	fmt.Printf("App is starting")
	audio.Listen()
}
