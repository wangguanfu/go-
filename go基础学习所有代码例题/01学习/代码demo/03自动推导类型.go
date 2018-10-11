package main

import "fmt"

func main0301() {
	//自动推导类型
	//在一个函数范围内变量名不允许重名
	//var a int = 10
	a:=10//根据数据的值  来确定变量的类型
	b:=3.14
	fmt.Println(a)
	fmt.Println(b)
	//%T 是一个占位符表示输出变量对应的数据类型
	fmt.Printf("%T\n",a)//int
	fmt.Printf("%T\n",b)//float64

}
func main0302(){
	a:=10
	b:=20

	//var c int
	//交换变量a b 的值
	//c:=a//自动推导类型创建变量
	//c=a
	//a=b
	//b=c

	a=a+b
	b=a-b
	a=a-b

	fmt.Println("a=",a)
	fmt.Println("b=",b)
}