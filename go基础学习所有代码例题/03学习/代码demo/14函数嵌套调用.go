package main

import "fmt"

func test5(a int, b int) {
	fmt.Println(a + b)
}

func test4(a int, b int) {
	test5(a, b)
}

func main1401() {
	test4(10, 20)
}

func test7(a ...int){
	for i:=0;i<len(a);i++{
		fmt.Println(a[i])
	}
}

//不定参函数调用
func test6(a ...int){
	//不定参调用格式
	//test7(a...)
	//test7(a[0],a[1],a[2])
	//test7(a[0:3]...)//1 2 3
	//a[起始位置下标：结束位置下标]不包含结束位置下标
	test7(a[2:5]...)//3 4 5
}
func main(){
	test6(1,2,3,4,5)
}