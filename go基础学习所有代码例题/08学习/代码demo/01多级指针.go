package main

import "fmt"

func main() {

	a := 10 //int
	b := 20

	p := &a //*int

	//pp := &p//**int
	//定义二级指针存储一级指针的地址
	var pp **int = &p


	var p3 ***int = &pp

	**pp = 123
	*pp = &b
	//二级指针的值
	//*p3
	//一级指针的值
	//**p3
	//变量的值
	***p3=123

	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", p)
	fmt.Printf("%T\n", pp)

	fmt.Println(*p)
	fmt.Println(a)
	fmt.Println(b)
}
