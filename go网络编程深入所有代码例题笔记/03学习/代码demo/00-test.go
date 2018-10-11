package main

import (
	"os"
	"fmt"
	"bufio"
	"io"
	"strings"
)

func main()  {

	f, err := os.Open("C:/itcast/test2/test.txt")
	if err != nil {
		fmt.Println("Open err:", err)
		return
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
			return
		}
		// 将一行数据中的所有单词 拆分
		//words := strings.Fields(string(buf))
		words := strings.Split(string(buf), " ")
		for _, word := range words {
			fmt.Printf("word:%q\n", word)
			if word == "love" {
				counter++	// 统计 个数。
			}
		}
		fmt.Println("===============================")
	}
	fmt.Println("count:", counter)
}
