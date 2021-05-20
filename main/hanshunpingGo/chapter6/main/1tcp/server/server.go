package main

import (
	"fmt"
	"net"
)

func process(conn net.Conn) {
	defer conn.Close()
	for {
		content := make([]byte, 1024)
		// 会阻塞等待客户端发送数据
		n, err := conn.Read(content)
		if err != nil {
			fmt.Println(conn.RemoteAddr().String(), ":", "服务端 read string err=", err)
			return
		}
		fmt.Print(conn.RemoteAddr().String(), ":", string(content[:n]))
	}
}

func main() {
	listen, err := net.Listen("tcp", "127.0.0.1:8888")
	if err != nil {
		fmt.Println("main() listen err=", err)
		return
	}
	defer listen.Close()

	fmt.Println("main() 服务器开始监听8888端口...")
	for {
		fmt.Println("main() 服务器等待客户端连接...")
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("main() Accept() err=", err)
		} else {
			fmt.Println("main() 连接成功，conn=", conn, " 客户端的IP=", conn.RemoteAddr().String())
			go process(conn)
		}
	}
}
