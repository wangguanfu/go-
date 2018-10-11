package main

import "fmt"

func main0201() {



	fmt.Println("hello world1")
	fmt.Println("hello world2")
	fmt.Println("hello world3")

	//程序可以运行 但是遇到panic停止
	//当程序遇到panic时  会自动崩溃
	panic("hello world")

	fmt.Println("hello world4")
	fmt.Println("hello world5")
	fmt.Println("hello world6")
}

func main(){
	//var p *int
	//
	////为空指针进行赋值  程序会自动调用panic异常
	//*p=123
	//
	//fmt.Println(*p)

	var arr [10]int
	var i int =10

	//数组下标越界
	arr[i]=123

	fmt.Println(arr)
}