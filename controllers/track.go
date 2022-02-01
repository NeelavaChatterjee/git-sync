package controllers

import (
	"github.com/NeelavaChatterjee/git-sync/database"
	"github.com/NeelavaChatterjee/git-sync/models"
	"github.com/NeelavaChatterjee/git-sync/utilities"
	"github.com/robfig/cron/v3"
)

func FindTrack(owner string, repository string, branch string) (*models.Track, error) {
	var track models.Track
	result := database.Db.Where(&models.Track{Owner: owner, Repository: repository, Branch: branch}).First(&track)
	if result.Error != nil {
		return nil, result.Error
	}
	return &track, nil
}

func FindTrackByID(track_id uint64) (*models.Track, error) {
	var track models.Track
	result := database.Db.First(&track, track_id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &track, nil
}

func FetchAllTracked() (*[]models.Track, error) {
	var tracks []models.Track
	result := database.Db.Find(&tracks)
	if result.Error != nil {
		return nil, result.Error
	}
	return &tracks, nil
}

func CreateTrackEntry(track *models.Track) error {
	cron_id := SchedulePoll(track)
	track.CronID = cron_id
	result := database.Db.Create(track)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func DeleteTrackById(track_id uint64) error {
	result := database.Db.Delete(&models.Track{}, track_id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func UnTrack(track *models.Track) error {
	utilities.Cron.Remove(track.CronID)
	track.IsTracked = false
	result := database.Db.Save(track)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func ReTrack(track *models.Track) error {
	SchedulePoll(track)
	track.IsTracked = true
	result := database.Db.Save(track)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func UpdatePollInterval(track *models.Track, new_poll_interval string) error {
	track.PollInterval = new_poll_interval
	utilities.Cron.Remove(track.CronID)
	cron_entry_id := SchedulePoll(track)
	track.CronID = cron_entry_id
	result := database.Db.Save(track)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func UpdateCronEntryID(track *models.Track, new_entry_id cron.EntryID) error {
	track.CronID = new_entry_id
	result := database.Db.Save(track)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
