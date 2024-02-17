package listener

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net"
	"net/http"

	conf "github.com/ad/vk-callbackapi-to-telegram/config"
	"github.com/ad/vk-callbackapi-to-telegram/models"
	"github.com/ad/vk-callbackapi-to-telegram/sender"
)

type Listener struct {
	lgr    *slog.Logger
	config *conf.Config
	Server *http.Server
	Sender *sender.Sender
}

type serverContextKey string

const keyServerAddr = "serverAddr"

func InitListener(lgr *slog.Logger, config *conf.Config, s *sender.Sender) (*Listener, error) {
	listener := &Listener{
		lgr:    lgr,
		config: config,
		Sender: s,
	}

	mux := http.NewServeMux()

	mux.HandleFunc("/", listener.handler)

	ctx, cancelCtx := context.WithCancel(context.Background())
	server := &http.Server{
		Addr:    net.JoinHostPort(config.ListenHost, config.ListenPort),
		Handler: mux,
		BaseContext: func(l net.Listener) context.Context {
			ctx = context.WithValue(ctx, serverContextKey(keyServerAddr), l.Addr().String())
			return ctx
		},
	}

	go func(*http.Server) {
		err := server.ListenAndServe()
		if errors.Is(err, http.ErrServerClosed) {
			lgr.Info("server closed")
		} else if err != nil {
			lgr.Error(fmt.Sprintf("error listening for server: %s", err))
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
			l.lgr.Error(fmt.Sprintf("error writing response: %s", err))
		}

		return
	}

	l.lgr.Debug(fmt.Sprintf("%s: %s", result.Type, string(bodyValue)))

	if l.config.VkSecret != "" && result.Secret != l.config.VkSecret {
		l.lgr.Debug(fmt.Sprintf("secret mistmatch %s != %s", l.config.VkSecret, result.Secret))
		l.lgr.Debug(string(bodyValue))

		if _, err := io.WriteString(w, "ok"); err != nil {
			l.lgr.Error(fmt.Sprintf("error writing response: %s", err))
		}

		return
	}

	if result.Type == "confirmation" {
		if _, err := io.WriteString(w, l.config.VkConfirmation); err != nil {
			l.lgr.Error(fmt.Sprintf("error writing response: %s", err))
		}

		return
	}

	if _, err := io.WriteString(w, "ok"); err != nil {
		l.lgr.Error(fmt.Sprintf("error writing response: %s", err))
	}

	_ = l.Sender.ProcessVKMessage(result)
}
