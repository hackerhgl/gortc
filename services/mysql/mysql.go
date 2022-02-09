package gortc_mysql_service

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	env "gortc/services/env"
)

var db *gorm.DB

func Connect() bool {
	var err error
	configs := env.E().MYSQL

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", configs.USER, configs.PASSWORD, configs.HOST, configs.PORT, configs.DATABASE)
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Errorf(err.Error())
		return false
	}

	x, _ := db.DB()
	err = x.Ping()
	if err != nil {
		fmt.Println("MYSQL: PING FAILED")
		return false
	}
	fmt.Println("MYSQL: PING SUCCESS")

	migration()

	return true
}

func Ins() *gorm.DB {
	return db
}
