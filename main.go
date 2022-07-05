package main

import (
	"psycho-dad/config"
	"psycho-dad/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	v1 := r.Group("/v1")

	routes.AddUserRouter(v1)
	routes.AddThumbsRouter(v1)
	routes.AddStoresRouter(v1)
	routes.AddFavoritesRouter(v1)

	go func() {
		// 連接資料庫
		config.StartUp()
	}()

	r.Run(":3088")
}
