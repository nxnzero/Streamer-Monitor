package main

import (
	"LycorisMonitor/internal/services"
	streamers "LycorisMonitor/internal/streamers"
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


// Никнеймы стримеров через запятую
func main() {
	channels := []string{
		"Nickname1",
		"Nickname2",
	}

	trc := trovo.NewTrovoClient()       // Инициализируем клиент для Trovo
	tlc := Telegram.NewClientTelegram() // Инициализируем клиент для Telegram

	allStreamers, err := streamers.ReadFromFile("./streamers.json")
	if err != nil {
		fmt.Printf("Error reading streamers file: %v\n", err)
		return
	}

	// Создаем мапу для быстрого поиска стримеров по имени
	streamerMap := make(map[string]*services.Streamer)
	for i := range *allStreamers {
		streamerMap[(*allStreamers)[i].Username] = &(*allStreamers)[i]
	}

	for _, channel := range channels {
		streamerInfo, err := trc.ChannelByUsername(channel)
		if err != nil {
			fmt.Printf("Error getting channel info for %s: %v\n", channel, err)
			continue
		}

		// Находим соответствующего стримера в нашем списке
		streamerFL, exists := streamerMap[channel]
		if !exists {
			fmt.Printf("Streamer %s not found in local data\n", channel)
			continue
		}

		// Проверяем изменение статуса
		if streamerInfo.IsLive && !streamerFL.IsLive {
			// Стример стал онлайн
			message := fmt.Sprintf("%s на связи!\nСсылка на стрим: %s", channel, streamerInfo.ChannelURL)
			err := tlc.SendMessageWithPhoto(Telegram.TelegramChatID, message, "https://your_photo_hosting/path_to_file")
			if err != nil {
				fmt.Printf("Error sending message: %v\n", err)
			} else {
				fmt.Printf("Notification sent: %s is now online\n", streamerInfo.Username)
			}

			streamerFL.IsLive = true
			err = streamers.WriteToFile(allStreamers, "./streamers.json")
			if err != nil {
				fmt.Printf("Error saving streamer status: %v\n", err)
			}
		} else if !streamerInfo.IsLive && streamerFL.IsLive {
			// Стример стал офлайн
			streamerFL.IsLive = false
			err = streamers.WriteToFile(allStreamers, "./streamers.json")
			if err != nil {
				fmt.Printf("Error saving streamer status: %v\n", err)
			} else {
				fmt.Printf("%s is now offline\n", streamerInfo.Username)
			}
		}
	}

}
