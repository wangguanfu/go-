package main

import "fmt"

//接口定义和使用

type Humaner interface {
	//方法的声明
	SayHi()
	//接口中定义的方法 必须有实现
	//PrintInfo()

}

type person struct {
	name string
	sex  string
	age  int
}

type student struct {
	person

	score int
}

type teacher struct {
	person

	subject string
}

//学生子类方法
func (stu *student) SayHi() {
	fmt.Printf("大叫好，我叫%s，我今年%d岁，我是%s生，我的成绩%d\n",
		stu.name, stu.age, stu.sex, stu.score)
}
func (stu *student)PrintInfo(){

}

//老师子类方法
func (tea *teacher) SayHi() {
	fmt.Printf("大叫好，我叫%s，我今年%d岁，我是%s生，我的学科%s\n",
		tea.name, tea.age, tea.sex, tea.subject)
}

func main() {

	//接口类型  数据类型  //先定义接口 在根据接口来创建方法
	//接口是虚的 方法的实的  接口定义规则  方法实现规则
	var i Humaner

	var stu student = student{person{"杨过", "男", 28}, 99}

	i = &stu

	i.SayHi()

	var tea teacher = teacher{person{"尹志平", "男", 30}, "驯龙"}

	i = &tea

	i.SayHi()

}
