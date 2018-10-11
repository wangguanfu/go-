package main

import (
	"net"
	"fmt"
)

func main()  {
	conn, err := net.Dial("tcp", "127.0.0.1: 8007")
	if err != nil {
		fmt.Println("Dial err:", err)
		return
	}
	defer conn.Close()

	// 模拟浏览器，组织 http 请求报文， 写给web服务器
	//requestHttpHeadr := "GET /hello HTTP/1.1\r\nHost:127.0.0.1:8007\r\n\r\n"
	requestHttpHeadr := "GET /hello http/1.1\r\nHost:127.0.0.1:8007\r\n\r\n"
	conn.Write([]byte(requestHttpHeadr))

	// 读取 web服务器回发的 http应答报文
	buf := make([]byte, 4096)
	n, err := conn.Read(buf)
	if err != nil {
		fmt.Println("Read err:", err)
		return
	}
	fmt.Printf("||\n%s||\n", string(buf[:n]))
}
