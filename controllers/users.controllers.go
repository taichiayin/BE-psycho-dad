package controllers

import (
	"net/http"
	"psycho-dad/models"

	"github.com/gin-gonic/gin"
)

func FindAllUsers(c *gin.Context) {
	users := models.GetAllUsers()
	c.JSON(http.StatusOK, users)
}

func PostUser(c *gin.Context) {
	user := models.User{}
	err := c.BindJSON(&user)
	if err != nil {
		c.JSON(http.StatusNotAcceptable, "Error: "+err.Error())
		return
	}
	newUser := models.CreateUser(user)
	c.JSON(http.StatusOK, newUser)
}
