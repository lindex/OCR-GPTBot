package telegram

import (
	"encoding/json"
	"log"

	"gopkg.in/telebot.v3"
)

type Logger interface {
	Println(args ...interface{})
}

type InfoLogger interface {
	Infoln(args ...interface{})
}

func PrivateChatOnly(next telebot.HandlerFunc) telebot.HandlerFunc {
	return func(c telebot.Context) error {
		if c.Chat().Type == telebot.ChatPrivate {
			return next(c)
		}
		return nil
	}
}

func Logging(logger ...Logger) telebot.MiddlewareFunc {
	var l Logger
	if len(logger) > 0 {
		l = logger[0]
	} else {
		l = log.Default()
	}

	return func(next telebot.HandlerFunc) telebot.HandlerFunc {
		return func(c telebot.Context) error {
			data, _ := json.MarshalIndent(c.Update(), "", "  ")
			if newLogger, ok := l.(InfoLogger); ok {
				newLogger.Infoln(string(data))
			} else {
				l.Println(string(data))
			}
			return next(c)
		}
	}
}
