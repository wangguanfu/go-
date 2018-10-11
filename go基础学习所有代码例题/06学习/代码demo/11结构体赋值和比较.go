package main

import "fmt"

type student struct {
	id    int
	name  string
	age   int
	sex   string
	score int
	addr  string
}

func main() {


	stu:=student{101,"秦海东",18,"女",100,"江苏启东"}

	//结构体赋值
	s1:=stu


	s1.sex="男"
	//结构体比较时 比较所有成员是否相同
	//结构体变量不允许 大小比较  结构体成员可以进行大小比较
	if s1==stu{
		fmt.Println("相同")
	}else{
		fmt.Println("不相同")
	}


	//fmt.Println(s1)
	//fmt.Println(stu)
	//
	//fmt.Printf("%p\n",&stu)
	//fmt.Printf("%p\n",&stu.id)
	//fmt.Printf("%p\n",&stu.name)
	//fmt.Printf("%p\n",&stu.age)
	//fmt.Printf("%p\n",&stu.sex)
	//fmt.Printf("%p\n",&s1)


}
