package models

import (
	"psycho-dad/config"
)

type User struct {
	Id     int    `json:"userId"`
	Name   string `json:"userName"`
	Avatar string `json:"userAvatar"`
}

func GetAllUsers() []User {

	var users []User
	config.Conn.Find(&users)

	return users
}

func GetUserById(userId int) User {
	var user User
	config.Conn.Where("id = ?", userId).First(&user)

	return user
}

func CreateUser(user User) string {
	err := config.Conn.Create(&user).Error
	if err != nil {
		return "Error" + err.Error()
	}

	return "CREATE SUCCESSFUL"
}

func UpdateUser(userId int, user User) string {
	err := config.Conn.Model(&user).Omit("id").Where("id = ?", userId).Updates(user).Error
	if err != nil {
		return "Error" + err.Error()
	}

	return "UPDATE SUCCESSFUL"
}

func DeleteUser(userId int) string {
	user := User{}
	err := config.Conn.Where("id = ?", userId).Delete(&user).Error
	if err != nil {
		return "Error" + err.Error()
	}

	return "DELETE SUCCESSFUL"
}
