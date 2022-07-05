package controllers

import (
	"net/http"
	"psycho-dad/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAllFavorites(c *gin.Context) {
	favorites := models.GetAllFavorites()
	c.JSON(http.StatusOK, favorites)
}

func CreateFavorite(c *gin.Context) {
	favorite := models.Favorite{}
	err := c.BindJSON(&favorite)
	if err != nil {
		c.JSON(http.StatusAccepted, "Error: "+err.Error())
	}

	res := models.CreateFavorite(favorite)

	c.JSON(http.StatusOK, res)
}

func DeleteFavorite(c *gin.Context) {
	favoriteId, _ := strconv.Atoi(c.Param("favoriteId"))
	res := models.DeleteFavorite(favoriteId)

	c.JSON(http.StatusOK, res)
}
