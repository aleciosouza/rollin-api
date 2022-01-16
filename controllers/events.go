package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/colombohenrique/rollin-api/database"
	"github.com/colombohenrique/rollin-api/models"
	"github.com/gorilla/mux"
)

func GetAllEvents(w http.ResponseWriter, r *http.Request) {
	var events []models.Events
	database.DB.Find(&events)
	json.NewEncoder(w).Encode(events)
}

func GetEvent(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var event models.Event
	database.DB.First(&event, id)
	json.NewEncoder(w).Encode(event)
}

func AddNewEvent(w http.ResponseWriter, r *http.Request) {
	var event models.Event
	json.NewDecoder(r.Body).Decode(&event)
	database.DB.Create(&event)
	json.NewEncoder(w).Encode(event)
}

func DeleteEvent(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var event models.Event
	database.DB.Delete(&event, id)
	json.NewEncoder(w).Encode(event)
}

func EditEvent(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var event models.Event
	database.DB.First(&event, id)
	json.NewDecoder(r.Body).Decode(&event)
	database.DB.Save(&event)
	json.NewEncoder(w).Encode(event)
}
