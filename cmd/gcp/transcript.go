package gcp

import (
	"context"
	"fmt"
	"log"
	"speech_to_text/cmd/audio"
	"strings"

	speech "cloud.google.com/go/speech/apiv1"
	"cloud.google.com/go/speech/apiv1/speechpb"
)

func TranscriptFile(filename string) {

	audioBytes := audio.ReadWavFile(filename)

	ctx := context.Background()
	client, err := speech.NewClient(ctx)
	if err != nil {
		log.Fatalf("Failed to create speech client: %v", err)
	}
	defer client.Close()

	audioContent := &speechpb.RecognitionAudio{
		AudioSource: &speechpb.RecognitionAudio_Content{
			Content: audioBytes,
		},
	}
	config := &speechpb.RecognitionConfig{
		Encoding:        speechpb.RecognitionConfig_LINEAR16,
		SampleRateHertz: 44100,
		LanguageCode:    "en-US",
	}

	resp, err := client.Recognize(ctx, &speechpb.RecognizeRequest{
		Config: config,
		Audio:  audioContent,
	})
	if err != nil {
		log.Fatalf("Failed to recognize speech: %v", err)
	}

	var transcription string
	for _, result := range resp.Results {
		for _, alt := range result.Alternatives {
			fmt.Printf("Recognized text: %s\n", alt.Transcript)
			transcription += alt.Transcript
			transcription += "\n"
		}
	}
	audio.SaveTranscription(strings.TrimRight(filename, ".wav"), transcription)
}
