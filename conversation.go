package intercom

// ConversationService handles interactions with the API through an ConversationRepository.
type ConversationService struct {
	Repository ConversationRepository
}

// ConversationList is a list of Conversations
type ConversationList struct {
	Pages         PageParams     `json:"pages"`
	Conversations []Conversation `json:"conversations"`
}

// A Conversation represents a conversation between users and admins in Intercom.
type Conversation struct {
	Type               string                  `json:"type"`
	ID                 string                  `json:"id"`
	CreatedAt          int64                   `json:"created_at"`
	UpdatedAt          int64                   `json:"updated_at"`
	WaitingSince       int64                   `json:"waiting_since"`
	SnoozedUntil       int64                   `json:"snoozed_until"`
	Source             Source                  `json:source""`
	Contacts           ConversationContactList `json:"contacts"`
	FirstContactReply  FirstContactReply       `json:"first_contact_reply"`
	AdminAssigneeID    int64                   `json:"admin_assignee_id"`
	TeamAssigneeID     string                  `json:"team_assignee_id"`
	Open               bool                    `json:"open"`
	Read               bool                    `json:"read"`
	Tags               ConversationTagList     `json:"tags"`
	Priority           string                  `json:"priority"`
	SLAApplied         SLAApplied              `json:"sla_applied"`
	Statistics         ConversationStatistics  `json:"statistics"`
	ConversationRating ConversationRating      `json:"conversation_rating"`
	Teammates          ConversationTeammate    `json:"teammates"`
	Title              string                  `json:"title"`
	CustomAttributes   map[string]interface{}  `json:"custom_attributes"`
}

// SourceAuthor ...
type SourceAuthor struct {
	Type string `json:"type"`
	ID   string `json:"id"`
}

// Source ...
type Source struct {
	Type        string       `json:"type"`
	ID          string       `json:"id"`
	DeliveredAs string       `json:"delivered_as"`
	Subject     string       `json:"subject"`
	Body        string       `json:"body"`
	Author      SourceAuthor `json:"author"`
	// Attachments map[string]interface{} `json:"attachments"`
	URL string `json:"url"`
}

// ConversationContactObj ...
type ConversationContactObj struct {
	Type string `json:"type"`
	ID   string `json:"id"`
}

// ConversationContactList ...
type ConversationContactList struct {
	Type     string                   `json:"type"`
	Contacts []ConversationContactObj `json:"contacts"`
}

// ConversationTagObj ...
type ConversationTagObj struct {
	Type string `json:"type"`
	ID   string `json:"id"`
	Name string `json:"name"`
}

// ConversationTagList ...
type ConversationTagList struct {
	Type string               `json:"type"`
	Tags []ConversationTagObj `json:"tags"`
}

// FirstContactReply ...
type FirstContactReply struct {
	Type      string `json:"type"`
	URL       string `json:"url"`
	CreatedAt int64  `json:"created_at"`
}

// SLAApplied ...
type SLAApplied struct {
	SLAName   string `json:"sla_name"`
	SLAStatus string `json:"sla_status"`
}

// ConversationStatistics ...
type ConversationStatistics struct {
	TimeToAssignment           int64                  `json:"time_to_assignment"`
	TimeToAdminReply           int64                  `json:"time_to_admin_reply"`
	TimeToFirstClose           int64                  `json:"time_to_first_close"`
	TimeToLastClose            int64                  `json:"time_to_last_close"`
	MedianTimeToReply          int64                  `json:"median_time_to_reply"`
	FirstContactReplyAt        int64                  `json:"first_contact_reply_at"`
	FirstAssignmentAt          int64                  `json:"first_assignment_at"`
	FirstAdminReplyAt          int64                  `json:"first_admin_reply_at"`
	FirstCloseAt               int64                  `json:"first_close_at"`
	LastAssignmentAt           int64                  `json:"last_assignment_at"`
	LastAssignmentAdminReplyAt int64                  `json:"last_assignment_admin_reply_at"`
	LastContactReplyAt         int64                  `json:"last_contact_reply_at"`
	LastAdminReplyAt           int64                  `json:"last_admin_reply_at"`
	LastCloseAt                int64                  `json:"last_close_at"`
	LastClosedBy               map[string]interface{} `json:"last_closed_by"`
	CountReopens               int64                  `json:"count_reopens"`
	CountAssignments           int64                  `json:"count_assignments"`
	CountConversationsParts    int64                  `json:"count_conversations_parts"`
}

// ConversationRating ...
type ConversationRating struct {
	Rating    int32                  `json:"rating"`
	Remark    string                 `json:"remark"`
	CreatedAt int64                  `json:"created_at"`
	Contact   map[string]interface{} `json:"contact"`
	Teammate  map[string]interface{} `json:"teammate"`
}

