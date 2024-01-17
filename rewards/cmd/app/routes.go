package main

import "github.com/gorilla/mux"

func (app *application) routes() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/rewards", app.all).Methods("GET")
	router.HandleFunc("/rewards/{id}", app.findById).Methods("GET")
	router.HandleFunc("/rewards", app.insertReward).Methods("POST")
	router.HandleFunc("/rewards/{id}", app.updateReward).Methods("PUT")
	router.HandleFunc("/rewards/{name}", app.findByName).Methods("GET")

	return router
}