package models

import "psycho-dad/config"

type File struct {
	Id         int    `json:"fileId"`
	DefaultImg string `json:"defaultImg"`
	Img1       string `json:"img1"`
	Img2       string `json:"img2"`
	StoreId    int    `json:"storeId"`
}

func CreateFiles(file *File) error {
	err := config.Conn.Create(&file).Error
	if err != nil {
		return err
	}

	return nil
}

func UpdateFiles(store_id int, file *File) error {
	err := config.Conn.Where(`store_id = ?`, store_id).Updates(&file).Error
	if err != nil {
		return err
	}

	return nil
}
