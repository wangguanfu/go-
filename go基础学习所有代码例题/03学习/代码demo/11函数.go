package main

import "fmt"

//     *
//    ***
//   *****
//  *******
// *********
//***********

//函数定义和调用
//func  函数名(函数参数列表) 返回值类型 {代码体}

//计算10+20和
//函数定义(函数实现过程)
func add(a int, b int) {
	sum := a - b
	fmt.Println(sum)
}

func main1101() {
	//返回值（int）=len(字符串变量(string()))

	a := 10
	b := 20
	//函数调用
	add(a, b)
	//函数优点  可以重复使用  修改迭代简单
	add(1, 2)
	add(55, 44)

}

//在函数定义时 函数形参  形式参数
//形参接收实参传递的值  形参是格式要求  根据形参的类型和个数来去调用函数
func test(a int, b int) {
	a = 100
	b = 200
	fmt.Println(a, b)
}
func main0111() {
	a := 10
	b := 20
	//在函数调用中 函数的参数称为实参 实际参数
	test(a, b)
}