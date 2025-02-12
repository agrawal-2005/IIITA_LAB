package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"github.com/agrawal-2005/iiita_lab/pkg/models"
)

// FeedbackController struct
type FeedbackController struct {
	DB *gorm.DB
}

// SubmitFeedback - Add new feedback
func (fc *FeedbackController) SubmitFeedback(c *gin.Context) {
	var feedback models.Feedback
	if err := c.ShouldBindJSON(&feedback); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := fc.DB.Create(&feedback).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to submit feedback"})
		return
	}
	c.JSON(http.StatusCreated, feedback)
}

// GetAllFeedback - Retrieve all feedbacks
func (fc *FeedbackController) GetAllFeedback(c *gin.Context) {
	var feedbacks []models.Feedback
	fc.DB.Find(&feedbacks)
	c.JSON(http.StatusOK, feedbacks)
}

// GetFeedbackByID - Retrieve feedback by ID
func (fc *FeedbackController) GetFeedbackByID(c *gin.Context) {
	var feedback models.Feedback
	id := c.Param("id")
	if err := fc.DB.First(&feedback, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Feedback not found"})
		return
	}
	c.JSON(http.StatusOK, feedback)
}

// DeleteFeedback - Delete feedback by ID
func (fc *FeedbackController) DeleteFeedback(c *gin.Context) {
	id := c.Param("id")
	if err := fc.DB.Delete(&models.Feedback{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete feedback"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Feedback deleted successfully"})
}