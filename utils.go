package edgetts

import (
	"bytes"
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
)

func uuidWithOutDashes() string {
	return strings.ReplaceAll(uuid.New().String(), "-", "")
}

func dateToString() string {
	return time.Now().UTC().Format("Mon Jan 02 2006 15:04:05 GMT-0700 (Coordinated Universal Time)")
}

func mkssml(text string, voice string, rate string, volume string) string {
	ssml := fmt.Sprintf("<speak version='1.0' xmlns='http://www.w3.org/2001/10/synthesis' xml:lang='en-US'><voice name='%s'><prosody pitch='+0Hz' rate='%s' volume='%s'>%s</prosody></voice></speak>", voice, rate, volume, text)
	return ssml
}

func ssmlHeadersPlusData(requestID string, timestamp string, ssml string) string {
	return fmt.Sprintf("X-RequestId:%s\r\nContent-Type:application/ssml+xml\r\nX-Timestamp:%sZ\r\nPath:ssml\r\n\r\n%s", requestID, timestamp, ssml)
}

func getHeadersAndData(dataBytes []byte) (map[string]string, []byte, error) {
	headers := make(map[string]string)
	lines := bytes.Split(dataBytes[:bytes.Index(dataBytes, []byte("\r\n\r\n"))], []byte("\r\n"))
	for _, line := range lines {
		parts := bytes.SplitN(line, []byte(":"), 2)
		if len(parts) < 2 {
			continue
		}
		key := string(parts[0])
		value := strings.TrimSpace(string(parts[1]))
		headers[key] = value
	}

	return headers, dataBytes[bytes.Index(dataBytes, []byte("\r\n\r\n"))+4:], nil
}
