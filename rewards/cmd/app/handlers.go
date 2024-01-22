package main

import (
	"encoding/json"
	"net/http"

	"github.com/cpprian/stucoin-backend/rewards/pkg/models"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (app *application) all(w http.ResponseWriter, r *http.Request) {
	// Get all rewards
	rewards, err := app.rewards.All()
	if err != nil {
		app.errorLog.Println("Error getting all rewards: ", err)
		app.serverError(w, err)
		return
	}

	// Convert reward list into json encoding
	b, err := json.Marshal(rewards)
	if err != nil {
		app.errorLog.Println("Error marshalling rewards: ", err)
		app.serverError(w, err)
		return
	}

	app.infoLog.Println("\nAll rewards were sent")

	// Send response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(b)
}

func (app *application) findById(w http.ResponseWriter, r *http.Request) {
	// Get reward id from request
	id := mux.Vars(r)["id"]
	app.infoLog.Printf("Getting reward with id %s\n", id)

	// Get reward
	reward, err := app.rewards.FindById(id)
	if err != nil {
		if err.Error() == "no reward found" {
			app.infoLog.Println("Reward not found")
			return
		}
		app.errorLog.Println("Error getting reward: ", err)
		app.serverError(w, err)
		return
	}
	app.infoLog.Println("\nReward:", reward)

	// Convert reward into json encoding
	b, err := json.Marshal(reward)
	if err != nil {
		app.errorLog.Println("Error marshalling reward: ", err)
		app.serverError(w, err)
		return
	}

	app.infoLog.Println("\nReward was sent")

	// Send response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(b)
}

func (app *application) findByName(w http.ResponseWriter, r *http.Request) {
	// Get name from request
	name := mux.Vars(r)["name"]
	app.infoLog.Printf("Getting reward with name %s\n", name)

	// Get reward
	reward, err := app.rewards.FindByName(name)
	if err != nil {
		if err.Error() == "no reward found" {
			app.infoLog.Println("Reward not found")
			return
		}
		app.errorLog.Println("Error getting reward: ", err)
		app.serverError(w, err)
		return
	}
	app.infoLog.Println("\nReward:", reward)

	// Convert reward into json encoding
	b, err := json.Marshal(reward)
	if err != nil {
		app.errorLog.Println("Error marshalling reward: ", err)
		app.serverError(w, err)
		return
	}

	app.infoLog.Println("\nReward was sent")

	// Send response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(b)
}

func (app *application) insertReward(w http.ResponseWriter, r *http.Request) {
	// Get reward from request
	var reward models.Reward
	err := json.NewDecoder(r.Body).Decode(&reward)
	if err != nil {
		app.errorLog.Println("Error:", err)
		app.serverError(w, err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	app.infoLog.Println("\nReward:", reward)

	// Insert reward
	resp, err := app.rewards.InsertReward(&reward)
	if err != nil {
		app.errorLog.Println("Error:", err)
		app.serverError(w, err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	app.infoLog.Println("Reward was inserted with data:", reward)

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

func (app *application) updateReward(w http.ResponseWriter, r *http.Request) {
	// Get reward from request
	var reward models.Reward
	err := json.NewDecoder(r.Body).Decode(&reward)
	if err != nil {
		app.errorLog.Println("Error:", err)
		app.serverError(w, err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	app.infoLog.Println("\nReward:", reward)

	// Update reward
	_, err = app.rewards.UpdateReward(&reward)
	if err != nil {
		app.errorLog.Println("Error:", err)
		app.serverError(w, err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	app.infoLog.Println("Reward was updated with id:", reward.ID)

	// Send response
	w.WriteHeader(http.StatusOK)
}