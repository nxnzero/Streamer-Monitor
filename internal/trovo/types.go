package trovo

import "net/http"

type ClientTrovo struct {
	Client *http.Client
}

type RefreshTokenRequest struct {
	ClientSecret string `json:"client_secret"`
	GrantType    string `json:"grant_type"`
	RefreshToken string `json:"refresh_token"`
}

type ResponseTokenRequest struct {
	AccessToken  string `json:"access_token"`
	TokenType    string `json:"token_type"`
	ExpiresIn    int    `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
}

type ChannelInfo struct {
	IsLive         bool         `json:"is_live"`
	CategoryID     string       `json:"category_id"`
	CategoryName   string       `json:"category_name"`
	LiveTitle      string       `json:"live_title"`
	AudiType       string       `json:"audi_type"`
	LanguageCode   string       `json:"language_code"`
	Thumbnail      string       `json:"thumbnail"`
	CurrentViewers int          `json:"current_viewers"`
	Followers      int          `json:"followers"`
	StreamerInfo   string       `json:"streamer_info"`
	ProfilePic     string       `json:"profile_pic"`
	ChannelURL     string       `json:"channel_url"`
	CreatedAt      string       `json:"created_at"`
	SubscriberNum  int          `json:"subscriber_num"`
	Username       string       `json:"username"`
	SocialLinks    []SocialLink `json:"social_links"`
	StartedAt      string       `json:"started_at"`
	EndedAt        string       `json:"ended_at"`
}

type SocialLink struct {
	Type string `json:"type"`
	URL  string `json:"url"`
}

type ChannelInfoByID struct {
	ChannelID string `json:"channel_id"`
}

type ChannelInfoByUsername struct {
	Username string `json:"username"`
}
