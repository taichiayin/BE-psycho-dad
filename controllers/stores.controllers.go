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
	store := models.Store{}
	err := c.BindQuery(&store)
	if err == nil {
		fmt.Println(store)
		fmt.Println(store.CountyId)
	}
	stores, _ := models.GetAllStores(&store)
	c.JSON(http.StatusOK, utils.RespSuccess(stores))
}

func GetAllStore(c *gin.Context) {
	myLon, _ := strconv.ParseFloat(c.Query("lon"), 64)
	mylat, _ := strconv.ParseFloat(c.Query("lat"), 64)

	store := models.Store{}

	err := c.ShouldBindQuery(&store)

	if err == nil {
		fmt.Println(store)
		// fmt.Println(store.PageIndex)
	}

	stores, p := models.GetAllStores(&store)
	fmt.Printf("經度: %f", myLon)
	fmt.Printf("緯度: %f", mylat)
	// 計算最短距離，若資料不足則顯示9999
	for i := 0; i < len(stores); i++ {
		if myLon != 0 && mylat != 0 && stores[i].Lon != 0 && stores[i].Lat != 0 {
			dis := utils.CalGeoDistance(float64(myLon), float64(mylat), stores[i].Lon, stores[i].Lat) / 1000
			disKM, _ := strconv.ParseFloat(fmt.Sprintf("%.3f", dis), 64)
			stores[i].Dis = disKM
		} else {
			stores[i].Dis = 9999
		}
	}

	c.JSON(http.StatusOK, utils.PagingRespSuccess(stores, &p))
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
	err = c.BindJSON(&store)
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
