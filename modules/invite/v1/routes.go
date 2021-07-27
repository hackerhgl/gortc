package gortc_invite_v1

import (
	middlewares "gortc/middlewares"
	models "gortc/models"

	"github.com/kataras/iris/v12"
)

func Routes(app *iris.Application) {
	admin := app.Party("/admin/invite/v1")
	admin.Use(middlewares.Auth(true), middlewares.Permission(models.RoleAdmin))
	{
		admin.Post("/list", list)
		admin.Post("/generate", generate)
		admin.Post("/generate_bulk", generateBulk)
	}
}
