package main

import (
	"log"
	_ "todo-app/docs"
	"todo-app/internal/database"
	"todo-app/internal/routes"
)

// @title           TodoAppAPI
// @version         1.0
// @host           localhost:8080
// @BasePath       /
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
