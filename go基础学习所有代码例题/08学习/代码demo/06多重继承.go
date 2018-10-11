package main

import "fmt"

type human struct {
	name string
	sex  string
}

type person4 struct {
	human
	age int
}

type student4 struct {
	person4

	id    int
	score int
	addr  string
}

func main() {

	//stu := student4{person4{human{"张学霸", "男"}, 20}, 100, 99, "上海浦东新区"}

	var stu student4

	//stu.person4.human.name="张学霸"
	stu.name = "张学霸"
	stu.sex = "男"
	stu.age = 18
	stu.score = 99
	stu.id = 100
	stu.addr = "上海"
	fmt.Println(stu)
}
