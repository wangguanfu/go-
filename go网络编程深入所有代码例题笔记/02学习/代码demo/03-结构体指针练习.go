package main

import "fmt"

type Test struct {
	name string
	flg bool
	age int
	interest []string
}

func initPerson (t *Test)  {
	t.age = 10
	t.name = "Nami"
	t.flg = true
	t.interest = append(t.interest, "shopping")
	t.interest = append(t.interest, "sleeping")
}


func initPerson2 () *Test {
	//t := new(Test)	// heap
	var t Test			// 局部变量 —— 随着栈帧的释放，消失（无意义）
	t.age = 10
	t.name = "Nami"
	t.flg = true
	t.interest = append(t.interest, "shopping")
	t.interest = append(t.interest, "sleeping")

	return &t
}

func main()  {
/*	var test Test	initPerson(&test)
	fmt.Println("test=", test)*/

	ret := initPerson2()

	fmt.Println("ret = ", ret)

}
