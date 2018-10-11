package main

import "fmt"

func main1101() {
	slice :=[]int{1,2,3,4,5}
	//[起始下标:结束下标]  不包含结束下标
	//切片的截取 截取后的切片也是原始切片的一部分
	//s:=slice[2:5]

	s:=slice[2:5]

	s[0]=123
	fmt.Printf("%p\n",s)
	fmt.Printf("%p\n",slice)
	//fmt.Println(s)
	fmt.Println(s)
	fmt.Println(slice)

}
func main(){
	slice :=[]int{1,2,3,4,5}
	var s []int=make([]int,10,10)
	//在使用copy进行拷贝时需要注意使用有足够的长度
	//使用copy 会在内存中创建另外一个独立的切片空间 两个空间不会相互影响
	//copy(s,slice)
	//copy(s,slice[0:2])
	//[起始下标:]从起始下标拷贝到长度结束
	//copy(s,slice[2:])
	//[:结束下标]从下标0开始拷贝到结束下标之前
	copy(s,slice[:4])
	//s[2]=123

	fmt.Println(s)
	fmt.Println(slice)
	fmt.Printf("%p\n",s)
	fmt.Printf("%p\n",slice)

	//slice :=[]int{1,2,3,4,5}
	////如果切片直接赋值 两个切片指向同一内存地址空间
	//s:=slice
	//
	//s[2]=123
	//fmt.Println(s)
	//fmt.Println(slice)
	//fmt.Printf("%p\n",s)
	//fmt.Printf("%p\n",slice)

}