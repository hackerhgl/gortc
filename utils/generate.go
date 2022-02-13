package gortc_utils

import (
	"crypto/rand"
	"encoding/hex"
	"io"
)

func GenHex(length int) string {
	bytes := make([]byte, length)
	io.ReadFull(rand.Reader, bytes)

	return hex.EncodeToString(bytes)
}
