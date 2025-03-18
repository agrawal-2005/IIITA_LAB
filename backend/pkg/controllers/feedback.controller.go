package controllers

import (
	"net/http"

	"github.com/agrawal-2005/iiita_lab/backend/pkg/config"
	"github.com/agrawal-2005/iiita_lab/backend/pkg/models"
	"github.com/gin-gonic/gin"
)

// SubmitFeedback handles the feedback form submission
func SubmitFeedback(c *gin.Context) {
	var feedback models.Feedback

	// Bind JSON request to feedback struct
	if err := c.ShouldBindJSON(&feedback); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	db := config.DB

	// Insert feedback into the database
	if err := db.Create(&feedback).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to submit feedback"})
		return
	}

	// Redirect to homepage
	c.File("./views/homepage.html")
}

// GetConferences retrieves all conferences from the database
func GetConferences(c *gin.Context) {
	db := config.DB
	var conferences []models.Conference

	// Fetch conferences ordered by date descending
	if err := db.Order("date DESC").Find(&conferences).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch conferences"})
		return
	}

	// Render conferences page
	c.HTML(http.StatusOK, "conferences.ejs", gin.H{"conferences": conferences})
}
