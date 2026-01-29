package main

import (
	"log"
	"os"

	"github.com/Aidajy111/API-chats-and-messages/internal/config"
	"github.com/Aidajy111/API-chats-and-messages/internal/database"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using system environment variables")
	}

	cfg := config.Load()
	log.Printf("Starting in %s mode on port %s", cfg.Env, cfg.Port)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8512"
	}

	// Connect db
	db, err := database.NewConnect(cfg)
	if err != nil {
		log.Fatal("Database connection failed: ", err)
	}
	defer db.Close()

	// Migrations
	if err := database.RunMigrations(db, cfg.Env); err != nil {
		log.Fatal("Migrations error: ", err)
	}

	log.Println("Running database migrations...")
}
