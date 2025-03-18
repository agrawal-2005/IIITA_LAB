package controllers

import (
	"net/http"
	"github.com/agrawal-2005/iiita_lab/backend/pkg/config"
	"github.com/agrawal-2005/iiita_lab/backend/pkg/models"
	"github.com/gin-gonic/gin"
)

func getProfile(c *gin.Context){
	email, exists := c.Get("email")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	userType, _ := c.Get("user_type")
	var userProfile map[string]interface{}

	if userType == "Student" {
		var student models.Student
		if err := config.DB.Where("Email = ?", email).First(&student).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Student profile not found"})
			return
		}
		userProfile = map[string]interface{}{
			"name":        student.Name,
			"email":       student.Email,
			"roll_number": student.RollNumber,
			"department":  student.Department,
			"batch":       student.Batch,
			"supervisor":  student.Supervisor,
			"research_id": student.ResearchID,
			"project_id":  student.ProjectID,
			"conference":  student.Conference,
			"publication": student.Publication,
		}
	} else if userType == "Supervisor" {
		var supervisor models.Supervisor
		if err := config.DB.Where("email = ?", email).First(&supervisor).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Supervisor profile not found"})
			return
		}
		userProfile = map[string]interface{}{
			"name":        supervisor.Name,
			"email":       supervisor.Email,
			"department":  supervisor.Department,
			"research_id": supervisor.ResearchID,
			"project_id":  supervisor.ProjectID,
		}
	} else {
		var person models.People
		if err := config.DB.Where("email = ?", email).First(&person).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "User profile not found"})
			return
		}
		userProfile = map[string]interface{}{
			"name":       person.Name,
			"email":      person.Email,
			"department": person.Department,
			"user_type":  person.User_Type,
			"domain":     person.Domain,
		}
	}

	c.JSON(http.StatusOK, gin.H{"profile": userProfile})
}