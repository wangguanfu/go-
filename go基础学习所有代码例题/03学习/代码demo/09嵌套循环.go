package main

import "fmt"

func main0901() {

	//外层执行一次  内层执行一周
	for i := 0; i < 5; i++ {

		for j := 0; j < 5; j++ {
			fmt.Println(i, j)
		}
	}
}

//九九乘法口诀
// 1*1=1
// 1*2=2 2*2=4
// 1*3=3 2*3=6 3*3=9
// 1*4=4 2*4=8 3*4=12 4*4=16

func main() {

	//外层控制行 内层控制列
	for i := 1; i <= 9; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d*%d=%d\t", j, i, i*j)
		}
		fmt.Println()
	}
}
