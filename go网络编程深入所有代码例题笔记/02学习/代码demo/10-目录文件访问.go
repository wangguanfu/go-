package main

import (
	"os"
	"fmt"
)

func main()  {
	var dir string
	fmt.Print("请输入一个待判定的目录：")
	fmt.Scan(&dir)

	// 打开目录
	f, err := os.OpenFile(dir, os.O_RDONLY, os.ModeDir)
	if err != nil {
		fmt.Println("OpenFile err:", err)
		return
	}
	defer f.Close()

	// 读取目录项 -- 将一个目录下的所有内容， 存入[]FileInfo
	fileInfo, err := f.Readdir(-1)
	if err != nil {
		fmt.Println("Readdir err:", err)
		return
	}

	// 依次取出目录项
	for _, dirInfo := range fileInfo {
		if dirInfo.IsDir() {
			fmt.Printf("%s 是一个 目录\n", dirInfo.Name())
		} else {
			fmt.Printf("%s is not 目录\n", dirInfo.Name())
		}
	}

}
