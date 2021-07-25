package gortc_middlewares

import (
	"time"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/jwt"

	models "gortc/models"
	env "gortc/services/env"
	mysql "gortc/services/mysql"
)

var sharedKey, secretKey []byte

func Generate(data iris.Map) (string, error) {
	initKeys()
	// signer := jwt.NewSigner(jwt.HS256, sharedKey, time.Second*30)
	signer := jwt.NewSigner(jwt.HS256, sharedKey, time.Hour*24)
	signer.WithEncryption(secretKey, nil)
	token, err := signer.Sign(data)
	if err != nil {
		return "", err
	}

	tokenString := string(token)

	result := mysql.Ins().Create(&models.AuthToken{
		Token:  tokenString,
		UserID: data["id"].(uint),
	})

	if result.Error != nil {
		return "", err
	}

	return "JWT " + tokenString, nil
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
