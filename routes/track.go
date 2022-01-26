package routes

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/NeelavaChatterjee/git-sync/controllers"
	"github.com/NeelavaChatterjee/git-sync/models"
)

// TODO remove all the dummy stuff and add stuff from requests

func FindTrack(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	w.Header().Set("Content-Type", "Application/json")
	track, err := controllers.FindTrack("repo", "branch")
	if err != nil {
		json.NewEncoder(w).Encode(&err)
	}
	json.NewEncoder(w).Encode(&track)
}

func FindTrackByID(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	w.Header().Set("Content-Type", "application/json")
	track, err := controllers.FindTrackByID(67)
	if err != nil {
		json.NewEncoder(w).Encode(&err)
	}
	json.NewEncoder(w).Encode(&track)
}

func FetchAllTracked(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	w.Header().Set("Content-Type", "application/json")
	tracks, err := controllers.FetchAllTracked()
	if err != nil {
		json.NewEncoder(w).Encode(&err)
	}
	json.NewEncoder(w).Encode(&tracks)
}

func CreateTrackEntry(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	w.Header().Set("Content-Type", "Application/json")
	// TODO get the actual data from request
	err := controllers.CreateTrackEntry(&models.Track{})
	if err != nil {
		json.NewEncoder(w).Encode(&err)
	}
	json.NewEncoder(w).Encode("The track was added")
}

func DeleteTrackById(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	w.Header().Set("Content-Type", "application/json")
	err := controllers.DeleteTrackById(67)
	if err != nil {
		json.NewEncoder(w).Encode(&err)
	}
	json.NewEncoder(w).Encode("The track was deleted")
}

func ToggleTrack(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	w.Header().Set("Content-Type", "application/json")
	err := controllers.ToggleTrack(&models.Track{})
	if err != nil {
		json.NewEncoder(w).Encode(&err)
	}
	json.NewEncoder(w).Encode("Track toggled")
}

func UpdatePollInterval(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	w.Header().Set("Content-Type", "application/json")
	err := controllers.UpdatePollInterval(&models.Track{}, time.Time{})
	if err != nil {
		json.NewEncoder(w).Encode(&err)
	}
	json.NewEncoder(w).Encode("Time interval updated")
}
