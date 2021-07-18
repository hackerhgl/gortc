package gortc_auth_v1

import "github.com/kataras/iris/v12"

func LogIn(ctx iris.Context) {
	ctx.WriteString("Login")

}
