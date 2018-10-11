package main

import (
	"net"
	"fmt"
)

func main()  {
	// 创建 UDPAddr 类型的 UDP 地址结构（IP + port）, 为了给ListenUDP传参
	udpAddr, err := net.ResolveUDPAddr("udp", "127.0.0.1:8009")
	if err != nil {
		fmt.Println("ResolveUDPAddr err:", err)
		return
	}
	// 创建用于 数据通信的 socket -- conn
	conn, err := net.ListenUDP("udp", udpAddr)
	if err != nil {
		fmt.Println("ListenUDP err:", err)
		return
	}
	defer conn.Close()

	buf := make([]byte, 4096)

	// 读取客户端数据
	n, cUDPAddr, err := conn.ReadFromUDP(buf)
	if err != nil {
		fmt.Println("ReadFromUDP err:", err)
		return
	}
	fmt.Printf("读到客户端 %v 发送数据:%s\n", cUDPAddr, string(buf[:n]))

	// 写数据给客户端
	conn.WriteToUDP([]byte("This is UDP server"), cUDPAddr)
}
