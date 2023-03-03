package main

import (
	"github.com/elencdeo/api-go-gin/database"
	"github.com/elencdeo/api-go-gin/routes"
)

func main() {
	database.ConectaComBancoDeDados()
	routes.HandleRequests()
}
