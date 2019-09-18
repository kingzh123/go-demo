package main

import "fmt"

func main() {
	A("string",1,1,3,4,5,6,7,8,9,9)
}

func A(s string, n...int) {
	for _, v := range n {
		fmt.Println(v)
	}
	fmt.Println(s)
}