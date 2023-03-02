package main

import (
	"fmt"

	"github.com/Elen0207/API-Rest-Golang/database"
	"github.com/Elen0207/API-Rest-Golang/routes"
)

func main() {
	database.ConectaComBancoDeDados()
	fmt.Println("Initializing server...")
	routes.HandleResquest()
}