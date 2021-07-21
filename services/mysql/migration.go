package gortc_mysql_service

import (
	models "gortc/models"
)

func Migration() {
	db.AutoMigrate(&models.User{})
}
