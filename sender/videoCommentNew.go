package sender

import (
	"fmt"

	"github.com/ad/vk-callbackapi-to-telegram/models"
)

func videoCommentNew(m *models.VkCallbackRequest) string {
	return fmt.Sprintf(
		"<a href=\"https://vk.com/nethouse?z=video-%d_%d\">Комментарий под видео</a> в группе от vk.com/id%d: %s",
		m.GroupID,
		m.Object.VideoID,
		m.Object.FromID,
		m.Object.Text,
	)
}
