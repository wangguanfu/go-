package main

import (
	"fmt"
	"unsafe"
)

// 定义函数指针【类型】FuncP, 该类型的指针，指向一类函数，有两个参数（int/bool) ，一个int 返回值的函数。
type FuncP func(x int, y bool) int
// 函数名，是函数第一条指定的地址。函数名——>地址（指针）
func addOne(x int, y bool) int {
	if y == true {
		x = x+1
	}
	return x
}
func subTen(x int, y bool) int {
	if y == true {
		x -= 10
	}
	return x
}
func useCallback(x int, y bool, p FuncP) int {

	return p(x, y)
}
func main()  {
	var p FuncP			// 空指针
	p = addOne
	fmt.Printf("type:%T, size:%d\n", p, unsafe.Sizeof(p))

	fmt.Println(useCallback(5, true, p))

	fmt.Println(useCallback(5, true, subTen))

	//addOne(60, false)  普通函数调用方法。
}
