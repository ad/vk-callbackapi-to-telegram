package sender

import (
	"fmt"
)

func (sender *Sender) SendResult(s SendResult) error {
	if s.Error != nil {
		fmt.Printf("message id %d sent to %d error %q: %s\n", s.MessageID, s.ChatID, s.Error, s.Msg)
		return s.Error
	}

	fmt.Printf("message id %d sent to %d: %s\n", s.MessageID, s.ChatID, s.Msg)

	return nil
}
