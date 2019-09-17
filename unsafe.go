package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

type V struct {
	i int32
	j int64
}

type W struct {
	b byte
	i int32
	j int64
}

func main() {
	WTest()
	//VTest()
	//Int()
	//Float64bits(12.00001)
}

func ZTest() {

}

//计算结构体占用内存空间
//结构体 成员变量占用内存为13个字节 时间占用了16个字节 原因：结构体占用空间大小必须是4的倍数，不足的话需要补全
func WTest()  {
	w := new(W)
	fmt.Println(unsafe.Sizeof(*w)) //计算 struct w 所占用的内存空间 16
	fmt.Println(unsafe.Alignof(w.b)) //计算 w b 占用的内存空间 1
	fmt.Println(unsafe.Alignof(w.i)) //计算 w i 占用的内存空间 4
	fmt.Println(unsafe.Alignof(w.j)) //计算 w j 占用的内存空间 8
	var b byte = 'A'
	fmt.Printf("%c\n", b) //输出b的值
	fmt.Println(unsafe.Alignof(b)) //计算b的占用内存空间 1
	c := []byte{'1'}
	fmt.Println(unsafe.Alignof(c)) //计算byte数组的占用内存空间 8
	fmt.Println("Offsetof:", unsafe.Offsetof(w.b)) //w.b的偏移量
	fmt.Println("Offsetof:", unsafe.Offsetof(w.i)) //w.i的偏移量
	fmt.Println("Offsetof:", unsafe.Offsetof(w.j)) //w.j的偏移量
}

//unsafe.Pointer 变量转换成指针类型
//结构体的内存地址和第一个成员变量的内存地址是一样的
//指针类型不能进行运算 所以要转换 uintptr 在进行指针计算 例如：偏移量等
func VTest() {
	v := new(V)
	i32 := (*int32)(unsafe.Pointer(v))
	*i32 = int32(99) //给内存地址重新复制 实际是给结构体第一个成员变量赋值
	fmt.Println(i32) //变量v的内存地址 实际32和64位是一样的
	fmt.Println(v.i) //变更后的i值
	j := (*int64)(unsafe.Pointer(uintptr(unsafe.Pointer(v)) + uintptr(unsafe.Sizeof(int(0))))) //计算v.j偏移量
	fmt.Println(j) //v.j的内存地址
	*j = int64(100) //v.j重新赋值
	fmt.Println(*j) //v.j的值
}

//unsafe.Pointer 是对指针进行类型转换的方法
func Float64bits(f float64) {
	//原始类型
	fmt.Println(reflect.TypeOf(unsafe.Pointer(&f)))  //unsafe.Pointer
	//转换后类型
	fmt.Println(reflect.TypeOf((*uint64)(unsafe.Pointer(&f))))  //*uint64
}

func Int() {
	v1 := uint(12)
	v2 := int(13)
	fmt.Println(reflect.TypeOf(v1)) //uint
	fmt.Println(reflect.TypeOf(v2)) //int
	fmt.Println(reflect.TypeOf(&v1)) //*uint
	fmt.Println(reflect.TypeOf(&v2)) //*int
	p := (*uint)(unsafe.Pointer(&v2)) // *int to *uint
	fmt.Println(p)
}













