package util

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

func Sha256(data string)string{
	hash := sha256.New()
	hash.Write([]byte(data))
	md := hash.Sum(nil)
	mdStr := hex.EncodeToString(md)
	fmt.Println(mdStr)
	return mdStr
}


