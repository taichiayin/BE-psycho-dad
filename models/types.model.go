package models

import (
	"fmt"
	"psycho-dad/config"

	"gorm.io/gorm"
)

type Type struct {
	Id   int    `json:"typeId"`
	Name string `json:"name"`
}

func GetAllTypes() []Type {
	var types []Type
	config.Conn.Find(&types)

	return types
}

func GetTypeById(userId int) Type {
	myType := Type{}
	err := config.Conn.Where("id = ?", userId).First(&myType).Error

	if err == gorm.ErrRecordNotFound {
		// return "Not Foun Record"
		// return
		fmt.Println("RECORD NOT FOUND")
	}

	return myType
}

func CreateType(myType Type) string {
	err := config.Conn.Create(&myType).Error
	if err != nil {
		return "Error :" + err.Error()
	}

	return "CREATE SUCCESSFUL"
}

func UpdateType(typeId int, myType Type) string {
	err := config.Conn.Model(&myType).Where("id = ?", typeId).Updates(myType).Error
	if err != nil {
		return "Error :" + err.Error()
	}

	return "UPDATE SUCCESSFUL"
}

func DeleteType(typeId int) string {
	myType := Type{}
	err := config.Conn.Where("id = ?", typeId).Delete(&myType).Error
	if err != nil {
		return "Error :" + err.Error()
	}
	return "DELETE SUCCESSFUL"
}
