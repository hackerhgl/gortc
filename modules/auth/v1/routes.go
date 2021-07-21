package gortc_auth_v1

import (
	"github.com/kataras/iris/v12"
)

func Routes(app *iris.Application) {
	auth := app.Party("/auth/v1")
	auth.Post("/log_in", logIn)
	auth.Post("/sign_up", signUp)
}
