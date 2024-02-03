package sender

import (
	"fmt"

	"github.com/ad/vk-callbackapi-to-telegram/models"
)

func photoCommentNew(m *models.VkCallbackRequest) string {
	return fmt.Sprintf(
		"<a href=\"https://vk.com/nethouse?z=photo-%d_%d\">Комментарий под фотографией</a> в группе от vk.com/id%d: %s",
		m.GroupID,
		m.Object.PhotoID,
		m.Object.FromID,
		m.Object.Text,
	)
}
