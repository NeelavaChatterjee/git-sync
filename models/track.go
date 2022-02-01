package models

import (
	"time"

	"github.com/robfig/cron/v3"
	"gorm.io/gorm"
)

type Track struct {
	gorm.Model   `json:"model"`
	ID           uint64       `json:"id"`
	Owner        string       `json:"owner"`
	Repository   string       `json:"repository"`
	Branch       string       `json:"branch"`
	PollInterval string       `json:"poll_interval"`
	CronID       cron.EntryID `json:"cron_id"`
	IsTracked    bool         `json:"is_tracked"`
	CreatedAt    time.Time    `json:"created_at"`
	UpdatedAt    time.Time    `json:"updated_st"`
}
