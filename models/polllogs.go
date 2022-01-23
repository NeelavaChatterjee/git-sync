package models

import (
	"time"

	"gorm.io/gorm"
)

// repo, branch, poll time, number of new files, number of new commits
type PollLogs struct {
	gorm.Model      `json:"model"`
	ID              uint64    `json:"id"`
	NumberOfFiles   int       `json:"number_of_files"`
	NumberOfCommits int       `json:"number_of_commits"`
	CreatedAt       time.Time `json:"created_at"`
	TrackID         uint64    `json:"track_id"`
	Track           Track
}
