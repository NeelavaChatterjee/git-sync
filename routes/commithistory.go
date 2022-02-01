package routes

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/NeelavaChatterjee/git-sync/controllers"
	"github.com/gorilla/mux"
)

func AllCommitHistory(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	w.Header().Set("Content-Type", "application/json")
	poll_logs, err := controllers.FetchAllCommitHistory()
	if err != nil {
		json.NewEncoder(w).Encode(&err)
		return
	}
	json.NewEncoder(w).Encode(&poll_logs)
}

func FilteredCommitHistory(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	w.Header().Set("Content-Type", "Application/json")
	params := mux.Vars(r)
	track_id, err := strconv.Atoi(params["trackid"])
	filtered_poll_logs, err := controllers.FetchFilteredCommitHistory(uint64(track_id))
	if err != nil {
		json.NewEncoder(w).Encode(&err)
		return
	}
	json.NewEncoder(w).Encode(&filtered_poll_logs)
}
