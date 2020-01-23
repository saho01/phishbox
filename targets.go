package phishbox

// Target contains common info for all target responses
type Target struct {
	FirstName string `json:"first_name,omitempty"`
	LastName  string `json:"last_name,omitempty"`
	Email     string `json:"email,omitempty"`
	Label     string `json:"label,omitempty"`
}

// Optional contains optional fields
type Optional struct {
	Optional1 string `json:"optional_1,omitempty"`
	Optional2 string `json:"optional_2,omitempty"`
	Optional3 string `json:"optional_3,omitempty"`
}

// OptionalShort contains optional fields in different format
type OptionalShort struct {
	Optional1 string `json:"o1,omitempty"`
	Optional2 string `json:"o2,omitempty"`
	Optional3 string `json:"o3,omitempty"`
}

// TargetIDs contains ID fields for various items
type TargetIDs struct {
	GroupID  string `json:"group_id,omitempty"`
	EmailID  string `json:"email_id,omitempty"`
	PhishKey string `json:"phish_key,omitempty"`
}

// ReportTestTarget contains common fields from report and test api calls
type ReportTestTarget struct {
	TempEmailID      string `json:"temp_email_id,omitempty"`
	SentID           string `json:"sent_id,omitempty"`
	FullName         string `json:"full_name,omitempty"`
	ListType         string `json:"list_type,omitempty"`
	OpenedEmail      string `json:"opened_email,omitempty"`
	OpenedNum        string `json:"opened_num,omitempty"`
	IsSent           string `json:"is_sent,omitempty"`
	LastIPAddress    string `json:"last_ip_address,omitempty"`
	RemoveDuplicates string `json:"remove_duplicates,omitempty"`
	DateSent         string `json:"date_sent,omitempty"`
	TemplateID       string `json:"template_id,omitempty"`
	SendTimestamp    string `json:"send_timestamp,omitempty"`
}

// TestTarget contains fields from Test api call
type TestTarget struct {
	Target
	TargetIDs
	Optional
	ReportTestTarget
	FailedAction  string `json:"failed_action,omitempty"`
	Failed        string `json:"failed,omitempty"`
	FailedDeliv   string `json:"failed_deliv,omitempty"`
	FailedSub     string `json:"failed_sub,omitempty"`
	TestingStatus string `json:"testing_status,omitempty"`
}

// ReportTarget contains fields from Report api call for targets
type ReportTarget struct {
	Target
	TargetIDs
	Optional
	ReportTestTarget
}

// GetTarget contains fields for get target api call
type GetTarget struct {
	Target
	TargetIDs
	Name             string `json:"name,omitempty"`
	MiddleName       string `json:"middle_name,omitempty"`
	Company          string `json:"company,omitempty"`
	Title            string `json:"title,omitempty"`
	AddressOne       string `json:"address_one,omitempty"`
	AddressTwo       string `json:"address_two,omitempty"`
	City             string `json:"city,omitempty"`
	State            string `json:"state,omitempty"`
	Zip              string `json:"zip,omitempty"`
	Country          string `json:"country,omitempty"`
	PhoneBusiness    string `json:"phone_business,omitempty"`
	PhoneBusinessFax string `json:"phone_business_fax,omitempty"`
	PhoneMobile      string `json:"phone_mobile,omitempty"`
	Optional
}

// CreateTarget contains fields for create target api call
type CreateTarget struct {
	Target
	GroupID string `json:"group_id,omitempty"`
	OptionalShort
}

// CreateBatchTarget contains fields for create batch target api call
type CreateBatchTarget struct {
	Target
	OptionalShort
}

// UpdateTarget contains fields for update target api call
type UpdateTarget struct {
	Target
	EmailID int64 `json:"email_id,omitempty"`
	OptionalShort
}

// TargetFailedDetails contains fields for failed targets in get report api call
type TargetFailedDetails struct {
	AccessID      string `json:"access_id,omitempty"`
	AccessDate    string `json:"access_date,omitempty"`
	SentID        string `json:"sent_id,omitempty"`
	GroupID       string `json:"group_id,omitempty"`
	AccessGroupID string `json:"accessed_group_id,omitempty"`
	Name          string `json:"name,omitempty"`
	Email         string `json:"email,omitempty"`
	Type          string `json:"type,omitempty"`
	IPAddress     string `json:"ip_address,omitempty"`
	UserAgent     string `json:"user_agent,omitempty"`
	RequestURI    string `json:"request_uri,omitempty"`
	Domain        string `json:"domain,omitempty"`
	PhishKey      string `json:"phish_key,omitempty"`
	IPAddressNum  string `json:"ip_address_num,omitempty"`
}
