package services

import (
	"time"
)

type Streamer struct {
	Username  string    `json:"username"`
	IsLive    bool      `json:"isLive"`
	LastCheck time.Time `json:"lastCheck"`
}
