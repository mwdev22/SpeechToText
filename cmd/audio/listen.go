package audio

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/gordonklaus/portaudio"
)

const (
	Rate            = 44100
	Channels        = 1
	secondsToRecord = 10
)

var audioBuffer []int32

func Listen() string {
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

	// list the audio devices for user to pick
	var inputDevice *portaudio.DeviceInfo
	for i, device := range devices {
		fmt.Printf("%v. %s\n", i, device.Name)
	}

	fmt.Print("Enter number of input audio device...\n")
	var devNum string
	fmt.Scanln(&devNum)

	dev, err := strconv.Atoi(devNum)
	if err != nil {
		log.Fatalf("invalid number of device: %d", dev)
	}

	inputDevice = devices[dev]

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

	fmt.Print("Enter number of seconds to determine recording time...\n")
	var sec string
	fmt.Scanln(&sec)
	seconds, err := strconv.Atoi(sec)
	if err != nil || seconds <= 0 {
		log.Fatalf("invalid input for seconds: %s", sec)
	}

	err = stream.Start()
	if err != nil {
		log.Fatalf("failed to start audio stream: %s", err)
	}

	defer stream.Close()

	fmt.Printf("Listening for %d seconds...\n", seconds)
	time.Sleep(time.Duration(seconds) * time.Second)

	err = stream.Stop()
	if err != nil {
		log.Fatalf("Failed to stop audio stream: %v", err)
	}

	fmt.Print("Audio capturing finished.\n")

	return SaveToWavFile(audioBuffer)
}

func processAudio(in []int32) {
	audioBuffer = append(audioBuffer, in...)
}
