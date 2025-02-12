package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"github.com/agrawal-2005/iiita_lab/pkg/models"
)

// PublicationController struct
type PublicationController struct {
	DB *gorm.DB
}

// CreatePublication - Add a new publication
func (pc *PublicationController) CreatePublication(c *gin.Context) {
	var publication models.Publication
	if err := c.ShouldBindJSON(&publication); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := pc.DB.Create(&publication).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create publication"})
		return
	}
	c.JSON(http.StatusCreated, publication)
}

// GetPublications - Retrieve all publications
func (pc *PublicationController) GetPublications(c *gin.Context) {
	var publications []models.Publication
	pc.DB.Find(&publications)
	c.JSON(http.StatusOK, publications)
}

// GetPublicationByID - Retrieve a single publication by ID
func (pc *PublicationController) GetPublicationByID(c *gin.Context) {
	var publication models.Publication
	id := c.Param("id")
	if err := pc.DB.First(&publication, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Publication not found"})
		return
	}
	c.JSON(http.StatusOK, publication)
}

// UpdatePublication - Update a publication by ID
func (pc *PublicationController) UpdatePublication(c *gin.Context) {
	var publication models.Publication
	id := c.Param("id")

	if err := pc.DB.First(&publication, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Publication not found"})
		return
	}

	if err := c.ShouldBindJSON(&publication); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	pc.DB.Save(&publication)
	c.JSON(http.StatusOK, publication)
}

// DeletePublication - Delete a publication by ID
func (pc *PublicationController) DeletePublication(c *gin.Context) {
	id := c.Param("id")
	if err := pc.DB.Delete(&models.Publication{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete publication"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Publication deleted successfully"})
}
