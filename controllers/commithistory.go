package controllers

import (
	"github.com/NeelavaChatterjee/git-sync/database"
	"github.com/NeelavaChatterjee/git-sync/models"
)

// Gets the whole commit history from db
func FetchAllCommitHistory() ([]models.CommitHistory, error) {
	var commit_history []models.CommitHistory
	result := database.Db.Find(&commit_history)
	if result.Error != nil {
		return nil, result.Error
	}
	return commit_history, nil
}

// Get a filtered commit history from db
func FetchFilteredCommitHistory(track_id uint64) ([]models.CommitHistory, error) {
	var filtered_commit_history []models.CommitHistory
	result := database.Db.Where(&models.CommitHistory{TrackID: track_id}).Find(&filtered_commit_history)
	if result.Error != nil {
		return nil, result.Error
	}
	return filtered_commit_history, nil
}

// Creates a new commit entry fetched from github
func CreateNewCommitHistoryEntry(commit_history_entry *models.CommitHistory) error {
	result := database.Db.Create(commit_history_entry)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// Deletes a commit entry from db
func DeleteCommitHistoryById(commit_history_id uint64) error {
	result := database.Db.Delete(&models.CommitHistory{}, commit_history_id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
