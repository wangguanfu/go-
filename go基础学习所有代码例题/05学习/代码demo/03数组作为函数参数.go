package main

import "fmt"

//数组作为函数参数
func test(arr [5]int) {

	//fmt.Println(arr)
	//sum:=0
	//for i := 0; i < len(arr); i++ {
	//	sum+=arr[i]
	//}
	//
	//fmt.Println(sum)
	fmt.Printf("%p\n", &arr)
	//arr[2]=123
	//fmt.Println(arr)

}
func main0301() {

	arr := [5]int{1, 2, 3, 4, 5}

	//值传递
	test(arr)
	fmt.Printf("%p\n", &arr)

}

//数组作为函数返回值
func BubbleSort(arr [10]int) [10]int {
	//冒泡排序
	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-1-i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	return arr

}
func main0302() {
	arr := [10]int{9, 1, 5, 6, 3, 7, 10, 8, 2, 4}
	//数组作为函数参数是值传递   形参不能改变实参的值
	//如果想改变实参的值  可以使用返回值赋值操作
	arr = BubbleSort(arr)

	fmt.Println(arr)

}

func StrLong(str [5]string){

	max:=len(str[0])
	i:=0
	for index,v :=range str{
		if len(v)>max{
			max=len(v)
			i=index
		}
	}

	fmt.Println(str[i])
}
func main(){
	str:=[5]string{ "马龙", "迈克尔乔丹", "雷吉米勒", "蒂姆邓肯", "科比布莱恩特" }
	StrLong(str)
}