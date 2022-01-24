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
	r.HandleFunc("/api/personalities", controllers.Home)
	log.Fatal(http.ListenAndServe(":8000", r))
}
