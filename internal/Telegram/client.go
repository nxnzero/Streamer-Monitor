package Telegram

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net"
	"net/http"
	"time"
)

const (
	TelegramURL         = "https://api.telegram.org/bot"
	TelegramToken       = "8279530994:AAEbSWILEqLZridXVIWBX89r0umgfB8Q7KY" // Telegram
	TelegramChatID      = "-1002551938305"
	TelegramSendMessage = "https://api.telegram.org/bot8279530994:AAEbSWILEqLZridXVIWBX89r0umgfB8Q7KY/sendMessage" // Telegram
	TelegramSendPhoto   = "https://api.telegram.org/bot8279530994:AAEbSWILEqLZridXVIWBX89r0umgfB8Q7KY/sendPhoto"   // Telegram
)

func NewClientTelegram() *ClientTelegram {
	transport := &http.Transport{
		MaxIdleConns:        100,
		MaxIdleConnsPerHost: 10,
		IdleConnTimeout:     90 * time.Second,
		DisableKeepAlives:   false,
		DialContext: (&net.Dialer{
			Timeout:   5 * time.Second,
			KeepAlive: 30 * time.Second,
		}).DialContext,
		TLSHandshakeTimeout:   10 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
	}

	return &ClientTelegram{
		Client: &http.Client{
			Timeout:   10 * time.Second,
			Transport: transport,
		},
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
		return fmt.Errorf("failed to marshal request body: %w", err)
	}

	request, err := http.NewRequest("POST", TelegramSendPhoto, bytes.NewBuffer(bodyJSON))
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}

	request.Header.Add("Content-Type", "application/json")

	response, err := tlc.Client.Do(request)
	if err != nil {
		return fmt.Errorf("failed to make request: %w", err)
	}
	defer response.Body.Close()

	if response.StatusCode < 200 || response.StatusCode >= 300 {
		dataResponse, _ := io.ReadAll(response.Body)
		return fmt.Errorf("telegram API error: status %d, body: %s", response.StatusCode, string(dataResponse))
	}

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
