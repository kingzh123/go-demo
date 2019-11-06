package main

import (
	"net"
	"net/http"
	"net/rpc"
	"os"
	"time"
)

type Args struct {
	A, B int
}

type Quotient struct {
	Quo, Rem int
}

type Arith int


//计算乘积
func (t *Arith) Multiply(args *Args, reply *int) error {
	time.Sleep(time.Second * 3) //睡三秒，同步调用会等待，异步会先往下执行
	*reply = args.A * args.B
	return nil
}

//计算商和余数
func (t *Arith) Divide(args *Args, quo *Quotient) error {
	time.Sleep(time.Second * 3)
	if args.B == 0 {
		return nil
	}
	quo.Quo = args.A / args.B
	quo.Rem = args.A % args.B
	return nil
}

func main() {
	arith := new(Arith)
	rpc.Register(arith)
	rpc.HandleHTTP()
	l, e := net.Listen("tcp", ":8887")
	if e != nil {
		panic(e)
	}
	go http.Serve(l, nil)
	os.Stdin.Read(make([]byte, 1))
}















