package main

import (
	"encoding/json"
	"fmt"
	"github.com/robfig/cron"
	"hgp/custom/redis"
	"hgp/net"
	"log"
	"net/http"
	"os"
	"strings"
)

//
var logger *log.Logger
var file *os.File
var err error

var c *cron.Cron
var r *redis.Redis
var ro *route.Route

//结构体 规范：名称首字母大写
type test struct {
	Id    int
	Name  string
	Title string
}

func init() {
	file, err = os.OpenFile("test.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 666)
	if err != nil {
		logger.Fatal(err)
	}
	logger = log.New(file, "", log.LstdFlags)
	logger.SetPrefix("Test- ") // 设置日志前缀
	logger.SetFlags(log.LstdFlags | log.Lshortfile)
}

func sayhallo(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Println(r.Form)
}

// net/http
func helloworld(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	//fmt.Println("Form: ", r.Form)
	//fmt.Println("Path: ", r.URL.Path)
	//fmt.Println(r.Form["a"])
	//fmt.Println(r.Form["b"])
	for k, v := range r.Form {
		fmt.Println(k, "=>", v, strings.Join(v, "-"))
	}
	fmt.Fprint(w, "hello world!")
}

//站内重定向
func httpRedirect(w http.ResponseWriter, r *http.Request) {
	//解析参数 不执此方法参数不能被解析
	r.ParseForm()
	fmt.Println(r.Form.Get("a"))
	url := "/say"
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

//json 编码
func jsons() {
	//切片
	arr := make([]redis.Redis, 0)
	re := &redis.Redis{}
	re.Name = "king"
	//加入到切片
	arr = append(arr, *re)
	fmt.Println(arr)
	if err != nil {
		fmt.Println(err)
	}
	re.Name = "s"
	//json编码
	data, _ := json.Marshal(re)
	fmt.Println(string(data))
}

/*
 * json 解码
 */
func dejson() {
	var s = `{"Name":"s"}`
	redis := redis.Redis{}
	json.Unmarshal([]byte(s), &redis)
	fmt.Println(redis)
}

func main() {
	//dejson()
	//jsons()
	http.HandleFunc("/say", sayhallo)
	http.HandleFunc("/", helloworld)
	http.HandleFunc("/redirect", httpRedirect)
	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		fmt.Println(err)
	}
	//i := 0
	//c = cron.New()
	//r.Hredis()
	//spec := "*/2 * * * * ?"
	//c.AddFunc(spec, func() {
	//	logger.Println("cron running:", i)
	//})
	//c.Start()
	//select {}
}
