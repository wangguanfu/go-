package main

import "fmt"

func main() {
	a := 10
	b := 20
	//比较运算符结果是bool类型
	// > 大于  < 小于 >=  大于等于 <= 小于等于 == 相等于 !=不等于
	//c := a < b
	c:=a == b
	fmt.Println(c)
}
