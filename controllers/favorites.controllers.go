package controllers

import (
	"net/http"
	"psycho-dad/models"
	"psycho-dad/utils"

	"github.com/gin-gonic/gin"
)

func GetAllFavorites(c *gin.Context) {
	favorites := models.GetAllFavorites()
	c.JSON(http.StatusOK, favorites)
}

func GetFavoritesByUser(c *gin.Context) {
	userId, _ := c.Get("UserId")
	userIdInt := userId.(int)
	// userId, _ := strconv.Atoi(c.Param("userId"))
	res, err := models.GetFavoritesByUser(userIdInt)
	if err != nil {
		c.JSON(http.StatusOK, utils.RespError(err.Error()))
		return
	}

	c.JSON(http.StatusOK, utils.RespSuccess(res))
}

func CreateFavorite(c *gin.Context) {
	favorite := models.Favorite{}
	err := c.BindJSON(&favorite)
	if err != nil {
		c.JSON(http.StatusAccepted, utils.RespError(err.Error()))
		return
	}

	err = models.CreateFavorite(favorite)

	if err != nil {
		c.JSON(http.StatusAccepted, utils.RespError(err.Error()))
		return
	}

	c.JSON(http.StatusOK, utils.RespSuccess(err))
}

func DeleteFavorite(c *gin.Context) {
	// favoriteId, _ := strconv.Atoi(c.Param("favoriteId"))
	favorite := models.Favorite{}

	err := c.ShouldBind(&favorite)
	if err != nil {
		c.JSON(http.StatusAccepted, utils.RespError(err.Error()))
		return
	}

	err = models.DeleteFavorite(favorite)

	if err != nil {
		c.JSON(http.StatusAccepted, utils.RespError(err.Error()))
		return
	}

	c.JSON(http.StatusOK, utils.RespSuccess(err))
}
