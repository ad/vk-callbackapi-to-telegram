package sender

import (
	"fmt"
	"time"

	"github.com/ad/vk-callbackapi-to-telegram/models"
)

func (s *Sender) PrepareMessage(m *models.VkCallbackRequest) string {
	switch m.Type {
	case "message_new":
		return messageNew(m)
	case "wall_reply_new":
		return wallReplyNew(m)
	case "photo_comment_new":
		return photoCommentNew(m)
	case "video_comment_new":
		return videoCommentNew(m)
	}

	return ""
}
func (s *Sender) ProcessVKMessage(m *models.VkCallbackRequest) error {
	if !checkIfTypeIsAllowed(m) {
		s.lgr.Debug(fmt.Sprintf("Type %s is not allowed", m.Type))
		return nil
	}

	// Check if the message is a duplicate
	if s.isDuplicateMessage(m) {
		// s.lgr.Debug("Duplicate message received")
		return nil
	}

	message := s.PrepareMessage(m)

	if message != "" {
		s.MakeRequestDeferred(DeferredMessage{
			Method: "sendMessageHTML",
			ChatID: s.config.TelegramTargetID,
			Text:   message,
		}, s.SendResult)
	}

	return nil
}

func (s *Sender) isDuplicateMessage(m *models.VkCallbackRequest) bool {
	s.Lock()
	defer s.Unlock()

	// Check if the message already exists in the cache
	if _, ok := s.messageCache[m.Object.ID]; ok {
		return true
	}

	// Add the message to the cache
	s.messageCache[m.Object.ID] = time.Now()

	// Remove old messages from the cache
	for id, timestamp := range s.messageCache {
		if time.Since(timestamp) > 5*time.Second {
			delete(s.messageCache, id)
		}
	}

	return false
}

func checkIfTypeIsAllowed(m *models.VkCallbackRequest) bool {
	return m.Type == "message_new" || m.Type == "wall_reply_new" || m.Type == "photo_comment_new" || m.Type == "video_comment_new"
}
