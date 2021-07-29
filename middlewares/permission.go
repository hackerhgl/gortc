package gortc_middlewares

import (
	"github.com/kataras/iris/v12"

	models "gortc/models"
)

func Permission(allowedRole models.UserRole) iris.Handler {
	return func(ctx iris.Context) {
		user := ctx.Values().Get("user").(models.User)
		userRoleIndex := 0
		allowedRoleIndex := 0

		for index, value := range models.RolesArray {
			if user.Role == value {
				userRoleIndex = index
			}
			if allowedRole == value {
				allowedRoleIndex = index
			}
		}

		if userRoleIndex >= allowedRoleIndex {
			ctx.Next()
		} else {
			ctx.StatusCode(401)
			ctx.JSON(iris.Map{
				"message": "Insufficient role",
			})
		}

	}
}
