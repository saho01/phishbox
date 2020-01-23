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
	allTestsResult = `
		{
		  "response_code": 200,
		  "success": true,
		  "data": [
			{
			  "sent_id": "10000",
			  "account_id": "1000",
			  "campaign_id": "1000",
			  "active": "1",
			  "test_name": "test api Test #1",
			  "date_created": "2019-11-21 13:12:48",
			  "date_authorization_requested": "2019-11-21 18:12:49",
			  "date_authorized": "2019-11-21 18:13:24",
			  "date_started": "2019-11-21 12:21:00",
			  "date_started_pbox": "2019-11-21 13:21:00",
			  "date_started_gmt": "2019-11-21 18:21:00",
			  "date_started_timestamp": "1574360460",
			  "date_send_ended_timestamp": "1574392860",
			  "date_ended": "2019-11-22 21:21:00",
			  "date_ended_pbox": "2019-11-22 22:21:00",
			  "date_ended_gmt": "2019-11-23 03:21:00",
			  "date_ended_timestamp": "1574479260",
			  "date_timezone": "-6",
			  "date_send_start": "2019-11-21 18:21:03",
			  "recurring_type": "one_time",
			  "recurring_length": "1",
			  "recurring_period": null,
			  "recurring_spawn": "0",
			  "send_length": "1",
			  "send_period": "day",
			  "business_day_start": "08:00:00",
			  "business_day_end": "17:00:00",
			  "business_days": "0:1:1:1:1:1:0",
			  "track_length": "1",
			  "track_period": "week",
			  "date_sent": null,
			  "test_send_email": null,
			  "send_to_these": null,
			  "emails_sent_to": "0",
			  "duplicates": "0",
			  "errors": "0",
			  "group_id": "10000",
			  "individual_target_selection": "0",
			  "message_type": "TEXT",
			  "template_id": "0",
			  "template_ids": null,
			  "template_options": null,
			  "current_step": "5",
			  "setup_complete": "1",
			  "is_complete": "1",
			  "test_type": "immediate",
			  "notify_frequency": "none",
			  "notify_email": "",
			  "email_rate_limit": "1000",
			  "auto_enroll_on_fail": "0",
			  "webhook_url": null,
			  "courses_to_be_enrolled_to": "0",
			  "smarteru_to_be_enrolled_to": "0",
			  "moodle_to_be_enrolled_to": "0",
			  "sufficient": "1",
			  "bridge_to_be_enrolled_to": null,
			  "course_id_to_be_auto_enrolled_to": null,
			  "template_options": {
				"type": "Update-Attachment",
				"name": "ExampleTemplate",
				"description": "",
				"email_from_email": "test@example.com",
				"email_from_name": "test@example.com",
				"email_replyto_email": "test@example.com",
				"email_subject": "Example Subject",
				"destination_title": "Update Software",
				"datetime_created": "1970-01-01 13:29:06",
				"datetime_updated": "1970-01-01 08:30:38",
				"cname": "compliance-central.com",
				"click_through_failure": "1",
				"submission_redirect_url": "",
				"submission_thankyou_text": "",
				"submission_action": "Performed Update",
				"submission_anydatafailure": "1",
				"submission_allfieldsrequired": "1",
				"lp_url_redirect": "",
				"lp_url_replicate": "",
				"hook_link_text": "Example Text",
				"hook_link_text_usage": "1",
				"download_option": "message",
				"download_filename": "report",
				"download_customfile": "",
				"open_tracking": "both"
			  },
			  "target_count": "3",
			  "targets": [
				{
				  "failed_action": "opened",
				  "failed": "1",
				  "temp_email_id": "1000000",
				  "account_id": "1000",
				  "campaign_id": "1000",
				  "sent_id": "10000",
				  "group_id": "10000",
				  "email_id": "1000000",
				  "full_name": "Joe Test",
				  "first_name": "Joe",
				  "last_name": "Test",
				  "email": "joe@test.com",
				  "label": "",
				  "phish_key": "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa",
				  "list_type": "NORMAL",
				  "opened_email": null,
				  "opened_num": "0",
				  "is_sent": "yes",
				  "last_ip_address": null,
				  "optional_1": "",
				  "optional_2": "",
				  "optional_3": "",
				  "remove_duplicates": null,
				  "date_sent": "2019-11-21 13:21:03",
				  "template_id": "10000",
				  "send_timestamp": "1574360460",
				  "failed_deliv": "0",
				  "failed_sub": null,
				  "testing_status": "failed"
				},
				{
				  "failed_action": "opened",
				  "failed": "1",
				  "temp_email_id": "1000000",
				  "account_id": "1000",
				  "campaign_id": "1000",
				  "sent_id": "10000",
				  "group_id": "10000",
				  "email_id": "1000000",
				  "full_name": "Joe Test",
				  "first_name": "Joe",
				  "last_name": "Test",
				  "email": "joe@test.com",
				  "label": "",
				  "phish_key": "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa",
				  "list_type": "NORMAL",
				  "opened_email": null,
				  "opened_num": "0",
				  "is_sent": "yes",
				  "last_ip_address": null,
				  "optional_1": "",
				  "optional_2": "",
				  "optional_3": "",
				  "remove_duplicates": null,
				  "date_sent": "2019-11-21 13:21:03",
				  "template_id": "35179",
				  "send_timestamp": "1574360460",
				  "failed_deliv": "0",
				  "failed_sub": null,
				  "testing_status": "failed"
				},
				{
				  "failed_action": "page-load",
				  "failed": "1",
				  "temp_email_id": "1000000",
				  "account_id": "1000",
				  "campaign_id": "1000",
				  "sent_id": "10000",
				  "group_id": "10000",
				  "email_id": "1000000",
				  "full_name": "Joe Test",
				  "first_name": "Joe",
				  "last_name": "Test",
				  "email": "joe@test.com",
				  "label": "",
				  "phish_key": "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa",
				  "list_type": "NORMAL",
				  "opened_email": null,
				  "opened_num": "0",
				  "is_sent": "yes",
				  "last_ip_address": null,
				  "optional_1": "",
				  "optional_2": "",
				  "optional_3": "",
				  "remove_duplicates": null,
				  "date_sent": "2019-11-21 13:21:03",
				  "template_id": "10000",
				  "send_timestamp": "1574360460",
				  "failed_deliv": "0",
				  "failed_sub": null,
				  "testing_status": "failed"
				},
				{
				  "failed_action": "Viewed Training Page",
				  "failed": "1",
				  "temp_email_id": "1000000",
				  "account_id": "1000",
				  "campaign_id": "1000",
				  "sent_id": "10000",
				  "group_id": "10000",
				  "email_id": "1000000",
				  "full_name": "Joe Test",
				  "first_name": "Joe",
				  "last_name": "Test",
				  "email": "joe@test.com",
				  "label": "",
				  "phish_key": "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa",
				  "list_type": "NORMAL",
				  "opened_email": null,
				  "opened_num": "0",
				  "is_sent": "yes",
				  "last_ip_address": null,
				  "optional_1": "",
				  "optional_2": "",
				  "optional_3": "",
				  "remove_duplicates": null,
				  "date_sent": "2019-11-21 13:21:03",
				  "template_id": "10000",
				  "send_timestamp": "1574360460",
				  "failed_deliv": "0",
				  "failed_sub": null,
				  "testing_status": "failed"
				},
				{
				  "failed_action": "opened",
				  "failed": "1",
				  "temp_email_id": "1768713",
				  "account_id": "1000",
				  "campaign_id": "1000",
				  "sent_id": "10000",
				  "group_id": "10000",
				  "email_id": "1000000",
				  "full_name": "Joe Test",
				  "first_name": "Joe",
				  "last_name": "Test",
				  "email": "joe@test.com",
				  "label": "",
				  "phish_key": "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa",
				  "list_type": "NORMAL",
				  "opened_email": null,
				  "opened_num": "0",
				  "is_sent": "yes",
				  "last_ip_address": null,
				  "optional_1": "",
				  "optional_2": "",
				  "optional_3": "",
				  "remove_duplicates": null,
				  "date_sent": "2019-11-21 13:21:03",
				  "template_id": "10000",
				  "send_timestamp": "1574360460",
				  "failed_deliv": "0",
				  "failed_sub": null,
				  "testing_status": "failed"
				},
				{
				  "failed_action": null,
				  "failed": "0",
				  "temp_email_id": "1000000",
				  "account_id": "1000",
				  "campaign_id": "1000",
				  "sent_id": "10000",
				  "group_id": "10000",
				  "email_id": "1000000",
				  "full_name": "Bill Test",
				  "first_name": "Bill",
				  "last_name": "Test",
				  "email": "bill.test@test.com",
				  "label": "",
				  "phish_key": "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa",
				  "list_type": "NORMAL",
				  "opened_email": null,
				  "opened_num": "0",
				  "is_sent": "yes",
				  "last_ip_address": null,
				  "optional_1": "",
				  "optional_2": "",
				  "optional_3": "",
				  "remove_duplicates": null,
				  "date_sent": "2019-11-21 13:21:03",
				  "template_id": "10000",
				  "send_timestamp": "1574360460",
				  "failed_deliv": "0",
				  "failed_sub": null,
				  "testing_status": ""
				},
				{
				  "failed_action": null,
				  "failed": "0",
				  "temp_email_id": "1000000",
				  "account_id": "1000",
				  "campaign_id": "1000",
				  "sent_id": "10000",
				  "group_id": "10000",
				  "email_id": "1000000",
				  "full_name": "Sammy Test",
				  "first_name": "Sammy",
				  "last_name": "Test",
				  "email": "sammy.test@test.com",
				  "label": "",
				  "phish_key": "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa",
				  "list_type": "NORMAL",
				  "opened_email": null,
				  "opened_num": "0",
				  "is_sent": "yes",
				  "last_ip_address": null,
				  "optional_1": "",
				  "optional_2": "",
				  "optional_3": "",
				  "remove_duplicates": null,
				  "date_sent": "2019-11-21 13:21:03",
				  "template_id": "10000",
				  "send_timestamp": "1574360460",
				  "failed_deliv": "0",
				  "failed_sub": null,
				  "testing_status": ""
				}
			  ]
			}
		  ]
		}
	`

	badJSONTestResult = `
		{
		  "response_code": 200,
		  "success": true,
		  "data": [
			{
			  "sent_id": "10000",
			  "account_id": "1000",
			  "campaign_id": "1000",
			  "active": "1",
			  "test_name": "test api Test #1",
			  "date_created": "2019-11-21 13:12:48",
			  "date_authorization_requested": "2019-11-21 18:12:49",
			  "date_authorized": "2019-11-21 18:13:24",
			  "date_started": "2019-11-21 12:21:00",
			  "date_started_pbox": "2019-11-21 13:21:00",
			  "date_started_gmt": "2019-11-21 18:21:00",
			  "date_started_timestamp": "1574360460",
			  "date_send_ended_timestamp": "1574392860",
			  "date_ended": "2019-11-22 21:21:00",
			  "date_ended_pbox": "2019-11-22 22:21:00",
			  "date_ended_gmt": "2019-11-23 03:21:00",
			  "date_ended_timestamp": "1574479260",
			  "date_timezone": "-6",
			  "date_send_start": "2019-11-21 18:21:03",
			  "recurring_type": "one_time",
			  "recurring_length": "1",
			  "recurring_period": null,
			  "recurring_spawn": "0",
			  "send_length": "1",
			  "send_period": "day",
			  "business_day_start": "08:00:00",
			  "business_day_end": "17:00:00",
			  "business_days": "0:1:1:1:1:1:0",
			  "track_length": "1",
			  "track_period": "week",
			  "date_sent": null,
			  "test_send_email": null,
			  "send_to_these": null,
			  "emails_sent_to": "0",
			  "duplicates": "0",
			  "errors": "0",
			  "group_id": "10000",
			  "individual_target_selection": "0",
			  "message_type": "TEXT",
			  "template_id": "0",
			  "template_ids": null,
			  "template_options": null,
			  "current_step": "5",
			  "setup_complete": "1",
			  "is_complete": "1",
			  "test_type": "immediate",
			  "notify_frequency": "none",
			  "notify_email": "",
			  "email_rate_limit": "1000",
			  "auto_enroll_on_fail": "0",
			  "webhook_url": null,
			  "courses_to_be_enrolled_to": "0",
			  "smarteru_to_be_enrolled_to": "0",
			  "moodle_to_be_enrolled_to": "0",
			  "sufficient": "1",
			  "bridge_to_be_enrolled_to": null,
			  "course_id_to_be_auto_enrolled_to": null,
			  "template_options": {
				"type": "Update-Attachment",
				"name": "ExampleTemplate",
				"description": "",
				"email_from_email": "test@example.com",
				"email_from_name": "test@example.com",
				"email_replyto_email": "test@example.com",
				"email_subject": "Example Subject",
				"destination_title": "Update Software",
				"datetime_created": "1970-01-01 13:29:06",
				"datetime_updated": "1970-01-01 08:30:38",
				"cname": "compliance-central.com",
				"click_through_failure": "1",
				"submission_redirect_url": "",
				"submission_thankyou_text": "",
				"submission_action": "Performed Update",
				"submission_anydatafailure": "1",
				"submission_allfieldsrequired": "1",
				"lp_url_redirect": "",
				"lp_url_replicate": "",
				"hook_link_text": "Example Text",
				"hook_link_text_usage": "1",
				"download_option": "message",
				"download_filename": "report",
				"download_customfile": "",
				"open_tracking": "both"
			  },
			  "target_count": "1"
			  "targets": [
				{
				  "failed_action": "opened",
				  "failed": "1",
				  "temp_email_id": "1000000",
				  "account_id": "1000",
				  "campaign_id": "1000",
				  "sent_id": "10000",
				  "group_id": "10000",
				  "email_id": "1000000",
				  "full_name": "Joe Test",
				  "first_name": "Joe",
				  "last_name": "Test",
				  "email": "joe@test.com",
				  "label": "",
				  "phish_key": "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa",
				  "list_type": "NORMAL",
				  "opened_email": null,
				  "opened_num": "0",
				  "is_sent": "yes",
				  "last_ip_address": null,
				  "optional_1": "",
				  "optional_2": "",
				  "optional_3": "",
				  "remove_duplicates": null,
				  "date_sent": "2019-11-21 13:21:03",
				  "template_id": "10000",
				  "send_timestamp": "1574360460",
				  "failed_deliv": "0",
				  "failed_sub": null,
				  "testing_status": "failed"
				}
			  ]
			}
		  ]
		}
	`
	badTestListJSON = `{
		"response_code":200,
		"success":true,
		"data":{}
	}`
)

