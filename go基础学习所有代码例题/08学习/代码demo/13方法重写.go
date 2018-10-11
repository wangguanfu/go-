package main

import "fmt"

type person8 struct {
	id   int
	name string
	sex  string
	age  int
}

type student8 struct {
	person8

	score int
	addr  string
}

type teacher8 struct {
	person8

	subject  string
	worktime int
}

//接收者类型是指针类型或者符合类型 如果表示同一个结构体信息  就是一个相同的方法
func (p *person8) SayHello() {
	fmt.Printf("大家好，我是%s，我是%s生，我今年%d岁", (*p).name, p.sex, p.age)
}

//func (p person8)SayHello(){
//	fmt.Printf("大家好，我是%s，我是%s生，我今年%d岁")
//}
//学生子类方法
func (stu *student8) SayHello() {
	fmt.Printf("大家好，我是%s，我是%s生，我今年%d岁，我的成绩为：%d\n",
		stu.name, stu.sex, stu.age, stu.score)
}

func (tea *teacher8) SayHello() {
	fmt.Printf("大家好，我是%s，我是%s生，我今年%d岁，我教的学科是：%s\n",
		tea.name, tea.sex, tea.age, tea.subject)
}



func main() {

	//var p person8 = person8{101, "渣渣辉", "男", 90}
	//p.SayHello()

	//如果子类的方法和父类的方法相同 方法重写
	var stu student8 = student8{person8{101, "渣渣辉", "男", 40}, 100, "香港九龙"}
	//子类方法
	//stu.SayHello()
	//父类方法
	stu.person8.SayHello()

	var tea teacher8 = teacher8{person8{102, "古天乐", "男", 39}, "贪玩蓝月", 4}
	//tea.SayHello()

	tea.person8.SayHello()
}
