package main

import "fmt"

func main0601() {
	var score int
	//通过键盘为score赋值  & 取地址运算符
	//fmt.Scan(&score)
	fmt.Scanf("%d", &score)
	fmt.Println(score)
}
func main() {
	//通过键盘输入三门成绩  计算总成绩和平均成绩
	var c, m, e int

	//通过键盘为三门成绩赋值  数据间隔空格 或者回车
	//fmt.Scan(&c, &m, &e)
	fmt.Scanf("%d%d%d", &c, &m, &e)
	//计算总成绩
	sum := c + m + e

	fmt.Println("总成绩：", sum)
	fmt.Println("平均成绩： ", sum/3) //整型数据相除 得到的结果是整型数据



}
