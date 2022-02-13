package gortc_auth_v1

import (
	"errors"
	models "gortc/models"
	jwt "gortc/services/jwt"
	mysql "gortc/services/mysql"
	gortc_utils "gortc/utils"
	"time"

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

	if !gortc_utils.VerifySaltNHash(user.Password, user.Salt, body.Password) {
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

	hash, salt := gortc_utils.SaltNHash(body.Password)

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

func forgetPasswordSendOTP(ctx iris.Context) {
	var body forgetPasswordSendOTPReq
	var err error = ctx.ReadBody(&body)
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
			"error": "Email doesn't exist",
		})
		return
	} else if !user.IsVerified {
		ctx.StatusCode(403)
		ctx.JSON(iris.Map{
			"error": "Email is not verified",
		})
		return
	}

	var existing models.UserResetPasswordOTP

	result = mysql.Ins().Where("user_id = ?", user.ID).Last(&existing)

	buffer := time.Now().Add(time.Minute * -1)
	expire := time.Now().Add(time.Hour * 4)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) || !existing.IsActive {
		mysql.Ins().Create(&models.UserResetPasswordOTP{
			Code:   gortc_utils.GenHex(3),
			UserID: user.ID,
		})
		// SEND SMTP
		ctx.StatusCode(201)
		ctx.JSON(iris.Map{
			"message": "Code sent to your email",
		})
		return
	} else if existing.IsActive && (buffer.Before(existing.CreatedAt) || buffer.Before(existing.UpdatedAt)) {
		ctx.StatusCode(429)
		ctx.JSON(iris.Map{
			"message": "Please wait 1 minute before resending the code again",
		})
		return
	} else if existing.IsActive && existing.CreatedAt.Before(expire) {
		mysql.Ins().Save(&existing)
		// SEND SMTP
		ctx.StatusCode(201)
		ctx.JSON(iris.Map{
			"message": "Code sent to your email",
		})
		return
	}
}
