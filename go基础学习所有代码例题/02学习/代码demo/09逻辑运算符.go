package main

import "fmt"

func main0901() {
	//a:=10
	//b:=20
	////逻辑非 !   非真为假 非假为真
	//c:=!!(a>b)
	a := false
	b := !a

	fmt.Println(b)
}

func main0902() {
	a := 10
	b := 20
	c := 30

	//逻辑与  表达式1 &&  表达式2  同真为真 其余为假
	d := a > b && a < c

	fmt.Println(d)
}
func main() {

	a := 10
	b := 20
	c := 30
	//逻辑或 表达式1 || 表达式2  同假为假 其余为真
	d := a < b || a > c

	fmt.Println(d)
	//
}
