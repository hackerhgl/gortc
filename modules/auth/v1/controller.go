package gortc_auth_v1

import (
	"encoding/json"

	"github.com/kataras/iris/v12"
)

type Book struct {
	Title string `json:"title"`
}

func LogIn(ctx iris.Context) {
	// var b Book
	// ctx.ReadBody(&b)
	// fmt.Println(b)
	// inter := map[string]interface{}{}
	inter := iris.Map{}
	body, _ := ctx.GetBody()

	json.Unmarshal(body, &inter)

	ctx.JSON(inter)

}
