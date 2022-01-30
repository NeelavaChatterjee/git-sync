package routes

import (
	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()
	// router.HandleFunc("/hotel", GetHotelDetails).Methods("GET")
	// router.HandleFunc("/hotel", CreateHotel).Methods("POST")
	// router.HandleFunc("/guest", CreateGuest).Methods("POST")
	// router.HandleFunc("/guest/{id}", GetGuestByID).Methods("GET")
	// router.HandleFunc("/guests", GetAllGuests).Methods("GET")
	// router.HandleFunc("/guest/{id}", UpdateCheckOutDate).Methods("PUT")
	// router.HandleFunc("/guest/{id}", DeleteGuest).Methods("DELETE")

	return router
}
