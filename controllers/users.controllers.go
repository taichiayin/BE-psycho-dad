package controllers

import (
	"net/http"
	"psycho-dad/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

func FindAllUsers(c *gin.Context) {
	users := models.GetAllUsers()
	c.JSON(http.StatusOK, users)
}

func FindByUserId(c *gin.Context) {
	userId, _ := strconv.Atoi(c.Param("id"))
	user, err := models.GetUserById(userId)

	if err != nil {
		c.String(http.StatusOK, err.Error())
		return
	}

	c.JSON(http.StatusOK, user)
}

func CreateUser(c *gin.Context) {
	user := models.User{}
	err := c.BindJSON(&user)
	if err != nil {
		c.JSON(http.StatusNotAcceptable, "Error: "+err.Error())
		return
	}

	newUser := models.CreateUser(user)
	c.JSON(http.StatusOK, newUser)
}

func UpdateUser(c *gin.Context) {
	userId, _ := strconv.Atoi(c.Param("id"))
	user := models.User{}
	err := c.BindJSON(&user)
	if err != nil {
		c.JSON(http.StatusAccepted, "Error:"+err.Error())
	}

	newUser := models.UpdateUser(userId, user)
	c.JSON(http.StatusOK, newUser)
}

func DeleteUser(c *gin.Context) {
	userId, _ := strconv.Atoi(c.Param("id"))
	res := models.DeleteUser(userId)

	c.JSON(http.StatusOK, res)
}
