package controllers

import "net/http"

// TODO
// add all the fetched commit details into db
// fetch all the commit history from db
// fetch commit history from db just for a repo and branch
// think about adding pagination

// TODO Gets the whole commit history from db
func GetAllCommitHistory(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	w.Header().Set("Content-Type", "Application/json")

}

// TODO Get a filtered commit history from db
// Filters to be considered: repository, branch, time frame, author
func GetFilteredCommitHistory(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	w.Header().Set("Content-Type", "Application/json")

}

// TODO Creates a new commit entry fetched from github
func CreateNewCommitEntry(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	w.Header().Set("Content-Type", "Application/json")

}

// TODO Deletes a commit entry from db
func DeleteCommitEntryById(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	w.Header().Set("Content-Type", "Application/json")

}
