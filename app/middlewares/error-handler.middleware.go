package middlewares

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		var errMessage string

		for _, ginErr := range c.Errors {
			switch ginErr.Error() {
			case "400":
				errMessage = "bad request"
			case "404":
				errMessage = "not found"
			case "500":
				errMessage = "server error"
			default:
				errMessage = "unknown error"
			}

			errCode, _ := strconv.Atoi(ginErr.Error())

			c.JSON(errCode, `error: `+errMessage)
			return
		}
	}
}
