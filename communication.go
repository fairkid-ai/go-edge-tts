package edgetts

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"log"

	"github.com/gorilla/websocket"
)

type Chunk struct {
	Type     string
	Data     []byte
	Offset   int
	Duration int
	Text     string
}

type turnMetaInnerText struct {
	Text         string `json:"Text"`
	Length       int    `json:"Length"`
	BoundaryType string `json:"BoundaryType"`
}

type turnMetaInnerData struct {
	Offset   int               `json:"Offset"`
	Duration int               `json:"Duration"`
	Text     turnMetaInnerText `json:"text"`
}

type turnMetadata struct {
	Type string            `json:"Type"`
	Data turnMetaInnerData `json:"Data"`
}

type turnMeta struct {
	Metadata []turnMetadata `json:"Metadata"`
}

type Communication struct {
	option Option
}

func NewCommunication(option Option) *Communication {
	return &Communication{option: option}
}

func (c *Communication) Close() error {
	return nil
}

func (c *Communication) openWebsocket() (*websocket.Conn, error) {
	dialer := websocket.Dialer{}
	conn, _, err := dialer.Dial(WSS_URL+"&ConnectionId="+uuidWithOutDashes(), gWssHeaders)
	return conn, err
}

func (c *Communication) Stream(text string) (chan Chunk, error) {
	date := dateToString()
	conn, err := c.openWebsocket()
	if err != nil {
		return nil, err
	}

	cmdStr := fmt.Sprintf("X-Timestamp:%s\r\nContent-Type:application/json;charset=utf-8\r\nPath:speech.config\r\n\r\n{\"context\":{\"synthesis\":{\"audio\":{\"metadataoptions\":{\"sentenceBoundaryEnabled\":false,\"wordBoundaryEnabled\":true},\"outputFormat\":\"audio-24khz-48kbitrate-mono-mp3\"}}}}\r\n", date)
	if err := conn.WriteMessage(websocket.TextMessage, []byte(cmdStr)); err != nil {
		return nil, err
	}

	ssmlStr := ssmlHeadersPlusData(uuidWithOutDashes(), date, mkssml(
		text, c.option.Voice, c.option.Rate, c.option.Volume))
	if err := conn.WriteMessage(websocket.TextMessage, []byte(ssmlStr)); err != nil {
		return nil, err
	}

	recvChan := make(chan Chunk, 100)
	go c.receive(conn, recvChan)

	return recvChan, nil
}

func (c *Communication) receive(conn *websocket.Conn, recvChan chan Chunk) {
	defer conn.Close()
	defer close(recvChan)

	// download indicates whether we should be expecting audio data,
	// this is so what we avoid getting binary data from the websocket
	// and falsely thinking it's audio data.
	downloadAudio := false

	// audio_was_received indicates whether we have received audio data
	// from the websocket. This is so we can raise an exception if we
	// don't receive any audio data.
	// audioWasReceived := false

	// finalUtterance := make(map[int]int)
	for {
		// read message
		messageType, data, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}

		switch messageType {
		case websocket.TextMessage:
			parameters, data, _ := getHeadersAndData(data)
			path := parameters["Path"]

			switch path {
			case "turn.start":
				downloadAudio = true

			case "turn.end":
				downloadAudio = false
				recvChan <- Chunk{Type: CHUNK_TYPE_END}
				return

			case "audio.metadata":
				var meta turnMeta
				if err := json.Unmarshal(data, &meta); err != nil {
					log.Println("We received a text message, but unmarshal failed:", string(data))
					return
				}
				for _, v := range meta.Metadata {
					switch v.Type {
					case CHUNK_TYPE_WORD_BOUNDARY:
						recvChan <- Chunk{
							Type:     v.Type,
							Offset:   v.Data.Offset,
							Duration: v.Data.Duration,
							Text:     v.Data.Text.Text,
						}

					case CHUNK_TYPE_SESSION_END:
						continue

					default:
						log.Println("Unknown metadata type:", v.Type)
					}
				}
			}

		case websocket.BinaryMessage:
			if !downloadAudio {
				log.Println("We received a binary message, but we are not expecting one.")
				return
			}
			if len(data) < 2 {
				log.Println("We received a binary message, but it is missing the header length.")
				return
			}
			headerLength := int(binary.BigEndian.Uint16(data[:2]))
			if len(data) < headerLength+2 {
				log.Println("We received a binary message, but it is missing the audio data.")
				return
			}
			recvChan <- Chunk{
				Type: CHUNK_TYPE_AUDIO,
				Data: data[headerLength+2:],
			}
		}
	}
}
