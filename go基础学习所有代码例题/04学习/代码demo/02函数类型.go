package main

import "fmt"

func add(a int, b int) {
	fmt.Println(a + b)
}

func sub(i int, j int){
	fmt.Println(i-j)
}
//创建函数类型
//type  1、创建函数类型 2、为已存在的数据类型起别名
type FUNCTYPE func (int,int) 
func main() {
	a := 10
	b := 20
	//add(a, b)
	//函数名是一个地址 指向代码区的地址
	//fmt.Println(add)
	//var f func (int, int)
	//var 变量名 数据类型（函数类型）
	//f是一个函数类型的变量 可以被赋值
	var f FUNCTYPE
	f = add
	f(a, b)
	//add=f//err add是一个常量(只读地址) 表示函数的地址 函数指针
	f=sub
	f(a,b)

	fmt.Printf("%T\n", f)
	fmt.Printf("%T\n", add)
	fmt.Println(f)
}
