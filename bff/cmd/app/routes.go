package main

import "github.com/gorilla/mux"

func (app *application) routes() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/tasks", app.getAllTasks).Methods("GET")
	router.HandleFunc("/tasks/create", app.createTask).Methods("POST")
	router.HandleFunc("/tasks/view/{id:[0-9]+}", app.getTaskById).Methods("GET")
	router.HandleFunc("/tasks/view/{title}", app.getTaskByTitle).Methods("GET")
	router.HandleFunc("/tasks/update/{id:[0-9]+}", app.updateTask).Methods("PUT")
	router.HandleFunc("/tasks/delete/{id:[0-9]+}", app.deleteTask).Methods("DELETE")

	return router
}