package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	conf "github.com/ad/vk-callbackapi-to-telegram/config"
	lstnr "github.com/ad/vk-callbackapi-to-telegram/listener"
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

	_, errInitListener := lstnr.InitListener(config)
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

	if config.Debug {
		fmt.Println("awaiting signal")
	}

	<-done
	fmt.Println("exiting")
}
