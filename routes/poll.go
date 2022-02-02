package routes

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/NeelavaChatterjee/git-sync/controllers"
	"github.com/NeelavaChatterjee/git-sync/utilities"
	"github.com/gorilla/mux"
	"github.com/robfig/cron/v3"
)

func TriggerManualPoll(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	track_id, err := strconv.Atoi(params["track_id"])
	if err != nil {
		panic(err)
	}
	track, err := controllers.FindTrackByID(uint64(track_id))
	if err != nil {
		json.NewEncoder(w).Encode(err.Error())
		log.Println("Couldn't find the requested tracks")
		return
	}

	// Doing this in a different go routine as it takes a lot of time
	go controllers.Poll(track)
	json.NewEncoder(w).Encode("Polling InProgress")
}

func StopScheduledPoll(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	track_cron_id, err := strconv.Atoi(params["track_cron_id"])
	track_cron_entry_id := cron.EntryID(track_cron_id)
	utilities.Cron.Remove(track_cron_entry_id)
	if err != nil {
		json.NewEncoder(w).Encode(err.Error())
	}
	json.NewEncoder(w).Encode("Next Poll Removed")
}

func ReSchedulePoll(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	track_id, err := strconv.Atoi(params["track_id"])
	if err != nil {
		panic(err)
	}

	track, err := controllers.FindTrackByID(uint64(track_id))
	if err != nil {
		panic(err)
	}

	cron_id := controllers.SchedulePoll(track)

	err = controllers.UpdateCronEntryID(track, cron_id)
	if err != nil {
		json.NewEncoder(w).Encode(err.Error())
	}
	json.NewEncoder(w).Encode("Poll added")
}
