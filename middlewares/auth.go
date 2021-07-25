package gortc_middlewares

import (
	"fmt"

	"github.com/kataras/iris/v12"
)

func Auth() iris.Handler {
	return func(ctx iris.Context) {
		token := ctx.GetHeader("authentication")

		fmt.Println(token)
		ctx.Next()
	}
}
