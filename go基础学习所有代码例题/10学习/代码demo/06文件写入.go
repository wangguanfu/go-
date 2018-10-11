package main

import (
	"os"
	"fmt"
	"io"
)

func main0601() {

	fp, err := os.Create("D:/a.txt")
	if err != nil {
		fmt.Println("文件创建失败")
		return
	}

	//关闭文件
	defer fp.Close()

	//将一个字符串写入到文件中
	//如果没有出现明确的换行 writestring 会将所有的字符串写在一行
	fp.WriteString("hello world")
	//使用换行 需要写 \r\n windows系统中换行是以\r\n为换行 在linux中是以\n为换行
	fp.WriteString("好好学习\r\n")
	fp.WriteString("高薪就业")

}
func main0602() {
	fp, err := os.Create("D:/a.txt")

	if err != nil {
		fmt.Println("文件创建失败")
		return
	}

	defer fp.Close()

	//s:=[]byte{'h','e','l','l','o'}
	//将字符串转成字符切片
	s := []byte("hello")
	fp.Write(s)

	s = []byte("好好学习高薪就业月薪过万")
	fp.Write(s)
	fp.Write([]byte("澳门在线赌场上线了"))
}
func main() {

	//只读方式打开文件
	//fp, err := os.Open("D:/a.txt")
	//读写方式打开文件  1、文件路径 2、打开模式 3、打开权限 r（4）w（2）x（1）
	fp, err := os.OpenFile("D:/a.txt", os.O_RDWR, 6)
	if err != nil {
		fmt.Println("打开文件失败\n")
		return
	}
	defer fp.Close()

	//fp.WriteAt([]byte("world"),0)
	//替换指定位置中的文字
	//fp.WriteAt([]byte("日薪越亿"), 17)
	//io.SeekStart 偏移到文件起始位置
	//io.SeekCurrent 偏移到文件末尾
	//io.SeekEnd 偏移到文件末尾
	//将文件偏移量作为写入值
	i, _ := fp.Seek(0, io.SeekEnd)

	fp.WriteAt([]byte("性感荷官在线发牌"),i)

}
