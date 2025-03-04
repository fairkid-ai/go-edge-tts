package edgetts

import "net/http"

var (
	gWssHeaders   http.Header
	gVoiceHeaders http.Header
)

func init() {
	// Base headers
	baseHeaders := http.Header{}
	baseHeaders.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/130.0.0.0 Safari/537.36 Edg/130.0.0.0")
	baseHeaders.Add("Accept-Encoding", "gzip, deflate, br")
	baseHeaders.Add("Accept-Language", "en-US,en;q=0.9")

	// Websocket headers
	wssHeaders := baseHeaders.Clone()
	wssHeaders.Add("Pragma", "no-cache")
	wssHeaders.Add("Cache-Control", "no-cache")
	wssHeaders.Add("Origin", "chrome-extension://jdiccldimpdaibmpdkjnbmckianbfold")
	gWssHeaders = wssHeaders

	// Voice headers
	voiceHeaders := baseHeaders.Clone()
	voiceHeaders.Add("Authority", "speech.platform.bing.com")
	voiceHeaders.Add("Sec-CH-UA", `" Not;A Brand";v="99", "Microsoft Edge";v="130", "Chromium";v="130"`)
	voiceHeaders.Add("Sec-CH-UA-Mobile", "?0")
	voiceHeaders.Add("Content-Type", "application/json")
	voiceHeaders.Add("Accept", "*/*")
	voiceHeaders.Add("Sec-Fetch-Site", "none")
	voiceHeaders.Add("Sec-Fetch-Mode", "cors")
	voiceHeaders.Add("Sec-Fetch-Dest", "empty")

	gVoiceHeaders = voiceHeaders
}
