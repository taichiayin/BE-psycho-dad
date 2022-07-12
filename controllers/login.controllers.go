package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/parnurzeal/gorequest"
)

type FbToken struct {
	AccessToken string `json:"accessToken"`
}

// FB secret
var app_id = "603757541009890"
var app_secret = "07954a0eca99605ccbfd7f073e2c552b"

func GetFBAppToken(c *gin.Context) {
	strURL := "https://graph.facebook.com/v14.0/oauth/access_token?client_id=" + app_id +
		"&client_secret=" + app_secret +
		"&grant_type=client_credentials"
	_, body, errs := gorequest.New().
		Get(strURL).
		End()

	if errs != nil {
		c.JSON(http.StatusOK, "error")
		// return "Error"
	}

	fmt.Println(body)
	c.JSON(http.StatusOK, body)
}

func ValidateFbToken(c *gin.Context) {
	fbToken := FbToken{}
	err := c.BindJSON(&fbToken)
	if err != nil {
		c.JSON(http.StatusOK, err.Error())
	}

	resp, body, errs := gorequest.New().
		Get("https://graph.facebook.com/v14.0/debug_token?input_token=" + fbToken.AccessToken +
			"&access_token=603757541009890|VZYiBRkI99kBBiIGdlbWfl7uJ14").
		End()

	if errs != nil {
		c.JSON(http.StatusOK, "Error")
	}

	fmt.Println(resp)
	fmt.Println(body)

	c.JSON(http.StatusOK, body)

}
