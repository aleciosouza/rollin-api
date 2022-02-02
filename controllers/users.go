package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/colombohenrique/rollin-api/database"
	"github.com/colombohenrique/rollin-api/models"
	"github.com/gorilla/mux"
)

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	var users []models.User
	database.DB.Find(&users)
	json.NewEncoder(w).Encode(users)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var user models.User
	database.DB.First(&user, id)
	// rows, err := database.DB.Table("users").Select("*").Joins("left join persons on persons.id = users.person_id").Where("users.id = " + id).Rows()

	// if err != nil {
	// 	fmt.Println("batata error: ", err)
	// }

	// for rows.Next() {
	// 	fmt.Println(rows.Scan(user))
	// 	rows.Scan(&user)
	// 	fmt.Println(user)
	// 	rows.Scan(&person)
	// 	fmt.Println(person)
	// }
	json.NewEncoder(w).Encode(user)
}

func AddNewUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	json.NewDecoder(r.Body).Decode(&user)
	database.DB.Create(&user)
	json.NewEncoder(w).Encode(user)
}

func EditUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var user models.User
	database.DB.First(&user, id)
	json.NewDecoder(r.Body).Decode(&user)
	database.DB.Save(&user)
	json.NewEncoder(w).Encode(user)
}
