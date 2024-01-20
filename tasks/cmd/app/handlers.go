package main

import (
	"encoding/json"
	"net/http"

	"github.com/cpprian/stucoin-backend/tasks/pkg/models"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (app *application) all(w http.ResponseWriter, r *http.Request) {
	// Get all tasks
	tasks, err := app.tasks.All()
	if err != nil {
		app.errorLog.Println("Error getting all tasks: ", err)
		app.serverError(w, err)
		return
	}

	// Convert task list into json encoding
	b, err := json.Marshal(tasks)
	if err != nil {
		app.errorLog.Println("Error marshalling tasks: ", err)
		app.serverError(w, err)
		return
	}

	app.infoLog.Println("\nAll tasks were sent")

	// Send response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(b)
}

func (app *application) getAllTeacherTasks(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["owner"]
	app.infoLog.Printf("Getting all tasks from teacher with id %s\n", id)

	tasks, err := app.tasks.AllTeacherTasks(id)
	if err != nil {
		app.errorLog.Println("Error getting all tasks: ", err)
		app.serverError(w, err)
		return
	}

	// Convert task list into json encoding
	b, err := json.Marshal(tasks)
	if err != nil {
		app.errorLog.Println("Error marshalling tasks: ", err)
		app.serverError(w, err)
		return
	}

	app.infoLog.Println("\nAll tasks were sent")
	app.infoLog.Println("\nTasks:", tasks)

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
	task, err := app.tasks.FindById(id)
	if err != nil {
		if err.Error() == "no task found" {
			app.infoLog.Println("Task not found")
			return
		}
		app.errorLog.Println("Error getting task: ", err)
		app.serverError(w, err)
		return
	}
	app.infoLog.Println("\nTask:", task)

	// Convert task into json encoding
	b, err := json.Marshal(task)
	if err != nil {
		app.errorLog.Println("Error marshalling task: ", err)
		app.serverError(w, err)
		return
	}

	app.infoLog.Println("\nTask was sent")

	// Send response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(b)
}

func (app *application) findByTitle(w http.ResponseWriter, r *http.Request) {
	// Get taskname from request
	taskname := mux.Vars(r)["taskname"]
	app.infoLog.Printf("Getting task with taskname %s\n", taskname)

	// Get task
	task, err := app.tasks.FindByTitle(taskname)
	if err != nil {
		if err.Error() == "no task found" {
			app.infoLog.Println("Task not found")
			return
		}
		app.errorLog.Println("Error getting task: ", err)
		app.serverError(w, err)
		return
	}
	app.infoLog.Println("\nTask:", task)

	// Convert task into json encoding
	b, err := json.Marshal(task)
	if err != nil {
		app.errorLog.Println("Error marshalling task: ", err)
		app.serverError(w, err)
		return
	}

	app.infoLog.Println("\nTask was sent")

	// Send response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(b)
}

func (app *application) insertTask(w http.ResponseWriter, r *http.Request) {
	// Get task from request
	var task models.Task
	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		app.errorLog.Println("Error:", err)
		app.serverError(w, err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	app.infoLog.Println("\nTask:", task)

	// Insert task
	resp, err := app.tasks.InsertTask(&task)
	if err != nil {
		app.errorLog.Println("Error:", err)
		app.serverError(w, err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	app.infoLog.Printf("Task was inserted with data %+v and id %p\n", task, resp.InsertedID)

	// Encode response as JSON
	responseData := struct {
		InsertedID string `json:"insertedID"`
	}{
		InsertedID: resp.InsertedID.(primitive.ObjectID).Hex(),
	}
	app.infoLog.Println("\nResponse:", responseData)

	encodedResponse, err := json.Marshal(responseData)
	if err != nil {
		app.errorLog.Println("Error encoding response:", err)
		app.serverError(w, err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Send JSON response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(encodedResponse)
}

func (app *application) updateTask(w http.ResponseWriter, r *http.Request) {
	// Get task from request
	var task models.Task
	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		app.errorLog.Println("Error:", err)
		app.serverError(w, err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	app.infoLog.Println("\nTask:", task)

	// Update task
	_, err = app.tasks.UpdateTask(&task)
	if err != nil {
		app.errorLog.Println("Error:", err)
		app.serverError(w, err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	app.infoLog.Println("Task was updated with id:", task.ID)

	// Send response
	w.WriteHeader(http.StatusOK)
}

func (app *application) deleteTask(w http.ResponseWriter, r *http.Request) {
	// Get task id from request
	id := mux.Vars(r)["id"]
	app.infoLog.Printf("Deleting task with id %s\n", id)

	// Delete task
	_, err := app.tasks.DeleteTask(id)
	if err != nil {
		app.errorLog.Println("Error:", err)
		app.serverError(w, err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	app.infoLog.Println("Task was deleted with id:", id)

	// Send response
	w.WriteHeader(http.StatusOK)
}

func (app *application) updateCoverImageById(w http.ResponseWriter, r *http.Request) {
	// Get task id from request
	id := mux.Vars(r)["id"]
	app.infoLog.Printf("Updating cover image from task with id %s\n", id)

	var coverImage models.CoverImage
	err := json.NewDecoder(r.Body).Decode(&coverImage)
	if err != nil {
		app.errorLog.Println("Error:", err)
		app.serverError(w, err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	app.infoLog.Println("\nCover image:", coverImage)
	_, err = app.tasks.UpdateCoverImageById(id, coverImage)
	if err != nil {
		app.errorLog.Println("Error:", err)
		app.serverError(w, err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	app.infoLog.Println("Cover image was updated from task with id:", id)

	w.WriteHeader(http.StatusOK)
}
