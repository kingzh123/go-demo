package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	//url := "http://www.ptsyy.com/robots.txt"
	url := "http://www.jfinfo.com/robots.txt"
	test, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer test.Body.Close()
	body, _ := ioutil.ReadAll(test.Body)
	fmt.Printf("%s\n", body)
}