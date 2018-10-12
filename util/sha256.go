package util

import (
	"encoding/hex"
	"crypto/sha256"
)


func Hmac_Sha1(str string) string {
	b := []byte(str)
	hash := sha256.New()
	hash.Write(b)
	md := hash.Sum(nil)
	mdStr := hex.EncodeToString(md)
	return mdStr
}

