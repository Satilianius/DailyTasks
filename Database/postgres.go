package Database

import (
	"DailyTasks/config"
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib"
)

// NewConnection Establishes a connection pool to the PostgreSQL database.
func NewConnection(cfg *config.Config) (*sql.DB, error) {
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		cfg.Db.Host,
		cfg.Db.Port,
		cfg.Db.User,
		cfg.Db.Password,
		cfg.Db.Name,
		cfg.Db.SSLMode,
	)

	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to open database connection: %w", err)
	}

	db.SetMaxOpenConns(cfg.Db.MaxOpenConnections)
	db.SetMaxIdleConns(cfg.Db.MaxIdleConnections)
	db.SetConnMaxLifetime(time.Duration(cfg.Db.ConnMaxLifetimeMinutes) * time.Minute)

	err = db.Ping()
	if err != nil {
		if err := db.Close(); err != nil {
			log.Printf("Error closing database connection: %v", err)
		} else {
			log.Println("Database connection closed.")
		}
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	log.Println("Database connection pool established successfully.")
	return db, nil
}
