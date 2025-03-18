package routes

import (
	"github.com/agrawal-2005/iiita_lab/backend/pkg/controllers"
	"github.com/agrawal-2005/iiita_lab/backend/pkg/middlewares"
	"github.com/gin-gonic/gin"
)

// ProfileRoutes sets up routes for user profiles
func SetupProfileRoutes(router *gin.Engine) {
	profileGroup := router.Group("/profile")
	profileGroup.Use(middleware.AuthMiddleware())
	{
		profileGroup.GET("/", controllers.GetProfile)
	}
}
