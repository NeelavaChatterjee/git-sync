package routes

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/NeelavaChatterjee/git-sync/controllers"
	"github.com/gorilla/mux"
)

func TriggerManualPoll(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	track, err := controllers.FindTrack(params["repository"], params["branch"])
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
