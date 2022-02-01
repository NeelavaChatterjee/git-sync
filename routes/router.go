package routes

import (
	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/manual-poll", TriggerManualPoll).Methods("GET")
	router.HandleFunc("/stop-schedule-poll", StopScheduledPoll).Methods("GET")

	router.HandleFunc("/commit-history", AllCommitHistory).Methods("GET")
	router.HandleFunc("/commit-history/{track_id}", FilteredCommitHistory).Methods("GET")

	router.HandleFunc("/poll-log", AllPollLogs).Methods("GET")
	router.HandleFunc("/poll-log/{track_id}", GetFilteredPollLogs).Methods("GET")

	router.HandleFunc("/track", FindTrack).Methods("GET")
	router.HandleFunc("/track/{track_id}", FindTrackByID).Methods("GET")
	router.HandleFunc("/list-tracks", FetchAllTracked).Methods("GET")
	router.HandleFunc("/track", CreateTrackEntry).Methods("POST")
	router.HandleFunc("/track/{track_id}", DeleteTrackById).Methods("DELETE")
	router.HandleFunc("/untrack/{track_id}", Untrack).Methods("PATCH")
	router.HandleFunc("/retrack/{track_id}", Retrack).Methods("PATCH")
	router.HandleFunc("/update-poll-interval/{track_id}", UpdatePollInterval).Methods("PATCH")

	return router
}
