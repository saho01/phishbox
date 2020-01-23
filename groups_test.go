package phishbox

import (
	"context"
	"net/http"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var (
	allGroupsResult = `{
		"response_code": 200,
		"success": true,
		"data": [
			{
				"group_id": "0",
				"account_id": "0",
				"name": "my group",
				"is_active": "1",
				"optional_field_1": "",
				"optional_field_2": "",
				"optional_field_3": "",
				"service_type": null
			}
	]}`
	badJSON = `{
		"response_code": 200,
		"success": true,
		"data": [
				"group_id": "0",
				"account_id": "0",
				"name": "my group",
				"is_active": "1",
				"optional_field_1": "",
				"optional_field_2": "",
				"optional_field_3": "",
				"service_type": null
			}
	]}`
	badGroupJSON = `{
		"response_code": 200,
		"success": true,
		"data": {
				"group_id": "0",
				"account_id": "0",
				"name": "my group",
				"is_active": "1",
				"optional_field_1": "",
				"optional_field_2": "",
				"optional_field_3": "",
				"service_type": null
			}
	}`
	unsuccessfulResult = `{
		"response_code":401,
		"success":false,
		"error_message":"Invalid API user token",
		"data":[]
	}`
	createGroupResult = `{
		"response_code": 200,
		"success": true,
		"data": {
			"group_id": "46",
			"account_id": "46",
			"name": "test group",
			"optional_field_1": "",
			"optional_field_2": "",
			"optional_field_3": "",
			"is_active": "1",
			"service_type": null
		}
	}`
)

func TestBadHTTPRequest(t *testing.T) {
	client, server := testClient(http.StatusOK, strings.NewReader(allGroupsResult))
	defer server.Close()
	req := newHTTPRequest("/%zz")
	req.setHTTPClient(client.Client, apiKey)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	_, err := req.makeRequest(ctx, "GET", nil)
	assert.Error(t, err, "should be error with bad url")
}

func TestSuccessful(t *testing.T) {
	client, server := testClient(http.StatusOK, strings.NewReader(allGroupsResult))
	defer server.Close()

	p := NewPhishBox(apiKey)
	p.SetClient(client.Client)
	p.SetAPIBase(server.URL)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	groups, err := p.ListGroups(ctx)
	assert.NoError(t, err, "Should be no error")
	assert.Len(t, groups, 1, "length of groups should be 1")
}

func TestBadJSONError(t *testing.T) {
	client, server := testClient(http.StatusOK, strings.NewReader(badJSON))
	defer server.Close()

	p := NewPhishBox(apiKey)
	p.SetClient(client.Client)
	p.SetAPIBase(server.URL)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	_, err := p.ListGroups(ctx)
	assert.Error(t, err, "Should be error, bad return json")
}

func TestBadGroupJSONError(t *testing.T) {
	client, server := testClient(http.StatusOK, strings.NewReader(badGroupJSON))
	defer server.Close()

	p := NewPhishBox(apiKey)
	p.SetClient(client.Client)
	p.SetAPIBase(server.URL)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	_, err := p.ListGroups(ctx)
	assert.Error(t, err, "Should be error, bad group json")
}

func TestUnsuccessfulResult(t *testing.T) {
	client, server := testClient(http.StatusOK, strings.NewReader(unsuccessfulResult))
	defer server.Close()

	p := NewPhishBox(apiKey)
	p.SetClient(client.Client)
	p.SetAPIBase(server.URL)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	_, err := p.ListGroups(ctx)
	assert.Error(t, err, "Should be error, unsuccessfulResult")
}

func TestCreateGroup(t *testing.T) {
	client, server := testClient(http.StatusOK, strings.NewReader(createGroupResult))
	defer server.Close()

	p := NewPhishBox(apiKey)
	p.SetClient(client.Client)
	p.SetAPIBase(server.URL)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	group, err := p.CreateGroup(ctx, "46", "test group")
	assert.NoError(t, err, "Should be no error")
	assert.Equal(t, group.Name, "test group", "group names should be equal")
	assert.Equal(t, group.GroupID, "46", "group id should be same")
}

func TestCreateGroupBadJSON(t *testing.T) {
	client, server := testClient(http.StatusOK, strings.NewReader(badJSON))
	defer server.Close()

	p := NewPhishBox(apiKey)
	p.SetClient(client.Client)
	p.SetAPIBase(server.URL)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	_, err := p.CreateGroup(ctx, "46", "test group")
	assert.Error(t, err, "Should be error")
}

func TestCreateGroupBadGroupJSON(t *testing.T) {
	client, server := testClient(http.StatusOK, strings.NewReader(listNotObjectJSON))
	defer server.Close()

	p := NewPhishBox(apiKey)
	p.SetClient(client.Client)
	p.SetAPIBase(server.URL)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	_, err := p.CreateGroup(ctx, "46", "test group")
	assert.Error(t, err, "Should be error")
}
