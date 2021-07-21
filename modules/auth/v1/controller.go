package gortc_auth_v1

import (
	"crypto/rand"
	"encoding/hex"
	"errors"
	"fmt"
	models "gortc/models"
	env "gortc/services/env"
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
