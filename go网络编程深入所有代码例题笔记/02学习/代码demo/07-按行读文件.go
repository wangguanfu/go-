package main

import (
	"os"
	"fmt"
	"bufio"
	"io"
)

func main()  {
	f, err := os.Open("test.txt")
	if err != nil {
		fmt.Println("Open err:", err)
		return
	}
	defer f.Close()

	// 获取阅读器 reader， 自带缓冲区（用户缓冲）。
	reader := bufio.NewReader(f)

	for {			// 循环读取文件， 当 err == io.EOF 结束循环
		// 使用 带分割符的 函数，读取指定数据 ‘\n’获取一行
		buf, err := reader.ReadBytes('\n')
		// 成功读取到的一行数据 保存在 buf中
		fmt.Printf("buf:%s", buf)
		if err != nil && err == io.EOF {
			break
		}
	}
	fmt.Println("文件读取完毕")
}