// ConversationTeammate ...
type ConversationTeammate struct {
	Type  string `json:"type"`
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

// A ConversationMessage is the message that started the conversation rendered for presentation
type ConversationMessage struct {
	ID      string         `json:"id"`
	Subject string         `json:"subject"`
	Body    string         `json:"body"`
	Author  MessageAddress `json:"author"`
	URL     string         `json:"url"`
}

// A ConversationPartList lists the subsequent Conversation Parts
type ConversationPartList struct {
	Parts []ConversationPart `json:"conversation_parts"`
}

// A ConversationPart is a Reply, Note, or Assignment to a Conversation
type ConversationPart struct {
	ID         string         `json:"id"`
	PartType   string         `json:"part_type"`
	Body       string         `json:"body"`
	CreatedAt  int64          `json:"created_at"`
	UpdatedAt  int64          `json:"updated_at"`
	NotifiedAt int64          `json:"notified_at"`
	AssignedTo Admin          `json:"assigned_to"`
	Author     MessageAddress `json:"author"`
}

// The state of Conversations to query
// SHOW_ALL shows all conversations,
// SHOW_OPEN shows only open conversations (only valid for Admin Conversation queries)
// SHOW_CLOSED shows only closed conversations (only valid for Admin Conversation queries)
// SHOW_UNREAD shows only unread conversations (only valid for User Conversation queries)
type ConversationListState int

const (
	SHOW_ALL ConversationListState = iota
	SHOW_OPEN
	SHOW_CLOSED
	SHOW_UNREAD
)

// List all Conversations
func (c *ConversationService) ListAll(pageParams PageParams) (ConversationList, error) {
	return c.Repository.list(ConversationListParams{PageParams: pageParams})
}

// List Conversations by Admin
func (c *ConversationService) ListByAdmin(admin *Admin, state ConversationListState, pageParams PageParams) (ConversationList, error) {
	params := ConversationListParams{
		PageParams: pageParams,
		Type:       "admin",
		AdminID:    admin.ID.String(),
	}
	if state == SHOW_OPEN {
		params.Open = Bool(true)
	}
	if state == SHOW_CLOSED {
		params.Open = Bool(false)
	}
	return c.Repository.list(params)
}

// List Conversations by User
func (c *ConversationService) ListByUser(user *User, state ConversationListState, pageParams PageParams) (ConversationList, error) {
	params := ConversationListParams{
		PageParams:     pageParams,
		Type:           "user",
		IntercomUserID: user.ID,
		UserID:         user.UserID,
		Email:          user.Email,
	}
	if state == SHOW_UNREAD {
		params.Unread = Bool(true)
	}
	return c.Repository.list(params)
}

// Find Conversation by conversation id
func (c *ConversationService) Find(id string) (Conversation, error) {
	return c.Repository.find(id)
}

// Mark Conversation as read (by a User)
func (c *ConversationService) MarkRead(id string) (Conversation, error) {
	return c.Repository.read(id)
}

func (c *ConversationService) Reply(id string, author MessagePerson, replyType ReplyType, body string) (Conversation, error) {
	return c.reply(id, author, replyType, body, nil)
}

// Reply to a Conversation by id
func (c *ConversationService) ReplyWithAttachmentURLs(id string, author MessagePerson, replyType ReplyType, body string, attachmentURLs []string) (Conversation, error) {
	return c.reply(id, author, replyType, body, attachmentURLs)
}

func (c *ConversationService) reply(id string, author MessagePerson, replyType ReplyType, body string, attachmentURLs []string) (Conversation, error) {
	addr := author.MessageAddress()
	reply := Reply{
		Type:           addr.Type,
		ReplyType:      replyType.String(),
		Body:           body,
		AttachmentURLs: attachmentURLs,
	}
	if addr.Type == "admin" {
		reply.AdminID = addr.ID
	} else {
		reply.IntercomID = addr.ID
		reply.UserID = addr.UserID
		reply.Email = addr.Email
	}
	return c.Repository.reply(id, &reply)
}

// Assign a Conversation to an Admin
func (c *ConversationService) Assign(id string, assigner, assignee *Admin) (Conversation, error) {
	assignerAddr := assigner.MessageAddress()
	assigneeAddr := assignee.MessageAddress()
	reply := Reply{
		Type:       "admin",
		ReplyType:  CONVERSATION_ASSIGN.String(),
		AdminID:    assignerAddr.ID,
		AssigneeID: assigneeAddr.ID,
	}
	return c.Repository.reply(id, &reply)
}

// Open a Conversation (without a body)
func (c *ConversationService) Open(id string, opener *Admin) (Conversation, error) {
	return c.reply(id, opener, CONVERSATION_OPEN, "", nil)
}

// Close a Conversation (without a body)
func (c *ConversationService) Close(id string, closer *Admin) (Conversation, error) {
	return c.reply(id, closer, CONVERSATION_CLOSE, "", nil)
}

type ConversationListParams struct {
	PageParams
	Type           string `url:"type,omitempty"`
	AdminID        string `url:"admin_id,omitempty"`
	IntercomUserID string `url:"intercom_user_id,omitempty"`
	UserID         string `url:"user_id,omitempty"`
	Email          string `url:"email,omitempty"`
	Open           *bool  `url:"open,omitempty"`
	Unread         *bool  `url:"unread,omitempty"`
	DisplayAs      string `url:"display_as,omitempty"`
}
