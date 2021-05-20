package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8888")
	if err != nil {
		fmt.Println("dial err=", err)
		return
	}
	defer conn.Close()
	fmt.Println("连接成功,conn=", conn)
	reader := bufio.NewReader(os.Stdin)

	for {
		content, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("read string err=", err)
			continue
		}

		if strings.Trim(content, " \r\n") == "exit" {
			fmt.Println("客户端退出...")
			return
		}
		_, err = conn.Write([]byte(content))
		if err != nil {
			fmt.Println("write string err=", err)
			continue
		}
	}
}
