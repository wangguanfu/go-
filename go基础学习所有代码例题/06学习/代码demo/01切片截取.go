package main

import "fmt"

func main() {
	slice:=[]int{1,2,3,4,5}

	//s:=slice[0:2]
	//s:=slice[2:]
	//s:=slice[:5]

	//切片全部截取
	//s:=slice[:]
	//s:=slice
	//[起始下标:结束下标:容量]
	//容量不能大于源切片容量  不能小于截取后长度
	s:=slice[0:3:5]
	fmt.Println(s)

}
