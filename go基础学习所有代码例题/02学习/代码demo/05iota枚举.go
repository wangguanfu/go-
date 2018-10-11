package main

import "fmt"

func main0501() {

	//枚举定义和使用
	const (
		//在定义枚举时默认值为0 每换一行值增加1
		//在定义枚举是可以不动写  =iota
		a = iota //0
		b = iota //1
		c = iota //2

		d
		e
	)
	//a=100//err
	fmt.Println(a, b, c, d, e)
}
func main0502() {
	const (
		//iota枚举里面数据可以赋值  不建议赋值
		//枚举里面数据都是整型数据
		a = iota
		b = 10
		c = iota
		d
		e
	)
	fmt.Println(a, b, c, d, e)
}
func main() {
	const (
		//枚举值里面数据如果不换行 一行内的值相同
		//枚举定义时如果没有赋值 会重复上一行格式
		a    = iota
		b, c = iota, iota
		d, e
	)
	fmt.Println(a, b, c, d, e)
}
