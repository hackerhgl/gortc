package main

import (
	"fmt"

	"github.com/kataras/iris/v12"
)

func main() {
	app := iris.New()
	app.Use(iris.Compression)

	app.Get("/", func(ctx iris.Context) {
		ctx.WriteString("Hot reload")
	})

	app.Listen(":8080")
	fmt.Println("Server running at port :8080")
}
