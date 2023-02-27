package main

import (
	"fmt"

	"github.com/Elen0207/API-Rest-Golang/models"
	"github.com/Elen0207/API-Rest-Golang/routes"
)

func main() {
	models.Personalidades = []models.Personalidade{
		{Id: 1, Nome: "Nome 1", Historia: "Historia 1"},
		{Id: 2, Nome: "Nome 2", Historia: "Historia 2"},
	}
	fmt.Println("Initializing server...")
	routes.HandleResquest()
}