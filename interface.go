package main

import "fmt"
//接口可以更灵活的构建程序，只要实现接口的方法就可以使用接口方法，也可以按接口定义的参数进行传参
type InterF interface {
	Add(i int) bool
	Update()
}
type Dba struct {
	Name string
}
//Dba 实现 InterF 的Add方法（如果想实现某个接口必须实现接口内的所有方法）
func (d Dba) Add(i int) bool {
	fmt.Println(i)
	return false
}
//实现InterF接口 Update方法
func (d Dba) Update(){}
func TypeInfo(i InterF) {
	fmt.Printf("Interface 类型 %T ,  值： %v\n", i, i)
}
//空接口(不需要实现方法的接口)
func NonInterface(i interface{}) {
	fmt.Printf("Non Interface 类型 %T ,  值： %v\n", i, i)
}
//断言参数类型
func AssertType(i interface{}){
	_, ok := i.(int)
	if !ok {
		fmt.Println("参数不是一个int类型")
	} else {
		fmt.Println("类型 int 匹配")
	}
}
//分析参数的类型
func AnalysisType(i interface{}) {
	switch i.(type) {
	case string:
		fmt.Println("param is string type")
	case int:
		fmt.Println("param is int type")
	case Dba:
		fmt.Println("param is Dba type")
	default:
		fmt.Println("param is Unknown type")
	}
}
func main() {
	inter := new(Dba)
	inter.Add(1)
	TypeInfo(inter)
	NonInterface(2)
	AssertType(2)
	AnalysisType(1)
	AnalysisType("asdfaf")
	AnalysisType(Dba{})
	AnalysisType(21.1)
}