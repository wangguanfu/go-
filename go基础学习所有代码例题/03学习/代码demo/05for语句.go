package main

import "fmt"

func main() {

	//for 表达式1;表达式2;表达式3{代码}
	//var i int = 0
	//for i:=0; i < 10; i++{
	//	fmt.Println("表达式1")
	//
	//}

	//计算1-100和
	//sum := 0
	//for i := 1; i <= 100; i++ {
	//	sum += i
	//}
	//
	//fmt.Println(sum)

	//计算1-100 偶数的和
	//for 循环可以嵌套if语句
	//sum := 0
	//for i := 1; i <= 100; i++ {
	//
	//	if i%2 == 0 {
	//		sum += i
	//	}
	//}
	//fmt.Println(sum)

	sum:=0
	for i:=0;i<=100;i+=2{
		sum+=i
	}
	fmt.Println(sum)

}
