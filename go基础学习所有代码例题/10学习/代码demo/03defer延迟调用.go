package main

import "fmt"

func main0301() {

	//defer是在函数结束时  从栈区最后进入的数据依次的向前进行函数调用

	fmt.Println("hello world1")
	defer fmt.Println("hello world2")
	fmt.Println("hello world3")
	defer fmt.Println("hello world4")
	fmt.Println("hello world5")
}

func main0302() {

	a := 10
	b := 20

	defer func(a int, b int) {
		fmt.Println("打印2",a + b)
	}(a, b)

	a = 100
	b = 200
	fmt.Println("打印1",a + b)
}



func main(){
	a := 10
	b := 20

	//现将函数加载到内存中 在主函数结束时在调用函数
	defer func() {
		fmt.Println("打印2",a + b)
	}()

	a = 100
	b = 200
	fmt.Println("打印1",a + b)
}
