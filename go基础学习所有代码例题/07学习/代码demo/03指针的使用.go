package main

import (
	"fmt"
)

func main0301() {

	//定义指针如果为赋初始值默认值为nil
	//nil指向内存地址编号为0的空间
	//invalid memory address or nil pointer dereference
	//无效的内存地址或者是一个空指针的引用
	var p *int //空指针
	//fmt.Printf("%p\n",p)
	//fmt.Println(p)
	//指针变量指向了内存地址编号为0的空间  0-255为系统占用不允许读写操作
	//*p=100//err //写入操作
	fmt.Println(*p) //err 读操作

	//var slice []int//0x0
	//slice=append(slice,1,2,3)
	//var slice =make([]int,5)
	//var m map[int]string=make(map[int]string)
	//m[101]="hello"

}
func main0302() {
	//野指针   指针变量指向内存中一个未知的空间
	//从操作野指针对应的内存空间会报错
	//var p *int =*int(0xff00)
	var a int = 10
	var p *int = &a
	//野指针
	//p = 100
	*p = 100

	fmt.Println(*p)

}

func main() {
	a := 10
	var p *int
	var p1 *int
	//fmt.Println(p)
	//new  创建一块新的空间 将空间的地址赋值给一个指针变量
	p = new(int) //new 新建内存默认存储的值和（）内的类型是一致的
	p1 = p

	*p = 100
	p = &a
	fmt.Println(*p)
	fmt.Println(*p1)
}
