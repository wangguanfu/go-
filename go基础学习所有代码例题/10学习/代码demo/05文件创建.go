package main

import (
	"os"
	"fmt"
)

func main() {

	//创建新文件  os.Create 如果文件不存在会创建新文件  如果文件存在会覆盖之前内容 变成空文件
	fp, err := os.Create("D:/a.txt")

	if err != nil {
		//文件创建失败 1、硬盘空间不足 2、没有写权限 3、没有指定的路径 4、应用程序打开文件上限  65535
		fmt.Println("创建失败")
		return
	}

	//可以通过文件指针操作文件


	//关闭文件
	fp.Close()


}
