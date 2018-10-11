package main

import "fmt"

func main0701() {

	slice := []int{1, 2, 3, 4, 5}

	//指向切片的指针
	var p *[]int = &slice
	//切片不允许直接通过指针来操作元素 err
	//	p[0]=123
	//	p[2]=222

	(*p)[0] = 123
	(*p)[2] = 222
	fmt.Println(slice)

	fmt.Printf("%p\n", slice)
}

//func test6(s []int) []int {
//	s = append(s, 1, 2, 3)
//	return s
//}

func test7(s *[]int){
	*s=append(*s,1,2,3,4)
}
func main() {
	var s []int
	//s = test6(s)
	//切片指针作为函数参数 地址传递
	test7(&s)
	fmt.Printf("%p\n", s)
	fmt.Println(s)
}
