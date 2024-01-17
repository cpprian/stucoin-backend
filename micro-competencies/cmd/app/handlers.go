package main

import (
	"encoding/json"
	"net/http"

	"github.com/cpprian/stucoin-backend/micro-competencies/pkg/models"
	"github.com/gorilla/mux"
)

func (app *application) all(w http.ResponseWriter, r *http.Request) {
	// Get all microCompetencies
	microCompetencies, err := app.microCompetencies.All()
	if err != nil {
		app.errorLog.Println("Error getting all microCompetencies: ", err)
		app.serverError(w, err)
		return
	}

	// Convert task list into json encoding
	b, err := json.Marshal(microCompetencies)
	if err != nil {
		app.errorLog.Println("Error marshalling microCompetencies: ", err)
		app.serverError(w, err)
		return
	}

	app.infoLog.Println("\nAll microCompetencies were sent")

	// Send response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(b)
}

func (app *application) findById(w http.ResponseWriter, r *http.Request) {
	// Get task id from request
	id := mux.Vars(r)["id"]
	app.infoLog.Printf("Getting task with id %s\n", id)

	// Get task
	task, err := app.microCompetencies.FindById(id)
	if err != nil {
		if err.Error() == "no task found" {
			app.infoLog.Println("MicroCompetence not found")
			return
		}
		app.errorLog.Println("Error getting task: ", err)
		app.serverError(w, err)
		return
	}
	app.infoLog.Println("\nMicroCompetence:", task)

	// Convert task into json encoding
	b, err := json.Marshal(task)
	if err != nil {
		app.errorLog.Println("Error marshalling task: ", err)
		app.serverError(w, err)
		return
	}

	app.infoLog.Println("\nMicroCompetence was sent")

	// Send response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(b)
}

func (app *application) findByName(w http.ResponseWriter, r *http.Request) {
	// Get name from request
	name := mux.Vars(r)["name"]
	app.infoLog.Printf("Getting task with name %s\n", name)

	// Get task
	task, err := app.microCompetencies.FindByName(name)
	if err != nil {
		if err.Error() == "no task found" {
			app.infoLog.Println("MicroCompetence not found")
			return
		}
		app.errorLog.Println("Error getting task: ", err)
		app.serverError(w, err)
		return
	}
	app.infoLog.Println("\nMicroCompetence:", task)

	// Convert task into json encoding
	b, err := json.Marshal(task)
	if err != nil {
		app.errorLog.Println("Error marshalling task: ", err)
		app.serverError(w, err)
		return
	}

	app.infoLog.Println("\nMicroCompetence was sent")

	// Send response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(b)
}

func (app *application) insertMicroCompetence(w http.ResponseWriter, r *http.Request) {
	// Get task from request
	var task models.MicroCompetence
	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		app.errorLog.Println("Error:", err)
		app.serverError(w, err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	app.infoLog.Println("\nMicroCompetence:", task)

	// Insert task
	_, err = app.microCompetencies.InsertMicroCompetence(&task)
	if err != nil {
		app.errorLog.Println("Error:", err)
		app.serverError(w, err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	app.infoLog.Println("MicroCompetence was inserted with data:", task)

	// Send response
	w.WriteHeader(http.StatusOK)
}

func (app *application) updateMicroCompetence(w http.ResponseWriter, r *http.Request) {
	// Get task from request
	var task models.MicroCompetence
	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		app.errorLog.Println("Error:", err)
		app.serverError(w, err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	app.infoLog.Println("\nMicroCompetence:", task)

	// Update task
	_, err = app.microCompetencies.UpdateMicroCompetence(&task)
	if err != nil {
		app.errorLog.Println("Error:", err)
		app.serverError(w, err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	app.infoLog.Println("MicroCompetence was updated with id:", task.ID)

	// Send response
	w.WriteHeader(http.StatusOK)
}