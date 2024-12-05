package main

import (
	"go-orm-postgresql/config"
	"go-orm-postgresql/models"
	"go-orm-postgresql/repositories"
	"go-orm-postgresql/services"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// Initialize the database
	config.InitDB()
	db := config.DB

	// Run database migrations
	if err := db.AutoMigrate(&models.User{}); err != nil {
		log.Fatalf("Error running migrations: %v", err)
	}
	log.Println("Database migrated successfully.")

	// Set up repository and service
	userRepo := repositories.NewUserRepository(db)
	userService := services.NewUserService(userRepo)

	// Example usage
	user := &models.User{Name: "John Doe", Email: "john.doe@example.com", Password: "securepassword"}
	if err := userService.CreateUser(user); err != nil {
		log.Fatalf("Error creating user: %v", err)
	}

	users, err := userService.GetAllUsers()
	if err != nil {
		log.Fatalf("Error retrieving users: %v", err)
	}

	log.Printf("Retrieved users: %+v", users)
}
