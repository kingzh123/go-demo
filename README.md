# go-demo

GO 命令

go build
go install
go run

#获得 go 包
go get

#查看 go 文件方法的汇编程序 需要先执行 go build 生成 main 文件
go tool objdump -s "main.sayhallo" main