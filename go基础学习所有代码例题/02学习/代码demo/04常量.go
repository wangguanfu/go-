package main

import "fmt"

func main0401() {
	//变量定义格式
	//var 变量名  数据类型
	//常量定义和使用
	//const 常量名 数据类型
	//常量是在程序运行过程中 其值不能发生改变的量 成为常量
	//常量一般用大写字母表示 和变量进行区分
	const PI float64 = 3.14
	//常量定以后值不能发生改变
	//PI=3.14159//err
	var b float64
	//代码写完后可以进行格式化 ctrl+alt+l
	b = PI
	fmt.Println(PI * 2.5 * 2.5)
	//常量可以参与程序计算或者赋值给变量
	fmt.Println(b)

	fmt.Printf("%p\n", &b)
	//常量的地址不允许访问 常量是在数据区下的常量区进行存储
	//fmt.Printf("%p\n",&PI)//err
}

func main(){
	//字面常量
	fmt.Println(3.14)
	fmt.Println("hello world")
	a:="hello"
	a=a+"world"
	fmt.Println(a)
}