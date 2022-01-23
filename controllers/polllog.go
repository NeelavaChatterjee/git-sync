package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/NeelavaChatterjee/git-sync/database"
	"github.com/NeelavaChatterjee/git-sync/models"
)

//TODO

// create a poll log
// delete a poll log(not to be used)
// fetch all poll logs
// fetch poll logs for a repo and branch
// think about adding pagination

// Fetches all poll logs from database
func GetAllPollLogs(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	w.Header().Set("Content-Type", "application/json")
	var poll_logs []models.PollLogs
	database.Db.Find(&poll_logs)
	json.NewEncoder(w).Encode(&poll_logs)
}

// TODO Fetches filtered poll logs based on repo and branch from db
// Filters to be considered: repository, branch, time frame
func GetFilteredPollLogs(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	w.Header().Set("Content-Type", "Application/json")

}

// TODO Creates a new Poll Log entry in db
func CreatePollLog(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	w.Header().Set("Content-Type", "Application/json")

}

// TODO Deletes a Poll Log from db based on its id.
func DeletePollLogById(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	w.Header().Set("Content-Type", "Application/json")

}
