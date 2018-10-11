//注释  行注释 可以注释一行

/*
快注释
可以注释多行
 */
 //程序所属包
package main

//导入系统fmt包 format 格式化  导入格式化输入输出包
import "fmt"

//func 函数格式  main  函数名  程序有且只有一个主函数（main）main 是程序的住入口
//()括起来的成为函数的参数列表 可以是0个 也可以是多个 （变量名 数据类型, 变量名 数据类型）{}代码体 程序体 函数体
func main01(){
	//在代码中 一条语句可以写在多行  但是多条语句不能写在一行
	//打印输出 fmt.Println  ""引起来的成为字符串
	fmt.Println("hello world!")
	fmt.Println("字符串")
	fmt.Println("你瞅啥")

}


