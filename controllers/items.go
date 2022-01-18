package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/colombohenrique/rollin-api/database"
	"github.com/colombohenrique/rollin-api/models"
	"github.com/gorilla/mux"
)

func GetAllItems(w http.ResponseWriter, r *http.Request) {
	var items []models.Item
	database.DB.Find(&items)
	json.NewEncoder(w).Encode(items)
}

func GetItem(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var item models.Item
	database.DB.First(&item, id)
	json.NewEncoder(w).Encode(item)
}

func AddNewItem(w http.ResponseWriter, r *http.Request) {
	var item models.Item
	json.NewDecoder(r.Body).Decode(&item)
	database.DB.Create(&item)
	json.NewEncoder(w).Encode(item)
}

func DeleteItem(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var item models.Item
	database.DB.Delete(&item, id)
	json.NewEncoder(w).Encode(item)
}

func EditItem(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var item models.Item
	database.DB.First(&item, id)
	json.NewDecoder(r.Body).Decode(&item)
	database.DB.Save(&item)
	json.NewEncoder(w).Encode(item)
}
