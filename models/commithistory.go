package models

import (
	"time"

	"gorm.io/gorm"
)

type CommitHistory struct {
	gorm.Model  `json:"model"`
	ID          uint64    `json:"id"`
	SHA         string    `json:"sha"`
	URL         string    `json:"url"`
	Date        time.Time `json:"date"`
	Message     string    `json:"message"`
	AuthorName  string    `json:"author_name"`
	AuthorEmail string    `json:"author_email"`
	// Files       map[string]interface `json:"files"`
	CreatedAt time.Time `json:"created_at"`
	TrackID   uint64    `json:"track_id"`
	Track     Track
}
