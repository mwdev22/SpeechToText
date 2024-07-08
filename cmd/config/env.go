package config

import "os"

var SpeechAPIKey = os.Getenv("SP_TO_TXT_KEY")

var DeviceName = os.Getenv("AUDIO_DEVICE_NAME")
