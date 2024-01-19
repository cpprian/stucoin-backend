package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/cpprian/stucoin-backend/rewards/pkg/models"
	"github.com/gorilla/mux"
)

type RewardData struct {
	Reward models.Reward
}

func (app *application) createReward(w http.ResponseWriter, r *http.Request) {
	var reward models.Reward
	err := json.NewDecoder(r.Body).Decode(&reward)
	if err != nil {
		app.errorLog.Println("Error decoding reward: ", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	app.infoLog.Printf("Creating reward: %v\n", reward)
	_, err = app.postApiContent(app.apis.rewards, reward)
	if err != nil {
		app.errorLog.Println("Error creating reward: ", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	app.infoLog.Println("Reward was created")
}

func (app *application) getRewardById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	app.infoLog.Println(vars)
	id, ok := vars["id"]
	if !ok {
		app.errorLog.Println("Error getting reward id")
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	app.infoLog.Printf(("Getting reward with id %s\n"), id)
	url := fmt.Sprintf("%s/%s", app.apis.rewards, id)
	app.getReward(w, r, url)
}

func (app *application) getRewardByName(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name, ok := vars["name"]
	if !ok {
		app.errorLog.Println("Error getting reward name")
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	app.infoLog.Printf(("Getting reward with name %s\n"), name)
	url := fmt.Sprintf("%s/%s", app.apis.rewards, name)
	app.getReward(w, r, url)
}

func (app *application) getReward(w http.ResponseWriter, r *http.Request, url string) {
	resp, err := app.getApiContent(url)
	if err != nil {
		app.errorLog.Println("Error getting reward: ", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	var reward models.Reward
	err = json.NewDecoder(resp.Body).Decode(&reward)
	if err != nil {
		app.errorLog.Println("Error decoding reward: ", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	app.infoLog.Printf("Reward: %v\n", reward)
}

func (app *application) getAllRewards(w http.ResponseWriter, r *http.Request) {
	app.infoLog.Println("URL: ", app.apis.rewards)
	resp, err := app.getApiContent(app.apis.rewards)
	if err != nil {
		app.errorLog.Println("Error getting rewards: ", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	var rewards []models.Reward
	err = json.NewDecoder(resp.Body).Decode(&rewards)
	if err != nil {
		app.errorLog.Println("Error decoding rewards: ", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	app.infoLog.Printf("Rewards: %v\n", rewards)
}

func (app *application) updateReward(w http.ResponseWriter, r *http.Request) {
	var reward models.Reward
	err := json.NewDecoder(r.Body).Decode(&reward)
	if err != nil {
		app.errorLog.Println("Error decoding reward: ", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	app.infoLog.Printf("Updating reward: %v\n", reward)
	url := fmt.Sprintf("%s/%d", app.apis.rewards, reward.ID)
	err = app.putApiContent(url, reward)
	if err != nil {
		app.errorLog.Println("Error updating reward: ", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	app.infoLog.Printf("Reward with id %d was updated\n", reward.ID)
}

func (app *application) deleteReward(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		app.errorLog.Println("Error getting reward id")
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	app.infoLog.Printf("Deleting reward with id %s\n", id)
	url := fmt.Sprintf("%s/%s", app.apis.rewards, id)
	err := app.deleteApiContent(url)
	if err != nil {
		app.errorLog.Println("Error deleting reward: ", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	app.infoLog.Printf("Reward with id %s was deleted\n", id)
}