package main

import "fmt"

//在递归函数中一定有出口  return表示函数结束
func test5(a int) {

	if a == 0 {
		return
	}

	fmt.Println(a)
	test5(a - 1)

}

func main0601() {
	a := 10
	test5(a)
}

var sum = 1

func mlt(num int) {

	if num == 0 {
		return
	}
	mlt(num - 1)
	sum*=num
}
func main() {
	a := 6
	mlt(a)

	fmt.Println(sum)
}
