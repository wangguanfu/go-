package main

import (
	"fmt"
	"os"
	"strings"
	"io"
)

func copyMp3ToDir(fileName, path string)  {
	pathName := path + "/" + fileName
	fr, err := os.Open(pathName)				// 打开 源文件
	if err != nil {
		fmt.Println("Open err:", err)
		return
	}
	defer fr.Close()

	fw, err :=os.Create("./" + fileName)		// 打开拷贝的文件
	if err != nil {
		fmt.Println("Create err:", err)
		return
	}
	defer fw.Close()

	// 创建一个用于read 存储数据的 buf
	buf := make([]byte, 4096)

	// 循环从 fr 中读取数据， 原封不动 写到 fw中
	for {
		n, err := fr.Read(buf)
		if n == 0 {
			fmt.Println("文件拷贝完毕")
			break
		}
		if err != nil && err != io.EOF {
			fmt.Println("err:", err)
		}
		fw.Write(buf[:n])  	// 读多少，写多少
	}

}

func main()  {
	// 请用户指定 目录
	var path string
	fmt.Print("请输入目录位置：")
	fmt.Scan(&path)

	// 打开指定目录位置
	dir_fp, err := os.OpenFile(path, os.O_RDONLY, os.ModeDir)
	if err != nil {
		fmt.Println("Openfile err:", err)
		return
	}
	defer dir_fp.Close()

	// 读取目录项
	dirsInfo, err := dir_fp.Readdir(-1)
	if err != nil {
		fmt.Println("Readdir err:", err)
		return
	}
	// 从目录项[] 中提取每一个目录项
	for _, dir := range dirsInfo {
		fileName := dir.Name()
		// 根据后缀名，筛选文件
		if strings.HasSuffix(fileName, ".mp3") {

			// 将该文件 copy 至指定目录
			copyMp3ToDir(fileName, path)
			//fmt.Println("mp3文件有：", fileName)
		}
	}
}
