package audio

import (
	"fmt"
	"log"
	"speech_to_text/cmd/config"
	"time"

	"github.com/gordonklaus/portaudio"
)

const (
	Rate            = 44100
	Channels        = 1
	secondsToRecord = 10
)

var audioBuffer []int32

func Listen() {
	err := portaudio.Initialize()
	if err != nil {
		log.Fatalf("failed to initialize: %s", err)
	}
	defer portaudio.Terminate()

	// list available devices
	devices, err := portaudio.Devices()
	if err != nil {
		log.Fatalf("failed to get devices: %v", err)
	}

	// find the specified in .env device
	var inputDevice *portaudio.DeviceInfo
	for _, device := range devices {
		if device.Name == config.DeviceName {
			inputDevice = device
		}
	}

	if inputDevice == nil {
		log.Fatalf("failed to find the specified audio device")
	}

	inputParameters := portaudio.StreamParameters{
		Input: portaudio.StreamDeviceParameters{
			Device:   inputDevice,
			Channels: Channels,
			Latency:  inputDevice.DefaultLowInputLatency,
		},
		SampleRate:      Rate,
		FramesPerBuffer: portaudio.FramesPerBufferUnspecified,
	}

	stream, err := portaudio.OpenStream(inputParameters, processAudio)
	if err != nil {
		log.Fatalf("failed to open audio stream: %v", err)
	}

	err = stream.Start()
	if err != nil {
		log.Fatalf("failed to start audio stream: %s", err)
	}

	defer stream.Close()

	fmt.Println("Listening for 10 seconds...")
	time.Sleep(10 * time.Second)

	err = stream.Stop()
	if err != nil {
		log.Fatalf("Failed to stop audio stream: %v", err)
	}

	fmt.Println("Audio capturing finished.")

	SaveToWavFile(audioBuffer)
}

func processAudio(in []int32) {
	audioBuffer = append(audioBuffer, in...)
}
