package routes

import (
	"log"
	"net/http"

	"github.com/colombohenrique/rollin-api/controllers"
	"github.com/colombohenrique/rollin-api/middleware"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func HandleRequest() {
	r := mux.NewRouter()
	r.Use(middleware.ContentTypeMiddleware)

	//Home
	r.HandleFunc("/", controllers.Home) //Page containing API documentation

	//Events
	r.HandleFunc("/api/event", controllers.GetAllEvents).Methods("Get") //Return all events
	r.HandleFunc("/api/event/{id}", controllers.GetEvent).Methods("Get")
	r.HandleFunc("/api/event", controllers.AddNewEvent).Methods("Post")
	r.HandleFunc("/api/event/{id}", controllers.DeleteEvent).Methods("Delete")
	r.HandleFunc("/api/event/{id}", controllers.EditEvent).Methods("Put")

	//Items
	r.HandleFunc("/api/item", controllers.GetAllItems).Methods("Get")
	r.HandleFunc("/api/item/{id}", controllers.GetItem).Methods("Get")
	r.HandleFunc("/api/item", controllers.AddNewItem).Methods("Post")
	r.HandleFunc("/api/item/{id}", controllers.DeleteItem).Methods("Delete")
	r.HandleFunc("/api/item/{id}", controllers.EditItem).Methods("Put")

	//Users
	r.HandleFunc("/api/user", controllers.GetAllUsers).Methods("Get")
	r.HandleFunc("/api/user/{id}", controllers.GetUser).Methods("Get")
	r.HandleFunc("/api/user", controllers.AddNewUser).Methods("Post")
	r.HandleFunc("/api/user/{id}", controllers.EditUser).Methods("Put")

	log.Fatal(http.ListenAndServe(":8000", handlers.CORS(handlers.AllowedOrigins([]string{"*"}))(r)))
}
