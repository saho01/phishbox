package phishbox

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

type httpRequest struct {
	Client     *http.Client
	URL        string
	Headers    map[string]string
	Parameters map[string]string
}

type httpResponse struct {
	ResponseCode int              `json:"response_code"`
	Success      bool             `json:"success"`
	ErrorMessage string           `json:"error_message"`
	Data         *json.RawMessage `json:"data"`
}

type payload interface {
	getValues() []keyValuePair
	getPayloadBuffer() (*bytes.Buffer, error)
}

type keyValuePair struct {
	key   string
	value string
}

type urlParams struct {
	Values []keyValuePair
}

type requestBody struct {
	Values []keyValuePair
}

func newURLParams() *urlParams {
	return &urlParams{}
}

func (u *urlParams) addValue(key, value string) {
	u.Values = append(u.Values, keyValuePair{key: key, value: value})
}

func (u *urlParams) getValues() []keyValuePair {
	return u.Values
}

func (u *urlParams) getPayloadBuffer() (*bytes.Buffer, error) {
	data := url.Values{}
	for _, keyVal := range u.Values {
		data.Add(keyVal.key, keyVal.value)
	}
	return bytes.NewBufferString(data.Encode()), nil
}

func newRequestBody() *requestBody {
	return &requestBody{}
}

func (r *requestBody) addValue(key, value string) {
	r.Values = append(r.Values, keyValuePair{key: key, value: value})
}

func (r *requestBody) getValues() []keyValuePair {
	return r.Values
}

func (r *requestBody) getPayloadBuffer() (*bytes.Buffer, error) {
	m := make(map[string]interface{})

	for _, keyVal := range r.Values {
		m[keyVal.key] = keyVal.value
	}
	b, err := json.Marshal(m)
	if err != nil {
		return nil, err
	}
	return bytes.NewBuffer(b), nil
}

func newHTTPRequest(url string) *httpRequest {
	return &httpRequest{
		URL:    url,
		Client: http.DefaultClient,
	}
}

func (h *httpRequest) setHTTPClient(c *http.Client, apiKey string) {
	h.Client = c
	h.addHeader("api-token", apiKey)
}

func (h *httpRequest) addHeader(header, value string) {
	if h.Headers == nil {
		h.Headers = make(map[string]string)
	}
	h.Headers[header] = value
}

func (h *httpRequest) makeRequest(ctx context.Context, method string, payload payload) (*httpResponse, error) {
	var body io.Reader
	var err error
	if payload != nil {
		if body, err = payload.getPayloadBuffer(); err != nil {
			return nil, err
		}
	} else {
		body = nil
	}

	req, err := http.NewRequest(method, h.URL, body)
	if err != nil {
		return nil, err
	}

	for h, v := range h.Headers {
		// server doesnt like capitalized header keys
		req.Header[h] = []string{v}
	}
	req = req.WithContext(ctx)

	resp, err := h.Client.Do(req)
	defer resp.Body.Close()
	if err != nil {
		return nil, err
	}
	var response httpResponse
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return nil, err
	}
	if !response.Success && response.ResponseCode != 200 {
		return nil, fmt.Errorf("%s", response.ErrorMessage)
	}
	return &response, nil
}

func (h *httpRequest) getRequest(ctx context.Context, payload payload) (*httpResponse, error) {
	return h.makeRequest(ctx, "GET", payload)
}

func (h *httpRequest) postRequest(ctx context.Context, payload payload) (*httpResponse, error) {
	return h.makeRequest(ctx, "POST", payload)
}

func (h *httpResponse) parseJSON(i interface{}) error {
	return json.Unmarshal(*h.Data, i)
}
