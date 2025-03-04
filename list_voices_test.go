package edgetts

import (
	"fmt"
	"testing"
)

func testListVoices(t *testing.T) {
	voices, err := ListVoices()
	if err != nil {
		t.Error(err)
	}
	for i, voice := range voices {
		fmt.Printf("%d, Name: %s, ShortName: %s, Gender: %s, Locale: %s, SuggestedCodec: %s, FriendlyName: %s, Status: %s, VoiceTag: %+v\n",
			i, voice.Name, voice.ShortName, voice.Gender, voice.Locale, voice.SuggestedCodec, voice.FriendlyName, voice.Status, voice.VoiceTag)
	}
}
