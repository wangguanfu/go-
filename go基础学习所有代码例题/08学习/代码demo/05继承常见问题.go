package main

import "fmt"

type person3 struct {
	student3

	name string
	id   int
	addr string
}
type student3 struct {
	//结构体继承于结构体本身  是错误的  无效的递归循环  无法确定结构体在内存中存储的大小
	//student3//err

	//结构体继承于结构体指针可以的  预先知道指针大小
	//结构体继承与结构体指针 是一种数据结构 链表表现方式
	//*student3


	//两个结构体不能互相继承  错误的
	person3

	name string
	sex  string
	age  int
}

func main() {

	var stu student3

	stu.name = "梁彦青"
	stu.sex = "男"
	stu.age = 22

	fmt.Println(stu)
}
