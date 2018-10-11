package main

import "fmt"

func main0201() {
	//变量 在程序运行过程中 其值给以发生改变的量 成为变量
	//变量的定义和使用
	//格式 var 变量名 数据类型
	//int 表示整型数据

	//var a int //声明 是没有具体的在 默认值为0
	//为变量赋值  =  称为赋值运算符
	//a = 10

	var a int = 10 //定义
	//变量参与程序计算
	a = a + 20
	fmt.Println(a)
}
func main0202() {
	//计算圆的面积和周长
	//面积 s=PI*r*r
	//周长 l=2*PI*r

	//PI 值为3.14
	//浮点型 可以存储带小数的数据 float64（小数位数为15位） float32（小数位数为7位）
	//默认小数位数保留6位为准确数据
	//浮点型数据都不是一个精准的数据
	var PI float64 = 3.14
	//半径
	var r float64 = 2.6
	//面积
	var s float64
	//周长
	var l float64
	//计算面积
	s = PI * r * r
	//计算周长
	l = PI * r * 2

	fmt.Println(s)
	fmt.Println(l)

}
