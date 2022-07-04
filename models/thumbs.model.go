package models

import (
	"psycho-dad/config"
)

type Thumb struct {
	Id      int `json:"thumbId"`
	StoreId int `json:"storeId"`
	UserId  int `json:"userId"`
}

func GetThumbCountByStore(stordId int) int64 {
	var thumbs []Thumb
	count := config.Conn.Where("store_id = ?", stordId).Find(&thumbs).RowsAffected

	return count
}

func CreateThumb(thumb Thumb) string {
	err := config.Conn.Create(&thumb).Error

	if err != nil {
		return "Error :" + err.Error()
	}
	return "Like!!"
}

func DeleteThumb(thumbId int) string {
	thumb := Thumb{}
	err := config.Conn.Where("id = ? ", thumbId).Delete(&thumb).Error

	if err != nil {
		return "Error :" + err.Error()
	}
	return "Unlike!!"
}
