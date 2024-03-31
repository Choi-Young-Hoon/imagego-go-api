package util

import (
	"crypto/sha512"
	"encoding/base64"
)

func Sha512(input string) string {
	hash := sha512.New()
	hash.Write([]byte(input))
	input = base64.StdEncoding.EncodeToString(hash.Sum(nil))

	return input
}
