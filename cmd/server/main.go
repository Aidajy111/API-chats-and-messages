package main

import (
	"database/sql"
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

	dsn := cfg.PostgresDSN()
	log.Printf("Using database: %s", dsn)

	// Connect db
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatal("Database connection failed: ", err)
	}

	// Migrations
	if err := database.RunMigrations(db); err != nil {
		log.Fatal("Migrations error: ", err)
	}

	log.Println("Migrations completed successfully")
}
