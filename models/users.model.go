package models

import (
	"fmt"
	"psycho-dad/config"
)

type User struct {
	Id       int    `json:"userId" gorm:"AUTO_INCREMENT"`
	Username string `json:"username"`
	Email    string `json:"email"`
	FbId     string `json:"fbId"`
	Password string `json:"password"`
	Avatar   string `json:"avatar"`
	IsAdmin  bool   `json:"isAdmin"`
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

func FindByUsername(username string) (*User, error) {
	user := &User{}

	err := config.Conn.Table("Users").Where("username = ?", username).First(user).Error
	if err != nil {
		return nil, err
	}

	return user, nil

}

func UpdateUserByFBId(id int, user User) string {
	err := config.Conn.Model(&user).Omit("id").Where("fb_id = ?", id).Updates(user).Error
	if err != nil {
		return err.Error()
	}

	return "UPDATE SUCCESSFUL"
}

func CreateUser(user *User) error {
	err := config.Conn.Create(&user).Error
	if err != nil {
		return err
	}

	return nil
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
