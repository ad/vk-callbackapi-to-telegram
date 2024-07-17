package sender

import (
	"context"
	"fmt"
	"log/slog"
	"strconv"
	"sync"

	conf "github.com/ad/vk-callbackapi-to-telegram/config"
	"github.com/go-telegram/bot"
	bm "github.com/go-telegram/bot/models"
)

type Sender struct {
	sync.RWMutex
	lgr              *slog.Logger
	config           *conf.Config
	Bot              *bot.Bot
	Config           *conf.Config
	deferredMessages map[int64]chan DeferredMessage
	lastMessageTimes map[int64]int64
}

func InitSender(lgr *slog.Logger, config *conf.Config) (*Sender, error) {
	sender := &Sender{
		lgr:              lgr,
		config:           config,
		deferredMessages: make(map[int64]chan DeferredMessage),
		lastMessageTimes: make(map[int64]int64),
	}

	opts := []bot.Option{
		bot.WithDefaultHandler(sender.handler),
		bot.WithSkipGetMe(),
	}

	// if config.Debug {
	// 	opts = append(opts, bot.WithDebug())
	// }

	b, newBotError := bot.New(config.TelegramToken, opts...)
	if newBotError != nil {
		return nil, fmt.Errorf("start bot error: %s", newBotError)
	}

	go b.Start(context.Background())
	go sender.sendDeferredMessages()

	sender.Bot = b

	return sender, nil
}

func (s *Sender) handler(ctx context.Context, b *bot.Bot, update *bm.Update) {
	if s.config.Debug {
		if update.Message != nil && update.Message.From != nil && update.Message.Chat.ID != 0 && update.Message.Text != "" {
			s.lgr.Debug(fmt.Sprintf("%s -> %d: %s", getMessageFromUsername(update), update.Message.Chat.ID, update.Message.Text))
		}
	}
}

func getMessageFromUsername(update *bm.Update) string {
	if update.Message != nil && update.Message.From != nil && update.Message.From.Username != "" {
		return update.Message.From.Username
	}

	return strconv.FormatInt(update.Message.From.ID, 10)
}
