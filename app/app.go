package app

import (
	"context"
	"io"
	"os"

	conf "github.com/ad/vk-callbackapi-to-telegram/config"
	lstnr "github.com/ad/vk-callbackapi-to-telegram/listener"
	"github.com/ad/vk-callbackapi-to-telegram/logger"
	sndr "github.com/ad/vk-callbackapi-to-telegram/sender"
)

var (
	config *conf.Config
)

func Run(ctx context.Context, w io.Writer, args []string) error {
	confLoad, errInitConfig := conf.InitConfig(os.Args)
	if errInitConfig != nil {
		return errInitConfig
	}

	config = confLoad

	lgr := logger.InitLogger(config.Debug)

	sender, errInitSender := sndr.InitSender(lgr, config)
	if errInitSender != nil {
		return errInitSender
	}

	if config.TelegramAdminID != 0 {
		sender.MakeRequestDeferred(sndr.DeferredMessage{
			Method: "sendMessage",
			ChatID: config.TelegramAdminID,
			Text:   "Bot restarted",
		}, sender.SendResult)
	}

	_, errInitListener := lstnr.InitListener(lgr, config, sender)
	if errInitListener != nil {
		return errInitListener
	}

	return nil
}
