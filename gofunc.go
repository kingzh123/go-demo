package main

import (
	"fmt"
	"strconv"
	"sync"
)

var wg sync.WaitGroup
var cc = make(chan string, 10)
func hello()  {
	l := len(cc)
	for i:=0;i<l ;i++{
		fmt.Println(<-cc)
	}
	wg.Done()
}
func main() {
	wg.Add(1)
	for i:=0;i<10 ;i++ {
		cc <- strconv.Itoa(i)
	}
	go hello()
	wg.Wait()
}