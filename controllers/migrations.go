package controllers

import (
	"context"
	"log"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

var db *pgxpool.Pool

func InitializeDB(database *pgxpool.Pool) {
	db = database
}

func RunMigrations(db *pgxpool.Pool) {
	migrationFiles := "./migrations/001_create_robots_table.sql"

	// Read migration file
	migrationQuery, err := os.ReadFile(migrationFiles)
	if err != nil {
		log.Fatalf("Failed to read migration file %v", err)
	}

	// Execute migration query
	_, err = db.Exec(context.Background(), string(migrationQuery))
	if err != nil {
		log.Fatalf("Failed to apply migration: %v", err)
	}

	log.Println("Database migration applied successfully")
}
