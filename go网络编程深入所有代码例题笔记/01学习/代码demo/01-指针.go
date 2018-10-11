package main

import "fmt"

/*func main()  {
	var a int = 10
	var p *int = &a

	fmt.Println("a = ", a)
	fmt.Println("*p = ", *p)
}*/

func test(d int) {
	var m int = 400
	m += d
}
func main() {
	//	var a int = 10
	//	fmt.Println("&a=", &a)
	var p *int = nil
	p = new(int)	// 给指针做有效初始化

	test(90)

	*p = 100
	fmt.Println("*p =", *p)
}




