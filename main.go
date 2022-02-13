package main

import (
	"fmt"

	mysql "gortc/services/mysql"

	env "gortc/services/env"

	"github.com/kataras/iris/v12"
)

func main() {
	env.Init()
	app := iris.New()
	mysql.Connect()
	app.Use(iris.Compression)

	routes(app)

	app.Listen(":8080")
	fmt.Println("Server running at port :8080")
}
