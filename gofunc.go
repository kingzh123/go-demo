package main

import (
	"fmt"
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

var ch = make(chan int)

func foo(id int) {
	fmt.Println(id)
	ch <- id
}

func main() {
	//DEMO 4
	//ch1 := make(chan int)
	//go func(s string) {
	//	for i:=0; i<500; i++ {
	//		fmt.Println(s + strconv.Itoa(i))
	//	}
	//	ch1 <- 0 //非缓缓存 channel 加入数据（目的让协程执行后挂起，带释放数据，保证业务正常跑完）
	//}("go func ")
	//<-ch1 //释放管道中的数据保证业务正常执行
	//DEMO 3
	//count := 1000
	//for i:=0; i<count; i++ {
	//	go foo(i)
	//}
	////取出信道中的数据,保证foo正常执行完成
	//for i := 0; i < count; i++ {
	//	<- ch
	//}
	//DEMO 2
	//ch1 := make(chan string)
	//go func() {
	//	ch1 <- "hello world"
	//}()
	//fmt.Println(<-ch1)
	//DEMO 1 利用 sync.WaitGroup 设定执行完成状态和协程数量
	//wg.Add(1)
	//for i:=0;i<10 ;i++ {
	//	cc <- strconv.Itoa(i)
	//}
	//go hello()
	//wg.Wait()
}