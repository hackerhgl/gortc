package gortc_auth_v1

import (
	middlewares "gortc/middlewares"

	"github.com/kataras/iris/v12"
)

func Routes(app *iris.Application) {
	v1 := app.Party("/auth/v1")
	v1.Post("/log_in", logIn)
	v1.Post("/sign_up", signUp)
	v1.Post("/user_profile", userProfile).Use(middlewares.Auth(false))
	v1.Post("/verification", verification).Use(middlewares.Auth(false))
}
