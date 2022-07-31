package main

import (
	"net/http"
	"psycho-dad/config"
	"psycho-dad/controllers"
	_ "psycho-dad/daemon"
	"psycho-dad/middleware"
	"psycho-dad/routes"

	"github.com/gin-gonic/gin"
	// "github.com/unrolled/secure"
)

func main() {
	r := gin.Default()

	// 處理cros options問題
	r.Use(middleware.CrosOptionsHandle())

	v1 := r.Group("/v1")

	// middlewar 驗證token
	v1.Use(middleware.JWTAuth())

	routes.AddUserRouter(v1)
	routes.AddThumbsRouter(v1)
	routes.AddStoresRouter(v1)
	routes.AddFavoritesRouter(v1)
	routes.AddTypesRouter(v1)
	routes.AddCountiesRouter(v1)
	routes.AddDistrictsRouter(v1)
	routes.AddUploadRouter(v1)

	r.POST("/v1/login", controllers.Login)
	r.POST("/v1/fblogin", controllers.Fblogin)
	r.Static("/img", "./img")

	// 測試api pass
	r.GET("/yeah", func(c *gin.Context) {
		c.JSON(http.StatusOK, "yeah")
	})

	go func() {
		// 連接資料庫
		config.StartUp()
	}()

	r.Run(":3088")
}
