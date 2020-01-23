package phishbox

// Template contains fields for template api calls
type Template struct {
	Type                        string `json:"type,omitempty"`
	Name                        string `json:"name,omitempty"`
	Description                 string `json:"description,omitempty"`
	EmailFromEmail              string `json:"email_from_email,omitempty"`
	EmailFromName               string `json:"email_from_name,omitempty"`
	EmailReplyToEmail           string `json:"email_replyto_email,omitempty"`
	EmailSubject                string `json:"email_subject,omitempty"`
	DestinationTitle            string `json:"destination_title,omitempty"`
	DatetimeCreated             string `json:"datetime_created,omitempty"`
	DetetimeUpdated             string `json:"datetime_updated,omitempty"`
	Cname                       string `json:"cname,omitempty"`
	ClickThroughFailure         string `json:"click_through_failure,omitempty"`
	SubmissionRedirectURL       string `json:"submission_redirect_url,omitempty"`
	SubmissinoThankyouText      string `json:"submission_thankyou_text,omitempty"`
	SubmissionAction            string `json:"submission_action,omitempty"`
	SubmissionAnyDataFailure    string `json:"submission_anydatafailure,omitempty"`
	SubmissionAllFieldsRequired string `json:"submission_allfieldsrequired,omitempty"`
	LPURLRedirect               string `json:"lp_url_redirect,omitempty"`
	LPURLReplicate              string `json:"lp_url_replicate,omitempty"`
	HookLinkText                string `json:"hook_link_text,omitempty"`
	HookLinkTextUsase           string `json:"hook_link_text_usage,omitempty"`
	DownloadOption              string `json:"download_option,omitempty"`
	DownloadFilename            string `json:"download_filename,omitempty"`
	DownloadCustomFile          string `json:"download_customfile,omitempty"`
	OpenTracking                string `json:"open_tracking,omitempty"`
}

// BaseTemplate contains common template fields
type BaseTemplate struct {
	ID                           string `json:"id,omitempty"`
	UUID                         string `json:"uuid,omitempty"`
	Predesigned                  string `json:"predesigned,omitempty"`
	Clone                        string `json:"clone,omitempty"`
	EverSaved                    string `json:"ever_saved,omitempty"`
	EmailBody                    string `json:"email_body,omitempty"`
	DestinationBody              string `json:"destination_body,omitempty"`
	SMTPUsername                 string `json:"smtp_username,omitempty"`
	SMTPPassword                 string `json:"smtp_password,omitempty"`
	SMTPServer                   string `json:"smtp_server,omitempty"`
	SMTPPort                     string `json:"smtp_port,omitempty"`
	SMTPEncryption               string `json:"smtp_encryption,omitempty"`
	BodyBackground               string `json:"bodybackground,omitempty"`
	LandingType                  string `json:"landing_type,omitempty"`
	LandingPageBodyBackground    string `json:"landingpage_bodybackground,omitempty"`
	LandingPageBootstrap         string `json:"landingpage_bootstrap,omitempty"`
	TrainingOption               string `json:"training_option,omitempty"`
	TrainingID                   string `json:"training_id,omitempty"`
	CustomAttachmentID           string `json:"custom_attachment_id,omitempty"`
	CustomAttachmentFilename     string `json:"custom_attachment_filename,omitempty"`
	CustomAttachmentOriginalName string `json:"custom_attachment_original_name,omitempty"`
	OptionalField1               string `json:"optional_field_1,omitempty"`
	OptionalField2               string `json:"optional_field_2,omitempty"`
	OptionalField3               string `json:"optional_field_3,omitempty"`
	Community                    string `json:"community,omitempty"`
	SubmittedTemplateID          string `json:"submitted_template_id,omitempty"`
	SubmitStatus                 string `json:"submit_status,omitempty"`
	SubmitterID                  string `json:"submitter_id,omitempty"`
	LandingPageImage             string `json:"landingpage_image,omitempty"`
	EmailImage                   string `json:"email_image,omitempty"`
	DatetimeImaged               string `json:"datetime_imaged,omitempty"`
	DestinationEditor            string `json:"destination_editor,omitempty"`
	DestinationHead              string `json:"destination_head,omitempty"`
	DestinationBodyClass         string `json:"destination_body_class,omitempty"`
	TrackAttachment              string `json:"track_attachment,omitempty"`
	TemplateCategory             string `json:"template_category,omitempty"`
	TrackReply                   string `json:"track_reply,omitempty"`
	EmailEditor                  string `json:"email_editor,omitempty"`
	PreviewToken                 string `json:"preview_token,omitempty"`
	IMAPServer                   string `json:"imap_server,omitempty"`
	IMAPPort                     string `json:"imap_port,omitempty"`
	IMAPUsername                 string `json:"imap_username,omitempty"`
	IMAPPassword                 string `json:"imap_password,omitempty"`
	IMAPEncryption               string `json:"imap_encryption,omitempty"`
	IMAPFolder                   string `json:"imap_folder,omitempty"`
	LibTemplateID                string `json:"lib_template_id,omitempty"`
	CustomHeaderName             string `json:"custom_header_name,omitempty"`
	CustomHeaderValue            string `json:"custom_header_value,omitempty"`
	LandingID                    string `json:"landing_id,omitempty"`
	ErrorScan                    string `json:"error_scan,omitempty"`
	CustomInstructions           string `json:"custom_instructions,omitempty"`
	MovedAccountID               string `json:"moved_account_id,omitempty"`
	MovedTemplateID              string `json:"moved_template_id,omitempty"`
}

// GetTemplate contains fields for get template api calls
type GetTemplate struct {
	Template
	BaseTemplate
	AccountID string `json:"account_id,omitempty"`
}
