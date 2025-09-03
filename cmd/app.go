package main

import (
	"LycorisMonitor/internal/trovo"
	"fmt"
	"os"
)

const (
	Nates13 = "Nates13"
)

var (
	ClientID     = os.Getenv("TROVO_CLIENT_ID")
	AccessToken  = os.Getenv("TROVO_ACCESS_TOKEN")
	RefreshToken = os.Getenv("TROVO_REFRESH_TOKEN")
)

func main() {
	trc := trovo.NewTrovoClient() // Инициализируем клиент для Трово

	UserInfo, err := trc.ChannelByUsername("Nates13") // Запрашиваем данные по стримеру
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(UserInfo) // Отображаем данные

}
