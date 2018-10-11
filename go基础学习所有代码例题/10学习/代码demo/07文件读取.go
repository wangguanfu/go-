package main

import (
	"os"
	"fmt"
	"io"
	"bufio"
)

func main0701() {

	//只读方式打开
	fp, err := os.Open("D:/a.txt")

	//读写追加方式打开
	//os.OpenFile()

	if err != nil {
		fmt.Println("文件打开失败")
		return
	}

	defer fp.Close()

	//读取指定文件内容的一块  块读取
	//s := make([]byte, 5)
	////读取文件
	//fp.Read(s)
	//
	//fmt.Println(string(s))
	//
	//str:=make([]byte,12)
	//fp.Read(str)
	//fmt.Println(string(str))

	//s := make([]byte, 1024)
	//n, _ := fp.Read(s)
	//
	//fmt.Println(string(s[:n]))

	//循环读到文件末尾
	s := make([]byte, 4)
	for {
		//文本文件是以ASCII进行存储的 0-127 128-255  32-126  ‘0’48 'a'97 'A'65
		//文件结束标志io.EOF (-1)

		n, err := fp.Read(s)
		if err == io.EOF {
			//到文件结尾
			break
		}
		fmt.Print(string(s[:n]))
	}

}
func main0702() {
	fp, err := os.Open("D:/a.txt")

	if err != nil {
		fmt.Println("文件打开失败")
		return
	}
	defer fp.Close()

	//创建读取缓冲区
	r := bufio.NewReader(fp)

	//创建切片

	s := []byte{}
	//行读取  包含\n
	s, _ = r.ReadBytes('\n')

	fmt.Print(string(s))
}

func main() {
	fp, err := os.Open("D:/a.txt")
	if err != nil {
		fmt.Println("打开文件失败")
		return
	}
	defer fp.Close()

	//	创建缓冲区

	r := bufio.NewReader(fp)

	for {
		//r.ReadBytes 可以进行截取字符
		s, err := r.ReadBytes(' ')
		fmt.Println(string(s))
		//判断文件末尾
		if err == io.EOF {
			break
		}

	}
}
