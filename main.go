package main

import (
	"fmt"

	"github.com/rfulgencio3/go-api-rest/routes"
)

func main() {
	fmt.Println("Starting API REST using Go")
	routes.HandleRequest()
}
