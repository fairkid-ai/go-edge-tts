package edgetts

import (
	"os"
	"testing"
)

func TestCommunication(t *testing.T) {
	option := DefaultOption()
	option.Voice = "en-US-AnaNeural"

	c := NewCommunication(option)
	ch, err := c.Stream(`(From Wikipedia)
	Speech synthesis is the artificial production of human speech.
	A computer system used for this purpose is called a speech synthesizer, 
	and can be implemented in software or hardware products. 
	A text-to-speech (TTS) system converts normal language text into speech; 
	other systems render symbolic linguistic representations like phonetic transcriptions into speech.
	The reverse process is speech recognition.`)
	if err != nil {
		t.Error(err)
		return
	}

	audioBytes := make([]byte, 0, 128)
	for chunk := range ch {
		if chunk.Type == CHUNK_TYPE_AUDIO {
			audioBytes = append(audioBytes, chunk.Data...)
		}
	}
	if err := os.MkdirAll(".local", 0755); err != nil {
		t.Error(err)
		return
	}

	if err := os.WriteFile(".local/tts.mp3", audioBytes, 0644); err != nil {
		t.Error(err)
		return
	}
}
