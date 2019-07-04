package main

import (
	"fmt"
	"sync"
	"time"
)

var(
	m *sync.RWMutex
	lock *sync.Mutex
	w []int
)

//Mutex 排它锁 一般用在不确定读或写的代码块中
//RWMutex 读写锁 一般在明确代码中读、写操作的时候 用此类型
//当使用 goroutine（协程）的时候，程序并行执行保证数据正确性 每个协程执行的时候需要锁定代码块
func main() {
	w = make([]int, 3)
	w[0] = 0
	w[1] = 1
	m = new(sync.RWMutex)
	lock = new(sync.Mutex)
	go write(1, 10)
	go write(2, 99)
	go write(3, 88)
	go write(4, 56)
	go read(1)
	go read(2)
	time.Sleep(1*time.Second)
	fmt.Println("last w value", w)
}

func read(i int) {
	lock.Lock()
	fmt.Printf("%d starting\n", i)
	fmt.Printf("%d doing\n", i)
	fmt.Printf("%d over\n", i)
	lock.Unlock()
}

func write(i int, v int) {
	m.Lock()
	fmt.Println("write before w:", w)
	fmt.Println("string write", i)
	fmt.Println("doing write", i)
	w[2] = w[2] +v
	fmt.Println("over write", i)
	fmt.Println("write after w:", w)
	m.Unlock()
}
