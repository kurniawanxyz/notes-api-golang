package utils

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func DBConnect() (db *gorm.DB, err error) {
	dsn := "root:2006@tcp(127.0.0.1:3306)/utschool_db?charset=utf8mb4&parseTime=True&loc=Local"
	db, error := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if error != nil {
		log.Fatal("Failed to connect database", err)
		return nil, err
	}

	return db, nil
}
