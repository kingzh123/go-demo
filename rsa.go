package main

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/hex"
	"encoding/pem"
	"fmt"
	"os"
)

func main() {
	//generatePrivateKey(1028)
	msg := []byte("hello world")
	//使用公钥进行加密
	cipherText := RSA_Encrypt("publicKey.pem", msg)
	fmt.Printf("RSA加密后密文：%X\n", cipherText)
	//对密文生成签名
	sign := RSA_Sign("privateKey.pem", cipherText)
	fmt.Printf("对密文生成的签名：%s\n", sign)
	isSign := VerifySign(cipherText, sign, "publicKey.pem")
	fmt.Printf("签名是否合法：%t\n", isSign)
	//使用私钥解密
	plain := RSA_Decrypt("privateKey.pem", cipherText)
	fmt.Printf("RSA解密后明文：%s\n", plain)
}

//验证签名
//msg rsa 加密后的密文
//sign 签名
//path 公钥路径
func VerifySign(msg []byte, sign string, path string) bool {
	unhex_sign, _ := hex.DecodeString(sign)
	//取得公钥
	publicKey := getRSAPublicKey(path)
	//计算消息散列值
	hash := sha256.New()
	hash.Write(msg)
	bytes := hash.Sum(nil)
	//验证数字签名
	err := rsa.VerifyPKCS1v15(publicKey, crypto.SHA256, bytes, []byte(unhex_sign))
	return err == nil
}

//通过私钥对密文进行签名
//privateKeyPath 私钥路径
//cipherText 密文文本
func RSA_Sign(privateKeyPath string, cipherText []byte) string {
	hashed := sha256.Sum256(cipherText)
	privateKey := getRSAPrivateKey(privateKeyPath)
	sign, err := rsa.SignPKCS1v15(rand.Reader, privateKey, crypto.SHA256, hashed[:])
	if err != nil {
		panic(err)
	}
	return hex.EncodeToString(sign)
}

func getRSAPublicKey(path string) *rsa.PublicKey {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	privateInfo, _ := file.Stat()
	//创建缓存
	buf := make([]byte, privateInfo.Size())
	//读取文件大缓存
	file.Read(buf)
	//pem解码
	block, _ := pem.Decode(buf)
	//x509解码
	publicKeyInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		panic(err)
	}
	publicKey := publicKeyInterface.(*rsa.PublicKey)
	return publicKey
}

func getRSAPrivateKey(path string) *rsa.PrivateKey {
	//打开文件
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	//获取文件内容
	info, _ := file.Stat()
	buf := make([]byte, info.Size())
	file.Read(buf)
	//pem解码
	block, _ := pem.Decode(buf)
	//X509解码
	privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	return privateKey
}

//RSA解码
func RSA_Decrypt(privateKeyPath string, cipherText []byte) []byte {
	//打开文件
	file, err := os.Open(privateKeyPath)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	//获取文件内容
	info, _ := file.Stat()
	buf := make([]byte, info.Size())
	file.Read(buf)
	//pem解码
	block, _ := pem.Decode(buf)
	//X509解码
	privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		panic(err)
	}
	//对密文进行解密
	plainText, _ := rsa.DecryptPKCS1v15(rand.Reader, privateKey, cipherText)
	//返回明文
	return plainText
}

//RSA加密
func RSA_Encrypt(publicKeyPath string, msg []byte) []byte {
	file, err := os.Open(publicKeyPath)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	privateInfo, _ := file.Stat()
	//创建缓存
	buf := make([]byte, privateInfo.Size())
	//读取文件大缓存
	file.Read(buf)
	//pem解码
	block, _ := pem.Decode(buf)
	//x509解码
	publicKeyInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		panic(err)
	}
	public := publicKeyInterface.(*rsa.PublicKey)
	//加密明文
	cipherText, err := rsa.EncryptPKCS1v15(rand.Reader, public, msg)
	if err != nil {
		panic(err)
	}
	return cipherText
}

//生成秘钥对（公钥和私钥）
func generatePrivateKey(bits int) {
	privateKey, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		panic(err)
	}
	//将私钥序列化为 PKCS#1 DER 编码形式
	derStream := x509.MarshalPKCS1PrivateKey(privateKey)
	block := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: derStream,
	}
	//创建私钥文件
	file, err := os.Create("privateKey.pem")
	if err != nil {
		panic(err)
	}
	//生成私钥到文件
	err = pem.Encode(file, block)
	if err != nil {
		panic(err)
	}
	//生成公钥
	publicKey := &privateKey.PublicKey
	//将公钥序列化为 DER 编码的 PKIX 格式
	derPkix, err := x509.MarshalPKIXPublicKey(publicKey)
	if err != nil {
		panic(err)
	}
	block = &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: derPkix,
	}
	//创建公钥文件
	file, err = os.Create("publicKey.pem")
	if err != nil {
		panic(err)
	}
	//写入公钥
	err = pem.Encode(file, block)
	if err != nil {
		panic(err)
	}
}
