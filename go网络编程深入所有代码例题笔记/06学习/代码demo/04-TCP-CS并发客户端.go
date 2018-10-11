package main

import (
	"net"
	"fmt"
	"os"
)

func main()  {
	conn, err := net.Dial("tcp", "127.0.0.1:8000")
	if err != nil {
		fmt.Println("Dial err:", err)
		return
	}
	defer conn.Close()

	// 读取用户的键盘输入。
	go func() {
		buf := make([]byte, 4096)
		for {
			// 获取键盘输入。 fmt.Scan --》 结束标记 \n 和 空格
			n, err := os.Stdin.Read(buf)			// buf[:n]
			if err != nil {
				fmt.Println("os.Stdin.Read err:", err)
				return
			}
			// 直接将读到键盘输入数据，写到 socket 中，发送给服务器
			conn.Write(buf[:n])
		}
	}()

	// 在 主go程中， 获取服务器回发数据。
	buf2 := make([]byte, 4096)
	for {
		// 借助 socket 从服务器读取 数据。
		n, err := conn.Read(buf2)
		if n == 0 {
			fmt.Println("客户端检查到服务器，关闭连接， 本端也退出")
			return
		}
		if err != nil {
			fmt.Println("os.Stdin.Read err:", err)
			return
		}
		fmt.Println("客户端读到：", string(buf2[:n]))
	}
}
