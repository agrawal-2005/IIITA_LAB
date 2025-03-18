package controllers

import (
	"log"
	"net/http"

	"github.com/agrawal-2005/iiita_lab/backend/pkg/models"
	"github.com/agrawal-2005/iiita_lab/backend/pkg/config"
	"github.com/gin-gonic/gin"
)

func CreateProject(c *gin.Context) {
	var request struct {
		Topic       string `json:"topic" binding:"required"`
		ProjectLead string `json:"project_lead" binding:"required,email"`
		Faculty     string `json:"project_faculty" binding:"required,email"`
		Description string `json:"description" binding:"required"`
		Domain      string `json:"domain" binding:"required"`
		Year        int    `json:"year" binding:"required"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body", "details": err.Error()})
		return
	}

	db := config.DB

	project := models.Project{
		Title:           request.Topic,
		Description:     request.Description,
		ProjectLead:     request.ProjectLead,
		InstructorEmail: request.Faculty,
		Domain:          request.Domain,
		Year:            request.Year,
		Status:          "ongoing",
		Citations:       0,
	}

	if err := db.Create(&project).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create project"})
		return
	}

	projectID := project.ProjectID

	var person models.People
	if err := db.Where("email = ?", request.ProjectLead).First(&person).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Project lead not found in people table"})
		return
	}

	coauthor := models.Coauthor{
		PaperID:  projectID,
		Coauthor: request.ProjectLead,
		Name:     person.Name,
	}

	if err := db.Create(&coauthor).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add coauthor"})
		return
	}

	log.Println("✅ Project created successfully")

	c.JSON(http.StatusOK, gin.H{
		"message":    "Project created successfully",
		"project_id": projectID,
	})
}