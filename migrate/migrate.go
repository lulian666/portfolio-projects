package main

import (
	"log"
	"portfolio/initializers"
	"portfolio/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectDB()
}

func main() {
	err := initializers.DB.AutoMigrate(&models.Project{})
	if err != nil {
		log.Fatal("Error migrating database")
	}
	log.Println("Migration successful")
}
