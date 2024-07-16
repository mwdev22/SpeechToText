package audio

import (
	"encoding/binary"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/youpy/go-wav"
)

var AudioDir string = getFileDir("audio_files")
var TranscriptionDir string = getFileDir("transcriptions")

func getFileDir(dirname string) string {
	cmd := exec.Command("go", "env", "GOMOD")
	output, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
	}

	goModPath := strings.TrimSpace(string(output))

	if goModPath == "" {
		log.Fatalf("go.mod file not found")
	}

	moduleDir := filepath.Dir(goModPath)

	path := filepath.Join(moduleDir, dirname)

	if _, err := os.Stat(path); os.IsNotExist(err) {
		err := os.Mkdir(path, os.ModePerm)
		if err != nil {
			log.Fatal(err)
		}
	}

	return path
}

func SaveToWavFile(data []int32) string {
	// saving file in audio directory
	var fname string
	fmt.Print("Enter the name of the recording file...\n")
	fmt.Scanln(&fname)
	fname = fname + ".wav"

	filePath := filepath.Join(AudioDir, fname)

	file, err := os.Create(filePath)
	if err != nil {
		log.Fatalf("Failed to create WAV file: %v", err)
	}
	defer file.Close()

	// uint types converting, cause NewWriter args are not generic
	w := wav.NewWriter(file, uint32(len(data)), uint16(Channels), uint32(Rate), 16)

	// convert int32 audio data to int16 byte slice with proper scaling
	buf := make([]byte, len(data)*2) // 2 bytes per sample for int16
	for i := 0; i < len(data); i++ {
		sample := data[i]
		// scale down from int32 to int16, with right shift
		scaledSample := int16(sample >> 16) // reduces the sample value to fit within the int16 range
		binary.LittleEndian.PutUint16(buf[i*2:], uint16(scaledSample))
	}

	// write audio data to WAV file
	if _, err := w.Write(buf); err != nil {
		log.Fatalf("Failed to write data to WAV file: %v", err)
	}

	log.Println("WAV file saved.")
	if _, err := w.Write(buf); err != nil {
		log.Fatalf("Failed to write data to WAV file: %v", err)
	}

	// close the file to ensure all data is written and flushed to disk
	err = file.Close()
	if err != nil {
		log.Fatalf("Failed to close WAV file: %v", err)
	}

	// let file info to determine the file size
	fileInfo, err := os.Stat(filePath)
	if err != nil {
		log.Fatalf("Failed to get file info: %v", err)
	}

	// log the file size in bytes
	fileSize := fileInfo.Size()
	log.Printf("WAV file %s saved. Size: %d bytes.\n", fileInfo.Name(), fileSize)

	return fileInfo.Name()
}

func ReadWavFile(filename string) *[]byte {

	if !filepath.IsAbs(filename) {
		filename = filepath.Join(AudioDir, filename)
	}

	file, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	return &file
}

func SaveTranscription(filename, transcription string) (string, error) {
	fullPath := filepath.Join(TranscriptionDir, filename+".txt")

	file, err := os.Create(fullPath)
	if err != nil {
		log.Fatal(err)
		return "", err
	}
	l, err := file.WriteString(transcription)
	if err != nil {
		log.Fatal(err)
		return "", err
	}
	fmt.Printf("%v bytes written successfully.", l)

	return fullPath, nil
}
