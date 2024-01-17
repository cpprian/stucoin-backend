package main

import "github.com/gorilla/mux"

func (app *application) routes() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/micro-competencies", app.all).Methods("GET")
	router.HandleFunc("/micro-competencies/{id}", app.findById).Methods("GET")
	router.HandleFunc("/micro-competencies", app.insertMicroCompetence).Methods("POST")
	router.HandleFunc("/micro-competencies/{id}", app.updateMicroCompetence).Methods("PUT")
	router.HandleFunc("/micro-competencies/{title}", app.findByTitle).Methods("GET")

	return router
}