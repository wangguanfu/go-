package main

import "fmt"

func test3(a int, b int) {
	a, b = b, a
}
func test4(a *int, b *int) {

	//temp := *a
	//*a = *b
	//*b = temp
	*a, *b = *b, *a
}
func main() {
	a := 10
	b := 20
	//变量作为函数参数是值传递
	//test3(a, b)
	//指针作为函数参数是地址传递
	test4(&a, &b)
	fmt.Println(a, b)
}
