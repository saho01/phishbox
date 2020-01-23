package phishbox

import (
	"context"
)

// Group holds group info
type Group struct {
	GroupID        string `json:"group_id,omitempty"`
	AccountID      string `json:"account_id,omitempty"`
	Name           string `json:"name,omitempty"`
	IsActive       string `json:"is_active,omitempty"`
	OptionalField1 string `json:"optional_field_1,omitempty"`
	OptionalField2 string `json:"optional_field_2,omitempty"`
	OptionalField3 string `json:"optional_field_3,omitempty"`
	ServiceType    string `json:"service_type,omitempty"`
}

// ListGroups gets list of groups
func (p *PhishBox) ListGroups(ctx context.Context) ([]Group, error) {
	var groups []Group
	req := newHTTPRequest(generateAPIUrl(p, groupEndpoint, allAction))
	req.setHTTPClient(p.Client(), p.APIKey())

	resp, err := req.getRequest(ctx, nil)
	if err != nil {
		return groups, err
	}

	err = resp.parseJSON(&groups)
	if err != nil {
		return groups, err
	}

	return groups, nil
}

// CreateGroup creates a group
func (p *PhishBox) CreateGroup(ctx context.Context, accountID, name string) (Group, error) {
	var group Group
	b := newRequestBody()
	b.addValue("account_id", accountID)
	b.addValue("name", name)
	req := newHTTPRequest(generateAPIUrl(p, groupEndpoint, createAction))
	req.setHTTPClient(p.Client(), p.APIKey())

	resp, err := req.postRequest(ctx, b)
	if err != nil {
		return group, err
	}

	err = resp.parseJSON(&group)
	if err != nil {
		return group, err
	}

	return group, nil
}
