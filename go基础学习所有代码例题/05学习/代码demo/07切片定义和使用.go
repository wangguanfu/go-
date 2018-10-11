package main

import "fmt"

func main0701() {

	//数组定义和使用
	//var i int
	//fmt.Scan(&i)
	//a:=10
	//const a int = 10
	////数组定义时必须事先知道大小
	//var arr [a]int
	//fmt.Println(arr)

	//切片的定义
	//切片定义时 不需要写元素个数
	var slice []int
	//slice[0]=123
	//fmt.Println(slice)

	//append 为切片添加元素
	slice=append(slice,1,2,3)
	slice=append(slice,2,3,4)
	//修改切片中下标为1元素的值
	slice[1]=123
	fmt.Println(len(slice))
	fmt.Println(slice)
}

func main(){
	//var slice []int=[]int{1,2,3,4,5,6}
	//自动推导类型
	//slice:=[]int{1,2,3}
	//使用make创建切片长度  make(切片类型,切片长度)
	slice:=make([]int,3)
	slice[0]=1
	slice[1]=123
	slice[2]=321
	//slice[3]=2//切片下标越界//err

	slice=append(slice,1)
	slice=append(slice,123,234,345)


	//fmt.Println(len(slice))

	//for i:=0;i<len(slice);i++{
	//	fmt.Println(slice[i])
	//}
	for i,v:=range slice{
		fmt.Println(i,v)
	}
}
