package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func (app *application) routes() http.Handler {
	router := mux.NewRouter()

	// tasks routes
	router.HandleFunc("/tasks", app.getAllTasks).Methods("GET")
	router.HandleFunc("/tasks", app.createTask).Methods("POST")
	router.HandleFunc("/tasks/{id}", app.getTaskById).Methods("GET")
	// router.HandleFunc("/tasks/title/{title}", app.getTaskByTitle).Methods("GET")
	router.HandleFunc("/tasks/{id}", app.updateTask).Methods("PUT")
	router.HandleFunc("/tasks/{id}", app.deleteTask).Methods("DELETE")
	router.HandleFunc("/tasks/teacher/{owner}", app.getAllTasksByOwnerId).Methods("GET")
	router.HandleFunc("/tasks/cover/{id}", app.updateCoverImageById).Methods("PUT")
	router.HandleFunc("/tasks/content/{id}", app.updateContentById).Methods("PUT")
	router.HandleFunc("/tasks/title/{id}", app.updateTitleById).Methods("PUT")
	router.HandleFunc("/tasks/files/{id}", app.saveFilesById).Methods("POST")
	router.HandleFunc("/tasks/files/{id}", app.deleteFilesById).Methods("DELETE")
	router.HandleFunc("/tasks/assign/{id}", app.assignTaskById).Methods("PUT")
	router.HandleFunc("/tasks/complete/{id}", app.completeTaskById).Methods("PUT")
	router.HandleFunc("/tasks/accept/{id}", app.acceptTaskById).Methods("PUT")
	router.HandleFunc("/tasks/reject/{id}", app.rejectTaskById).Methods("PUT")

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

	return cors.Default().Handler(router)
}