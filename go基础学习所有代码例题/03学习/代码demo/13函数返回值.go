package main

import "fmt"

//func  plus(a int, b int) int{
//
//	sum:=a+b
//
//	//return 如果在函数中出现return 表示函数的结束
//	return sum
//
//}
//在函数定义过程中 可以将函数的返回值 直接定义
func plus(a int, b int) (sum int) {

	sum = a + b
	return
}
func main1301() {

	a := 10
	b := 20

	var value int
	value = plus(a, b)

	fmt.Println(value)

}

//go 语言 允许程序有多个范围值
func test3() (a int, b int, c int) {
	a, b, c = 10, 20, 30
	return
}
func main1302() {
	// _ 匿名变量 用来在接收函数返回值不需要的值
	// 匿名变量可以使用多次
	a,_,_:= test3()

	fmt.Println(a)

}
