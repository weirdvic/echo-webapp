package models

import (
	"errors"
	"time"
)

var ErrNoRecord = errors.New("models: records not found")

type Snippet struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
	ExpiresAt time.Time `json:"expires_at"`
}
