package main

import (
	"fmt"
	"speech_to_text/cmd/config"
	"speech_to_text/cmd/control"
)

func main() {
	config.GetEnv()

chooseOption:
	fmt.Println("1 - Read an audio file and transcribe it.")
	fmt.Println("2 - Record your voice from an audio device and then transcribe it.")
	fmt.Println("0 - Exit.")

	var input string

	fmt.Print("Enter your choice: ")
	fmt.Scanln(&input)

	switch input {
	case "1":
		control.Transcript()
	case "2":
		control.ListenAndTranscript()
	case "0":
		fmt.Println("Exiting program.")
		return
	default:
		fmt.Println("Invalid choice. Please try again.")
		goto chooseOption
	}
}
