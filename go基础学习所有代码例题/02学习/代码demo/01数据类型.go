package main

import "fmt"

func main0101() {
	//数据类型

	//var 变量名 数据类型
	//bool

	var a bool //声明  默认值为 false
	//bool类型值只有两个 真 true 假 false
	a=true

	fmt.Println(a)
	//%t是一个占位符 表示输出一个bool类型的值
	fmt.Printf("%t",a)

}
func main0102(){
	//整型数据
	var a int
	//当赋值给int8时超出了取值范围会报错
	//与符号数据取值范围
	//-2^（8-1）~ 2^(8-1)-1
	a=140
	fmt.Println(a)

	//无符号整型数据定义
	//u -> unsigned
	// 0 ~ (2^8)-1
	//无符号只能存储0和正整数
	var b uint8
	//b=0
	fmt.Println(b)

}

func main0103(){
	//浮点型数据定义和使用
	//浮点型数据小数位数都是不精准的 float32 默认保留7位
	var a float32
	a=3.14//3.140000000124
	a=a*a
	fmt.Println(a)
	//%f是一个占位符 表示输出浮点数据  默认保留六位小数
	//%.2f小数点默认保留两位 会对第三位四舍五入
	//fmt.Printf("%.2f",a)


	//float64 小数位数默认保留15位 比float32更精准
	//var b float64 =3.14
	b:=3.14//float64
	b=b*b
	fmt.Println(b)
	fmt.Printf("%.2f",a)

}

func main(){
	//字符类型定义和使用
	//字符型
	//var a byte  //默认值为0
	//a:='0'
	//a='a'
	//a='0'//字符0


	//将小写字母转成大写字母 减32
	a:='a'
	a=a-32//允许参与整型计算
	//如果使用println打印数据时 显示byte对应的ASCII值
	//fmt.Println(a)
	//如果定义类型为byte  打印变量类型为uint8
	//如果通过自动推导类型创建变量 打印类型为int32
	fmt.Printf("%T",a)

	//fmt.Printf("%c",a)
}

func main0105(){
	//字符串类型
	//var a string  //默认值为 ""  "\0"  \0是字符串的结束标志 有系统自动添加
	//a="hello world"
	////字符串在打印时 打印\0之前内容
	//fmt.Println(a)
	//fmt.Printf("%s",a)

	var a string="相加"
	b:="操作"
	//字符串类型也允许相加操作  理解为字符串拼接
	a=a+b
	fmt.Println(a)
}