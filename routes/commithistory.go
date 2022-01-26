package routes

import (
	"encoding/json"
	"net/http"

	"github.com/NeelavaChatterjee/git-sync/controllers"
)

func AllCommitHistory(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	w.Header().Set("Content-Type", "application/json")
	poll_logs, err := controllers.FetchAllCommitHistory()
	if err != nil {
		json.NewEncoder(w).Encode(&err)
	}
	json.NewEncoder(w).Encode(&poll_logs)
}

func FilteredCommitHistory(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	w.Header().Set("Content-Type", "Application/json")
	// TODO get the actual data from request and 56 here is a dummy data
	filtered_poll_logs, err := controllers.FetchFilteredCommitHistory(56)
	if err != nil {
		json.NewEncoder(w).Encode(&err)
	}
	json.NewEncoder(w).Encode(&filtered_poll_logs)
}
