package Database

import (
	"DailyTasks/config"
	"database/sql"
	"errors"
	"fmt"
	"log"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

func InitDb(cfg *config.Config) {
	db, err := getDbConnection(cfg)
	defer db.Close()

	// Create a new migration instance
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		log.Fatalf("Could not create migration driver: %v", err)
	}

	// Point to migration files
	m, err := migrate.NewWithDatabaseInstance(
		"file://Database/migrations", // Path to migration files
		"DailyTasks",                 // Database type
		driver,
	)
	if err != nil {
		log.Fatalf("Migration error: %v", err)
	}

	// Run all up migrations
	if err := m.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		log.Fatalf("Failed to run migrations: %v", err)
	}

	fmt.Println("Migrations completed successfully!")
}

func getDbConnection(cfg *config.Config) (*sql.DB, error) {
	// Get connection details from environment variables
	connString := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		cfg.Db.User,
		cfg.Db.Password,
		cfg.Db.Host,
		cfg.Db.Port,
		cfg.Db.Name)

	// Connect to the Database
	db, err := sql.Open("postgres", connString)
	if err != nil {
		log.Fatalf("Could not connect to Database: %v", err)
	}
	return db, err
}
