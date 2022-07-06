package routes

import (
	"log"
	"net/http"

	"github.com/gessecarneiro/api-go-rest/controllers"
	"github.com/gorilla/mux"
)

func HandleRequest() {
	r := mux.NewRouter()
	r.HandleFunc("/", controllers.Home)
	r.HandleFunc("/api/personalities", controllers.AllPersonalities).Methods("Get")
	r.HandleFunc("/api/personalities/{id}", controllers.ReturnPersonality).Methods("Get")
	log.Fatal(http.ListenAndServe(":8000", r))
}
