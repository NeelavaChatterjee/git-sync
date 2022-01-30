package routes

import (
	"encoding/json"
	"net/http"

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
	track, err := controllers.FindTrack(params["repository"], params["branch"])
	if err != nil {
		json.NewEncoder(w).Encode(&err)
		return
	}
	filtered_poll_logs, err := controllers.FetchFilteredCommitHistory(track.ID)
	if err != nil {
		json.NewEncoder(w).Encode(&err)
		return
	}
	json.NewEncoder(w).Encode(&filtered_poll_logs)
}
