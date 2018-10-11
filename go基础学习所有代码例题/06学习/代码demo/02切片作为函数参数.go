package main

import "fmt"

//切片作为函数参数
func test(s []int) {

	//for i:=0;i<len(s);i++{
	//	fmt.Println(s[i])
	//}

	fmt.Printf("%p\n", s)

	s[2] = 123
}

func main0201() {

	s := []int{1, 2, 3, 4, 5, 6}
	//切片作为函数参数是地址传递  形参可以改变实参的值
	test(s)
	fmt.Printf("%p\n", s)
	fmt.Println(s)
}

func BubbleSort(s []int) {

	//切片排序
	for i := 0; i < len(s)-1; i++ {
		for j := 0; j < len(s)-1-i; j++ {
			if s[j] > s[j+1] {
				s[j], s[j+1] = s[j+1], s[j]
			}
		}
	}

}

func main0202() {
	s := []int{9, 1, 5, 6, 10, 2, 8, 4, 7, 3}
	//切片作为函数参数 地址传递 形参可以改变实参值
	BubbleSort(s)

	fmt.Println(s)
}

//切片作为函数的返回值
func test1(s []int) []int {

	//append 添加切片元素 会改变切片的地址
	s = append(s, 1, 2, 3, 4, 5)
	fmt.Println(s)
	return s

}
func main() {

	s := []int{1, 2, 3, 4, 5, 6}
	s = test1(s)

	fmt.Println(s)

}
