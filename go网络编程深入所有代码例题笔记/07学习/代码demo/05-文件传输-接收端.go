package main

import (
	"net"
	"fmt"
	"os"
)

func main()  {
	// 创建 监听 socket
	listener, err := net.Listen("tcp", "127.0.0.1:8010")
	if err != nil {
		fmt.Println("Listen err:", err)
		return
	}
	defer listener.Close()

	// 创建 通信socket
	conn, err := listener.Accept()
	if err != nil {
		fmt.Println("Accept err:", err)
		return
	}
	defer conn.Close()

	buf := make([]byte, 4096)
	// 读取 发送端发送的 “文件名”
	n, err := conn.Read(buf)
	if err != nil {
		fmt.Println("Read err:", err)
		return
	}
	fileName := string(buf[:n])		// 无路径的文件名

	conn.Write([]byte("ok"))	// 保存文件名后，应答“ok”

	// 封装函数， 获取文件内容
	RecvFile(fileName, conn)
}

func RecvFile(fileName string, conn net.Conn)  {
	// 根据保存的文件名，创建文件
	f, err := os.Create(fileName)			// ./fileName
	if err != nil {
		fmt.Println("Create err:", err)
		return
	}
	defer f.Close()

	buf := make([]byte, 4096)
	for {							// 循环从 socket 中读取文件内容
		n, err := conn.Read(buf)
		if n == 0 {					// 判断是否读到文件末尾
			fmt.Println("读取文件结束")		// 对端 Close
			break
		}
		if err != nil {
			fmt.Println("Read err:", err)
			return
		}
		f.Write(buf[:n])			// 读多少，写多少
	}
}