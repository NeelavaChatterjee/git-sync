package controllers

import (
	"time"

	"github.com/NeelavaChatterjee/git-sync/database"
	"github.com/NeelavaChatterjee/git-sync/models"
)

func FindTrack(repository string, branch string) (*models.Track, error) {
	var track models.Track
	r := database.Db.Where(&models.Track{Repository: repository, Branch: branch}).First(&track)
	if r.Error != nil {
		return nil, r.Error
	}
	return &track, nil
}

func FindTrackByID(track_id uint64) (*models.Track, error) {
	var track models.Track
	r := database.Db.First(&track, track_id)
	if r.Error != nil {
		return nil, r.Error
	}
	return &track, nil
}

func FetchAllTracked() (*[]models.Track, error) {
	var tracks []models.Track
	r := database.Db.Find(&tracks)
	if r.Error != nil {
		return nil, r.Error
	}
	return &tracks, nil
}

func CreateTrackEntry(track *models.Track) error {
	r := database.Db.Create(track)
	if r.Error != nil {
		return r.Error
	}
	return nil
}

func DeleteTrackById(track_id uint64) error {
	r := database.Db.Delete(&models.Track{}, track_id)
	if r.Error != nil {
		return r.Error
	}
	return nil
}

// Updates is_tracking field
func ToggleTrack(track *models.Track) error {
	track.IsTracked = !(track.IsTracked)
	r := database.Db.Save(track)
	if r.Error != nil {
		return r.Error
	}
	return nil
}

func UpdatePollInterval(track *models.Track, new_poll_interval time.Time) error {
	track.PollInterval = new_poll_interval
	r := database.Db.Save(track)
	if r.Error != nil {
		return r.Error
	}
	return nil
}
