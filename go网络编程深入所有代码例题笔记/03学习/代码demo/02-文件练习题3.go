package main

import (
	"os"
	"fmt"
	"strings"
	"bufio"
	"io"
)

func countLove(fileName, path string) int {
	// 打开文件，txt
	f, err := os.Open(path + "/" + fileName)
	if err != nil {
		fmt.Println("Open err:", err)
		return -1
	}
	defer f.Close()

	// 1 创建 reader
	reader := bufio.NewReader(f)
	var counter int = 0

	// for 循环 按行读取文件，--- this is a  test for love  --》[]string
	for {
		// 2 指定分隔符，按行提取数据
		buf, err := reader.ReadBytes('\n')			// buf中保存读到的一行数据
		if err != nil && err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("ReadBytes err:", err)
			return -1
		}
		// 将一行数据中的所有单词 拆分
		words := strings.Fields(string(buf))
		for _, word := range words {
			if word == "love" {
				counter++	// 统计 个数。
			}
		}
	}
	fmt.Println("counter:", counter)

	return counter
}

func main()  {

	// 测试统计 函数是否生效
	//countLove("test.txt", "C:/itcast/test2")

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

	// 统计目录下所有 txt 文件的 love 个数
	var allLove int = 0

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
		if strings.HasSuffix(fileName, ".txt") {
			// 打开文件， 统计 该文件中有多少个 “love”
			allLove += countLove(fileName, path)
		}
	}
	fmt.Println("指定的目录中，一共包含 love 个数：", allLove)
}
