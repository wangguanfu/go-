package main

import (
	"os"
	"fmt"
)

/*func main()  {
	// 提取命令行参数
	list := os.Args

	for i:=0; i<len(list); i++ {
		fmt.Printf("list[%d]:%q\n", i, list[i])
	}
}
*/

func main()  {
	// 提取命令行参数
	list := os.Args

	// 判断命令行参数个数。
	if len(list) != 2 {
		fmt.Println("输入错误：格式为：go run xxx.go 文件绝对路径")
		return
	}
	filePath := list[1]	// 提取第二个参数 —— 文件名

	// 根据文件名获取文件属性
	fileInfo, err := os.Stat(filePath)
	if err != nil {
		fmt.Println("Stat err:", err)
		return
	}

	// 从 fileInfo 中获取 不带路径的“文件名”
	fmt.Println("name: ", fileInfo.Name())
	fmt.Println("size: ", fileInfo.Size())
}