package routes

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/NeelavaChatterjee/git-sync/controllers"
	"github.com/gorilla/mux"
)

func AllPollLogs(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	w.Header().Set("Content-Type", "application/json")
	poll_logs, err := controllers.FetchAllPollLogs()
	if err != nil {
		json.NewEncoder(w).Encode(&err)
		log.Println("Couldn't fetch poll logs from database.")
	}
	json.NewEncoder(w).Encode(&poll_logs)
}

func GetFilteredPollLogs(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	w.Header().Set("Content-Type", "Application/json")
	params := mux.Vars(r)
	track, err := controllers.FindTrack(params["repository"], params["branch"])
	if err != nil {
		json.NewEncoder(w).Encode(&err)
		log.Println("Couldn't find the requested tracks")
		return
	}
	filtered_poll_logs, err := controllers.FetchFilteredPollLogs(track.ID)
	if err != nil {
		json.NewEncoder(w).Encode(&err)
		log.Println("Couldn't fetch poll logs from database.")
		return
	}
	json.NewEncoder(w).Encode(&filtered_poll_logs)
}
