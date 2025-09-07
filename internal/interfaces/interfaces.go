package interfaces

import (
	"LycorisMonitor/internal/trovo"
)

// TrovoClient интерфейс для работы с Trovo API
type TrovoClient interface {
	ChannelByUsername(username string) (*trovo.ChannelInfo, error)
	ChannelByID(channelID string) (*trovo.ChannelInfo, error)
	RefreshAccessToken(accessToken, refreshToken string) (string, string, error)
	ConfigureAuthorizationURL(ClientID, ResponseType string, Scope []string, RedirectURL string) (string, error)
}

// TelegramClient интерфейс для работы с Telegram API
type TelegramClient interface {
	SendMessageWithPhoto(chatId, caption, imageUrl string) error
	SendTextMessage(text string) error
	SendMessageWithError(chatId string, err error) error
}
