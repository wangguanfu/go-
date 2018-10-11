package main

import "fmt"

func main() {
	//水仙花数
	//一个三位数 各个位数的立方和等于本身
	for i := 100; i < 1000; i++ {
		//个位 十位  百位
		a := i % 10
		b := i % 100 / 10 //b:=i/10%10
		c := i / 100

		if a*a*a+b*b*b+c*c*c == i {
			fmt.Println(i)
		}
	}
}
