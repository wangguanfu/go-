package main

import "fmt"

//func test(a int, b int) int {
//	return a + b
//}

//type  为已存在数据类型起别名

//系统规定不能为基本数据类型绑定方法
type Int int


//func (方法接收者 数据类型) 方法名(方法参数列表) 返回值列表{代码体}
func (a Int) add(b Int) Int {

	return a + b
}

func main() {

	//a := 10
	//b := 20
	//v := test(a, b)

	//fmt.Println(v)



	var a Int = 10
	var b Int = 20
	//对象.方法
	//c:=add(b)//err  方法调用必须有接收者
	//c := a.add(b)
	//方法的绑定是根据类型相关 跟变量无关
	c:=b.add(a)

	fmt.Println(c)
	//c := a + Int(b)
	//
	//fmt.Println(c)
	//fmt.Println(a)

	//fmt.Printf("%T\n", a)
}
