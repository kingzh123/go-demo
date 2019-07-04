package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var (
	once = sync.Once{}
	c = make(chan int)
)

func main() {
	TempSyncPool()
	//onces()
}

func TempSyncPool(){
	rand.Seed(time.Now().UnixNano())
	sp := sync.Pool{
		New: func() interface{} {
			i := make([]int, 10)
			return i;
		},
	}
	//从池中获得数据
	item := sp.Get()
	fmt.Println(item)
	//修改获得数据
	for i := 0; i < len(item.([]int)); i++ {
		item.([]int)[i] = i;
	}
	fmt.Println("item :", item);
	//把获得数据放回池中
	sp.Put(item)
	item2 := sp.Get()
	for i:=0; i<10 ;i++  {
		item2.([]int)[i] = rand.Intn(100)
	}
	fmt.Println("item2 :", item2)
	//第三次从池中获得数据，但是第二次获取后没有放回池中，所有第三获取的时候获得了新的初始数据
	item3 := sp.Get();
	fmt.Println("item3 :", item3)
}

//sync.Once Do 多次执行 仅会被执行一次
func onces(){
	for i:=0; i<10 ;i++  {
		go once.Do(doOnce)
	}
	<-c
}
func doOnce()  {
	fmt.Println("1")
	c <- 1
}