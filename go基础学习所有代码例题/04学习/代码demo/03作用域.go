package main

import "fmt"

//在函数外部定义的变量成为全局变量
//全局变量和局部变量可以重名
//全局变量可以在项目中所有文件进行使用
//在任意位置修改全局变量的值都会影响其他位置使用
//全局变量存储在数据区

var a int =100

//全局变量定义时不能使用自动推到类型
//c:=100//err

//全局常量
const b int = 10

func test(){
	a=123
	fmt.Println(a)
}
func test1(){
	fmt.Println(a)
}
func main() {
	fmt.Printf("全局变量地址：%p\n",&a)
	//fmt.Printf("全局常量地址：%p\n",&b)//err

	//在函数内部定义的变量成为局部变量
	var a  int =10
	//在局部变量作用域范围内 全局变量不起作用  就近原则
	a=234
	fmt.Println(a)
	test()
	test1()
	//不能获取常量地址
	//const b int =10
	//fmt.Printf("常量地址：%p\n",&b)

	fmt.Printf("局部变量地址：%p\n",&a)

}
