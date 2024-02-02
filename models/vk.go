package models

type VkCallbackRequest struct {
	Type   string `json:"type"`
	Object struct {
		ID           int     `json:"id"`
		AdminID      int     `json:"admin_id,omitempty"`
		UserID       int     `json:"user_id,omitempty"`
		FromID       int     `json:"from_id,omitempty"`
		AlbumID      int     `json:"album_id,omitempty"`
		OwnerID      int     `json:"owner_id,omitempty"`
		DeleterID    int     `json:"deleter_id,omitempty"`
		ToID         int     `json:"to_id,omitempty"`
		Date         int     `json:"date,omitempty"`
		UnblockDate  int     `json:"unblock_date,omitempty"`
		Reason       int     `json:"reason,omitempty"`
		Comment      string  `json:"comment,omitempty"`
		Key          string  `json:"key,omitempty"`
		Text         string  `json:"text,omitempty"`
		PostponedID  string  `json:"postponed_id,omitempty"`
		PostOwnerID  int     `json:"post_owner_id,omitempty"`
		PostID       int     `json:"post_id,omitempty"`
		PhotoOwnerID int     `json:"photo_owner_id,omitempty"`
		PhotoID      int     `json:"photo_id,omitempty"`
		TopicID      int     `json:"topic_id,omitempty"`
		TopicOwnerID int     `json:"topic_owner_id,omitempty"`
		Self         int     `json:"self,omitempty"`
		Body         string  `json:"body,omitempty"`
		JoinType     string  `json:"join_type,omitempty"`
		Message      Message `json:"message,omitempty"`
	} `json:"object,omitempty"`
	GroupID int    `json:"group_id"`
	Secret  string `json:"secret,omitempty"`
}

type Message struct {
	ID                    int    `json:"id"`
	ConversationMessageID int    `json:"conversation_message_id,omitempty"`
	Date                  int    `json:"date,omitempty"`
	UserID                int    `json:"user_id,omitempty"`
	FromID                int    `json:"from_id,omitempty"`
	Title                 string `json:"title,omitempty"`
	Text                  string `json:"text,omitempty"`
	Body                  string `json:"body,omitempty"`
	Action                Action `json:"action,omitempty"`
	AdminAuthorID         int    `json:"admin_author_id,omitempty"`
	IsCropped             bool   `json:"is_cropped,omitempty"`
	MembersCount          int    `json:"members_count,omitempty"`
	PinnedAt              int    `json:"pinned_at,omitempty"`
	Out                   int    `json:"out,omitempty"`
	Deleted               int    `json:"deleted,omitempty"`
}

type Action struct {
	Type     string `json:"type,omitempty"`
	MemberID int    `json:"member_id,omitempty"`
	Text     string `json:"text,omitempty"`
	Email    string `json:"email,omitempty"`
}
