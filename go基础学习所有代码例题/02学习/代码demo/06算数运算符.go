package main

import "fmt"

func main0601() {
	a := 10
	b := 2
	//fmt.Println(a + b)
	//fmt.Println(a - b)
	//fmt.Println(a * b)
	fmt.Println(a / b) //0.5  整数相除得到的结果也是整数  0不能作为除数
}
func main0602() {
	a := 10
	b := 6
	c := a % b
	//取模运算符 除数不能为0
	//取模不能对浮点型进行操作
	fmt.Println(c)
}

func main0603() {
	//自增和自减
	a := 10
	//a=a+1
	//a++//自增  数本身+1操作
	a-- //a=a-1
	a-- //a=a-1
	//自增和自减不能出现在表达式中  引起程序的二义性
	a--
	//a=a++ - a-- * a++//err
	fmt.Println(a)
}

func main064() {
	a := 10
	b := 20
	c := 3.99
	//类型转换  格式 数据类型(变量) 数据类型(表达式)
	//sum:=float64(a)+float64(b)+c
	//sum:=float64(a+b)+c
	//将浮点型转成整型 值保留整数部分 忽略小数部分
	sum := a + b + int(c)
	fmt.Println(sum)
	fmt.Printf("%T", sum)
}

func main0604() {
	//int32 和int64都是整型 但是不允许类型转换
	//var a int32 = 10
	//var b int64 = 20
	//fmt.Println(int64(a) + b)

	//数据类型在转换时 如果高位转成低位可能会导致溢出 数据发生变化
	var a int8 = 127
	var b int16 = 250
	//fmt.Println(int8(b))溢出
	fmt.Println(int16(a) + b)

}
func main() {
	//107653是几天几小时几分几秒
	var s int = 107653
	m := s / 60
	h := m / 60
	d := h / 24

	//fmt.Println("天：", timer/60/60/24%365)
	//fmt.Println("时：", timer/60/60%24)
	//fmt.Println("分：", timer/60%60)
	//fmt.Println("秒：", timer%60)

	fmt.Println("天：", d % 365)
	fmt.Println("时：", h % 24)
	fmt.Println("分：", m % 60)
	fmt.Println("秒：", s % 60)

	//13+54*60+5*60*60+1*24*60*60=107653

}
