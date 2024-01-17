package main

import "github.com/gorilla/mux"

func (app *application) routes() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/tasks", app.all).Methods("GET")
	router.HandleFunc("/tasks/{id}", app.findById).Methods("GET")
	router.HandleFunc("/tasks", app.insertTask).Methods("POST")
	router.HandleFunc("/tasks/{id}", app.updateTask).Methods("PUT")
	router.HandleFunc("/tasks/{title}", app.findByTitle).Methods("GET")

	return router
}