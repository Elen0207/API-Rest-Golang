package main

import (
	"github.com/Elen0207/API-Rest-Golang/database"
	"github.com/Elen0207/API-Rest-Golang/routes"
)

func main() {
	database.ConectaComBancoDeDados()
	routes.HandleRequest()
}
