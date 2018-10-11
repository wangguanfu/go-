package main

import "fmt"

type Student struct {
	id    int
	name  string
	age   int
	sex   string
	score int
	addr  string
}

func main0801() {

	stu := Student{101, "段永超", 18, "男", 59, "河南周口"}

	var p *Student = &stu
	//p:=&stu

	//fmt.Printf("%T\n",p)

	//通过指针间接修改结构体成员的值
	//(*p).score = 60

	p.score = 60
	fmt.Println(*p)
}

func test8(stu *Student) {

	stu.score=60
	stu.name="郑强"
}

func main() {
	stu := Student{101, "段永超", 18, "男", 59, "河南周口"}

	//结构体指针作为函数参数  地址传递  引用传递
	test8(&stu)

	fmt.Println(stu)
}
