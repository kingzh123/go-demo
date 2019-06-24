package main

import (
	"fmt"
	"runtime"
	"sync"
)

var (
	g = make(chan int, 5)
	s = sync.WaitGroup{}
)

func main() {
	fmt.Println("GOROOT:", runtime.GOROOT())
	fmt.Println("NumCgoCall:", runtime.NumCgoCall())
	fmt.Println("NumCPU:", runtime.NumCPU())
	s.Wait()
	fmt.Println("NumGoroutine:", runtime.NumGoroutine())
}

