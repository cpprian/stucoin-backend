package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/cpprian/stucoin-backend/tasks/pkg/models"
	"github.com/gorilla/mux"
)

type TaskData struct {
	Task models.Task
}

func (app *application) createTask(w http.ResponseWriter, r *http.Request) {
	var task models.Task
	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		app.errorLog.Println("Error decoding task: ", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	app.infoLog.Printf("Creating task: %v\n", task)
	resp, err := app.postApiContent(app.apis.tasks, task)
	if err != nil {
		app.errorLog.Println("Error creating task: ", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()
	app.infoLog.Println("Task was created")

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		app.errorLog.Println("Error reading response body: ", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	app.infoLog.Println("Response: ", string(body))
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(body)
}

func (app *application) getTaskById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	app.infoLog.Println(vars)
	id, ok := vars["id"]
	if !ok {
		app.errorLog.Println("Error getting task id")
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	app.infoLog.Printf(("Getting task with id %s\n"), id)
	url := fmt.Sprintf("%s/%s", app.apis.tasks, id)
	app.getTask(w, r, url)
}

func (app *application) getTask(w http.ResponseWriter, r *http.Request, url string) {
	resp, err := app.getApiContent(url)
	if err != nil {
		app.errorLog.Println("Error getting task: ", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	var task models.Task
	err = json.NewDecoder(resp.Body).Decode(&task)
	if err != nil {
		app.errorLog.Println("Error decoding task: ", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	body, err := json.Marshal(task)
	if err != nil {
		app.errorLog.Println("Error marshalling tasks: ", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	app.infoLog.Println("Body to send: ", string(body))
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(body)
}

func (app *application) getAllTasks(w http.ResponseWriter, r *http.Request) {
	app.infoLog.Println("URL: ", app.apis.tasks)
	resp, err := app.getApiContent(app.apis.tasks)
	if err != nil {
		app.errorLog.Println("Error getting tasks: ", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	var tasks []models.Task
	err = json.NewDecoder(resp.Body).Decode(&tasks)
	if err != nil {
		app.errorLog.Println("Error decoding tasks: ", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	app.infoLog.Printf("Tasks: %+v\n", tasks)

	body, err := json.Marshal(tasks)
	if err != nil {
		app.errorLog.Println("Error marshalling tasks: ", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	app.infoLog.Println("Body to send: ", string(body))
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(body)
}

func (app *application) getAllTasksByOwnerId(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	ownerId, ok := vars["owner"]
	if !ok {
		app.errorLog.Println("Error getting task owner id")
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	app.infoLog.Println("URL: ", app.apis.tasks)
	url := fmt.Sprintf("%s/teacher/%s", app.apis.tasks, ownerId)
	resp, err := app.getApiContent(url)
	if err != nil {
		app.errorLog.Println("Error getting tasks: ", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close();

	var tasks []models.Task
	err = json.NewDecoder(resp.Body).Decode(&tasks)
	if err != nil {
		app.errorLog.Println("Error decoding tasks: ", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	app.infoLog.Printf("Tasks: %+v\n", tasks)

	body, err := json.Marshal(tasks)
	if err != nil {
		app.errorLog.Println("Error marshalling tasks: ", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	app.infoLog.Println("Body to send: ", string(body))
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(body)
}

func (app *application) updateTask(w http.ResponseWriter, r *http.Request) {
	var task models.Task
	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		app.errorLog.Println("Error decoding task: ", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	app.infoLog.Printf("Updating task: %v\n", task)
	url := fmt.Sprintf("%s/%d", app.apis.tasks, task.ID)
	err = app.putApiContent(url, task)
	if err != nil {
		app.errorLog.Println("Error updating task: ", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	app.infoLog.Printf("Task with id %v was updated\n", task.ID)
}

func (app *application) deleteTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		app.errorLog.Println("Error getting task id")
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	app.infoLog.Printf("Deleting task with id %s\n", id)
	url := fmt.Sprintf("%s/%s", app.apis.tasks, id)
	err := app.deleteApiContent(url, nil)
	if err != nil {
		app.errorLog.Println("Error deleting task: ", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	app.infoLog.Printf("Task with id %s was deleted\n", id)
}

func (app *application) updateCoverImageById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		app.errorLog.Println("Error getting task id")
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	var coverImage models.CoverImage
	err := json.NewDecoder(r.Body).Decode(&coverImage)
	if err != nil {
		app.errorLog.Println("Error decoding task: ", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	app.infoLog.Printf("Updating cover image for task with id %s\n", id)
	url := fmt.Sprintf("%s/cover/%s", app.apis.tasks, id)
	err = app.putApiContent(url, coverImage)
	if err != nil {
		app.errorLog.Println("Error updating cover image for task: ", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	app.infoLog.Printf("Cover image for task with id %s was updated\n", id)
}

func (app *application) updateContentById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		app.errorLog.Println("Error getting task id")
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	var content models.Content
	err := json.NewDecoder(r.Body).Decode(&content)
	if err != nil {
		app.errorLog.Println("Error decoding task: ", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	app.infoLog.Printf("Updating content for task with id %s\n", id)
	url := fmt.Sprintf("%s/content/%s", app.apis.tasks, id)
	err = app.putApiContent(url, content)
	if err != nil {
		app.errorLog.Println("Error updating content for task: ", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	app.infoLog.Printf("Content for task with id %s was updated\n", id)
}

func (app *application) updateTitleById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		app.errorLog.Println("Error getting task id")
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	var title models.Title
	err := json.NewDecoder(r.Body).Decode(&title)
	if err != nil {
		app.errorLog.Println("Error decoding task: ", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	app.infoLog.Printf("Updating title for task with id %s\n", id)
	url := fmt.Sprintf("%s/title/%s", app.apis.tasks, id)
	err = app.putApiContent(url, title)
	if err != nil {
		app.errorLog.Println("Error updating title for task: ", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	app.infoLog.Printf("Title for task with id %s was updated\n", id)
}

func (app *application) saveFilesById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		app.errorLog.Println("Error getting task id")
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	app.infoLog.Printf("Saving file: %v\n", r.Body)

	var files models.File
	err := json.NewDecoder(r.Body).Decode(&files)
	if err != nil {
		app.errorLog.Println("Error decoding task: ", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	app.infoLog.Printf("Saving files for task with id %s\n", id)
	url := fmt.Sprintf("%s/files/%s", app.apis.tasks, id)
	_, err = app.postApiContent(url, files)
	if err != nil {
		app.errorLog.Println("Error saving files for task: ", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	app.infoLog.Printf("Files for task with id %s were saved\n", id)
}

func (app *application) deleteFilesById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		app.errorLog.Println("Error getting task id")
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	app.infoLog.Printf("Deleting file: %v\n", r.Body)

	var file models.File
	err := json.NewDecoder(r.Body).Decode(&file)
	if err != nil {
		app.errorLog.Println("Error decoding task: ", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	app.infoLog.Printf("Deleting files for task with id %s\n", id)
	url := fmt.Sprintf("%s/files/%s", app.apis.tasks, id)
	err = app.deleteApiContent(url, file)
	if err != nil {
		app.errorLog.Println("Error deleting files for task: ", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	app.infoLog.Printf("Files for task with id %s were deleted\n", id)
}

func (app *application) assignTaskById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		app.errorLog.Println("Error getting task id")
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	app.infoLog.Printf("Assigning task: %v\n", r.Body)

	var inCharge models.InCharge
	err := json.NewDecoder(r.Body).Decode(&inCharge)
	if err != nil {
		app.errorLog.Println("Error decoding task: ", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	app.infoLog.Printf("Assigning task with id %s\n", id)
	url := fmt.Sprintf("%s/assign/%s", app.apis.tasks, id)
	err = app.putApiContent(url, inCharge)
	if err != nil {
		app.errorLog.Println("Error assigning task: ", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	app.infoLog.Printf("Task with id %s was assigned\n", id)
}

func (app *application) completeTaskById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		app.errorLog.Println("Error getting completed task id")
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	app.infoLog.Printf("Completing task: %v\n", r.Body)

	app.infoLog.Printf("Completing task with id %s\n", id)
	url := fmt.Sprintf("%s/complete/%s", app.apis.tasks, id)
	err := app.putApiContent(url, nil)
	if err != nil {
		app.errorLog.Println("Error completing task: ", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	app.infoLog.Printf("Task with id %s was completed\n", id)
}

func (app *application) acceptTaskById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		app.errorLog.Println("Error accepting task id")
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	app.infoLog.Printf("Accepting task: %v\n", r.Body)

	app.infoLog.Printf("Accepting task with id %s\n", id)
	url := fmt.Sprintf("%s/accept/%s", app.apis.tasks, id)
	err := app.putApiContent(url, nil)
	if err != nil {
		app.errorLog.Println("Error accepting task: ", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	app.infoLog.Printf("Task with id %s was accepted\n", id)
}

func (app *application) rejectTaskById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		app.errorLog.Println("Error rejecting task id")
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	app.infoLog.Printf("Rejecting task: %v\n", r.Body)

	app.infoLog.Printf("Rejecting task with id %s\n", id)
	url := fmt.Sprintf("%s/reject/%s", app.apis.tasks, id)
	err := app.putApiContent(url, nil)
	if err != nil {
		app.errorLog.Println("Error rejecting task: ", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	app.infoLog.Printf("Task with id %s was rejected\n", id)
}

func (app *application) updatePointsById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		app.errorLog.Println("Error updating task points id")
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	app.infoLog.Printf("Updating task points: %v\n", r.Body)

	var points models.Points
	err := json.NewDecoder(r.Body).Decode(&points)
	if err != nil {
		app.errorLog.Println("Error decoding task: ", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	app.infoLog.Printf("Updating points for task with id %s\n", id)
	url := fmt.Sprintf("%s/points/%s", app.apis.tasks, id)
	err = app.putApiContent(url, points)
	if err != nil {
		app.errorLog.Println("Error updating points for task: ", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	app.infoLog.Printf("Points for task with id %s were updated\n", id)
}

func (app *application) getUserActiveTasks(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		app.errorLog.Println("Error getting user id")
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	app.infoLog.Println("URL: ", app.apis.tasks)
	url := fmt.Sprintf("%s/active/%s", app.apis.tasks, id)
	resp, err := app.getApiContent(url)
	if err != nil {
		app.errorLog.Println("Error getting tasks: ", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close();

	var tasks []models.Task
	err = json.NewDecoder(resp.Body).Decode(&tasks)
	if err != nil {
		app.errorLog.Println("Error decoding tasks: ", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	app.infoLog.Printf("Tasks: %+v\n", tasks)

	body, err := json.Marshal(tasks)
	if err != nil {
		app.errorLog.Println("Error marshalling tasks: ", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	app.infoLog.Println("Body to send: ", string(body))
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(body)
}

func (app *application) getUserHistoryTasks(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		app.errorLog.Println("Error getting user id")
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	app.infoLog.Println("URL: ", app.apis.tasks)
	url := fmt.Sprintf("%s/history/%s", app.apis.tasks, id)
	resp, err := app.getApiContent(url)
	if err != nil {
		app.errorLog.Println("Error getting tasks: ", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close();

	var tasks []models.Task
	err = json.NewDecoder(resp.Body).Decode(&tasks)
	if err != nil {
		app.errorLog.Println("Error decoding tasks: ", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	app.infoLog.Printf("Tasks: %+v\n", tasks)

	body, err := json.Marshal(tasks)
	if err != nil {
		app.errorLog.Println("Error marshalling tasks: ", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	app.infoLog.Println("Body to send: ", string(body))
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(body)
}