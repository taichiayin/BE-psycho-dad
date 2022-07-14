package controllers

import (
	"fmt"
	"net/http"
	"psycho-dad/models"
	"psycho-dad/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

func FindAllUsers(c *gin.Context) {
	users := models.GetAllUsers()
	fmt.Println("FindAllUsers")
	c.JSON(http.StatusOK, utils.RespSuccess(users))
}

func FindByUserId(c *gin.Context) {
	userId, _ := strconv.Atoi(c.Param("id"))
	user, err := models.GetUserById(userId)

	if err != nil {
		c.String(http.StatusOK, err.Error())
		return
	}

	c.JSON(http.StatusOK, utils.RespSuccess(user))
}

func CreateUser(c *gin.Context) {
	user := models.User{}
	err := c.BindJSON(&user)
	if err != nil {
		c.JSON(http.StatusNotAcceptable, "Error: "+err.Error())
		return
	}

	models.CreateUser(user)
	c.JSON(http.StatusOK, "CREATE SUCCESSFUL")
}

func UpdateUser(c *gin.Context) {
	userId, _ := strconv.Atoi(c.Param("id"))
	user := models.User{}
	err := c.BindJSON(&user)
	if err != nil {
		c.JSON(http.StatusAccepted, "Error:"+err.Error())
	}

	models.UpdateUser(userId, user)
	c.JSON(http.StatusOK, "UPDATE SUCCESSFUL")
}

func DeleteUser(c *gin.Context) {
	userId, _ := strconv.Atoi(c.Param("id"))
	models.DeleteUser(userId)

	c.JSON(http.StatusOK, "DELETE SUCCESSFUL")
}
