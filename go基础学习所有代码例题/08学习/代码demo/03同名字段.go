package main

import "fmt"

type person1 struct {
	name string
	sex  string
	age  int
}

type student1 struct {
	person1
	//子类和父类有相同名字的结构体成员
	name  string
	id    int
	score int
	addr  string
}

func main() {

	//var stu student1
	//
	//stu.id = 101
	//stu.sex = "男"
	//stu.age = 18
	/////采用就近原则为子类赋值
	//stu.name = "法师"
	////为父类结构体成员赋值
	//stu.person1.name="法师"
	//stu.addr = "黑龙江大庆"
	//stu.score=1000


	var stu student1=student1{person1{"法师","男",31},"大法师",101,100,"黑龙江大庆"}


	fmt.Println(stu.name)
	fmt.Println(stu.person1.name)
}
