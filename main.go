package main

import (
	"psycho-dad/config"
	"psycho-dad/controllers"
	"psycho-dad/middleware"
	"psycho-dad/routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	// "github.com/unrolled/secure"
)

func main() {
	r := gin.Default()
	r.Use(cors.Default())
	// https
	// r.Use(TlsHandler())
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
	// routes.AddImgRouter(v1)

	r.POST("/v1/login", controllers.Login)
	r.POST("/v1/fblogin", controllers.Fblogin)
	r.Static("/img", "./img")

	go func() {
		// 連接資料庫
		config.StartUp()
	}()

	// r.RunTLS(":3088", "./certs/localhost.pem", "./certs/localhost-key.pem")
	r.Run(":3088")
}

// func TlsHandler() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		secureMiddleware := secure.New(secure.Options{
// 			SSLRedirect: true,
// 			SSLHost:     "localhost:3088",
// 		})
// 		err := secureMiddleware.Process(c.Writer, c.Request)

// 		// If there was an error, do not continue.
// 		if err != nil {
// 			return
// 		}

// 		c.Next()
// 	}
// }
