package app

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

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

	r.GET("/test", func(c *gin.Context) {
		processId := os.Getpid()

		log.Println("/test 1")
		time.Sleep(time.Second * 5)
		log.Println("/test 2")

		c.JSON(200, `processId: `+fmt.Sprint(processId))
	})

	r.Run("localhost:" + os.Getenv("PORT"))
}
