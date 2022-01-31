package controllers

import (
	"github.com/NeelavaChatterjee/git-sync/database"
	"github.com/NeelavaChatterjee/git-sync/models"
)

func FindTrack(repository string, branch string) (*models.Track, error) {
	var track models.Track
	result := database.Db.Where(&models.Track{Repository: repository, Branch: branch}).First(&track)
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

// TODO
// Updates is_tracking field
func ToggleTrack(track *models.Track) (bool, error) {
	track.IsTracked = !(track.IsTracked)
	result := database.Db.Save(track)
	if result.Error != nil {
		// Reverting it back incase it is not saved
		track.IsTracked = !(track.IsTracked)
		return track.IsTracked, result.Error
	}
	return track.IsTracked, nil
}

// TODO
func UpdatePollInterval(track *models.Track, new_poll_interval string) error {
	track.PollInterval = new_poll_interval
	result := database.Db.Save(track)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
