package routes

import (
	"github.com/agrawal-2005/iiita_lab/backend/pkg/controllers"
	"github.com/gin-gonic/gin"
)

func SetupAuthRoutes(router *gin.Engine){
	auth := router.Group("/auth")
	{
		auth.POST("/login", controllers.LoginUser)
		auth.POST("/register", controllers.Register)
		auth.POST("/logout", controllers.Logout)
	}
}