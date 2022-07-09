package models

import "psycho-dad/config"

type District struct {
	Id   int `json:"favoriteId"`
	Name int `json:"name"`
}

func GetAllDistrict() []District {
	var districts []District
	config.Conn.Find(&districts)

	return districts
}

func GetDistrictsByUser(districtId int) District {
	district := District{}
	config.Conn.Where("user_id = ?", districtId).Find(&district)

	return district
}

func CreateDistrict(district District) string {
	err := config.Conn.Create(&district).Error
	if err != nil {
		return "Error :" + err.Error()
	}

	return "CREATE SUCCESSFUL"
}

func DeleteDistrict(typeId int) string {
	district := District{}
	err := config.Conn.Where("id = ?", typeId).Delete(&district).Error
	if err != nil {
		return "Error :" + err.Error()
	}
	return "DELETE SUCCESSFUL"
}
