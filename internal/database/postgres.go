package database

import (
	"fmt"
	"log"

	"github.com/Aidajy111/API-chats-and-messages/internal/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// NewConnect - соединение с БД
func NewConnect(cfg *config.Config) (*gorm.DB, error) {
	dsn := cfg.PostgresDSN()
	log.Printf("Using database: %s", dsn)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("open db: %w", err)
	}

	return db, nil
}
