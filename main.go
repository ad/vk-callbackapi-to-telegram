package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	conf "github.com/ad/vk-callbackapi-to-telegram/config"
	lstnr "github.com/ad/vk-callbackapi-to-telegram/listener"
	sndr "github.com/ad/vk-callbackapi-to-telegram/sender"
)

var (
	config *conf.Config
)

func main() {
	sigs := make(chan os.Signal, 1)

	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	done := make(chan bool, 1)

	confLoad, errInitConfig := conf.InitConfig()
	if errInitConfig != nil {
		log.Fatal(errInitConfig)
	}

	config = confLoad

	sender, errInitSender := sndr.InitSender(config)
	if errInitSender != nil {
		log.Fatal(errInitSender)
	}

	sender.MakeRequestDeferred(sndr.DeferredMessage{
		Method: "sendMessage",
		ChatID: config.TelegramTargetID,
		Text:   "started",
	}, sender.SendResult)

	_, errInitListener := lstnr.InitListener(config, sender)
	if errInitListener != nil {
		log.Fatal(errInitListener)
	}

	// defer listener.Disconnect()

	go func() {
		sig := <-sigs
		fmt.Println()
		fmt.Println(sig)
		done <- true
	}()

	fmt.Println("started")

	<-done
	fmt.Println("exiting")
}
