package main

import "fmt"

func main0501() {
	//fmt.Print 如果没有ln打印数据不会换行
	fmt.Print("打印数据")

	fmt.Println("赢取白富美")
	fmt.Print("整型数据")
}
func main0502(){
	a:=10
	//\n 是一个转义字符 表示换行
	fmt.Printf("%T\n",a)
	//%d是一个占位符 表示输出一个10进制整型数据
	fmt.Printf("%d\n",a)
	b:=3.595
	//%f是一个占位符 表示输出一个浮点型数据 默认小数位数保留6位
	fmt.Printf("%f\n",b)
	//%.2f表示数据保留两位小数  会对第三位小数四舍五入
	fmt.Printf("%.2f\n",b)

	c:="hello world"
	fmt.Printf("%T\n",c)//string  字符串类型
	//%s是一个占位符表示输出一个字符串类型数据
	fmt.Printf("%s\n",c)

	d:='a'
	fmt.Printf("%T\n",d)//byte 字符类型  int32  ASCII 对应整数的值
	//%c是一个占位符表示输出一个字符类型数据
	fmt.Printf("%c\n",d)

	e:=true//false true
	fmt.Printf("%T\n",e)//bool 布尔类型 用作于条件判断
	//%t是一个占位符表示输出一个bool类型数据
	fmt.Printf("%t\n",e)
}
func main(){
	name:="法师"
	age:=31
	sex:="男"

	fmt.Printf("大叫好，我叫%s，我今年%d岁了，我是%s生\n",name,age,sex)
}