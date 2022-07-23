package models

import (
	"psycho-dad/config"
)

type Favorite struct {
	Id      int    `json:"favoriteId"`
	StoreId int    `json:"storeId"`
	UserId  string `json:"userId"`
}

type FavoriteApi struct {
	Id       int    `json:"favoriteId"`
	StoreId  int    `json:"storeId"`
	UserId   string `json:"userId"`
	StoreApi `gorm:"embedded"`
}

func GetAllFavorites() []Favorite {
	var favorites []Favorite
	config.Conn.Find(&favorites)

	return favorites
}

func GetFavoritesByUser(userId int) (*[]FavoriteApi, error) {
	// var favorites []Favorite
	favoriteApis := &[]FavoriteApi{}
	err := config.Conn.Table("Favorites").Select("stores.*", "favorites.*", "files.default_img").
		Joins("left join stores on favorites.store_id = stores.id").
		Joins("left join files on stores.id = files.store_id").
		Where("favorites.user_id = ?", userId).Find(favoriteApis).Error

	if err != nil {
		return nil, err
	}

	return favoriteApis, nil
}

func CreateFavorite(favorite Favorite) error {
	err := config.Conn.Table("Favorites").Create(&favorite).Error
	if err != nil {
		return err
	}

	return nil
}

func DeleteFavorite(favorite Favorite) error {
	storeId := favorite.StoreId
	userId := favorite.UserId

	err := config.Conn.Table("Favorites").Where("user_id = ? and store_id = ?", userId, storeId).Delete(&favorite).Error
	if err != nil {
		return err
	}
	return nil
}
