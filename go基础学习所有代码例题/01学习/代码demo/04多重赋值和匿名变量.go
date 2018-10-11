package main

import "fmt"

func main0401() {
	//多重赋值
	//在多重赋值时变量和值的个数一一对应
	//在多重赋值时变量的类型可以不相同
	//a:=10
	//b:=20
	//c:=30

	//1、编辑时异常
	//2、编译时异常
	//3、运行时异常
	//a,b,c:=10,20,30
	a,b,c:=10,20.345,"你瞅啥"
	fmt.Println(a,b,c)
}


func main0402(){
	//_表示匿名变量 数据不接收
	a,_,c:=10,20,30

	fmt.Println(a,c)
}
func main(){
	a:=10
	b:=20

	//通过多重赋值
	a,b=b,a

	fmt.Println(a)
	fmt.Println(b)
}