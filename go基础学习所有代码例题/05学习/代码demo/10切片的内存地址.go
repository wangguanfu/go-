package main

import "fmt"

func main1001() {

	//如果定义切片没有长度  内存地址编号为0
	//切片是在内存中连续存储的
	//var slice [] int
	var slice []int=make([]int,1)
	fmt.Printf("%p\n",slice)
	//通过append添加数据时可能会改变切片的地址
	slice=append(slice,1)//slice[0]=0 slice[1]=1
	fmt.Printf("%p\n",slice)


}

func main(){
	//make(切片类型，长度，容量)
	//超出容量 可以进行扩充
	var slice []int =make([]int,3,10)
	//len(切片名)计算已经使用的长度
	fmt.Println(len(slice))
	//cap(切片名)计算切片容量
	fmt.Println(cap(slice))
	//在打印地址是 切片名本身就是一个地址
	fmt.Printf("%p\n",slice)

	slice=append(slice,1,2,3,4)

	//len(切片名)计算已经使用的长度
	fmt.Println(len(slice))
	//cap(切片名)计算切片容量
	fmt.Println(cap(slice))
	fmt.Printf("%p\n",slice)

	//超出容量 程序不会报错  容量和长度会自动扩充
	//容量扩充一般是上一次的2倍 如果超过1024字节 是上一次的1/4
	slice=append(slice,1,2,3,4)

	//len(切片名)计算已经使用的长度
	fmt.Println(len(slice))
	//cap(切片名)计算切片容量
	fmt.Println(cap(slice))
	fmt.Printf("%p\n",slice)

	slice=append(slice,1,2,3,4,1,2,3,4,5,6)

	//len(切片名)计算已经使用的长度
	fmt.Println(len(slice))
	//cap(切片名)计算切片容量
	fmt.Println(cap(slice))
	fmt.Printf("%p\n",slice)

	fmt.Println(slice)
}
