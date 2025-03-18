package routes

import (
	"github.com/agrawal-2005/iiita_lab/backend/pkg/controllers"
	"github.com/gin-gonic/gin"
)

// SetupRoutes registers all API routes
func SetupFeedbackRoutes(router *gin.Engine) {
	// Serve static files (e.g., HTML files)
	router.Static("/views", "./views")

	// Routes for feedback
	router.POST("/feedback", controllers.SubmitFeedback)

	// Routes for conferences
	router.GET("/conferences.ejs", controllers.GetConferences)
}
