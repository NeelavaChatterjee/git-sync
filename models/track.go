package models

import (
	"time"

	"gorm.io/gorm"
)

type Track struct {
	gorm.Model   `json:"model"`
	ID           uint64    `json:"id"`
	Repository   string    `json:"repository"`
	Branch       string    `json:"branch"`
	PollInterval time.Time `json:"poll_interval"`
	IsTracked    bool      `json:"is_tracked"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_st"`
}
