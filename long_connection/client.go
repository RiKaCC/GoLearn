package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	server := "127.0.0.1:8982"
	tcpAddr, err := net.ResolveTCPAddr("tcp4", server)
	if err != nil {
		fmt.Println("resolve error,", err)
		return
	}

	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		fmt.Println("dial TCP error:", err)
		return
	}

	fmt.Println("connect success")
	sender(conn)
}

func sender(conn net.Conn) {
	msg := "long connection"
	i := 0
	for i < 10 {
		conn.Write([]byte(msg))
		time.Sleep(time.Duration(5) * time.Second)
	}
	fmt.Println("over")
}
