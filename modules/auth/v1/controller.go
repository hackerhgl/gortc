package gortc_auth_v1

import (
	"github.com/kataras/iris/v12"
)

type Book struct {
	Title string `json:"title"`
}

func LogIn(ctx iris.Context) {
	var body LogInReq
	err := ctx.ReadBody(&body)

	if err != nil {
		ctx.JSON(iris.Map{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(body)

}
