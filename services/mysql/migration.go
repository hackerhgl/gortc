package gortc_mysql_service

import (
	models "gortc/models"
)

func migration() {
	db.AutoMigrate(&models.User{})
	db.Migrator().AlterColumn(&models.User{}, "Role")
}
