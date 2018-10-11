package main

import (
	"fmt"
)

//定义接口
type Opter interface {
	GetResult() int
}

type Operate struct {
	num1 int
	num2 int
}
type Add struct {
	Operate
}

type Sub struct {
	Operate
}

func (a Add) GetResult() int {
	return a.num1 + a.num2
}

func (s Sub) GetResult() int {
	return s.num1 - s.num2
}

//多态  将接口作为函数参数
func GetResult(i Opter) (v int) {
	v = i.GetResult()
	return
}

type Factory struct {
}

func (fac *Factory) Calc(num1 int, num2 int, ch string) (v int) {

	switch ch {

	case "+":
		var a Add = Add{Operate{num1, num2}}
		v = GetResult(&a)
	case "-":
		var s Sub = Sub{Operate{num1, num2}}
		v = GetResult(&s)
	}
	return
}

func main(){

	var f Factory


	v:=f.Calc(10,20,"-")
	fmt.Println(v)
}


func main0201() {
	var a Add = Add{Operate{10, 20}}
	v := GetResult(&a)
	//var i Opter
	//i = &a
	//v := i.GetResult()
	////v := a.GetResult()
	fmt.Println(v)

	var s Sub = Sub{Operate{10, 20}}
	v = GetResult(&s)
	//i = &s
	//v = i.GetResult()
	////v := s.GetResult()
	fmt.Println(v)
}
