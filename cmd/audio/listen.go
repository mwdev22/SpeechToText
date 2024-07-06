package audio

import (
	"fmt"

	"github.com/gordonklaus/portaudio"
)

const (
	Rate            = 44100
	Channels        = 1
	secondsToRecord = 10
)

func Listen() {
	err := portaudio.Initialize()
	if err != nil {
		fmt.Errorf("error initializing portaudio, %s", err)
	}
}

func GenerateFile() {}

func SaveToFile() {

}
