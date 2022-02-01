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
	track_id, err := strconv.Atoi(params["trackid"])
	if err != nil {
		panic(err)
	}
	track, err := controllers.FindTrackByID(uint64(track_id))
	if err != nil {
		json.NewEncoder(w).Encode(err.Error())
		log.Println("Couldn't find the requested tracks")
		return
	}

	controllers.Poll(track)
	if err != nil {
		json.NewEncoder(w).Encode(err.Error())
	}
	json.NewEncoder(w).Encode("Polling Done")
}

// TODO Needs to be cross checked
func StopScheduledPoll(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	track_cron_id, err := strconv.Atoi(params["trackcronid"])
	track_cron_entry_id := cron.EntryID(track_cron_id)
	utilities.Cron.Remove(track_cron_entry_id)
	if err != nil {
		json.NewEncoder(w).Encode(err.Error())
	}
	json.NewEncoder(w).Encode("Next Poll Removed")
}
