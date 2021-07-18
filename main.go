package main

import (
	"fmt"

	gortc_auth_v1 "gortc/modules/auth/v1"

	"github.com/kataras/iris/v12"
)

func main() {
	app := iris.New()
	app.Use(iris.Compression)

	app.Get("/", func(ctx iris.Context) {
		ctx.WriteString("Hot reload")
	})

	// AuthControllerV1()
	gortc_auth_v1.Routes(app)

	app.Listen(":8080")
	fmt.Println("Server running at port :8080")
}
