package main

import (
	"log"
	"os"

	edgetts "github.com/fairkid-ai/go-edge-tts"
)

func main() {
	text := `(From Wikipedia)
	Speech synthesis is the artificial production of human speech.
	A computer system used for this purpose is called a speech synthesizer, 
	and can be implemented in software or hardware products. 
	A text-to-speech (TTS) system converts normal language text into speech; 
	other systems render symbolic linguistic representations like phonetic transcriptions into speech.
	The reverse process is speech recognition.`

	option := edgetts.DefaultOption()
	option.Voice = "en-US-AnaNeural"

	c := edgetts.NewCommunication(option)
	ch, err := c.Stream(text)
	if err != nil {
		log.Fatal(err)
	}

	audioBytes := make([]byte, 0, 128)
	for chunk := range ch {
		if chunk.Type == edgetts.CHUNK_TYPE_AUDIO {
			audioBytes = append(audioBytes, chunk.Data...)
		}
	}

	if err := os.WriteFile("tts.mp3", audioBytes, 0644); err != nil {
		log.Fatal(err)
	}
}
