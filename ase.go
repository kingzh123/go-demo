package main

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"fmt"
	"math/rand"
	"time"
)

func main() {
	key := []byte("1234567890hhhhhh")
	content := []byte("8877767qwe")
	content = padding(content, aes.BlockSize)
	fmt.Printf("需要加密的字符串(填充后)：%s\n", content)
	ciphertext := aesCBCEncrypt(key, content)
	fmt.Printf("加密后的字符串：%x\n", ciphertext)
	deContent := aesCBCDecrypt(key, ciphertext)
	fmt.Printf("解密后的字符串：%s\n", deContent)
	uppad := unPadding(deContent)
	fmt.Printf("解密后的字符串（取消填充）：%s\n", uppad)
}

//aes CBC 解密
//加密的时需要一个 iv 随机字符串参与加密
func aesCBCDecrypt(key []byte, ciphertext []byte) []byte {
	if len(ciphertext)%aes.BlockSize != 0 {
		panic("ciphertext is not a multiple of the block size")
	}
	if len(ciphertext) < aes.BlockSize {
		panic("ciphertext too short")
	}
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}
	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]
	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(ciphertext, ciphertext)
	return ciphertext
}

//aes CBC 加密（对称加密）
//加密是需要 iv(随机字符串) 参与加密的
//加密规则：加密前需要生成由 key 创建的加密块，然后由 加密块 和 iv(随机字符串) 生成 mode，之后在生成密文
//注意：被加密的文件必须是aes.BlockSize的倍数
func aesCBCEncrypt(key []byte, content []byte) []byte {
	if len(content)%aes.BlockSize != 0 {
		panic("加密的密文非aes块的倍数")
	}
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}
	//iv随机数
	iv := []byte(getRandomString(aes.BlockSize))
	//密码文本
	ciphertext := make([]byte, aes.BlockSize+len(content))
	//密文增加随机数
	copy(ciphertext, iv)
	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(ciphertext[aes.BlockSize:], content)
	return ciphertext
}

//pkcs5 填充数据
/**
 * 为不满足 BlockSize 倍数的文本补充剩余字节
 * 原因：aes加密时，被加密的文本必须是 BlockSize的倍数
 */
func padding(content []byte, BlockSize int) []byte {
	//与blocksize倍数差异数
	diffNumber := BlockSize - len(content)%BlockSize
	//生成填充文本
	padtext := bytes.Repeat([]byte{byte(diffNumber)}, diffNumber)
	return append(content, padtext...)
}

//pkcs5 取消填充
func unPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	uppad := origData[:(length - unpadding)]
	return uppad
}

//生成随机字符串
func getRandomString(l int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyz"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < l; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}
