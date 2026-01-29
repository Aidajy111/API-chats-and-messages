package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq" // для PostgreSQL
	"github.com/pressly/goose/v3"
)

func RunMigrations(db *sql.DB) error {
	// Устанавливаем диалект БД
	goose.SetDialect("postgres")

	// Запускаем миграции из папки migrations
	if err := goose.Up(db, "migrations"); err != nil {
		return fmt.Errorf("миграции не удались: %w", err)
	}

	return nil
}
