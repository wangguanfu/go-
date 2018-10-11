package main

import "fmt"

//func test4() {
//	var i int
//	i++
//	fmt.Println(i)
//}

//type FUNCTEST func() int
func test4() func() int{
	var i int
	return func() int{
		i++
		return i
	}
}
func main() {

	//for i := 0; i < 10; i++ {
	//	//	//函数结束会从栈区内存销毁
	//	//	test4()
	//	//}

	//func() func() int
	//v:=test4

	//func() int
	//v是一个函数类型的变量
	v:=test4()
	//函数调用
	//fmt.Println(v())
	for i := 0; i < 10; i++ {
		fmt.Println(v())
	}
	//fmt.Printf("%T\n",v)//int   func() int  func() func()int
}
