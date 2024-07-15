package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var SpeechAPIKey string
var DefDeviceName string

func GetEnv() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
	SpeechAPIKey = os.Getenv("SP_TO_TXT_KEY")
	DefDeviceName = os.Getenv("AUDIO_DEVICE_NAME")
}
