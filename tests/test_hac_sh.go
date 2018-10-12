package main

import (
	"crypto/hmac"
	"fmt"
	"encoding/hex"
	"crypto/sha1"
)
func main(){
	h := hmac.New(sha1.New, []byte("Centili"))
	h.Write([]byte("content"))
	fmt.Println(hex.EncodeToString(h.Sum(nil)))
}
