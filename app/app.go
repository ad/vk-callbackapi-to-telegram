package app

import (
	"context"
	// "fmt"
	"io"
	// "os"
	// "sync"

	conf "github.com/ad/vk-callbackapi-to-telegram/config"
	lstnr "github.com/ad/vk-callbackapi-to-telegram/listener"
	sndr "github.com/ad/vk-callbackapi-to-telegram/sender"
)

var (
	config *conf.Config
)

func Run(ctx context.Context, w io.Writer, args []string) error {
	confLoad, errInitConfig := conf.InitConfig()
	if errInitConfig != nil {
		return errInitConfig
	}

	config = confLoad

	sender, errInitSender := sndr.InitSender(config)
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

	_, errInitListener := lstnr.InitListener(config, sender)
	if errInitListener != nil {
		return errInitListener
	}

	// var wg sync.WaitGroup
	// wg.Add(1)
	// go func() {
	// 	defer wg.Done()
	// 	<-ctx.Done()
	// 	if err := listener.Server.Shutdown(ctx); err != nil {
	// 		fmt.Fprintf(os.Stderr, "error shutting down http server: %s\n", err)
	// 	}
	// }()
	// wg.Wait()

	return nil
}
