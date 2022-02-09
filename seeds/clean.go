package main

import (
	models "gortc/models"
	mysql "gortc/services/mysql"
)

func Clean() {
	mysql.Ins().Where("id > 0").Delete(&models.InviteCode{})

	mysql.Ins().Where("id > 0").Delete(&models.User{})
}
