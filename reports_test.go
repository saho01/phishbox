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
	getReportResult = `
	{
	  "response_code": 200,
	  "success": true,
	  "data": {
		"sent_id": "10000",
		"account_id": null,
		"campaign_id": "1000",
		"active": "1",
		"test_name": "Test #1",
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
		"uuid": null,
		"name": null,
		"datetime_created": null,
		"datetime_updated": null,
		"url_string": "",
		"intro_page_text": null,
		"has_text": "No",
		"footer_text": null,
		"logo_resize": "0",
		"image_id": null,
		"image_sm_id": null,
		"username_field_name": null,
		"password_field_name": null,
		"submit_field_name": null,
		"authorized_domain": null,
		"smtp_custom_relay": "0",
		"smtp_user": null,
		"smtp_pass": null,
		"smtp_host": null,
		"send_as_from": null,
		"send_reply_to": null,
		"login_template_id": null,
		"install_template_id": null,
		"is_active": "1",
		"optional_field_1": null,
		"optional_field_2": null,
		"optional_field_3": null,
		"service_type": null,
		"service_configuration": null,
		"default_course_configuration": null,
		"moved_account_id": null,
		"moved_group_id": null,
		"auto_sync": "0",
		"smart_sync": "0",
		"allow_auto_enroll": "1",
		"id": null,
		"predesigned": null,
		"clone": null,
		"ever_saved": null,
		"type": null,
		"landing_type": null,
		"description": null,
		"email_from_email": null,
		"email_from_name": null,
		"email_replyto_email": null,
		"email_subject": null,
		"email_body": "",
		"destination_title": null,
		"destination_body": "",
		"smtp_username": null,
		"smtp_server": null,
		"smtp_port": null,
		"smtp_encryption": null,
		"bodybackground": null,
		"landingpage_bodybackground": null,
		"landingpage_bootstrap": null,
		"cname": null,
		"click_through_failure": null,
		"submission_redirect_url": null,
		"submission_thankyou_text": null,
		"submission_action": null,
		"submission_anydatafailure": null,
		"submission_allfieldsrequired": null,
		"lp_url_redirect": null,
		"lp_url_replicate": null,
		"hook_link_text": null,
		"hook_link_text_usage": null,
		"download_option": null,
		"download_filename": null,
		"download_customfile": null,
		"open_tracking": null,
		"training_option": null,
		"training_id": null,
		"custom_attachment_id": null,
		"custom_attachment_filename": null,
		"custom_attachment_original_name": null,
		"community": null,
		"submitted_template_id": null,
		"submit_status": null,
		"submitter_id": null,
		"landingpage_image": null,
		"email_image": null,
		"datetime_imaged": null,
		"destination_editor": null,
		"destination_head": null,
		"destination_body_class": null,
		"track_attachment": null,
		"track_reply": null,
		"template_category": null,
		"email_editor": null,
		"preview_token": null,
		"imap_server": null,
		"imap_port": null,
		"imap_username": null,
		"imap_password": null,
		"imap_encryption": null,
		"imap_folder": null,
		"lib_template_id": null,
		"custom_header_name": null,
		"custom_header_value": null,
		"landing_id": null,
		"error_scan": null,
		"custom_instructions": null,
		"moved_template_id": null,
		"targets_tested": [
		  {
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
			"temp_email_id": "1000000",
			"account_id": "1000",
			"campaign_id": "1000",
			"sent_id": "10000",
			"group_id": "10000",
			"email_id": "1000000",
			"full_name": "Bill Test",
			"first_name": "Test",
			"last_name": "Test",
			"email": "bill.test@test.com",
			"label": "",
			"phish_key": "1111111111111111111111111111111111111111",
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
			"testing_status": ""
		  },
		  {
			"temp_email_id": "1000000",
			"account_id": "1000",
			"campaign_id": "1000",
			"sent_id": "10000",
			"group_id": "10000",
			"email_id": "1000000",
			"full_name": "Jane test",
			"first_name": "Jane",
			"last_name": "Test",
			"email": "jane.test@test.com",
			"label": "",
			"phish_key": "eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee",
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
			"testing_status": ""
		  }
		],
		"targets_failed_details": [
		  {
			"access_id": "1000001",
			"access_date": "2019-11-21 13:25:23",
			"sent_id": "10000",
			"group_id": null,
			"accessed_group_id": "10000",
			"name": null,
			"email": "joe@test.com",
			"type": "opened",
			"ip_address": "66.102.6.212",
			"user_agent": "Mozilla/5.0 (Windows NT 5.1; rv:11.0) Gecko Firefox/11.0 (via ggpht.com GoogleImageProxy)",
			"request_uri": "/phlogo_1e069dedff51a6f71fe08bdeb094467cb825f741.gif",
			"domain": "shipment-status.com",
			"phish_key": "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa",
			"ip_address_num": "66102006212"
		  },
		  {
			"access_id": "1000000",
			"access_date": "2019-11-21 13:25:35",
			"sent_id": "10000",
			"group_id": null,
			"accessed_group_id": "10000",
			"name": null,
			"email": "joe@test.com",
			"type": "opened",
			"ip_address": "10.10.10.10",
			"user_agent": "Mozilla/5.0 (X11; Ubuntu; Linux x86_64; rv:70.0) Gecko/20100101 Firefox/70.0",
			"request_uri": "/route.php?k=1e069dedff51a6f71fe08bdeb094467cb825f741&training_id=3833&group_id=17454&target_id=6032026",
			"domain": "protocol46.testingcenter.net",
			"phish_key": "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa",
			"ip_address_num": "10010010010"
		  }
		]
	  }
	}
	`

	getReportBadJSONResult = `
	{
	  "response_code": 200,
	  "success": true,
	  "data": {
		"sent_id": "10000",
		"custom_instructions": null
		"moved_template_id": null,
		"targets_tested": [
		  {
			"temp_email_id": "1000000",
			"account_id": "1000",
			"campaign_id": "1000",
			"sent_id": "10000",
			"group_id": "10000",
			"email_id": "1000000",
			"full_name": "Jane test",
			"first_name": "Jane",
			"last_name": "Test",
			"email": "jane.test@test.com",
			"label": "",
			"phish_key": "eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee",
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
			"testing_status": ""
		  }
		],
		"targets_failed_details": [
		  {
			"access_id": "1000000",
			"access_date": "2019-11-21 13:25:35",
			"sent_id": "10000",
			"group_id": null,
			"accessed_group_id": "10000",
			"name": null,
			"email": "joe@test.com",
			"type": "opened",
			"ip_address": "10.10.10.10",
			"user_agent": "Mozilla/5.0 (X11; Ubuntu; Linux x86_64; rv:70.0) Gecko/20100101 Firefox/70.0",
			"request_uri": "/route.php?k=1e069dedff51a6f71fe08bdeb094467cb825f741&training_id=3833&group_id=17454&target_id=6032026",
			"domain": "protocol46.testingcenter.net",
			"phish_key": "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa",
			"ip_address_num": "10010010010"
		  }
		]
	  }
	}
	`
	listNotObjectJSON = `{
		"response_code":200,
		"success":true,
		"data":[]
	}`
)

