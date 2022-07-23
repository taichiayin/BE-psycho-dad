package controllers

import (
	"fmt"
	"net/http"
	"psycho-dad/models"
	"psycho-dad/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAllStoreForEdit(c *gin.Context) {
	page, _ := strconv.Atoi(c.Query("page"))
	size, _ := strconv.Atoi(c.Query("size"))

	storeApi := models.StoreApi{}
	err := c.BindQuery(&storeApi)
	if err == nil {
		fmt.Println(storeApi)
		fmt.Println(storeApi.CountyId)
	}
	storeApis, _ := models.GetAllStores(&storeApi, page, size, "")
	c.JSON(http.StatusOK, utils.RespSuccess(storeApis))
}

func GetAllStore(c *gin.Context) {
	myLon, _ := strconv.ParseFloat(c.Query("lon"), 64)
	mylat, _ := strconv.ParseFloat(c.Query("lat"), 64)
	page, _ := strconv.Atoi(c.Query("page"))
	size, _ := strconv.Atoi(c.Query("size"))
	userId := c.Query("id")

	storeApi := models.StoreApi{}

	err := c.ShouldBindQuery(&storeApi)

	if err == nil {
		fmt.Println(storeApi)
		// fmt.Println(store.PageIndex)
	}

	storeApis, p := models.GetAllStores(&storeApi, page, size, userId)
	fmt.Printf("經度: %f", myLon)
	fmt.Printf("緯度: %f", mylat)
	// 計算最短距離，若資料不足則顯示9999
	for i := 0; i < len(storeApis); i++ {
		if myLon != 0 && mylat != 0 && storeApis[i].Lon != 0 && storeApis[i].Lat != 0 {
			dis := utils.CalGeoDistance(float64(myLon), float64(mylat), storeApis[i].Lon, storeApis[i].Lat) / 1000
			disKM, _ := strconv.ParseFloat(fmt.Sprintf("%.3f", dis), 64)
			storeApis[i].Dis = disKM
		} else {
			storeApis[i].Dis = 9999
		}
	}

	c.JSON(http.StatusOK, utils.PagingRespSuccess(storeApis, &p))
}

func GetStoreById(c *gin.Context) {
	storeId, _ := strconv.Atoi(c.Param("id"))
	user, err := models.GetStoreById(storeId)
	if err != nil {
		c.JSON(http.StatusOK, utils.RespError(err.Error()))
		return
	}
	c.JSON(http.StatusOK, utils.RespSuccess(user))
}

func CreateStore(c *gin.Context) {
	store := models.Store{}
	err := c.BindJSON(&store)
	if err != nil {
		c.JSON(http.StatusNotAcceptable, "Error: "+err.Error())
		return
	}

	err = models.CreateStore(store)
	if err != nil {
		c.JSON(http.StatusOK, utils.RespError(err.Error()))
		return
	}

	c.JSON(http.StatusOK, utils.RespSuccess(nil))
}

func UpdateStore(c *gin.Context) {
	storeId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusOK, "Api path 'ID' has error.")
	}

	store := models.Store{}
	err = c.ShouldBind(&store)
	if err != nil {
		c.JSON(http.StatusOK, "Error:"+err.Error())
		return
	}

	err = models.UpdateStore(storeId, store)
	if err != nil {
		c.JSON(http.StatusOK, utils.RespError(err.Error()))
		return
	}

	c.JSON(http.StatusOK, utils.RespSuccess(nil))
}

func DeleteStore(c *gin.Context) {
	storeId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusOK, "Api path 'ID' has error.")
	}

	res := models.DeleteStore(storeId)
	c.JSON(http.StatusOK, res)
}
