package gortc_auth_v1

import (
	"crypto/rand"
	"encoding/hex"
	"errors"
	models "gortc/models"
	env "gortc/services/env"
	jwt "gortc/services/jwt"
	mysql "gortc/services/mysql"

	"io"
	"log"

	"github.com/kataras/iris/v12"
	"golang.org/x/crypto/bcrypt"
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

	if !verifySaltNHash(user.Password, user.Salt, body.Password) {
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

	hash, salt := saltNHash(body.Password)

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

func saltNHash(password string) (string, string) {
	bytes := make([]byte, 6)
	_, err := io.ReadFull(rand.Reader, bytes)
	if err != nil {
		log.Fatal(err)
	}
	salt := hex.EncodeToString(bytes)

	hash, err := bcrypt.GenerateFromPassword([]byte(password+salt+env.E().APP.PEPPER), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal(err)
	}

	return string(hash), salt
}

func verifySaltNHash(hash string, salt string, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password+salt+env.E().APP.PEPPER))
	return err == nil
}
