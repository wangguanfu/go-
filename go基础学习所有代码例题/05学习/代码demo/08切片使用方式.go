package main

import "fmt"

func main0801() {
	//数组
	var arr [10]int=[10]int{1,2,3,4,5,6,7,8,9,10}
	//切片
	//var slice []int=make([]int,10)
	var slice []int


	for i:=0;i<len(arr);i++{
		//slice[i]=arr[i]
		slice=append(slice,arr[i])
	}
	fmt.Println(slice)
}
func main(){
	//字符切片和字符串类型 可以进行转换

	////字符串
	//var str string="hello"
	////字符切片
	var ch []byte
	////类型转换
	//ch=[]byte(str)

	//var ch[]byte=[]byte{'h','e','l','l','o'}
	ch=[]byte("hello")

	//for i:=0;i<len(ch) ;i++  {
	//	fmt.Printf("%s",ch[i])
	//}

	fmt.Printf("%s",ch)


}