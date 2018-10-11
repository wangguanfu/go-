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

//结构体变量作为函数参数
func test1(stu student) student {
	//fmt.Println(stu)
	stu.age = 16
	return stu
}
func main0101() {

	//结构体变量作为函数参数
	stu := student{101, "惠静", 18, "女", 100, "长达小区"}

	//结构体变量作为函数参数是值传递 形参不能改变实参的值
	stu = test1(stu)

	fmt.Println(stu)

}

func test2(stu []student) {
	//fmt.Println(stu)
	stu[2].name = "老衲"
}
func main() {
	stu := []student{student{101, "惠静", 18, "女", 100, "长达小区"},
		student{102, "刘菁", 18, "女", 100, "长达小区"},
		student{103, "师太", 18, "女", 100, "长达小区"}}
	//数组作为函数参数是值传递
	//切片作为函数参数是地址传递
	test2(stu)

	fmt.Println(stu)

}

func test(stu map[int]student) {
	//stu[101].name="段丁方"//err
	stu[103] = student{101, "段丁方", 18, "女", 100, "长达小区"}

	delete(stu, 102)
}
func main0103() {
	stu := make(map[int]student)
	stu[101] = student{101, "惠静", 18, "女", 100, "长达小区"}
	stu[102] = student{102, "刘菁", 18, "女", 100, "长达小区"}

	//map作为函数参数 地址传递
	test(stu)
	fmt.Println(stu)
}
