package main

import (
	"net"
	"fmt"
	"time"
)

func main()  {
	// 创建用于通信socket
	conn, err := net.Dial("tcp", "127.0.0.1:7020")	// IP地址/prot端口号--服务器的
	if err != nil {
		fmt.Println("Dial err:", err)
		return
	}
	// 关闭连接
	defer conn.Close()

	for {
		// 发送数据 write
		_, err = conn.Write([]byte("hello socket"))
		if err != nil {
			fmt.Println("Write err:", err)
			return
		}
		buf := make([]byte, 4096)
		// 接收数据 read
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println("Read err:", err)
			return
		}
		fmt.Println("读到 服务器 回发：", string(buf[:n]))

		time.Sleep(time.Second)
	}
}
