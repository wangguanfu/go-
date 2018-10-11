package main

import "fmt"

func main() {
	//字符类型定义 ''单引号引起来字符
	//var a byte
	////a='a'
	////转义字符 也是字符的一种用\开头  \0
	////a='\n'
	////a='男'//err 一个汉字占3个字节
	//fmt.Printf("==%c==",a)
	//fmt.Println(a)
	////当定义\0是 默认将\0当做八进制数据
	//a=0//相当于\0

	//字符串类型
	//在字符串结尾有\0表示结尾
	//var b string="表示结尾"
	var b string = "hello world"
	//len(string)  计算字符串中字符个数  不包含\0  返回值为int类型 返回值为个数
	var count int
	count = len(b)
	//fmt.Println(b)
	//fmt.Println(len(b))
	fmt.Println(count)

	//\t转义字符  水平制表符 一次跳过8个空格  打印数据对齐会使用到
	//fmt.Printf("\t你好")
}
