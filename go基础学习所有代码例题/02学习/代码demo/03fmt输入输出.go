package main

import "fmt"

func main0301() {
	//打印 %  需要使用%%
	//fmt.Printf("35%%")
	//fmt.Println("35%")
	//十进制数据
	//a:=123
	//八进制数据是以0开头
	//var a int
	//a=0123
	//十六进制数据 是以0x开头的
	var a int
	a=0xABC
	//二进制  逢二进一 所有数据都是1 0 组成的
	//%b  占位符 表示输出二进制
	fmt.Printf("%b\n",a)

	//八进制 逢八进一 所有数据用0-7组成
	//%o  占位符 表示输出八进制
	fmt.Printf("%o\n",a)

	//十六进制 逢十六进一 所有数据都是 0-9 A-F 10A 11B 12C 13D 14E 15F
	//%x  占位符 表示输出十六进制  小写字母
	fmt.Printf("%x\n",a)
	//%X  占位符 表示输出十六进制  大写字母
	fmt.Printf("%X\n",a)
	//十进制
	fmt.Printf("%d\n",a)

	//%p占位符 表示数据的内存地址
	fmt.Printf("a的内存地址：%p\n",&a)
}

func main(){
	//所有数据的内存地址都是一个无符号十六进制整型数据
	a:="hello"
	fmt.Printf("%p\n",&a)
}
