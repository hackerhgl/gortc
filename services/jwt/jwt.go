package gortc_middlewares

import (
	"fmt"
	"time"

	"github.com/kataras/iris/v12/middleware/jwt"

	env "gortc/services/env"
)

var sharedKey, secretKey []byte

func Generate(data interface{}) string {
	initKeys()
	signer := jwt.NewSigner(jwt.HS256, sharedKey, time.Hour*24)
	signer.WithEncryption(secretKey, nil)

	token, err := signer.Sign(data)

	if err != nil {
		fmt.Println("err", err)
	}

	return "JWT " + string(token)
}

func Decode(rawToken string) ([]byte, error) {
	initKeys()
	verifier := jwt.NewVerifier(jwt.HS256, sharedKey)
	verifier.WithDecryption(secretKey, nil)

	token, err := verifier.VerifyToken([]byte(rawToken))

	if err != nil {
		return nil, err
	}

	return token.Payload, nil
}

func initKeys() {
	if sharedKey == nil && secretKey == nil {
		sharedKey = []byte(env.E().JWT.SHARED)
		secretKey = []byte(env.E().JWT.SECRET)
	}
}
