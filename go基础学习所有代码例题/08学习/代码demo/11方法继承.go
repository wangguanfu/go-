package main

import "fmt"

type person6 struct {
	id   int
	name string
	sex  string
	age  int
}

type student6 struct {
	person6

	score int
	addr  string
}

//接收者的类型如果不相同 就是不同的方法
//函数名 可以和方法重名
//func Print() {
//	fmt.Print("hello")
//}

func (p *person6) Print() {
	fmt.Printf("大家好，我是%s,我今年%d岁\n", p.name, p.age)
}

func main() {

	var stu student6=student6{person6{101,"段永超","男",10},100,"北京五道口"}

	//子类可以从父类继承结构体成员  也可以继承父类的方法
	//stu.Print()
	stu.person6.Print()


	var per person6=person6{102,"挣钱","男",10}

	per.Print()

}
