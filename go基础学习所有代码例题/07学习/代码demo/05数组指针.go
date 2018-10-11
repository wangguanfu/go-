package main

import "fmt"

func main0501() {

	//数组
	var arr [10]int = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}

	//定义指针指向数组
	//p := &arr
	//定义指针是要和数组元素个数数据类型保持一致
	var p *[10]int = &arr

	//数组首地址
	//fmt.Printf("%p\n",p)
	//fmt.Printf("%p\n",&arr)
	//fmt.Printf("%p\n",&arr[0])

	//fmt.Printf("%T\n", p)
	//通过指针间接修改数组元素的值
	//(*p)[0]=123
	//(*p)[1]=222
	//数组指针加下标就可以直接修改数组元素的值
	//p[0]=123
	//p[1]=222
	//
	//fmt.Println(arr)

	//让数组指针和数组建立关系
	//for i := 0; i < len(p); i++ {
	//	fmt.Println(p[i])
	//}

	for i, v := range p {
		fmt.Println(i, v)
	}

}

//数组指针作为函数参数
func test5(p *[10]int) {

	//fmt.Println(p)
	//修改数组元素的值
	p[0]=123
	p[1]=222
}

func main() {
	//数组
	var arr [10]int = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}

	//数组指针作为函数参数是地址传递
	test5(&arr)

	fmt.Println(arr)
}
