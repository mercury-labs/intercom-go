package intercom

import "fmt"

// ContactService handles interactions with the API through a ContactRepository.
type ContactService struct {
	Repository ContactRepository
}

// ContactList holds a list of Contacts and paging information
type ContactList struct {
	Pages       PageParams
	Contacts    []Contact `json:"data"`
	ScrollParam string    `json:"scroll_param,omitempty"`
}

// Contact represents a Contact within Intercom.
// Not all of the fields are writeable to the API, non-writeable fields are
// stripped out from the request. Please see the API documentation for details.
type Contact struct {
	Type                   string                 `json:"type,omitempty"`
	ID                     string                 `json:"id,omitempty"`
	WorkspaceID            string                 `json:"workspace_id,omitempty"`
	ExternalID             string                 `json:"external_id,omitempty"`
	role                   string                 `json:"role,omitempty"`
	Email                  string                 `json:"email,omitempty"`
	Phone                  string                 `json:"phone,omitempty"`
	Name                   string                 `json:"name,omitempty"`
	Avatar                 string                 `json:"avatar,omitempty"`
	OwnerID                int64                  `json:"owner_id,omitempty"`
	SocialProfiles         *SocialProfileList     `json:"social_profiles,omitempty"`
	HasHardBounced         *bool                  `json:"has_hard_bounced,omitempty"`
	MarkedEmailAsSpam      *bool                  `json:"marked_email_as_spam,omitempty"`
	UnsubscribedFromEmails *bool                  `json:"unsubscribed_from_emails,omitempty"`
	CreatedAt              int64                  `json:"created_at,omitempty"`
	UpdatedAt              int64                  `json:"updated_at,omitempty"`
	SignedUpAt             int64                  `json:"signed_up_at,omitempty"`
	LastSeenAt             int64                  `json:"last_seen_at,omitempty"`
	LastRepliedAt          int64                  `json:"last_replied_at,omitempty"`
	LastContactedAt        int64                  `json:"last_contacted_at,omitempty"`
	LastEmailOpenedAt      int64                  `json:"last_email_opened_at,omitempty"`
	LastEmailClickedAt     int64                  `json:"last_email_clicked_at,omitempty"`
	LanguageOverride       string                 `json:"language_override,omitempty"`
	Browser                string                 `json:"browser,omitempty"`
	BrowserVersion         string                 `json:"browser_version,omitempty"`
	BrowserLanguage        string                 `json:"browser_language,omitempty"`
	OS                     string                 `json:"os,omitempty"`
	Location               *ContactLocation       `json:"location,omitempty"`
	AndroidAppName         string                 `json:"android_app_name,omitempty"`
	AndroidAppVersion      string                 `json:"android_app_version,omitempty"`
	AndroidDevice          string                 `json:"android_device,omitempty"`
	AndroidOSVersion       string                 `json:"android_os_version,omitempty"`
	AndroidSDKVersion      string                 `json:"android_sdk_version,omitempty"`
	AndroidLastSeenAt      int64                  `json:"android_last_seen_at,omitempty"`
	IOSAppName             string                 `json:"ios_app_name,omitempty"`
	IOSAppVersion          string                 `json:"ios_app_version,omitempty"`
	IOSDevice              string                 `json:"ios_device,omitempty"`
	IOSOSVersion           string                 `json:"ios_os_version,omitempty"`
	IOSSDKVersion          string                 `json:"ios_sdk_version,omitempty"`
	IOSLastSeenAt          int64                  `json:"ios_last_seen_at,omitempty"`
	CustomAttributes       map[string]interface{} `json:"custom_attributes,omitempty"`
	Tags                   *AddressableList       `json:"tags,omitempty"`
	Notes                  *AddressableList       `json:"notes,omitempty"`
	Companies              *AddressableList       `json:"companies,omitempty"`
}

type contactListParams struct {
	PageParams
	SegmentID string `url:"segment_id,omitempty"`
	TagID     string `url:"tag_id,omitempty"`
	Email     string `url:"email,omitempty"`
}

type ContactLocation struct {
	Type    string `json:"type,omitempty"`
	Country string `json:"country,omitempty"`
	Region  string `json:"region,omitempty"`
	City    string `json:"city,omitempty"`
}

type Addressable struct {
	Type string `json:"type,omitempty"`
	ID   string `json:"id,omitempty"`
	URL  string `json:"url,omitempty"`
}

type AddressableList struct {
	Type       string         `json:"type,omitempty"`
	Data       *[]Addressable `json:"data,omitempty"`
	URL        string         `json:"url,omitempty"`
	TotalCount int64          `json:"total_count,omitempty"`
	HasMore    *bool          `json:"has_more,omitempty"`
}

// FindByID looks up a Contact by their Intercom ID.
func (c *ContactService) FindByID(id string) (Contact, error) {
	return c.findWithIdentifiers(UserIdentifiers{ID: id})
}

// FindByUserID looks up a Contact by their UserID (automatically generated server side).
func (c *ContactService) FindByUserID(userID string) (Contact, error) {
	return c.findWithIdentifiers(UserIdentifiers{UserID: userID})
}

func (c *ContactService) findWithIdentifiers(identifiers UserIdentifiers) (Contact, error) {
	return c.Repository.find(identifiers)
}

// List all Contacts for App.
func (c *ContactService) List(params PageParams) (ContactList, error) {
	return c.Repository.list(contactListParams{PageParams: params})
}

// List all Contacts for App via Scroll API
func (c *ContactService) Scroll(scrollParam string) (ContactList, error) {
	return c.Repository.scroll(scrollParam)
}

// ListByEmail looks up a list of Contacts by their Email.
func (c *ContactService) ListByEmail(email string, params PageParams) (ContactList, error) {
	return c.Repository.list(contactListParams{PageParams: params, Email: email})
}

// List Contacts by Segment.
func (c *ContactService) ListBySegment(segmentID string, params PageParams) (ContactList, error) {
	return c.Repository.list(contactListParams{PageParams: params, SegmentID: segmentID})
}

// List Contacts By Tag.
func (c *ContactService) ListByTag(tagID string, params PageParams) (ContactList, error) {
	return c.Repository.list(contactListParams{PageParams: params, TagID: tagID})
}

// Create Contact
func (c *ContactService) Create(contact *Contact) (Contact, error) {
	return c.Repository.create(contact)
}

// Update Contact
func (c *ContactService) Update(contact *Contact) (Contact, error) {
	return c.Repository.update(contact)
}

// Convert Contact to User
func (c *ContactService) Convert(contact *Contact, user *User) (User, error) {
	return c.Repository.convert(contact, user)
}

// Delete Contact
func (c *ContactService) Delete(contact *Contact) (Contact, error) {
	return c.Repository.delete(contact.ID)
}

// MessageAddress gets the address for a Contact in order to message them
func (c Contact) MessageAddress() MessageAddress {
	return MessageAddress{
		Type:  "contact",
		ID:    c.ID,
		Email: c.Email,
		// UserID: c.UserID,
	}
}

func (c Contact) String() string {
	return fmt.Sprintf("[intercom] contact { id: %s name: %s, email: %s ... }", c.ID, c.Name, c.Email)
}
