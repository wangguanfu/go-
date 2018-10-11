package main

import "fmt"

func main() {
	//计算1-100之间 不包含7（个位 十位） 或者7的倍数

	for i := 1; i <= 100; i++ {

		if i%7 == 0 || i%10 == 7 || i/10 == 7 {
			fmt.Println("敲桌子")
		}else{
			fmt.Println(i)
		}

	}
}
