package gortc_auth_v1

import (
	"errors"
	"fmt"
	models "gortc/models"
	mysql "gortc/services/mysql"

	"github.com/kataras/iris/v12"
	"gorm.io/gorm"
)

func logIn(ctx iris.Context) {
	var body logInReq
	err := ctx.ReadBody(&body)
	if err != nil {
		ctx.JSON(iris.Map{
			"error": err.Error(),
		})
		return
	}

	var user models.User
	result := mysql.Ins().Where("email = ?", body.Email).First(&user)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		ctx.JSON(iris.Map{
			"error": "Email not found",
		})
		return
	}

	fmt.Println("uu", user.Email, user)
	ctx.JSON(body)
}

func signUp(ctx iris.Context) {
	var body signUpReq
	err := ctx.ReadBody(&body)
	if err != nil {
		ctx.JSON(iris.Map{
			"error": err.Error(),
		})
		return
	}

	var user models.User
	result := mysql.Ins().Where("email = ?", body.Email).First(&user)
	if result.RowsAffected > 0 {
		ctx.JSON(iris.Map{
			"error": "Email already exist!",
		})
		return
	}

	newUser := models.User{
		Email:    body.Email,
		Name:     body.Name,
		Password: body.Password,
	}
	mysql.Ins().Create(&newUser)

	ctx.JSON(iris.Map{
		"message": "User registered successfully",
	})

}
