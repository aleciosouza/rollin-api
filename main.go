package main

import (
	"fmt"

	"github.com/colombohenrique/rollin-api/routes"
)

func main() {
	//database.ConectaComBancoDeDados()
	fmt.Println("Iniciando API")
	routes.HandleRequest()
}
