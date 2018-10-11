package main

import "fmt"


// 该示例—— 切片做函数参数，传地址（引用)
func testFunc(s []int) { // 切片做函数参数
	s[0] = -1 // 直接修改 main中的 slice
}

func main() {
	slice := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	fmt.Println(slice)

	testFunc(slice)

	fmt.Println(slice)
}
