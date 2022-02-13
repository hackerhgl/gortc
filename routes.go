package main

import (
	gortc_auth_v1 "gortc/modules/auth/v1"
	gortc_invite_v1 "gortc/modules/invite/v1"

	"github.com/kataras/iris/v12"
)

func routes(app *iris.Application) {
	app.Get("/", func(ctx iris.Context) {
		ctx.WriteString("Home")
	})

	gortc_auth_v1.Routes(app)
	gortc_invite_v1.Routes(app)
}
