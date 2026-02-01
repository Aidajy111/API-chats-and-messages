package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Aidajy111/API-chats-and-messages/internal/config"
	"github.com/Aidajy111/API-chats-and-messages/internal/database"
	"github.com/Aidajy111/API-chats-and-messages/internal/handlers"
	"github.com/Aidajy111/API-chats-and-messages/internal/repo"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using system environment variables")
	}

	cfg := config.Load()
	log.Printf("Starting in %s mode on port %s", cfg.Env, cfg.Port)

	// Connect db
	db, err := database.NewConnect(cfg)
	if err != nil {
		log.Fatal("Database connection failed: ", err)
	}

	// Migrations
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Running database migrations...")
	if err := database.RunMigrations(sqlDB, cfg.Env); err != nil {
		log.Fatal("Migrations error: ", err)
	}

	// Server
	chatRepo := repo.NewChatRepository(db)
	messageRepo := repo.NewMessagesRepository(db)
	handlers := handlers.NewChatHandler(chatRepo, messageRepo)

	r := mux.NewRouter()
	r.HandleFunc("/chats", handlers.Create).Methods("POST")
	r.HandleFunc("/chats/{id}/messages", handlers.CreateMessage).Methods("POST")
	r.HandleFunc("/chats/{id}", handlers.Get).Methods("GET")
	r.HandleFunc("/chats/{id}", handlers.Delete).Methods("DELETE")

	fmt.Printf("Сервер запущен на http://localhost:%s", cfg.Port)
	if err := http.ListenAndServe(":"+cfg.Port, r); err != nil {
		log.Fatal("Server failed:", err)
	}
}
