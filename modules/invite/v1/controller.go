package gortc_invite_v1

import (
	"crypto/rand"
	"encoding/hex"
	models "gortc/models"
	mysql "gortc/services/mysql"
	"io"

	"github.com/kataras/iris/v12"
)

func list(ctx iris.Context) {

	ctx.JSON(iris.Map{})

}

func generate(ctx iris.Context) {
	user := ctx.Values().Get("user").(models.User)
	bytes := make([]byte, 3)
	_, err := io.ReadFull(rand.Reader, bytes)

	if err != nil {
		ctx.StatusCode(500)
		ctx.JSON(iris.Map{
			"message": err.Error(),
		})
		return
	}

	code := hex.EncodeToString(bytes)

	entry := models.InviteCode{
		Code:      code,
		CreatedBy: user.ID,
	}

	result := mysql.Ins().Create(&entry)

	if result.Error != nil {
		ctx.StatusCode(500)
		ctx.JSON(iris.Map{
			"message": result.Error.Error(),
		})
		return
	}

	ctx.JSON(iris.Map{
		"message": "Code generated successfully",
		"data":    entry,
	})
}

func generateBulk(ctx iris.Context) {
	user := ctx.Values().Get("user").(models.User)
	var body generateBulkReq
	ctx.ReadBody(&body)
	if body.Amount > 24 {
		ctx.StatusCode(400)
		ctx.JSON(iris.Map{
			"message": "Max amount exceed",
		})
		return
	}
	codes := make([]models.InviteCode, body.Amount)

	for i := 0; i < body.Amount; i++ {
		bytes := make([]byte, 3)
		io.ReadFull(rand.Reader, bytes)
		codes[i] = models.InviteCode{
			Code:      hex.EncodeToString(bytes),
			CreatedBy: user.ID,
		}
	}

	mysql.Ins().CreateInBatches(&codes, body.Amount)

	ctx.JSON(iris.Map{
		"data": codes,
	})
}
