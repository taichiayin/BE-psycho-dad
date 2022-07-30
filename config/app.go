package config

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var Conn *gorm.DB

var err error

func StartUp() {
	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	// local
	dsn := "root:taichi0804@tcp(127.0.0.1:3306)/psychodad?charset=utf8mb4&parseTime=True&loc=Local"
	// deploy
	// dsn := "root:taichi0804@tcp(127.0.0.1:3306)/psychodad?charset=utf8mb4&parseTime=True&loc=Local"
	Conn, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatal(err)
	}

}
