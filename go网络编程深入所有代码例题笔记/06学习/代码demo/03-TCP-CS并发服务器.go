package main

import (
	"net"
	"fmt"
	"strings"
)

// 处理客户端数据通信
func HandlerConnect(conn net.Conn)  {
	defer conn.Close()

	// 获取当前客户端的 Addr
	cltAddr := conn.RemoteAddr()
	fmt.Printf("%v客户端连接成功.\n", cltAddr)

	buf := make([]byte, 4096)
	for {		// 在一个客户端连接上，进行循环多次通信
		// 读
		n, err := conn.Read(buf)
		fmt.Println(buf[:n])
		if n == 0 {
			fmt.Printf("服务器检测到 %v 客户端关闭", cltAddr)
			break
		}
		if err != nil {
			fmt.Println("Read err:", err)
			return
		}
		if string(buf[:n]) == "exit\n" || string(buf[:n]) == "exit\r\n" {	//  string(buf[:n-1]) == "exit"
			fmt.Printf("服务器发现 %v 客户端 exit ", cltAddr)
			break
		}

		fmt.Printf("服务器读到 %v 客户端发送数据：%s", cltAddr, string(buf[:n]))
		// 处理数据: 小 -- 大
		UpperWord := strings.ToUpper(string(buf[:n]))
		// 写
		conn.Write([]byte(UpperWord))
	}
}

func main()  {
	listener, err := net.Listen("tcp", "127.0.0.1:8000")
	if err != nil {
		fmt.Println("Listen err:", err)
		return
	}
	defer listener.Close()

	for {		// 循环建立连接
		conn, err := listener.Accept()		// 阻塞
		if err != nil {
			fmt.Println("Accept err:", err)
			return
		}
		// 创建一个新go程，用来与客户端进行数据通信。 并且将 conn 传入。
		go HandlerConnect(conn)
	}
}
