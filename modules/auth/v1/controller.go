package gortc_auth_v1

import (
	"errors"
	models "gortc/models"
	jwt "gortc/services/jwt"
	mysql "gortc/services/mysql"
	utils "gortc/utils"

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
			"message": "Invalid credentials",
		})
		return
	}

	if !utils.VerifySaltNHash(user.Password, user.Salt, body.Password) {
		ctx.JSON(iris.Map{
			"error": "Invalid credentials",
		})
		return
	}

	token, err := jwt.Generate(iris.Map{
		"id": user.ID,
	})

	if err != nil {
		ctx.StatusCode(401)
		ctx.JSON(iris.Map{
			"error": "Error while singin please try again later",
		})

		return

	}

	ctx.JSON(iris.Map{
		"message": "User logged in successfully",
		"user":    user,
		"token":   token,
	})
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

	hash, salt := utils.SaltNHash(body.Password)

	newUser := models.User{
		Email:    body.Email,
		Name:     body.Name,
		Password: hash,
		Salt:     salt,
	}
	mysql.Ins().Create(&newUser)

	ctx.JSON(iris.Map{
		"message": "User registered successfully",
	})
}

func userProfile(ctx iris.Context) {
	user := ctx.Values().Get("user").(models.User)
	ctx.JSON(iris.Map{
		"message": "successfull",
		"user":    user,
	})

}
