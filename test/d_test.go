package main

import (
	"go-demo/data"
	"testing"
)

const url = "https://github.com/EDDYCJY"

//性能测试 Demo
func TestAxx(t *testing.T) {
	s := datax.Add(url)
	if s == "" {
		t.Errorf("Test.Add error!")
	}
}

//压力测试 Demo 进行压力测试时 命令需要添加 -test.bench=".*"
//压力测试参数需要是 *testing.B
func BenchmarkAdd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		datax.Add(url)
	}
}