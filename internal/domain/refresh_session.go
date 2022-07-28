package domain

import "time"

type RefreshSession struct {
	Token     string    `json:"token" db:"token"`
	ExpiresAt time.Time `json:"expiresAt" db:"expires_at"`
	UserId    int       `json:"-" db:"user_id"`
}