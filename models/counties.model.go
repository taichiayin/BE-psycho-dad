package models

import "psycho-dad/config"

type County struct {
	Id   int    `json:"countyId"`
	Name string `json:"name"`
}

func GetAllCounties() []County {
	var counties []County
	config.Conn.Find(&counties)

	return counties
}

func GetCountyById(countyId int) County {
	var county County
	config.Conn.Where("id = ?", countyId).Find(&county)

	return county
}

func CreateCounty(county County) string {
	err := config.Conn.Create(&county).Error
	if err != nil {
		return "Error :" + err.Error()
	}

	return "CREATE SUCCESSFUL"
}

func UpdateCounty(countyId int, county County) string {
	err := config.Conn.Model(&county).Where("id = ?", countyId).Updates(county).Error
	if err != nil {
		return "Error :" + err.Error()

	}
	return "UPDATE SUCCESSFUL"
}

func DeleteCounty(countyId int) string {
	county := County{}
	err := config.Conn.Where("id = ?", countyId).Delete(&county).Error
	if err != nil {
		return "Error :" + err.Error()
	}
	return "DELETE SUCCESSFUL"
}
