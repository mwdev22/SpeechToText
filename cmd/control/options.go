package control

import (
	"fmt"
	"speech_to_text/cmd/audio"
	"speech_to_text/cmd/gcp"
)

func ListenAndTranscript() {
	filename := audio.Listen()
	gcp.TranscriptFile(filename)
}

func Transcript() {
	var input string
	fmt.Print("Enter absolute filepath or filename from /audio directory...\n")
	fmt.Scanln(&input)
	fmt.Printf("Opening %s file...\n", input)
	gcp.TranscriptFile(input)
}
