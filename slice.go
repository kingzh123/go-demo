package main

import "fmt"

func main() {
	//变量 a 和 b 共享底层同一个数组 只是取值范围不一样
	a := []string{"a","b","c","d","e"}
	b := a[2:5]
	b[0] = "king"
	fmt.Printf("%+v\n", a)
	fmt.Println(b)
}