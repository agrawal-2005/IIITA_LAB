package routes

import (
	"github.com/agrawal-2005/iiita_lab/backend/pkg/controllers"
	"github.com/gin-gonic/gin"
)


func SetupProjectRoutes(router *gin.Engine) {
	projectRoutes := router.Group("/projects")
	{
		projectRoutes.POST("/create", controllers.CreateProject)
	}
}
