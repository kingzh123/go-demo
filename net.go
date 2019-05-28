package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	reqGetSetHead()
	//reqPostForm("https://ihaogu.tougub.com/e/extend/ahgapi/index.php?m=stocks&c=list")
	//reqPost("https://ihaogu.tougub.com/e/extend/ahgapi/index.php?m=stocks&c=list&classid=38", "application/x-www-form-urlencoded")
	//reqGet("https://ihaogu.tougub.com/e/extend/ahgapi/index.php?m=stocks&c=list&classid=38")
}

//控制HTTP客户端标题，重定向策略和其他设置， get 请求 TODO 没有实现目的
func reqGetSetHead() {
	//要是设置header信息需要创建 http client
	client := &http.Client{
		CheckRedirect: redirectPolicyFunc,
	}
	req, _ := http.NewRequest("GET", "https://ihaogu.tougub.com/e/extend/ahgapi/index.php?m=stocks&c=list&classid=38", nil)
	req.Header.Add("If-None-Match", `W/"wyzzy"`)
	resp, _ := client.Do(req)
	content, _ := ioutil.ReadAll(resp.Body)
	fmt.Printf("%s\n", content)
}

func redirectPolicyFunc(req *http.Request, via []*http.Request) error {
	return nil
}

//get 请求
func reqGet(url string){
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	content, err := ioutil.ReadAll(resp.Body)
	fmt.Printf("%s\n", content)
}

//post 请求
func reqPost(url string, contentType string){
	resp, err := http.Post(url, contentType, strings.NewReader("name=lee"))
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	fmt.Printf("%s\n", body)
}

//post 表单请求
func reqPostForm(urlPath string) {
	data := url.Values{"classid": {"38"}}
	resp, err := http.PostForm(urlPath, data)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	content, err := ioutil.ReadAll(resp.Body)
	fmt.Printf("%s\n", content)
}
