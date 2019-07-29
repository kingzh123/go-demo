package main

import (
	"fmt"
	"io/ioutil"
	"os"
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
	/*
		1.声明一个Template对象并解析模板文本
		func New(name string) *Template
		func (t *Template) Parse(text string) (*Template, error)

		2.从html文件解析模板
		func ParseFiles(filenames ...string) (*Template, error)

		3.模板生成器的包装
		template.Must(*template.Template, error )会在Parse返回err不为nil时，调用panic。
		func Must(t *Template, err error) *Template

		t := template.Must(template.New("name").Parse("html"))
		*/

	t, _ := template.ParseFiles("./html/test.tpl")
	t.Execute(os.Stdout, []string{"bgbiao","biaoge"})



	//templateText := OpenFile("index")
	//funcMap := template.FuncMap{
	//	"title": strings.Title,
	//}
	//text := templateText[:]
	//tmpl, err := template.New("test").Funcs(funcMap).Parse(string(text))
	//err = tmpl.Execute(os.Stdout, "the go programming language")
	//if err != nil {
	//	log.Fatalf("execution: %s", err)
	//}
}

func t1()  {
	sweaters := Inventory{"wool", 17}
	tmpl, err := template.New("test").Parse("{{.Count}} items<br /> are made of {{.Material}} {{23 -}} < {{- 45}}")
	if err != nil { panic(err) }
	err = tmpl.Execute(os.Stdout, sweaters)
	if err != nil { panic(err) }
}