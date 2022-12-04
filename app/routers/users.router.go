package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/ilies-a/go-googleauth-example/app/handlers"
)

func InitUsersRouter(rg *gin.RouterGroup) {
	usersGroup := rg.Group("/users")
	{
		usersGroup.GET("/", handlers.GetAllUsers)
		usersGroup.GET("/:id", handlers.GetUser)
		usersGroup.POST("/", handlers.SaveUser)
		usersGroup.DELETE("/:id", handlers.DeleteUser)

	}
}
