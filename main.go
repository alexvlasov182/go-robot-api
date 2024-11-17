package main

import (
	"context"
	"go-robot-api/controllers"
	"go-robot-api/routes"
	"log"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Connect to PostgreSQL
	dbURL := os.Getenv("DATABASE_URL")
	config, err := pgxpool.ParseConfig(dbURL)
	if err != nil {
		log.Fatal("Failed to parse database URL:", err)
	}

	db, err := pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
		log.Fatal("Failed to connect to database", err)
	}
	defer db.Close()

	// Run migrations
	controllers.RunMigrations(db)

	// Initialize database in controllers
	controllers.InitializeDB(db)

	// Create Gin router
	router := gin.Default()
	router.Use(cors.Default())
	routes.RegisterRobotRoutes(router)

	// Start the server
	router.Run(":3000")
}
