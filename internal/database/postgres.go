package database

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/Aidajy111/API-chats-and-messages/internal/config"
)

// NewConnect - соединение с БД
func NewConnect(cfg *config.Config) (*sql.DB, error) {
	dsn := cfg.PostgresDSN()
	log.Printf("Using database: %s", dsn)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("open db: %w", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("ping db: %w", err)
	}

	return db, nil
}
