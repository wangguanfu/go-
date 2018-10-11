package main

import (
	"net"
	"fmt"
	"os"
	"io"
)

func main()  {
	// 提示用户输入带传送文件名， 获取命令行参数。
	list := os.Args
	if len(list) != 2 {
		fmt.Println("格式错误！参照：go run xxx.go 文件绝对路径")
		return
	}
	filePath := list[1]

	// 提取文件属性, fileInfo.Name() 提取不带路径文件名
	fileInfo, err := os.Stat(filePath)
	if err != nil {
		fmt.Println("Stat err:", err)
		return
	}
	fileName := fileInfo.Name()		// 不带路径文件名

	// 与接收端建立连接
	conn, err := net.Dial("tcp", "127.0.0.1:8010")
	if err != nil {
		fmt.Println("Dial err:", err)
		return
	}
	defer conn.Close()

	// 发送文件名 给接收端
	_, err = conn.Write([]byte(fileName))
	if err != nil {
		fmt.Println("Write err:", err)
		return
	}

	buf := make([]byte, 4096)
	// 读取接收端回发“ok”
	n, err := conn.Read(buf)
	if err != nil {
		fmt.Println("Read err:", err)
		return
	}
	// 判断是否是 ok
	if string(buf[:n]) == "ok" {
		// 封装函数，专门发送数据给 接收端
		SendFile(filePath, conn)
	}
}

func SendFile(filePath string, conn net.Conn)  {
	// 只读打开文件
	f, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Open err:", err)
		return
	}
	defer f.Close()

	buf := make([]byte, 4096)

	for {
		// 读文件
		n, err := f.Read(buf)
		if err != nil {
			if err == io.EOF {
				fmt.Println("读取到文件结尾")
				break
			} else {
				fmt.Println("Read err:", err)
				return
			}
		}
		// 借助socket 写回给 接收端,
		_, err = conn.Write(buf[:n])
		if err != nil {
			fmt.Println("Write err:", err)
			return
		}
	}
}
