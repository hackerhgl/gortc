package gortc_auth_v1

import (
	middlewares "gortc/middlewares"

	"github.com/kataras/iris/v12"
)

func Routes(app *iris.Application) {
	auth := app.Party("/auth/v1")
	auth.Post("/log_in", logIn)
	auth.Post("/sign_up", signUp)
	auth.Post("/user_profile", userProfile).Use(middlewares.Auth(false))
	auth.Post("/verification", verification).Use(middlewares.Auth(false))
}
