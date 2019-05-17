package main

import (
	"bytes"
	"encoding/base64"
	"encoding/gob"
	"encoding/json"
	"fmt"
	"os"
)

type testx struct {
	A int
	B string
	C bool
}

type returns struct {
	Code int
	Msg string
	Data string
}

type returns2 struct {
	Code int
	Msg string
	Da string
}

func main() {
	//json
	re := returns{0, "lllll", "json"}
	jsonText, err := json.Marshal(re)
	if err != nil {
		panic(err)
	}
	fmt.Printf("json:%s\n", jsonText)
	//json 解码
	unjson1 := returns{}
	//unjson2 := returns2{} //json解码的时候会根据 struct 字段进行填值，字段相同写入，不同则不会被赋值
	json.Unmarshal(jsonText, &unjson1)
	fmt.Printf("unjson:%#v\n", unjson1)

	os.Exit(1)
	//base64
	base64Text := base64Encrypt([]byte("hello world"))
	fmt.Printf("base64: %s\n", base64Text)
	msg := base64Decrypt(base64Text)
	fmt.Printf("unbase64 msg: %s\n  ", msg)

	//gob
	t := testx{A:1,B:"test",C:true}
	byteBuffer := gobEncrypt(t)
	fmt.Printf("gob编码的字节流：%#v\n", byteBuffer)
	tx := gobDecrypte(byteBuffer)
 	fmt.Println("gob解码数据", tx)

}

//base64编码
func base64Encrypt(msg []byte) string {
	base64Text := base64.StdEncoding.EncodeToString(msg)
	return base64Text
}

//base64解码
func base64Decrypt(base64Text string) string {
	msg, _:= base64.StdEncoding.DecodeString(base64Text)
	return string(msg)
}


//gob是go语言自定义的编码格式，其性能要高于json、xml，go本身可以对此编码解码成 struct结构体，一般go和go进行数据通信的话，官方建议使用gob进行解码、编码。
//gob 编码
func gobEncrypt(str testx) bytes.Buffer{
	//字节buff
	var network bytes.Buffer
	//创建gob编码器 传递一个 bytes.buffer 类型的指针
	en := gob.NewEncoder(&network)
	//对 struct 进行编码
	err := en.Encode(str)
	if err != nil {
		panic(err)
	}
	return network
}
//gob 解码
func gobDecrypte(buffer bytes.Buffer) testx {
	//创建解码器
	de := gob.NewDecoder(&buffer)
	//被赋值的 struct
	t := testx{}
	//解码
	de.Decode(t)
	return t
}