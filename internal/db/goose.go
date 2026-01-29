package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	"github.com/pressly/goose/v3"
)

// RunMigrations - запуск миграций
func RunMigrations(db *sql.DB) error {
	goose.SetDialect("postgres")

	if err := goose.Up(db, "migrations"); err != nil {
		return fmt.Errorf("миграции не удались: %w", err)
	}

	return nil
}
