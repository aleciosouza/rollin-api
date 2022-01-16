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
	//r.HandleFunc("/", controllers.Home)

	//Events
	r.HandleFunc("/api/event", controllers.GetAllEvents).Methods("Get") //Return all events
	r.HandleFunc("/api/event/{id}", controllers.GetEvent).Methods("Get")
	r.HandleFunc("/api/event", controllers.AddNewEvent).Methods("Post")
	r.HandleFunc("/api/event/{id}", controllers.DeleteEvent).Methods("Delete")
	r.HandleFunc("/api/event/{id}", controllers.EditEvent).Methods("Put")

	//Items

	log.Fatal(http.ListenAndServe(":8000", handlers.CORS(handlers.AllowedOrigins([]string{"*"}))(r)))
}
