package gortc_auth_v1

import (
	"github.com/kataras/iris/v12"
)

func Routes(app *iris.Application) {
	auth := app.Party("/auth/v1")
	auth.Post("/login", logIn)
}
