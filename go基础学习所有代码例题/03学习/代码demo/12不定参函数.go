package main

import "fmt"

//在整个项目中函数名是唯一的
//...不定参  a相当于一组数据的集合

func test1(a ...int) {
	//len(a) 计算集合数据个数
	//for i:=0;i<len(a);i++{
	//	//a[下标]找到具体元素
	//	fmt.Println(a[i])
	//}
	fmt.Println(a)
}
func main1201() {

	test1(1, 2, 3, 4, 5, 6, 7, 8)
	test1(6, 6, 6)
	//在不定参调用中可以不传值
	test1()
	//test1 函数指针  函数对应的代码区地址
}

func test2(arr ...int) {
	sum := 0
	//for i := 0; i < len(arr); i++ {
	//	sum += arr[i]
	//}
	//i index 下标  v value 值
	for _, v := range arr {
		sum += v
		//fmt.Println(v)
	}
	fmt.Println(sum)
}
func main() {
	test2(1, 2, 3)
	test2(4, 5, 6, 7, 8, 9, 10)
}
