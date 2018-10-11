package main

import "fmt"

func main() {

	a := 10
	b := 20
	c := 30
	//指针数组
	//每一个元素都是指针类型的  可以存储多个变量的地址
	var arr [3]*int
	arr[0] = &a
	arr[2] = &b
	arr[1] = &c

	fmt.Println(arr)
	//
	//fmt.Println(&a)
	//fmt.Println(&b)
	//fmt.Println(&c)

	//通过指针数组间接修改变量的值
	*arr[1]=222
	*arr[0]=123

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

	//for i:=0;i<3;i++{
	//	fmt.Printf("%p\n",&arr[i])
	//}
	//遍历指针数组
	//for i:=0;i<len(arr);i++{
	//	fmt.Println(*arr[i])
	//}

	for i,v:=range arr {
		//fmt.Printf("%T\n",v)
		fmt.Println(i,*v)
	}

}
