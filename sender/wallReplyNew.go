package sender

import (
	"fmt"

	"github.com/ad/vk-callbackapi-to-telegram/models"
)

func wallReplyNew(m *models.VkCallbackRequest) string {
	return fmt.Sprintf(
		"<a href=\"https://vk.com/wall-%d_%d\">Комментарий на стене</a> в группе от vk.com/id%d: %s",
		m.GroupID,
		m.Object.PostID,
		m.Object.FromID,
		m.Object.Text,
	)
}
