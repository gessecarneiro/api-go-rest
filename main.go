package main

import (
	"fmt"

	"github.com/gessecarneiro/api-go-rest/models"
	"github.com/gessecarneiro/api-go-rest/routes"
)





func main() {

	models.Personalities =[]models.Personality {
		{Id: 1,Name: "Name 1", History: "History 1"},
		{Id: 2,Name: "Name 2", History: "History 2"},
	}

	fmt.Println("Server on with Go")
	routes.HandleRequest()

}