package models

import (
	"errors"
	"fmt"
	"os"
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
	DefaultImg         string  `json:"defaultImg" gorm:"-"`
}

type StoreApi struct {
	Id                 int     `json:"storeId"`
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
	UserId             int     `json:"userId"`
	IsDads             bool    `json:"isDads"`
	IsDadsRecommend    bool    `json:"isDadsRecommend"`
	WebSite            string  `json:"webSite"`
	Lon                float64 `json:"lon"`
	Lat                float64 `json:"lat"`
	Dis                float64 `json:"dis" gorm:"-"`
	IsClosePermanently bool    `json:"isClosePermanently"`
	DefaultImg         string  `json:"defaultImg"`
	FavoriteId         int     `json:"favoriteId"`
}

type StoreList struct {
	Id                 int     `json:"storeId"`
	Name               string  `json:"storeName" form:"storeName"`
	CountyId           int     `json:"countyId" form:"countyId"`
	CountyName         string  `json:"countyName"`
	DistrictId         int     `json:"districtId" form:"districtId"`
	DistrictName       string  `json:"districtName"`
	TypeId             int     `json:"typeId" form:"typeId"`
	IsClosePermanently bool    `json:"isClosePermanently"`
	DefaultImg         string  `json:"defaultImg"`
	Dis                float64 `json:"dis" gorm:"-"`
	FavoriteId         int     `json:"favoriteId"`
}

func GetAllStores(storeApi *StoreApi, page int, size int, userId string) ([]StoreApi, proto.Paging) {

	storeApis := []StoreApi{}
	db := config.Conn.Table("stores").
		Select("stores.*, files.default_img, types.name as type_name, counties.name as county_name, districts.name as district_name , favorites.id as favorite_id")

	if storeApi.Id != 0 {
		db = db.Where("stores.id = ?", storeApi.Id)
	}
	if storeApi.Name != "" {
		db = db.Where("stores.name like ?", "%"+storeApi.Name+"%")
	}
	if storeApi.TypeId != 0 {
		db = db.Where("stores.type_id = ?", storeApi.TypeId)
	}
	if storeApi.CountyId != 0 {
		db = db.Where("stores.county_id = ?", storeApi.CountyId)
	}
	if storeApi.TypeId != 0 {
		db = db.Where("stores.type_id = ?", storeApi.TypeId)
	}
	if storeApi.DistrictId != 0 {
		db = db.Where("stores.district_id = ?", storeApi.DistrictId)
	}
	if storeApi.UserId != 0 {
		db = db.Where("stores.user_id = ?", storeApi.UserId)
	}
	if storeApi.IsDads {
		db = db.Where("stores.is_dads = ?", storeApi.IsDads)
	}
	if storeApi.IsDadsRecommend {
		db = db.Where("stores.is_dads_recommend = ?", storeApi.IsDadsRecommend)
	}
	if storeApi.IsClosePermanently {
		db = db.Where("stores.is_close_permanently = ?", storeApi.IsClosePermanently)
	}

	// Join files Table
	db.Joins("left join files on files.store_id = stores.id").
		Joins("left join types on types.id = stores.type_id").
		Joins("left join counties on counties.id = stores.county_id ").
		Joins("left join districts on districts.id = stores.district_id ").
		Joins("left join favorites on favorites.store_id = stores.id and favorites.user_id = ?", userId)

	paging := proto.Paging{}
	paging.PageIndex = int64(page)
	paging.PageSize = int64(size)

	FindPage(db, &storeApis, &paging)

	return storeApis, paging
}

