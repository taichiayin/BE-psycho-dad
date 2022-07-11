package main

import (
	"psycho-dad/config"
	"psycho-dad/routes"

	"github.com/gin-gonic/gin"
	// "github.com/unrolled/secure"
)

func main() {
	r := gin.Default()
	// https
	// r.Use(TlsHandler())
	v1 := r.Group("/v1")

	routes.AddUserRouter(v1)
	routes.AddThumbsRouter(v1)
	routes.AddStoresRouter(v1)
	routes.AddFavoritesRouter(v1)
	routes.AddTypesRouter(v1)
	routes.AddCountiesRouter(v1)
	routes.AddDistrictsRouter(v1)

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
