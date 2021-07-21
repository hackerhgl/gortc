package main

import (
	"fmt"

	gortc_auth_v1 "gortc/modules/auth/v1"
	mysql_service "gortc/services/mysql"

	env_service "gortc/services/env"

	"github.com/kataras/iris/v12"
)

func main() {
	app := iris.New()
	env_service.Init()
	mysql_service.Connect()
	app.Use(iris.Compression)

	app.Get("/", func(ctx iris.Context) {
		ctx.WriteString("Hot reload")
	})

	// AuthControllerV1()
	gortc_auth_v1.Routes(app)

	app.Listen(":8080")
	fmt.Println("Server running at port :8080")
}