//
func FindAllList(storeList *StoreList, page int, size int, userId int) ([]StoreList, proto.Paging) {
	storeLists := []StoreList{}
	db := config.Conn.Table("stores").
		Select("stores.*, files.default_img, counties.name as county_name, districts.name as district_name , favorites.id as favorite_id")

	if storeList.Id != 0 {
		db = db.Where("stores.id = ?", storeList.Id)
	}
	if storeList.Name != "" {
		db = db.Where("stores.name like ?", "%"+storeList.Name+"%")
	}
	if storeList.TypeId != 0 {
		db = db.Where("stores.type_id = ?", storeList.TypeId)
	}
	if storeList.CountyId != 0 {
		db = db.Where("stores.county_id = ?", storeList.CountyId)
	}
	if storeList.TypeId != 0 {
		db = db.Where("stores.type_id = ?", storeList.TypeId)
	}
	if storeList.DistrictId != 0 {
		db = db.Where("stores.district_id = ?", storeList.DistrictId)
	}
	if storeList.IsClosePermanently {
		db = db.Where("stores.is_close_permanently = ?", storeList.IsClosePermanently)
	}
	db.Joins("left join files on files.store_id = stores.id").
		Joins("left join types on types.id = stores.type_id").
		Joins("left join counties on counties.id = stores.county_id ").
		Joins("left join districts on districts.id = stores.district_id ").
		Joins("left join favorites on favorites.store_id = stores.id and favorites.user_id = ?", userId)

	paging := proto.Paging{}
	paging.PageIndex = int64(page)
	paging.PageSize = int64(size)

	FindPage(db, &storeLists, &paging)

	return storeLists, paging
}

func FindById(storeId int, userId int) (*StoreApi, error) {
	storeApi := &StoreApi{}

	err := config.Conn.Table("stores").
		Select("stores.*, files.default_img, types.name as type_name, counties.name as county_name, districts.name as district_name , favorites.id as favorite_id").
		Joins("left join files on files.store_id = stores.id").
		Joins("left join types on types.id = stores.type_id").
		Joins("left join counties on counties.id = stores.county_id ").
		Joins("left join districts on districts.id = stores.district_id ").
		Joins("left join favorites on favorites.store_id = stores.id and favorites.user_id = ?", userId).
		Where("stores.id = ?", storeId).First(storeApi).Error
	// err := config.Conn.Where("id = ?", storeId).First(store).Error
	if err != nil {
		return nil, err
	}

	return storeApi, nil
}

func CreateStore(store Store) error {

	// begin a transaction
	tx := config.Conn.Begin()
	// 防止沒有rollback被卡死
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	err := tx.Create(&store).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	file := &File{}
	file.DefaultImg = store.DefaultImg
	// file.Img1 = store.File.Img1
	// file.Img2 = store.File.Img2
	file.StoreId = store.Id
	err = tx.Create(file).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	err = tx.Commit().Error
	if err != nil {
		tx.Rollback()
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
	// transiction
	tx := config.Conn.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	err := tx.Where("id = ?", storeId).Save(store).Error

	if err != nil {
		tx.Rollback()
		return err
	}

	file := &File{}

	err = tx.Table("files").Where(`store_id = ?`, store.Id).First(file).Error
	fmt.Println("file.DefaultImg", file.DefaultImg)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// files沒有資料，新增
			file.DefaultImg = store.DefaultImg
			file.StoreId = store.Id

			err := tx.Table("files").Create(file).Error

			if err != nil {
				tx.Rollback()
				return err
			}

			err = tx.Commit().Error
			if err != nil {
				tx.Rollback()
				return err
			}

			return nil
		}
		tx.Rollback()
		return err
	}

	// 先刪掉原本的圖片
	_ = os.Remove("." + file.DefaultImg)

	file.DefaultImg = store.DefaultImg
	file.StoreId = store.Id
	// 找到files，更新
	err = tx.Table("files").Where(`store_id = ?`, file.StoreId).Updates(file).Error

	if err != nil {
		tx.Rollback()
		return err
	}

	err = tx.Commit().Error
	if err != nil {
		tx.Rollback()
		return err
	}

	return nil
}

func FindPage(db *gorm.DB, dest interface{}, p *proto.Paging) error {
	if p == nil {
		return errors.New("paging is nil")
	} else {
		err := db.Count(&p.AllCount).Error
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
