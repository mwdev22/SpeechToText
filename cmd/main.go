package main

import (
	"fmt"
	"speech_to_text/cmd/audio"
	"speech_to_text/cmd/config"
)

func main() {
	config.GetEnv()

	var input string
	var optionString string = "Choose and option:\n1 - Read an audio file and transcript it.\n2 - Record your voice from audio device and then transciprt it."
	fmt.Print(optionString)
	fmt.Scanln(&input)

	switch input {
	case "2":
		audio.Listen()
	}
	audio.Listen()
}
