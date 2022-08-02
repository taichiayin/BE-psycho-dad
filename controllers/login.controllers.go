package controllers

import (
	"errors"
	"net/http"
	"psycho-dad/models"
	"psycho-dad/utils"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	secret                 = "!p@s#y$c%h.o&d?a,d"
	TokenExpired     error = errors.New("Token過期，請重新登入")
	TokenNotValidYet error = errors.New("Token未啟用")
	TokenMalformed   error = errors.New("Token格式錯誤")
	TokenInvalid     error = errors.New("Toekn無效")
)

// fb登入請求結構
type FbLoginReq struct {
	FbId   string `json:"id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	Avatar string `json:"avatarUrl"`
}

// api請求結構
type Req struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// api反饋結構
type Resp struct {
	UserId   int    `json:"userId"`
	Username string `json:"username"`
	Avatar   string `json:"avatar"`
	Email    string `json:"email"`
	IsAdmin  bool   `json:"isAdmin"`
	Token    string `json:"token"`
}

// jwt payload結構
type Claims struct {
	UserId int `json:"userId"`
	jwt.StandardClaims
}

// key結構
type JWT struct {
	SigningKey []byte
}

// 建立signkey
func NewJWT() *JWT {
	return &JWT{
		[]byte(secret),
	}
}

// 建立token
// HS256
// claims + key(sign)生成token
func (j *JWT) CreateToken(claims Claims) (string, error) {
	// 返回一個token的指標
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(j.SigningKey)
}

// 生成token
func generateToken(user *models.User) (string, error) {
	// 建立signkey的JWT結構
	j := NewJWT()
	now := time.Now()

	// 建立payload
	claims := Claims{
		user.Id,
		jwt.StandardClaims{
			NotBefore: now.Add(-3 * time.Second).Unix(), // 生效時間
			ExpiresAt: now.Add(24 * time.Hour).Unix(),   // 過期時間
			IssuedAt:  now.Unix(),
			Issuer:    "ginJWT",
		},
	}

	// 依據claims設定生成token
	token, err := j.CreateToken(claims)
	if err != nil {
		return "", err
	}

	return token, nil
}

// 解析token 返回Claims指標及錯誤
func (j *JWT) ParserToken(tokenStr string) (*Claims, error) {
	// Keyfunc是匿名函数类型: type Keyfunc func(*Token) (interface{}, error)
	// func ParseWithClaims(tokenString string, claims Claims, keyFunc Keyfunc) (*Token, error)
	token, err := jwt.ParseWithClaims(tokenStr, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return j.SigningKey, nil
	})

	if err != nil {
		// jwt.ValidationError 是一個無效token的錯誤結構
		if ve, ok := err.(*jwt.ValidationError); ok {
			// 表示token不可用
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, TokenMalformed
				// 表示token過期
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				return nil, TokenExpired
				// 表示token格式錯誤
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, TokenNotValidYet
				// 歸類表示token無效
			} else {
				return nil, TokenInvalid
			}
		}
	}

	// 解析token中的claims
	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}

	return nil, TokenInvalid
}

// 登入
func Login(c *gin.Context) {
	req := Req{}

	// 解析body json
	err := c.BindJSON(&req)
	if err != nil {
		c.JSON(http.StatusOK, utils.RespError(err.Error()))
		return
	}

	// find by username
	user, err := models.FindByUsername(req.Username)
	if err != nil {
		c.JSON(http.StatusOK, utils.RespError("登入失敗，帳號密碼輸入不正確，請重新輸入！"))
		return
	}

	// 比對密碼是否正確
	if utils.Md5(req.Password) != user.Password {
		c.JSON(http.StatusOK, utils.RespError("登入失敗，帳號密碼輸入不正確，請重新輸入！"))
		return
	}

	// 產生token
	token, err := generateToken(user)

	if err != nil {
		c.JSON(http.StatusOK, utils.RespError(err.Error()))
		return
	}

	// 組成反饋api
	resp := Resp{
		UserId:   user.Id,
		Username: user.Username,
		Avatar:   user.Avatar,
		Email:    user.Email,
		IsAdmin:  user.IsAdmin,
		Token:    token,
	}

	c.JSON(http.StatusOK, utils.RespSuccess(resp))
}

func Fblogin(c *gin.Context) {
	fbLoginReq := &FbLoginReq{}
	user := &models.User{}

	err := c.ShouldBind(fbLoginReq)
	if err != nil {
		c.JSON(http.StatusOK, utils.RespError(err.Error()))
	}

	fbId, _ := strconv.Atoi(fbLoginReq.FbId)
	user, err = models.GetUserByFBId(fbId)
	if err != nil {
		// fb_id沒找到user
		if errors.Is(err, gorm.ErrRecordNotFound) {
			myUser := &models.User{
				FbId:     fbLoginReq.FbId,
				Email:    fbLoginReq.Email,
				Username: fbLoginReq.Name,
			}

			err = models.CreateUser(myUser)
			if err != nil {
				// 新增用戶失敗
				c.JSON(http.StatusOK, utils.RespError(err.Error()))
				return
			}

			// 新增用戶成功
			// c.JSON(http.StatusOK, utils.RespSuccess(user))
			// return
		} else {
			c.JSON(http.StatusOK, utils.RespError(err.Error()))
			return
		}
	}

	// 產生token
	token, err := generateToken(user)

	if err != nil {
		c.JSON(http.StatusOK, utils.RespError(err.Error()))
		return
	}

	// 組成反饋api
	resp := Resp{
		UserId:   user.Id,
		Username: user.Username,
		Avatar:   user.Avatar,
		Email:    user.Email,
		IsAdmin:  user.IsAdmin,
		Token:    token,
	}

	c.JSON(http.StatusOK, utils.RespSuccess(resp))
}
