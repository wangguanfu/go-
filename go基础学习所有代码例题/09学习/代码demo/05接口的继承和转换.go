package main

import "fmt"

//子集
type Human interface {
	SayHello()
}

//超集
type Personer interface {
	//继承字节接口
	Human
	sing(name string)
}

type person1 struct {
	name string
	age  int
	sex  string
}

type student1 struct {
	person1
	score int
}

type teacher1 struct {
	person1

	subject string
}

func (stu *student1) SayHello() {
	fmt.Printf("大家好，我是学生%s,我今年%d岁，我是%s生\n",
		stu.name, stu.age, stu.sex)
}

func (tea *teacher1) SayHello() {
	fmt.Printf("大家好，我是老师%s,我今年%d岁，我是%s生\n",
		tea.name, tea.age, tea.sex)
}

func (tea *teacher1) sing(name string) {
	fmt.Printf("我为大家唱首歌：%s\n", name)
}

func main() {

	var i Human    //子集
	var p Personer //超集

	var stu = student1{person1{"姜维", 40, "男"}, 80}

	i = &stu
	//字节中缺少sing()方法
	//p=i//err
	i.SayHello()

	var tea = teacher1{person1{"诸葛亮", 58, "男"}, "五行八卦"}

	//i = &tea
	//i.SayHello()
	//p = &tea
	//
	//p.SayHello()
	//p.sing("夜上海")
	p = &tea
	//子集中的方法 在超集中都存在
	i = p

	i.SayHello()
	//i.sing()//err
}
