package models

import (
	"psycho-dad/config"
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

func GetTypeById(userId int) (*Type, error) {
	myType := &Type{}
	err := config.Conn.Where("id = ?", userId).First(myType).Error

	if err != nil {
		return nil, err
	}

	return myType, nil
}

func CreateType(myType Type) error {
	err := config.Conn.Create(&myType).Error
	if err != nil {
		return err
	}

	return nil
}

func UpdateType(typeId int, myType Type) error {
	err := config.Conn.Model(&myType).Where("id = ?", typeId).Updates(myType).Error
	if err != nil {
		return err
	}
	return nil
}

func DeleteType(typeId int) error {
	myType := Type{}
	err := config.Conn.Where("id = ?", typeId).Delete(&myType).Error
	if err != nil {
		return err
	}
	return nil
}
