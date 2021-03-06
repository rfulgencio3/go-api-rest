package routes

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rfulgencio3/go-api-rest/controllers"
)

func HandleRequest() {
	r := mux.NewRouter()
	r.HandleFunc("/", controllers.Home)
	r.HandleFunc("/api/personalities", controllers.GetAllPersonalities).Methods("Get")
	r.HandleFunc("/api/personalities/{id}", controllers.GetPersonalityById).Methods("Get")
	log.Fatal(http.ListenAndServe(":8000", r))
}
