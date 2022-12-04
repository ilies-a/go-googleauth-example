package routers

import "github.com/gin-gonic/gin"

func InitAPIRouter(r *gin.Engine) {
	apiGroup := r.Group("/api")
	{
		InitUsersRouter(apiGroup)
	}
}
