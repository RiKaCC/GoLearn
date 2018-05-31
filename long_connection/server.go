package main

import (
	"fmt"
	"net"
)

func main() {
	l, err := net.Listen("tcp", "localhost:8982")
	if err != nil {
		fmt.Println("listen error:", err)
	}

	fmt.Println("waiting for the message")
	for {
		conn, err := l.Accept()
		if err != nil {
			continue
		}

		fmt.Println(conn.RemoteAddr().String(), "connect success")
		// 处理连接
		handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	buffer := make([]byte, 2048)

	for {
		n, err := conn.Read(buffer)
		if err != nil {
			fmt.Println(conn.RemoteAddr().String(), "connect error:")
			return
		}

		fmt.Println(conn.RemoteAddr().String(), "receive data string:\n", string(buffer[:n]))
	}
}
