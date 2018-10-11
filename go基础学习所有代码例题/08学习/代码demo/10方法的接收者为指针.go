package main

import "fmt"

type students struct {
	id    int
	name  string
	sex   string
	age   int
	score int
}

func (s *students)Print(){
	s.name="法师"
}
func main() {

	stu:=students{101,"刘某某","男",31,100}

	//对象绑定的方法可以是指针类型  使用时可以结构体变量（对象）.方法
	stu.Print()

	fmt.Println(stu)

}
