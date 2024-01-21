package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/cpprian/stucoin-backend/micro-competencies/pkg/models"
	"github.com/gorilla/mux"
)

type MicroCompetenceData struct {
	MicroCompetence models.MicroCompetence
}

func (app *application) createMicroCompetency(w http.ResponseWriter, r *http.Request) {
	var microCompetence models.MicroCompetence
	err := json.NewDecoder(r.Body).Decode(&microCompetence)
	if err != nil {
		app.errorLog.Println("Error decoding micro competence: ", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	app.infoLog.Printf("Creating micro competence: %v\n", microCompetence)
	_, err = app.postApiContent(app.apis.microCompetencies, microCompetence)
	if err != nil {
		app.errorLog.Println("Error creating micro competence: ", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	app.infoLog.Println("Micro competence was created")
}

func (app *application) getMicroCompetencyById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	app.infoLog.Println(vars)
	id, ok := vars["id"]
	if !ok {
		app.errorLog.Println("Error getting micro competence id")
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	app.infoLog.Printf(("Getting micro competence with id %s\n"), id)
	url := fmt.Sprintf("%s/%s", app.apis.microCompetencies, id)
	app.getMicroCompetence(w, r, url)
}

func (app *application) getMicroCompetencyByName(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name, ok := vars["name"]
	if !ok {
		app.errorLog.Println("Error getting micro competence name")
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	app.infoLog.Printf(("Getting micro competence with name %s\n"), name)
	url := fmt.Sprintf("%s/%s", app.apis.microCompetencies, name)
	app.getMicroCompetence(w, r, url)
}

func (app *application) getMicroCompetence(w http.ResponseWriter, r *http.Request, url string) {
	resp, err := app.getApiContent(url)
	if err != nil {
		app.errorLog.Println("Error getting micro competence: ", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	var microCompetence models.MicroCompetence
	err = json.NewDecoder(resp.Body).Decode(&microCompetence)
	if err != nil {
		app.errorLog.Println("Error decoding micro competence: ", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	app.infoLog.Printf("Micro competence: %v\n", microCompetence)
}

func (app *application) getAllMicroCompetencies(w http.ResponseWriter, r *http.Request) {
	app.infoLog.Println("URL: ", app.apis.microCompetencies)
	resp, err := app.getApiContent(app.apis.microCompetencies)
	if err != nil {
		app.errorLog.Println("Error getting micro competencies: ", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	var microCompetencies []models.MicroCompetence
	err = json.NewDecoder(resp.Body).Decode(&microCompetencies)
	if err != nil {
		app.errorLog.Println("Error decoding micro competencies: ", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	app.infoLog.Printf("Micro competencies: %v\n", microCompetencies)
}

func (app *application) updateMicroCompetency(w http.ResponseWriter, r *http.Request) {
	var microCompetence models.MicroCompetence
	err := json.NewDecoder(r.Body).Decode(&microCompetence)
	if err != nil {
		app.errorLog.Println("Error decoding micro competence: ", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	app.infoLog.Printf("Updating micro competence: %v\n", microCompetence)
	url := fmt.Sprintf("%s/%d", app.apis.microCompetencies, microCompetence.ID)
	err = app.putApiContent(url, microCompetence)
	if err != nil {
		app.errorLog.Println("Error updating micro competence: ", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	app.infoLog.Printf("MicroCompetence with id %d was updated\n", microCompetence.ID)
}

func (app *application) deleteMicroCompetency(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		app.errorLog.Println("Error getting micro competence id")
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	app.infoLog.Printf("Deleting micro competence with id %s\n", id)
	url := fmt.Sprintf("%s/%s", app.apis.microCompetencies, id)
	err := app.deleteApiContent(url, nil)
	if err != nil {
		app.errorLog.Println("Error deleting micro competence: ", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	app.infoLog.Printf("MicroCompetence with id %s was deleted\n", id)
}