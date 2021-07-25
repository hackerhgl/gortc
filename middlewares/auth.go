package gortc_middlewares

import (
	"encoding/json"
	"errors"
	"strings"

	"github.com/kataras/iris/v12"
	"gorm.io/gorm"

	models "gortc/models"
	jwt "gortc/services/jwt"
	mysql "gortc/services/mysql"
)

type JWTDecoded struct {
	ID uint `json:"id"`
}

func Auth(strict bool) iris.Handler {
	return func(ctx iris.Context) {
		rawToken := ctx.GetHeader("authentication")
		if rawToken == "" {
			ctx.StatusCode(401)
			ctx.JSON(iris.Map{
				"mesage": "Un Authorized",
			})
			return
		}
		split := strings.Split(rawToken, " ")
		if split[0] != "JWT" || len(split) != 2 {
			ctx.StatusCode(400)
			ctx.JSON(iris.Map{
				"mesage": "Invalid token",
			})
			return
		}
		decoded, err := jwt.Decode(split[1])
		if err != nil {
			ctx.StatusCode(400)
			ctx.JSON(iris.Map{
				"mesage": "Error validating the token",
			})
			return
		}
		var parsed JWTDecoded
		err = json.Unmarshal(decoded, &parsed)
		if err != nil {
			ctx.StatusCode(400)
			ctx.JSON(iris.Map{
				"mesage": "Unknown error",
			})
			return
		}

		var user models.User

		result := mysql.Ins().Where("id = ?", parsed.ID).First(&user)

		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			ctx.StatusCode(404)
			ctx.JSON(iris.Map{
				"message": "User not found",
			})
			return
		}

		if strict && !user.IsVerified {
			ctx.StatusCode(401)
			ctx.JSON(iris.Map{
				"message": "Please verify your account",
			})
			return
		}

		ctx.Values().Set("user", user)

		ctx.Next()
	}
}
