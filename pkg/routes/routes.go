package routes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"github.com/agrawal-2005/iiita_lab/pkg/controllers"
)

func SetupRoutes(router *gin.Engine, db *gorm.DB) {
	// Initialize controllers with DB instance
	publicationController := controllers.PublicationController{DB: db}
	feedbackController := controllers.FeedbackController{DB: db}
	authController := controllers.AuthController{DB: db}

	// Authentication routes
	auth := router.Group("/auth")
	{
		auth.POST("/login", authController.Login)
		auth.POST("/register", authController.Register)
		auth.POST("/logout", authController.Logout)
	}

	// Publications routes
	publications := router.Group("/publications")
	{
		publications.POST("/", publicationController.CreatePublication)
		publications.GET("/", publicationController.GetPublications)
		publications.GET("/:id", publicationController.GetPublicationByID)
		publications.PUT("/:id", publicationController.UpdatePublication)
		publications.DELETE("/:id", publicationController.DeletePublication)
	}

	// Feedback routes
	feedback := router.Group("/feedback")
	{
		feedback.POST("/", feedbackController.SubmitFeedback)
		feedback.GET("/", feedbackController.GetAllFeedback)
		feedback.GET("/:id", feedbackController.GetFeedbackByID)
		feedback.DELETE("/:id", feedbackController.DeleteFeedback)
	}
}
