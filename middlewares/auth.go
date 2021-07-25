package gortc_middlewares

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/kataras/iris/v12"

	jwt "gortc/services/jwt"
)

type JWTDecoded struct {
	ID uint `json:"id"`
}

func Auth() iris.Handler {
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

		fmt.Println(parsed, "pp")

		ctx.Next()
	}
}
