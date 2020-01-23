package phishbox

import (
	"context"
)

// BaseTest contains common test fields
type BaseTest struct {
	SentID                     string `json:"sent_id,omitempty"`
	AccountID                  string `json:"account_id,omitempty"`
	CampaignID                 string `json:"campaign_id,omitempty"`
	Active                     string `json:"active,omitempty"`
	TestName                   string `json:"test_name,omitempty"`
	DateCreated                string `json:"date_created,omitempty"`
	DateAuthoriztionRequested  string `json:"date_authorization_requested,omitempty"`
	DateAuthorized             string `json:"date_authorized,omitempty"`
	DateStarted                string `json:"date_started,omitempty"`
	DateStartedPbox            string `json:"date_started_pbox,omitempty"`
	DateStartedGMT             string `json:"date_started_gmt,omitempty"`
	DateStartedTimestamp       string `json:"date_started_timestamp,omitempty"`
	DateSendStart              string `json:"date_send_start,omitempty"`
	DateSendEndedTimestamp     string `json:"date_send_ended_timestamp,omitempty"`
	DateEnded                  string `json:"date_ended,omitempty"`
	DateEndedPbox              string `json:"date_ended_pbox,omitempty"`
	DateEndedGMT               string `json:"date_ended_gmt,omitempty"`
	DateEndedTimestamp         string `json:"date_ended_timestamp,omitempty"`
	DateTimezone               string `json:"date_timezone,omitempty"`
	RecurringType              string `json:"recurring_type,omitempty"`
	RecurringLength            string `json:"recurring_length,omitempty"`
	RecurringPeriod            string `json:"recurring_period,omitempty"`
	RecurringSpawn             string `json:"recurring_spawn,omitempty"`
	SendLength                 string `json:"send_length,omitempty"`
	SendPeriod                 string `json:"send_period,omitempty"`
	BusinessDayStart           string `json:"business_day_start,omitempty"`
	BusinessDayEnd             string `json:"business_day_end,omitempty"`
	BusinessDays               string `json:"business_days,omitempty"`
	TrackLength                string `json:"track_length,omitempty"`
	TrackPeriod                string `json:"track_period,omitempty"`
	DateSent                   string `json:"date_sent,omitempty"`
	TestSendEmail              string `json:"test_send_email,omitempty"`
	SendToThese                string `json:"send_to_these,omitempty"`
	EmailsSentTo               string `json:"emails_sent_to,omitempty"`
	Duplicates                 string `json:"duplicates,omitempty"`
	Errors                     string `json:"errors,omitempty"`
	GroupID                    string `json:"group_id,omitempty"`
	IndividualTargetSelection  string `json:"individual_target_selection,omitempty"`
	MessageType                string `json:"message_type,omitempty"`
	CurrentStep                string `json:"current_step,omitempty"`
	SetupComplete              string `json:"setup_complete,omitempty"`
	IsComplete                 string `json:"is_complete,omitempty"`
	TestType                   string `json:"test_type,omitempty"`
	NotifyFrequency            string `json:"notify_frequency,omitempty"`
	NotifyEmail                string `json:"notify_email,omitempty"`
	EmailRateLimit             string `json:"email_rate_limit,omitempty"`
	AutoEnrollOnFail           string `json:"auto_enroll_on_fail,omitempty"`
	WebhookURL                 string `json:"webhook_url,omitempty"`
	CoursesToBeEnrolledTo      string `json:"courses_to_be_enrolled_to,omitempty"`
	SmarteruToBeEnrolledTo     string `json:"smarteru_to_be_enrolled_to,omitempty"`
	MoodleToBeEnrolledTo       string `json:"moodle_to_be_enrolled_to,omitempty"`
	Sufficient                 string `json:"sufficient,omitempty"`
	BridgeToBeEnrolledTo       string `json:"bridge_to_be_enrolled_to,omitempty"`
	CourseIDToBeAutoEnrolledTo string `json:"course_id_to_be_auto_enrolled_to,omitempty"`
	AttachFile                 string `json:"attach_file,omitempty"`
	Domain                     string `json:"domain,omitempty"`
	LinkDomain                 string `json:"link_domain,omitempty"`
	AttachedImage              string `json:"attached_image,omitempty"`
	AttachedImageLinkDomain    string `json:"attached_image_link_domain,omitempty"`
	AttachedHTML               string `json:"attached_html,omitempty"`
	AttachedHTMLLinkDomain     string `json:"attached_html_link_domain,omitempty"`
	ClientName                 string `json:"client_name,omitempty"`
	URLString                  string `json:"url_string,omitempty"`
	LogoDefault                string `json:"logo_default,omitempty"`
	LogoID                     string `json:"logo_id,omitempty"`
	LogoSMID                   string `json:"logo_sm_id,omitempty"`
	AuthorizedDomain           string `json:"authorized_domain,omitempty"`
	DateInstructedToSend       string `json:"date_instructed_to_send,omitempty"`
	SMTPCustomRelay            string `json:"smtp_custom_relay,omitempty"`
	SMTPUser                   string `json:"smtp_user,omitempty"`
	SMTPHost                   string `json:"smtp_host,omitempty"`
	SendAsFrom                 string `json:"send_as_from,omitempty"`
	SendReplyTo                string `json:"send_reply_to,omitempty"`
	IntroPageText              string `json:"intro_page_text,omitempty"`
	FooterText                 string `json:"footer_text,omitempty"`
	UsernameFieldName          string `json:"username_field_name,omitempty"`
	PasswordFieldName          string `json:"password_field_name,omitempty"`
	SubmitFieldName            string `json:"submit_field_name,omitempty"`
	Subject                    string `json:"subject,omitempty"`
	Message                    string `json:"message,omitempty"`
	TemplateID                 string `json:"template_id,omitempty"`
	TemplateIDs                string `json:"template_ids,omitempty"`
}

// Test cotains items for Test components
type Test struct {
	BaseTest
	TemplateOptions Template     `json:"template_options"`
	TargetCount     string       `json:"target_count,omitempty"`
	Targets         []TestTarget `json:"targets"`
}

// ListTests gets list of tests
func (p *PhishBox) ListTests(ctx context.Context, accountID string) ([]Test, error) {
	var tests []Test

	params := newURLParams()
	params.addValue(accountIDParam, accountID)
	url, err := generateParameterizedURL(p, testEndpoint, allAction, params)
	if err != nil {
		return tests, err
	}

	req := newHTTPRequest(url)
	req.setHTTPClient(p.Client(), p.APIKey())

	resp, err := req.getRequest(ctx, nil)
	if err != nil {
		return tests, err
	}

	err = resp.parseJSON(&tests)
	if err != nil {
		return tests, err
	}

	return tests, nil
}
