package main

import (
	"fmt"
	"sort"
)
type Person struct {
	Name string
	Age int
}
type personSlice []Person
type SortBy func(p, q *Person) bool
//
type PersonWrapper struct {
	people [] Person
	by func(p, q * Person) bool
}
//重写 Len() 方法
func (pw PersonWrapper) Len() int {
	return len(pw.people)
}
//重写 Swap() 方法
func (pw PersonWrapper) Swap(i, j int){
	pw.people[i], pw.people[j] = pw.people[j], pw.people[i]
}
//重写 Less() 方法
func (pw PersonWrapper) Less(i, j int) bool {
	return pw.by(&pw.people[i], &pw.people[j])
}
//排序方法
func SortPerson(peoples []Person, by SortBy){// SortPerson 方法
	sort.Sort(PersonWrapper{peoples, by})
}

func main() {
	//int string 默认为升序
	//int sort
	intArr := []int{1,2,4,5,1,2,3,4,5,-2}
	sort.Ints(intArr)
	fmt.Println(intArr)
	//int 反向排序
	sort.Sort(sort.Reverse(sort.IntSlice(intArr)))
	fmt.Println(intArr)

	//string
	stringArr := []string{"king", "lee", "zhang", "liu"}
	sort.Strings(stringArr)
	fmt.Println(stringArr)
	//string 反向排序
	sort.Sort(sort.Reverse(sort.StringSlice(stringArr)))
	fmt.Println(stringArr)

	//struct 排序 要排序的struct 要实现sort的Len、Swap、Less方法
	s := personSlice{
		{
			Name: "zz",
			Age: 55,
		},
		{
			Name: "AAA",
			Age: 55,
		},
		{
			Name: "BBB",
			Age: 22,
		},
		{
			Name: "CCC",
			Age: 0,
		},
		{
			Name: "DDD",
			Age: 22,
		},
		{
			Name: "EEE",
			Age: 11,
		},
	}
	fmt.Println(s)
	SortPerson(s, func(p, q *Person) bool {
		return p.Name < q.Name    // Name 递增排序
	})
	fmt.Println(s)
	SortPerson(s, func(p, q *Person) bool {
		return p.Age > q.Age    // Name 递增排序
	})
	fmt.Println(s)
}
