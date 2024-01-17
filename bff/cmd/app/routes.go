package main

import "github.com/gorilla/mux"

func (app *application) routes() *mux.Router {
	router := mux.NewRouter()

	// tasks routes
	router.HandleFunc("/tasks", app.getAllTasks).Methods("GET")
	router.HandleFunc("/tasks", app.createTask).Methods("POST")
	router.HandleFunc("/tasks/{id:[0-9]+}", app.getTaskById).Methods("GET")
	router.HandleFunc("/tasks/{title}", app.getTaskByTitle).Methods("GET")
	router.HandleFunc("/tasks/{id:[0-9]+}", app.updateTask).Methods("PUT")
	router.HandleFunc("/tasks/{id:[0-9]+}", app.deleteTask).Methods("DELETE")

	// rewards routes
	router.HandleFunc("/rewards", app.getAllRewards).Methods("GET")
	router.HandleFunc("/rewards", app.createReward).Methods("POST")
	router.HandleFunc("/rewards/{id:[0-9]+}", app.getRewardById).Methods("GET")
	router.HandleFunc("/rewards/{name}", app.getRewardByName).Methods("GET")
	router.HandleFunc("/rewards/{id:[0-9]+}", app.updateReward).Methods("PUT")
	router.HandleFunc("/rewards/{id:[0-9]+}", app.deleteReward).Methods("DELETE")

	// micro-competencies routes
	router.HandleFunc("/micro-competencies", app.getAllMicroCompetencies).Methods("GET")
	router.HandleFunc("/micro-competencies", app.createMicroCompetency).Methods("POST")
	router.HandleFunc("/micro-competencies/{id:[0-9]+}", app.getMicroCompetencyById).Methods("GET")
	router.HandleFunc("/micro-competencies/{name}", app.getMicroCompetencyByName).Methods("GET")
	router.HandleFunc("/micro-competencies/{id:[0-9]+}", app.updateMicroCompetency).Methods("PUT")
	router.HandleFunc("/micro-competencies/{id:[0-9]+}", app.deleteMicroCompetency).Methods("DELETE")

	return router
}