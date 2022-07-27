package middleware

import (
	"net/http"
	"psycho-dad/controllers"
	"psycho-dad/utils"

	"github.com/gin-gonic/gin"
)

func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 取header裡的token
		token := c.Request.Header.Get("token")
		if token == "" {
			c.JSON(http.StatusOK, utils.RespError("請求未帶token，請重新登入"))
			c.Abort()
			return
		}
		// fmt.Println("token =", token)

		// 解析出payload
		j := controllers.NewJWT()

		claims, err := j.ParserToken(token)

		if err != nil {
			// token過期
			if err == controllers.TokenExpired {
				c.JSON(http.StatusOK, gin.H{
					"code":    2099,
					"message": "token已過期，請重新登入",
					"data":    nil,
				})
				c.Abort()
				return
			}
			// 其他錯誤
			c.JSON(http.StatusOK, gin.H{
				"code":    2099,
				"message": err.Error(),
				"data":    nil,
			})
			c.Abort()
			return
		}

		// 解析成功，取userId, 可使用c.get("UserId")得到UserId
		c.Set("UserId", claims.UserId)
	}
}
