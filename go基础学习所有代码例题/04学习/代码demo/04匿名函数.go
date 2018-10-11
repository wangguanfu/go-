package main

import "fmt"

func main0401() {
	//var a int = 10
	////匿名函数
	//{
	//	var a int
	//	a = 123
	//	fmt.Println(a)
	//}
	//fmt.Println(a)
	var i int
	for i = 0; i < 10; i++ {

	}
	fmt.Println(i)
}
//将函数类型作为函数参数
func test2(f func(int,int)){
	f(10,20)
}
func main0402() {
	a := 10
	b := 20

	//匿名函数 在代码区进行存储
	f := func(a int, b int) {
		fmt.Println(a + b)
	}

	//fmt.Printf("%p\n",f)

	f(a, b)
	//f(1, 2)
	test2(f)

}
func main(){
	a:=10
	b:=20

	//匿名函数如果在{}后面有()表示函数调用
	value:=func(a int , b int)int{
		return a+b
	}

	//value:=func(a int , b int)int{
	//	return a+b
	//}(a,b)

	v:=value(a,b)

	fmt.Printf("%T\n",value)
	fmt.Printf("%T\n",v)
}