package phishbox

import (
	"context"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var (
	basicResponse = `{
		"response_code": 200,
		"success": true,
		"data": []
	}`
)

func testClient(code int, body io.Reader, validators ...func(*http.Request)) (*httpRequest, *httptest.Server) {

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		for _, v := range validators {
			v(r)
		}
		w.WriteHeader(code)
		io.Copy(w, body)
		r.Body.Close()
		if closer, ok := body.(io.Closer); ok {
			closer.Close()
		}
	}))
	client := newHTTPRequest(testURL)
	return client, server
}

func TestBaseMethods(t *testing.T) {
	h := httpRequest{
		Client: &http.Client{
			Timeout: time.Duration(int64(1)),
		},
		URL: testURL,
	}
	req := newHTTPRequest(testURL)
	assert.Equal(t, req.URL, testURL, "URL should be equal")
	assert.NotEqual(t, req.Client, h.Client, "Clients should not be equal")
	assert.NotEqual(t, req.Headers["api-token"], apiKey, "Headers should not be equal")

	req.setHTTPClient(h.Client, apiKey)
	assert.Equal(t, req.Client, h.Client, "Clients should be equal")
	assert.Equal(t, req.Headers["api-token"], apiKey, "Headers should be equal")

	req.addHeader("api-token", diffAPIKey)
	assert.Equal(t, req.Headers["api-token"], diffAPIKey, "Headers should be equal")
}

func TestGetValues(t *testing.T) {
	p := newURLParams()
	p.addValue("key", "value")
	v := p.getValues()
	assert.Equal(t, v, []keyValuePair{keyValuePair{key: "key", value: "value"}})
	assert.Len(t, v, 1, "length should be 1")

	pl, err := p.getPayloadBuffer()
	assert.NoError(t, err, "should be no error")
	assert.Equal(t, pl.String(), "key=value", "payload string should be equal")
}

func TestRequestBody(t *testing.T) {
	r := newRequestBody()
	v := r.getValues()
	assert.Len(t, v, 0, "No values added yet")
	r.addValue("account_id", "1234")
	assert.Len(t, r.Values, 1, "should be 1 set of key/value")
	r.addValue("name", "test name")
	assert.Len(t, r.Values, 2, "should be 2 set of key/value")
	b, err := r.getPayloadBuffer()
	assert.NoError(t, err, "should be no err on getPayloadBuffer")
	assert.JSONEq(t, b.String(), `{"account_id":"1234","name":"test name"}`, "json strings should equal")
}

func TestPostRequest(t *testing.T) {
	client, server := testClient(http.StatusOK, strings.NewReader(basicResponse))
	defer server.Close()
	req := newHTTPRequest(server.URL)
	req.setHTTPClient(client.Client, apiKey)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	resp, err := req.postRequest(ctx, nil)
	assert.NoError(t, err, "should be no error in post request")
	assert.Equal(t, resp.ResponseCode, 200, "Should be 200 response")
}
