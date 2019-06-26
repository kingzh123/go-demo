package main

import (
	"fmt"
	"runtime"
	"runtime/debug"
	"sync"
)

var (
	g = make(chan int, 5)
	s = sync.WaitGroup{}
	fc = make(chan int)
)

func fooo(i int) {
	fmt.Println(i)
	fc <- i
}

func main() {
	for i:=0; i<100; i++ {
		go fooo(i)
	}
	for i:=0; i<100; i++ {
		<-fc
	}
	fmt.Println("GOROOT:", runtime.GOROOT())
	fmt.Println("NumCgoCall:", runtime.NumCgoCall())
	fmt.Println("NumCPU:", runtime.NumCPU()) // cpu 核心数
	fmt.Println("NumGoroutine:", runtime.NumGoroutine()) //当前 goroutine 数量
	debug.PrintStack()
}

