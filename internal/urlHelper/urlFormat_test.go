package urlhelper

import (
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRequestURLFormatter(t *testing.T) {
	interf := "ISteamNews"
	endpoint := "GetNewsForApp"
	version := "0002"

	urlValues := url.Values{}

	expectedURL := "https://api.steampowered.com/ISteamNews/GetNewsForApp/v0002?"
	endPoint := VersionedURLEndpoint{EndpointPath: endpoint, Version: version}

	t.Run("Test URL without Params", func(t *testing.T) {
		result := RequestURLFormatter(interf, endPoint, urlValues)
		if result != expectedURL {
			t.Errorf("expected %s but got %s", expectedURL, result)
		}
		assert.Equal(t, expectedURL, result, "The URLs should be the same")
	})
}
