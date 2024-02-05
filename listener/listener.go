package listener

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net"
	"net/http"

	conf "github.com/ad/vk-callbackapi-to-telegram/config"
	"github.com/ad/vk-callbackapi-to-telegram/models"
	"github.com/ad/vk-callbackapi-to-telegram/sender"
)

type Listener struct {
	config *conf.Config
	Server *http.Server
	Sender *sender.Sender
}

type serverContextKey string

const keyServerAddr = "serverAddr"

func InitListener(config *conf.Config, s *sender.Sender) (*Listener, error) {
	listener := &Listener{
		config: config,
		Sender: s,
	}

	mux := http.NewServeMux()

	mux.HandleFunc("/", listener.handler)

	ctx, cancelCtx := context.WithCancel(context.Background())
	server := &http.Server{
		Addr:    ":3333",
		Handler: mux,
		BaseContext: func(l net.Listener) context.Context {
			ctx = context.WithValue(ctx, serverContextKey(keyServerAddr), l.Addr().String())
			return ctx
		},
	}

	go func(*http.Server) {
		err := server.ListenAndServe()
		if errors.Is(err, http.ErrServerClosed) {
			fmt.Printf("server closed\n")
		} else if err != nil {
			fmt.Printf("error listening for server: %s\n", err)
		}
		cancelCtx()
	}(server)

	listener.Server = server

	return listener, nil
}

func (l *Listener) handler(w http.ResponseWriter, r *http.Request) {
	bodyValue, _ := io.ReadAll(r.Body)
	r.Body.Close()

	result := &models.VkCallbackRequest{}
	errUnmarshal := json.Unmarshal(bodyValue, result)

	if errUnmarshal != nil {
		if _, err := io.WriteString(w, "ok"); err != nil {
			fmt.Printf("error writing response: %s\n", err)
		}

		return
	}

	if l.config.Debug {
		fmt.Printf("%s: %s\n", result.Type, string(bodyValue))
	}

	if l.config.VkSecret != "" && result.Secret != l.config.VkSecret {
		fmt.Printf("secret mistmatch %s != %s \n", l.config.VkSecret, result.Secret)
		if l.config.Debug {
			fmt.Printf("%s\n", string(bodyValue))
		}

		if _, err := io.WriteString(w, "ok"); err != nil {
			fmt.Printf("error writing response: %s\n", err)
		}

		return
	}

	if result.Type == "confirmation" {
		if _, err := io.WriteString(w, l.config.VkConfirmation); err != nil {
			fmt.Printf("error writing response: %s\n", err)
		}

		return
	}

	if _, err := io.WriteString(w, "ok"); err != nil {
		fmt.Printf("error writing response: %s\n", err)
	}

	l.Sender.ProcessVKMessage(result)
}
