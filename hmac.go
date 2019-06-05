package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

func main() {
	key := []byte("abcd")
	content := []byte("content")
	hmacmsg := []byte(hmacSha256(key, content))
	fmt.Printf("%s\n", hmacmsg)
	check := checkHmacSha256(hmacmsg, content, key)
	fmt.Printf("%t\n", check)
}

//生成签名
func hmacSha256(key []byte, content []byte) string {
	h := hmac.New(sha256.New, key)
	h.Write(content)
	return hex.EncodeToString(h.Sum(nil))
}

//验证hmac签名
func checkHmacSha256(srcHmac []byte, content []byte, key []byte) bool {
	dscHmac := []byte(hmacSha256(key, content))
	fmt.Printf("原始hmac：%s\n", srcHmac)
	fmt.Printf("目标hmac：%s\n", dscHmac)
	return hmac.Equal(srcHmac, dscHmac)
}
