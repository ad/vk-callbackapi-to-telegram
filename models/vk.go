package models

type VkCallbackRequest struct {
	Type   string `json:"type"`
	Object struct {
		ID           int    `json:"id"`
		FromID       int    `json:"from_id,omitempty"`
		Date         int    `json:"date,omitempty"`
		Text         string `json:"text,omitempty"`
		PostOwnerID  int    `json:"post_owner_id,omitempty"`
		PostID       int    `json:"post_id,omitempty"`
		PhotoOwnerID int    `json:"photo_owner_id,omitempty"`
		PhotoID      int    `json:"photo_id,omitempty"`
		TopicID      int    `json:"topic_id,omitempty"`
		Body         string `json:"body,omitempty"`
	} `json:"object,omitempty"`
	GroupID int    `json:"group_id"`
	Secret  string `json:"secret,omitempty"`
}
