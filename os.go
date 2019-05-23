package main

import (
	"errors"
	"fmt"
	"os"
	"path"
	"runtime"
)

func main() {
	//创建文件夹
	dirPath := "/Users/fenghuacaijing/go/src/go-demo/xx"
	createDir(dirPath)
	os.Exit(0)
	//获得当前文件路径
	fileInfo, ok := getPath()
	//是否是文件夹
	b, err := isDir(fileInfo)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%t\n", b)
	if ok {
		b := isExist(fileInfo)
		if !b {
			fmt.Printf("文件夹或文件不存在")
		} else {
			fmt.Printf("文件或文件夹存在！")
		}
	}
}

//创建文件夹
func createDir(path string) {
	if !isExist(path) {
		//MkdirAll 创建多级文件夹目录 Mkdir:创建单个目录 一般创建日期文件夹时 使用前者
		err := os.MkdirAll(path, os.ModePerm) //os.Mkdir(path, 0655)
		if err != nil {
			panic(err)
		}
	}
}

//给定的文件或文件夹是否存在
func isExist(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		return false
	}
	if os.IsNotExist(err) {
		return false
	}
	return true
}

//判断是否是文件夹
func isDir(path string) (b bool, err2 error) {
	//验证文件或文件夹是否存在
	if !isExist(path) {
		return false, errors.New("file or dir is not exist!")
	}
	//文件或文件夹统计信息
	fileInfo, err := os.Stat(path)
	if err != nil {
		return false, err
	}
	//是否是文件夹
	if fileInfo.IsDir() {
		return true, nil
	}
	return false, errors.New("custom err: not dir")
}

//获得当前文件路径
func getPath() (pathInfo string, ok bool) {
	//后的当前路径+文件名
	_, filename, _, ok := runtime.Caller(1)
	var filePath string
	if ok {
		//获得文件路径
		filePath = path.Dir(filename)
		fmt.Printf("filename is %s\n", filename)
	} else {
		panic("runtime is failed")
	}
	fmt.Printf("file path is %s\n", filePath)
	return filePath, ok
}
