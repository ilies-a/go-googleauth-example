package handlers

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ilies-a/go-googleauth-example/app/models"
	"github.com/ilies-a/go-googleauth-example/app/utils"
)

func GetAllUsers(c *gin.Context) {
	users := models.GetAllUsers()
	c.JSON(http.StatusOK, users)
}

func GetUser(c *gin.Context) {
	userId := c.Param("id")
	if userId == "" {
		c.Errors = append(c.Errors, c.Error(errors.New("400")))
		return
	}

	user, err := models.GetUser(userId)
	if err != nil {
		c.Errors = append(c.Errors, c.Error(errors.New("404")))
		return
	}

	c.JSON(http.StatusOK, user)
}

func SaveUser(c *gin.Context) {
	var user models.User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.Errors = append(c.Errors, c.Error(errors.New("400")))
		return
	}

	user.Id = utils.GenerateStringUUID()

	models.SaveUser(&user)

	c.JSON(201, user)
}

func DeleteUser(c *gin.Context) {
	userId := c.Param("id")
	if userId == "" {
		c.Errors = append(c.Errors, c.Error(errors.New("400")))
		return
	}

	err := models.DeleteUser(userId)
	if err != nil {
		c.Errors = append(c.Errors, c.Error(errors.New("404")))
		return
	}

	c.JSON(200, "user "+userId+" successfully deleted")
}
