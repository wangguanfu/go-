package main

import (
	"net"
	"fmt"
	"os"
)

func showErr(err error, info string)  {
	if err != nil {
		fmt.Println(info, " err:", err)
		//return		// 返回当前函数调用
		os.Exit(-1)
	}
}

// 获取 浏览器发送的 请求数据包

func main()  {
	listener, err := net.Listen("tcp", "127.0.0.1:8000")
	showErr(err, "Listen")
	defer listener.Close()

	conn, err := listener.Accept()
	showErr(err, "Accept")
	defer conn.Close()

	buf := make([]byte, 4096)
	n, err := conn.Read(buf)
	showErr(err, "read")

	fmt.Printf("||\n%s||\n", string(buf[:n]))
}
