# go-demo

GO 命令

go build
go install
go run

#获得 go 包
go get

#查看 go 文件方法的汇编程序 需要先执行 go build 生成 main 文件
go tool objdump -s "main.sayhallo" main

#服务器环境
go build后直接吧包上传到服务器后运行，不需要在服务器上安装go环境，go build后生成的是二进制包

#go test
会执行当前目录下的 *_test.go 文件, 执行文件中TestFunc(t *Testing.T){} 方法
参考：https://studygolang.com/articles/7051

#go test -bench=. -cpuprofile=cpu.prof
通过此命令生成cpu性能报告
go test -run=文件名字 -bench=bench名字 -cpuprofile=生产的cprofile文件名称

#go tool pprof -http=:8080 cpu.prof
此命令可以在web浏览器中查看性能报告

#go tool pprof cpu.prof
通过命令行查看生成的性能报告

#go tool trace trace.out
通过 trace 工具分析程序生成的 trace 数据
一般通过trace，分析协程的性能