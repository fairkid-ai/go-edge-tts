package edgetts

const (
	DEFAULT_VOICE           = "en-US-EmmaMultilingualNeural"
	DEFAULT_RATE            = "+0%"
	DEFAULT_VOLUME          = "+0%"
	DEFAULT_PITCH           = "+0Hz"
	DEFAULT_CONNECT_TIMEOUT = 10
	DEFAULT_RECEIVE_TIMEOUT = 60
)

const (
	TRUSTED_CLIENT_TOKEN = "6A5AA1D4EAFF4E9FB37E23D68491D6F4"
	WSS_URL              = "wss://speech.platform.bing.com/consumer/speech/synthesize/readaloud/edge/v1?TrustedClientToken=" + TRUSTED_CLIENT_TOKEN
	VOICE_LIST_URL       = "https://speech.platform.bing.com/consumer/speech/synthesize/readaloud/voices/list?trustedclienttoken=" + TRUSTED_CLIENT_TOKEN
)

const (
	CHUNK_TYPE_AUDIO         = "Audio"
	CHUNK_TYPE_WORD_BOUNDARY = "WordBoundary"
	CHUNK_TYPE_SESSION_END   = "SessionEnd"
	CHUNK_TYPE_END           = "ChunkEnd"
)

var gDefaultOption = Option{
	Voice:          DEFAULT_VOICE,
	Rate:           DEFAULT_RATE,
	Volume:         DEFAULT_VOLUME,
	Pitch:          DEFAULT_PITCH,
	ConnectTimeout: DEFAULT_CONNECT_TIMEOUT,
	ReceiveTimeout: DEFAULT_RECEIVE_TIMEOUT,
}
