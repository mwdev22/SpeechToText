package gcp

import (
	"context"
	"fmt"
	"log"
	"speech_to_text/cmd/audio"
	"strconv"
	"strings"

	speech "cloud.google.com/go/speech/apiv1"
	"cloud.google.com/go/speech/apiv1/speechpb"
)

var languages = map[string]string{
	"Polish":      "pl-PL",
	"English(GB)": "en-GB",
	"English(US)": "en-US",
}

func TranscriptFile(filename string) {

	fmt.Println("Available languages:")
	langs := make([]string, 0, len(languages))
	for lang := range languages {
		langs = append(langs, lang)
	}
	for i, lang := range langs {
		fmt.Printf("%d. %s\n", i+1, lang)
	}

	// getting language from user
	fmt.Print("Enter the number for the language: ")
	var selectedNumber string
	fmt.Scanln(&selectedNumber)

	// convert input to integer
	index, err := strconv.Atoi(selectedNumber)
	if err != nil || index < 1 || index > len(langs) {
		log.Fatalf("Invalid selection: %s", selectedNumber)
	}

	// Get the selected language code
	selectedLanguage := langs[index-1]
	languageCode := languages[selectedLanguage]

	audioBytes := audio.ReadWavFile(filename)

	ctx := context.Background()
	client, err := speech.NewClient(ctx)
	if err != nil {
		log.Fatalf("Failed to create speech client: %v", err)
	}
	defer client.Close()

	audioContent := &speechpb.RecognitionAudio{
		AudioSource: &speechpb.RecognitionAudio_Content{
			Content: *audioBytes,
		},
	}
	config := &speechpb.RecognitionConfig{
		Encoding:        speechpb.RecognitionConfig_LINEAR16,
		SampleRateHertz: 44100,
		LanguageCode:    languageCode,
	}

	resp, err := client.Recognize(ctx, &speechpb.RecognizeRequest{
		Config: config,
		Audio:  audioContent,
	})
	if err != nil {
		log.Fatalf("Failed to recognize speech: %v", err)
	}

	var transcription string

	// gcp api sometimes return multiple same values, need to eliminate it
	seenTexts := make(map[string]bool)

	for _, result := range resp.Results {
		if len(result.Alternatives) > 0 {
			bestAlternative := result.Alternatives[0]
			text := strings.TrimSpace(bestAlternative.Transcript)
			if !seenTexts[text] {
				fmt.Printf("Recognized text: %s\n", text)
				transcription += text + "\n"
				seenTexts[text] = true
			}
		}
	}
	audio.SaveTranscription(strings.TrimRight(filename, ".wav"), transcription)
}
