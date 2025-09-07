package Telegram

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	TelegramURL         = "https://api.telegram.org/bot"
	TelegramToken       = "8279530994:AAEbSWILEqLZridXVIWBX89r0umgfB8Q7KY" // Telegram
	TelegramChatID      = "-1002551938305"
	TelegramSendMessage = "https://api.telegram.org/bot8279530994:AAEbSWILEqLZridXVIWBX89r0umgfB8Q7KY/sendMessage" // Telegram
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
