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
	fmt.Print("Enter absolut filepath or filename from /audio directory.")
	fmt.Scanln(&input)
	fmt.Printf("Transcripting %s file...", input)
	gcp.TranscriptFile(input)
}
