package models

import "psycho-dad/config"

type District struct {
	Id       int    `json:"districtId"`
	CountyId int    `json:"countyId"`
	Name     string `json:"name"`
}

func GetAllDistricts() (*[]District, error) {
	districts := &[]District{}
	err := config.Conn.Find(&districts).Error
	if err != nil {
		return nil, err
	}

	return districts, nil
}

func GetDistrictById(districtId int) (*District, error) {
	district := &District{}
	err := config.Conn.Where("id = ?", districtId).Find(district).Error

	if err != nil {
		return nil, err
	}

	return district, nil
}

func GetDistrictByCountyId(countyId int) (*[]District, error) {
	districts := &[]District{}
	err := config.Conn.Where("county_id = ?", countyId).Find(&districts).Error
	if err != nil {
		return nil, err
	}
	return districts, nil
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
