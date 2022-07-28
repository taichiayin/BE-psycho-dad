package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"psycho-dad/models"
	"psycho-dad/utils"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/parnurzeal/gorequest"
)

type FbToken struct {
	AccessToken string `json:"accessToken"`
}

type AppToken struct {
	AppToken string `json:"access_token"`
}

type ValidResp struct {
	Data *ValidRespData `json:"data"`
}

type ValidRespData struct {
	IsValid bool   `json:"is_valid"`
	UserId  string `json:"user_id"`
}

type UserInfoResp struct {
	Id      string       `json:"id"`
	Name    string       `json:"name"`
	Email   string       `json:"email"`
	Picture *PictureData `json:"picture"`
}

type PictureData struct {
	Data *Picture `json:"data"`
}

type Picture struct {
	Url string `json:"url"`
}

// FB secret
const app_id = "603757541009890"
const app_secret = "07954a0eca99605ccbfd7f073e2c552b"
const fbApiUrl_accessToken = "https://graph.facebook.com/v14.0/oauth/access_token?client_id=%s&client_secret=%s&grant_type=%s"
const fbApiUrl_debugToken = "https://graph.facebook.com/v14.0/debug_token?input_token=%s&access_token=%s"

// const fbApiUrl_userInfo = "https://graph.facebook.com/v14.0/%s"
const fbApiUrl_userInfo = "https://graph.facebook.com/%s?fields=id,name,email,picture&access_token=%s"

// 獲取應用程式權杖
func getFBAppToken() (string, error) {
	strUrl := fmt.Sprintf(fbApiUrl_accessToken, app_id, app_secret, "client_credentials")

	_, body, errs := gorequest.New().
		Get(strUrl).
		End()
	if errs != nil {
		return "", errors.New("請求失敗")
	}

	appToken := AppToken{}

	err := json.Unmarshal([]byte(body), &appToken)
	if err != nil {
		return "", errors.New("返回格式錯誤")
	}

	return appToken.AppToken, nil
}

// 獲取用戶數據
func getUserInfo(userId string, accessToken string) (*UserInfoResp, error) {
	strUrl := fmt.Sprintf(fbApiUrl_userInfo, userId, accessToken)

	_, body, errs := gorequest.New().
		Get(strUrl).
		End()

	if errs != nil {
		return nil, errors.New("請求失敗")
	}

	userInfo := UserInfoResp{}

	err := json.Unmarshal([]byte(body), &userInfo)
	if err != nil {
		return nil, errors.New(err.Error())
	}

	return &userInfo, nil
}

// db新增用戶
func addUserHandle(userInfo *UserInfoResp) {
	fmt.Println("addUserHanlde", userInfo)
	fbId, _ := strconv.Atoi(userInfo.Id)

	_, err := models.GetUserByFBId(fbId)

	userStrcut := models.User{}
	userInfoJson, _ := json.Marshal(userInfo)
	json.Unmarshal([]byte(userInfoJson), &userStrcut)
	userStrcut.FbId = userInfo.Id
	// 找不到新增用戶，已存在就更新
	if err != nil {
		fmt.Println("create user")
		models.CreateUser(userStrcut)
	} else {
		fmt.Println("update user")
		models.UpdateUserByFBId(fbId, userStrcut)
	}
}

// 驗證前端給的token是否有效，並且判斷新增用戶詳情
func ValidateFbToken(c *gin.Context) {
	fbToken := FbToken{}
	err := c.BindJSON(&fbToken)
	if err != nil {
		c.JSON(http.StatusOK, err.Error())
	}

	//
	app_token, err := getFBAppToken()
	if err != nil {
		c.JSON(http.StatusOK, err.Error())
	}

	// 驗證token
	strUrl := fmt.Sprintf(fbApiUrl_debugToken, fbToken.AccessToken, app_token)
	_, body, errs := gorequest.New().
		Get(strUrl).
		End()

	if errs != nil {
		errStr := fmt.Sprintf("%v", errs)
		c.JSON(http.StatusOK, errStr)
	}

	validResp := ValidResp{}

	err = json.Unmarshal([]byte(body), &validResp)
	if err != nil {
		c.JSON(http.StatusOK, err.Error())
	}

	// 無效token返回
	if !validResp.Data.IsValid {
		c.JSON(http.StatusOK, validResp)
		return
	}

	// 獲取用戶詳情
	userInfo, err := getUserInfo(validResp.Data.UserId, fbToken.AccessToken)
	if err != nil {
		c.String(http.StatusOK, err.Error())
		return
	}

	// db新增用戶
	addUserHandle(userInfo)

	c.JSON(http.StatusOK, utils.RespSuccess(userInfo))

}
