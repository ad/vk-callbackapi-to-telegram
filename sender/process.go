package sender

import (
	"fmt"

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
		if s.config.Debug {
			fmt.Printf("Type %s is not allowed\n", m.Type)
		}

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

func checkIfTypeIsAllowed(m *models.VkCallbackRequest) bool {
	return m.Type == "message_new" || m.Type == "wall_reply_new" || m.Type == "photo_comment_new" || m.Type == "video_comment_new"
}
