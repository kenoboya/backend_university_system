package model

import "time"

type Session struct {
	RefreshToken string
	ExpiresAt    time.Time
}
