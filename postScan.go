package main

import (
	"fmt"
	"net"
)

var(
	//ip = "112.5.125.141"
	ip = "www.tougub.com"
	host = "tapp.ihaogu.com"
	maxPort int
	c chan int
)
func main() {
	c = make(chan int)
	ip, _ := net.LookupHost(host)
	maxPort = 65535
	for i:=1;i<maxPort;i++  {
		netadd := new(net.TCPAddr)
		netadd.IP = net.ParseIP(ip[0])
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