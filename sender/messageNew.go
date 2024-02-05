package sender

import (
	"fmt"

	"github.com/ad/vk-callbackapi-to-telegram/models"
)

func messageNew(m *models.VkCallbackRequest) string {
	return fmt.Sprintf(
		"<a href=\"https://vk.com/gim%d?sel=%d&msgid=%d\">Сообщение(#%d)</a> в <a href=\"https://vk.com/gim%d\">группе</a> от vk.com/id%d: %s",
		m.GroupID,
		m.Object.Message.FromID,
		m.Object.Message.ID,
		m.Object.Message.ConversationMessageID,
		m.GroupID,
		m.Object.Message.FromID,
		m.Object.Message.Text,
	)
}
