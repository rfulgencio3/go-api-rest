package main

import (
	"fmt"

	"github.com/rfulgencio3/go-api-rest/routes"
)

func main() {
	fmt.Println("Iniciando o Servidor Rest com Go")
	routes.HandleRequest()
}
