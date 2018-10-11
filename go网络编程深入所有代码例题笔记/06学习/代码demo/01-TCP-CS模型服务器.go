package main

import (
	"net"
	"fmt"
)

func main() {
	// 创建用于监听的 socket
	listener, err := net.Listen("tcp", "127.0.0.1:7020")
	if err != nil {
		fmt.Println("Listen err:", err)
		return
	}
	fmt.Println("监听套接字，创建成功。。。")
	// 服务器结束前关闭 listener
	defer listener.Close()

	// 创建用户数据通信的socket
	conn, err := listener.Accept()		// 阻塞等待...
	if err != nil {
		fmt.Println("Accept err:", err)
		return
	}
	defer conn.Close()

	fmt.Println("通信套接字，创建成功。。。")

	// 创建一个用保存数据的缓冲区
	buf := make([]byte, 4096)

	for {
		// 获取客户端发送的数据内容
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println("Read err:", err)
			return
		}
		// 处理 客户端 的数据
		fmt.Println("读到客户端发送：", string(buf[:n]))

		// 回写数据给客户端
		_, err = conn.Write([]byte("This is Server\n"))
		if err != nil {
			fmt.Println("Write err:", err)
			return
		}
	}
}