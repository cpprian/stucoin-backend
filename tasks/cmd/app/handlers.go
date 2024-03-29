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

func (app *application) updateContentById(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	app.infoLog.Printf("Updating content from task with id %s\n", id)

	var content models.Content
	err := json.NewDecoder(r.Body).Decode(&content)
	if err != nil {
		app.errorLog.Println("Error:", err)
		app.serverError(w, err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	app.infoLog.Println("\nContent:", content)
	_, err = app.tasks.UpdateContentById(id, content)
	if err != nil {
		app.errorLog.Println("Error:", err)
		app.serverError(w, err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	app.infoLog.Println("Content was updated from task with id:", id)

	w.WriteHeader(http.StatusOK)
}

func (app *application) updateTitleById(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	app.infoLog.Printf("Updating title from task with id %s\n", id)

	var title models.Title
	err := json.NewDecoder(r.Body).Decode(&title)
	if err != nil {
		app.errorLog.Println("Error:", err)
		app.serverError(w, err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	app.infoLog.Println("\nTitle:", title)
	_, err = app.tasks.UpdateTitleById(id, title)
	if err != nil {
		app.errorLog.Println("Error:", err)
		app.serverError(w, err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	app.infoLog.Println("Title was updated from task with id:", id)

	w.WriteHeader(http.StatusOK)
}

func (app *application) saveFilesById(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	app.infoLog.Printf("Saving files from task with id %s\n", id)

	var file models.File
	err := json.NewDecoder(r.Body).Decode(&file)
	if err != nil {
		app.errorLog.Println("Error:", err)
		app.serverError(w, err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	app.infoLog.Println("\nFiles:", file)
	_, err = app.tasks.SaveFilesById(id, file)
	if err != nil {
		app.errorLog.Println("Error:", err)
		app.serverError(w, err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	app.infoLog.Println("Files were saved from task with id:", id)

	w.WriteHeader(http.StatusOK)
}

func (app *application) deleteFileById(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	app.infoLog.Printf("Deleting file from task with id %s\n", id)

	var file models.File
	err := json.NewDecoder(r.Body).Decode(&file)
	if err != nil {
		app.errorLog.Println("Error:", err)
		app.serverError(w, err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	app.infoLog.Println("\nFiles:", file)
	_, err = app.tasks.DeleteFileById(id, file)
	if err != nil {
		app.errorLog.Println("Error:", err)
		app.serverError(w, err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	app.infoLog.Println("File was deleted from task with id:", id)

	w.WriteHeader(http.StatusOK)
}

func (app *application) assignTaskById(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	app.infoLog.Printf("Assigning task with id %s\n", id)

	var inCharge models.InCharge
	err := json.NewDecoder(r.Body).Decode(&inCharge)
	if err != nil {
		app.errorLog.Println("Error:", err)
		app.serverError(w, err)
		return
	}

	app.infoLog.Println("\nInCharge:", inCharge)
	_, err = app.tasks.AssignTaskById(id, inCharge)
	if err != nil {
		app.errorLog.Println("Error:", err)
		app.serverError(w, err)
		return
	}

	app.infoLog.Println("Task was assigned with id:", id)

	w.WriteHeader(http.StatusOK)
}

func (app *application) completeTaskById(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	app.infoLog.Printf("Completing task with id %s\n", id)

	_, err := app.tasks.CompleteTaskById(id)
	if err != nil {
		app.errorLog.Println("Error:", err)
		app.serverError(w, err)
		return
	}

	app.infoLog.Println("Task was completed with id:", id)

	w.WriteHeader(http.StatusOK)
}

func (app *application) acceptTaskById(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	app.infoLog.Printf("Accepting task with id %s\n", id)

	_, err := app.tasks.AcceptTaskById(id)
	if err != nil {
		app.errorLog.Println("Error:", err)
		app.serverError(w, err)
		return
	}

	app.infoLog.Println("Task was accepted with id:", id)

	w.WriteHeader(http.StatusOK)
}

func (app *application) rejectTaskById(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	app.infoLog.Printf("Rejecting task with id %s\n", id)

	_, err := app.tasks.RejectTaskById(id)
	if err != nil {
		app.errorLog.Println("Error:", err)
		app.serverError(w, err)
		return
	}

	app.infoLog.Println("Task was rejected with id:", id)

	w.WriteHeader(http.StatusOK)
}

func (app *application) updatePointsById(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	app.infoLog.Printf("Updating points from task with id %s\n", id)

	var points models.Points
	err := json.NewDecoder(r.Body).Decode(&points)
	if err != nil {
		app.errorLog.Println("Error:", err)
		app.serverError(w, err)
		return
	}

	app.infoLog.Println("\nPoints:", points)
	_, err = app.tasks.UpdatePointsById(id, points)
	if err != nil {
		app.errorLog.Println("Error:", err)
		app.serverError(w, err)
		return
	}

	app.infoLog.Println("Points were updated from task with id:", id)

	w.WriteHeader(http.StatusOK)
}

func (app *application) getActiveUserTasks(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	app.infoLog.Printf("Getting active tasks from user with id %s\n", id)

	tasks, err := app.tasks.GetActiveUserTasks(id)
	if err != nil {
		app.errorLog.Println("Error getting active tasks: ", err)
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

	app.infoLog.Println("\nActive tasks were sent")

	// Send response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(b)
}

func (app *application) getHistoryUserTasks(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	app.infoLog.Printf("Getting history tasks from user with id %s\n", id)

	tasks, err := app.tasks.GetHistoryUserTasks(id)
	if err != nil {
		app.errorLog.Println("Error getting history tasks: ", err)
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

	app.infoLog.Println("\nHistory tasks were sent")

	// Send response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(b)
}