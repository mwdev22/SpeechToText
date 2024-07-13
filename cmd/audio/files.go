package audio

import (
	"encoding/binary"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/youpy/go-wav"
)

var FILES_DIR string = getFileDir()

func getFileDir() string {
	cmd := exec.Command("go", "env", "GOMOD")
	output, err := cmd.Output()
	if err != nil {
		panic(err)
	}

	goModPath := strings.TrimSpace(string(output))

	if goModPath == "" {
		panic("go.mod file not found")
	}

	moduleDir := filepath.Dir(goModPath)

	path := filepath.Join(moduleDir, "audio_files")

	if _, err := os.Stat(path); os.IsNotExist(err) {
		err := os.Mkdir(path, os.ModePerm)
		if err != nil {
			panic(err)
		}
	}

	return path
}

func SaveToWavFile(data []int32) string {

	file, err := os.Create("output.wav")
	if err != nil {
		log.Fatalf("Failed to create WAV file: %v", err)
	}
	defer file.Close()

	w := wav.NewWriter(file, uint32(len(data)), uint16(Channels), uint32(Rate), 16)

	// convert int32 audio data to int16 byte slice with proper scaling
	buf := make([]byte, len(data)*2) // 2 bytes per sample for int16
	for i := 0; i < len(data); i++ {
		sample := data[i]
		// scale down from int32 to int16
		scaledSample := int16(sample >> 16) // reduces the sample value to fit within the int16 range
		binary.LittleEndian.PutUint16(buf[i*2:], uint16(scaledSample))
	}

	// write audio data to WAV file
	if _, err := w.Write(buf); err != nil {
		log.Fatalf("Failed to write data to WAV file: %v", err)
	}

	log.Println("WAV file saved.")
	return file.Name()
}

func ReadWavFile(filename string) []byte {

	if !filepath.IsAbs(filename) {
		filename = filepath.Join(FILES_DIR, filename)
	}

	file, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	return file
}
