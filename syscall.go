package main

import (
	"fmt"
	"syscall"
)

//err linux操作系统执行失败返回的不是go err类型 而是 非零的值
func main()  {
	pwd()
}

func pwd() {
	path, _ := syscall.Getwd()
	fmt.Println(path)
}