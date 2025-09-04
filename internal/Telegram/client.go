package Telegram

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const (
	TelegramURL         = "https://api.telegram.org/bot"
	TelegramToken       = "8279530994:AAEbSWILEqLZridXVIWBX89r0umgfB8Q7KY" // Telegram
	TelegramChatID      = "-1002551938305"
	TelegramSendMessage = "https://api.telegram.org/bot8279530994:AAEbSWILEqLZridXVIWBX89r0umgfB8Q7KY/sendMessage" // Telegram
	TelegramSendPhoto   = "https://api.telegram.org/bot8279530994:AAEbSWILEqLZridXVIWBX89r0umgfB8Q7KY/sendPhoto"   // Telegram
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

	request, err := http.NewRequest("POST", TelegramSendPhoto, bytes.NewBuffer(bodyJSON))
	if err != nil {
		fmt.Println(err)
	}

	request.Header.Add("Content-Type", "application/json")

	response, err := tlc.Client.Do(request)
	if err != nil {
		fmt.Println(err)
	}

	dataResponse, _ := io.ReadAll(response.Body)
	fmt.Println(string(dataResponse))

	defer response.Body.Close()

	return nil
}

func (tlc *ClientTelegram) SendMessageWithError(chatId string, err error) error {
	body := MessagePhoto{
		ChatID: chatId,
	}

	bodyJSON, err := json.Marshal(body)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(bodyJSON)

	request, err := http.NewRequest("POST", TelegramSendPhoto, bytes.NewBuffer(bodyJSON))
	if err != nil {
		fmt.Println(err)
	}

	response, err := tlc.Client.Do(request)
	if err != nil {
		fmt.Println(err)
	}

	dataResponse, _ := io.ReadAll(response.Body)
	fmt.Println(string(dataResponse))

	defer response.Body.Close()

	return nil
}
