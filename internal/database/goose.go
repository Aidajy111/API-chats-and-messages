package database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	"github.com/pressly/goose/v3"
)

// RunMigrations - запуск миграций. В проде миграции не запускаем
func RunMigrations(db *sql.DB, env string) error {
	if env == "production" {
		return nil
	}

	goose.SetDialect("postgres")

	if err := goose.Up(db, "./migrations"); err != nil {
		return fmt.Errorf("migrations failed: %w", err)
	}
	return nil
}
