package main

import (
	"os"
	"fmt"
	"io"
)

func main()  {
	// 创建 待读取的文件，和 待写入的文件
	f_r, err := os.Open("C:/itcast/01-结构体的定义和初始化.avi")
	if err != nil {
		fmt.Println("open err:", err)
		return
	}
	defer f_r.Close()

	f_w, err := os.Create("./test.avi")
	if err != nil {
		fmt.Println("Create err:", err)
		return
	}
	defer f_w.Close()

	buf := make([]byte, 4096)	//创建缓冲区，存储读到的数据

	// 循环从 f_r 对应文件中读取数据，
	for {
		n, err := f_r.Read(buf)
		if err == io.EOF {
			fmt.Println("读取文件完成")
			break
		}
		if err != nil {
			fmt.Println("Read err:", err)
			return
		}
		// 原封不动写到 f_w 文件中, 读多少，写多少
		n, err = f_w.Write(buf[:n])
		if err != nil {
			fmt.Println("Write err:", err)
			return
		}
	}

	os.OpenFile()
}
