package main

import "fmt"

func main0701() {

	//var  变量名  数据类型

	//数组定义和使用
	//var 数组名 [元素个数] 数据类型
	//var score [8200000]int
	var arr [10]int //数组声明后默认的值为0
	//可以通过 数组名[下标] 访问具体的元素
	//数组下标是从0开始的到数组元素最大个数-1为最大下标
	//为数组元素赋值
	arr[0] = 10
	arr[1] = 20
	arr[2] = 30

	//打印数组元素

	//fmt.Println(arr[0])
	//fmt.Println(arr[1])
	//fmt.Println(arr[2])
	//fmt.Println(arr[3])
	//fmt.Println(arr[4])
	//fmt.Println(arr[5])
	//fmt.Println(arr[6])
	//fmt.Println(arr[7])
	//数组元素个数为10个
	for i := 0; i < 10; i++ {
		fmt.Println(arr[i])
	}
}
func main0702() {

	//数组初始化
	//var arr [10]int=[10]int{1,2,3,4,5,6,7,8,9,10}
	//var arr [10]int=[10]int{1,2,3,4,5}
	var arr [10]int = [10]int{1: 123, 3: 456, 6: 789}

	//for i := 0; i < 10; i++ {
	//	fmt.Println(arr[i])
	//}

	//计算数组元素个数
	//len(数组名)计算数组元素个数
	fmt.Println(len(arr))

	//for i := 0; i < len(arr); i++ {
	//	fmt.Println(arr[i])
	//}

	//i 表示数组下标  v表示数组元素的值
	//range 数组名  遍历数组的值
	for i, v := range arr {
		fmt.Println(i, v)
	}

}
func main0703() {
	var arr [5]int = [5]int{1, 2, 3, 4}
	var arr1 [5]int

	//数组下标越界
	//arr[5]=5//err
	//arr[-1]=123//err

	//数组名 表示数组的首地址  不能将值赋值给数组名
	//arr=1
	//数组如果类型相同 可以赋值
	arr1 = arr

	fmt.Println(arr1)

}

func main() {
	var arr [5]int = [5]int{1, 2, 3, 4, 5}

	//数组元素在内存中连续存储的
	for i := 0; i < len(arr); i++ {
		fmt.Printf("%p\n", &arr[i])
	}
	//数组名代表数组元素的首地址
	fmt.Printf("%p\n",&arr)
}
