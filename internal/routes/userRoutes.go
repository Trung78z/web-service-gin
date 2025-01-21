package routes

import (
	"github/Trung78z/web-service-gin/internal/controllers"
	"github/Trung78z/web-service-gin/internal/repositories"
	"github/Trung78z/web-service-gin/internal/services"

	"github.com/gin-gonic/gin"
)

func addUserRoutes(rg *gin.RouterGroup) {
	userGroup := rg.Group("/users")
	UserService := services.NewUserService(repositories.Queries)
	UserController := controllers.NewUserController(UserService)
	userGroup.POST("/", UserController.CreateUser)
	userGroup.GET("/:id", UserController.GetUserByID)
	userGroup.GET("/", UserController.ListUsers)

}
