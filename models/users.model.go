package models

import (
	"fmt"
	"psycho-dad/config"
)

type User struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	FbId  string `json:"fbId"`
}

func GetAllUsers() []User {
	var users []User
	config.Conn.Find(&users)

	return users
}

func GetUserById(id int) (*User, error) {
	user := &User{}
	err := config.Conn.Where("id = ?", id).First(user).Error

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return user, nil
}

func GetUserByFBId(id int) (*User, error) {
	user := &User{}
	err := config.Conn.Where("fb_id = ?", id).First(user).Error

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return user, nil
}

func CreateUser(user User) string {
	err := config.Conn.Create(&user).Error
	if err != nil {
		return err.Error()
	}

	return "CREATE SUCCESSFUL"
}

func UpdateUser(userId int, user User) string {
	err := config.Conn.Model(&user).Omit("id").Where("id = ?", userId).Updates(user).Error
	if err != nil {
		return err.Error()
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
