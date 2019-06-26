package main

import (
	"fmt"
	"os"
	"runtime/trace"
)

var (
	g = make(chan int)
)

func main() {
	f, err := os.Create("./data/trace.out")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	_ = trace.Start(f)
	for i:=0 ;i<3000 ;i++  {
		go Ta(i)
	}
	trace.Stop()
	<- g
}

func Ta(i int) {
	fmt.Println(i)
	g <- i
}