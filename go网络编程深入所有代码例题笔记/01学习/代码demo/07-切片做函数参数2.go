package main

import "fmt"


// 切片做函数参数，传值
func testFunc2(s []int) []int { // 切片做函数参数
	s = append(s, 777)
	fmt.Println("testFunc2: s=", s)

	return s
}

func main() {
	slice := []int{0, 1, 2}

	fmt.Println("main before: s=", slice)

	aferData := testFunc2(slice)
	fmt.Println("main after: s=", slice)
	fmt.Println("main after: aferData=", aferData)
}
