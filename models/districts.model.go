package models

import "psycho-dad/config"

type District struct {
	Id       int    `json:"districtId"`
	CountyId int    `json:"countyId"`
	Name     string `json:"name"`
}

func GetAllDistricts() []District {
	var districts []District
	config.Conn.Find(&districts)

	return districts
}

func GetDistrictById(districtId int) District {
	district := District{}
	config.Conn.Where("id = ?", districtId).Find(&district)

	return district
}

func CreateDistrict(district District) string {
	err := config.Conn.Create(&district).Error
	if err != nil {
		return "Error :" + err.Error()
	}

	return "CREATE SUCCESSFUL"
}

func UpdateDistrict(districtId int, district District) string {
	err := config.Conn.Model(&district).Where("id = ?", districtId).Updates(district).Error
	if err != nil {
		return "Error :" + err.Error()
	}
	return "UPDATE SUCCESSFUL"
}

func DeleteDistrict(typeId int) string {
	district := District{}
	err := config.Conn.Where("id = ?", typeId).Delete(&district).Error
	if err != nil {
		return "Error :" + err.Error()
	}
	return "DELETE SUCCESSFUL"
}
