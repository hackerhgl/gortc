package gortc_auth_v1

import (
	"errors"
	models "gortc/models"
	jwt "gortc/services/jwt"
	mysql "gortc/services/mysql"
	utils "gortc/utils"

	"github.com/kataras/iris/v12"
	"gopkg.in/nullbio/null.v4"
	"gorm.io/gorm"
)

func logIn(ctx iris.Context) {
	var body logInReq
	err := ctx.ReadBody(&body)
	if err != nil {
		ctx.StatusCode(400)
		ctx.JSON(iris.Map{
			"error": err.Error(),
		})
		return
	}

	var user models.User
	result := mysql.Ins().Where("email = ?", body.Email).First(&user)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		ctx.StatusCode(404)
		ctx.JSON(iris.Map{
			"message": "Invalid credentials",
		})
		return
	}

	if !utils.VerifySaltNHash(user.Password, user.Salt, body.Password) {
		ctx.StatusCode(404)
		ctx.JSON(iris.Map{
			"error": "Invalid credentials",
		})
		return
	}

	token, err := jwt.Generate(iris.Map{
		"id": user.ID,
	})

	if err != nil {
		ctx.StatusCode(400)
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
		ctx.StatusCode(400)
		ctx.JSON(iris.Map{
			"error": err.Error(),
		})
		return
	}

	var user models.User
	result := mysql.Ins().Where("email = ?", body.Email).First(&user)
	if result.RowsAffected > 0 {
		ctx.StatusCode(409)
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

func verification(ctx iris.Context) {
	user := ctx.Values().Get("user").(models.User)
	if user.IsVerified {
		ctx.StatusCode(400)
		ctx.JSON(iris.Map{
			"message": "User is already verified",
		})
		return
	}

	var body verificationReq
	ctx.ReadBody(&body)
	if len(body.Code) != 6 {
		ctx.StatusCode(400)
		ctx.JSON(iris.Map{
			"message": "invalid status code",
		})
		return
	}

	var inviteCode models.InviteCode
	result := mysql.Ins().Where("code = ?", body.Code).First(&inviteCode)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		ctx.StatusCode(404)
		ctx.JSON(iris.Map{
			"message": "Invite code not found",
		})
		return
	}

	if result.RowsAffected > 0 && inviteCode.RedeemedBy.Valid {
		ctx.StatusCode(409)
		ctx.JSON(iris.Map{
			"message": "Invite code is already redeemed",
		})
		return
	}

	err := mysql.Ins().Transaction(func(tx *gorm.DB) error {
		inviteCode.RedeemedBy = null.UintFrom(user.ID)
		result := tx.Save(&inviteCode)
		if result.Error != nil {
			return result.Error
		}
		user.IsVerified = true
		result = tx.Save(&user)
		if result.Error != nil {
			return result.Error
		}
		return nil
	})

	if err != nil {
		ctx.StatusCode(409)
		ctx.JSON(iris.Map{
			"message": "Error while redeeming the code",
			"reason":  err.Error(),
		})
		return
	}

	ctx.JSON(iris.Map{
		"message": "Invite code redeemed successfully",
	})
}
