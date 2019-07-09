package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"text/template"
)
var (
	defPath string = "./html/"
)
type Inventory struct {
	Material string
	Count    uint
}
func OpenFile(name string) []byte{
	path := defPath + name + ".html"
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	content, err := ioutil.ReadAll(file)
	if err != nil{
		panic(err)
	}
	fmt.Printf("%s\n", content)
	return content
}

func main() {
	templateText := OpenFile("index")
	funcMap := template.FuncMap{
		"title": strings.Title,
	}
	text := templateText[:]
	tmpl, err := template.New("test").Funcs(funcMap).Parse(text)
	err = tmpl.Execute(os.Stdout, "the go programming language")
	if err != nil {
		log.Fatalf("execution: %s", err)
	}
}

func t1()  {
	sweaters := Inventory{"wool", 17}
	tmpl, err := template.New("test").Parse("{{.Count}} items<br /> are made of {{.Material}} {{23 -}} < {{- 45}}")
	if err != nil { panic(err) }
	err = tmpl.Execute(os.Stdout, sweaters)
	if err != nil { panic(err) }
}