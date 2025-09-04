package main

import (
	Telegram "LycorisMonitor/internal/telegram"
	"LycorisMonitor/internal/trovo"
	"fmt"
	"os"
)

var (
	ClientID     = os.Getenv("TROVO_CLIENT_ID")
	AccessToken  = os.Getenv("TROVO_ACCESS_TOKEN")
	RefreshToken = os.Getenv("TROVO_REFRESH_TOKEN")
)

func main() {
	channels := []string{
		"Nates13",
		"Мирилит",
		"Леший",
		"illusiveHope",
	}

	trc := trovo.NewTrovoClient()       // Инициализируем клиент для Trovo
	tlc := Telegram.NewClientTelegram() // Инициализируем клиент для Telegram

	// Добавить функцию проверки статуса

	for _, channel := range channels {
		UserInfo, err := trc.ChannelByUsername(channel) // Запрашиваем данные по стримеру
		if err != nil {
			fmt.Println(err)
		}

		if UserInfo.IsLive { // запрашиваем статус
			textMessage := fmt.Sprintf("@%v на связи!\nПриходи посмотреть: %v", UserInfo.Username, UserInfo.ChannelURL)
			tlc.SendMessageWithPhoto("-1002551938305", textMessage, "disk.yandex.ru/i/XccYcEQK7TOq_g")
		} else {

		}
	}
}