func TestSuccessfulGetReport(t *testing.T) {
	client, server := testClient(http.StatusOK, strings.NewReader(getReportResult))
	defer server.Close()

	p := NewPhishBox(apiKey)
	p.SetClient(client.Client)
	p.SetAPIBase(server.URL)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	report, err := p.GetReport(ctx, "123", "456", "789")
	assert.NoError(t, err, "Should be no error")
	assert.Len(t, report.TargetsTested, 3, "Tested 3 targets")
	assert.Len(t, report.TargetsFailedDetails, 2, "2 failed target details")
	assert.Equal(t, report.TargetsTested[0].Email, "joe@test.com", "emails should be equal")
	assert.Equal(t, report.TargetsFailedDetails[0].AccessID, "1000001", "Access ids should be equal")
}

func TestBADJSONGetReport(t *testing.T) {
	client, server := testClient(http.StatusOK, strings.NewReader(getReportBadJSONResult))
	defer server.Close()

	p := NewPhishBox(apiKey)
	p.SetClient(client.Client)
	p.SetAPIBase(server.URL)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	_, err := p.GetReport(ctx, "123", "456", "789")
	assert.Error(t, err, "Should be error, Bad JSON")
}

func TestUnsuccessfulReportListResult(t *testing.T) {
	client, server := testClient(http.StatusOK, strings.NewReader(listNotObjectJSON))
	defer server.Close()

	p := NewPhishBox(apiKey)
	p.SetClient(client.Client)
	p.SetAPIBase(server.URL)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	_, err := p.GetReport(ctx, "123", "456", "789")
	assert.Error(t, err, "Should be error bad report json, listNotObjectJSON")
}
