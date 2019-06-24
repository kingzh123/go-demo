package main

import (
	"fmt"
	"reflect"
)

type test1 struct {
	Name string
}

type test2 struct {
	Sex int
}

type Users struct {
	Id int
	Name string
	Age string
}

func (u Users) UserFunc () {
	fmt.Println("this is UserFunc func")
}

func (u Users) UserFuncArgs (s string) {
	fmt.Println("this is UserFuncArgs", s)
}

func (u *Users) UserNoFunc() {
	fmt.Println("reflect no method")
}

func main() {
	u := Users{1, "king", "20"}
	CallFunc(u)
	reflectMethod(u)
	reflectField(u)
}

//通过反射调用方法
func CallFunc(types interface{}) {
	v := reflect.ValueOf(types)
	//有参数的调用
	method := v.MethodByName("UserFuncArgs")
	args := []reflect.Value{reflect.ValueOf("uuuuuuuuuuuuuuu")}
	method.Call(args)
	//无参数调用
	m2 := v.MethodByName("UserFunc")
	args = make([]reflect.Value, 0)
	m2.Call(args)
}

//反射类型的键值及相关信息
func reflectField(types interface{}) {
	//获得 interface 的 type
	t := reflect.TypeOf(types)
	//获得 interface 的 value
	v := reflect.ValueOf(types)
	//interface 的类型
	fmt.Println("interface type is ", t.Name())
	//遍历 interface
	for i:=0;i < t.NumField();i++ {
		fmt.Print("key is", t.Field(i).Name)
		fmt.Print(",")
		fmt.Print("type is ", t.Field(i).Type)
		fmt.Print(",")
		fmt.Println("value is ", v.Field(i).Interface())
	}
}

//获得类型的方法
//如果 struct 定义的方法 是带 * 的类型，reflect 将不能把该方法反射出来 【func (u *Users) UserNoFunc()】
func reflectMethod (types interface{}) {
	t := reflect.TypeOf(types)
	for i:=0; i<t.NumMethod(); i++ {
		fmt.Println("method is ", t.Method(i).Name)
		fmt.Println("type is ", t.Method(i).Type)
	}
}

func valueO() {
	i := 45
	j := "king"
	fmt.Println("i type is ", reflect.TypeOf(i))
	fmt.Println("i value is ", reflect.ValueOf(i))
	fmt.Println("j type is ", reflect.TypeOf(j))
	fmt.Println("j value is ", reflect.ValueOf(j))
}

func typeO() {
	t := test1{"test"}
	tt := reflect.ValueOf(t)
	fmt.Println("t type is ", reflect.TypeOf(t))
	fmt.Println("tt is ", tt)
	fmt.Println("t value is ", reflect.ValueOf(t))
}