package models

import "psycho-dad/config"

type Favorite struct {
	Id      int `json:"favoriteId"`
	StoreId int `json:"storeId"`
	UserId  int `json:"userId"`
}

func GetAllFavorites() []Favorite {
	var favorites []Favorite
	config.Conn.Find(&favorites)

	return favorites
}

func GetFavoriteByUser(userId int) []Favorite {
	var favorites []Favorite
	config.Conn.Where("user_id = ?", userId).Find(&favorites)

	return favorites
}

func CreateFavorite(favorite Favorite) string {
	err := config.Conn.Create(&favorite).Error
	if err != nil {
		return "Error :" + err.Error()
	}

	return "CREATE SUCCESSFUL"
}

func DeleteFavorite(favoriteId int) string {
	var favorites []Favorite
	err := config.Conn.Where("id = ?", favoriteId).Delete(&favorites).Error
	if err != nil {
		return "Error :" + err.Error()
	}
	return "DELETE SUCCESSFUL"
}
