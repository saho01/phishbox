package phishbox

import (
	"fmt"
	"net/http"
)

const (
	// APIBase base url of api
	APIBase              = ""
	allAction            = "all"
	deleteAction         = "delete"
	cancelAction         = "cancel"
	createAction         = "create"
	getAction            = "get"
	statusAction         = "status"
	updateAction         = "update"
	accountEndpoint      = "Account"
	groupEndpoint        = "Group"
	reportEndpoint       = "Report"
	targetEndpoint       = "Target"
	templateEndpoint     = "Template"
	testEndpoint         = "Test"
	trainingPageEndpoint = "TrainingPage"
	accountIDParam       = "account_id"
	testIDParam          = "test_id"
	groupIDParam         = "group_id"
	templateIDParam      = "template_id"
	emailIDParam         = "email_id"
	trainingPageIDParam  = "training_page_id"
)

// Phishboxer Base interface
type Phishboxer interface {
	APIBase() string
	APIKey() string
	Client() *http.Client
	SetAPIBase(url string)
	SetAPIKey(key string)
	SetClient(client *http.Client)
	ListGroups()
}

// PhishBox main struct for phishingbox api items
type PhishBox struct {
	apiBase string
	apiKey  string
	client  *http.Client
}

// NewPhishBox creates a new gophish client instance
func NewPhishBox(apikey string) *PhishBox {
	return &PhishBox{
		apiBase: APIBase,
		apiKey:  apikey,
		client:  http.DefaultClient,
	}
}

// APIBase returns base api url
func (p *PhishBox) APIBase() string {
	return p.apiBase
}

// APIKey returns api key
func (p *PhishBox) APIKey() string {
	return p.apiKey
}

// Client returns http client
func (p *PhishBox) Client() *http.Client {
	return p.client
}

// SetAPIBase sets base api url
func (p *PhishBox) SetAPIBase(a string) {
	p.apiBase = a
}

// SetAPIKey sets api key
func (p *PhishBox) SetAPIKey(k string) {
	p.apiKey = k
}

// SetClient sets http client
func (p *PhishBox) SetClient(c *http.Client) {
	p.client = c
}

func generateAPIUrl(p *PhishBox, endpoint, action string) string {
	return fmt.Sprintf("%s/%s/%s", p.APIBase(), endpoint, action)
}

func generateParameterizedURL(p *PhishBox, endpoint, action string, payload payload) (string, error) {
	paramBuffer, err := payload.getPayloadBuffer()
	if err != nil {
		return "", err
	}
	params := string(paramBuffer.Bytes())
	return fmt.Sprintf("%s?%s", generateAPIUrl(p, endpoint, action), params), nil
}
