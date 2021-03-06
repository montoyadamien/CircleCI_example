package controllers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"tree/entities"
	"tree/services"
)


func HandleRoutes(r *mux.Router) {
	r.Handle("/car-booking/findAll", findAll(),
	).Methods("GET", "OPTIONS").Name("findAllBookings")
}

func findAll() http.Handler{
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		jsonError := json.NewEncoder(w).Encode(services.GetService().FindAll())
		if jsonError != nil {
			e := entities.JSONError{Message: "Internal Server Error"}
			w.WriteHeader(http.StatusInternalServerError)
			err2 := json.NewEncoder(w).Encode(e)
			log.Panic(jsonError, err2)
		}
	})
}
