package routes

import (
	"Basketball_academy/controllers"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRoutes(r *gin.Engine, db *gorm.DB) {
	api := r.Group("/api")
	{

		userController := controllers.NewUserController(db)
		api.GET("/users", userController.GetUsers)
		api.POST("/users", userController.CreateUser)

	}
}
