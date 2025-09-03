package Telegram

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	TelegramURL         = "" // Telegram
	TelegramToken       = "" // Telegram
	TelegramChatID      = "" // Telegram
	TelegramSendMessage = "" // Telegram
)

func NewClientTelegram() *ClientTelegram {
	return &ClientTelegram{
		Client: &http.Client{},
	}
}

func (tlc *ClientTelegram) SendTextMessage(text string) error {
	return nil
}

func (tlc *ClientTelegram) SendMessageWithPhoto(chatId, caption, imageUrl string) error {
	body := MessagePhoto{
		ChatID:  chatId,
		Caption: caption,
		Photo:   imageUrl,
	}

	bodyJSON, err := json.Marshal(body)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(bodyJSON)

	return nil
}
