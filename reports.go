package phishbox

import (
	"context"
)

// Report contains report fields from get report api call
type Report struct {
	BaseTest
	BaseTemplate
	Template
	HasText                   string                `json:"has_text,omitempty"`
	LogoResize                string                `json:"logo_resize,omitempty"`
	ImageID                   string                `json:"image_id,omitempty"`
	ImageSMID                 string                `json:"image_sm_id,omitempty"`
	LoginTemplateID           string                `json:"login_template_id,omitempty"`
	InstallTemplateID         string                `json:"install_template_id,omitempty"`
	IsActive                  string                `json:"is_active,omitempty"`
	ServiceType               string                `json:"service_type,omitempty"`
	ServiceConfiguration      string                `json:"service_configuration,omitempty"`
	DefautCourseConfiguration string                `json:"default_course_configuration,omitempty"`
	MovedGroupID              string                `json:"moved_group_id,omitempty"`
	AutoSync                  string                `json:"auto_sync,omitempty"`
	SmartSync                 string                `json:"smart_sync,omitempty"`
	AllowAutoEnroll           string                `json:"allow_auto_enroll,omitempty"`
	TargetsTested             []TestTarget          `json:"targets_tested"`
	TargetsFailedDetails      []TargetFailedDetails `json:"targets_failed_details"`
}

// GetReport gets a report
func (p *PhishBox) GetReport(ctx context.Context, accountID, testID, groupID string) (Report, error) {
	var report Report

	params := newURLParams()
	params.addValue(accountIDParam, accountID)
	params.addValue(testIDParam, testID)
	params.addValue(groupIDParam, groupID)
	url, err := generateParameterizedURL(p, reportEndpoint, getAction, params)
	if err != nil {
		return report, err
	}

	req := newHTTPRequest(url)
	req.setHTTPClient(p.Client(), p.APIKey())

	resp, err := req.getRequest(ctx, nil)
	if err != nil {
		return report, err
	}

	err = resp.parseJSON(&report)
	if err != nil {
		return report, err
	}

	return report, nil
}
