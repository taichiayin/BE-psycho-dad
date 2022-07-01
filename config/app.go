package config

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Conn *gorm.DB

var err error

func StartUp() {
	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	dsn := "root:abc123456@tcp(127.0.0.1:3306)/psychoDad?charset=utf8mb4&parseTime=True&loc=Local"
	Conn, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
}
