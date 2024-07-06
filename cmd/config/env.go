package config

import "os"

var SpeechAPIKey = os.Getenv("SP_TO_TXT_KEY")
