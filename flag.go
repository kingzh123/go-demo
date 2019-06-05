package main

import (
	"flag"
	"fmt"
	"os"
)

var name string

func init() {
	//flag 引用命令行传入的参数 例: flag -n king 或 flag -n=king
	flag.StringVar(&name, "n", "default", "")
}

func main() {
	//flag用来解析命令行参数 flag.Parse() 执行这个方法可以将参数赋值给引用好的变量
	flag.Parse()
	fmt.Fprint(os.Stderr, "hello world\n")
	fmt.Println(name)
}
