package controllers

import (
	"context"
	"go-robot-api/models"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// CreateRobot handles the creation of a new robot
func CreateRobot(c *gin.Context) {
	var robot models.Robot

	if err := c.ShouldBindJSON(&robot); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload", "details": err.Error()})
		return
	}

	robot.LastTested = time.Now()

	query := `INSERT INTO robots (name, status, last_tested, test_results) VALUES ($1, $2, $3, $4) RETURNING id`
	err := db.QueryRow(context.Background(), query, robot.Name, robot.Status, robot.LastTested, robot.TestResults).Scan(&robot.ID)
	if err != nil {
		log.Printf("Failed to create robot: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create robot"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Robot created successfully", "data": robot})
}

// GetAllRobots retrieves all robots from the database
func GetAllRobots(c *gin.Context) {
	rows, err := db.Query(context.Background(), "SELECT id, name, status, last_tested, test_results FROM robots")
	if err != nil {
		log.Printf("Failed to query robots: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve robots"})
		return
	}

	defer rows.Close()

	var robots []models.Robot
	for rows.Next() {
		var robot models.Robot
		err := rows.Scan(&robot.ID, &robot.Name, &robot.Status, &robot.LastTested, &robot.TestResults)
		if err != nil {
			log.Printf("Error scanning robot: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse robot data"})
			return
		}
		robots = append(robots, robot)
	}
	c.JSON(http.StatusOK, gin.H{"data": robots})
}
