package main

import "fmt"

type stu struct {
	id   int
	name string
	age  int
}

func (s *stu) Print(a int, b int) int {
	fmt.Println("我是", s.name, "我学号", s.id)
	return 0
}

func (s *stu) test() {

}

func test(a int, b int) int {
	return 0
}

type FUNC func(int, int) int

func main0101() {

	s := stu{101, "貂蝉", 16}

	//s.Print()

	//代码区地址 方法也会在代码区进行存储
	//fmt.Println(s.Print)

	//fmt.Printf("%T\n",s.Print)

	//方法的类型和函数的类型一致的
	var f FUNC

	f = s.Print

	f(1, 2)
	//f = s.test/err
	//f()

	f = test
	f(1, 2)

	fmt.Println(test)
	fmt.Println(s.test)

}
func main() {
	var s stu

	var f FUNC
	f = s.Print
	f(1,1)

	fmt.Printf("%T\n",f)
}
