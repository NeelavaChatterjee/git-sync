package routes

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

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
	track_id, err := strconv.Atoi(params["trackid"])
	filtered_poll_logs, err := controllers.FetchFilteredPollLogs(uint64(track_id))
	if err != nil {
		json.NewEncoder(w).Encode(&err)
		log.Println("Couldn't fetch poll logs from database.")
		return
	}
	json.NewEncoder(w).Encode(&filtered_poll_logs)
}
