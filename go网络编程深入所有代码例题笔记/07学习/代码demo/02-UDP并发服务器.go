package main

import (
	"net"
	"fmt"
)

func main()  {

	srvAddr, err := net.ResolveUDPAddr("udp", "127.0.0.1:8000")
	if err != nil {
		fmt.Println("ResolveUDPAddr err:", err)
		return
	}
	conn, err := net.ListenUDP("udp", srvAddr)
	if err != nil {
		fmt.Println("ListenUDP err:", err)
		return
	}
	defer conn.Close()

	for {
		buf := make([]byte, 4096)
		n, cltAddr, err := conn.ReadFromUDP(buf)
		if err != nil {
			fmt.Println("ReadFromUDP err:", err)
			return
		}
		go func() {
			fmt.Printf("客户端%v,发送数据：%s\n", cltAddr, string(buf[:n]))
			conn.WriteToUDP(buf[:n], cltAddr)
		}()
	}
}
