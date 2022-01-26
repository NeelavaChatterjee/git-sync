package controllers

import (
	"github.com/NeelavaChatterjee/git-sync/database"
	"github.com/NeelavaChatterjee/git-sync/models"
)

// Fetches all poll logs from database
func FetchAllPollLogs() ([]models.PollLogs, error) {
	var poll_logs []models.PollLogs
	r := database.Db.Find(&poll_logs)
	if r.Error != nil {
		return nil, r.Error
	}
	return poll_logs, nil
}

// Fetches filtered poll logs based on repo and branch from db
// TODO Filters yet to be considered: time frame
func FetchFilteredPollLogs(track_id uint64) ([]models.PollLogs, error) {
	var filtered_poll_logs []models.PollLogs
	r := database.Db.Where(&models.PollLogs{TrackID: track_id}).Find(&filtered_poll_logs)
	if r.Error != nil {
		return nil, r.Error
	}
	return filtered_poll_logs, nil
}

// Creates a new Poll Log entry in db
func CreatePollLog(poll_log_entry *models.PollLogs) error {
	r := database.Db.Create(poll_log_entry)
	if r.Error != nil {
		return r.Error
	}
	return nil
}

func DeletePollLogById(poll_log_id uint64) error {
	r := database.Db.Delete(&models.PollLogs{}, poll_log_id)
	if r.Error != nil {
		return r.Error
	}
	return nil
}
