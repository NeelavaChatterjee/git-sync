package routes

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/NeelavaChatterjee/git-sync/controllers"
	"github.com/NeelavaChatterjee/git-sync/models"
	"github.com/gorilla/mux"
)

func FindTrack(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	w.Header().Set("Content-Type", "Application/json")
	params := mux.Vars(r)
	track, err := controllers.FindTrack(params["repository"], params["branch"])
	if err != nil {
		json.NewEncoder(w).Encode(err.Error())
		log.Println("Couldn't find the requested tracks")
		return
	}
	json.NewEncoder(w).Encode(&track)
}

func FindTrackByID(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	track_id, err := strconv.Atoi(params["track_id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err.Error())
		log.Println(err)
		return
	}
	track, err := controllers.FindTrackByID(uint64(track_id))
	if err != nil {
		json.NewEncoder(w).Encode(err.Error())
		log.Println(err, "Couldn't find the requested tracks")
		return
	}
	json.NewEncoder(w).Encode(&track)
}

func FetchAllTracked(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	w.Header().Set("Content-Type", "application/json")
	tracks, err := controllers.FetchAllTracked()
	if err != nil {
		json.NewEncoder(w).Encode(err.Error())
	}
	json.NewEncoder(w).Encode(&tracks)
}

func CreateTrackEntry(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	w.Header().Set("Content-Type", "Application/json")
	var track models.Track

	if err := json.NewDecoder(r.Body).Decode(&track); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err.Error())
		log.Println(err)
		return
	}

	if err := controllers.CreateTrackEntry(&track); err != nil {
		json.NewEncoder(w).Encode(err.Error())
		return
	}
	json.NewEncoder(w).Encode("The track was added")
}

func DeleteTrackById(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	track_id, err := strconv.Atoi(params["track_id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err.Error())
		log.Println(err)
		return
	}

	if err := controllers.DeleteTrackById(uint64(track_id)); err != nil {
		json.NewEncoder(w).Encode(err.Error())
		return
	}
	json.NewEncoder(w).Encode("The track was deleted")
}

func ToggleTrack(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	track, err := controllers.FindTrack(params["repository"], params["branch"])
	if err != nil {
		json.NewEncoder(w).Encode(err.Error())
		log.Println("Couldn't find the requested tracks")
		return
	}

	is_tracked, err := controllers.ToggleTrack(track)
	if err != nil {
		json.NewEncoder(w).Encode(err.Error())
		return
	}
	log.Println("is_tracked:", is_tracked)
	json.NewEncoder(w).Encode(is_tracked)
}

func UpdatePollInterval(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	track, err := controllers.FindTrack(params["repository"], params["branch"])
	if err != nil {
		json.NewEncoder(w).Encode(err.Error())
		log.Println("Couldn't find the requested tracks")
		return
	}
	// TODO Fetch the time from the request
	err = controllers.UpdatePollInterval(track, time.Time{})
	if err != nil {
		json.NewEncoder(w).Encode(err.Error())
	}
	json.NewEncoder(w).Encode("Time interval updated")
}
