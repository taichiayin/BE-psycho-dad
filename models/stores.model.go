package models

import "psycho-dad/config"

type Store struct {
	Id              int     `json:"storeId"`
	Name            string  `json:"storeName"`
	Introduce       string  `json:"introduce"`
	TypeId          int     `json:"typeId"`
	CountyId        int     `json:"countyId"`
	DistrictId      int     `json:"districtId"`
	Address         string  `json:"address"`
	Phone           string  `json:"phone"`
	Mobile          string  `json:"mobile"`
	Email           string  `json:"email"`
	OpenTime1       string  `json:"openTime1"`
	OpenTime2       string  `json:"openTime2"`
	OpenTime3       string  `json:"openTime3"`
	UserId          int     `json:"userId"`
	IsDads          bool    `json:"isDads"`
	IsDadsRecommend bool    `json:"isDadsRecommend"`
	WebSite         string  `json:"webSite"`
	Lon             float64 `json:"lon"`
	Lat             float64 `json:"lat"`
	Dis             float64 `json:"dis"`
}

func GetAllStores() []Store {
	var stores []Store
	config.Conn.Find(&stores)

	return stores
}

func GetStoreById(storeId int) (*Store, error) {
	store := &Store{}
	err := config.Conn.Where("id = ?", storeId).First(store).Error
	if err != nil {
		return nil, err
	}

	return store, nil
}

func CreateStore(store Store) string {
	err := config.Conn.Create(&store).Error

	if err != nil {
		return "Error" + err.Error()
	}
	return "CREATE SUCCESSFUL"
}

func DeleteStore(storeId int) string {
	var store Store
	err := config.Conn.Where("id = ?", storeId).Delete(&store).Error

	if err != nil {
		return "Error" + err.Error()
	}
	return "DELETE SUCCESSFUL"
}

func UpdateStore(storeId int, store Store) string {
	err := config.Conn.Where("id = ?", storeId).Updates(store).Error
	if err != nil {
		return "Error" + err.Error()
	}
	return "Update SUCCESSFUL"
}
