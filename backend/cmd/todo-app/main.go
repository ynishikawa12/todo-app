package main

import (
	"log"
	"todo-app/internal/database"
	"todo-app/internal/routes"
)

func main() {
	err := database.Connect()
	if err != nil {
		log.Fatal("Failed to connect to MySQL database")
	}

	r := routes.SetupRouter()
	if err := r.Run(":8080"); err != nil {
		log.Fatal("Failed to run server")
	}
}
