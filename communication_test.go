package edgetts

import (
	"os"
	"testing"
)

func TestCommunication(t *testing.T) {
	option := DefaultOption()
	option.Voice = "en-US-AnaNeural"

	c := NewCommunication(option)
	ch, err := c.Stream(`Enhanced throughput and concurrent ray tracing and shading capabilities improve ray tracing performance, 
	accelerating renders for product design and architecture, engineering, and construction workflows. 
	See lifelike designs in action with hardware accelerated motion blur to deliver stunning real-time animations.`)
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

	if err := os.WriteFile(".local/nvidia.mp3", audioBytes, 0644); err != nil {
		t.Error(err)
		return
	}
}
