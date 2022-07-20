package models

import (
	"errors"
	"fmt"
	"psycho-dad/config"
	"psycho-dad/proto"

	"gorm.io/gorm"
)

type Store struct {
	Id                 int     `json:"storeId"`
	Name               string  `json:"storeName"`
	Introduce          string  `json:"introduce"`
	TypeId             int     `json:"typeId"`
	CountyId           int     `json:"countyId"`
	DistrictId         int     `json:"districtId"`
	Address            string  `json:"address"`
	Phone              string  `json:"phone"`
	Mobile             string  `json:"mobile"`
	Email              string  `json:"email"`
	UserId             int     `json:"userId"`
	IsDads             bool    `json:"isDads"`
	IsDadsRecommend    bool    `json:"isDadsRecommend"`
	WebSite            string  `json:"webSite"`
	Lon                float64 `json:"lon"`
	Lat                float64 `json:"lat"`
	Dis                float64 `json:"dis" gorm:"-"`
	IsClosePermanently bool    `json:"isClosePermanently"`
	PageIndex          int
	PageSize           int
}

func GetAllStores(c *Store) ([]Store, proto.Paging) {

	stores := []Store{}
	db := config.Conn.Model(&stores)

	if c.Id != 0 {
		db = db.Where("id = ?", c.Id)
	}
	if c.Name != "" {
		db = db.Where("name like ?", "%"+c.Name+"%")
	}
	if c.TypeId != 0 {
		db = db.Where("type_id = ?", c.TypeId)
	}
	if c.CountyId != 0 {
		db = db.Where("county_id = ?", c.CountyId)
	}
	if c.TypeId != 0 {
		db = db.Where("type_id = ?", c.TypeId)
	}
	if c.DistrictId != 0 {
		db = db.Where("district_id = ?", c.DistrictId)
	}
	if c.UserId != 0 {
		db = db.Where("user_id = ?", c.UserId)
	}
	if c.IsDads {
		db = db.Where("is_dads = ?", c.IsDads)
	}
	if c.IsDadsRecommend {
		db = db.Where("is_dads_recommend = ?", c.IsDadsRecommend)
	}
	if c.IsClosePermanently {
		db = db.Where("is_close_permanently = ?", c.IsClosePermanently)
	}

	page := proto.Paging{}
	page.PageIndex = int64(c.PageIndex)
	page.PageSize = int64(c.PageSize)

	FindPage(db, &stores, &page)

	return stores, page
}

func GetStoreById(storeId int) (*Store, error) {
	store := &Store{}
	err := config.Conn.Where("id = ?", storeId).First(store).Error
	if err != nil {
		return nil, err
	}

	return store, nil
}

func CreateStore(store Store) error {
	err := config.Conn.Create(&store).Error

	if err != nil {
		return err
	}
	return nil
}

func DeleteStore(storeId int) string {
	var store Store
	err := config.Conn.Where("id = ?", storeId).Delete(&store).Error

	if err != nil {
		return "Error" + err.Error()
	}
	return "DELETE SUCCESSFUL"
}

func UpdateStore(storeId int, store Store) error {
	err := config.Conn.Where("id = ?", storeId).Save(store).Error
	if err != nil {
		return err
	}
	return nil
}

func FindPage(db *gorm.DB, dest interface{}, p *proto.Paging) error {
	if p == nil {
		return errors.New("paging is nil")
	} else {
		err := db.Count(&p.AllCount).Error
		fmt.Println(p)
		if err != nil {
			return err
		}

		err = db.Limit(int(p.PageSize)).Offset(int((p.PageIndex - 1) * p.PageSize)).Find(dest).Error
		if err != nil {
			return err
		}
	}

	return nil
}
