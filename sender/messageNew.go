package sender

import (
	"fmt"

	"github.com/ad/vk-callbackapi-to-telegram/models"
)

func messageNew(m *models.VkCallbackRequest) string {
	return fmt.Sprintf(
		"<a href=\"https://vk.com/gim%d\">Сообщение(#%d)</a> в группе от vk.com/id%d: %s",
		m.Object.Message.ConversationMessageID,
		m.GroupID,
		m.Object.Message.FromID,
		m.Object.Message.Text,
	)
}
