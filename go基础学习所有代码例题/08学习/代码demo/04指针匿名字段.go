package main

import "fmt"

type person2 struct {
	name string
	age  int
	sex  string
}

type student2 struct {
	//指针匿名字段
	*person2

	score int
	id    int
	addr  string
}

func main() {

	//p:=person2{"段永超",18,"男"}
	//stu:=student2{&person2{"段永超",18,"男"},59,101,"北京五道口"}
	//
	//fmt.Println(stu)
	//fmt.Println(stu.person2.name)
	//fmt.Println(stu.age)

	var stu student2

	//stu.person2指针变量位空 nil
	//通过new创建空间 存储时数据
	stu.person2 = new(person2)
	//stu.person2.name = "边振龙"
	////stu.name = "边振龙"
	//stu.person2.age = 18
	//stu.person2.sex = "男"
	stu.name = "郑强"
	stu.sex = "男"
	stu.age = 20

	stu.id = 102
	stu.score = 100
	stu.addr = "河南道口"

	fmt.Println(stu)
	fmt.Println(stu.name)
	fmt.Println(stu.sex)
	fmt.Println(stu.age)
}
