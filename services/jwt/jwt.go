package gortc_middlewares

import (
	"fmt"
	"time"

	"github.com/kataras/iris/v12/middleware/jwt"

	env "gortc/services/env"
)

var sharedKey, secretKey []byte

func Generate(data interface{}) string {
	if sharedKey == nil && secretKey == nil {
		sharedKey = []byte(env.E().JWT.SHARED)
		secretKey = []byte(env.E().JWT.SECRET)
	}
	var signer *jwt.Signer = jwt.NewSigner(jwt.HS256, sharedKey, time.Hour*24)
	signer.WithEncryption(secretKey, nil)

	token, err := signer.Sign(data)

	if err != nil {
		fmt.Println("err", err)
	}

	return "JWT " + string(token)
}

func Verify() bool {

	return true
}
