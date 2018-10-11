package main

import "fmt"

type stu struct {
	id    int
	name  string
	sex   string
	age   int
	score int
}

//为结构体 数据类型绑定方法
//结构体名可以直接作为方法的接收者
func (s stu) PrintInfo() {

	s.name="大法师"
	fmt.Println("大家好我叫", s.name, "我今年", s.age)


}

func main() {
	s1 := stu{101, "张强伟", "男", 20, 0}


	//对象.方法
	s1.PrintInfo()

	fmt.Println(s1)

}
