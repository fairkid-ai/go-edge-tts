package edgetts

import (
	"encoding/json"
	"io"
	"net/http"
)

type VoiceTag struct {
	ContentCategories  []string `json:"ContentCategories"`
	VoicePersonalities []string `json:"VoicePersonalities"`
}

type Voice struct {
	Name           string `json:"Name"`
	ShortName      string `json:"ShortName"`
	Gender         string `json:"Gender"`
	Locale         string `json:"Locale"`
	SuggestedCodec string `json:"SuggestedCodec"`
	FriendlyName   string `json:"FriendlyName"`
	Status         string `json:"Status"`
	Language       string
	VoiceTag       VoiceTag `json:"VoiceTag"`
}

func ListVoices() ([]Voice, error) {
	// Send GET request to retrieve the list of voices.
	client := http.Client{}
	req, err := http.NewRequest("GET", VOICE_LIST_URL, nil)
	if err != nil {
		return nil, err
	}
	req.Header = gVoiceHeaders

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Read the response body.
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// Parse the JSON response.
	var voices []Voice
	err = json.Unmarshal(body, &voices)
	if err != nil {
		return nil, err
	}

	return voices, nil
}
