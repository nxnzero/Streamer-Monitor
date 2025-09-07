package services

import (
	"time"
)

type Streamer struct {
	Username     string    `json:"username"`
	IsLive       bool      `json:"isLive"`
	LastCheck    time.Time `json:"lastCheck"`
	GamePreviews []Preview `json:"previews,omitempty"`
}

type Preview struct {
	Game string `json:"game"`
	URL  string `json:"url"`
}
