package main

import "fmt"

//type 结构体名 struct{
//结构体成员列表
// 成员名 数据类型
// name string
// score int
//}

// () [] 结构体成员选择 .

//代码格式化对齐 ctrl+alt+l
//结构体定义在函数外  可以被项目中所有函数使用  结构体名称是唯一的
type Student struct {
	id    int
	name  string
	score int
	age   int
	sex   string
	addr  string
}

func main() {

	//结构体是一种数据类型
	//var stu Student=Student{101,"刘提纲",101,31,"女","上海浦东新区"}
	//自动推导类型创建结构体
	//stu:=Student{102,"结构体",1000,24,"男","山东龙口"}
	stu := Student{name: "马克", age: 18, sex: "女", score: 10, addr: "长达小区", id: 103}
	//结构体变量
	//结构体变量.成员名
	//stu.id =101
	//stu.score=100
	//stu.sex="男"
	//stu.name="刘法师"
	//stu.age=31
	//stu.addr="上海浦东新区"

	fmt.Println(stu)
	fmt.Printf("%T\n", stu)
}
