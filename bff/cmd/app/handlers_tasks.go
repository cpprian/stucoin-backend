package main

import (
	"encoding/json"
	"fmt"
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

func (app *application) getTaskByTitle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	title, ok := vars["title"]
	if !ok {
		app.errorLog.Println("Error getting task title")
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	app.infoLog.Printf(("Getting task with title %s\n"), title)
	url := fmt.Sprintf("%s/%s", app.apis.tasks, title)
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

	app.infoLog.Printf("Task: %v\n", task)
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

	app.infoLog.Printf("Tasks: %v\n", tasks)
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
}