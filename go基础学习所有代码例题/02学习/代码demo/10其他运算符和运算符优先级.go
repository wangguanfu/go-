package main

import "fmt"

func main1001() {
	//& 取地址运算符 fmt.Scan()
	a := 10
	p := &a //p是一个指针变
	fmt.Printf("%p\n", p)
	//* 取值运算符
	fmt.Println(*p)
	// >> << 位移运算符 ^  &  |  ~  位运算符
}

//运算符优先级

//括号 () 数组下标[] 结构体成员选择 .

//单目
//算数 ++ -- 逻辑 !  取地址 &  取值 *

//双目
//算数 * / %
//算数 + -
//比较 > < >= <= == !=
//逻辑 || &&
//赋值 = += -= *= /= %=

func main() {
	a := 10
	b := 20
	c := a > b && b > a
	fmt.Println(c)
}
