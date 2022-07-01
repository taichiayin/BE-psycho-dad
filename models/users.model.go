package models

import (
	"psycho-dad/config"
	"psycho-dad/utils"
)

type User struct {
	Id     int    `json:userId`
	Name   string `json:userName`
	Avatar string `json:userAvatar`
}

func GetAllUsers() utils.ResponseStruct {
	var users []User
	config.Conn.Find(&users)
	// utils.HandleResponse(users)
	// fmt.Println(users)
	return utils.HandleResponse(users)
}

// func CreateUser(user User) string{
// 	res := config.Conn.Create(&user)

// 	if res.Error != nil {
// 		log.Fatalln(res.Error)
// 	}

// 	return "Su"
// }
