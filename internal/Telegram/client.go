package Telegram

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	TelegramURL         = "https://api.telegram.org/bot"
	TelegramToken       = "TELEGRAM_TOKEN" // Telegram
	TelegramChatID      = "CHATID"
	TelegramSendMessage = "https://api.telegram.org/bot<TOKEN>/sendMessage" // Telegram
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
