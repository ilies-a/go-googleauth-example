package app

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/ilies-a/go-googleauth-example/app/middlewares"
	"github.com/ilies-a/go-googleauth-example/app/routers"
)

func StartGinServer() {
	r := gin.New()

	r.NoRoute(gin.WrapH(http.FileServer(gin.Dir("./app/dist", false))))

	r.Use(
		gin.Recovery(),
		middlewares.ErrorHandler())

	routers.InitAPIRouter(r)

	r.Run("localhost:" + os.Getenv("PORT"))
}
