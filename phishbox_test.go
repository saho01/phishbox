package phishbox

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

const testURL = "test-url"
const apiKey = "test-key"
const diffAPIKey = "anotherkey"

func TestPhishbox(t *testing.T) {
	p := NewPhishBox(apiKey)

	// assert api key
	assert.Equal(t, p.APIKey(), apiKey, "apikeys should be equal")

	// assert base url
	assert.Equal(t, p.APIBase(), APIBase, "These should be equal")

	// assert different apikey
	p.SetAPIKey(diffAPIKey)
	assert.Equal(t, p.APIKey(), diffAPIKey, "keys should be equal")

	// set base to different url
	p.SetAPIBase(testURL)
	assert.Equal(t, p.APIBase(), testURL, "url should be equal")

	// assert client
	c := new(http.Client)
	p.SetClient(c)
	assert.Equal(t, p.Client(), c, "http client should be equal")
}

func TestGenerateAPI(t *testing.T) {
	p := NewPhishBox(apiKey)
	url := generateAPIUrl(p, groupEndpoint, allAction)
	assert.Equal(t, url, fmt.Sprintf("%s/%s/%s", p.APIBase(), "Group", "all"), "URL's should be equal")
}

func TestGenerateParameterizedURL(t *testing.T) {
	pb := NewPhishBox(apiKey)
	p := newURLParams()
	p.addValue(accountIDParam, "1234")
	url, err := generateParameterizedURL(pb, groupEndpoint, allAction, p)
	assert.Equal(t, url, fmt.Sprintf("%s/%s/%s?%s", pb.APIBase(), "Group", "all", "account_id=1234"), "URL's should be equal")
	assert.NoError(t, err, "should not be error")
}
