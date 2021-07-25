package gortc_utils

import (
	"crypto/rand"
	"encoding/hex"
	"io"
	"log"

	env "gortc/services/env"

	"golang.org/x/crypto/bcrypt"
)

func SaltNHash(password string) (string, string) {
	bytes := make([]byte, 6)
	_, err := io.ReadFull(rand.Reader, bytes)
	if err != nil {
		log.Fatal(err)
	}
	salt := hex.EncodeToString(bytes)

	hash, err := bcrypt.GenerateFromPassword([]byte(password+salt+env.E().APP.PEPPER), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal(err)
	}

	return string(hash), salt
}

func VerifySaltNHash(hash string, salt string, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password+salt+env.E().APP.PEPPER))
	return err == nil
}
