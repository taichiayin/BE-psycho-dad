package models

import (
	"psycho-dad/config"
)

type Favorite struct {
	Id      int `json:"favoriteId"`
	StoreId int `json:"storeId"`
	UserId  int `json:"userId"`
}

type FavoriteApi struct {
	Id                 int     `json:"favoriteId"`
	StoreId            int     `json:"storeId"`
	UserId             int     `json:"userId"`
	Name               string  `json:"storeName" form:"storeName"`
	Introduce          string  `json:"introduce"`
	TypeId             int     `json:"typeId" form:"typeId"`
	TypeName           string  `json:"typeName"`
	CountyId           int     `json:"countyId" form:"countyId"`
	CountyName         string  `json:"countyName"`
	DistrictId         int     `json:"districtId" form:"districtId"`
	DistrictName       string  `json:"districtName"`
	Address            string  `json:"address"`
	Phone              string  `json:"phone"`
	Mobile             string  `json:"mobile"`
	Email              string  `json:"email"`
	IsDads             bool    `json:"isDads"`
	IsDadsRecommend    bool    `json:"isDadsRecommend"`
	WebSite            string  `json:"webSite"`
	Lon                float64 `json:"lon"`
	Lat                float64 `json:"lat"`
	Dis                float64 `json:"dis" gorm:"-"`
	IsClosePermanently bool    `json:"isClosePermanently"`
	DefaultImg         string  `json:"defaultImg"`
}

func GetAllFavorites() []Favorite {
	var favorites []Favorite
	config.Conn.Find(&favorites)

	return favorites
}

func GetFavoritesByUser(userId int) (*[]FavoriteApi, error) {
	// var favorites []Favorite
	favoriteApis := &[]FavoriteApi{}
	err := config.Conn.Table("favorites").
		Select("stores.*, files.default_img, types.name as type_name, counties.name as county_name, districts.name as district_name , stores.id as store_id").
		Joins("left join stores on favorites.store_id = stores.id").
		Joins("left join files on files.store_id = stores.id").
		Joins("left join types on types.id = stores.type_id").
		Joins("left join counties on counties.id = stores.county_id ").
		Joins("left join districts on districts.id = stores.district_id ").
		Where("favorites.user_id = ?", userId).Find(favoriteApis).Error

	if err != nil {
		return nil, err
	}

	return favoriteApis, nil
}

func CreateFavorite(favorite Favorite) error {
	err := config.Conn.Table("favorites").Create(&favorite).Error
	if err != nil {
		return err
	}

	return nil
}

func DeleteFavorite(favorite Favorite) error {
	storeId := favorite.StoreId
	userId := favorite.UserId

	err := config.Conn.Table("favorites").Where("user_id = ? and store_id = ?", userId, storeId).Delete(&favorite).Error
	if err != nil {
		return err
	}
	return nil
}
