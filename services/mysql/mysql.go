package gortc_mysql_service

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func Connect() bool {
	var err error

	dsn := "root:root@tcp(127.0.0.1:3306)/gortc_dev?charset=utf8mb4&parseTime=True&loc=Local"
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println("Failed to connect with MYSQL Database")
		return false
	}

	fmt.Println("Successfully connected to MYSQL")

	return true
}

func Ins() *gorm.DB {
	return db
}
