package sender

import (
	"context"
	"encoding/json"
	"fmt"
	"sync"

	conf "github.com/ad/vk-callbackapi-to-telegram/config"
	"github.com/go-telegram/bot"
	bm "github.com/go-telegram/bot/models"
)

type Sender struct {
	sync.RWMutex
	config           *conf.Config
	Bot              *bot.Bot
	Config           *conf.Config
	deferredMessages map[int64]chan DeferredMessage
	lastMessageTimes map[int64]int64
}

func InitSender(config *conf.Config) (*Sender, error) {
	sender := &Sender{
		config:           config,
		deferredMessages: make(map[int64]chan DeferredMessage),
		lastMessageTimes: make(map[int64]int64),
	}

	opts := []bot.Option{
		bot.WithDefaultHandler(handler),
		bot.WithSkipGetMe(),
	}

	if config.Debug {
		opts = append(opts, bot.WithDebug())
	}

	b, newBotError := bot.New(config.TelegramToken, opts...)
	if newBotError != nil {
		return nil, fmt.Errorf("start bot error: %s", newBotError)
	}

	sender.Bot = b

	go sender.sendDeferredMessages()

	return sender, nil
}

func handler(ctx context.Context, b *bot.Bot, update *bm.Update) {
	cbdToSend, err := json.Marshal(update)
	if err != nil {
		fmt.Printf("%#v, err %s", update, err)

		return
	}

	fmt.Printf("получено сообщение: %s", string(cbdToSend))
}
