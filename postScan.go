package main

import (
	"fmt"
	"net"
)

var(
	ip = "112.5.125.141"
	maxPort int
	c chan int
)
func main() {
	c = make(chan int)
	maxPort = 65535
	for i:=1;i<maxPort;i++  {
		netadd := new(net.TCPAddr)
		netadd.IP = net.ParseIP(ip)
		netadd.Port = i
		go scan(netadd)
	}
	<-c
}

func scan(addr *net.TCPAddr) {
	conn, err:= net.DialTCP("tcp", nil, addr)
	if err != nil {
		//fmt.Println("loss:", addr.Port)
	} else {
		fmt.Printf("%s:%d\n", addr.IP, addr.Port)
		conn.Close()
	}
	c<-1
}