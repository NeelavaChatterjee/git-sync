package controllers

import (
	"github.com/NeelavaChatterjee/git-sync/database"
	"github.com/NeelavaChatterjee/git-sync/models"
)

// Fetches all poll logs from database
func FetchAllPollLogs() ([]models.PollLogs, error) {
	var poll_logs []models.PollLogs
	result := database.Db.Find(&poll_logs)
	if result.Error != nil {
		return nil, result.Error
	}
	return poll_logs, nil
}

// Fetches filtered poll logs based on repo and branch from db
func FetchFilteredPollLogs(track_id uint64) ([]models.PollLogs, error) {
	var filtered_poll_logs []models.PollLogs
	result := database.Db.Where(&models.PollLogs{TrackID: track_id}).Find(&filtered_poll_logs)
	if result.Error != nil {
		return nil, result.Error
	}
	return filtered_poll_logs, nil
}

func FetchLastPollLog(track_id uint64) (*models.PollLogs, error) {
	var last_poll_log models.PollLogs
	result := database.Db.Where(&models.PollLogs{TrackID: track_id}).Last(&last_poll_log)
	if result.Error != nil {
		return nil, result.Error
	}
	return &last_poll_log, nil
}

// Creates a new Poll Log entry in db
func CreatePollLog(poll_log_entry *models.PollLogs) error {
	result := database.Db.Create(poll_log_entry)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func DeletePollLogById(poll_log_id uint64) error {
	result := database.Db.Delete(&models.PollLogs{}, poll_log_id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
