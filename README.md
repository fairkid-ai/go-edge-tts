# go-edge-tts [![GoDoc](https://godoc.org/github.com/fairkid-ai/go-edge-tts?status.svg)](https://godoc.org/github.com/fairkid-ai/go-edge-tts)
Golang implementation to use Microsoft Edge's online Text-To-Speech (TTS) service without an API key.

## Installation
```bash
go get -u github.com/fairkid-ai/go-edge-tts
```

## Usage
```go
package main

import (
	"os"
	"testing"
)

func main() {
    option := edgetts.DefaultOption()
	option.Voice = "en-US-AnaNeural"

	c := edgetts.NewCommunication(option)
	ch, err := c.Stream("Speech synthesis is the artificial production of human speech.")
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
```

## License
[MIT License](https://github.com/fairkid-ai/go-edge-tts/blob/main/LICENSE)

## Credits
This library takes inspiration from [edge-tts](https://github.com/jncraton/edge-tts) and [edge-tts-go](https://github.com/surfaceyu/edge-tts-go).
