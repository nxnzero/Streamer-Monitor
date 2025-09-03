package Telegram

import "net/http"

type ClientTelegram struct {
	Client *http.Client
}

type MessagePhoto struct {
	ChatID  string `json:"chat_id"`
	Photo   string `json:"photo"`
	Caption string `json:"caption"`
}
