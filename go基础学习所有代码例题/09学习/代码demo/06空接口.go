package main

import "fmt"

func main0601() {

	//定义接口类型 可以存储各个类型的数据
	var i interface{}

	i = 0
	fmt.Println(i)

	i = 3.14
	fmt.Println(i)

	i = "接口类型"
	fmt.Println(i)

	i = []int{1, 2, 3, 4}
	fmt.Println(i)

	//i=student{person{"张三","男",12},100}
	//fmt.Println(i)

}

func main0602() {
	//接口类型默认为nil （void*  万能指针）
	//会根据存储的数据变化 类型随之变化
	var i interface{}
	fmt.Printf("%T\n", i)

	i = 0
	fmt.Printf("%T\n", i)

	i = 3.14
	fmt.Printf("%T\n", i)

	i = []int{1, 2, 3, 4}
	fmt.Printf("%T\n", i)

}

func main() {
	var i []interface{}

	i = append(i, 0, 3.14, "接口类型", [3]int{1, 2, 3}, map[int]string{101: "张三丰", 102: "张君宝"})

	//fmt.Println(i)

	//for j := 0; j < len(i); j++ {
	//	fmt.Println(i[j])
	//}

	//for idx,v:=range i{
	//	fmt.Println(idx,v)
	//}

	fmt.Printf("%T\n", i)
}
