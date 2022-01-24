package routes

import (
	"log"
	"net/http"

	"github.com/rfulgencio3/go-api-rest/controllers"
)

func HandleRequest() {
	http.HandleFunc("/", controllers.Home)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