func TestSuccessfulTestsList(t *testing.T) {
	client, server := testClient(http.StatusOK, strings.NewReader(allTestsResult))
	defer server.Close()

	p := NewPhishBox(apiKey)
	p.SetClient(client.Client)
	p.SetAPIBase(server.URL)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	tests, err := p.ListTests(ctx, "1234")
	assert.NoError(t, err, "Should be no error")
	assert.Equal(t, 1, len(tests), "length of tests should be 1")
	assert.Equal(t, tests[0].Targets[0].Email, "joe@test.com", "emails should be equal")
	assert.Equal(t, tests[0].TemplateOptions.EmailSubject, "Example Subject", "template email subject should be the same")
}

func TestBADJSONTestsList(t *testing.T) {
	client, server := testClient(http.StatusOK, strings.NewReader(badJSONTestResult))
	defer server.Close()

	p := NewPhishBox(apiKey)
	p.SetClient(client.Client)
	p.SetAPIBase(server.URL)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	tests, err := p.ListTests(ctx, "1234")
	assert.Error(t, err, "Should be error, Bad JSON")
	assert.Equal(t, 0, len(tests), "length of tests should be 0")
}

func TestUnsuccessfulTestListResult(t *testing.T) {
	client, server := testClient(http.StatusOK, strings.NewReader(badTestListJSON))
	defer server.Close()

	p := NewPhishBox(apiKey)
	p.SetClient(client.Client)
	p.SetAPIBase(server.URL)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	_, err := p.ListTests(ctx, "1234")
	assert.Error(t, err, "Should be error, badTestListJSON")
}
